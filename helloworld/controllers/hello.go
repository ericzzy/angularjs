package controllers

import (
    "net/http"

    web "github.com/gin-gonic/gin"
)

func ShowHelloController(c *web.Context) {
    c.JSON(http.StatusOK, web.H{"message": "Hello World", "status": http.StatusOK})
}
