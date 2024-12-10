package controllers

import (
	"api/src/models"
	"api/src/repository"
	"api/src/views"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {

		views.Error(w, http.StatusUnprocessableEntity, err)
		return

	}

	repo, err := repository.NewUserRepository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err := repo.Create(user); err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}

	views.ToJSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	repo, err := repository.NewUserRepository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	users, err := repo.Get()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	views.ToJSON(w, http.StatusOK, users)

}

func FindUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	repo, err := repository.NewUserRepository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	user, err := repo.Find(id)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	views.ToJSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		views.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	repo, err := repository.NewUserRepository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	user.ID = id

	if err := repo.Update(user); err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	views.ToJSON(w, http.StatusOK, user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	repo, err := repository.NewUserRepository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
	}
	if err := repo.Delete(id); err != nil {
		views.Error(w, http.StatusInternalServerError, err)
	}
	views.ToJSON(w, http.StatusNoContent, nil)
}
