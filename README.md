# transaction-routine

Esse projeto permite o gerenciamento de contas e transações de usuários através de requisições HTTP.

# Executando o projeto

Caso queira conectar em um banco de dados diferente, você precisará editar o arquivo de configuração do `docker-compose.yaml` ou passar as variáveis de ambiente diretamente pro comando docker que preferir

```bash
docker-compose up -d
```

Você estará inicializando dois containers:

- Um para o banco de dados PostgreSQL que realiza a persistência de contas e transações
- O contêiner da aplicação rodando por default na porta `3000`
