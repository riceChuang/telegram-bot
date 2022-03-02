package model

import (
	"fmt"
	"time"
)

const (
	Level_Info    = "Information"
	Level_Warning = "Warning"
	Level_System  = "System"
	Level_Error   = "Error"
)

type RespTelegramMessage struct {
	Level      string
	EventName  string
	Content    string
	serverID   int32
	vendorID   int32
	vendorName string
	serverName string
	time       time.Time
}

func (rtm *RespTelegramMessage) String() string {
	return fmt.Sprintf(
		//"服務ID: %v \n"+
		//	"代理ID: %v \n"+
		"代理名稱: %v \n"+
			"服務名稱: %v \n"+
			"告警時間: %v \n"+
			"告警等級: %v \n"+
			"告警項目: %v \n"+
			"訊息: %v \n",
		rtm.vendorName, rtm.serverName, rtm.time.String(), rtm.Level, rtm.EventName, rtm.Content)
}

func (rtm *RespTelegramMessage) SetServerID(id int32) {
	rtm.serverID = id
}
func (rtm *RespTelegramMessage) SetServerName(name string) {
	rtm.serverName = name
}
func (rtm *RespTelegramMessage) SetTime() {
	rtm.time = time.Now()
}
func (rtm *RespTelegramMessage) SetVendorID(id int32) {
	rtm.vendorID = id
}
func (rtm *RespTelegramMessage) SetVendorName(name string) {
	rtm.vendorName = name
}
