package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// http handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("remote address=[%s]\n", r.RemoteAddr)
	fmt.Printf("msg=[%s]\n", r.FormValue("msg"))

	// get message parameter
	msg := r.FormValue("msg")

	if msg != "" {
		// get line token
		line_token := r.FormValue("token") // first try from parameter
		if line_token == "" {
			line_token = os.Getenv("LINE_TOKEN") // second try from environment parameter
			if line_token == "" {
				fmt.Printf("can't get token\n")
			}
		}

		URL := "https://notify-api.line.me/api/notify"
		u, err := url.ParseRequestURI(URL)
		if err != nil {
			fmt.Printf("url parameter error\n")
		}

		client := &http.Client{}

		form := url.Values{}
		form.Add("message", msg)

		body := strings.NewReader(form.Encode())

		req, err := http.NewRequest("POST", u.String(), body)
		if err != nil {
			fmt.Printf("request error\n")
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", "Bearer "+line_token)

		_, err = client.Do(req)
		if err != nil {
			fmt.Printf("request error\n")
		}
	}
}

func main() {
	fmt.Printf("start server\n")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
