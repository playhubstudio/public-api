API_GEN_RUN := go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
LINT_RUN:= go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest
TEST_REPORTS_RUN := go run github.com/jstemmer/go-junit-report/v2@latest


V10_API_PATH:=./api/v1.0
V10_API_GEN_IMPORT_COMMON = "-import-mapping=../common-components.yaml:github.com/playhubstudio/public-api/api/v1.0"

V11_API_PATH:=./api/v1.1
V11_API_GEN_IMPORT_COMMON = "-import-mapping=../common-components.yaml:github.com/playhubstudio/public-api/api/v1.1"

V12_API_PATH:=./api/v1.2
V12_API_GEN_IMPORT_COMMON = "-import-mapping=../common-components.yaml:github.com/playhubstudio/public-api/api/v1.2"

.PHONY: gen-api
gen-api: gen-api-v12

# gen-api-v10:
# 	$(API_GEN_RUN) -package apiv1 -generate types,skip-prune,spec $(V10_API_PATH)/common-components.yaml > $(V10_API_PATH)/common-components.gen.go
# 	$(API_GEN_RUN) $(V10_API_GEN_IMPORT_COMMON) -package coreapiv1 -generate server,types,skip-prune,spec,client $(V10_API_PATH)/core/core-api.yaml > $(V10_API_PATH)/core/core-api.gen.go
# 	$(API_GEN_RUN) $(V10_API_GEN_IMPORT_COMMON) -package playhubintegrationapiv1 -generate types,skip-prune,spec,client $(V10_API_PATH)/playhubintegration/playhub-integration-api.yaml > $(V10_API_PATH)/playhubintegration/playhub-integration-api.gen.go

# gen-api-v11:
# 	$(API_GEN_RUN) -package apiv1 -generate types,skip-prune,spec $(V11_API_PATH)/common-components.yaml > $(V11_API_PATH)/common-components.gen.go
# 	$(API_GEN_RUN) $(V11_API_GEN_IMPORT_COMMON) -package coreapiv1 -generate server,types,skip-prune,spec,client $(V11_API_PATH)/core/core-api.yaml > $(V11_API_PATH)/core/core-api.gen.go
# 	$(API_GEN_RUN) $(V11_API_GEN_IMPORT_COMMON) -package playhubintegrationapiv1 -generate types,skip-prune,spec,client $(V11_API_PATH)/playhubintegration/playhub-integration-api.yaml > $(V11_API_PATH)/playhubintegration/playhub-integration-api.gen.go

gen-api-v12:
	$(API_GEN_RUN) -package apiv1 -generate types,skip-prune,spec $(V12_API_PATH)/common-components.yaml > $(V12_API_PATH)/common-components.gen.go
	$(API_GEN_RUN) $(V12_API_GEN_IMPORT_COMMON) -package coreapiv1 -generate server,types,skip-prune,spec,client $(V12_API_PATH)/core/core-api.yaml > $(V12_API_PATH)/core/core-api.gen.go
	$(API_GEN_RUN) $(V12_API_GEN_IMPORT_COMMON) -package playhubintegrationapiv1 -generate types,skip-prune,spec,client $(V12_API_PATH)/playhubintegration/playhub-integration-api.yaml > $(V11_API_PATH)/playhubintegration/playhub-integration-api.gen.go

.PHONY:publish-docs
publish-docs:
	# cd api/v1.0/core && rdme openapi core-api.yaml --version=v1.0 --key=$(PLAYHUB_README_API_KEY)  --id=65bc99b3a8f20d00395a35e7 --update
	# cd api/v1.0/playhubintegration && rdme openapi playhub-integration-api.yaml --version=v1.0 --key=$(PLAYHUB_README_API_KEY) --id=65bc99c69969910010b1426c --update
	# cd docs/v1.0 && rdme docs . --version=v1.0 --key=$(PLAYHUB_README_API_KEY) 
	
	# cd api/v1.1/core && rdme openapi core-api.yaml --version=v1.1 --key=$(PLAYHUB_README_API_KEY)  --id=671f27705c56c90011464dd6 --update
	# cd api/v1.1/playhubintegration && rdme openapi playhub-integration-api.yaml --version=v1.1 --key=$(PLAYHUB_README_API_KEY) --id=671f277f38a4f0001eebd066 --update
	# cd docs/v1.1 && rdme docs . --version=v1.1 --key=$(PLAYHUB_README_API_KEY) 

	cd $(V12_API_PATH)/core && rdme openapi core-api.yaml --version=v1.2 --key=$(PLAYHUB_README_API_KEY)  --id=67989aa203e47000307149fe --update
	cd $(V12_API_PATH)/playhubintegration && rdme openapi playhub-integration-api.yaml --version=v1.2 --key=$(PLAYHUB_README_API_KEY) --id=67989aabc6cdcd005bcfb973 --update
	cd docs/v1.2 && rdme docs . --version=v1.2 --key=$(PLAYHUB_README_API_KEY) 

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


