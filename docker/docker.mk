# Docker image build and push setting
DOCKER:=docker
DOCKERFILE_DIR?=./docker

APP_SYSTEM_IMAGE_NAME=$(RELEASE_NAME)
APP_RUNTIME_IMAGE_NAME=dashboard

# build docker image for linux
BIN_PATH=$(OUT_DIR)/$(TARGET_OS)_$(TARGET_ARCH)

DOCKERFILE:=Dockerfile
ifeq ($(TARGET_OS), windows)
  DOCKERFILE:=Dockerfile-windows
endif

# Supported docker image architecture
DOCKERMULTI_ARCH=linux-amd64 linux-arm linux-arm64 windows-amd64

################################################################################
# Target: docker-build, docker-push                                            #
################################################################################

LINUX_BINS_OUT_DIR=$(OUT_DIR)/linux_$(GOARCH)
DOCKER_IMAGE_TAG=$(APP_REGISTRY)/$(APP_SYSTEM_IMAGE_NAME):$(APP_TAG)
APP_RUNTIME_DOCKER_IMAGE_TAG=$(APP_REGISTRY)/$(APP_RUNTIME_IMAGE_NAME):$(APP_TAG)
APP_PLACEMENT_DOCKER_IMAGE_TAG=$(APP_REGISTRY)/$(APP_PLACEMENT_IMAGE_NAME):$(APP_TAG)
APP_SENTRY_DOCKER_IMAGE_TAG=$(APP_REGISTRY)/$(APP_SENTRY_IMAGE_NAME):$(APP_TAG)

ifeq ($(LATEST_RELEASE),true)
DOCKER_IMAGE_LATEST_TAG=$(APP_REGISTRY)/$(APP_SYSTEM_IMAGE_NAME):$(LATEST_TAG)
APP_RUNTIME_DOCKER_IMAGE_LATEST_TAG=$(APP_REGISTRY)/$(APP_RUNTIME_IMAGE_NAME):$(LATEST_TAG)
APP_PLACEMENT_DOCKER_IMAGE_LATEST_TAG=$(APP_REGISTRY)/$(APP_PLACEMENT_IMAGE_NAME):$(LATEST_TAG)
APP_SENTRY_DOCKER_IMAGE_LATEST_TAG=$(APP_REGISTRY)/$(APP_SENTRY_IMAGE_NAME):$(LATEST_TAG)
endif


# To use buildx: https://github.com/docker/buildx#docker-ce
export DOCKER_CLI_EXPERIMENTAL=enabled

# check the required environment variables
check-docker-env:
ifeq ($(APP_REGISTRY),)
	$(error APP_REGISTRY environment variable must be set)
endif
ifeq ($(APP_TAG),)
	$(error APP_TAG environment variable must be set)
endif

check-arch:
ifeq ($(TARGET_OS),)
	$(error TARGET_OS environment variable must be set)
endif
ifeq ($(TARGET_ARCH),)
	$(error TARGET_ARCH environment variable must be set)
endif


docker-build: check-docker-env check-arch
	$(info Building $(APP_RUNTIME_DOCKER_IMAGE_TAG) docker image ...)
	$(DOCKER) build --build-arg BIN_PATH=$(BIN_PATH) -f $(DOCKERFILE_DIR)/$(DOCKERFILE) . -t $(APP_RUNTIME_DOCKER_IMAGE_TAG)-$(TARGET_OS)-$(TARGET_ARCH) --platform $(TARGET_OS)/$(TARGET_ARCH)

# push docker image to the registry
docker-push: docker-build
	$(info Pushing $(APP_RUNTIME_DOCKER_IMAGE_TAG) docker image ...)
	$(DOCKER) push $(APP_RUNTIME_DOCKER_IMAGE_TAG)-$(TARGET_OS)-$(TARGET_ARCH)

# publish multi-arch docker image to the registry
docker-manifest-create: check-docker-env
	$(DOCKER) manifest create $(APP_RUNTIME_DOCKER_IMAGE_TAG) $(DOCKERMULTI_ARCH:%=$(APP_RUNTIME_DOCKER_IMAGE_TAG)-%)
ifeq ($(LATEST_RELEASE),true)
	$(DOCKER) manifest create $(APP_RUNTIME_DOCKER_IMAGE_LATEST_TAG) $(DOCKERMULTI_ARCH:%=$(APP_RUNTIME_DOCKER_IMAGE_TAG)-%)
endif

docker-publish: docker-manifest-create
	$(DOCKER) manifest push $(APP_RUNTIME_DOCKER_IMAGE_TAG)
ifeq ($(LATEST_RELEASE),true)
	$(DOCKER) manifest push $(APP_RUNTIME_DOCKER_IMAGE_LATEST_TAG)
endif

check-windows-version:
ifeq ($(WINDOWS_VERSION),)
	$(error WINDOWS_VERSION environment variable must be set)
endif

docker-windows-base-build: check-windows-version
	$(DOCKER) build --build-arg WINDOWS_VERSION=$(WINDOWS_VERSION) -f $(DOCKERFILE_DIR)/$(DOCKERFILE)-base . -t $(APP_REGISTRY)/windows-base:$(WINDOWS_VERSION)

docker-windows-base-push: check-windows-version
	$(DOCKER) push $(APP_REGISTRY)/windows-base:$(WINDOWS_VERSION)