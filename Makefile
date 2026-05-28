BUILD_DIR := build
APP_NAME  := infocli
VERSION   := $(shell grep 'gVersion' version.go | grep -oP 'v[\d.]+')

LDFLAGS := -ldflags '-extldflags "-static" -s -w -X main.gVersion=$(VERSION)'

DEV_OUTPUT        := $(BUILD_DIR)/$(APP_NAME)-dev
LINUX_OUTPUT      := $(BUILD_DIR)/$(APP_NAME)
MACOS_AMD64_OUTPUT := $(BUILD_DIR)/$(APP_NAME)-darwin-amd64
MACOS_ARM64_OUTPUT := $(BUILD_DIR)/$(APP_NAME)-darwin-arm64

.PHONY: all dev linux darwin-amd64 darwin-arm64 install release test clean fmt tidy

all: linux darwin-amd64 darwin-arm64

# local quick build for current platform
dev: $(BUILD_DIR)
	go build -o $(DEV_OUTPUT) .

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

linux: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a $(LDFLAGS) -o $(LINUX_OUTPUT)
	@which upx >/dev/null 2>&1 && upx $(LINUX_OUTPUT) || echo "upx not found, skipping compression"

darwin-amd64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a $(LDFLAGS) -o $(MACOS_AMD64_OUTPUT)

darwin-arm64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a $(LDFLAGS) -o $(MACOS_ARM64_OUTPUT)

install:
	@OS=$$(uname -s); ARCH=$$(uname -m); \
	if [ "$$OS" = "Darwin" ] && [ "$$ARCH" = "arm64" ]; then \
		$(MAKE) darwin-arm64; \
		sudo cp $(MACOS_ARM64_OUTPUT) /usr/local/bin/$(APP_NAME); \
	elif [ "$$OS" = "Darwin" ]; then \
		$(MAKE) darwin-amd64; \
		sudo cp $(MACOS_AMD64_OUTPUT) /usr/local/bin/$(APP_NAME); \
	else \
		$(MAKE) linux; \
		sudo cp $(LINUX_OUTPUT) /usr/local/bin/$(APP_NAME); \
	fi
	@echo "Installed $(APP_NAME) $(VERSION) to /usr/local/bin/$(APP_NAME)."

release: all
	git tag $(VERSION)
	git push origin $(VERSION)
	@echo "Released $(VERSION)."

test: dev
	@DB=/tmp/$(APP_NAME)-test.db; BIN=$(DEV_OUTPUT); \
	rm -f $$DB; \
	echo "--- add ---"; \
	$$BIN -f $$DB a key1 "value1"; \
	$$BIN -f $$DB a key2 "value2"; \
	echo "--- duplicate ---"; \
	$$BIN -f $$DB a key1 "dup"; \
	echo "--- query name ---"; \
	$$BIN -f $$DB q key; \
	echo "--- query id ---"; \
	$$BIN -f $$DB q id 1 -d; \
	echo "--- update data ---"; \
	$$BIN -f $$DB u data -i 1 "updated"; \
	echo "--- query after update ---"; \
	$$BIN -f $$DB q id 1; \
	echo "--- delete -i after subcommand ---"; \
	$$BIN -f $$DB d -i 2; \
	echo "--- delete -i before subcommand (id=1) ---"; \
	$$BIN -f $$DB -i 1 d; \
	echo "--- count ---"; \
	$$BIN -f $$DB c; \
	rm -f $$DB; \
	echo "--- test done ---"; \
	echo ""; \
	echo "=== encrypt test ==="; \
	EDB=/tmp/$(APP_NAME)-enc-test.db; \
	rm -f $$EDB; \
	echo "--- add with key ---"; \
	$$BIN -f $$EDB -k secret a server1 "192.168.1.1"; \
	$$BIN -f $$EDB -k secret a server2 "192.168.1.2"; \
	$$BIN -f $$EDB -k secret a other "not a server"; \
	echo "--- fuzzy query with key (decrypted) ---"; \
	$$BIN -f $$EDB -k secret q name server; \
	echo "--- fuzzy query without key (raw encrypted) ---"; \
	$$BIN -f $$EDB q name server; \
	echo "--- query with wrong key (raw encrypted) ---"; \
	$$BIN -f $$EDB -k wrongkey q name server; \
	rm -f $$EDB; \
	echo "--- encrypt test done ---"

clean:
	rm -rf $(BUILD_DIR)

fmt:
	go fmt ./...

tidy:
	go mod tidy
