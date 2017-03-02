package SlackAlpha

import (
	"fmt"
	"net/http"
 	"io/ioutil"
	"appengine"
 	"appengine/urlfetch"
	"net/url"
)


func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	//Read the Request Parameter "command". This should be /compute
	command := r.FormValue("command")

	//Read the Request Parameter "text". This will contain the Query
	query := r.FormValue("text")

  //properQuery := strings.Replace(query, " ", "+", -1)

	if command == "/compute" {

		var Url *url.URL
		    Url, err := url.Parse("http://api.wolframalpha.com/v2/result?")
		    if err != nil {
		        panic("boom")
		    }

		    parameters := url.Values{}
		    parameters.Add("appid", "APP-ID-HERE")
		    parameters.Add("i", query)
		    Url.RawQuery = parameters.Encode()

    ctx := appengine.NewContext(r)
    client := urlfetch.Client(ctx)
    resp, err := client.Get(Url.String())

    if err != nil {
      fmt.Fprint(w, "API Call Failed")
    }
    body, err := ioutil.ReadAll(resp.Body)

    s := (string(body))
		fmt.Fprint(w, s)

   }
 }
