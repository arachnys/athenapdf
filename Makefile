CLI_DIR ?= cli
CLI_IMAGE ?= "simplycredit/athenapdf"
CLI_DOCKER_ARTIFACT_DIR ?= "/athenapdf/build/"

SERVICE_DIR ?= weaver
SERVICE_IMAGE ?= "simplycredit/athenapdf-service"
SERVICE_DOCKER_ARTIFACT_FILE ?= "/go/src/github.com/arachnys/athenapdf/weaver"

P="\\033[34m[+]\\033[0m"

help:
	@echo
	@echo "  \033[34mbuildcli\033[0m – builds athenapdf (cli) docker image"
	@echo "  \033[34mtestcli\033[0m – tests athenapdf (cli) standard streams"
	@echo "  \033[34mbuildservice\033[0m – builds athenapdf-service docker image"
	@echo "  \033[34mtestservice\033[0m – tests athenapdf-service Go source"
	@echo "  \033[34mbuild\033[0m – builds both the cli, and service docker image"
	@echo

buildcli:
	@echo "  $(P) buildcli"
	@rm -rf $(CLI_DIR)/build/
	@docker build --rm=false -t $(CLI_IMAGE)-build -f $(CLI_DIR)/Dockerfile.build $(CLI_DIR)/
	@docker run -t --name cli-build $(CLI_IMAGE)-build /bin/true
	@docker cp cli-build:$(CLI_DOCKER_ARTIFACT_DIR) $(CLI_DIR)/build/
	@docker build --rm=false -t $(CLI_IMAGE) -f $(CLI_DIR)/Dockerfile $(CLI_DIR)/
	@rm -rf $(CLI_DIR)/build/

testcli:
	@echo "  $(P) testcli"
	@docker run $(CLI_IMAGE) athenapdf -S https://www.traviscistatus.com/ | grep -a "PDF-1.4"
	@echo "<h1>stdin test</h1>" | docker run -i $(CLI_IMAGE) athenapdf -S - | grep -a "PDF-1.4"

buildservice:
	@echo "  $(P) buildservice"
	@rm -rf $(SERVICE_DIR)/build/
	@docker build --rm=false -t $(SERVICE_IMAGE)-build -f $(SERVICE_DIR)/Dockerfile.build $(SERVICE_DIR)/
	@docker run -t --name svc-build $(SERVICE_IMAGE)-build /bin/true
	@docker cp svc-build:$(SERVICE_DOCKER_ARTIFACT_FILE) $(SERVICE_DIR)/build/
	@chmod +x $(SERVICE_DIR)/build/weaver
	@docker build --rm=false -t $(SERVICE_IMAGE) -f $(SERVICE_DIR)/Dockerfile $(SERVICE_DIR)/
	@rm -rf $(SERVICE_DIR)/build/

testservice:
	@echo "  $(P) testservice"
	@docker build --rm=false -t $(SERVICE_IMAGE)-test -f $(SERVICE_DIR)/Dockerfile.build $(SERVICE_DIR)/
	@docker run -t $(SERVICE_IMAGE)-test go test ./...

build:
	@echo "  $(P) build"
	@make buildcli
	@make buildservice

.PHONY: help buildcli testcli buildservice testservice build
