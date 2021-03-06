package dao

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
  "fmt"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Alias    string `json:"alias"`
}

type Session struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func AuthenticateUser(cred UserCredentials) string {
	fmt.Println("INSIDE AUTHENTICAEUSER FUNCTION")
  session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("user")
  query := clctn.Find(bson.M{"username": cred.Username, "password": cred.Password})
	err := query.One(&response)
  cn := 0
  cn,err= clctn.Find(bson.M{"username":"Adetee"}).Count()
	uuidStr := uuid.Must(uuid.NewV4()).String()
  if (cn ==0){
    err = clctn.Insert(&UserCredentials{Username:"Adetee",Password:"Adetee",Alias:"Adi"})
  }
    fmt.Println("UUID",uuidStr)
  err = clctn.Insert(&Session{Username:"Adetee",Token:uuidStr})
  fmt.Println("Successfully inserted")
  sessionStruct := Session{cred.Username, uuidStr}
	if err != nil {
		return ""
	}

	sessionClctn := session.DB("simplesurveys").C("session")
	sessionClctn.Insert(sessionStruct)
	return uuidStr
}

func GetSessionDetails(token string) UserCredentials {
	session := MgoSession.Clone()
	defer session.Close()

	var response Session
	sessionClctn := session.DB("simplesurveys").C("session")
	query := sessionClctn.Find(bson.M{"token": token})
	err := query.One(&response)
	if err != nil {
		return UserCredentials{}
	}

	var cred UserCredentials
	clctn := session.DB("simplesurveys").C("user")
	query = clctn.Find(bson.M{"username": response.Username})
	err = query.One(&cred)
	return cred
}
