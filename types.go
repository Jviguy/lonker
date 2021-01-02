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

/*
Used for the get endpoint returns a Entry
*/
type GetRequest struct {
    Id string `json:"id"`
}

/*
Used in storing urls into the database and returned when a valid get endpoint is getted.
*/
type Entry struct {
    Url string `json:"url"`
    Added *time.Time `json:"added"`
}
