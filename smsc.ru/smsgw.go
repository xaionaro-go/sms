package smscRu

import (
	"github.com/goodsign/gosmsc"
	"github.com/xaionaro-go/sms"
)

type smscru struct {
	impl *gosmsc.SenderFetcherImpl
}

func New(login, password string) (sms.GW, error) {
	var err error
	var gw smscru
	gw.impl, err = gosmsc.NewSenderFetcherImpl(&gosmsc.SmscClientOptions{
		User:     login,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return &gw, nil
}
func (gw smscru) Send(sender, destination, message string) error {
	_, err := gw.impl.Send(destination, message)
	return err
}
