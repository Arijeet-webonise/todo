package app

import (
	"errors"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/Arijeet-webonise/todo/app/domain"
	"github.com/Arijeet-webonise/todo/pkg/framework"
)

func (a App) GetAllTodos(w *framework.Response, r *framework.Request) {

	todo, err := a.TodoSeviceImpl.GetAllTodos()
	if err != nil {
		a.Log.Info(err.Error())
		w.NotFound(errors.New("could not find the todo"))
		return
	}
	a.Log.Info(todo)
	if err != nil {
		w.BadRequest(err)
		return
	}

	w.PutInData("todo", todo)
}

//RenderIndex renders the index page
func (app *App) DisplayTodos(w http.ResponseWriter, r *http.Request) {
	roles, err := app.TodoSeviceImpl.GetAllTodos()
	tmplList := []string{"./web/views/base.html",
		"./web/views/todos/todo.html"}
	data := struct {
		Roles []*domain.Todo
	}{roles}
	res, err := app.TplParser.ParseTemplate(tmplList, data)
	if err != nil {
		app.Log.Info(err)
	}
	io.WriteString(w, res)
}

func (app *App) CreateTodo(w *framework.Response, r *framework.Request) {
	t := &domain.Todo{
		Title: randomString(),
	}
	createErr := app.TodoSeviceImpl.InsertTodo(t)
	if createErr != nil {
		app.Log.Info(createErr.Error())
	}
	w.PutInData("todo", t)
}

func randomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 5)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}
