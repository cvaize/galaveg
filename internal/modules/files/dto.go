package files

import (
	"time"
)

type FileDto struct {
	Id            int64
	Filename      string
	Path          string
	Mime          string
	Hash          string
	Size          int64
	CreatorUserId int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeleteAt      time.Time
	DeletedAt     time.Time
	IsDelete      bool
	IsDeleted     bool
	Disk          string
}
