
export GO111MODULE := on
export GOPROXY=https://goproxy.cn,direct
GOENV  := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
SNRCLIENT_PKG := snr-client
TAG?=dev
EDITION ?= Commercial
LDFLAGS += -X "$(SNRCLIENT_PKG)/pkg/version.releaseVersion=$(shell git describe --tags --always)"
LDFLAGS += -X "$(SNRCLIENT_PKG)/pkg/version.buildDate=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "$(SNRCLIENT_PKG)/pkg/version.gitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "$(SNRCLIENT_PKG)/pkg/version.gitBranch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "$(SNRCLIENT_PKG)/pkg/version.edition=$(EDITION)"
LDFLAGS += -w -s
TOOLEXEC?=

GO_BUILD := $(GOENV) go build -trimpath -gcflags "all=-N -l" -ldflags '$(LDFLAGS)' $(TOOLEXEC)
SNRCLIENT ?= leovamwong/snr-client:${TAG}
snr-client:
	$(GO_BUILD) -o bin/snr-client main.go

.PHONY: snr-client-release
snr-client-release: snr-client
	cd bin && tar -czf SonatypeNexusRepositoryClient-linux-amd64.tar.gz snr-client

image-build-snr-client:
	docker build . -f build/Dockerfile -t ${SNRCLIENT}

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...
test:
	go test ./... -coverprofile coverage.out
	go tool cover -func=coverage.out

yaml-check:
	kubeconform

