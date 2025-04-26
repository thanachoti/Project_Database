package object

import "time"

type WatchHistory struct {
	HistoryID          int       `json:"history_id"`
	UserID             int       `json:"user_id"`
	ContentID          int       `json:"content_id"`
	WatchedTimestamp   time.Time `json:"watched_timestamp"`
	Progress           int       `json:"progress"`
	LanguagePreference string    `json:"language_preference"`
	CcPreference       string    `json:"cc_preference"`
}
