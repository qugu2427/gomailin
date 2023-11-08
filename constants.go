package mailin

import "net"

type EmailHandler = func(Email)
type ConnHandler = func(net.Conn, EmailHandler)

var (
	PktBufferSizeLimit int = 1025
)

const (
	Crlf       string = "\r\n"
	BodyFinish string = "\r\n.\r\n"
)

const (
	CodeDontRespond     uint16 = 0 // this is not an actual SMTP code
	CodeReady           uint16 = 220
	CodeBye             uint16 = 221
	CodeOk              uint16 = 250
	CodeStartMail       uint16 = 354
	CodeSyntaxErr       uint16 = 500
	CodeNotImplemented  uint16 = 502
	CodeInvalidSequence uint16 = 503
)

const (
	MsgCmdNotImplemented string = "COMMAND NOT IMPLEMENTED"
	MsgOk                string = "OK"
	MsgSyntaxErr         string = "SYNTAX ERROR"
	MsgInvalidSequence   string = "INVALID SEQUENCE"
	MsgStartMail         string = "START MAIL"
	MsgBye               string = "GOODBYE"
)

const (
	CmdHelo string = "HELO"
	CmdEhlo string = "EHLO"
	CmdMail string = "MAIL"
	CmdRcpt string = "RCPT"
	CmdData string = "DATA"
	CmdQuit string = "QUIT"
	CmdRset string = "RSET"
	CmdVrfy string = "VRFY"
	CmdNoop string = "NOOP"
	CmdTurn string = "TURN"
	CmdExpn string = "EXPN"
	CmdHelp string = "HELP"
	CmdSend string = "SEND"
	CmdSaml string = "SAML"
)
