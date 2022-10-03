package main 

import (
	"log"
	"net"
)

func main() {
	//iniciando o server centralizado
	s := newServer()
	go s.run() 

	//iniciando o server tcp
	listener, err := net.Listen("tcp",":8888")
	if err != nil {
		log.Fatal("Impossivel inciar o servidor: %s", err.Error())
	} 

	defer listener.Close()
	log.Printf("Servidor iniciado em :8888")
	
	// aceitando todos os clientes novos
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Não foi possivel conectar: %s", err.Error())
			continue
		}
	
	//iniciando o cliente
		go s.newClient(conn)
	}
}