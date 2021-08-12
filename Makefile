.PHONY: build/container changelog ci/build ci/release/dryrun release test

export GOVERSION = $(shell cat .go-version)

build/container:
	docker build . -t semtag --build-arg GOVERSION

changelog:
	git-chglog -o CHANGELOG.md

ci/build:
	docker run -it --rm -v $(PWD):/src -w /src goreleaser/goreleaser:latest build --snapshot --rm-dist

ci/release/dryrun:
	docker run -it --rm -v $(PWD):/src -w /src goreleaser/goreleaser:latest release --snapshot --rm-dist --skip-publish

install:
	go install

test:
	go test -v github.com/leb4r/semtag/cmd

release:
	semtag final -s minor
