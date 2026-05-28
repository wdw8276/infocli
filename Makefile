BUILD_DIR := build
APP_NAME  := infocli
VERSION   := $(shell grep 'gVersion' version.go | grep -oP 'v[\d.]+')

LDFLAGS := -ldflags '-extldflags "-static" -s -w -X main.gVersion=$(VERSION)'

WINDOWS_OUTPUT    := $(BUILD_DIR)/$(APP_NAME).exe
LINUX_OUTPUT      := $(BUILD_DIR)/$(APP_NAME)-linux
MACOS_AMD64_OUTPUT := $(BUILD_DIR)/$(APP_NAME)-darwin-amd64
MACOS_ARM64_OUTPUT := $(BUILD_DIR)/$(APP_NAME)-darwin-arm64

.PHONY: all dev windows linux darwin-amd64 darwin-arm64 install clean fmt tidy

all: windows linux darwin-amd64 darwin-arm64

# 本地快速构建（当前平台）
dev:
	go build -o $(APP_NAME) .

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

windows: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a $(LDFLAGS) -o $(WINDOWS_OUTPUT)

linux: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a $(LDFLAGS) -o $(LINUX_OUTPUT)
	@which upx >/dev/null 2>&1 && upx $(LINUX_OUTPUT) || echo "upx not found, skipping compression"

darwin-amd64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a $(LDFLAGS) -o $(MACOS_AMD64_OUTPUT)

darwin-arm64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a $(LDFLAGS) -o $(MACOS_ARM64_OUTPUT)

install: linux
	@echo "Installing $(APP_NAME) $(VERSION) to /usr/local/bin/$(APP_NAME)..."
	@sudo cp $(LINUX_OUTPUT) /usr/local/bin/$(APP_NAME)
	@echo "Done."

clean:
	rm -rf $(BUILD_DIR) $(APP_NAME)

fmt:
	go fmt ./...

tidy:
	go mod tidy
