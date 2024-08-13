package main

import (
	"codewars/remote"
)

func main() {

	rm := remote.Connect()

	rm.Storage.UpdateUrl("update")
}
