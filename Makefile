API_GEN_RUN := go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
LINT_RUN:= go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest
TEST_REPORTS_RUN := go run github.com/jstemmer/go-junit-report/v2@latest


V1_API_PATH:=./api/v1
V1_API_GEN_IMPORT_COMMON = "-import-mapping=../common-components.yaml:github.com/playhubstudio/public-api/api/v1"

.PHONY: gen-api
gen-api:
	$(API_GEN_RUN) -package apiv1 -generate types,skip-prune,spec $(V1_API_PATH)/common-components.yaml > $(V1_API_PATH)/common-components.gen.go
	$(API_GEN_RUN) $(V1_API_GEN_IMPORT_COMMON) -package coreapiv1 -generate server,types,skip-prune,spec,client $(V1_API_PATH)/core/core-api.yaml > $(V1_API_PATH)/core/core-api.gen.go
	$(API_GEN_RUN) $(V1_API_GEN_IMPORT_COMMON) -package playhubintegrationapiv1 -generate types,skip-prune,spec,client $(V1_API_PATH)/playhubintegration/playhub-integration-api.yaml > $(V1_API_PATH)/playhubintegration/playhub-integration-api.gen.go

.PHONY:publish-docs
publish-docs:
	cd api/v1/core && rdme openapi core-api.yaml --version=v1.0 --key=$(PLAYHUB_README_API_KEY)  --id=65bc99b3a8f20d00395a35e5 --update
	cd api/v1/playhubintegration && rdme openapi playhub-integration-api.yaml --version=v1.0 --key=$(PLAYHUB_README_API_KEY) --id=65bc99c69969910010b1426a --update
	cd docs && rdme docs . --version=v1.0 --key=$(PLAYHUB_README_API_KEY)

.PHONY:lint
lint:
	$(LINT_RUN) cache clean
	$(LINT_RUN) run -v

.PHONY:build
build:
	go build ./...

.PHONY: test-ci
test-ci:
	mkdir test-reports || true
	go test -v -count 1 ./... 2>&1 | $(TEST_REPORTS_RUN) -set-exit-code > test-reports/report.xml

