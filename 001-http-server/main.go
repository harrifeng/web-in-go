package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type MemberReport struct {
	MemberID string `json:"member_id"`
	TK       string `json:"tk"`
	Status   int    `json:"status"`
}

type InputBody struct {
	MemberID string `json:"member_id"`
}

type OutputBody struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    []MemberReport `json: "data"`
}

func MemberRelationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output := OutputBody{0, "", []MemberReport{}}

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
			output.Data = append(output.Data, MemberReport{one.MemberID, "", 0})
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
