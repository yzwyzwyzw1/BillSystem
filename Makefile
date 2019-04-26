.PHONY: all dev env-clean build env-up env-down run
all: env-clean build env-up run
dev: build run
##### BUILD
build:
	@echo "Build ..."
#	@dep ensure
	@go build
	@echo "Build done"
##### ENV
env-up:
	@echo "Start environment ..."
	@docker-compose -f  fixtures/docker-compose-cli.yaml  up -d
	@echo "Sleep 1 seconds in order to let the environment setup correctly"
	@sleep 1
	@docker ps
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@docker-compose -f fixtures/docker-compose-cli.yaml down
	@docker ps
	@echo "Environment down"
##### RUN
run:
	@echo "Start app ..."
	@./BillSystem

##### CLEAN
env-clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/state-store
	@rm -rf /tpm/msp
	@rm BillSystem
	@docker volume prune -f #  清理挂载卷
	@docker network prune -f #  来清理没有再被任何容器引用的networks
	@docker rm -f -v `docker ps -a --no-trunc | grep "mycc" | cut -d ' ' -f 1` 2>/dev/null || true  # 清除链码
	@docker rmi `docker images --no-trunc | grep "mycc" | cut -d ' ' -f 1` 2>/dev/null || true
	@docker ps
	@echo "Clean up done"
