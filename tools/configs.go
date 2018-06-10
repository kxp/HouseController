package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

const AlarmSettingsFilename string = "alarmsettings.json"
const SettingsFilename string = "settings.json"

// Singleton stuff

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

/* Server settings */

type Config struct {
	// Database struct {
	// 	Host     string `json:"host"`
	// 	Password string `json:"password"`
	// } `json:"database"`
	HttpPort   string `json:"httpPort"`
	SerialPort string `json:"serialPort"`
}

func (singleton) LoadConfiguration(file string) (*Config, error) {

	var config *Config
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		log.Println(err.Error())
		config = generateEmptySettings()
		return config, nil
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(config)
	if config == nil {
		log.Println("Fail to de-serialize the options, using default settings")
		config = generateEmptySettings()
	}
	return config, nil
}

func generateEmptySettings() *Config {
	var settings *Config = new(Config)
	settings.HttpPort = ":8000"
	settings.SerialPort = "/dev/ttyUSB0"
	return settings
}

/*Alarm settings */

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

func (singleton) LoadAlarmSettings(file string) (*AlarmConfig, error) {

	var config *AlarmConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		config = generateEmptyAlarm()
		return config, nil
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(config)
	if config == nil {
		log.Println("Fail to de-serialize the alarm settings, using default settings")
		config = generateEmptyAlarm()
	}
	return config, nil
}

func generateEmptyAlarm() *AlarmConfig {

	var alarm = new(AlarmConfig)
	alarm.SelectedDays = make([]WeekDay, 7)
	alarm.LightColor = 5
	alarm.AlarmTime = "07:51:00"

	for day := 0; day < 7; day++ {
		alarm.SelectedDays[day].Active = false
		alarm.SelectedDays[day].Day = day
	}
	return alarm
}
