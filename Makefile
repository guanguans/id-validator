# note: call scripts from /scripts

.PHONY: \
	golangci-lint \
	gosec \
	fmt \
	vet \
	test \
	bench \
	cover \
	actionlint \
	checks \
	checks-required \
	checks-optional \
	git-chglog \
	gitleaks \
	gitleaks-generate-baseline \
	lint-md \
	lint-md-fix \
	lint-md-prototype \
	todo-lint \
	trufflehog \
	typos \
	typos-write-changes \
	vhs \
	zhlint \
	zhlint-fix \
	zhlint-prototype \
	zizmor

golangci-lint:
	golangci-lint run ./... --color=always --verbose

gosec:
	gosec ./... -color -verbose

fmt:
	go fmt ./...

fumpt:
	gofumpt -d -e -l -w -extra -modpath $(shell find $$PWD -iname "*.go" -not -iname "*pb.go" -not -iwholename "*vendor*")

vet:
	go vet ./...

test:
	go test ./... -covermode=atomic -coverprofile=coverage -cover -race -v

bench:
	go test ./... -bench=. -benchmem -v

cover:
	$(MAKE) test
	go tool cover -html=coverage

actionlint:
	actionlint -ignore=SC2035 -ignore=SC2086 -color -oneline -verbose

checks:
	$(MAKE) checks-required
	$(MAKE) checks-optional

checks-required:
	$(MAKE) fmt
	$(MAKE) test

checks-optional:
	$(MAKE) actionlint
	$(MAKE) gitleaks
	$(MAKE) lint-md
	$(MAKE) todo-lint
	$(MAKE) typos
	$(MAKE) zhlint

git-chglog:
	git-chglog $$(git describe --tags $$(git rev-list --tags --max-count=1))

gitleaks:
	gitleaks git --report-path=.build/gitleaks-report.json -v

gitleaks-generate-baseline:
	gitleaks git --report-path=gitleaks-baseline.json -v

lint-md:
	if ! command -v lint-md >/dev/null 2>&1; then echo 'lint-md not found, installing...'; npm install -g @lint-md/cli; fi
	$(MAKE) lint-md-prototype

lint-md-fix:
	$(MAKE) lint-md-prototype LINT_MD_EXTRA=--fix

lint-md-prototype:
	lint-md --suppress-warnings *.md .github/ docs/ $(LINT_MD_EXTRA)

todo-lint:
	! git --no-pager grep --extended-regexp --ignore-case 'todo|fixme' -- '*.go' ':!*.blade.go' ':(exclude)resources/'

trufflehog:
	trufflehog git https://github.com/guanguans/id-validator --only-verified

typos:
	typos --color=always --sort --verbose

typos-write-changes:
	typos --write-changes

vhs:
	vhs < id-validator.tape

yamlfmt:
	yamlfmt $(shell find $$PWD -iname "*.yml") -gitignore_excludes

zhlint:
	if ! command -v zhlint >/dev/null 2>&1; then echo 'zhlint not found, installing...'; npm install -g zhlint; fi
	$(MAKE) zhlint-prototype

zhlint-fix:
	$(MAKE) zhlint-prototype ZHLINT_EXTRA=--fix

zhlint-prototype:
	zhlint {,docs/,docs/**/}*-zh_CN.md $(ZHLINT_EXTRA)

zizmor:
	zizmor .github/ --verbose
