package files

import (
	"fmt"
	"galaveg/internal/infrastructures/db"
	dbModule "galaveg/internal/modules/db"
)

type DbRepo = *DbRepoImpl

type DbRepoImpl struct {
	dbRepo *dbModule.DbRepo[UserDto, int64]
}

type DbRepoImplSettings struct {
	Db     db.Db
	Table  string
	Prefix string
}

func NewDbRepoImpl(settings DbRepoImplSettings) (*DbRepoImpl, error) {
	if settings.Table == "" {
		settings.Table = "users_files"
	}
	dbRepo, e := dbModule.NewDbRepo[UserDto, int64](dbModule.DbRepoSettings[UserDto, int64]{
		Db:          settings.Db,
		Table:       settings.Table,
		Prefix:      settings.Prefix,
		IdColumnKey: "id",
		Columns:     []string{"id", "file_id", "filename", "path", "upload_filename", "mime", "user_id", "created_at", "updated_at", "deleted_at", "is_deleted", "is_public", "disk"},
		//
		ColumnsThatShouldNotBeUpdatedByDefault: []string{"created_at", "updated_at"},
		DtoMapFun: func(columns []string, values []interface{}) (*UserDto, error) {
			var dto UserDto
			for i, column := range columns {
				value := values[i]
				var e error
				switch column {
				case "id":
					dto.Id, e, _ = dbModule.ToInt64(value)
					break
				case "file_id":
					dto.FileId, e, _ = dbModule.ToInt64(value)
					break
				case "filename":
					dto.Filename, e, _ = dbModule.ToString(value)
					break
				case "path":
					dto.Path, e, _ = dbModule.ToString(value)
					break
				case "upload_filename":
					dto.UploadFilename, e, _ = dbModule.ToString(value)
					break
				case "mime":
					dto.Mime, e, _ = dbModule.ToString(value)
					break
				case "user_id":
					dto.UserId, e, _ = dbModule.ToInt64(value)
					break
				case "created_at":
					dto.CreatedAt, e, _ = dbModule.ToTime(value)
					break
				case "updated_at":
					dto.UpdatedAt, e, _ = dbModule.ToTime(value)
					break
				case "deleted_at":
					dto.DeletedAt, e, _ = dbModule.ToTime(value)
					break
				case "is_deleted":
					dto.IsDeleted, e, _ = dbModule.ToBool(value)
					break
				case "is_public":
					dto.IsPublic, e, _ = dbModule.ToBool(value)
					break
				case "disk":
					dto.Disk, e, _ = dbModule.ToString(value)
					break
				default:
					continue
				}
				if e != nil {
					return nil, e
				}
			}
			return &dto, nil
		},
		QueryMapFun: func(column string, dto *UserDto) (interface{}, error) {
			switch column {
			case "id":
				return dto.Id, nil
			case "file_id":
				return dto.FileId, nil
			case "filename":
				return dbModule.NilIfEmptyString(dto.Filename), nil
			case "path":
				return dbModule.NilIfEmptyString(dto.Path), nil
			case "upload_filename":
				return dbModule.NilIfEmptyString(dto.UploadFilename), nil
			case "mime":
				return dbModule.NilIfEmptyString(dto.Mime), nil
			case "user_id":
				return dto.UserId, nil
			case "created_at":
				return dto.CreatedAt, nil
			case "updated_at":
				return dto.UpdatedAt, nil
			case "deleted_at":
				return dbModule.NilIfZeroTime(dto.DeletedAt), nil
			case "is_deleted":
				return dto.IsDeleted, nil
			case "is_public":
				return dto.IsPublic, nil
			case "disk":
				return dto.Disk, nil
			}
			return nil, fmt.Errorf("unknown column: %s", column)
		},
	})
	if e != nil {
		return nil, e
	}
	return &DbRepoImpl{dbRepo}, nil
}
