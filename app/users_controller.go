package app

import (
	"net/http"
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
