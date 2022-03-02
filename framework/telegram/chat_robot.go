package telegram

import (
	"context"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/riceChuang/telegram-bot/model"
	log "github.com/sirupsen/logrus"
	"sync"
)

var robotInstance *TelegramRobot
var telegramRobotOnce sync.Once

type TelegramBot interface {
	Name() string
	SetChatID(chatID int64) int64
	SendMessage(respMessage *model.RespTelegramMessage) error
}

type TelegramRobot struct {
	token      string
	chatID     int64
	serverID   int32
	vendorID   int32
	serverName string
	vendorName string
	Robot      *tgbotapi.BotAPI
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewTelegramRobot(token string, chatID int64, serverID int32, serverName string, vendorID int32, vendorName string) TelegramBot {
	if robotInstance != nil {
		return robotInstance
	}

	telegramRobotOnce.Do(func() {
		robotInstance = &TelegramRobot{}
		robotInstance.token = token
		robotInstance.chatID = chatID
		robotInstance.serverName = serverName
		robotInstance.vendorName = vendorName
		robotInstance.serverID = serverID
		robotInstance.vendorID = vendorID
		robotInstance.ConnectBotServer(token)
		robotInstance.ctx, robotInstance.cancel = context.WithCancel(context.Background())
	})

	return robotInstance
}

func GetTelegramInstance() TelegramBot {
	if robotInstance != nil && robotInstance.Robot != nil {
		return robotInstance
	}
	return nil
}

func (bot *TelegramRobot) ConnectBotServer(token string)  {
	if bot.Robot == nil {
		botConnect, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Error("初始化 tg_bot error:%v", err)
			return
		}
		log.Info("Authorized on account %s", botConnect.Self.UserName)
		bot.Robot = botConnect
	}
}

func (bot *TelegramRobot) Name() string {
	return "baifu_gameserver_bot"
}

func (bot *TelegramRobot) Memo() string {
	return ""
}

func (bot *TelegramRobot) SetChatID(chatID int64) int64 {
	bot.chatID = chatID
	return bot.chatID
}

func (bot *TelegramRobot) SendMessage(respMessage *model.RespTelegramMessage) error {
	bot.setRespMessage(respMessage)
	msg := tgbotapi.NewMessage(bot.chatID, respMessage.String())
	if _, err := bot.Robot.Send(msg); err != nil {
		log.Error("[telegramRobot] send message fail error:%v", err)
		return err
	}
	return nil
}

func (bot *TelegramRobot) setRespMessage(respMessage *model.RespTelegramMessage) {
	respMessage.SetServerName(bot.serverName)
	respMessage.SetTime()
	respMessage.SetServerID(bot.serverID)
	respMessage.SetVendorID(bot.vendorID)
	respMessage.SetVendorName(bot.vendorName)
}
