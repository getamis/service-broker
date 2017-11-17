# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: all docker grpc test clean

CURDIR := $(shell pwd)
GOBIN = $(shell pwd)/build/bin
DOCKER_REPOSITORY := quay.io/amis
DOCKER_IMAGE := $(DOCKER_REPOSITORY)/broker
ifeq ($(REV),)
REV := $(shell git rev-parse --short HEAD 2> /dev/null)
endif

TARGETS := $(sort $(notdir $(wildcard ./cmd/*)))
.PHONY: $(TARGETS)

all: $(TARGETS)

.SECONDEXPANSION:
$(TARGETS): $(addprefix $(GOBIN)/,$$@)

$(GOBIN):
	@mkdir -p $@

$(GOBIN)/%: $(GOBIN) FORCE
	@go build -o $@ ./cmd/$(notdir $@)
	@echo "Done building."
	@echo "Run \"$(subst $(CURDIR),.,$@)\" to launch $(notdir $@)."

coverage.txt:
	@touch $@

test: coverage.txt FORCE
	@for d in `go list ./... | grep -v vendor | grep -v mock`; do		\
		go test -v -coverprofile=profile.out -covermode=atomic $$d;	\
		if [ $$? -eq 0 ]; then						\
			echo "\033[32mPASS\033[0m:\t$$d";			\
			if [ -f profile.out ]; then				\
				cat profile.out >> coverage.txt;		\
				rm profile.out;					\
			fi							\
		else								\
			echo "\033[31mFAIL\033[0m:\t$$d";			\
			exit -1;						\
		fi								\
	done;

BROKER_GRPC_PROTOS := $(CURDIR)/broker/pb/*.proto
PROTOS := \
	$(BROKER_GRPC_PROTOS)

grpc: FORCE
	@protoc \
		-I$(CURDIR)/vendor/github.com/golang/protobuf/ptypes \
		-I$(CURDIR)/vendor/github.com/golang/protobuf/ptypes/any \
		-I$(CURDIR)/vendor/github.com/golang/protobuf/ptypes/struct \
		-I$(GOPATH)/src \
		--go_out=plugins=grpc:$(GOPATH)/src $(PROTOS)

docker:
	@docker build -t $(DOCKER_IMAGE):$(REV) .

dockerpush:
	@docker push $(DOCKER_IMAGE):$(REV)

clean:
	rm -fr $(GOBIN)/*

PHONY: help
help:
	@echo  'Generic targets:'
	@echo  '  all               - Build all targets marked with [*]'
	@echo  '* broker            - Build Broker'
	@echo  ''
	@echo  'Protobuf targets:'
	@echo  '  grpc              - Generate gRPC go bindings from .proto files'
	@echo  ''
	@echo  'Docker targets:'
	@echo  '  docker            - Build docker image'
	@echo  ''
	@echo  'Test targets:'
	@echo  '  test              - Run all unit tests'
	@echo  ''
	@echo  'Cleaning targets:'
	@echo  '  clean             - Remove built executables'
	@echo  ''
	@echo  'Execute "make" or "make all" to build all targets marked with [*] '
	@echo  'For further info see the ./README.md file'

.PHONY: FORCE
FORCE:
