package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MemberReport struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	MemberID  int64         `json:"member_id"`
	Telephone string        `json:"telephone"`
}

type InputBody struct {
	MemberID int64 `json:"member_id"`
}

type OutputBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Data    []MemberReport `json: "data"`
}

func MemberRelationHandler(w http.ResponseWriter, r *http.Request) {
	url := "127.0.0.1"
	session, err := mgo.Dial(url)
	fmt.Println(err)
	c := session.DB("member_report").C("members")
	result := MemberReport{}

	err = c.Find(bson.M{"member_id": 1}).One(&result)
	fmt.Println(err, result)

	w.Header().Set("Content-Type", "application/json")

	output := OutputBody{0, ""}

	switch r.Method {
	case "POST":
		w.WriteHeader(200)

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var t []InputBody
		err := decoder.Decode(&t)
		if err != nil {
			w.WriteHeader(400)
			output.Message = "Incorrect json format"
			json.NewEncoder(w).Encode(output)
			break
		}
		for _, one := range t {
			_ = one
			// output.Data = append(output.Data)
		}

		json.NewEncoder(w).Encode(output)

	case "HEAD":
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(output)
	default:
		w.WriteHeader(400)
		output.Message = "Unsupport HTTP Method"
		json.NewEncoder(w).Encode(output)
	}
}

func main() {

	fmt.Println()

	http.HandleFunc("/member_relation", MemberRelationHandler)
	http.ListenAndServe("0.0.0.0:5679", nil)
	os.Exit(0)
}
