# runai-clients

This repository contains versioned Run.ai client packages. These packages are used as client libraries for various Run.ai services and infrastructure components.

## Release Process

### 📦 Production Release

**Production releases can be performed from a synced master branch only.**

To release a new **production** version, use:

```sh
make bump-version
```

This command will:

1. **Check Tagging Permissions:**
Attempts to create and push a new production tag. If you have permission, the tag is created and pushed automatically.
    - Will Fail if you are not checked out to the `master` branch
    - Will Fail if `master` is not synced with `origin/master`
2. **Open a Pull Request if Tagging is Not Allowed:**
If you **do not have permission** to push tags, the script will:
    - Prompt for a JIRA ticket if not found in your branch name.
    - Create a new branch from `master` and apply your `runai-clients` changes.
    - Push the branch and open a Pull Request (PR) for your changes.
    - Print the PR link in the terminal and exit early (no tag is created).

#### Example Output (Success: Tag Created)

```text
Checking push permission
You have permission to push a production tag
Latest dev tag: runai-clients/v1.0.19-dev.2
Latest prod tag: runai-clients/v1.0.18
Bumped tag: runai-clients/v1.0.19
```


#### Example Output (No Tagging Permission, PR Opened)

```text
Checking push permission
No push permission, lets open a PR
Target branch: RUN-29880-clients
Branch RUN-29880-clients already exists.
Applying runai-clients changes from
No runai-clients changes to commit.
Pushing branch RUN-29880-clients to origin...
Branch RUN-29880-clients pushed to remote.
Checking for existing pull request for branch
Existing PR found. Updating branch...
Pull request created! - https://github.com/dsx-ai-factory/run-ai/pull/10387
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
Finds the latest dev tag and increments it, or creates a new dev tag based on the latest production tag.
2. **Create and Push the Dev Tag:**
Tags the current commit in the format:

```
runai-clients/vX.Y.Z-dev.N
```

and pushes it to the remote.

#### Example Output

```text
Latest tag: runai-clients/v1.0.19
Bumped tag: runai-clients/v1.0.20-dev.0
```

Or, if the latest tag is already a dev tag:

```text
Latest tag: runai-clients/v1.0.20-dev.0
Bumped tag: runai-clients/v1.0.20-dev.1
```


---

## 📝 Summary of Makefile Commands

| Command | Description |
| :-- | :-- |
| `make bump-version` | Bump and release a new production version (opens PR if not permitted to tag) |
| `make bump-version-dev` | Bump and release a new development version |
