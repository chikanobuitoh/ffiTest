.DEFAULT_GOAL := help

.PHONY: help
help: ## このMakefileのヘルプを表示します

.PHONY: buildWin
buildWin: ## 共有ライブラリを作成します
	go build -buildmode=c-shared -o libTest.dll ./lib/main.go

.PHONY: buildMac
buildMac: ## 共有ライブラリを作成します
	go build -buildmode=c-shared -o libTest.dylib ./lib/main.go