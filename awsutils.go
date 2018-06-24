package main

import (
	"awsutils/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
