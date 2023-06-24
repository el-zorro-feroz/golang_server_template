package main

import cmd "cleanarch/src/cmd/server"

func main() {
	if err := cmd.RunServer(); err != nil {
		panic(err)
	}
}
