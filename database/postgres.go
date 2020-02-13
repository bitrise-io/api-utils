package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	// Postgres driver
	_ "github.com/lib/pq"
)

const (
	dbDialect string = "postgres"
)

var (
	sqlDB *gorm.DB
)

// PostgresDatabase ...
type PostgresDatabase struct {
	Host     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

// GetDB ...
func (psql PostgresDatabase) GetDB() *gorm.DB {
	return sqlDB
}

// Close ...
func (psql PostgresDatabase) Close() {
	closeDB(sqlDB)
	setDB(nil)
}

// InitializeConnection ...
func (psql PostgresDatabase) InitializeConnection(withDB bool) error {
	connString, err := psql.connectionString(withDB)
	if err != nil {
		return errors.WithStack(err)
	}
	db, err := gorm.Open(dbDialect, connString)
	if err != nil {
		return errors.Wrap(err, "Failed to open database")
	}
	setDB(db)
	if err = sqlDB.DB().Ping(); err != nil {
		closeDB(sqlDB)
		return errors.Wrap(err, "Failed to ping database")
	}
	isLogModeEnabled, err := strconv.ParseBool(os.Getenv("GORM_LOG_MODE_ENABLED"))
	if err == nil && isLogModeEnabled {
		sqlDB.LogMode(true)
	}
	return nil
}

func setDB(sdb *gorm.DB) {
	sqlDB = sdb
}

func closeDB(dbToClose *gorm.DB) {
	if dbToClose != nil {
		if err := dbToClose.Close(); err != nil {
			log.Printf(" [!] Exception: Failed to close DB: %+v", err)
		}
	}
}

func (psql PostgresDatabase) validate() error {
	if psql.Host == "" {
		return errors.New("No database host specified")
	}
	if psql.DBName == "" {
		return errors.New("No database name specified")
	}
	if psql.User == "" {
		return errors.New("No database user specified")
	}
	if psql.Password == "" {
		return errors.New("No database password specified")
	}
	return nil
}

func (psql PostgresDatabase) connectionString(withDB bool) (string, error) {
	if psql.Host == "" {
		psql.Host = os.Getenv("DB_HOST")
	}
	if psql.DBName == "" {
		psql.DBName = os.Getenv("DB_NAME")
	}
	if psql.User == "" {
		psql.User = os.Getenv("DB_USER")
	}
	if psql.Password == "" {
		psql.Password = os.Getenv("DB_PWD")
	}
	if psql.SSLMode == "" {
		psql.SSLMode = os.Getenv("DB_SSL_MODE")
	}
	if err := psql.validate(); err != nil {
		return "", err
	}
	connString := fmt.Sprintf("host=%s user=%s password=%s",
		psql.Host, psql.User, psql.Password)
	if withDB {
		connString += " dbname=" + psql.DBName
	}
	// optionals
	if psql.SSLMode != "" {
		connString += " sslmode=" + psql.SSLMode
	}
	return connString, nil
}
