# === CONFIG ===
NODE_MASTER_DIR=apps/node-master
NODE_CLIENT_DIR=apps/node-client
CLIENT_DIR=apps/client


# === BUILD ===
build-master:
	go build -o bin/node-master $(NODE_MASTER_DIR)/cmd/main.go

build-client:
	go build -o bin/node-client $(NODE_CLIENT_DIR)/cmd/main.go

build-desktop:
	cd $(CLIENT_DIR) && wails build

build-all: build-master build-client

# === RUN ===
run-master:
	go run $(NODE_MASTER_DIR)/cmd/main.go

run-client:
	go run $(NODE_CLIENT_DIR)/cmd/main.go

run-desktop:
	cd $(CLIENT_DIR) && wails dev

# === CLEAN ===
clean:
	rm -rf bin/*

# === UTILS ===
format:
	gofmt -w $(NODE_MASTER_DIR) $(NODE_CLIENT_DIR) libs/

tidy:
	cd $(NODE_MASTER_DIR) && go mod tidy
	cd $(NODE_CLIENT_DIR) && go mod tidy
	cd libs/logger && go mod tidy

.PHONY: build-master build-client build-all run-master run-client run-desktop clean format tidy
