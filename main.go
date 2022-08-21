package main

import (
"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/characters", GetCharacters)
	r.GET("/characters/:id", GetCharacter)
	r.POST("/characters", CreateCharacter)
	r.PUT("/characters/:id", UpdateCharacters)
	r.DELETE("/characters/:id", DeleteCharacter)
	r.Run(":8080")
}

