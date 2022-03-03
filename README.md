
<p align="center"><img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/RickAndMorty.png" width="360"></p>

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-93%25-brightgreen.svg?longCache=true&style=flat)</a>

# Go + Iris - Base API <a href="README_PTBR.md"><img width="20px" src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Flag_of_Brazil.svg/800px-Flag_of_Brazil.svg.png?20220111182208" /></a>

This is an example API, built with [Go](https://go.dev/) and [Iris](https://www.iris-go.com/). The codebase was created to show a working backend application in Go, with different options of database clients.

## Table of Contents

1. [How it works](#how-it-works)
2. [Getting started](#getting-started)
3. [Additional](#additional)
## How it works

The repository is organized in 3 different branches: 

- master
- bun
- pgx
- sqlx

### master
The master branch contains only the file structres and example controllers. No DB, just the Iris framework.

#### [Iris](https://github.com/kataras/iris)

> Iris is a fast, simple yet fully featured and very efficient web framework for Go.
> It provides a beautifully expressive and easy to use foundation for your next website or API.


### [bun](https://bun.uptrace.dev/)
> Simple and performant DB client for PostgreSQL, MySQL, and SQLite.

### [pgx](https://github.com/jackc/pgx)
> pgx is a pure Go driver and toolkit for PostgreSQL.

### [sqlx](https://github.com/jmoiron/sqlx) 
> sqlx is a library which provides a set of extensions on go's standard database/sql library.

## Getting started

### Running for development

```sh
make run
```

### Building the app

Building just for current OS
```sh
make build
```

Buiding for multiple distributions
```sh
make build-all
```

### Testing the app

```sh
make test
```

## Additional

### Generating coverage in HTML
```sh
make cover-html
```

### Generating  badge
Everytime you call ```make test``` or ```make cover-html```, the badge in README.md gets updated, but you can do it manually:
```sh
make badge
```

### Cleaning the repo
```sh
make clean
```


---
Feel free to contribute.

Made with ❤ in Brazil

The logo is from @ashleymcnamara and can be found [here](https://github.com/ashleymcnamara/gophers).