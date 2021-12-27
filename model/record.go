package model

import "time"

type Record struct {
	Id   int
	Pid  int
	Txt  string
	Time time.Time
}
