# Scylla Cloud POC

```bash
make install-dependencies
make gen
make build
source <(./scylla-cloud completion bash)
```

## Using Rancher Desktop

- [Install Rancher Desktop](https://rancherdesktop.io)
- run `make -C k8s setup-rancher`
- To access services from your dev machine open rancher desktop, find the appropriate port in the `Port Forwarding` tab and click `Forward`. This will forward the service port to a port on your local machine.

## Using Minikube

- [Install minikube](https://minikube.sigs.k8s.io/docs/start)
- run `make -C k8s setup-minikube`

## Running

To run everything in a local Kubernetes cluster using [Skaffold](https://skaffold.dev) you the option of deploying to Rancher Desktop or Minikube.

* `make start-dev-env` - start development environment
* `make k8s-build` - build and redeploy
