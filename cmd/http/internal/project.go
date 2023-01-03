package internal

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"template-go/internal/config"
	log "template-go/internal/logger"
	"template-go/internal/postgre"
	ginHttp "template-go/internal/server"
	db "template-go/internal/sqlc/repositories"
	"template-go/pkg/makertoken/paseto"
)

type Flags struct {
	Debug             debugFlag
	TimeZone          string
	ConfigurationFile string
	LogFile           string
	Version           bool
}

// Run the project
func Run(f Flags) error {
	// set default timezone
	err := os.Setenv("TZ", f.TimeZone)
	if err != nil {
		return err
	}

	loadConfig, err := config.LoadConfig(f.ConfigurationFile)
	if err != nil {
		return err
	}

	// Config log System
	log.Initialize("project.log", "INFO")

	conn, err := postgre.NewConnect(loadConfig)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Cannot configure Connection postgre Error. %s", err))
		return err
	}

	store := db.NewStore(conn)

	runDBMigration(loadConfig.MigrationURL, loadConfig.DBSource)

	tokenMaker, err := paseto.New(loadConfig.TokenSymmectricKey)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Cannot configure Token Maker Error. %s", err))
		return err
	}

	server, err := ginHttp.NewServer(loadConfig, store, tokenMaker)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Cannot configure Http Server Error. %s", err))
		return err
	}

	err = server.Start()
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Cannot start Http Server Error. %s", err))
		return err
	}

	return nil
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Logger.Fatal(fmt.Sprintf("cannot create new migrate instance: %v", err))
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Logger.Fatal(fmt.Sprintf("failed to run migrate up: %v", err))
	}
	log.Logger.Info("db migrated successfully!")
}
