package swagger

import (
	"embed"
	"net/http"
)

//go:embed index.html
var ui embed.FS

func Handler() http.Handler {
	return http.FileServer(http.FS(ui))
}
