package mailin

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn, emailHandler EmailHandler) {
	defer conn.Close()
	senderAddr := conn.RemoteAddr()

	// Greet the client
	readyMsg := fmt.Sprintf("%d SMTP SERVICE READY%s", CodeReady, Crlf)
	conn.Write([]byte(readyMsg))

	// Initialize smtp session
	session := session{}
	session.senderAddr = senderAddr.String()
	session.fillDefaultVals()

	// For each tcp packet (i.e each request)
	for {
		pktBuffer := make([]byte, PktBufferSizeLimit)

		pktSize, err := conn.Read(pktBuffer)
		if err != nil {
			logMsg := fmt.Sprintf("ERROR: failed connection with %s (%s)", conn.RemoteAddr(), err)
			LogHandler(logMsg)
			break
		} else if pktSize >= PktBufferSizeLimit {
			errMsg := fmt.Sprintf("%d PACKET SIZE MUST BE <%d%s", CodeReady, PktBufferSizeLimit, Crlf)
			conn.Write([]byte(errMsg))
			break
		}

		req := string(pktBuffer[:pktSize])

		resCode, resMsg := session.handleRequest(req, emailHandler)
		res := fmt.Sprintf("%d %s%s", resCode, resMsg, Crlf)
		if resCode != CodeDontRespond {
			_, _ = conn.Write([]byte(res))
			if resCode == CodeBye {
				return
			}
		}
	}
}
