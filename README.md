# TESTE GOBRAX
### CRUD de veículos e motoristas

### Descrição
uma API simples em Golang e MySQL paracadastro de motoristas, veículos e a possibilidade de víncular um motorista a um veículo

### Ferramentas
```
Docker
MySql
Golang 1.21.3
Fiber Framework
PhpMyAdmin
```

## Executar localmente via docker
use:
```
# clone o projeto, (git clone https://github.com/neiltonrodriguez/teste-gobrax)
# acesse a pasta do projeto (cd teste-gobrax)
# docker-compose up -d --build
# renomei o arquivo .env-example para .env


o docker se encarregará de instalar todas as dependências do projeto, incluindo o Go e Mysql
```

#### se não tiver workbech instalado, pode usar o endereço do phpMyadmin para acessar o banco:
```
# PhpMyAdmin: http://localhost:8888/ 
# usuário: user
# senha: password
# execute o script sql que está dentro de ./docs/model.sql para criar as tabelas
```

#### se tiver workbech acesse o banco depois de executar o docker com esses dados:
```
Host: 127.0.0.1:3307
User: user
Password: password
Database: teste_gobrax

# execute o script sql que está dentro de ./docs/model.sql para criar as tabelas
```
### Regra de negócio
```
Não ficou explícito no teste qual era o relacionamento entre motorista e veíclo, considerei que seja 1:1
portanto, um motorista só pode ser vinculado a um veículo e vice-versa.

o cadastro do veículo pode ser feito informando já o id do motorista, que pode ser listado no get da rota 'v1/driver/'
todos os endpoints de getAll
```

##  Rotas e Modelo de requisição:
GetAll:
nessa rota, é possível obter todos os motoristas, e essa mesma rota pode ser usada para listar os motoristas que não estão vinculados a um veículo na hora de cadastrar o veículo, para isso, basta usar a flag 'available=true'.

esse endpoint tem paginação e filtro por nome.
```
curl --request GET \
  --url 'http://localhost:8080/v1/driver?available=false&limit=0&page=0&name='
```

Create:
essa rota é para cadastro de motoristas.

```
curl --request POST \
  --url http://localhost:8080/v1/driver/ \
  --data '{
  "name": "joão",
  "drivers_license": "D",
  "phone": "61987652399",
  "age": 20
}'
```


GetAll:
essa rota é get de todos os veículos, tbm tem paginação e tem filtro porplaca e marca.

```
curl --request GET \
  --url 'http://127.0.0.1:8080/v1/vehicle?plate=&brand=&limit=10&page=1'
```

Create:
essa rota é para cadastro de veículos, o 'driver_id' pode ser informado de acordo com o get na rota de motoristas usando a flag 'available=true' ou pode ser mandado como null e depois usado o update para fazer o vínculo.

```
curl --request POST \
  --url http://localhost:8080/v1/vehicle/ \
  --data '{
  "driver_id": 5,
  "plate": "ABC123",
  "brand": "FIAT",
  "model": "uno"
}'
```

além disso, a API tem os endpoints de GetById, Update e Delete, tanto para motorista como veículo

Developed by Neilton Rodrigues