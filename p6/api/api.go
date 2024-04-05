package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/sensors", rUmine)
    router.GET("/sensors/:id", rFluorescent)
    router.POST("/sensors", rDoIWannaKnow)

    fmt.Println("Server will run on http://localhost:8000")
    router.Run(":8000")
}

type Sensor struct {
    Sensor string `json:"sensor"`
    Value  string `json:"value"`
}

// rUmine responde com a lista de todos os sensores em JSON.
func rUmine(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, Fireside())
}

// rDoIWannaKnow adiciona um sensor a partir do JSON recebido no corpo da solicitação.
func rDoIWannaKnow(c *gin.Context) {
    var data Sensor

    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, IWantItAll(data.Sensor, data.Value))
}

func rFluorescent(c *gin.Context) {
    id := c.Param("id")

    sensors := Fireside()

    // Percorre a lista de sensores, procurando
    // um sensor cujo valor do id corresponda ao parâmetro.
    for _, a := range sensors {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sensor not found"})
}
