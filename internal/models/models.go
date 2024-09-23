package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	Post      []Post
	Photo     []Photo
	Review    []Review
	Message   []Message
	PageIds   []PageId
	CreatedAt time.Time
}

type Post struct {
	Id             int
	Title          string
	Content        string
	CreatedAt      time.Time
	UpdatedAt      sql.NullTime
	AiPhotoDesc    string
	Image          string
	User           []User
	Photos         []Photo
	PublishedPosts []PublishedPost
	ScheduledPosts []SchedulePost
}

type Photo struct {
	Id          int
	Title       string
	Description string
	Url         string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
	Post        []Post
}

type Review struct {
	Id        int
	ReviewId  string
	Platform  string
	Status    string
	CreatedAt time.Time
	User      []User
}

type Message struct {
	Id           int
	Timestamp    string
	Sender_id    string
	Recipient_id string
	Message      string
	Mid          string
	Is_read      bool
	Replied      bool
	Thread_id    int
	User         []User
}

type PageId struct {
	Id          int
	PageId      string
	Title       string
	AccessToken string
	User        []User
}

type PublishedPost struct {
	Id             int
	Platform       string
	ExternalPostId string
	PublishedAt    time.Time
	Post           []Post
}

type SchedulePost struct {
	Id                int
	ScheduledAt       time.Time
	Platform          string
	PlatformAccountId string
	Published         bool
	Post              []Post
}
