# GoTime Release Process

This document describes the step-by-step process for releasing a new version of GoTime. It is designed to be repeatable for all future releases (major, minor, or patch).

## 1. Pre-Release Checklist

- Ensure you are on the correct branch (usually `master` or `main`).
- All code changes must be committed and pushed.
- All tests must pass with 100% coverage.
- Documentation (including Godoc and `/docs`) must be up to date.
- The `RELEASENOTES.md` file must be updated with all major changes.
- Version numbers in documentation and code (if any) must be correct.

## 2. Final Quality Checks

```bash
# Run all tests
 go test ./...

# Check 100% test coverage
 go test -coverprofile=final_coverage.out ./...
 go tool cover -func=final_coverage.out | grep "total:"

# Run benchmarks (optional)
 go test -bench=. ./...

# Build verification
 go build ./...
```

## 3. Documentation Verification

- Generate and review Godoc for all public APIs:
  - `go doc ./...`
  - `godoc -http=:6060` (browse at http://localhost:6060/pkg/github.com/maniartech/gotime/)
- Review `/docs` for completeness and accuracy.
- Ensure all new features and changes are documented.

## 4. Final Commit & Push

```bash
git add .
git commit -m "Final preparations for vX.Y.Z release"
git push origin master
```

## 5. Tag the Release

```bash
# Replace vX.Y.Z with the new version
 git tag -a vX.Y.Z -m "Release vX.Y.Z: <short summary of major changes>"
 git push origin vX.Y.Z
```

## 6. Create GitHub Release

- Go to https://github.com/maniartech/gotime/releases
- Click "Draft a new release"
- Set the tag to `vX.Y.Z`
- Title: `GoTime vX.Y.Z - <short summary>`
- Description: Paste the relevant section from `RELEASENOTES.md`
- Click "Publish release"

## 7. Module Publication Verification

```bash
# Wait a few minutes for Go proxy to update
 go list -m -versions github.com/maniartech/gotime

# Test installation in a new directory
 mkdir /tmp/gotime-test && cd /tmp/gotime-test
 go mod init test
 go get github.com/maniartech/gotime@vX.Y.Z
 go list -m github.com/maniartech/gotime
```

## 8. Post-Release Tasks

- Announce the release (GitHub Discussions, Go community, social media, etc.)
- Monitor for issues or feedback
- Update any external documentation or references

## 9. Rollback Plan (if needed)

```bash
# Delete the tag locally and remotely
 git tag -d vX.Y.Z
 git push --delete origin vX.Y.Z
# Delete the GitHub release via the web interface
# Fix issues and re-release as vX.Y.Z or increment to vX.Y.(Z+1) if needed
```

## 10. Success Criteria

- [ ] Tag `vX.Y.Z` exists on GitHub
- [ ] GitHub release is published with release notes
- [ ] `go get github.com/maniartech/gotime@vX.Y.Z` works
- [ ] pkg.go.dev shows correct documentation
- [ ] All tests pass for the tagged version
- [ ] Announcement is made

## Notes
- For major releases, ensure all deprecated/breaking APIs are removed or properly documented.
- For minor/patch releases, focus on backward compatibility and bug fixes.
- Always update `RELEASENOTES.md` before tagging.
- This process is designed to be repeatable for all future releases.
