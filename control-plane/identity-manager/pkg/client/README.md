## Identity Manager Client: Usage and Versioning Policy

### AuthnManager
`AuthnManager` manages OAuth2 client-credentials tokens and implements `http.RoundTripper`, so it can be used as a client's transport to automatically inject a Bearer token on every outgoing request. It refreshes tokens before expiry (using the configured refresh buffer) and retries once on 401 by refreshing the token and re-sending the request.

**Important:** When using an HTTP client whose transport is an `AuthnManager`, the manager's token is always set on the request. That takes precedence over any Bearer token already present on the request (e.g. from context or pre-set headers). Callers should not rely on passing a bearer token in context or headers when using this client; the manager overwrites the `Authorization` header.

### Importing the Client
Services requiring the identity-manager client must import it as a versioned Go package. For example:

```
import "control-plane/identity-manager/pkg/client/v1.0.0"
```

### Development Version Tagging
Developers may create a development version of the client within their working branch by executing the following command from the service's main directory:

```
make client-bump-version-dev
```

This process will generate a development version tag in the following format:

```
control-plane/identity-manager/pkg/client/v1.0.0-dev.0
```

Development versions are intended solely for use in feature branches or any branch other than `master`. **It is strictly prohibited to merge a pull request into the `master` branch if it references a development package version.** The continuous integration (CI) system will automatically fail any such pull request to enforce this policy.

### Production Release Procedure
To publish a production release of the client package, proceed as follows:
1. Merge all relevant client changes into the `master` branch.
2. From the service's main directory, execute:

```
make client-bump-version
```

This will create a production package release with the following format:

```
control-plane/identity-manager/pkg/client/v1.0.0
```

Production versions may be referenced in other branches and pull requests as required.

You can then use this prod version for your other PRs.