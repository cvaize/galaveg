package roles

import (
	"fmt"
	"galaveg/internal/infrastructures/db"
	dbModule "galaveg/internal/modules/db"
	"strings"
)

type DbRepo = *DbRepoImpl

type DbRepoImpl struct {
	dbRepo *dbModule.DbRepo[*RoleDto]
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
	dbRepo, e := dbModule.NewDbRepo[*RoleDto](dbModule.DbRepoSettings[*RoleDto]{
		Db:      settings.Db,
		Table:   settings.Table,
		Prefix:  settings.Prefix,
		Columns: []string{"code", "name", "description", "permissions"},
		MapFun: func(column string, dto *RoleDto) (interface{}, error) {
			switch column {
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

// Insert creates a new role in the database
func (r *DbRepoImpl) Insert(role *RoleDto, columns []string) error {
	return r.dbRepo.Insert(role, columns)
}

// Update updates role entries according to filters
func (r *DbRepoImpl) Update(role *RoleDto, filters map[string]interface{}, columns []string) error {
	return r.dbRepo.Update(role, filters, columns)
}

// Delete удаляет записи ролей согласно фильтрам
func (r *DbRepoImpl) Delete(filters map[string]interface{}) error {
	return r.dbRepo.Delete(filters)
}
