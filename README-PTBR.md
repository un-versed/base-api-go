
<p align="center"><img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/RickAndMorty.png" width="360"></p>

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-93%25-brightgreen.svg?longCache=true&style=flat)</a>

# Go + Iris - Base API
Essa é uma API de exemplo, feita em [Go](https://go.dev/) e [Iris](https://www.iris-go.com/). O repositório foi criado para exemplificar uma aplicação backend funcional em Go, com diferentes opções de clientes de banco de dados.

## Table of Contents

1. [Como funciona](#como-funciona)
2. [Começando a usar](#começando-a-usar)
3. [Adicional](#Adicional)
## Como funciona

O repositório é organizado em 3 branches diferentes: 

- master
- bun
- pgx
- sqlx

### master
A branch master contém apenas a estrutura de pastas e arquivo, e controladores de exemplo. Sem banco de dados, somente o Iris framework.
#### [Iris](https://github.com/kataras/iris)

> Iris is a fast, simple yet fully featured and very efficient web framework for Go.
> It provides a beautifully expressive and easy to use foundation for your next website or API.


### [bun](https://bun.uptrace.dev/)
> Simple and performant DB client for PostgreSQL, MySQL, and SQLite.

### [pgx](https://github.com/jackc/pgx)
> pgx is a pure Go driver and toolkit for PostgreSQL.

### [sqlx](https://github.com/jmoiron/sqlx) 
> sqlx is a library which provides a set of extensions on go's standard database/sql library.

## Começando a usar

### Rodando para desenvolvimento

```sh
make run
```

### Compilando a aplicação

Compilando apenas para o sistema operacional atual
```sh
make build
```

Compilando para multiplas distribuições
```sh
make build-all
```

### Testando a aplicação

```sh
make test
```

## Adicional

### Gerando um arquivo de cobertura de testes em HTML
```sh
make cover-html
```

### Gerando o banner de cobertura de testes
Sempre que rodar ```make test``` ou ```make cover-html```, o banner é atualizado no README.md, mas você pode rodar manualmente:
```sh
make badge
```

### Limpando o código
```sh
make clean
```


---
Sinta-se livre para contribuir.

Feito com ❤ no Brasil

O logo é propriedade de @ashleymcnamara e pode ser encontrado [aqui](https://github.com/ashleymcnamara/gophers).