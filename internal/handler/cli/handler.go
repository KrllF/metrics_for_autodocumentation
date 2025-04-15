package cli

import "github.com/KrllF/metrics_for_autodocumentation/internal/entity"

type (
	Service interface {
		GetMetrics(sourceFile, mdFile string) (entity.Stat, error)
	}
	Handler struct {
		serv Service
	}
)

func NewHandler(serv Service) *Handler {
	return &Handler{serv: serv}
}
