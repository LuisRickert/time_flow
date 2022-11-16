package main


import https
 
var token = "963da1510feb0c0ca888fdacee1b0a48"

type Posts []struct {
	Userid int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func get_report() {
	resp, err := https.Get("https:toggle.com") + params.Encode
}
