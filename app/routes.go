package app

import "net/http"

// InitRouter will intialise the router
func (app *App) InitRouter() {
	initialiseV1API(app)
}

func initialiseV1API(app *App) {
	//REST API
	app.Router.Get("/api/ping", app.handle(app.ping))
	app.Router.Get("/api/todo/", app.handle(app.GetAllTodos))
	//VIEW
	app.Router.Get("/", app.renderView(app.RenderIndex))
	app.Router.Get("/todo/", app.renderView(app.DisplayTodos))

	//STATIC
	fs := http.FileServer(http.Dir("assets/"))
	app.Router.Post("/static/", http.StripPrefix("/static/", fs))
}
