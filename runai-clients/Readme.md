# runai-clients

This repository contains versioned Run.ai client packages. These packages are used as client libraries for various Run.ai services and infrastructure components.

## Release Process

> [!IMPORTANT]
> Tagging and releasing should only be performed by team leads or authorized personnel.

To release a Go package from runai-clients, follow these steps:

> [!IMPORTANT]
> This is a production release, for a dev release check out [here](#development-release-process)

1. **Tag the Branch**:  
   To release a package, tag the master branch using the following format:

   ```
   runai-clients/vX.X.X
   ```

   - Replace `vX.X.X` with the desired version number (following semantic versioning rules).

   Example tag:

   ```
   runai-clients/v1.2.3
   ```

2. **Import the Package**:  
   Once the branch is tagged, it can be imported by other services or repositories using the standard Go module import syntax:

   ```go
   import "github.com/dsx-ai-factory/run-ai/runai-clients/your-client-name"
   ```

3. **Follow Versioning Rules**:  
   When tagging a release, adhere to [semantic versioning](https://semver.org/) guidelines to maintain consistency and stability across your packages. The format for versioning is `vMAJOR.MINOR.PATCH`.

   - **MAJOR**: Introduced Incompatible API changes
   - **MINOR**: Added Backward-compatible functionality
   - **PATCH**: Backward-compatible bug fixes

   Example:
    ```yaml
   - v1.0.0
   - v2.1.0
   - v2.1.1
    ```
<br/><br/>

> [!IMPORTANT] 
> For runai-clients, updating the **major version** should usually be avoided, even for breaking changes. Prefer incrementing the minor or patch version and communicate breaking changes clearly in the release notes. Only update the major version in exceptional cases after consulting with the team leads.
<br/><br/>

## Development Release Process
**A dev tag can be performed from any branch**

To release a development version of runai-clients, use the same tagging process as above but append `-dev.x` to the version number.

Example:
```
runai-clients/v1.2.3-dev.0
```

This tag will mark the package as a development release, indicating that it is a work-in-progress version, not yet intended for production use.

### Versioning Development Releases

Development releases should also follow proper versioning guidelines. However, the `-dev.x` suffix signals that the release is not considered stable and may still undergo changes before the official release.

- Use `vX.X.X-dev.x` for development versions leading up to a full release.
- Increment `DEV` version, until you ready to release the next patch/minor (e.g., `v1.2.3-dev.1`).
- Once the package is stable, remove the `-dev.x` suffix and create the final release tag (e.g., `v1.2.3`).
<br/><br/>

## Importing Private Go Modules in Developer Workspace

To import private Go modules, you need to configure both your Git settings and Go environment. Follow these steps:

### Steps:

1. Run the following command to configure your Git settings:

   ```sh
   git config --file ${HOME}/.gitconfig url.https://<YOUR_GITHUB_TOKEN>@github.com/.insteadOf https://github.com/
   ```

2. Set the **GOPRIVATE** environment variable:

   ```sh
   export GOPRIVATE=github.com/run-ai/
   ```
> [!IMPORTANT]  
> Make sure to add this to your .bashrc or .zshrc

3. After configuring **Git** and setting **GOPRIVATE**, you can fetch private Go modules by running:

   `go get` or `go mod tidy`

---

### Explanation:

- **Git Configuration**:  
  The first command configures Git to use your GitHub token when accessing `https://github.com/` so that private repositories can be fetched without manually entering credentials.

- **<YOUR_GITHUB_TOKEN>**:  
  Replace this placeholder with your personal GitHub token. Ensure the token has the necessary scopes, such as `repo`, to access private repositories.

- **GOPRIVATE=github.com/run-ai**:  
  This tells Go that any modules under `github.com/run-ai` are private, so it bypasses the public proxy and checksum database when fetching these modules.

