package app

import (
	"github.com/KrllF/metrics_for_autodocumentation/internal/entity"
	"github.com/KrllF/metrics_for_autodocumentation/internal/handler/cli"
	goServ "github.com/KrllF/metrics_for_autodocumentation/internal/service/golang"
)

type (
	Handler interface {
		Run(sourceFile, mdFile string) (entity.Stat, error)
	}
	App struct {
		hand Handler
	}
)

func NewApp() *App {
	serv := goServ.NewService()
	hand := cli.NewHandler(serv)

	return &App{hand: hand}
}
