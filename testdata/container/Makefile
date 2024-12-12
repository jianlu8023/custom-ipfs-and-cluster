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

clean: ipfs-clean cluster-clean compose-down image-clean
	@echo "starting clean"
.PHONY: clean

ipfs-clean:
	@echo "starting ipfs clean"
	$(MAKE) -C ../../go-ipfs clean
.PHONY: ipfs-clean

cluster-clean:
	@echo "starting cluster clean"
	$(MAKE) -C ../../ipfs-cluster clean
.PHONY: cluster-clean

image-clean:
	@echo "starting image clean"
	docker rmi ipfs-image
	docker rmi cluster-image
	docker rmi `docker images -f dangling=true`
	docker builder prune -a -f
.PHONY: image-clean

compose-down:
	@echo "starting compose down"
	docker-compose down
.PHONY: compose-down
