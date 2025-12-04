package roles

import (
	"fmt"
	"galaveg/internal/infrastructures/db"
	dbModule "galaveg/internal/modules/db"
	"strings"
)

type DbRepo = *DbRepoImpl

type DbRepoImpl struct {
	dbRepo *dbModule.DbRepo[RoleDto, ID]
}

type DbRepoImplSettings struct {
	Db     db.Db
	Table  string
	Prefix string
}

func NewDbRepoImpl(settings DbRepoImplSettings) (*DbRepoImpl, error) {
	if settings.Table == "" {
		settings.Table = "roles"
	}
	dbRepo, e := dbModule.NewDbRepo[RoleDto, ID](dbModule.DbRepoSettings[RoleDto, ID]{
		Db:          settings.Db,
		Table:       settings.Table,
		Prefix:      settings.Prefix,
		IdColumnKey: "id",
		Columns:     []string{"id", "code", "name", "description", "permissions", "created_at", "updated_at"},
		//
		ColumnsThatShouldNotBeUpdatedByDefault: []string{"created_at", "updated_at"},
		DtoMapFun: func(columns []string, values []interface{}) (*RoleDto, error) {
			var dto RoleDto
			for i, column := range columns {
				value := values[i]

				var e error
				switch column {
				case "id":
					dto.ID, e, _ = dbModule.ToInt64(value)
					break
				case "name":
					dto.Name, e, _ = dbModule.ToString(value)
					break
				case "code":
					dto.Code, e, _ = dbModule.ToString(value)
					break
				case "description":
					dto.Description, e, _ = dbModule.ToString(value)
					break
				case "permissions":
					var vStr string
					vStr, e, _ = dbModule.ToString(value)
					if vStr != "" {
						dto.Permissions, e = dbModule.JsonToArrayString(vStr)
					}
					break
				case "created_at":
					dto.CreatedAt, e, _ = dbModule.ToTime(value)
					break
				case "updated_at":
					dto.UpdatedAt, e, _ = dbModule.ToTime(value)
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
		QueryMapFun: func(column string, dto *RoleDto) (interface{}, error) {
			switch column {
			case "id":
				return dto.ID, nil
			case "code":
				return dto.Code, nil
			case "name":
				return dto.Name, nil
			case "description":
				return dbModule.NilIfEmptyString(dto.Description), nil
			case "permissions":
				if len(dto.Permissions) == 0 {
					return nil, nil
				}
				return "[\"" + strings.Join(dto.Permissions, `","`) + "\"]", nil
			case "created_at":
				return dto.CreatedAt, nil
			case "updated_at":
				return dto.UpdatedAt, nil
			}
			return nil, fmt.Errorf("unknown column: %s", column)
		},
	})
	if e != nil {
		return nil, e
	}
	return &DbRepoImpl{dbRepo}, nil
}
