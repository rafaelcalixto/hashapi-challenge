package hash_api

import (
    // Core libraries
    "testing"
    "net/http"
    "net/http/httptest"
    "io/ioutil"
)

var (
    test string
)

func TestIndexHandler(t *testing.T) {
    // Test for the Welcome Message of the API
    t.Run("Index_handler test", func(t *testing.T) {
        req, err := http.NewRequest("GET", "localhost:8000", nil)
        if err != nil { t.Fatalf("Request fail: %v", err) }

        rec := httptest.NewRecorder()
        Index_handler(rec, req)

        res := rec.Result()
        defer res.Body.Close()
        if res.StatusCode != http.StatusOK {
            t.Errorf("Expected Status OK, got %v", res.Status)
        }

        b_msg, err := ioutil.ReadAll(res.Body)
        if err != nil {
            t.Fatalf("Could not read response: %v", err)
        }
        msg := string(b_msg)
        if msg != "This is the API for the Hash API test" {
            t.Fatalf("Message returned is not the Welcome Message")
        }
    })
}
func TestInsertToken(t *testing.T) {
    // Test to Insert a token on the database
    t.Run("Insert Token test", func(t *testing.T) {
        test = "t=testtoken"
        req, err := http.NewRequest("GET", "localhost:8000/api/create_hash?" +
                                    test, nil)
        if err != nil { t.Fatalf("Request fail: %v", err) }

        rec := httptest.NewRecorder()
        CreateHash(rec, req)

        res := rec.Result()
        defer res.Body.Close()
        if res.StatusCode != http.StatusOK {
            t.Errorf("Expected Status OK, got %v", res.Status)
        }

        b_msg, err := ioutil.ReadAll(res.Body)
        if err != nil {
            t.Fatalf("Could not read response: %v", err)
        }

        msg := string(b_msg)
        switch {
        case msg == "The hash \"ada63e98fe50eccb55036d88eda4b2c3709f53c2b65bc0335797067e9a2a5d8b\" was included":
        case msg == "The hash \"ada63e98fe50eccb55036d88eda4b2c3709f53c2b65bc0335797067e9a2a5d8b\" already was been included":
        default:
            t.Fatalf("Message returned is not the right message")
        }
    })
}

func TestColsultHashs(t *testing.T) {
    // Test for querying a company in the API
    t.Run("Consult Hashs test", func(t *testing.T) {
        req, err := http.NewRequest("GET", "localhost:8080/api/return_hashs", nil)
        if err != nil { t.Fatalf("Request fail: %v", err) }

        rec := httptest.NewRecorder()
        ReturnHash(rec, req)
        res := rec.Result()
        defer res.Body.Close()
        if res.StatusCode != http.StatusOK {
            t.Errorf("Expected Status OK, got %v", res.Status)
        }

        b_msg, err := ioutil.ReadAll(res.Body)
        if err != nil {
            t.Fatalf("Could not read response: %v", err)
        }

        msg := string(b_msg)
        if len(msg) < 1 {
            t.Fatalf("Message returned is not the right message")
        }
    })
}

func TestColsultToken(t *testing.T) {
    // Test for querying a company in the API
    t.Run("Consult a token test", func(t *testing.T) {
        test = "t=ada63e98fe50eccb55036d88eda4b2c3709f53c2b65bc0335797067e9a2a5d8b"
        req, err := http.NewRequest("GET", "localhost:8080/api/return_text?" +
                                                                      test, nil)
        if err != nil { t.Fatalf("Request fail: %v", err) }

        rec := httptest.NewRecorder()
        ReturnText(rec, req)
        res := rec.Result()
        defer res.Body.Close()
        if res.StatusCode != http.StatusOK {
            t.Errorf("Expected Status OK, got %v", res.Status)
        }

        b_msg, err := ioutil.ReadAll(res.Body)
        if err != nil {
            t.Fatalf("Could not read response: %v", err)
        }

        msg := string(b_msg)
        if msg != "{\n\t\"hash\": \"ada63e98fe50eccb55036d88eda4b2c3709f53c2b65bc0335797067e9a2a5d8b\",\n\t\"token\": \"testtoken\"\n}" {
            t.Fatalf("Message returned is not the right message")
        }
    })
}

func TestColsultMissingToken(t *testing.T) {
    // Test for querying a company in the API
    t.Run("Consult a missing token test", func(t *testing.T) {
        test = "t=0000"
        req, err := http.NewRequest("GET", "localhost:8080/api/return_text?" +
                                                                      test, nil)
        if err != nil { t.Fatalf("Request fail: %v", err) }

        rec := httptest.NewRecorder()
        ReturnText(rec, req)
        res := rec.Result()
        defer res.Body.Close()
        if res.StatusCode != http.StatusOK {
            t.Errorf("Expected Status OK, got %v", res.Status)
        }

        b_msg, err := ioutil.ReadAll(res.Body)
        if err != nil {
            t.Fatalf("Could not read response: %v", err)
        }

        msg := string(b_msg)
        if msg != "{\n\t\"msg\": \"404\"\n}" {
            t.Fatalf("Message returned is not the right message")
        }
    })
}
