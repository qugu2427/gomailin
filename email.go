package mailin

import "net"

type Email struct {
	Helo         string
	MailFromAddr net.IP
	MailFrom     string
	Recipients   []string
	Body         string
}
