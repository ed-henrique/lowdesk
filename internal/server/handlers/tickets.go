package handlers

import (
	"log/slog"
	"net/http"

	"github.com/ed-henrique/lowdesk/internal/models"
	"github.com/ed-henrique/lowdesk/internal/ui/pages"
)

func TicketsCreate(handlerName string, q *models.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := pages.TicketsCreate().Render(r.Context(), w)
		if err != nil {
			errMessage := "could not render '" + handlerName + "'"
			slog.Error(errMessage, slog.String("handler", handlerName), slog.String("err", err.Error()))
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})
}
