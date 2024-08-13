package main

import (
	"codewars/remote"
	"time"
)

func main() {

	rm := remote.Connect()

	fieldConnType := rm.Connection.Url
	childStructMethod := rm.Connection.URLWithID("connection url")
	parentStructMethod := rm.Connection.OtherFunction(12)

	parentMethodAvailableForEveryType := rm.Images.OtherFunction(2)
	fieldImagesType := rm.Images.Url
	encapsulatedMethodHiddenUsingField := rm.Images.Download(233)

	time.Sleep(4 * time.Second)

}
