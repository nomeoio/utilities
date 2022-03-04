package utilities

import "time"

func ConvertUnixTime(unixTs int64, timezone, timeFormat string) (tm string, err error) {
	// ConvertUnixTime("Asia/Shanghai","2006-01-02 15:04:05")
	var loc *time.Location
	if loc, err = time.LoadLocation(timezone); err != nil {
		return
	}
	tm = time.Unix(unixTs, 0).In(loc).Format(timeFormat)
	return
}

func TimeNowString(timezone, timeFormat string) (nowString string, err error) {
	// TimeNow("Asia/Shanghai", "2006-01-02 15:04:05")
	var loc *time.Location
	if loc, err = time.LoadLocation(timezone); err != nil {
		return
	}
	var now time.Time = time.Now().In(loc)
	nowString = now.Format(timeFormat)
	return
}

func TimeNowUnix(timezone, timeFormat string) (timeUnix int64, err error) {
	// TimeNow("Asia/Shanghai", "2006-01-02 15:04:05")
	var loc *time.Location
	if loc, err = time.LoadLocation(timezone); err != nil {
		return
	}
	var now time.Time = time.Now().In(loc)
	timeUnix = now.Unix()

	return
}
