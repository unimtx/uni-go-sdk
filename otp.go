package uni

type OtpService struct {
	client *UniClient
}

type OtpSendParams struct {
	To				string
	Code			string
	Ttl				int
	Digits			int
	Intent			string
	Channel			string
	Signature		string
	TemplateId		string
}

type OtpVerifyParams struct {
	To				string
	Code			string
	Ttl				int
	Intent			string
}

func (service *OtpService) Send(params *OtpSendParams) (response *UniResponse, err error) {
	data := StructToMap(*params)
	return service.client.Request("otp.send", data)
}

func (service *OtpService) Verify(params *OtpVerifyParams) (response *UniResponse, err error) {
	data := StructToMap(*params)
	res, err := service.client.Request("otp.verify", data)

	if res != nil {
		res.Valid = res.Data["valid"].(bool)
	}

	return res, err
}
