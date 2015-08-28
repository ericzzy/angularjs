package main

import (
    "net/http"

    web "github.com/gin-gonic/gin"
)

func main() {
    router := web.Default()
    router.StaticFS("/self/", http.Dir("./static"))

    router.Run(":8090")
}

