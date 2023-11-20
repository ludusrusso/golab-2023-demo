PROJECT_DIR = ${PWD}

## Location to install dependencies to
LOCALBIN ?= $(PROJECT_DIR)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

# go-install-tool will 'go install' any package $2 and install it to $1.
define go-install-tool
@[ -f $(1) ] || { \
set -e ;\
echo "Installing $(2)" ;\
echo "GOBIN=$(LOCALBIN) go install $(2)" ;\
GOBIN=$(LOCALBIN) go install $(2) ;\
}
endef

SQLC ?= $(LOCALBIN)/sqlc
SQLC_VERSION ?= v1.23.0

.PHONY: sqlc
sqlc: $(LOCALBIN)
	$(call go-install-tool,$(SQLC),github.com/sqlc-dev/sqlc/cmd/sqlc@$(SQLC_VERSION))

DBMATE ?= $(LOCALBIN)/dbmate
DBMATE_VERSION ?= v2.7.0

.PHONY: dbmate
dbmate: $(LOCALBIN)
	$(call go-install-tool,$(DBMATE),github.com/amacneil/dbmate/v2@$(DBMATE_VERSION))
	
BUF ?= $(LOCALBIN)/buf
BUF_VERSION ?= v1.28.0

.PHONY: buf
buf: $(LOCALBIN)
	$(call go-install-tool,$(BUF),github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION))

generate: sqlc buf
	$(SQLC) generate
	$(BUF) generate

dbup: dbmate
	$(DBMATE) up

test:
	go clean -testcache
	go test -v ./...