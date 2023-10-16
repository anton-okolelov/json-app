package httpserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/anton-okolelov/json-app/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4"
)

func (s Server) createUser(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	var u domain.User
	err := json.Unmarshal(body, &u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.userService.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := struct {
		Id int `json:"id"`
	}{
		Id: id,
	}

	render.JSON(w, r, response)
}

func (s Server) getUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := s.userService.GetUser(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			render.HTML(w, r, "User not found")
		} else {
			fmt.Printf("%+v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	render.JSON(w, r, u)
}

func (s Server) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.userService.GetAllUsers()
	if err != nil {
		fmt.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, users)
}
