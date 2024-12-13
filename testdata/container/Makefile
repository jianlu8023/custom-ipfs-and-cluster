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
	@if docker images | grep -q ipfs-image; then \
  		echo "ipfs-image already exists"; \
 	else \
		$(MAKE) -C ../../go-ipfs docker; \
	fi
	docker builder prune -a -f
.PHONY: ipfs

cluster:
	@echo "starting cluster"
	@if docker images | grep -q cluster-image; then \
  		echo "cluster-image already exists"; \
  	else \
  		$(MAKE) -C ../../ipfs-cluster docker; \
  	fi
	@if docker images | grep -q cluster-image-test; then \
  		docker rmi cluster-image-test; \
  	fi
	docker builder prune -a -f
.PHONY: cluster

clean: compose-down ipfs-clean cluster-clean image-clean
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
	@if docker images | grep -q ipfs-image; then \
          echo "Removing ipfs-image"; \
          docker rmi ipfs-image; \
    else \
          echo "ipfs-image does not exist, skipping remove"; \
    fi
	@if docker images | grep -q cluster-image; then \
    	echo "Removing cluster-image"; \
        docker rmi cluster-image; \
    else \
        echo "cluster-image does not exist, skipping remove"; \
    fi
	 @if docker images -f dangling=true  --format ".table"  | grep -q .; then \
    	echo "Removing dangling images"; \
        docker rmi `docker images -f dangling=true --format "{{.Repository}}:{{.Tag}}"`; \
    else \
        echo "No dangling images found, skipping remove"; \
    fi
	@if docker volume ls -q -f dangling=true --format ".table" | grep -q .; then \
          echo "Removing all Docker volumes"; \
          docker volume rm `docker volume ls -q -f dangling=true`; \
    else \
          echo "No Docker volumes found, skipping remove"; \
    fi
	docker builder prune -a -f
.PHONY: image-clean

compose-down:
	@echo "starting compose down"
	docker-compose down
.PHONY: compose-down
