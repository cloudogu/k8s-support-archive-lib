PROJECT_NAME=k8s-support-archive-lib
ARTIFACT_ID=k8s-support-archive-operator-crd
APPEND_CRD_SUFFIX=false
VERSION=0.1.2
## Image URL to use all building/pushing image targets
IMAGE=cloudogu/${ARTIFACT_ID}:${VERSION}
GOTAG?=1.24.1
MAKEFILES_VERSION=10.1.1

ADDITIONAL_CLEAN=dist-clean

PRE_COMPILE = generate-deepcopy
CRD_POST_MANIFEST_TARGETS = crd-add-labels

include build/make/variables.mk
include build/make/self-update.mk
include build/make/dependencies-gomod.mk
include build/make/build.mk
include build/make/test-common.mk
include build/make/test-unit.mk
include build/make/static-analysis.mk
include build/make/clean.mk
include build/make/digital-signature.mk
include build/make/mocks.mk
include build/make/k8s-controller.mk

##@ Debug

.PHONY: print-debug-info
print-debug-info: ## Generates info and the list of environment variables required to start the operator in debug mode.
	@echo "The target generates a list of env variables required to start the operator in debug mode. These can be pasted directly into the 'go build' run configuration in IntelliJ to run and debug the operator on-demand."
	@echo "STAGE=$(STAGE);LOG_LEVEL=$(LOG_LEVEL);KUBECONFIG=$(KUBECONFIG);NAMESPACE=$(NAMESPACE)"

# Override make target to use k8s-support-archive-lib as label
.PHONY: crd-add-labels
crd-add-labels: $(BINARY_YQ)
	@echo "Adding labels to CRD..."
	@for file in ${HELM_CRD_SOURCE_DIR}/templates/*.yaml ; do \
		$(BINARY_YQ) -i e ".metadata.labels.app = \"ces\"" $${file} ;\
		$(BINARY_YQ) -i e ".metadata.labels.\"app.kubernetes.io/name\" = \"${PROJECT_NAME}\"" $${file} ;\
	done
