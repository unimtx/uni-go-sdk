package main

import (
	"fmt"

	"github.com/unimtx/uni-go-sdk"
)

// send a verification code to a recipient
func sendOtp(client *uni.UniClient) {
	res, err := client.Otp.Send(&uni.OtpSendParams{
		To: "your phone number",  // in E.164 format
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

// verify a verification code
func verifyOtp(client *uni.UniClient) {
	res, err := client.Otp.Verify(&uni.OtpVerifyParams{
		To: "your phone number",  // in E.164 format
		Code: "the code you received",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Valid)
	}
}

func main() {
	client := uni.NewClient("your access key id", "your access key secret")

	sendOtp(client)

	verifyOtp(client)
}
