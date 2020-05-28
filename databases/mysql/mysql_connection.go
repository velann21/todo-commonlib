package mysql

import (
	"database/sql"
	"fmt"
	"github.com/velann21/todo-commonlib/commons/helpers"
	"sync"
)

var connection *SQLConnection;
type SQLConnection struct{
	db *sql.DB
	mutex sync.Mutex
}
func (sqlConn *SQLConnection) OpenSqlConnection() (error) {
	defer sqlConn.mutex.Unlock()
	if sqlConn.db == nil {
		sqlConn.mutex.Lock()
		if sqlConn.db == nil {
			fmt.Println(helpers.ReadEnv(helpers.MYSQLCONNECTIONSTRING))
			db, err := sql.Open("mysql", helpers.ReadEnv(helpers.MYSQLCONNECTIONSTRING))
			if err != nil {
				fmt.Println(err)
				return helpers.SomethingWrong
			}
			db.SetMaxIdleConns(10)
			db.SetMaxOpenConns(10)
			db.SetConnMaxLifetime(60)
			sqlConn.db = db
			connection = sqlConn
			return nil
		}
	}
	return nil
}


func GetSqlConnection() *sql.DB{
	return connection.db
}

