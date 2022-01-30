package port

import (
	"fmt"
	"net"
	"time"
)

func Scan(ip string, port uint) bool {
	s := fmt.Sprintf("%s:%d", ip, port)
	d := net.Dialer{Timeout: 500 * time.Millisecond}
	conn, err := d.Dial("tcp", s)
	if err != nil {
		return false
	}
	conn.Close()

	return true
}
