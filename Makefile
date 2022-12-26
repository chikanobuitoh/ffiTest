.DEFAULT_GOAL := help

.PHONY: help
help: ## このMakefileのヘルプを表示します

.PHONY: setup
setup: ## 初期設定を行います
	go mod tidy
	npm install ffi-napi

.PHONY: buildWin
buildWin: ## 共有ライブラリを作成します
	go build -buildmode=c-shared -o libTest.dll ./lib/main.go

.PHONY: buildMac
buildMac: ## 共有ライブラリを作成します
	go build -buildmode=c-shared -o libTest.dylib ./lib/main.go