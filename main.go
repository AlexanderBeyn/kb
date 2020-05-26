package main

import (
	"github.com/AlexanderBeyn/kb/cmd"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()
}
