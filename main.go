package main

import (
	"envdir/envdir"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	flag.Parse()
	args := flag.Args()
	if flag.NArg() == 0 || len(args) < 2 {
		flag.Usage()
		os.Exit(2)
	}

	pathToEnv := args[0]
	cmd := args[1:]

	env, err := envdir.ReadDir(pathToEnv)
	if err != nil {
		log.Fatal(err)
	}
	out, code := envdir.RunCmd(cmd, env)
	fmt.Println(out)
	os.Exit(code)
}
