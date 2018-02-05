package app

import (
	"database/sql"
	"io"
	"net/http"
	"time"

	"github.com/Arijeet-webonise/todo/app/models"
	"github.com/lib/pq"
)

func (a App) ForgetPassword(w http.ResponseWriter, r *http.Request) {

	var to []string
	to = append(to, "arijeet.baruah@weboniselab.com")

	err := a.Mailer.TestSMTP(to)

	if err != nil {
		a.Log.Error(err)
	} else {
		a.Log.Info("Mail send")
	}
}

func (a App) UserRegister(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/user/register.html"}
	res, err := a.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		a.Log.Error(err)
	}
	io.WriteString(w, res)

}

func (a App) UserRegisterSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var userService models.UserServiceImpl
		userService.DB = a.DB

		var user models.User

		user.Username = sql.NullString{String: r.FormValue("username"), Valid: true}
		user.Password = sql.NullString{String: r.FormValue("password"), Valid: true}
		user.Email = sql.NullString{String: r.FormValue("email"), Valid: true}
		user.Firstname = sql.NullString{String: r.FormValue("firstname"), Valid: true}
		user.Lastname = sql.NullString{String: r.FormValue("lastname"), Valid: true}
		user.LastLogin = pq.NullTime{Time: time.Now().Local(), Valid: true}
		err := userService.Register(&user)
		if err != nil {
			a.Log.Error(err)
		}
	}
}
