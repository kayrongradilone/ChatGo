# Iniciar Servidor

- Usar o comando go build na pasta root do projeto
- Usar o comando ./ChatGo(nome da pasta root)

# Conectar Cliente ao Servidor

- Ativar os comandos de telnet no painel de de controle
- Passo 1 - acessar painel de controle
- Passo 2 - clicar na opcao de Programas
- Passo 3 - clicar na opcao de ativar ou desativar recursos do Windows
- Passo 4 - Procurar opcao Telnet e marcar o checkbox
- Passo 5 - Abrir o terminal e digitar o comando "telnet localhost 8888(porta de escolha)"

# Comandos

- `/nick <nome> ` - cadastra um nome, se não for cadastrado irá continuar como anônimo.
- `/entrar <nome> ` - entra em uma sala, se a sala não existir criará uma nova sala, Um usuário só pode estar em uma sala por vez.
- `/salas` - mostra a lista de salas disponíveis.
- `/msg <msg> ` - envia a mensagem para todos na sala
- `/sair` - disconecta do chat server.