package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userRepository := repositories.NewUserRepository(db)
	userRepository.Create(user)

	jsonUser, _ := json.Marshal(user)
	w.Write(jsonUser)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding user"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}
