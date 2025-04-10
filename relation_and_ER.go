package main

import "time"

type USER struct {
	user_id           int `gorm:"primaryKey"`
	names             string
	email             string
	password          string
	username          string
	SubscriptionPlan  string `gorm:"type:ENUM('Free', 'Basic', 'Premium')"`
	registration_date time.Time
	profile_picture   string
	age               int
	Reviews           []REVIEW  `gorm:"foreignKey:user_id"`
	profileUser       []PROFILE `gorm:"foreignKey:user_id"`
}
type PROFILE struct {
	profile_id     int `gorm:"primaryKey"`
	user_id        int
	profile_name   string
	age_restrition int
	WatchHistory   []WATCH_HISTORY `gorm:"foreignKey:profile_id"`
	FavoriteList   []FAVORITE_LIST `gorm:"foreignKey:profile_id"`
}
type MOVIE_SHOW struct {
	content_id        int `gorm:"primaryKey"`
	title             string
	description       string
	release_year      int
	duration          int
	content_type      string `gorm:"type:ENUM('Movie', 'Show')"`
	total_seasons     int
	thumbnail_url     string
	video_url         string
	rating            float64
	ContentCategories []CONTENT_CATEGORY `gorm:"foreignKey:content_id"`
	WatchHistory      []WATCH_HISTORY    `gorm:"foreignKey:content_id"`
	Reviews           []REVIEW           `gorm:"foreignKey:content_id"`
	MovieLanguages    []MOVIE_LANGUAGE   `gorm:"foreignKey:content_id"`
	MovieSubtitles    []MOVIE_SUBTITLE   `gorm:"foreignKey:content_id"`
}
type CATEGORY struct {
	catrgory_id       int `gorm:"primaryKey"`
	category_name     string
	ContentCategories []CONTENT_CATEGORY `gorm:"foreignKey:category_id"`
}
type CONTENT_CATEGORY struct {
	content_category_id int `gorm:"primaryKey"`
	content_id          int
	category_id         int
	MovieShow           MOVIE_SHOW `gorm:"foreignKey:content_id"`
	Category            CATEGORY   `gorm:"foreignKey:category_id"`
}
type WATCH_HISTORY struct {
	history_id          int `gorm:"primaryKey"`
	user_id             int
	content_id          int
	watched_timestamp   time.Time
	progress            int
	language_preference string
	cc_preference       string
	User                USER       `gorm:"foreignKey:user_id"`
	MovieShow           MOVIE_SHOW `gorm:"foreignKey:content_id"`
}
type FAVORITE_LIST struct {
	favorite_id int `gorm:"primaryKey"`
	user_id     int
	content_id  int
	User        USER       `gorm:"foreignKey:user_id"`
	MovieShow   MOVIE_SHOW `gorm:"foreignKey:content_id"`
}
type REVIEW struct {
	review_id   int `gorm:"primaryKey"`
	user_id     int
	content_id  int
	rating      float64
	review_text string
	review_date time.Time
	User        USER       `gorm:"foreignKey:user_id"`
	MovieShow   MOVIE_SHOW `gorm:"foreignKey:content_id"`
}
type LANGUAGE struct {
	language_id    int `gorm:"primaryKey"`
	language_name  string
	MovieLanguages []MOVIE_LANGUAGE `gorm:"foreignKey:language_id"`
	MovieSubtitles []MOVIE_SUBTITLE `gorm:"foreignKey:language_id"`
}
type MOVIE_LANGUAGE struct {
	id          int `gorm:"primaryKey"`
	content_id  int
	language_id int
	MovieShow   MOVIE_SHOW `gorm:"foreignKey:content_id"`
	Language    LANGUAGE   `gorm:"foreignKey:canguage_id"`
}
type MOVIE_SUBTITLE struct {
	id          int `gorm:"primaryKey"`
	content_id  int
	language_id int
	MovieShow   MOVIE_SHOW `gorm:"foreignKey:content_id"`
	Language    LANGUAGE   `gorm:"foreignKey:canguage_id"`
}

func main() {
	//fmt.Println("Hello, Go!")
}
