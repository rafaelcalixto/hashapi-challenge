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

func TestAPILoad(t *testing.T) {
    // Update Test
    t.Run("Update Test", func(t *testing.T) {
        db_test = OpenConn()
        defer CloseConn(db_test)
        t_return = APILoad("test", "99999", "www.test.com", db_test)
        if t_return != 0 { t.Errorf("Error in the update answer") }
    })
}

func TestAPIQuery(t *testing.T) {
    // Select Test
    t.Run("Select Test", func(t *testing.T) {
        db_test = OpenConn()
        defer CloseConn(db_test)
        comp_data_test = APIQuery("test", "99999", db_test)
        if comp_data_test["msg"] != "No match results" {
            t.Errorf("Wrog return from database query") 
        }
    })
}
