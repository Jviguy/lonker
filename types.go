package main

import (
   "time"
)

type UploadRequest struct {
    Url string `json:"url"`
}

type UploadResponse struct {
    Id string `json:"id"`
    Time *time.Time `json:"time"`
    Url string `json:"url"`
}

type Entry struct {
    Url string `json:"url"`
    Added *time.Time `json:"added"`
}
