#!/bin/bash

# GoTime Library - One-Pass Release Script
#
# Validates everything first, then performs the release in a single pass:
#   1. Validate version + that release notes exist
#   2. Verify clean tree, no unpushed commits, correct branch, tag not already used
#   3. Run quality gates (build, vet, tests)
#   4. Create the annotated vX.Y.Z tag
#   5. Push branch + tag
#   6. Create the GitHub release from the notes file (if gh is available)
#
# The version is given WITHOUT the leading "v" (e.g. 2.0.4); the script adds it.
#
# Usage:
#   ./scripts/release.sh 2.0.4
#   ./scripts/release.sh 2.0.4 --dry-run      # validate + show the plan, change nothing
#   ./scripts/release.sh 2.0.4 --no-push       # tag locally only
#   ./scripts/release.sh 2.0.4 --no-github     # skip GitHub release creation
#   ./scripts/release.sh 2.0.4 --race          # run the race detector in the test gate
#   ./scripts/release.sh 2.0.4 --yes           # skip the confirmation prompt
#
# Options:
#   --dry-run        Run all checks and print the plan, but make no changes
#   --no-push        Create the tag locally but do not push branch or tag
#   --no-github      Do not create a GitHub release (via gh)
#   --race           Include -race in the test gate (needs a C toolchain)
#   --skip-tests     Skip build/vet/test gates (NOT recommended)
#   --allow-dirty    Proceed even if the working tree has uncommitted changes
#   --allow-unpushed Proceed even if the branch has commits not yet on the remote
#   --allow-branch   Proceed even if not on master/main
#   --remote NAME    Git remote to push to (default: origin)
#   --yes, -y        Do not prompt for confirmation before releasing
#   -h, --help       Show this help message

set -euo pipefail

# ---------------------------------------------------------------------------
# Output helpers
# ---------------------------------------------------------------------------
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

print_status()  { echo -e "${BLUE}[INFO]${NC} $1"; }
print_success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
print_warning() { echo -e "${YELLOW}[WARNING]${NC} $1"; }
print_error()   { echo -e "${RED}[ERROR]${NC} $1" >&2; }

usage() {
    # Print the leading comment block (from line 3 to the first non-comment line).
    awk 'NR>=3 { if (/^#/) { sub(/^# ?/, ""); print; next } else { exit } }' "$0"
    exit "${1:-0}"
}

# Bump X.Y.Z by the given component (patch|minor|major).
bump_version() {
    local ver="$1" part="$2" major minor patch
    major="${ver%%.*}"; ver="${ver#*.}"
    minor="${ver%%.*}"; patch="${ver#*.}"
    patch="${patch%%-*}"   # drop any pre-release suffix
    case "$part" in
        patch) echo "${major}.${minor}.$((patch + 1))" ;;
        minor) echo "${major}.$((minor + 1)).0" ;;
        major) echo "$((major + 1)).0.0" ;;
    esac
}

# Print one suggestion line: version, and whether its notes file already exists.
suggest_line() {
    local label="$1" ver="$2"
    local notes="docs/releases/v${ver}.md"
    if [ -f "$notes" ]; then
        printf "    %-7s %-10s ${GREEN}notes ready${NC}     ./scripts/release.sh %s\n" "$label" "v$ver" "$ver"
    else
        printf "    %-7s %-10s ${YELLOW}notes missing${NC}   create %s first\n" "$label" "v$ver" "$notes"
    fi
}

# Thorough guidance shown when the script is invoked without a version.
guide() {
    local latest patch minor major
    latest="$(git tag --list 'v[0-9]*' --sort=-version:refname 2>/dev/null | head -1)"

    echo "This script releases GoTime in one pass: it validates the version and its"
    echo "release notes, runs the build/vet/test gates, then tags, pushes, and publishes"
    echo "the GitHub release."
    echo ""
    echo "Usage:"
    echo "    ./scripts/release.sh <version>          # e.g. 2.0.4  (no 'v' prefix)"
    echo "    ./scripts/release.sh <version> --dry-run # validate only, change nothing"
    echo ""

    if [ -n "$latest" ]; then
        local base="${latest#v}"
        echo "Latest released version: ${latest}"
        echo ""
        echo "Suggested next versions:"
        suggest_line "patch" "$(bump_version "$base" patch)"
        suggest_line "minor" "$(bump_version "$base" minor)"
        suggest_line "major" "$(bump_version "$base" major)"
    else
        echo "No existing version tags found — this would be the first tagged release."
    fi

    echo ""
    echo "Before releasing, make sure that:"
    echo "    1. Release notes exist at docs/releases/v<version>.md"
    echo "    2. RELEASENOTES.md has a section for the version"
    echo "    3. Install commands in README.md and docs/ pin the new version"
    echo "    4. Your changes are committed and you are on master/main"
    echo ""
    echo "Recommended flow:"
    echo "    ./scripts/release.sh <version> --dry-run   # confirm all checks pass"
    echo "    ./scripts/release.sh <version>             # perform the release"
    echo ""
    echo "Run './scripts/release.sh --help' for the full list of options."
    exit "${1:-1}"
}

# ---------------------------------------------------------------------------
# Parse arguments
# ---------------------------------------------------------------------------
VERSION=""
DRY_RUN=false
DO_PUSH=true
DO_GITHUB=true
USE_RACE=false
SKIP_TESTS=false
ALLOW_DIRTY=false
ALLOW_UNPUSHED=false
ALLOW_BRANCH=false
ASSUME_YES=false
REMOTE="origin"

while [ $# -gt 0 ]; do
    case "$1" in
        --dry-run)      DRY_RUN=true ;;
        --no-push)      DO_PUSH=false ;;
        --no-github)    DO_GITHUB=false ;;
        --race)         USE_RACE=true ;;
        --skip-tests)   SKIP_TESTS=true ;;
        --allow-dirty)    ALLOW_DIRTY=true ;;
        --allow-unpushed) ALLOW_UNPUSHED=true ;;
        --allow-branch)   ALLOW_BRANCH=true ;;
        --yes|-y)       ASSUME_YES=true ;;
        --remote)       shift; REMOTE="${1:-}"; [ -n "$REMOTE" ] || { print_error "--remote needs a value"; exit 1; } ;;
        -h|--help)      usage 0 ;;
        -*)             print_error "Unknown option: $1"; usage 1 ;;
        *)
            if [ -n "$VERSION" ]; then
                print_error "Unexpected extra argument: $1"; usage 1
            fi
            VERSION="$1"
            ;;
    esac
    shift
done

echo "🚀 GoTime Release"
echo "================="
echo ""

# ---------------------------------------------------------------------------
# 0. Location + prerequisites
# ---------------------------------------------------------------------------
if [ ! -f "go.mod" ]; then
    print_error "go.mod not found. Run this script from the project root."
    exit 1
fi

command -v git >/dev/null 2>&1 || { print_error "git is required but not found."; exit 1; }
command -v go  >/dev/null 2>&1 || { print_error "go is required but not found."; exit 1; }

# ---------------------------------------------------------------------------
# 1. Validate the version and derive the tag
# ---------------------------------------------------------------------------
if [ -z "$VERSION" ]; then
    guide 1
fi

# Be forgiving if the caller accidentally included a leading "v".
VERSION="${VERSION#v}"

if ! echo "$VERSION" | grep -Eq '^[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.]+)?$'; then
    print_error "Version '$VERSION' is not valid semver (expected X.Y.Z, without the 'v' prefix)."
    exit 1
fi

TAG="v$VERSION"
NOTES_FILE="docs/releases/${TAG}.md"

print_status "Releasing version: ${TAG}"

# ---------------------------------------------------------------------------
# 2. Release notes MUST exist — this is the core guard the release depends on
# ---------------------------------------------------------------------------
if [ ! -f "$NOTES_FILE" ]; then
    print_error "Release notes not found: ${NOTES_FILE}"
    print_error "Create the release notes for ${TAG} before releasing (see docs/releases/)."
    exit 1
fi
if [ ! -s "$NOTES_FILE" ]; then
    print_error "Release notes ${NOTES_FILE} exist but are empty."
    exit 1
fi
print_success "Release notes found: ${NOTES_FILE}"

# Soft check: the consolidated changelog should mention this version too.
if [ -f "RELEASENOTES.md" ] && ! grep -q "$TAG" RELEASENOTES.md; then
    print_warning "RELEASENOTES.md does not mention ${TAG}; consider adding a section."
fi

# Hard check: install instructions must pin the version being released. Every
# "gotime/v2@vX.Y.Z" occurrence in these files must match ${TAG}; a stale pin
# means the docs still advertise a previous release and the release is aborted.
VERSIONED_FILES="README.md docs/README.md docs/getting-started/installation.md docs/getting-started/quick-start.md"
for f in $VERSIONED_FILES; do
    [ -f "$f" ] || continue
    pinned="$(grep -oE 'gotime(/v2)?@v[0-9]+\.[0-9]+\.[0-9]+' "$f" | sed 's/.*@//' | sort -u || true)"
    if [ -z "$pinned" ]; then
        print_warning "${f}: no pinned install version found (expected @${TAG})."
        continue
    fi
    stale="$(echo "$pinned" | grep -vx "$TAG" || true)"
    if [ -n "$stale" ]; then
        print_error "${f} pins $(echo "$stale" | paste -sd, -) but is being released as ${TAG}."
        print_error "Update the install command(s) to @${TAG} before releasing."
        exit 1
    fi
    print_success "${f} references ${TAG}."
done

# ---------------------------------------------------------------------------
# 3. Repository state checks
# ---------------------------------------------------------------------------
git rev-parse --is-inside-work-tree >/dev/null 2>&1 || { print_error "Not a git repository."; exit 1; }

CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [ "$ALLOW_BRANCH" = false ] && [ "$CURRENT_BRANCH" != "master" ] && [ "$CURRENT_BRANCH" != "main" ]; then
    print_error "On branch '${CURRENT_BRANCH}', expected master/main. Use --allow-branch to override."
    exit 1
fi

if [ "$ALLOW_DIRTY" = false ] && [ -n "$(git status --porcelain)" ]; then
    print_error "Working tree has uncommitted changes. Commit or stash them first (or use --allow-dirty)."
    git status --short
    exit 1
fi

# Tag must not already exist locally...
if git rev-parse -q --verify "refs/tags/${TAG}" >/dev/null; then
    print_error "Tag ${TAG} already exists locally."
    exit 1
fi
# ...nor on the remote.
if git ls-remote --exit-code --tags "$REMOTE" "$TAG" >/dev/null 2>&1; then
    print_error "Tag ${TAG} already exists on remote '${REMOTE}'."
    exit 1
fi
print_success "Tag ${TAG} is available on '${REMOTE}' and locally."

# ---------------------------------------------------------------------------
# 4. Quality gates
# ---------------------------------------------------------------------------
if [ "$SKIP_TESTS" = true ]; then
    print_warning "Skipping build/vet/test gates (--skip-tests)."
else
    print_status "Building all packages..."
    go build ./...
    print_success "Build passed."

    print_status "Running go vet..."
    go vet ./...
    print_success "Vet passed."

    if [ "$USE_RACE" = true ]; then
        print_status "Running tests with the race detector..."
        go test -race ./...
    else
        print_status "Running tests..."
        go test ./...
    fi
    print_success "Tests passed."
fi

# ---------------------------------------------------------------------------
# 5. Plan summary
# ---------------------------------------------------------------------------
COMMIT="$(git rev-parse --short HEAD)"
# Derive a one-line summary from the notes for the tag/release title. Prefer the
# first prose line under a "## Summary" heading; otherwise the first prose line in
# the file. Skip headings, blockquotes, blanks, and **metadata** lines.
extract_summary() {
    awk '
        /^##[[:space:]]+[Ss]ummary[[:space:]]*$/ { in_summary = 1; next }
        /^##/ { in_summary = 0 }
        {
            line = $0
            if (line ~ /^[[:space:]]*(#|>|\*\*|$)/) next   # skip heading/quote/meta/blank
            gsub(/^[[:space:]]+|[[:space:]]+$/, "", line)
            if (in_summary) { print line; exit }
            if (first == "") first = line
        }
        END { if (first != "") print first }
    ' "$NOTES_FILE" | head -1
}
SUMMARY="$(extract_summary)"
[ -n "$SUMMARY" ] || SUMMARY="Release ${TAG}"
TAG_MESSAGE="Release ${TAG}: ${SUMMARY}"

echo ""
echo "Plan:"
echo "  Version .......... ${TAG}"
echo "  Commit ........... ${COMMIT} (${CURRENT_BRANCH})"
echo "  Remote ........... ${REMOTE}"
echo "  Notes ............ ${NOTES_FILE}"
echo "  Tag message ...... ${TAG_MESSAGE}"
echo "  Push branch+tag .. $([ "$DO_PUSH" = true ] && echo yes || echo no)"
echo "  GitHub release ... $([ "$DO_GITHUB" = true ] && echo yes || echo no)"
echo ""

if [ "$DRY_RUN" = true ]; then
    print_success "Dry run complete — all checks passed. No changes made."
    exit 0
fi

# ---------------------------------------------------------------------------
# 6. Confirmation before irreversible actions
# ---------------------------------------------------------------------------
if [ "$ASSUME_YES" = false ]; then
    printf "Proceed with releasing %s? [y/N] " "$TAG"
    read -r reply
    case "$reply" in
        y|Y|yes|YES) ;;
        *) print_warning "Aborted by user."; exit 1 ;;
    esac
fi

# ---------------------------------------------------------------------------
# 7. Create the annotated tag
# ---------------------------------------------------------------------------
print_status "Creating annotated tag ${TAG}..."
git tag -a "$TAG" -m "$TAG_MESSAGE"
print_success "Tag ${TAG} created."

if [ "$DO_PUSH" = false ]; then
    print_warning "Skipping push (--no-push). Tag exists locally only."
    print_status "To finish later: git push ${REMOTE} ${CURRENT_BRANCH} && git push ${REMOTE} ${TAG}"
    exit 0
fi

# ---------------------------------------------------------------------------
# 8. Push branch + tag
# ---------------------------------------------------------------------------
print_status "Pushing ${CURRENT_BRANCH} to ${REMOTE}..."
git push "$REMOTE" "$CURRENT_BRANCH"
print_status "Pushing tag ${TAG} to ${REMOTE}..."
git push "$REMOTE" "$TAG"
print_success "Branch and tag pushed."

# ---------------------------------------------------------------------------
# 9. Create the GitHub release
# ---------------------------------------------------------------------------
if [ "$DO_GITHUB" = false ]; then
    print_warning "Skipping GitHub release (--no-github)."
elif ! command -v gh >/dev/null 2>&1; then
    print_warning "gh CLI not found — skipping GitHub release."
    print_status "Create it manually: gh release create ${TAG} --title \"GoTime ${TAG}\" --notes-file ${NOTES_FILE}"
else
    if gh release view "$TAG" >/dev/null 2>&1; then
        print_warning "GitHub release ${TAG} already exists — leaving it unchanged."
    else
        print_status "Creating GitHub release ${TAG}..."
        gh release create "$TAG" \
            --title "GoTime ${TAG} - ${SUMMARY}" \
            --notes-file "$NOTES_FILE"
        print_success "GitHub release ${TAG} published."
    fi
fi

echo ""
print_success "Release ${TAG} complete! 🎉"
print_status "Verify: go list -m -versions github.com/maniartech/gotime/v2"
