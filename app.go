package main

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
	"vk-bot/handlers"
	"vk-bot/keyboards"
	"vk-bot/middlewares"
)

func main() {
	key, exist := os.LookupEnv("BOT_KEY")
	log.Print(key)
	if exist == false {
		log.Panic("Key doesn't exist")
	}
	pref := tele.Settings{
		Token:  key,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	keyboardYesNo, yes, no := keyboards.NewKeyboardYesNo()
	keyboardVariants, rightButton := keyboards.NewKeyboardVariants()
	keyboardStart := keyboards.NewKeyboardStart()

	hello := handlers.NewHandlerHello(keyboardYesNo)
	test := handlers.NewHandlerTest(keyboardVariants, keyboardStart)
	catch := handlers.NewHandlerCatchAnswer(keyboardStart, "3")

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Panic(err)
	}

	bot.Use(middlewares.Logger())
	bot.Use(middlewares.Private())

	bot.Handle("/start", hello.Hello)
	bot.Handle(no, test.Test)
	bot.Handle(yes, test.Test)
	bot.Handle(rightButton, catch.CatchRight)
	bot.Handle(tele.OnText, catch.CatchWrong)
	bot.Start()
}
