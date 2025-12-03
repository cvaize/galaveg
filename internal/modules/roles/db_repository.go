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
		Columns:     []string{"id", "code", "name", "description", "permissions"},
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
						dto.Permissions = strings.Split(vStr, ",")
						for i2, permission := range dto.Permissions {
							dto.Permissions[i2] = strings.Trim(permission, " []\"")
						}
					}
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
				str := strings.TrimSpace(dto.Description)
				if str == "" {
					return nil, nil
				}
				return str, nil
			case "permissions":
				// Convert the permissions array to a JSON string
				permissionsJSON := "[]"
				if len(dto.Permissions) > 0 {
					permissionsJSON = fmt.Sprintf(`["%s"]`, strings.Join(dto.Permissions, `","`))
				}
				if permissionsJSON == "[]" {
					return nil, nil
				}
				return permissionsJSON, nil
			}
			return nil, fmt.Errorf("unknown column: %s", column)
		},
	})
	if e != nil {
		return nil, e
	}
	return &DbRepoImpl{dbRepo}, nil
}

func (r *DbRepoImpl) All(filters map[string]interface{}, orderBy string, columns []string) ([]*RoleDto, error) {
	return r.dbRepo.All(filters, orderBy, columns)
}

func (r *DbRepoImpl) AllIds(filters map[string]interface{}, orderBy string) ([]ID, error) {
	return r.dbRepo.AllIds(filters, orderBy)
}

func (r *DbRepoImpl) First(filters map[string]interface{}, columns []string) (*RoleDto, error) {
	return r.dbRepo.First(filters, columns)
}

func (r *DbRepoImpl) Exists(filters map[string]interface{}) (bool, error) {
	return r.dbRepo.Exists(filters)
}

func (r *DbRepoImpl) Insert(role *RoleDto, columns []string) error {
	return r.dbRepo.Insert(role, columns)
}

func (r *DbRepoImpl) Update(role *RoleDto, filters map[string]interface{}, columns []string) error {
	return r.dbRepo.Update(role, filters, columns)
}

func (r *DbRepoImpl) Delete(filters map[string]interface{}) error {
	return r.dbRepo.Delete(filters)
}
