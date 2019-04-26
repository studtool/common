package queues

//go:generate easyjson

//easyjson:json
type CreatedUserData struct {
	UserId string `json:"userId"`
}

type DeletedUserData struct {
	UserId string `json:"userId"`
}
