package postgre

import (
	"database/sql"
	"template-go/internal/config"
)

func NewConnect(config config.Config) (*sql.DB, error) {
	return sql.Open(config.DBDriver, config.DBSource)
}
