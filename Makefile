
export GO111MODULE := on
export GOPROXY=https://goproxy.cn,direct
GOENV  := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
AL_CLOUD_PKG := al-cloud
EDITION ?= Commercial
LDFLAGS += -X "$(AL_CLOUD_PKG)/pkg/version.releaseVersion=$(shell git describe --tags --always)"
LDFLAGS += -X "$(AL_CLOUD_PKG)/pkg/version.buildDate=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "$(AL_CLOUD_PKG)/pkg/version.gitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "$(AL_CLOUD_PKG)/pkg/version.gitBranch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "$(AL_CLOUD_PKG)/pkg/version.edition=$(EDITION)"
LDFLAGS += -w -s
TOOLEXEC?=

GO_BUILD := $(GOENV) go build -trimpath -gcflags "all=-N -l" -ldflags '$(LDFLAGS)' $(TOOLEXEC)

snr-client:
	$(GO_BUILD) -o bin/snr-client main.go