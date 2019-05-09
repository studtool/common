package queues

//go:generate easyjson

//easyjson:json
type AvatarToCreateData struct {
	UserId string `json:"userId"`
}

//easyjson:json
type AvatarToDeleteData struct {
	UserId string `json:"userId"`
}
