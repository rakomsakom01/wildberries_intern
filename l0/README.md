#### Задание L0

### Структура проекта
```
.
├── application
│   ├── cmd
│   │   └── app
│   │       └── app.go
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   ├── database
│   │   │   ├── db.go
│   │   │   ├── freeids.go
│   │   │   ├── query.go
│   │   │   └── query_helpers.go
│   │   ├── nats
│   │   │   └── nats.go
│   │   └── web
│   │       ├── handlers.go
│   │       └── web.go
│   ├── main
│   ├── main.go
│   └── ui
│       ├── html
│       │   ├── footer.html
│       │   ├── header.html
│       │   ├── json.html
│       │   ├── main.html
│       │   └── query.html
│       └── static
│           └── css
│               └── style.css
├── docker-compose.yml
├── go.mod
├── go.sum
├── L0
│   ├── L0.pdf
│   └── model.json
├── README.md
├── send
└── send.go
```
### Файлы

## `/application/internal/database`

Содержит в себе всё необходимое для работы с базами данных

## `/application/internal/nats`

Содержит в себе всё необходимое для работы с nats streaming

## `/application/internal/web`

Содержит в себе веб сервер

## `/application/cmd/app`

Основное приложение, заставляет всё работать вместе

## `/application/ui`

Содержит в себе шаблоны для страниц, отвечает за визуал

## `/docker-compose.yml`

Запускает бд и nats streaming

## `/send`

Отправляет данные в nats streaming

## `/send.go`

Исходники send
