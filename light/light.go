package light

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var results []string

func Light(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		//does the JSON parsing internally
		results = append(results, string(body))

		fmt.Fprint(w, "POST done")
	} else {
		//http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		io.WriteString(w, "Hello world! from Light")
		fmt.Println("Light: ", r.Body)
	}
}
