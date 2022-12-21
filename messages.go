package uni

import (
	"regexp"
)

type MessageService struct {
	client *UniClient
}

type MessageSendParams struct {
	To				string
	Signature		string
	TemplateId		string
	TemplateData	map[string]string
	Content			string
	Text			string
}

func (service *MessageService) Send(params *MessageSendParams) (response *UniResponse, err error) {
	data := StructToMap(*params)
	data["to"] = regexp.MustCompile(`\s*,\s*`).Split(params.To, -1)
	return service.client.Request("sms.message.send", data)
}
