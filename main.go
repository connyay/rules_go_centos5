package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fmt.Printf("Hello, %s_%s!\nuname=%q\ngo=%q\nip=%q\n",
		runtime.GOOS, runtime.GOARCH, mustGetUnameOutput(), runtime.Version(), mustGetIpAddress(),
	)
}

func mustGetUnameOutput() string {
	out, err := exec.Command("uname", "-a").CombinedOutput()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func mustGetIpAddress() string {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		panic(err)
	}
	ip, _ := io.ReadAll(res.Body)
	return string(ip)
}
