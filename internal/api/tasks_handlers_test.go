package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCreateTask(t *testing.T) {
	api := Application{} // Inicialização

	payload := map[string]any{ // Mock de um payload
		"title":       "Learn TDD",
		"description": "Get hands-on exp with TDD in Go",
		"priority":    8000,
	}

	body, err := json.Marshal(payload) // conversão em Json para usar na request
	if err != nil {
		t.Fatal("Failed to parse our request payload")
	}

	req := httptest.NewRequest("POST", "/api/v1/tasks", bytes.NewReader(body)) // Criando uma request
	req.Header.Set("Content-Type", "application/json")                         // Setando content-type

	rec := httptest.NewRecorder() // Criando um response

	handler := http.HandlerFunc(api.handleCreateTAsk)
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusCreated {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusCreated)
	}

	var resBody map[string]any

	err = json.Unmarshal(rec.Body.Bytes(), &resBody)

	if err != nil {
		t.Fatalf("failed to parse response body: %s\n", err.Error())
	}

	if resBody["title"] != payload["title"] {
		t.Errorf("title differs; got: %q | want: %q", resBody["title"], payload["title"])
	}
}
