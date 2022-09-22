package main

import (
	"fmt"
	"lyrics-extraction/services"
)

func main() {
	//command := cmd.NewLyricsCmd()
	//err := command.Execute()
	//if err != nil {
	//	log.Fatalf("Lyrics extraction failed, err : %v \n", err)
	//}
	testGetResult()
}

func testUpload() {
	filePath := `D:\resources\Never_on_Sunday.mp3`
	duration := 162
	_, err := services.NewTransferService().Upload(filePath, duration)
	if err != nil {
		fmt.Println(err)
	}
}

func testGetResult() {
	_, err := services.NewTransferService().GetResult("DKHJQ202209221722092217249EC77BDD00041")
	if err != nil {
		fmt.Println(err)
	}
}
