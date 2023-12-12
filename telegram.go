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
	var charList = []string{"{", "}", "<", ">", "/", ":", "!", "_", "-", "=", "+", ".", "|", "#"}
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
	log.Printf("text: %+v", text)
	b, _ := json.Marshal(msg)

	var err error
	var respBody []byte
	if respBody, err = HttpRequest(
		"POST", b,
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", apiToken),
		[][]string{{"Accept", "application/json"}, {"Content-Type", "application/json"}},
	); err != nil {
		log.Fatalln(err)
	}

	if err = json.Unmarshal(respBody, &resp); err != nil {
		log.Fatalln(err)
	}
	log.Printf("resp: %+v", resp)
	log.Println("respbody", string(respBody))
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
