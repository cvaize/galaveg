package users

import (
	"fmt"
	"galaveg/internal/infrastructures/db"
	dbModule "galaveg/internal/modules/db"
	"strings"
)

type DbRepo = *DbRepoImpl

type DbRepoImpl struct {
	dbRepo *dbModule.DbRepo[UserDto, ID]
}

type DbRepoImplSettings struct {
	Db     db.Db
	Table  string
	Prefix string
}

func NewDbRepoImpl(settings DbRepoImplSettings) (*DbRepoImpl, error) {
	if settings.Table == "" {
		settings.Table = "users"
	}
	dbRepo, e := dbModule.NewDbRepo[UserDto, ID](dbModule.DbRepoSettings[UserDto, ID]{
		Db:          settings.Db,
		Table:       settings.Table,
		Prefix:      settings.Prefix,
		IdColumnKey: "id",
		Columns:     []string{"id", "email", "locale", "surname", "name", "patronymic", "is_super_admin", "roles_ids", "avatar_id", "created_at", "updated_at"},
		//
		ColumnsThatShouldNotBeUpdatedByDefault: []string{"created_at", "updated_at"},
		DtoMapFun: func(columns []string, values []interface{}) (*UserDto, error) {
			var dto UserDto
			for i, column := range columns {
				value := values[i]

				var e error
				switch column {
				case "id":
					dto.ID, e, _ = dbModule.ToInt64(value)
					break
				case "email":
					dto.Email, e, _ = dbModule.ToString(value)
					break
				case "locale":
					dto.Locale, e, _ = dbModule.ToString(value)
					break
				case "surname":
					dto.Surname, e, _ = dbModule.ToString(value)
					break
				case "name":
					dto.Name, e, _ = dbModule.ToString(value)
					break
				case "patronymic":
					dto.Patronymic, e, _ = dbModule.ToString(value)
					break
				case "is_super_admin":
					dto.IsSuperAdmin, e, _ = dbModule.ToBool(value)
					break
				case "roles_ids":
					var vStr string
					vStr, e, _ = dbModule.ToString(value)
					if vStr != "" {
						dto.RolesIds, e = dbModule.JsonToArrayInt64(vStr)
					}
					break
				case "avatar_id":
					dto.AvatarId, e, _ = dbModule.ToInt64(value)
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
		QueryMapFun: func(column string, dto *UserDto) (interface{}, error) {
			switch column {
			case "id":
				return dto.ID, nil
			case "email":
				return dto.Email, nil
			case "locale":
				return dbModule.NilIfEmptyString(dto.Locale), nil
			case "surname":
				return dbModule.NilIfEmptyString(dto.Surname), nil
			case "name":
				return dbModule.NilIfEmptyString(dto.Name), nil
			case "patronymic":
				return dbModule.NilIfEmptyString(dto.Patronymic), nil
			case "is_super_admin":
				return dto.IsSuperAdmin, nil
			case "roles_ids":
				if len(dto.RolesIds) == 0 {
					return nil, nil
				}
				return strings.ReplaceAll(fmt.Sprint(dto.RolesIds), " ", ","), nil
			case "avatar_id":
				return dto.AvatarId, nil
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
