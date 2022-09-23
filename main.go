package main

import (
	"fmt"
	"log"
	"lyrics-extraction/cmd"
	"lyrics-extraction/services"
)

func main() {
	command := cmd.NewLyricsCmd()
	err := command.Execute()
	if err != nil {
		log.Fatalf("Lyrics extraction failed, err : %v \n", err)
	}
	//services.ParseResult()
}

func testUpload() {
	filePath := `D:\resources\Never_on_Sunday.mp3`
	duration := 162000
	_, err := services.NewTransferService().Upload(filePath, duration)
	if err != nil {
		fmt.Println(err)
	}
}

func testGetResult() {
	_, err := services.NewTransferService().GetResult("DKHJQ2022092314220923145309AAEA5B00079")
	if err != nil {
		fmt.Println(err)
	}
}
