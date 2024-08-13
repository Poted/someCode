package main

import (
	"codewars/remote"
	"fmt"
)

func main() {

	rm := remote.Connect()

	rm.Storage.FileID = 2
	strg := rm.Storage

	newfile, err := strg.FileDownload(20)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("newfile: %v\n", newfile)

	fmt.Printf("strg.FileID: %v\n", strg.FileID)

	conn := rm.Connection

	fmt.Printf("conn.URL: %v\n", conn.URL())
	fmt.Printf("conn.URLWithID(30): %v\n", conn.UpdateUrl("some").URL())
	fmt.Printf("conn.URL: %v\n", conn.URL())
}
