// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Chat struct {
	ID        int32
	CreatedAt sql.NullTime
}

type ChatMessage struct {
	ID       int32
	IDChat   int32
	Sender   string
	Receiver string
	DateSent time.Time
	Content  string
}

type ChatParticipant struct {
	IDChat int32
	IDUser int32
}

type User struct {
	ID             int32
	FirstName      string
	LastName       string
	Username       string
	HashedPassword string
}

type UserFriend struct {
	IDUser   int32
	IDFriend int32
}
