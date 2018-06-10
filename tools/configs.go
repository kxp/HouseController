package tools

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
	HttpPort   string `json:"httpPort"`
	SerialPort string `json:"serialPort"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		config.HttpPort = "8080"
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}

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

func LoadAlarmSettings(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		config.HttpPort = "8080"
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}
