package main

import (
	"fmt"
	"log"
)

func main() {
	builder := newNotificationBuilder()

	builder.SetTitle("Title")

	builder.SetMessage("Message")
	builder.SetImage("Image")
	builder.SetIcon("Icon")
	builder.SetPriority(1)
	builder.SetType("Type")

	notification, err := builder.Build()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(notification)
}
