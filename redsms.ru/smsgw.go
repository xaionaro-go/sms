package redsmsRu

import (
	"github.com/kodix/redSmsClient"
	"github.com/xaionaro-go/sms"
)

type redsmsru struct {
	impl *redSmsClient.RedSmsClient
}

func New(login, apiKey string) (sms.GW, error) {
	return &redsmsru{
		impl: &redSmsClient.RedSmsClient{
			Url:    "https://cp.redsms.ru/api",
			Login:  login,
			ApiKey: apiKey,
			Sender: "default",
		},
	}, nil
}

func (gw redsmsru) Send(sender, destination, message string) error {
	gw.impl.Sender = sender
	_, err := gw.impl.Send(destination, message)
	return err
}
