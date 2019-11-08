package main

import "fmt"

func main() {
	fmt.Println("Creating WrebAssembly code from Go")
}

// ➜  chapter2 git:(master) ✗ GOOS=js GOARCH=wasm go build -o main.wasm to_wasm.go
// ➜  chapter2 git:(master) ✗ ls main.wasm
// main.wasm
// ➜  chapter2 git:(master) ✗ ls -alh  main.wasm
// -rwxr-xr-x  1 r  staff   2.2M 11  8 14:25 main.wasm
// ➜  chapter2 git:(master) ✗ file main.wasm
// main.wasm: WebAssembly (wasm) binary module version 0x1 (MVP)
