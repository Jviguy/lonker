package main

import (
   "github.com/go-redis/redis/v8"
   "context"
   "github.com/gin-gonic/gin"
   "encoding/json"
   "time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
    Addr: ":6379",
})

func main() {
    r := gin.Default()
    r.Static("/assets", "./assets")
    r.LoadHTMLGlob("assets/html/*")
    r.GET("/", Index)
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
    r.POST("/api/v0/upload", UploadEnpoint)
    r.Run()
}

func Index(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
        "title": "Lonker",
    })
    return
}

func UploadEnpoint(c *gin.Context) {
    json := UploadRequest{}
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    entry, err := AddUrl(json.Url)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, entry)
    return
}

func AddUrl(url string) (UploadResponse, error) {
    t := time.Now()
    id := GenerateLinkId()
    entry := Entry{Url: url, Added: &t}
    data, _ := json.Marshal(entry)
    err := rdb.Set(ctx, id, data, 0).Err()
    if err != nil {
        return UploadResponse{}, err
    }
    return UploadResponse{Id: id, Time: &t, Url: "http://104.243.44.32:8080/u/"+id}, err
}

func GetEndPoint(c *gin.Context) {
    json := GetRequest{}
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    entry, err := GetEntry(json.Id)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, entry)
}

func GetEntry(id string) (Entry, error) {
    entry := Entry{}
    data, err := rdb.Get(ctx, id).Bytes()
    if err != nil {
        //return empty entry and the error
        return entry, err
    }
    err = json.Unmarshal(data, &entry)
    if err != nil {
        return entry, err
    }
    return entry, err
}
