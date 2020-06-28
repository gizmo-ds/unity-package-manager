NAME=UnityPackageManager
VERSION=$(shell echo "$${GITHUB_REF\#refs/heads/}")
GIT_HASH=$(shell git show -s --format=%h)
LINKFLAGS=-H windowsgui -X upm/cmd.gitHash=${GIT_HASH} -X upm/cmd.version=${VERSION}

.PHONY: all

all: windows-amd64

update-versioninfo:
	cat versioninfo.json | jq -c ".StringFileInfo.ProductVersion=\"$(VERSION)\"" | jq -c ".StringFileInfo.FileVersion=\"$version (build_$(GIT_HASH))\"" > versioninfo.json

windows-amd64: update-versioninfo
	cat versioninfo.json
	go generate
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o "./out/unity-package-manager.exe" -ldflags "$(LINKFLAGS)"
