package main

import (
 _ "github.com/gin-gonic/gin"
 "github.com/gin-gonic/gin"
)


//func main() {
// r := gin.Default()
// r.GET(“/”, func(c *gin.Context) {
// c.String(200, “Hello World”)
// })
// r.Run()
//}

func main()  {

 r := gin.Default()
 r.GET("/", func(c *gin.Context) {

  c.JSON(200, gin.H{
   "message" : "SB",
  })

 })

 r.Run()

}