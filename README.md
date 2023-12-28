# Unimatrix Go SDK

[![PkgGoDev](https://pkg.go.dev/badge/github.com/unimtx/uni-go-sdk)](https://pkg.go.dev/github.com/unimtx/uni-go-sdk) [![Release](https://img.shields.io/github/release/unimtx/uni-go-sdk.svg)](https://github.com/unimtx/uni-go-sdk/releases/latest) [![GitHub license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://github.com/unimtx/uni-go-sdk/blob/main/LICENSE)

The Unimatrix Go SDK provides convenient access to integrate communication capabilities into your Go applications using the Unimatrix HTTP API. The SDK provides support for sending SMS, 2FA verification, and phone number lookup.

## Getting started

Before you begin, you need an [Unimatrix](https://www.unimtx.com/) account. If you don't have one yet, you can [sign up](https://www.unimtx.com/signup?s=go.sdk.gh) for an Unimatrix account and get free credits to get you started.

## Documentation

Check out the documentation at [unimtx.com/docs](https://www.unimtx.com/docs) for a quick overview.

## Installation

The Unimatrix SDK for Golang uses Go Modules, which is available from the public [Github repository](https://github.com/unimtx/uni-go-sdk).

Run the following command to add `uni-go-sdk` as a dependency to your project:

```bash
go get github.com/unimtx/uni-go-sdk
```

## Usage

The following example shows how to use the Unimatrix Go SDK to interact with Unimatrix services.

### Initialize a client

```go
package main

import (
    "fmt"
    "github.com/unimtx/uni-go-sdk"
)

func main() {
    client := uni.NewClient()
}
```

or you can configure your credentials by environment variables:

```sh
export UNIMTX_ACCESS_KEY_ID=your_access_key_id
export UNIMTX_ACCESS_KEY_SECRET=your_access_key_secret
```

### Send SMS

Send a text message to a single recipient.

```go
package main

import (
    "fmt"
    "github.com/unimtx/uni-go-sdk"
)

func main() {
    client := uni.NewClient()

    res, err := client.Messages.Send(&uni.MessageSendParams{
        To: "+1206880xxxx",  // in E.164 format
        Text: "Your verification code is 2048.",
    })
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(res)
    }
}
```

### Send verification code

Send a one-time passcode (OTP) to a recipient. The following example will send a automatically generated verification code to the user.

```go
package main

import (
    "fmt"
    "github.com/unimtx/uni-go-sdk"
)

func main() {
    client := uni.NewClient()

    res, err := client.Otp.Send(&uni.OtpSendParams{
        To: "+1206880xxxx",
    })
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(res)
    }
}
```

### Check verification code

Verify the one-time passcode (OTP) that a user provided. The following example will check whether the user-provided verification code is correct.

```go
package main

import (
    "fmt"
    "github.com/unimtx/uni-go-sdk"
)

func main() {
    client := uni.NewClient()

    res, err := client.Otp.Verify(&uni.OtpVerifyParams{
        To: "+1206880xxxx",
        Code: "123456", // the code user provided
    })
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(res.Valid)
    }
}
```

## Reference

### Other Unimatrix SDKs

To find Unimatrix SDKs in other programming languages, check out the list below:

- [Java](https://github.com/unimtx/uni-java-sdk)
- [Node.js](https://github.com/unimtx/uni-node-sdk)
- [Python](https://github.com/unimtx/uni-python-sdk)
- [PHP](https://github.com/unimtx/uni-php-sdk/)
- [Ruby](https://github.com/unimtx/uni-ruby-sdk)
- [.NET](https://github.com/unimtx/uni-dotnet-sdk)

## License

This library is released under the [MIT License](https://github.com/unimtx/uni-go-sdk/blob/main/LICENSE).
