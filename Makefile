.PHONY: build test

GOOS := windows darwin linux
GOARCH := amd64 arm64
TARGETS := $(foreach os,$(GOOS),$(foreach arch,$(GOARCH),build/$(os)/$(arch)/cogen))

all: $(TARGETS)

define build-target
build/$(1)/$(2)/cogen: main.go
	GOOS=$(1) GOARCH=$(2) go build -o $$@ main.go
endef

$(foreach os,$(GOOS),$(foreach arch,$(GOARCH),$(eval $(call build-target,$(os),$(arch)))))

test:
	go test -v ./...

clean:
	rm -rf build

.DEFAULT_GOAL := all