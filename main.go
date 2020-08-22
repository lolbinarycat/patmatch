// patmatch is a tool for pattern matching in bash scripts
package main

import "os"

const Shell = "/bin/bash"

func main() {
	exec.Setenv("foo","1")
}
