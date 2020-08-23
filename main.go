// patmatch is a tool for pattern matching in bash scripts
package main

import (
	"os"
	"os/exec"
	"log"
)

const Shell = "/bin/bash"

//go:generate peg grammar.peg

func main() {
	src := `[$a]
	echo "got $a"
	exit 0
`
	p := Prog{
		Cmd:&exec.Cmd{
			Path:Shell,
			Stdout:os.Stdout,
		},
		MStat: NewMatchStatus(1),
		Args:[]string{"t"},
		Buffer:src}
	var err error
	p.In, err = p.Cmd.StdinPipe()
	fatal(err)
	logln("starting shell")
	err = p.Cmd.Start()	
	fatal(err)
	logln("initalizing parser")
	p.Init()
	fatal(err)
	logln("parsing")
	err = p.Parse()
	fatal(err)
	logln("executing")
	p.Execute()
	logln("done")
}

func fatal(err error) {
	if err != nil {panic(err)}
} 

func logln(args ...interface{}) {
	log.Println(args...)
}
