package main

import (
   "github.com/go-redis/redis/v8"
   "context"
   "github.com/gin-gonic/gin"
   "encoding/json"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
    Addr: ":6379",
})

func main() {
    r := gin.Default()
    r.GET("/u/:linkId", func(c *gin.Context){
        data, err := rdb.Get(ctx, c.Param("linkId")).Bytes()
        if err != nil {
            c.JSON(404, gin.H{"message": "Unkown Link",})
            return
        }
        entry := Entry{}
        json.Unmarshal(data, &entry)
        c.Redirect(301, entry.Url)
    })
    r.POST("/api/v0/upload", func(c *gin.Context) {
        c.JSON(200, gin.H{"url": ""})
        return
    })
    r.Run()
}
