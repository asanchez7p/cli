
FILES = $(shell find . -type f -name '*.go' -not -path './vendor/*')

gofmt:
	@impsort cmd internal -p github.com/altipla-sites/cli
	@gofmt -s -w $(FILES)
	@gofmt -r '&α{} -> new(α)' -w $(FILES)
