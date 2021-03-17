package main

import (
	"log"
	"fmt"
	"time"
	"strconv"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"database/sql"
	"github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

type Baseinfo struct {
	TgId int
	ServerID int 
	Userdate time.Time
	}
	
func init() {
	
		// loads values from .env into the system
		if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		}
	}

func main() {

	intCh := make(chan time.Time) 
	var envs map[string]string
    envs, err := godotenv.Read(".env")
	bot, err := tgbotapi.NewBotAPI(envs["TG_API_KEY"])
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

	var telegramIdS string = strconv.FormatInt(update.Message.Chat.ID,10)
	var telegramIdI int = int(update.Message.Chat.ID)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	/// Start.Goroutine
	
	var existbool bool = uniqueID(telegramIdI)
	go connectBd(telegramIdI,intCh,existbool) 
	fmt.Println("Go routine starts")
	clientDate := <-intCh
	var clienDates string = clientDate.String()
	var elements string = update.Message.Command()
	if (existbool == true) && (elements == "knd") {
		msg.Text = "Дата окончания абонплаты" + " " + clienDates
	}	else if (existbool == false) && (elements == "knd") {
		msg.Text = "На" + " " + clienDates+ " Вы не зарегистрированы в сервисе" + "\nрегистрация /reg"
	}	else if elements == "weather" {
		msg.Text = weather("706524")
	}	else if elements == "extend" {
		msg.Text = "Продлить абонплату" + envs["EXTEND_API"]
	} 	else if elements == "reg" {
		msg.Text = "Регистрация на сайте https://onlinecab.net/auth/" + " " + "Ваш TelegramID " + telegramIdS
	}	else {
		msg.Text = "Чтобы узнать дату окончания абонплаты используйте /knd, а для продления /extend" + "\n Прогноз погоды /weather"
		}
		
		
        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }

	}

	}

func connectBd (parI int, parII chan time.Time, parIII bool ){
	if parIII == true {
		var envs map[string]string
		envs, err := godotenv.Read(".env")
		db, err := sql.Open("mysql", envs["MYSQL_CONNECT"]) //db_tg?parseTime trick for time.Time
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
		
		for res.Next() {
		var baseinfo Baseinfo

		err := res.Scan(&baseinfo.TgId, &baseinfo.ServerID, &baseinfo.Userdate)

		parII <- baseinfo.Userdate
		if err != nil {
			log.Fatal(err)
			}
		}
	} else {
		//const layoutISO = "2006-01-02"
		t := time. Now()
		parII <- t

} }

