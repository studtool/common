package queues

//go:generate easyjson

//easyjson:json
type CreatedUserData struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
}

//easyjson:json
type DeletedUserData struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
}

//easyjson:json
type VerifiedUsersData struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
}

//easyjson:json
type UpdatedEmailData struct {
	UserID   string `json:"userId"`
	OldEmail string `json:"oldEmail"`
	NewEmail string `json:"newEmail"`
}

//easyjson:json
type UpdatedPasswordData struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
}

//easyjson:json
type RegistrationEmailData struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

//easyjson:json
type VerificationEmailData struct {
	Email string `json:"email"`
}

//easyjson:json
type EmailUpdateData struct {
	Email string `json:""`
}
