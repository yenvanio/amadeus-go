package models

import (
	"fmt"
	"strings"
	"time"
)

type AmadeusTime struct {
	time.Time
}

const atLayout = "2006-01-02T15:04:00"

func (at *AmadeusTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		at.Time = time.Time{}
		return
	}
	at.Time, err = time.Parse(atLayout, s)
	return
}

func (at *AmadeusTime) MarshalJSON() ([]byte, error) {
	if at.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", at.Time.Format(atLayout))), nil
}
