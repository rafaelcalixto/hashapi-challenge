package hashapi_db_conn

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
    "log"
    "encoding/json"
    "io/ioutil"
)

const (
    credfile string = "db_cred.json"
)

type connect struct {
    conn  *sql.DB
}

type credentials struct {
    Host    string `json:"host"`
    Port    string `json:"port"`
    User    string `json:"user"`
    DBName  string `json:"dbname"`
    PW      string `json:"pw"`
    SSLMode string `json:"sslmode"`
}

var (
    db        connect
    cs        string
    hash      string
    token     string
    query     string
    feedback  sql.Result
    work      int64
    err       error
    inter_js  credentials
    db_cred   map[string]string
    comp_data map[string]string
)

func ReadCred() map[string]string {
    jsonfile, err := ioutil.ReadFile(credfile)
    if err != nil { fmt.Println(err) }

    if err := json.Unmarshal(jsonfile, &inter_js); err != nil {
        panic(err)
    }

    db_cred = make(map[string]string)
    db_cred["host"] = inter_js.Host
    db_cred["port"] = inter_js.Port
    db_cred["user"] = inter_js.User
    db_cred["dbname"] = inter_js.DBName
    db_cred["pw"] = inter_js.PW
    db_cred["sslmode"] = inter_js.SSLMode

    return db_cred
}

func OpenConn() (connect) {
    db_cred = ReadCred()
    cs = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
                     db_cred["host"], db_cred["port"], db_cred["user"],
                     db_cred["dbname"], db_cred["pw"], db_cred["sslmode"])

    db.conn, err = sql.Open("postgres", cs)
    if err != nil { fmt.Println("Conldn't connect to the database", err) }

    return db
}

func CloseConn(db connect) {
    db.conn.Close()
}

func InsertToken(hash string, token string, db connect) (int64) {
    query = fmt.Sprintf("insert into tokens (hash, token) values " +
                      "('%s', '%s') on conflict do nothing", hash, token)
    feedback, err = db.conn.Exec(query)
    if err != nil { fmt.Println(err) }
    aff_rows, err := feedback.RowsAffected()
    if err != nil { fmt.Println(err) }
    return aff_rows
}

func ReturnAll(db connect) ( map[string]string ) {
    query = fmt.Sprintf("select hash, token from tokens")
    rows, err := db.conn.Query(query)
    if err != nil { fmt.Println(err) }

    comp_data = make(map[string]string)

    switch {
    case err == sql.ErrNoRows:
        comp_data["msg"] = "No token registered yet"
    case err != nil:
        log.Fatalf("Error while querying: %v", err)
    default:
        for rows.Next() {
            if err := rows.Scan(&hash, &token); err != nil {
                log.Fatalf("Error while extracting values: %v", err)
            }
            comp_data[hash] = token
        }
    }
    return comp_data
}

func ReturnToken(token string, db connect) ( map[string]string ) {
    query = fmt.Sprintf("select hash, token from tokens where hash = '%s'",
                        token)
    err := db.conn.QueryRow(query).Scan(&hash, &token)
    comp_data = make(map[string]string)

    switch {
    case err == sql.ErrNoRows:
        comp_data["msg"] = "404"
    case err != nil:
        log.Fatalf("Error while querying: %v", err)
    default:
        comp_data[hash] = token
    }
    return comp_data
}
