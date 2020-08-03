package mysql

import (
	"fmt"
	"github.com/velann21/todo-commonlib/commons/helpers"
	"testing"
)
import "github.com/sirupsen/logrus"

func TestSQLConnection_OpenSqlConnection(t *testing.T) {
	helpers.SetEnv()
	sqlConn := SQLConnection{}
	err := sqlConn.OpenSqlConnection()
	if err != nil{
		logrus.Error("Error occured")
		return
	}
	con := GetSqlConnection()
	rows, err := con.Query("select * from users;")
	if err != nil{
		logrus.Error(err)
	}
	for rows.Next() {
		fmt.Println("sdksjdkjd")
	}

}
