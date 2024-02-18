package server

import "fmt"

type ErrorBase struct {
	field string
	msg   string
}

func (e ErrorBase) Error() string {
	return fmt.Sprintf("field: %v, msg: %v", e.field, e.msg)
}

type UdpError struct {
	field string
	msg   string
	ErrorBase
}

type PacketDataError struct {
	field string
	msg   string
	ErrorBase
}
