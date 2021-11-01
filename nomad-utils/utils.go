package apiserver

import (
	"strconv"
	"time"
)

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Time
func GetTimeOutSeconds() time.Duration {
	return time.Duration(AppConfig().AppTimeout) * time.Second
}

func GetTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
