package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	UUID_SERVICE = "https://api.mojang.com/users/profiles/minecraft/"
	FACE_SERVICE = "https://crafatar.com/avatars/"
	HEAD_SERVICE = "https://crafatar.com/renders/head/"
	BODY_SERVICE = "https://crafatar.com/renders/body/"
)

type UUIDResponse struct {
	Username string `json:"name"`
	Uuid     string `json:"id"`
}

func GetUUID(username string) string {
	// Get UUID from username
	res, err := http.Get(UUID_SERVICE + username)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	uuid := UUIDResponse{}
	json.NewDecoder(res.Body).Decode(&uuid)
	if uuid.Uuid == "" {
		log.Fatal("UUID for ", username, " not found")
	}

	return uuid.Uuid
}

func GetFace(username string) {
	//Get UUID
	uuid := GetUUID(username)

	// Get face image from uuid
	res, err := http.Get(FACE_SERVICE + uuid + "?overlay")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	ioutil.WriteFile("public/images/"+username+"_face.png", data, 0666)
}
