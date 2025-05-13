package app

import (
	"github.com/KrllF/metrics_for_autodocumentation/internal/entity"
	"github.com/KrllF/metrics_for_autodocumentation/internal/handler/cli"
	"github.com/KrllF/metrics_for_autodocumentation/internal/service/checkStruct"
	goServ "github.com/KrllF/metrics_for_autodocumentation/internal/service/golang"
)

type (
	Handler interface {
		Run(sourceFile, mdFile string) (entity.StructStat, entity.Stat, error)
	}
	App struct {
		hand Handler
	}
)

func NewApp() *App {
	servGo := goServ.NewService()
	servStruct := checkStruct.NewService()
	hand := cli.NewHandler(servGo, servStruct)

	return &App{hand: hand}
}
