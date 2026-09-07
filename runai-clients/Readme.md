# runai-clients

This repository contains versioned Run.ai client packages. These packages are used as client libraries for various Run.ai services and infrastructure components.

## 🏗️ Developer Workflow (Quick Steps)

When developing with runai-clients:

1. **On your working branch:**
    - Run `make bump-version-dev` to create a new dev version release.
    - Update the relevant services to use the printed dev version.
2. **When your changes are ready:**
    - Run `make bump-version` (from your working branch) to open a PR for production release (if not on master).
    - Merge the new PR after review.
    - After merging, switch to `master` and run `make bump-version` again to create and push the production release.
    - Update your original service PR to use the new production version.

---

## Release Process

### 📦 Production Release

**Production releases can be performed from a synced master branch only.**

To release a new **production** version, use:

```sh
make bump-version
```

This command will:

1. **Check if you're on the master branch:**
    - If you are on `master` and it is synced with `origin/master`, attempts to create and push a new production tag.
    - Will fail if you are not checked out to the `master` branch or if `master` is not synced with `origin/master`.
2. **Open a Pull Request if you are NOT on master:**
    - If you are not on the `master` branch, the script will:
        - Prompt for a JIRA ticket if not found in your branch name.
        - Create a new branch from `master`.
        - Apply your `runai-clients` changes to the new branch.
        - Push the branch to the remote.
        - Open a Pull Request (PR) for your changes.
        - Print the PR link in the terminal and exit early (no tag is created).

#### Example Output (Success: Tag Created)

```text
Checking push permission...
You have permission to push a production tag.
Latest dev tag: runai-clients/v1.0.19-dev.2
Latest prod tag: runai-clients/v1.0.18
Bumped tag: runai-clients/v1.0.19
Pushed tag runai-clients/v1.0.19 to origin.
```

#### Example Output (No Tagging Permission, PR Opened)

```text
Branch is NOT up to date with origin/master. Opening a PR instead.
Target branch: RUN-29880-my-branch
Creating branch RUN-29880-my-branch from master...
Applying runai-clients changes from RUN-29880-improve-readme...
Committed runai-clients changes to RUN-29880-my-branch.
Pushing branch RUN-29880-my-branch to origin...
Branch RUN-29880-my-branch pushed to remote.
Checking for existing pull request for branch RUN-29880-my-branch...
No existing PR found. Creating a new pull request...
Pull request created! - https://github.com/dsx-ai-factory/run-ai/pull/10529
```

---

### 🧪 Development Release Process

**Development releases can be performed from any branch.**

To release a **development** version, use:

```sh
make bump-version-dev
```

This will:

1. **Determine the Next Dev Tag:**
    - Finds the latest dev tag and increments it, or creates a new dev tag based on the latest production tag.
2. **Create and Push the Dev Tag:**
    - Tags the current commit in the format:

```
runai-clients/vX.Y.Z-dev.N
```

    - Pushes the tag to the remote.

#### Example Output

```text
Latest tag: runai-clients/v1.0.3
Published tag: runai-clients/v1.0.4-dev.1
```

Or, if the latest tag is already a dev tag:

```text
Latest tag: runai-clients/v1.0.4-dev.1
Published tag: runai-clients/v1.0.4-dev.2
```

---

## 📝 Summary of Makefile Commands

| Command | Description |
| :-- | :-- |
| `make bump-version` | Bump and release a new production version (opens PR if not permitted to tag) |
| `make bump-version-dev` | Bump and release a new development version |
