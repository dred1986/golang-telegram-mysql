package main

import (
	"log"
	"strconv"
	"database/sql"
    "github.com/joho/godotenv"
)
func uniqueID (parI int) (bool) {

	var envs map[string]string
	envs, err := godotenv.Read(".env")
	db, err := sql.Open("mysql", envs["MYSQL_CONNECT"]) 
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var telegramS string = strconv.Itoa(int(parI))
	res, err := db.Query("SELECT * FROM ListID WHERE TgId="+telegramS)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}
	var returnbool bool
	for res.Next() {
	var baseinfo Baseinfo
	
	err := res.Scan(&baseinfo.TgId, &baseinfo.ServerID, &baseinfo.Userdate)
		if err != nil {
		log.Fatal(err)
		}
		
		if parI == baseinfo.TgId {
			returnbool = true
			}	
	}
	return returnbool
}