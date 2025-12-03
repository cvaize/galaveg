package roles

import (
	"fmt"
	"galaveg/internal/infrastructures/db"
	dbModule "galaveg/internal/modules/db"
	"github.com/goforj/godump"
	"strings"
)

type DbRepo = *DbRepoImpl

type DbRepoImpl struct {
	dbRepo *dbModule.DbRepo[RoleDto]
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
	dbRepo, e := dbModule.NewDbRepo[RoleDto](dbModule.DbRepoSettings[RoleDto]{
		Db:      settings.Db,
		Table:   settings.Table,
		Prefix:  settings.Prefix,
		Columns: []string{"id", "code", "name", "description", "permissions"},
		DtoMapFun: func(columns []string, values []interface{}) (*RoleDto, error) {
			var dto RoleDto
			for i, column := range columns {
				value := values[i]

				//godump.Dump(column)
				//godump.Dump(value)
				var ok bool
				switch column {
				case "id":
					dto.ID, ok = (*value.(*interface{})).(ID)
					break
				case "name":
					dto.Name, ok = (*value.(*interface{})).(string)
					break
				case "code":
					godump.Dump(value)
					var v []uint8
					v, ok = (*value.(*interface{})).([]uint8)
					if ok {
						dto.Code = string(v)
					}
					break
				case "description":
					dto.Description, ok = (*value.(*interface{})).(string)
					break
				case "permissions":
					var permissions string
					permissions, ok = (*value.(*interface{})).(string)
					godump.Dump(permissions)
					break
				//case "permissions":
				default:
					continue
					//return nil, fmt.Errorf("no type conversion is specified in the column: %s", column)
				}
				if !ok {
					return nil, fmt.Errorf("couldn't convert the value to the type in the column: %s", column)
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
				return dto.Description, nil
			case "permissions":
				// Convert the permissions array to a JSON string
				permissionsJSON := "[]"
				if len(dto.Permissions) > 0 {
					permissionsJSON = fmt.Sprintf(`["%s"]`, strings.Join(dto.Permissions, `","`))
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
