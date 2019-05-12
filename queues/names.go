package queues

const (
	CreatedUsersQueueName = "created_users"
	DeletedUsersQueueName = "deleted_users"

	VerifiedUsersQueueName = "verified_users"

	UpdatedEmailsQueueName   = "updated_emails"
	UpdatedPasswordQueueName = "updated_passwords"

	RegistrationEmailsQueueName   = "registration_emails"
	VerificationEmailsQueueName   = "verification_emails"
	EmailUpdateEmailsQueueName    = "email_update_emails"
	PasswordUpdateEmailsQueueName = "password_update_emails"

	ProfilesToCreateQueueName = "profiles_to_create"
	ProfilesToDeleteQueueName = "profiles_to_delete"

	AvatarsToCreateQueueName = "avatars_to_create"
	AvatarsToDeleteQueueName = "avatars_to_delete"
)
