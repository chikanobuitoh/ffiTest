.DEFAULT_GOAL := help

.PHONY: help
help: ## このMakefileのヘルプを表示します

.PHONY: setup
setup: ## 初期設定を行います
	go mod tidy
	npm install ffi-napi
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: makeproto
makeproto:
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/test.proto

.PHONY: buildWin
buildWin: ## 共有ライブラリを作成します
	go build -buildmode=c-shared -o libTest.dll ./lib/main.go

.PHONY: buildMac
buildMac: ## 共有ライブラリを作成します
	go build -buildmode=c-shared -o libTest.dylib ./lib/main.go