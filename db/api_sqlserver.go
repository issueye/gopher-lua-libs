package db

import (
	"database/sql"
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
)

type luaSqlServer struct {
	config *dbConfig
	sync.Mutex
	db *sql.DB
}

var (
	sharedSqlServer     = make(map[string]*luaSqlServer, 0)
	sharedSqlServerLock = &sync.Mutex{}
)

func init() {
	RegisterDriver(`sqlserver`, &luaSqlServer{})
}

func (sqlserver *luaSqlServer) constructor(config *dbConfig) (luaDB, error) {
	sharedSqlServerLock.Lock()
	defer sharedSqlServerLock.Unlock()

	if config.sharedMode {
		result, ok := sharedSqlServer[config.connString]
		if ok {
			return result, nil
		}
	}

	db, err := sql.Open(`sqlserver`, config.connString)
	if err != nil {
		return nil, err
	}
	result := &luaSqlServer{config: config}
	db.SetMaxIdleConns(config.maxOpenConns)
	db.SetMaxOpenConns(config.maxOpenConns)
	result.db = db

	if config.sharedMode {
		sharedSqlServer[config.connString] = result
	}

	return result, nil
}

func (sqlserver *luaSqlServer) getDB() *sql.DB {
	sqlserver.Lock()
	defer sqlserver.Unlock()
	return sqlserver.db
}

func (sqlserver *luaSqlServer) getTXOptions() *sql.TxOptions {
	return &sql.TxOptions{ReadOnly: sqlserver.config.readOnly}
}

func (sqlserver *luaSqlServer) closeDB() error {
	err := sqlserver.db.Close()
	if err != nil {
		return err
	}
	if sqlserver.config.sharedMode {
		sharedMySQLLock.Lock()
		delete(sharedMySQL, sqlserver.config.connString)
		sharedMySQLLock.Unlock()
	}
	return nil
}
