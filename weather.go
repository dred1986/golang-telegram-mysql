package main

import (
	"log"
	"strconv"
    "github.com/joho/godotenv"
	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)



func weather(weatherS string) (string) {
	var envs map[string]string
    envs, err := godotenv.Read(".env")
    w, err := owm.NewCurrent("C", "ru", envs["OWM_API_KEY"])
	if err != nil {
		log.Fatalln(err)
	}
	weatherID,err := strconv.Atoi(weatherS)
	w.CurrentByID(weatherID)
	return ("В Керчи " + w.Weather[0].Description + "\nТемпература: " +strconv.Itoa(int(w.Main.Temp)) +"\nОщущается как: " +  strconv.Itoa(int(w.Main.FeelsLike))) + "\nВетер: " + strconv.Itoa(int(w.Wind.Speed)) + " м/с"
}
