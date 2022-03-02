# telegram-bot

## 介紹
此功能用於串接telegram

## 事前工作
1. 新增聯絡人 BotFather，https://t.me/botfather
2. 輸入/newbot
3. 輸入 Bot 名稱
4. 輸入 Bot 使用者名稱，必須為 bot 結尾
5. 成功建立後取得 Telegram API token
6. 傳送訊息給 Bot /start
7. 若是需要群組訊息則先將bot加入群組
8. 透過 https://api.telegram.org/bot"YOUR_BOT_TOKEN"/getUpdates 取得 Chat ID，請將紅字部分替
   換成機器人的 token，ex:
   https://api.telegram.org/bot12345678:AAAABBBBBCCCCC/getUpdates
   Option 1. 私人訊息(type=private)，取 id 值
   Option 2. 群組訊息(type=group)，取 id 值
   
## 使用功能
    
將token 跟 chatID 輸入後即可使用

	var token = "test:abcd"
	var chatID int64 = -584533863
建立出需要送出的messge RespTelegramMessage