package main

type comandosID int

const (
		CMD_NOME comandosID = iota
		CMD_ENTRAR
		CMD_SALAS
		CMD_MSG
		CMD_SAIR
)

type comando struct {
	id comandosID
	cliente *cliente
	args []string
}