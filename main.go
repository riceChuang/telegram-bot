package main

import (
	"fmt"
	"github.com/riceChuang/telegram-bot/framework/telegram"
	"github.com/riceChuang/telegram-bot/model"
)

func main() {

	var token = "test:abcd"
	var chatID int64 = -584533863
	telegramBot := telegram.NewTelegramRobot(token, chatID, 12345, "sczjh_1_A_01", 1, "abc")

	for i := 0; i < 10; i++ {
		sendText := &model.RespTelegramMessage{
			Level:     model.Level_Info,
			EventName: "Test",
			Content:   fmt.Sprintf("testing, value:%v", i),
		}
		err := telegramBot.SendMessage(sendText)
		if err != nil {
			fmt.Println(err)
		}
	}
}
