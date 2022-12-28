package postgres

import (
	"database/sql"
	"template-go/internal/config"
)

func NewConnectPostGres(config config.Config) (*sql.DB, error) {
	return sql.Open(config.DBDriver, config.DBSource)
}
