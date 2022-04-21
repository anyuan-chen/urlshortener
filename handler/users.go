package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anyuan-chen/urlshortener/server/auth"
	"github.com/anyuan-chen/urlshortener/server/users"
)

 func GetUser(w http.ResponseWriter, r *http.Request )  {
	client := r.Context().Value("client")
	provider := r.Context().Value("provider").(string)
	var resp []byte
	var err error
	if provider == "github"{
		resp, err = auth.GetGithubUserInfo(client.(*http.Client))
	} else if provider == "google" {
		resp, err = auth.GetGoogleUserInfo(client.(*http.Client))
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var data map[string]interface{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(data["id"].(string))
	id := users.GetUser(data["id"].(string))
	if (id == ""){
		users.CreateUser(data["id"].(string), "")
	}
	id = users.GetUser(data["id"].(string))
	type IdResponse struct {
		Id string `json:"id"`
	}
	fmt.Println(id)
	json.NewEncoder(w).Encode(IdResponse{id})
 }
 func GetLinksForUser(w http.ResponseWriter, r *http.Request) {
	client := r.Context().Value("client")
	provider := r.Context().Value("provider").(string)
	var resp []byte
	var err error
	if provider == "github"{
		resp, err = auth.GetGithubUserInfo(client.(*http.Client))
	} else if provider == "google" {
		resp, err = auth.GetGoogleUserInfo(client.(*http.Client))
	} else{
		http.Error(w, "wrong oauth",http.StatusInternalServerError )
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var data map[string]interface{}
	err = json.Unmarshal(resp, &data)
	for key, value := range data {
		fmt.Println(key, "corresponds", value)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	id := users.GetUser(data["id"].(string))
	if id == "" {
		http.Error(w, "user doesn't exist", http.StatusInternalServerError)
	}
	allLinks := users.GetLinksByUser(id)
	fmt.Fprint(w, allLinks)
 }

 func AddLinkForUser(w http.ResponseWriter, r *http.Request) {

 }


