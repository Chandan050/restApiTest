package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

var users []User

type User struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
	Age   int    `json: "age"`
}

func Additem(q http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newProfile User
	json.NewDecoder(r.Body).Decode(&newProfile)
	q.Header().Set("Content-Type", "application/json")
	users = append(users, newProfile) //adding the profile to our
	json.NewEncoder(q).Encode(users)
}

func getAllProfiles(q http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(users)

}

func getprofile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID cound not be convered to integer"))
		return
	}
	var profile User
	for a, i := range users {
		if id == i.ID {
			profile = users[a]
		}
	}
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)

}
func updateprofile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID cound not be convered to integer"))
		return
	}

	var updateprofile User
	json.NewDecoder(r.Body).Decode(&updateprofile)

	users[id] = updateprofile
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(updateprofile)

}
func deleteprofile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID cound not be convered to integer"))
		return
	}
	users = append(users[:id], users[:id+1]...)
	w.WriteHeader(200)

}

func main() {
	router := httprouter.New()
	// router.HandlerFunc("profiles", additem).Methods("POST")
	router.POST("/profiles", Additem) // create

	router.GET("/profiles", getAllProfiles) // read
	router.GET("/profile/(id)", getprofile)
	router.PUT("/profile/(id)", updateprofile) // update

	router.DELETE("/profile/(id)", deleteprofile) // delete

	http.ListenAndServe(":8081", router)
}
