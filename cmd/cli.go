package main

import (
	"log"

	"github.com/dollarkillerx/dog"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	dog.Execute()
}
