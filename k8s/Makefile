.PHONY: setup-rancher
setup-rancher:
	@kubectl config use-context rancher-desktop

.PHONY: setup-minikube
setup-minikube:
	@minikube start --profile local-dev
	@eval $(minikube -p local-dev docker-env)

.PHONY: install-prometheus
install-prometheus:
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo update
	helm install prom prometheus-community/prometheus -f ./prom-values.yaml -n scylla-cloud
