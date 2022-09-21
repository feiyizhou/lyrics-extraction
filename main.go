package main

import (
	"log"
	"lyrics-extraction/cmd"
)

func main() {
	command := cmd.NewLyricsCmd()
	err := command.Execute()
	if err != nil {
		log.Fatalf("Lyrics extraction failed, err : %v \n", err)
	}
}
