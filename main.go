package main

import (
   "github.com/go-redis/redis/v8"
   "context"
   "log"
   "github.com/gin-gonic/gin"
   "time"
   "encoding/json"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
    Addr: ":6379",
})

func main() {
    time := time.Now()
    entry := Entry{Url: "https://youtube.com/", Added: &time}
    data, _ := json.Marshal(entry)
    err := rdb.Set(ctx, "TEST", data, 0).Err()
    if err != nil {
        log.Fatal(err)
    }
    r := gin.Default()
    r.GET("/u/:linkId", func(c *gin.Context){
        data, err := rdb.Get(ctx, c.Param("linkId")).Bytes()
        if err != nil {
            c.JSON(400, gin.H{"message": "Unkown Link",})
            return
        }
        entry := Entry{}
        json.Unmarshal(data, &entry)
        c.JSON(200, entry)
    })
    r.Run()
}
