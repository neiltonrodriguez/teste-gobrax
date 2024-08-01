# TESTE GOBRAX
### CRUD de produtos

### Descrição
uma simples api para CRUD de produtos, usando golang e mysql

### Contexto de negócio
```
Docker
MySql
Golang 1.21.3
Fiber Framework
```

## Executar localmente via docker
use:
```
# clone o projeto, (git clone https://github.com/neiltonrodriguez/teste-gobrax)
# acesse a pasta do projeto (cd teste-gobrax)
# docker-compose up -d --build
# renomei o arquivo .env-example para .env

```
se der conflito com a porta 3306, pare o seu serviço mysql, ou mude a porta no Dockerfile
```
# sudo service mysqld stop
ou
# sudo /etc/init.d/mysqld stop

o docker se encarregará de instalar todas as dependências do projeto, incluindo o Go e Mysql
```

#### se não tiver workbech instalado, pode usar o endereço do phpMyadmin para acessar o banco:
```
# PhpMyAdmin: http://localhost:8888/ 
# execute o script sql que está dentro de ./docs/model.sql para criar as tabelas
```

#### se tiver workbech acesse o banco depois de executar o docker com esses dados:
```
Host: 127.0.0.1:3306
User: user
Password: password
Database: gobrax

# execute o script sql que está dentro de ./script-sql/model.sql para criar as tabelas
```


##  Rotas e Modelo de requisição:
cURL create:
```
curl --request POST \
  --url http://localhost:8080/product/ \
  --data '{
  "name": "play station 4",
  "description": "teste de produto",
  "value": 5000.99
}'
```

cURL GetAll:
```
curl --request GET \
  --url http://localhost:8080/product/
```

cURL getById:
```
curl --request GET \
  --url http://localhost:8080/product/:id
```

cURL Update:
```
curl --request PUT \
  --url http://localhost:8080/product/:id \
  --data '{
  "name": "Notebook i7",
  "description": "teste de produto lenovo",
  "value": 1999.33
}'
```

cURL Delete:
```
curl --request DELETE \
  --url http://localhost:8080/product/:id
```

Developed by Neilton Rodrigues