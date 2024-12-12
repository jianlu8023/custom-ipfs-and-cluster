all: run
.PHONY: all

run: build
	@echo "starting run"
	docker-compose up -d
	@echo "done"
.PHONY: run

build: ipfs cluster
	@echo "starting build"
.PHONY: build

ipfs:
	@echo "starting ipfs"
	$(MAKE) -C ../../go-ipfs docker
	docker builder prune -a -f
.PHONY: ipfs

cluster:
	@echo "starting cluster"
	$(MAKE) -C ../../ipfs-cluster docker
	docker rmi cluster-image-test
	docker builder prune -a -f
.PHONY: cluster