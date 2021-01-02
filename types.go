package main

import (
   "time"
)


type Entry struct {
    Url string `json:url`
    Added *time.Time `json:added`
}
