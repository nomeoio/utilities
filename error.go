package utilities

// func DealWithError(err error) {
// 	var now, errMsg, wd string
// 	var e error
// 	if wd, e = os.Getwd(); e != nil {
// 		log.Panic(e)
// 	}
// 	now, _ = TimeNow()
// 	errMsg = fmt.Sprintf("%s\nerr: %s\ndetail: %s", now, err.Error(), strings.Replace(string(debug.Stack()), wd, ".", -1))
// 	if IsTestMode && IsLocal {
// 		log.Fatalln(errMsg)
// 	} else {
// 		sc.SendPlainText(errMsg, os.Getenv("SlackWebHookUrlTest"))
// 	}
// }
