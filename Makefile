MOCKERY := mockery
SRC_ROOT := internal
MOCK_SUFFIX := mocks

.PHONY: mocks clean-mocks

# Generate mocks for all interfaces inside internal/*
mocks:
	@echo "🔍 Scanning interfaces in $(SRC_ROOT)..."
	@for dir in $(SRC_ROOT)/*; do \
		if [ -d $$dir ]; then \
			for iface in `grep -E '^type [A-Z][A-Za-z0-9_]+ interface' $$dir/*.go | awk '{print $$2}'`; do \
				echo "✨ Generating mock for $$iface in $$dir..."; \
				$(MOCKERY) --name=$$iface --dir=$$dir --output=$$dir/$(MOCK_SUFFIX) --structname=Mock$$iface --quiet; \
			done \
		fi \
	done

# Clean mocks from all modules
clean-mocks:
	@echo "🧹 Cleaning mocks..."
	@find $(SRC_ROOT) -type f -path "*/$(MOCK_SUFFIX)/Mock*.go" -delete


# === TARGET: Run All Tests ===
.PHONY: test
test:
	@echo "Running unit tests..."
	@go test ./internal/... -v
	@echo "Tests complete."

# === TARGET: Format ===
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@goimports -w .
	@echo "Formatting complete."

# === TARGET: All ===
.PHONY: all
all: mocks test