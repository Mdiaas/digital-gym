Para executar a aplicação é necessário possuir o docker instalado.
Necessário criar o arquivo .env no diretório raiz do projeto com as informações presentes dentro do arquivo .env.example

```
docker-compose up
```

rotas da aplicação local:

Cadastro de usuário
POST: /user
LOGIN
POST: /user/login

ROTAS DAS AULAS
CRIAÇÃO
POST: /gymclass
EDIÇÃO
PUT: /gymclass/{gymclass-id}
DELEÇÃO
PUT: /gymclass/{gymclass-id}
LISTAGEM
GET: /gymclass/page={number}?limit={total}&sort={asc/desc}
BUSCA DE AULA
GET: /gymclass/{gymclass-id}