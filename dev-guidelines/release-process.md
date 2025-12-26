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
 git push origin master
 git push origin vX.Y.Z
```

> **Why push both `master` and the tag?**
>
> - `git push origin master` updates the main branch with your latest commits, ensuring the codebase is current for all collaborators and CI/CD systems.
> - `git push origin vX.Y.Z` pushes the release tag, creating an immutable reference to the exact commit for this release. This is what users and Go modules fetch for versioned installs.
>
> Both are required: the branch for ongoing development, and the tag for versioned releases.

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

### Speeding up pkg.go.dev updates

To speed up the appearance of your new version on pkg.go.dev:

1. Visit: https://pkg.go.dev/github.com/maniartech/gotime@vX.Y.Z (replace with your version)
2. Click the "Request" or "Refresh" button (if available) to trigger a re-index.
3. Wait a few minutes and refresh the page.

If it still doesn't update after 30–60 minutes:
- Double-check that the tag is pushed and visible on GitHub.
- Make sure the tag is annotated (not lightweight).
- Ensure your repository is public.

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

### Deleting Invalid or Obsolete Releases

If you need to remove a release (for example, if a tag was created with an incorrect module path or for any other reason):

1. Delete the tag locally and remotely as shown above.
2. Go to your repository’s "Releases" page on GitHub.
3. Find the release associated with the tag you want to remove.
4. Click "Delete" (trash icon or "Delete this release").
5. Confirm the deletion.

After correcting the issue (such as updating the module path or fixing release content), create a new tag and draft a new GitHub release as usual. This ensures only valid, working releases are visible to users and tools like pkg.go.dev.
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
