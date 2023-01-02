package sessions

type SessionRepository interface {
	CreateSession(session *Session) (*Session, error)
	//GetSession(id string) (*domain.Session, error)
}
