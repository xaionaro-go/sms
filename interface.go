package sms

type GW interface {
	Send(sender, destination, message string) error
}
