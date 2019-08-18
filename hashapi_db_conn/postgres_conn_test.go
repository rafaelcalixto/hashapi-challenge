package neoway_db_conn

import (
    "testing"
)

var (
    db_test        connect
    test_test      string
    err_test       error
    t_return       int64
    comp_data_test map[string]string
)

func TestOpenConn(t *testing.T) {
    // Connection Test
    t.Run("Connection Test", func(t *testing.T) {
        db_test = OpenConn()
        defer CloseConn(db_test)
        err_test = db.conn.QueryRow("select 1").Scan(&t_return)
        if err_test != nil { t.Errorf("Connection error") }
        if t_return != 1 { t.Errorf("Error in the database answer") }
    })
}
