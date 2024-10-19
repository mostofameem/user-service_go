package db

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"
	"user-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var CntOnce = sync.Once{}
var db *sqlx.DB

func NewDB(dbCnf *config.DBConfig) *sqlx.DB {
	CntOnce.Do(func() {

	})
	return db
}
func connect(dbConfig *config.DBConfig) *sqlx.DB {
	dbSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
	if !dbConfig.EnableSSLMode {
		dbSource += " sslmode=disable"
	}

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dbCon.SetConnMaxIdleTime(
		time.Duration(dbConfig.MaxIdleTimeInMinute * int(time.Minute)),
	)

	return dbCon
}
