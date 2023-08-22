package main

import (
	"db-tool/pkg/app"
	"db-tool/pkg/common/server"
	"embed"
	"io/fs"
	"log"

	"github.com/pietjan/template"

	_ "embed"
)

//go:embed web
var web embed.FS

func main() {
	server := server.New(
		server.Static(mustSubFS(web, `web/static`)),
		server.Template(template.New(
			template.FS(mustSubFS(web, `web/template`)),
		)),
		server.Application(app.Application{
			Command: app.Commands{},
			Query:   app.Queries{},
		}),
	)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}

func mustSubFS(fsys fs.FS, folder string) fs.FS {
	sub, err := fs.Sub(fsys, folder)
	if err != nil {
		panic(err)
	}

	return sub
}
