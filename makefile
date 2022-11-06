.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o static/template.wasm
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" static/wasm_exec.js