package cli

import "github.com/KrllF/metrics_for_autodocumentation/internal/entity"

type (
	Service interface {
		GetMetrics(sourceFile, mdFile string) (entity.Stat, error)
	}
	ServiceStruct interface {
		EqualStruct(sourceFile, mdFile string) (entity.StructStat, error)
	}
	Handler struct {
		servGo     Service
		servStruct ServiceStruct
	}
)

func NewHandler(serv Service, servStruct ServiceStruct) *Handler {
	return &Handler{servGo: serv, servStruct: servStruct}
}
