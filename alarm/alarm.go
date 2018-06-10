package alarm

import (
	"HouseController/tools"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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

func alarmProcess() (string, error) {

	return "", nil
}

func generateEmptyAlarm() *tools.AlarmConfig {

	var alarm = new(tools.AlarmConfig)
	alarm.SelectedDays = make([]tools.WeekDay, 7)
	alarm.LightColor = 5
	alarm.AlarmTime = "07:50:00"

	for day := 0; day < 7; day++ {
		alarm.SelectedDays[day].Active = false
		alarm.SelectedDays[day].Day = day
	}
	return alarm
}
