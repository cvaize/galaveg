# Galaveg [Статус: В разработке]

Galaveg - это заготовка для веб-проектов на Golang. Я вдохновлялся Laravel, его простотой и количеством уже написанных функций.
Название Galaveg появилось в следствии замены букв L на G в слове Laravel, так и получилось Galaveg.

В конце Readme.md вы можете увидеть скриншоты работающего приложения.

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

Чтобы вывести debug информацию в HTML:
```go
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goforj/godump"
)

func Index(c *gin.Context) {
    c.Data(200, "text/html; charset=utf-8", []byte(godump.DumpHTML(err)))
    return
}
```