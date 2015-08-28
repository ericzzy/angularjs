package main

import (
    "net/http"

    c "github.com/ericzzy/angularjs/helloworld/controllers"

    web "github.com/gin-gonic/gin"
)

func main() {
    router := web.Default()
    router.GET("/api/v1/hello", c.ShowHelloController)
    router.StaticFS("/service/", http.Dir("./static"))
    
    router.Run(":8090")
}
