package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	out, _ := exec.Command("uname", "-a").CombinedOutput()
	fmt.Printf("Hello, %s_%s!\nuname=%q\ngo=%q\n", runtime.GOOS, runtime.GOARCH, strings.TrimSpace(string(out)), runtime.Version())
}
