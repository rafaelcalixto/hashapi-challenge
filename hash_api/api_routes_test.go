package neoway_api

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

func TestIndex_handler(t *testing.T) {
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
        if msg != "This is the API for the Neoway test" {
            t.Fatalf("Message returned is not the Welcome Message")
        }
    })
}
func TestAssociateLink(t *testing.T) {
    // Test for the Update of website in the API
    t.Run("AssociateLink test", func(t *testing.T) {
        test = "c=test&z=99999&w=www.test.com"
        req, err := http.NewRequest("GET", "localhost:8000/api/clientdata?" + test, nil)
        if err != nil { t.Fatalf("Request fail: %v", err) }

        rec := httptest.NewRecorder()
        AssociateLink(rec, req)

        res := rec.Result()
        defer res.Body.Close()
        if res.StatusCode != http.StatusOK {
            t.Errorf("Expected Status OK, got %v", res.Status)
        }

        b_msg, err := ioutil.ReadAll(res.Body)
        if err != nil {
            t.Fatalf("Could not read response: %v", err)
        }
        t.Fatalf("Message returned is not the Welcome Message")
        msg := string(b_msg)
        if msg != "Parameter not found" {
            t.Fatalf("Message returned is not the right message")
        }
    })
}
func TestConsultCompanies(t *testing.T) {
    // Test for querying a company in the API
    t.Run("ConsultCompanies test", func(t *testing.T) {
        test = "c=test&z=99999"
        req, err := http.NewRequest("GET", "localhost:8080/api/companies?" + test, nil)
        if err != nil { t.Fatalf("Request fail: %v", err) }

        rec := httptest.NewRecorder()
        AssociateLink(rec, req)
        t.Fatalf("Message returned is not the right message")
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
        if msg != "{\n\t\"msg\": \"No match results\"\n}" {
            t.Fatalf("Message returned is not the right message")
        }
    })
}
