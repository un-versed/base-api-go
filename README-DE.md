
<p align="center"><img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/RickAndMorty.png" width="360"></p>

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-93%25-brightgreen.svg?longCache=true&style=flat)</a>

# Go + Iris - Base API

Dies ist eine Beispiel API, die mit [Go](https://go.dev/) und [Iris](https://www.iris-go.com/) erstellt wurde. Die Codebase wurde erstellt, um eine funktionierende Backend Anwendung in Go mit verschiedenen Optionen für database clients zu zeigen.

## Inhalt

 1. [Wie es läuft](#wie-es-läuft)
 2. [Anfängst](#anfängst)
 3. [Zusätzlich](#zusätzlich)

## Wie es läuft

Das Repository ist in 3 verschiedene Branches organisiert:

- master
- bun
- pgx
- sqlx


### master
Der Master branch enthält nur die Dateistrukturen und Beispiel Controller. Nur das Iris-Framework, keine KB.

#### [Iris]

 > Iris ist ein schnelles, einfaches, aber voll funktionsfähiges und sehr effizientes Web Framework für Go.
 > Es bietet eine schön ausdrucksstarke und einfach zu verwendende Grundlage für Ihre nächste Website oder API.

### [bun](https://bun.uptrace.dev/)
 > Einfacher und leistungsfähiger DB client für PostgreSQL, MySQL, und SQLite


### [pgx](https://github.com/jackc/pgx)
 > pgx ist ein reiner Go driver und Toolkit für PostgreSQL.

### [sqlx](https://github.com/jmoiron/sqlx)
 > sqlx ist eine Sammlung, die eine Reihe von Erweiterungen für die standard database/sql-Sammlung von go bereitstellt.

## Anfängst

### Laufende development

```sh
make run
```

### Erstellen der App

Nur für das aktuelle OS erstellen
```sh
make build
```

Für mehrere Distributionen erstellen
```sh
make build-all
```

### Testen der App

```sh
make test
```

## Zusätzlich

### Erzeugen coverage in HTML

```sh
make cover-html
```

### Erzeugen badge
Jedes Mal, wenn Sie ```make test``` oder ```make cover-html``` aufrufen, wird das Abzeichen in README.md aktualisiert, aber Sie können dies manuell tun:

```sh
make badge
```

### Reinemachen des Repo

```sh
make clean
```

---
Du kannst gerne beitragen.

Das Logo stammt von @ashleymcnamara und kann [hier](https://github.com/ashleymcnamara/gophers) gefunden werden.

In Brasilien hergestellt mit ❤️.