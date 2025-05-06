package handlers

import (
	"log/slog"
	"net/http"

	"github.com/ed-henrique/lowdesk/internal/models"
	"github.com/ed-henrique/lowdesk/internal/ui/pages"
)

func TicketsGetAll(handlerName string, q *models.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tickets, err := q.GetAllTickets(r.Context())
		if err != nil {
			errMessage := "could not get all tickets"
			slog.Error(errMessage, slog.String("handler", handlerName), slog.String("err", err.Error()))
			http.Error(w, errMessage, http.StatusInternalServerError)
		}

		err = pages.TicketsGetAll(tickets).Render(r.Context(), w)
		if err != nil {
			errMessage := "could not render get all tickets"
			slog.Error(errMessage, slog.String("handler", handlerName), slog.String("err", err.Error()))
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})
}

func TicketsCreate(handlerName string, q *models.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := pages.TicketsCreate().Render(r.Context(), w)
		if err != nil {
			errMessage := "could not render create tickets"
			slog.Error(errMessage, slog.String("handler", handlerName), slog.String("err", err.Error()))
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})
}

func APITicketsCreate(handlerName string, q *models.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			errMessage := "could not parse form"
			slog.Error(errMessage, slog.String("handler", handlerName), slog.String("err", err.Error()))
			http.Error(w, errMessage, http.StatusInternalServerError)
		}

		title := r.PostForm.Get("title-form")
		content := r.PostForm.Get("content-form")

		if _, err := q.InsertTicket(r.Context(), models.InsertTicketParams{
			Title:   title,
			Content: content,
		}); err != nil {
			errMessage := "could not insert ticket"
			slog.Error(errMessage, slog.String("handler", handlerName), slog.String("err", err.Error()))
			http.Error(w, errMessage, http.StatusInternalServerError)
		}

		http.Redirect(w, r, "/tickets/", http.StatusSeeOther)
	})
}
