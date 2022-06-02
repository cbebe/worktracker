PORT := 5000

web:
	PORT=$(PORT) go run cmd/webserver/main.go

webserver: bin/webserver
cli: bin/cli
bin/%: cmd/%/main.go
	go build $<
	mv main $@

USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)
image:
	docker build \
		--build-arg USER_ID=$(USER_ID) \
		--build-arg GROUP_ID=$(GROUP_ID) \
		--build-arg PORT=$(PORT) \
		-t worktracker .
	
run:
	docker run -p $(PORT):$(PORT) -v $(PWD)/work.db:/app/work.db worktracker

.PHONY: cli webserver web image run
