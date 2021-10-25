ifeq ($(GOPATH),)
GOPATH := $(HOME)/go
endif

ifeq ($(OS),Windows_NT)
	LDFLAGS="extldflags=-static"
	OPERATING_SYSTEM="windows"
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
		LDFLAGS="-linkmode external -extldflags=-static"
		OPERATING_SYSTEM="linux"
    endif
    ifeq ($(UNAME_S),Darwin)
		LDFLAGS="extldflags=-static"
		OPERATING_SYSTEM="macos"
    endif
endif


all: test lint vet build

build: reg

reg:
	@echo "*** building $@"
	@cd cmd/$@ && go build -trimpath -o ../../bin/$@ -tags osusergo,netgo,sqlite_omit_load_extension -ldflags=${LDFLAGS}

clean:
	@rm -f bin/*
	@go clean -testcache

test:
	@echo "*** test"
	@go test ./...

test_clean:
	@go clean -testcache 
	@go test ./... -timeout 30s

test_verbose:
	@go test ./... -v

test_race:
	@go test ./... -race

lint:
	@echo "*** linting"
	@revive ./... 

vet:
	@echo "*** vetting"
	@go vet ./...

staticcheck:
	@staticcheck ./...

benchmark:
	@go test -bench . ./...

count:
	@echo "Linecounts excluding generated and third party code"
	@gocloc --not-match-d='apipb|openapiv2' .
