package authentication

type AuthenticationService interface {
	Authenticate(
		username,
		password string,
	) (*Authentication, error)
}
