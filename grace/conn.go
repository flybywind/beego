package grace

import "fmt"
import "net"

type graceConn struct {
	net.Conn
	server *graceServer
}

func (c graceConn) Close() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			fmt.Println("unexpected error:", err)
		}
	}()
	c.server.wg.Done()
	err = c.Conn.Close()
	return
}
