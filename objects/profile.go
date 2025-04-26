package object

type Profile struct {
	ProfileID      int    `json:"profile_id"`
	UserID         int    `json:"user_id"`
	ProfileName    string `json:"profileName"`
	AgeRestriction int    `json:"age_restriction"`
}
