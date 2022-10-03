package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)


type server struct {
	salas map[string]*sala
	comandos chan comando
}

func newServer() *server {
	return &server{
		salas: make(map[string]*sala),
		comandos: make(chan comando),
	 }
}

func (s *server) run(){
	for cmd := range s.comandos {
		switch cmd.id {
		case CMD_NOME:
			s.nick(cmd.cliente, cmd.args)
		case CMD_ENTRAR:
			s.entrar(cmd.cliente, cmd.args)
		case CMD_SALAS:
			s.listaSalas(cmd.cliente, cmd.args)
		case CMD_MSG:
			s.msg(cmd.cliente, cmd.args)
		case CMD_SAIR:
			s.sair(cmd.cliente, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn){
	log.Printf("Novo cliente esta conectado: %s", conn.RemoteAddr().String())

	
	c := &cliente{
		conn: conn,
		nome: "anonimo",
		comandos: s.comandos,
	  }
	  c.readInput()	
}

func (s *server) nick(c *cliente, args []string){
	c.nome = args[1]
	c.msg(fmt.Sprintf("Tudo certo, te chamaremos de %s", c.nome))
}
func (s *server) entrar(c *cliente, args []string){
	nomeSala := args[1]
	r, ok := s.salas[nomeSala]
	if !ok {
		r = &sala{
			nome: nomeSala,
			membros: make(map[net.Addr]*cliente),
		}
		s.salas[nomeSala] = r
	}
	r.membros[c.conn.RemoteAddr()] = c 

	if c.sala != nil {

	}

	s.sairSala(c)

	c.sala = r

	r.broadcast(c, fmt.Sprintf("%s entrou no sala", c.nome))
	c.msg(fmt.Sprintf("Bem vindo a %s", r.nome))
}
func (s *server) listaSalas(c *cliente, args []string){
	var salas []string	
	for nome := range s.salas {
		salas = append(salas, nome)
	}
	c.msg(fmt.Sprintf("As salas disponiveis sao: %s", strings.Join(salas, ", ")))
}
func (s *server) msg(c *cliente, args []string){
	if len(args) < 2 {
		c.msg("Pra mandar mensagem use o comando: /msg MSG")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.sala.broadcast(c, c.nome +  ": "+ msg)
}
func (s *server) sair(c *cliente, args []string){
	log.Printf("Cliente foi desconectado: %s", c.conn.RemoteAddr().String())

	s.sairSala(c)

	c.msg("Que pena voce sair :(")

	c.conn.Close()
}

func (s *server) sairSala(c *cliente){
	if c.sala != nil {
		delete(c.sala.membros, c.conn.RemoteAddr())
		c.sala.broadcast(c, fmt.Sprintf("%s saiu da sala", c.nome))
	}
}