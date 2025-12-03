package db

import (
	"database/sql"
	"errors"
	"fmt"
	"galaveg/internal/infrastructures/db"
	"strconv"
	"strings"
)

var emptyDbRepoQuery = &DbRepoQuery{}

type DbRepo[DTO any, IdType any] struct {
	db          db.Db
	table       string
	columns     []string
	idColumnKey string
	dtoMapFun   func(columns []string, values []interface{}) (*DTO, error)
	queryMapFun func(column string, dto *DTO) (interface{}, error)
}

type DbRepoSettings[DTO any, IdType any] struct {
	Db          db.Db
	Table       string
	Prefix      string
	Columns     []string
	IdColumnKey string
	DtoMapFun   func(columns []string, values []interface{}) (*DTO, error)
	QueryMapFun func(column string, dto *DTO) (interface{}, error)
}

func NewDbRepo[DTO any, IdType any](settings DbRepoSettings[DTO, IdType]) (*DbRepo[DTO, IdType], error) {
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
	if settings.IdColumnKey == "" {
		settings.IdColumnKey = "id"
	}
	table := settings.Table
	if settings.Prefix != "" && !strings.HasPrefix(table, settings.Prefix) {
		table = settings.Prefix + table
	}
	return &DbRepo[DTO, IdType]{
		settings.Db,
		table,
		settings.Columns,
		settings.IdColumnKey,
		settings.DtoMapFun,
		settings.QueryMapFun,
	}, nil
}

func (r *DbRepo[DTO, IdType]) prepare(q *DbRepoQuery) (*DbRepoQuery, []string) {
	if q == nil {
		q = emptyDbRepoQuery
	}
	columns := q.Columns
	if columns == nil || len(columns) == 0 {
		columns = r.columns
	}
	return q, columns
}

func (r *DbRepo[DTO, IdType]) First(q *DbRepoQuery) (*DTO, error) {
	q, columns := r.prepare(q)
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), r.table)

	query += q.getQuery()
	values := q.getValues()

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

// All
//
// Example:
//
//	search := "%ADMIN%"
//	values := make([]interface{}, 2)
//	values[0] = search
//	values[1] = search
//	filters := &dbModule.DbRepoFilters{[]string{"(name like ? or description like ?)"}, values}
//	query := &dbModule.DbRepoQuery{
//	    Filters: filters,
//	}
func (r *DbRepo[DTO, IdType]) All(q *DbRepoQuery) ([]*DTO, error) {
	q, columns := r.prepare(q)
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), r.table)

	query += q.getQuery()
	values := q.getValues()

	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dtos []*DTO
	for rows.Next() {
		vals := make([]interface{}, len(columns))
		for i := range columns {
			var ii interface{}
			vals[i] = &ii
		}

		err = rows.Scan(vals...)
		if err != nil {
			return nil, err
		}

		dto, e := r.dtoMapFun(columns, vals)
		if e != nil {
			return nil, e
		}
		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (r *DbRepo[DTO, IdType]) AllIds(q *DbRepoQuery) ([]IdType, error) {
	q, _ = r.prepare(q)
	query := fmt.Sprintf("SELECT id FROM %s", r.table)

	query += q.getQuery()
	values := q.getValues()

	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []IdType
	for rows.Next() {
		var id IdType
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (r *DbRepo[DTO, IdType]) Exists(q *DbRepoQuery) (bool, error) {
	q, _ = r.prepare(q)
	// Building a query
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s", r.table)

	query += q.getQuery()
	values := q.getValues()

	query += " LIMIT 1)"

	var exists bool
	err := r.db.QueryRow(query, values...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *DbRepo[DTO, IdType]) Insert(dto *DTO, columns []string) error {
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

func (r *DbRepo[DTO, IdType]) Update(dto *DTO, filters *DbRepoFilters, columns []string) error {
	// If there are no filters, prevent all records from being updated
	if filters == nil || len(filters.WhereClauses) == 0 {
		return fmt.Errorf("filters are required for the UPDATE operation")
	}
	if columns == nil || len(columns) == 0 {
		columns = r.columns
	}

	// Prepare the SET part of the query
	setClauses := make([]string, len(columns))
	filterValues := filters.getValues()
	values := make([]interface{}, len(columns)+len(filterValues))

	var index int
	for i, col := range columns {
		v, e := r.queryMapFun(col, dto)
		if e != nil {
			return e
		}
		values[i] = v
		index = i
	}

	for _, value := range filterValues {
		index++
		values[index] = value
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s %s",
		r.table,
		strings.Join(setClauses, ", "),
		filters.getWhereClauses(),
	)

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DbRepo[DTO, IdType]) Delete(filters *DbRepoFilters) error {
	if filters == nil || len(filters.WhereClauses) == 0 {
		// If there are no filters, prevent all records from being deleted
		return fmt.Errorf("filters are required for the delete operation")
	}

	query := fmt.Sprintf("DELETE FROM %s", r.table)

	query += filters.getWhereClauses()
	values := filters.getValues()

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

type DbRepoFilters struct {
	WhereClauses []string
	Values       []interface{}
}

type DbRepoQuery struct {
	Filters *DbRepoFilters
	Columns []string
	OrderBy string
	Limit   int
}

func (q *DbRepoQuery) getValues() []interface{} {
	var values []interface{}

	if q.Filters != nil {
		values = q.Filters.getValues()
	}

	return values
}

func (q *DbRepoQuery) getQuery() string {
	query := ""

	if q.Filters != nil {
		query += q.Filters.getWhereClauses()
	}

	if q.OrderBy != "" {
		query += " ORDER BY " + q.OrderBy
	}

	if q.Limit > 0 {
		query += " LIMIT " + strconv.Itoa(q.Limit)
	}

	return query
}

func (q *DbRepoFilters) getValues() []interface{} {
	return q.Values
}

func (q *DbRepoFilters) getWhereClauses() string {
	query := ""
	if len(q.WhereClauses) > 0 {
		query += " WHERE " + strings.Join(q.WhereClauses, " AND ")
	}

	return query
}
