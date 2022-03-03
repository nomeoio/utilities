// "nomeoio-log" slack hook url: https://api.slack.com/apps/A02ER3QHX43/incoming-webhooks

package utilities

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"
	"time"
)

type Utils struct{}

func (u Utils) HttpRequest(requestMethod string, reqBody []byte, url string, headers [][]string) (respBody []byte, err error) {
	var req *http.Request
	var resp *http.Response
	req, err = http.NewRequest(requestMethod, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}
	if len(headers) > 0 {
		for _, header := range headers {
			req.Header.Add(header[0], header[1])
		}
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	var buf *bytes.Buffer = new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return
	}

	respBody = buf.Bytes()
	return
}

//
// Files IO
//

func (u Utils) ReadFile(fp string) (b []byte) {
	var err error
	if b, err = ioutil.ReadFile(fp); err != nil {
		log.Printf("%s not found: %s", fp, err)
	}
	return
}

func (u Utils) WriteFile(b []byte, filename string) (err error) {
	return ioutil.WriteFile(filename, b, 0644)
}

func (u Utils) DoesFileExist(fp string) (exist bool) {
	var err error
	if _, err = os.Stat(fp); err == nil {
		exist = true
	} else if errors.Is(err, os.ErrNotExist) {
		exist = false
	}
	return
}

func (u Utils) CreateDirIfNotExist(fp string) (err error) {
	var dp string = filepath.Dir(fp) // dir path
	if _, err = os.Stat(dp); os.IsNotExist(err) {
		err = os.MkdirAll(dp, os.ModePerm)
		if err == nil {
			return
		}
	}
	return
}

//
// Json Utils
//

func (u Utils) PrettyJsonString(body []byte) (respJson string) {
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, body, "", "  "); err != nil {
		log.Panic(err)
	}
	respJson = dst.String()
	return
}

//
// random Utils
//

func (u Utils) RandomString() string {
	var length int = 6
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

//
// Enctyption
//

func (u Utils) Encrypt(message string) (encodedMsg string, err error) {
	var plainText = []byte(message)

	var block cipher.Block
	var key = []byte(os.Getenv("ENCRYPTION_KEY"))
	if block, err = aes.NewCipher(key); err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	var cipherText = make([]byte, aes.BlockSize+len(plainText))
	var iv []byte = cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(crand.Reader, iv); err != nil {
		log.Panicln(err)
		return
	}

	var stream cipher.Stream = cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encodedMsg = base64.RawStdEncoding.EncodeToString(cipherText)
	return
}

func (u Utils) Decrypt(encodedMsg string) (decodedMsg string, err error) {
	var cipherText []byte
	if cipherText, err = base64.RawStdEncoding.DecodeString(encodedMsg); err != nil {
		return
	}

	var block cipher.Block
	var key = []byte(os.Getenv("ENCRYPTION_KEY"))
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	var iv []byte = cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	var stream cipher.Stream = cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	decodedMsg = string(cipherText)
	return
}

//
// Hash
//

func (u Utils) GenerateHash() string {
	var s string = fmt.Sprint(time.Now().UnixNano())
	var h [32]byte = sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", h)
}

func (u Utils) EmailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// func (u Utils) DownloadFile(url, fn string, ignoreErr bool) {

// 	// Create blank file
// 	var file *os.File
// 	var err error
// 	if file, err = os.Create(fn); err != nil {
// 		u.DealWithError(err)
// 	}
// 	client := http.Client{
// 		CheckRedirect: func(r *http.Request, via []*http.Request) error {
// 			r.URL.Opaque = r.URL.Path
// 			return nil
// 		},
// 	}

// 	// Put content on file
// 	var resp *http.Response
// 	if resp, err = client.Get(url); err != nil {
// 		u.DealWithError(err)
// 	}
// 	defer resp.Body.Close()

// 	if _, err = io.Copy(file, resp.Body); err != nil {
// 		u.DealWithError(err)
// 	}
// 	defer file.Close()
// }

// func (u Utils) DealWithError(err error) {
// 	var now, errMsg, wd string
// 	var e error
// 	if wd, e = os.Getwd(); e != nil {
// 		log.Panic(e)
// 	}
// 	now, _ = utils.TimeNow()
// 	errMsg = fmt.Sprintf("%s\nerr: %s\ndetail: %s", now, err.Error(), strings.Replace(string(debug.Stack()), wd, ".", -1))
// 	if IsTestMode && IsLocal {
// 		log.Fatalln(errMsg)
// 	} else {
// 		sc.SendPlainText(errMsg, os.Getenv("SlackWebHookUrlTest"))
// 	}
// }

// func (u Utils) CheckUrl(url string) (finalUrl string, contentLength int64, err error) {
// 	// check redirected final url & remove file size
// 	var resp *http.Response
// 	if resp, err = http.Head(url); err != nil {
// 		return
// 	}

// 	finalUrl = resp.Request.URL.String()
// 	contentLength = resp.ContentLength
// 	return
// }
