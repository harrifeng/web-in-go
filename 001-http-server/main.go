package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type IDFAInfo struct {
	IDFA   string `json:"idfa"`
	TK     string `json:"tk"`
	Status int    `json:"status"`
}

type InputInfo struct {
	IDFA string `json:"idfa"`
}

type OutputInfo struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []IDFAInfo `json: "data"`
}

func MemberRelationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output := OutputInfo{0, "", []IDFAInfo{}}

	switch r.Method {
	case "POST":
		w.WriteHeader(200)

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var t []InputInfo
		err := decoder.Decode(&t)
		if err != nil {
			w.WriteHeader(400)
			output.Message = "Incorrect json format"
			json.NewEncoder(w).Encode(output)
			break
		}
		for _, one := range t {
			output.Data = append(output.Data, IDFAInfo{one.IDFA, "", 0})
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
