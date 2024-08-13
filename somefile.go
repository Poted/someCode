package main

import (
	"codewars/remote"
	"fmt"
	"time"
)

func main() {

	rm := remote.Connect()

	fmt.Printf("rm.Connection.Url: %v\n", rm.Connection.Url)

	fmt.Printf("rm.Connection.ConnectionURLWithID(\"20\"): %v\n", rm.Connection.URLWithID("20").Url)

	time.Sleep(4 * time.Second)

}
