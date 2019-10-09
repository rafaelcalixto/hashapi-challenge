package hash_api

import (
    // Core libraries
    "fmt"
    "net/http"
    //"strings"
    //"strconv"
    "crypto/sha256"
    // Proprietary libraries
    dt "hashapi_db_conn"
)

var (
    msg          string
    token        string
    query_return map[string]string
    ok           bool
)

// This function returns for the Browsers some informations about the API
// This is mandary for some Browsers allows the access to the API
func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// This function returns a "Welcome message" to the API
func Index_handler(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    fmt.Fprintf(w, "This is the API for the Hash API test")
}

// This function take the parameters passed on the API and update the website on
// the database
func CreateHash(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)

    // Opening the Database Connection
    db := dt.OpenConn()
    defer dt.CloseConn(db)

    // Here the function to create the hashs is initialized
    hasher := sha256.New()

    // Here the token is captured from the URL parameters
    param, ok := r.URL.Query()["t"]
    if !ok || len(param[0]) < 1 {
        fmt.Println("URL token parameter incorrect")
        return
    }
    token = param[0]
    hasher.Write([]byte(token))
    msg = fmt.Sprintf("%x", hasher.Sum(nil) )

    // This block returns a message if the database was successfull updated
    work := dt.InsertToken(msg, token, db)
    if work == 1 {
        msg = fmt.Sprintf( "The hash \"%s\" was included", msg)
    } else {
        msg = fmt.Sprintf( "The hash \"%s\" already was been included", msg)
    }
    fmt.Fprintf(w, msg)
}

// This function take the parameters passed on the API and search for matchs on
// the database
func ReturnHash(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)

    // Opening the Database Connection
    db := dt.OpenConn()
    defer dt.CloseConn(db)

    // The below function returns the datafrom the database (if matched) and
    // the below if steatment builds a string to be printed on the API
    query_return = dt.ReturnAll(db)
    if _, ok := query_return["msg"]; ok {
        msg = fmt.Sprintf("{\n\t\"msg\": \"" + query_return["msg"] + "\"\n}")
    } else {
        msg = "{\n"
        for key, val := range query_return {
            msg = msg + fmt.Sprintf( "\t\"hash\": \"" + key + "\",\n" +
                                     "\t\"token\": \"" + val + "\",\n" )
        }
        msg += "\n}"
    }
    fmt.Fprintf(w, msg)
}

func ReturnText(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)

    // Opening the Database Connection
    db := dt.OpenConn()
    defer dt.CloseConn(db)

    // Here the token is captured from the URL parameters
    param, ok := r.URL.Query()["t"]
    if !ok || len(param[0]) < 1 {
        fmt.Println("URL token parameter incorrect")
        return
    }
    token = param[0]

    // The below function returns the datafrom the database (if matched) and
    // the below if steatment builds a string to be printed on the API
    query_return = dt.ReturnToken(token, db)
    if _, ok := query_return["msg"]; ok {
        msg = fmt.Sprintf("{\n\t\"msg\": \"" + query_return["msg"] + "\"\n}")
    } else {
        msg = "{\n"
        for key, val := range query_return {
            msg = msg + fmt.Sprintf( "\t\"hash\": \"" + key + "\",\n" +
                                     "\t\"token\": \"" + val + "\"" )
        }
        msg += "\n}"
    }

    // Opening the Database Connection
    fmt.Fprintf(w, msg)
}
