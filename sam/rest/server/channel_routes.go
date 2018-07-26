package server

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `channel.go`, `channel.util.go` or `channel_test.go` to
	implement your API calls, helper functions and tests. The file `channel.go`
	is only generated the first time, and will not be overwritten if it exists.
*/

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (ch *ChannelHandlers) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Route("/channels", func(r chi.Router) {
			r.Get("/", ch.List)
			r.Put("/", ch.Create)
			r.Post("/{channelID}", ch.Edit)
			r.Get("/{channelID}", ch.Read)
			r.Delete("/{channelID}", ch.Delete)
			r.Get("/{channelID}/members", ch.Members)
			r.Post("/{channelID}/members/{userID}", ch.Join)
			r.Delete("/{channelID}/members/{userID}", ch.Part)
			r.Post("/{channelID}/invite", ch.Invite)
		})
	})
}
