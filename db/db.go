package db

import (
	"context"
	"database/sql"
	"sync"

	model "github.com/bayarindevteam/bayaringo/graph/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	// Build with postgres database driver
	_ "github.com/lib/pq"
)

var (
	Store Storage
	once  sync.Once
)

// Storage interface defines required storage functions for synclib to operate
type Storage interface {
	Init(serviceName string, host string, port string, database string, username string, password string) error
	Ping() error
	Begin() (*sqlx.Tx, error)
	SetDB(db *sqlx.DB) // for unittest purposes

	CreateLink(ctx context.Context, l *model.Link) (string, error)
}

// DatabaseStorage implements Storage and enables storing synclib data in a postgres database
type DatabaseStorage struct {
	db *sqlx.DB              // database connection
	st map[string]*sqlx.Stmt // map of prepared sql statements
}

// newDatabase returns a DatabaseStorage object
func newDatabase() Storage {
	return &DatabaseStorage{}
}

// Handle returns a Storage Interface for interacting with the Storage layer
func Handle() Storage {
	once.Do(func() {
		Store = newDatabase()
	})

	return Store
}

// Init ...
func (d *DatabaseStorage) Init(serviceName string, host string, port string, database string, username string, password string) error {
	log.WithFields(log.Fields{
		"host":     host,
		"port":     port,
		"database": database,
	}).Info("Init database")

	var db *sqlx.DB
	var err error

	log.Debug("user=" + username + " password=" + password + " host=" + host + " port=" + port + " dbname=" + database + " application_name=" + serviceName + " sslmode=disable connect_timeout=2")
	db, err = d.connect("postgres", "user="+username+" password="+password+" host="+host+" port="+port+" dbname="+database+" application_name="+serviceName+" sslmode=disable connect_timeout=2")
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

func (d *DatabaseStorage) connect(driverName, dataSourceName string) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	db, err = sqlx.Connect(driverName, dataSourceName)

	return db, err
}

// Begin a database transaction
func (d *DatabaseStorage) Begin() (*sqlx.Tx, error) {
	return d.db.Beginx()
}

// Ping the database to check connectivity
func (d *DatabaseStorage) Ping() error {
	_, err := d.st["health"].Exec()
	if err != nil {
		return err
	}

	return nil
}

// SetDB set db for unit test only
func (d *DatabaseStorage) SetDB(db *sqlx.DB) {
	d.db = db
}

// ToNullString converts string types into sql.NullString types
func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}
