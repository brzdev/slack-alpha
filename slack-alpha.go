package SlackAlpha

import (
	"fmt"
	"net/http"
  //"bytes"
  "io/ioutil"
  //"strings"
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
		    parameters.Add("appid", "A72QEU-JR4WL54R8Y")
		    parameters.Add("i", query)
		    Url.RawQuery = parameters.Encode()

/*
    apiUrl := "http://api.wolframalpha.com/v2/result"
    var buffer bytes.Buffer
    buffer.WriteString(apiUrl)
    buffer.WriteString("?appid=A72QEU-JR4WL54R8Y")
    buffer.WriteString("&i=")
    buffer.WriteString(properQuery)
    getURL := buffer.String() */

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
