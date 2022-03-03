package utilities

import "time"

func ConvertUnixTime(unixTs int64) (tm string) {
	tm = time.Unix(unixTs, 0).Format("2006-01-02")
	return
}

func TimeNow(timezone, timeFormat string) (string, int64) {
	// TimeNow("Asia/Shanghai", "2006-01-02 15:04:05")
	var loc *time.Location
	var err error
	if loc, err = time.LoadLocation(timezone); err != nil {
		panic(err)
	}
	var now time.Time = time.Now().In(loc)
	return now.Format(timeFormat), now.Unix()
}
