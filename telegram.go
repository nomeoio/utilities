package utilities

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Telegram struct{}

type TGSendMsg struct {
	ChatId          string `json:"chat_id,omitempty"`
	Text            string `json:"text,omitempty"`
	ParseMode       string `json:"parse_mode,omitempty"`
	MessageThreadID string `json:"message_thread_id,omitempty"`
}

type TGEvent struct {
	UpdateID int       `json:"update_id,omitempty"`
	Message  TGMessage `json:"message,omitempty"`
}
type TGMessage struct {
	MessageID      int           `json:"message_id,omitempty"`
	From           TGMessageFrom `json:"from,omitempty"`
	Chat           TGMessageChat `json:"chat,omitempty"`
	ReplyToMessage *TGMessage    `json:"reply_to_message,omitempty"`
}

type TGMessageFrom struct { // the user info
	ID        int    `json:"id,omitempty"`
	IsBot     bool   `json:"is_bot,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	Username  string `json:"username,omitempty"`
}

type TGMessageChat struct {
	Date int    `json:"date,omitempty"`
	Text string `json:"text,omitempty"`
}

type TGResp struct {
	OK          bool   `json:"ok,omitempty"`
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
}

func (tg Telegram) EscapeChars(text string) string {
	var charList = []string{"(", ")", "!", "_", "-", "=", "+", ".", "{", "}", "<", ">", "|", "#"}
	for _, char := range charList {
		if strings.Contains(text, char) {
			text = strings.ReplaceAll(text, char, `\`+char)
		}
	}
	return text
}

func (tg Telegram) SendMessage(apiToken, text, chatId, threadID string) (resp TGResp) {
	text = tg.EscapeChars(text)
	var msg = TGSendMsg{
		ChatId:          chatId,
		Text:            text,
		ParseMode:       "MarkdownV2",
		MessageThreadID: threadID,
	}
	b, _ := json.Marshal(msg)

	var err error
	var respBody []byte
	if respBody, err = HttpRequest(
		"GET", b,
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", apiToken),
		[][]string{{"Accept", "application/json"}, {"Content-Type", "application/json"}},
	); err != nil {
		log.Fatalln(err)
	}

	if err = json.Unmarshal(respBody, &resp); err != nil {
		log.Fatalln(err)
	}
	// if !resp.OK {
	// 	SC.SendPlainText(string(respBody), os.Getenv("SlackWebHookNomeoHQErrs"))
	// }
	return
}

// func (amw ApiMiddlewares) TelegramAction(c *fiber.Ctx) (err error) {
// 	var s string = string(c.Body())
// 	SC.SendPlainText(s, os.Getenv("SlackWebHookTest"))

// 	var event TGEvent
// 	if err := json.Unmarshal(c.Body(), &event); err != nil {
// 		log.Fatalln(err)
// 	}

// 	// SC.SendPlainText(s, os.Getenv("SlackWebHookTest"))
// 	s = fmt.Sprintf("Telegram (%s): %+v\n", c.Method(), event)
// 	SC.SendPlainText(s, os.Getenv("SlackWebHookTest"))
// 	return c.SendString("ok")
// }

// var regex Regex

type Regex struct{}

const RegexUserName string = `^[a-z0-9_]{6,20}$`
const RegexUserNameSpecial string = `^[a-z0-9_]{2,20}$`
const RegexPassphrase string = `^.{8,128}$`
const RegexEmail string = `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`

func (reg Regex) VerifyUserName(username string) bool {
	r := regexp.MustCompile(RegexUserName)
	return r.MatchString(username)
}

func (reg Regex) VerifyUserNameSpecial(username string) bool {
	r := regexp.MustCompile(RegexUserNameSpecial)
	return r.MatchString(username)
}

func (reg Regex) VerifyPassphrase(passphrase string) bool {
	r := regexp.MustCompile(RegexPassphrase)
	return r.MatchString(passphrase)
}

func (reg Regex) VerifyEmail(email string) bool {
	r := regexp.MustCompile(RegexEmail)
	return r.MatchString(email)
}
