package sqlc

type SqlcStore struct {
	//store db.Store
}

func NewSqlcStore() *SqlcStore {
	return &SqlcStore{}
}
