package files

import (
	"time"
)

type UserDto struct {
	Id             int64
	FileId         int64
	Filename       string
	Path           string
	UploadFilename string
	Mime           string
	UserId         int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	IsDeleted      bool
	IsPublic       bool
	Disk           string
}
