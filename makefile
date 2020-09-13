DOCKER_REGISTRY   ?= 
IMAGE_PREFIX      ?= 
DEV_IMAGE         ?= 
SHORT_NAME        ?= 
TARGETS           ?= linux/amd64
TARGET_OBJS       ?= linux-amd64.tar.gz linux-amd64.tar.gz.sha256
DIST_DIRS         = find * -type d -exec

# go option
GO        ?= go
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   := -w -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
BINARIES  := rofi-leo-dict
BINARIES_BUMP = $(patsubst %,bump-%, $(BINARIES))

# tools
CP := cp -u -v

# Required for globs to work correctly
SHELL=/usr/bin/env sh

.PHONY: all
all: build

.PHONY: release
release: build bump-all

.PHONY: install
install: build
	$(CP) dist/$(BINARIES) ~/.bin/
	$(CP) scripts/* ~/.bin/

.PHONY: build
build: $(BINARIES)
	go version

.PHONY: $(BINARIES)
$(BINARIES):
	echo "building $@"
	$(eval LDF_LOCAL := ${LDFLAGS})
	$(eval VERSION_FILE := cmd/$@/VERSION.txt)
	$(eval VERSION := $(shell cat ${VERSION_FILE}))
	$(eval LDF_LOCAL += -X main.version=${VERSION})
	$(GO) build $(GOFLAGS) -tags '$(TAGS)' -o ./dist/$@  -ldflags "$(LDF_LOCAL)" ./cmd/$@
	echo "done building $@ -> dist/$@"

# github.com/jessfraz/junk/sembump download to gopath externally
.PHONY: bump-all
bump-all: $(BINARIES_BUMP)

.PHONY: $(BINARIES_BUMP)
BUMP := patch
$(BINARIES_BUMP):
	@echo "(${BUMP})ing, target: $(patsubst bump-%,%, $@)"

	$(eval VERSION_FILE := cmd/$(patsubst bump-%,%, $@)/VERSION.txt)
	$(eval VERSION := $(shell cat ${VERSION_FILE}))
	$(eval NEW_VERSION = $(shell sembump --kind $(BUMP) $(VERSION)))

	@echo "Bumping VERSION.txt from $(VERSION) to $(NEW_VERSION)"
	@echo $(NEW_VERSION) > ${VERSION_FILE}
	@git add ${VERSION_FILE}
	@git commit -vsam "Bump '$(patsubst bump-%,%, $@)' version to $(NEW_VERSION)"

.PHONY: test
test: build
test: TESTFLAGS += -race -v

.PHONY: test-unit
test-unit:
	@echo
	@echo "==> Running unit tests <=="
	$(GO) test $(GOFLAGS) -run $(TESTS) $$(go list ./... | grep -v /vendor/) $(TESTFLAGS)
	
.PHONY: generate
generate:
	$(GO) generate ./...

.PHONY: clean
clean:
	@rm -rf $(BINDIR) ./_dist
	rm -rf dist/*
	find static -type f -name '*.br' -delete
	find static -type f -name '*.gz' -delete

include versioning.mk
