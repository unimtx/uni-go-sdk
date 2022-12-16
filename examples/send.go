package main

import (
	"fmt"

	"github.com/unimtx/uni-go-sdk"
)

func main() {
	client := uni.NewClient("your access key id", "your access key secret")

	// send a text message to a single recipient
	message := client.Messages.BuildMessage()
	message.SetTo("your phone number") // in E.164 format
	message.SetSignature("your sender name")
	message.SetContent("Your verification code is 2048.")

	res, err := client.Send(message)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
