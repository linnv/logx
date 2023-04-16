GOFMT ?= gofmt "-s"
SOURCES ?= $(shell find . -path ./vendor -prune -o -name "*.go" -type f)

all:
	go test 
	# go test -i
fmt:
	$(GOFMT) -w $(SOURCES)
