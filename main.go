package main

import (
	"fmt"
)

var (
	PROGRAM     string
	VERSION     string
	BUILD       string
	COMMIT_SHA1 string
)

func init() {
	fmt.Printf("-----------------------------------------------------\nProgram: %s\nVersion: %s\nBuild: %s\nCommit_sha1: %s\n-----------------------------------------------------\n", PROGRAM, VERSION, BUILD, COMMIT_SHA1)
}

func main() {

	cli := CLI{}
	cli.Run()
}
