package queues

//go:generate easyjson

//easyjson:json
type CreatedAvatarData struct {
	UserId string `json:"userId"`
}

//easyjson:json
type DeletedAvatarData struct {
	UserId string `json:"userId"`
}
