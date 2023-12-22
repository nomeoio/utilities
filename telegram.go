package utilities

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Telegram struct{}

type TGSendMsg struct {
	ChatId          string `json:"chat_id"`
	Text            string `json:"text"`
	ParseMode       string `json:"parse_mode"`
	MessageThreadID string `json:"message_thread_id"`
}

type TGEvent struct {
	UpdateID int       `json:"update_id"`
	Message  TGMessage `json:"message"`
}
type TGMessage struct {
	MessageID      int           `json:"message_id"`
	From           TGMessageFrom `json:"from"`
	Chat           TGMessageChat `json:"chat"`
	ReplyToMessage *TGMessage    `json:"reply_to_message"`
}

type TGMessageFrom struct { // the user info
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type TGMessageChat struct {
	Date int    `json:"date"`
	Text string `json:"text"`
}

type TGResp struct {
	OK          bool   `json:"ok"`
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
}

func (tg Telegram) EscapeChars(text string) string {
	var charList = []string{"{", "}", "(", ")", "[", "]", "<", ">", "/", ":", "!", "-", "=", "+", ".", "|", "#"}
	for _, char := range charList {
		if strings.Contains(text, char) {
			text = strings.ReplaceAll(text, char, `\`+char)
		}
	}
	return text
}

func (tg Telegram) SendMessage(apiToken, text, chatId, threadID string) (resp TGResp, err error) {
	// text = tg.EscapeChars(text)
	var msg = TGSendMsg{
		ChatId:          chatId,
		Text:            text,
		ParseMode:       "Markdown",
		MessageThreadID: threadID,
	}
	b, _ := json.Marshal(msg)

	var respBody []byte
	if respBody, err = HttpRequest(
		"POST", b,
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", apiToken),
		[][]string{{"Accept", "application/json"}, {"Content-Type", "application/json"}},
	); err != nil {
		return
	}

	err = json.Unmarshal(respBody, &resp)
	return
}
