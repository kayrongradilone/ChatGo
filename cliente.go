package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type cliente struct {
	conn net.Conn
	nome string
	sala *sala
	comandos chan<- comando

 }

 func (c *cliente) readInput(){
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return 
		}
		msg = strings.Trim(msg, "\r\n" )
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		
		switch cmd {
		case "/nome":
			c.comandos <- comando {
				id: CMD_NOME,
				cliente: c,
				args: args,
			}
		case "/entrar":
			c.comandos <- comando {
				id: CMD_ENTRAR,
				cliente: c,
				args: args,
			}
		case "/salas":
			c.comandos <- comando {
				id: CMD_SALAS,
				cliente: c,
				args: args,
			}
		case "/msg":
			c.comandos <- comando {
				id: CMD_MSG,
				cliente: c,
				args: args,
			}
		case "/sair":
			c.comandos <- comando {
				id: CMD_SAIR,
				cliente: c,
				args: args,
			}
		default:
			c.err(fmt.Errorf("Comando nao reconhecido %s", cmd))
		}
	}
 }

 func (c *cliente) err(err error){
	c.conn.Write([]byte("ERR" + err.Error() + "\n"))
 }
 func (c *cliente) msg(msg string){
	c.conn.Write([]byte(">" + msg + "\n"))
 }