package serverhandler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type WeekDay struct {
	Day    int `json:"Weekday"`
	Active bool
}

// AlarmConfig Received JSON structure From the app.
type AlarmConfig struct {
	AlarmTime    string
	SelectedDays []WeekDay
	LightColor   int32
}

var results []string

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
	fmt.Println("we arrive on hello")
}
func Alarm(w http.ResponseWriter, r *http.Request) {

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
		var alarm = generateEmptyAlarm()
		b, err := json.Marshal(alarm)
		if err != nil {
			fmt.Println(err)
			return
		}
		io.WriteString(w, string(b))
		//io.WriteString(w, "Hello world! from alarm")
		fmt.Println("Alarm: ", string(b))
	}

}

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

func alarmProcess() (string, error) {

	return "", nil
}

func generateEmptyAlarm() AlarmConfig {

	var alarm = new(AlarmConfig)
	alarm.SelectedDays = make([]WeekDay, 7)
	alarm.LightColor = 5
	alarm.AlarmTime = "07:50:00"

	for day := 0; day < 7; day++ {
		alarm.SelectedDays[day].Active = false
		alarm.SelectedDays[day].Day = day
	}
	return alarm
}
