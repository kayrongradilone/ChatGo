package main

import (
	"net"
)


type sala struct {
	nome string
	membros map[net.Addr]*cliente
}

func (r *sala) broadcast(sender *cliente, msg string){
	for addr, m := range r.membros {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}