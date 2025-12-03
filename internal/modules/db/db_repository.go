package db

import (
	"database/sql"
	"errors"
	"fmt"
	"galaveg/internal/infrastructures/db"
	"strings"
)

type DbRepo[DTO any] struct {
	db          db.Db
	table       string
	columns     []string
	dtoMapFun   func(columns []string, values []interface{}) (*DTO, error)
	queryMapFun func(column string, dto *DTO) (interface{}, error)
}

type DbRepoSettings[DTO any] struct {
	Db          db.Db
	Table       string
	Prefix      string
	Columns     []string
	DtoMapFun   func(columns []string, values []interface{}) (*DTO, error)
	QueryMapFun func(column string, dto *DTO) (interface{}, error)
}

func NewDbRepo[DTO any](settings DbRepoSettings[DTO]) (*DbRepo[DTO], error) {
	if settings.Db == nil {
		return nil, fmt.Errorf("when creating a DB repository, no Db was specified")
	}
	if settings.DtoMapFun == nil {
		return nil, fmt.Errorf("when creating a DB repository, no DtoMapFun was specified")
	}
	if settings.QueryMapFun == nil {
		return nil, fmt.Errorf("when creating a DB repository, no QueryMapFun was specified")
	}
	if settings.Table == "" {
		return nil, fmt.Errorf("when creating a DB repository, no Table was specified")
	}
	table := settings.Table
	if settings.Prefix != "" && !strings.HasPrefix(table, settings.Prefix) {
		table = settings.Prefix + table
	}
	return &DbRepo[DTO]{settings.Db, table, settings.Columns, settings.DtoMapFun, settings.QueryMapFun}, nil
}

func (r *DbRepo[DTO]) First(filters map[string]interface{}, columns []string) (*DTO, error) {
	if columns == nil || len(columns) == 0 {
		columns = r.columns
	}
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), r.table)

	values := make([]interface{}, len(filters))

	if filters != nil && len(filters) > 0 {
		whereClauses := make([]string, len(filters))

		i := -1
		for field, value := range filters {
			i++
			whereClauses[i] = fmt.Sprintf("%s = ?", field)
			values[i] = value
		}

		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += " LIMIT 1"

	vals := make([]interface{}, len(columns))
	for i := range columns {
		var ii interface{}
		vals[i] = &ii
	}

	err := r.db.QueryRow(query, values...).Scan(vals...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Entry not found
		}
		return nil, err
	}

	dto, e := r.dtoMapFun(columns, vals)
	if e != nil {
		return nil, e
	}

	return dto, nil
}

func (r *DbRepo[DTO]) Exists(filters map[string]interface{}) (bool, error) {
	// Building a query
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s", r.table)

	// Adding filters
	values := make([]interface{}, len(filters))

	if filters != nil && len(filters) > 0 {
		whereClauses := make([]string, len(filters))

		i := -1
		for field, value := range filters {
			i++
			whereClauses[i] = fmt.Sprintf("%s = ?", field)
			values[i] = value
		}
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += " LIMIT 1)"

	var exists bool
	err := r.db.QueryRow(query, values...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *DbRepo[DTO]) Insert(dto *DTO, columns []string) error {
	if columns == nil || len(columns) == 0 {
		columns = r.columns
	}

	// Preparing an SQL query
	placeholders := make([]string, len(columns))
	values := make([]interface{}, len(columns))

	for i, col := range columns {
		placeholders[i] = "?"
		v, e := r.queryMapFun(col, dto)
		if e != nil {
			return e
		}
		values[i] = v
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		r.table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DbRepo[DTO]) Update(dto *DTO, filters map[string]interface{}, columns []string) error {
	// If there are no filters, prevent all records from being updated
	if filters == nil || len(filters) == 0 {
		return fmt.Errorf("filters are required for the UPDATE operation")
	}
	if columns == nil || len(columns) == 0 {
		columns = r.columns
	}

	// Prepare the SET part of the query
	setClauses := make([]string, len(columns))
	values := make([]interface{}, len(columns)+len(filters))

	var index int
	for i, col := range columns {
		v, e := r.queryMapFun(col, dto)
		if e != nil {
			return e
		}
		values[i] = v
		index = i
	}

	// Build the WHERE part of the query
	whereClauses := make([]string, len(filters))

	i := -1
	for field, value := range filters {
		i++
		index++
		values[index] = value
		whereClauses[i] = fmt.Sprintf("%s = ?", field)
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		r.table,
		strings.Join(setClauses, ", "),
		strings.Join(whereClauses, " AND "),
	)

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DbRepo[DTO]) Delete(filters map[string]interface{}) error {
	if filters == nil || len(filters) == 0 {
		// If there are no filters, prevent all records from being deleted
		return fmt.Errorf("filters are required for the delete operation")
	}

	// Building a query
	query := fmt.Sprintf("DELETE FROM %s", r.table)

	// Adding filters
	whereClauses := make([]string, len(filters))
	values := make([]interface{}, len(filters))

	i := -1
	for field, value := range filters {
		i++
		whereClauses[i] = fmt.Sprintf("%s = ?", field)
		values[i] = value
	}
	query += " WHERE " + strings.Join(whereClauses, " AND ")

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}
