package timekit

import (
	"log"
	"time"
)

var loc *time.Location

func init() {
	var err error
	loc, err = time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Panicln("invalid location")
	}
}

func Location() *time.Location {
	return loc
}

func Now() time.Time {
	return time.Time(time.Now().In(loc))
}

func ToWIB(t time.Time) time.Time {
	return t.In(loc)
}
