VERSION ?= `git rev-parse --short HEAD`
CURRENT_DATE ?= `date -u +"%Y-%m-%dT%H:%M:%SZ"`

DOCKER_GO_PATH ?= "/go/src/github.com/arachnys/athenapdf"

CLI_IMAGE_FILE ?= "./docker/Dockerfile.cli"
CLI_IMAGE_NAME ?= "arachnysdocker/athenapdf"

SERVICE_IMAGE_FILE ?= "./docker/Dockerfile.service"
SERVICE_IMAGE_NAME ?= "arachnysdocker/athenapdf-service"

DEV_IMAGE_FILE ?= "./docker/Dockerfile.dev"
DEV_IMAGE_NAME ?= "arachnysdocker/athenapdf-dev"

P="\\033[34m[+]\\033[0m"

help:
	@echo
	@echo "  \033[34mbuild/dev\033[0m – builds dev docker image"
	@echo "  \033[34mbuild/cli\033[0m – builds athenapdf cli assembly docker image"
	@echo "  \033[34mbuild/service\033[0m – builds athenapdf service assembly docker image"
	@echo "  \033[34mrun/dev\033[0m – runs dev docker image (with code)"
	@echo "  \033[34mrun/cli\033[0m – runs athenapdf cli assembly docker image"
	@echo

build/dev:
	@echo "  $(P) build/dev"
	@make file=${DEV_IMAGE_FILE} image=${DEV_IMAGE_NAME} tag=latest _build

build/cli:
	@echo "  $(P) build/cli"
	@make build/dev
	@make file=${CLI_IMAGE_FILE} image=${CLI_IMAGE_NAME} tag=latest _build

build/service:
	@echo "  $(P) build/service"
	@make build/cli
	@make file=${SERVICE_IMAGE_FILE} image=${SERVICE_IMAGE_NAME} tag=latest _build

run/dev:
	@echo "  $(P) run/dev"
	docker run --rm -it \
				-v `pwd`:${DOCKER_GO_PATH} \
				-w ${DOCKER_GO_PATH} \
				${DEV_IMAGE_NAME}:latest \
				${args}

run/cli:
	@echo "  $(P) run/cli"
	VERSION=latest ./bin/athenapdf.sh ${args}

version:
	@echo "${VERSION}"

_build:
	@echo "building: ${image}:${tag}"
	@docker build --rm \
				--build-arg BUILD_DATE=${CURRENT_DATE} \
				--build-arg VCS_REF=${VERSION} \
				-f ${file} \
				-t ${image}:${tag} \
				.

.PHONY: help build/dev build/cli build/service run/dev run/cli version _build