package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CallbackBody struct {
	Text        string `json:"text"`
	UserName    string `json:"user_name"`
	LinkToTweet string `json:"link_to_tweet"`
	CreatedAt   string `json:"created_at"`
}

func main() {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body := &CallbackBody{}
		err := json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			fmt.Printf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Printf("%#v\n", body)
	})

	http.ListenAndServe(":8080", nil)
}
