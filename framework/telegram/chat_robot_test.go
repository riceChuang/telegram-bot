package telegram

import (
	"github.com/riceChuang/telegram-bot/model"
	"testing"
)

var token = "1852217926:AAGt0kLHus1scj5zIUVxdOrnd2VjCd15kNs"
var chatID int64 = -584533863

func TestTelegramRobot_SendMessage(t *testing.T) {

	telegramBot := NewTelegramRobot(token, chatID,12345,"sczjh_1_A_01",1,"abc")

	type fields struct {
		TelegramBot
	}
	type args struct {
		text *model.RespTelegramMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_telegram_bot",
			fields: fields{
				telegramBot,
			},
			args: args{
				text: &model.RespTelegramMessage{
					Level: model.Level_Error,
					EventName: "RTP",
					Content:  "testing",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := tt.fields.TelegramBot
			if err := bot.SendMessage(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
