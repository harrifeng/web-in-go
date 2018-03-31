package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/harrifeng/web-in-go/db/mongodb"

	"gopkg.in/mgo.v2/bson"
)

type Member struct {
	ID        bson.ObjectId `json:"-" bson:"_id,omitempty"`
	MemberID  int64         `json:"member_id" bson:"member_id"`
	Telephone string        `json:"telephone" bson:"telephone"`
}

type InputBody struct {
	MemberID int64 `json:"member_id"`
}

type OutputBody struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []Member `json: "data"`
}

func MemberRelationHandler(w http.ResponseWriter, r *http.Request) {
	session := mongodb.CloneSession()
	defer session.Close() // Return to the pool
	c := session.DB("member_report").C("members")

	w.Header().Set("Content-Type", "application/json")

	output := OutputBody{0, "", []Member{}}

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
			result := Member{}
			err := c.Find(bson.M{"member_id": one.MemberID}).One(&result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("result", result)
			}
			output.Data = append(output.Data, result)
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
