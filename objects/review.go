package object

import "time"

type Review struct {
	ReviewID   int       `json:"review_id"`
	UserID     int       `json:"user_id"`
	ContentID  int       `json:"content_id"`
	Rating     int       `json:"rating"`
	ReviewText string    `json:"review_text"`
	ReviewDate time.Time `json:"review_date"`
}
