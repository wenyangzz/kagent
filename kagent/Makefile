
# Image URL to use all building/pushing image targets
IMG ?= kagent:latest
# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.22

all: agent

# Build agent binary
agent: fmt vet
	go build -o bin/agent ./cmd/agent

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	# Remove the certmanager certificate if KSERVE_ENABLE_SELF_SIGNED_CA is not false
	cd config/default && if [ ${KSERVE_ENABLE_SELF_SIGNED_CA} != false ]; then \
	kustomize edit remove resource certmanager/certificate.yaml; \
	else kustomize edit add resource certmanager/certificate.yaml; fi;
	kustomize build config/default | kubectl apply -f -
	if [ ${KSERVE_ENABLE_SELF_SIGNED_CA} != false ]; then ./hack/self-signed-ca.sh; fi;

undeploy:
	kustomize build config/default | kubectl delete -f -
	kubectl delete validatingwebhookconfigurations.admissionregistration.k8s.io inferenceservice.serving.kserve.io
	kubectl delete validatingwebhookconfigurations.admissionregistration.k8s.io trainedmodel.serving.kserve.io
	kubectl delete validatingwebhookconfigurations.admissionregistration.k8s.io inferencegraph.serving.kserve.io
	kubectl delete mutatingwebhookconfigurations.admissionregistration.k8s.io inferenceservice.serving.kserve.io

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Build the docker image
docker-build:
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}

