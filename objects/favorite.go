package object

type Favorite struct {
	FavoriteID int `json:"favorite_id"`
	UserID     int `json:"user_id"`
	ContentID  int `json:"content_id"`
}
