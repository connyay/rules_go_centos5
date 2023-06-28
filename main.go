package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Hello, %s_%s!\nuname=%q\ngo=%q\nip=%q\n",
		runtime.GOOS, runtime.GOARCH, mustGetUnameOutput(), runtime.Version(), mustGetIpAddress(),
	)
	fmt.Println()
	fmt.Printf("File Contents:\n%s\n", mustGetFileContents())

	fmt.Println()
	fmt.Printf("Temp Contents:\n%s\n", writeAndReadTempfile())

	fmt.Println()
	fmt.Println("Testing tcp server/client")
	addr := startLocalServer()
	readWriteClient(addr)
}

func mustGetUnameOutput() string {
	out, err := exec.Command("uname", "-a").CombinedOutput()
	checkErr(err)
	return strings.TrimSpace(string(out))
}

func mustGetIpAddress() string {
	res, err := http.Get("https://api.myip.com")
	checkErr(err)
	ip, _ := io.ReadAll(res.Body)
	return string(ip)
}

func mustGetFileContents() string {
	content, err := os.ReadFile("example.txt")
	checkErr(err)
	return string(content)
}

func writeAndReadTempfile() string {
	tempFile, err := os.CreateTemp("", "example-*")
	checkErr(err)
	_, err = tempFile.WriteString(time.Now().String())
	checkErr(err)
	if err := tempFile.Close(); err != nil {
		panic(err)
	}
	fmt.Printf("Temp File: %q\n", tempFile.Name())
	content, err := os.ReadFile(tempFile.Name())
	checkErr(err)
	return string(content)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
