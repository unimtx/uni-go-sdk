package main

import (
	"fmt"

	"github.com/unimtx/uni-go-sdk"
)

func main() {
	client := uni.NewClient("your access key id", "your access key secret")

	res, err := client.Messages.Send(&uni.MessageSendParams{
		To: "your phone number",  // in E.164 format
		Signature: "your sender name",
		Content: "Your verification code is 2048.",
	})
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
