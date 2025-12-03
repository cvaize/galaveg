package db

import (
	"fmt"
	"galaveg/internal/infrastructures/db"
	"strings"
)

type DbRepo[DTO any] struct {
	db      db.Db
	table   string
	columns []string
	mapFun  func(column string, dto DTO) (interface{}, error)
}

type DbRepoSettings[DTO any] struct {
	Db      db.Db
	Table   string
	Prefix  string
	Columns []string
	MapFun  func(column string, dto DTO) (interface{}, error)
}

func NewDbRepo[DTO any](settings DbRepoSettings[DTO]) (*DbRepo[DTO], error) {
	if settings.Table == "" {
		return nil, fmt.Errorf("when creating a DB repository, no Table was specified")
	}
	table := settings.Table
	if settings.Prefix != "" && !strings.HasPrefix(table, settings.Prefix) {
		table = settings.Prefix + table
	}
	return &DbRepo[DTO]{settings.Db, table, settings.Columns, settings.MapFun}, nil
}

// Insert creates a new row in the database
func (r *DbRepo[DTO]) Insert(dto DTO, columns []string) error {
	if columns == nil {
		columns = r.columns
	}

	// Preparing an SQL query
	placeholders := make([]string, len(columns))
	values := make([]interface{}, len(columns))

	for i, col := range columns {
		placeholders[i] = "?"
		v, e := r.mapFun(col, dto)
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

// Update updates rows entries according to filters
func (r *DbRepo[DTO]) Update(dto DTO, filters map[string]interface{}, columns []string) error {
	// If there are no filters, prevent all records from being updated
	if filters == nil || len(filters) == 0 {
		return fmt.Errorf("filters are required for the UPDATE operation")
	}
	if columns == nil {
		columns = r.columns
	}

	// Prepare the SET part of the query
	setClauses := make([]string, len(columns))
	values := make([]interface{}, len(columns)+len(filters))

	var index int
	for i, col := range columns {
		v, e := r.mapFun(col, dto)
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

// Delete deletes rows entries according to filters
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
