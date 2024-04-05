package main

import (
    "testing"
    "io/ioutil"
    "net/http"
)

func testGet(t *testing.T) {
    url := "http://localhost:8000/sensors"

    // Faça a solicitação GET
    response, err := http.Get(url)
    if err != nil {
        t.Errorf("Erro ao fazer solicitação GET: %v", err)
    }
    defer response.Body.Close()

    // Leia o corpo da resposta
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        t.Errorf("Erro ao ler o corpo da resposta: %v", err)
    }

    t.Logf("Corpo da resposta: %s", body)
}

func TestRequests(t *testing.T) {
    testGet(t)
}
