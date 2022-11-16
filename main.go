package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type time_entry struct {
	start, stops time.Time
	desc, p_name string
}

type toggle_data struct {
	url, apiToken, apiPass, user_agent, workspace_id string
}

func main() {
	data := toggle_data{
		url:          "https://api.track.toggl.com/reports/api/v2/details",
		apiToken:     "f14f54ed7cb124dfd12e38f53e62d049",
		apiPass:      "api_token",
		user_agent:   "rickert.luis@googlemail.com",
		workspace_id: "3921935",
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, data.url+"?"+"user_agent="+data.user_agent+"&"+"workspace_id="+data.workspace_id, nil)
	req.SetBasicAuth(data.apiToken, data.apiPass)
	//q := req.URL.Query()
	//q.Add("user_agent", data.user_agent)
	//q.Add("workspace_id", data.workspace_id)

	//req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.RawQuery)
	res, _ := client.Do(req)

	defer res.Body.Close()
	fmt.Println(httputil.DumpResponse(res, true))

	res.Body.Close()

}

// Todo: read data from toggle
// Todo: read wbs mapping from disk
// Todo: create
