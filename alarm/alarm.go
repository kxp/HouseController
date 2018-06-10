package alarm

import (
	"HouseController/tools"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var results []string

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
		instanceTools := tools.GetInstance()
		alarmSetting, err := instanceTools.LoadAlarmSettings(tools.AlarmSettingsFilename)
		if err != nil {
			log.Println(err)
			return
		}

		b, err := json.Marshal(alarmSetting)
		if err != nil {
			log.Println(err)
			return
		}
		io.WriteString(w, string(b))
		log.Println("Alarm: ", string(b))
	}

}

func alarmProcess() (string, error) {

	return "", nil
}
