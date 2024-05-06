package models

import "time"

type AuthHeaders struct {
	SessionKey string
	UserAgent  string
}

type Session struct {
	SessionKey string
	TTL        TTL
}

type CacheUserSession struct {
	SessionKey string
	UserAgent  string
	Duration   time.Duration
	ID         int
}
