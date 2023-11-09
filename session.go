package mailin

import (
	"net"
	"strings"
)

/*
since smtp requests are not sent all at once
this is basically a struct to keep track of an
in-propgress smtp request
*/
type session struct {
	senderAddr       string
	hasSaidHello     bool
	helloFrom        string
	mailFrom         string
	rcptTos          []string
	hasBodyStarted   bool
	body             string
	hasBodyCompleted bool
}

func (s *session) toEmail() Email {
	return Email{
		s.helloFrom,
		net.IP(s.senderAddr),
		s.mailFrom,
		s.rcptTos,
		s.body,
	}
}

func (s *session) fillDefaultVals() {
	s.hasSaidHello = false
	s.helloFrom = ""
	s.mailFrom = ""
	s.rcptTos = []string{}
	s.hasBodyStarted = false
	s.body = ""
	s.hasBodyCompleted = false
}

func (s *session) isComplete() bool {
	return s.hasSaidHello &&
		s.hasBodyStarted &&
		s.hasBodyCompleted
}

/*
takes a request and delegates to sub functions
*/
func (s *session) handleRequest(req string, emailHandler EmailHandler) (resCode uint16, resMsg string) {
	if s.hasBodyStarted && !s.hasBodyCompleted {
		return s.handleBody(req)
	}
	req = strings.TrimSuffix(req, Crlf)
	cmd := findCmd(req)
	switch cmd {
	case CmdEhlo:
		return CodeNotImplemented, MsgCmdNotImplemented
	case CmdHelo:
		return s.handleHelo(req)
	case CmdMail:
		return s.handleMail(req)
	case CmdRcpt:
		return s.handleRcpt(req)
	case CmdData:
		return s.handleData(req)
	case CmdQuit:
		return s.handleQuit(req, emailHandler)
	case CmdRset:
		s.fillDefaultVals()
		return CodeOk, MsgOk
	case CmdVrfy:
		return CodeNotImplemented, MsgCmdNotImplemented
	case CmdNoop:
		return CodeOk, MsgOk
	case CmdTurn:
		return CodeNotImplemented, MsgCmdNotImplemented
	case CmdExpn:
		return CodeNotImplemented, MsgCmdNotImplemented
	case CmdHelp:
		return CodeNotImplemented, MsgCmdNotImplemented
	case CmdSend:
		return CodeNotImplemented, MsgCmdNotImplemented
	case CmdSaml:
		return CodeNotImplemented, MsgCmdNotImplemented
	}
	return CodeSyntaxErr, MsgSyntaxErr
}

// HELO<SP><domain><CRLF>
func (s *session) handleHelo(req string) (resCode uint16, resMsg string) {
	if s.hasSaidHello {
		return CodeInvalidSequence, MsgInvalidSequence
	} else {
		words := strings.Split(req, " ")
		if len(words) == 2 {
			s.helloFrom = strings.TrimSpace(words[1])
			s.hasSaidHello = true
			return CodeOk, MsgOk
		} else {
			return CodeSyntaxErr, MsgSyntaxErr
		}
	}
}

// MAIL<SP>FROM: <reverse-path><CRLF>
func (s *session) handleMail(req string) (resCode uint16, resMsg string) {
	if !s.hasSaidHello || s.mailFrom != "" {
		return CodeInvalidSequence, MsgInvalidSequence
	} else {
		sections := strings.Split(req, ":")
		if len(sections) == 2 && strings.ToUpper(sections[0]) == "MAIL FROM" {
			isEmailInLine, email := findEmailInLine(sections[1])
			if isEmailInLine {
				s.mailFrom = email
				return CodeOk, MsgOk
			} else {
				return CodeSyntaxErr, MsgSyntaxErr
			}
		} else {
			return CodeSyntaxErr, MsgSyntaxErr
		}
	}
}

// RCPT<SP>TO: <forward-path><CRLF>
func (s *session) handleRcpt(req string) (resCode uint16, resMsg string) {
	if s.mailFrom == "" || s.hasBodyStarted || !s.hasSaidHello {
		return CodeInvalidSequence, MsgInvalidSequence
	} else {
		sections := strings.Split(req, ":")
		if len(sections) == 2 && strings.ToUpper(sections[0]) == "RCPT TO" {
			isFound, email := findEmailInLine(sections[1])
			if isFound {

				// Check if rcpt is found
				if RcptVerifyHandler != nil &&
					!RcptVerifyHandler(email) {
					return CodeRcptNotFound, MsgRcptNotFound
				} else {
					s.rcptTos = append(s.rcptTos, email)
					return CodeOk, MsgOk
				}

			} else {
				return CodeSyntaxErr, MsgSyntaxErr
			}
		} else {
			return CodeSyntaxErr, MsgSyntaxErr
		}
	}
}

// DATA<CRLF>
func (s *session) handleData(req string) (resCode uint16, resMsg string) {
	if s.hasBodyStarted || len(s.rcptTos) == 0 || s.mailFrom == "" || !s.hasSaidHello {
		return CodeInvalidSequence, MsgInvalidSequence
	} else {
		s.hasBodyStarted = true
		return CodeStartMail, MsgStartMail
	}
}

func (s *session) handleBody(req string) (resCode uint16, resMsg string) {
	s.body += req
	if strings.HasSuffix(s.body, BodyFinish) {
		s.hasBodyCompleted = true
		return CodeOk, MsgOk
	} else {
		return CodeDontRespond, ""
	}
}

func (s *session) handleQuit(req string, emailHandler EmailHandler) (resCode uint16, resMsg string) {
	if s.isComplete() {
		emailHandler(s.toEmail()) // TODO decide how to handle return
	}
	return CodeBye, MsgBye
}
