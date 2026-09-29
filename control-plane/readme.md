# Run:ai Control Plane

Control plane services and the `control-plane-chart` Helm chart. Each service directory has its own `README.md` with build, test, and run instructions.

To spin up a full local environment, run `./local-dev.sh` from the repository root.

## Creating a New Service

Create a new service from a service template using cookiecutter. The template is located under `control-plane/infra/service-cookiecutter`.

1. Install cookiecutter
1. Go to the `control-plane` directory
1. Run `cookiecutter infra/service-cookiecutter`

For more information, see `control-plane/infra/service-cookiecutter/README.md`.
