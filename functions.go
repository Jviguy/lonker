package main

import (
   "math/rand"
   "time"
)

func GenerateLinkId() string {
        rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
        Chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
        var key = make([]rune, 7)
        for i := range key {
         key[i] = rune(Chars[rnd.Intn(len(Chars))])
     }
     return string(key)
}
