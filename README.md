# Galaveg [Статус: Архив]

## Проект заброшен по причине наличия уже устоявшихся фреймворков

Galaveg - это заготовка для веб-проектов на Golang.

Для запуска локальной среды разработки используйте команду:
```shell
docker compose -f docker-compose.local.yaml up -d
```
, и в дальнейшем используйте docker compose конфигурацию docker-compose.local.yaml для локальной разработки.

Для запуска миграций используйте команду:
```shell
go run . migrate
```

Для запуска одиночного теста используйте команду:
```shell
go test <test_file_path> -v
```
, например: `go test ./tests/unit/utils/path/path_test.go -v`.

В разработке использовалась утилита [watchexec](https://github.com/watchexec/watchexec), команда:
```shell
watchexec -r -e go,gohtml,html go run . serve
```

Production компиляция:
```shell
go build -ldflags="-s -w -X main.version=1.0.0" -o ./deploy/prod/chat ./cmd/chat/main.go 
go build -ldflags="-s -w -X main.version=1.0.0" -o ./deploy/prod/http ./cmd/http/main.go 
go build -ldflags="-s -w -X main.version=1.0.0" -o ./deploy/prod/serve ./cmd/serve/main.go 
go build -ldflags="-s -w -X main.version=1.0.0" -o ./deploy/prod/cmd ./main.go
```