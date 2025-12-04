package files

import (
	"fmt"
	"galaveg/internal/infrastructures/db"
	dbModule "galaveg/internal/modules/db"
)

type DbRepo = *DbRepoImpl

type DbRepoImpl struct {
	dbRepo *dbModule.DbRepo[FileDto, int64]
}

type DbRepoImplSettings struct {
	Db     db.Db
	Table  string
	Prefix string
}

func NewDbRepoImpl(settings DbRepoImplSettings) (*DbRepoImpl, error) {
	if settings.Table == "" {
		settings.Table = "files"
	}
	dbRepo, e := dbModule.NewDbRepo[FileDto, int64](dbModule.DbRepoSettings[FileDto, int64]{
		Db:          settings.Db,
		Table:       settings.Table,
		Prefix:      settings.Prefix,
		IdColumnKey: "id",
		Columns:     []string{"id", "filename", "path", "mime", "hash", "size", "creator_user_id", "created_at", "updated_at", "delete_at", "deleted_at", "is_delete", "is_deleted", "disk"},
		//
		ColumnsThatShouldNotBeUpdatedByDefault: []string{"created_at", "updated_at"},
		DtoMapFun: func(columns []string, values []interface{}) (*FileDto, error) {
			var dto FileDto
			for i, column := range columns {
				value := values[i]
				var e error
				switch column {
				case "id":
					dto.Id, e, _ = dbModule.ToInt64(value)
					break
				case "filename":
					dto.Filename, e, _ = dbModule.ToString(value)
					break
				case "path":
					dto.Path, e, _ = dbModule.ToString(value)
					break
				case "mime":
					dto.Mime, e, _ = dbModule.ToString(value)
					break
				case "hash":
					dto.Hash, e, _ = dbModule.ToString(value)
					break
				case "size":
					dto.Size, e, _ = dbModule.ToInt64(value)
					break
				case "creator_user_id":
					dto.CreatorUserId, e, _ = dbModule.ToInt64(value)
					break
				case "created_at":
					dto.CreatedAt, e, _ = dbModule.ToTime(value)
					break
				case "updated_at":
					dto.UpdatedAt, e, _ = dbModule.ToTime(value)
					break
				case "delete_at":
					dto.DeleteAt, e, _ = dbModule.ToTime(value)
					break
				case "deleted_at":
					dto.DeletedAt, e, _ = dbModule.ToTime(value)
					break
				case "is_delete":
					dto.IsDelete, e, _ = dbModule.ToBool(value)
					break
				case "is_deleted":
					dto.IsDeleted, e, _ = dbModule.ToBool(value)
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
		QueryMapFun: func(column string, dto *FileDto) (interface{}, error) {
			switch column {
			case "id":
				return dto.Id, nil
			case "filename":
				return dto.Filename, nil
			case "path":
				return dto.Path, nil
			case "mime":
				return dbModule.NilIfEmptyString(dto.Mime), nil
			case "hash":
				return dbModule.NilIfEmptyString(dto.Hash), nil
			case "size":
				return dbModule.NilIfZeroInt64(dto.Size), nil
			case "creator_user_id":
				return dbModule.NilIfZeroInt64(dto.CreatorUserId), nil
			case "created_at":
				return dto.CreatedAt, nil
			case "updated_at":
				return dto.UpdatedAt, nil
			case "delete_at":
				return dbModule.NilIfZeroTime(dto.DeleteAt), nil
			case "deleted_at":
				return dbModule.NilIfZeroTime(dto.DeletedAt), nil
			case "is_delete":
				return dto.IsDelete, nil
			case "is_deleted":
				return dto.IsDeleted, nil
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
