package app

import (
	"database/sql"

	"github.com/Arijeet-webonise/todo/app/domain"
	"github.com/Arijeet-webonise/todo/pkg/logger"
	"github.com/Arijeet-webonise/todo/pkg/mailer"
	"github.com/Arijeet-webonise/todo/pkg/templates"
	"github.com/go-zoo/bone"
)

// App enscapsulates the App environment
type App struct {
	Router         *bone.Mux
	Cfg            *Config
	Log            logger.ILogger
	TplParser      templates.ITemplateParser
	DB             *sql.DB
	TodoSeviceImpl domain.TodoService
	Mailer         *mailer.SMTPMailer
}
