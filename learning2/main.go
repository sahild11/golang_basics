package main

import (
	learning2 "learning2/importfolder"
	"log"
)

func main() {
	log.Println("main() of learning2/main.go")
	learning2.Talk()
	learning2.Drawpic(256, 256)
}
