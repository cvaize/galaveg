package db

import (
	"database/sql"
	"errors"
	"fmt"
	"galaveg/internal/infrastructures/db"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type DbRepo[DTO any, IdType any] struct {
	db               db.Db
	table            string
	columns          []string
	withoutIdColumns []string
	idColumnKey      string
	dtoMapFun        func(columns []string, values []interface{}) (*DTO, error)
	queryMapFun      func(column string, dto *DTO) (interface{}, error)
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
	withoutIdColumns := lo.Filter(settings.Columns, func(s string, _ int) bool {
		return s != settings.IdColumnKey
	})

	return &DbRepo[DTO, IdType]{
		settings.Db,
		table,
		settings.Columns,
		withoutIdColumns,
		settings.IdColumnKey,
		settings.DtoMapFun,
		settings.QueryMapFun,
	}, nil
}

func (r *DbRepo[DTO, IdType]) getColumns(columns []string) []string {
	if columns == nil || len(columns) == 0 {
		columns = r.columns
	}
	return columns
}

func (r *DbRepo[DTO, IdType]) getWithoutIdColumns(columns []string) []string {
	if columns == nil || len(columns) == 0 {
		columns = r.withoutIdColumns
	}
	return columns
}

func queryWhereClauses(query *string, whereClauses []string) {
	if whereClauses != nil && len(whereClauses) > 0 {
		*query += " WHERE " + strings.Join(whereClauses, " AND ")
	}
}

func queryOrderBy(query *string, orderBy string) {
	if len(orderBy) > 0 {
		*query += " ORDER BY " + orderBy
	}
}

func queryLimit(query *string, limit int) {
	if limit > 0 {
		*query += " LIMIT " + strconv.Itoa(limit)
	}
}

func queryOffset(query *string, offset int) {
	if offset > 0 {
		*query += " OFFSET " + strconv.Itoa(offset)
	}
}

func makeValues(values []interface{}) []interface{} {
	if values != nil {
		return values
	}
	return make([]interface{}, 0)
}

func (r *DbRepo[DTO, IdType]) First(filterValues []interface{}, whereClauses []string, columns []string, orderBy string) (*DTO, error) {
	filterValues = makeValues(filterValues)
	columns = r.getColumns(columns)
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), r.table)

	queryWhereClauses(&query, whereClauses)
	queryOrderBy(&query, orderBy)
	query += " LIMIT 1"

	vals := make([]interface{}, len(columns))
	for i := range columns {
		var ii interface{}
		vals[i] = &ii
	}

	err := r.db.QueryRow(query, filterValues...).Scan(vals...)
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

func (r *DbRepo[DTO, IdType]) All(filterValues []interface{}, whereClauses []string, columns []string, orderBy string, limit int, offset int) ([]*DTO, error) {
	filterValues = makeValues(filterValues)
	columns = r.getColumns(columns)
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), r.table)

	queryWhereClauses(&query, whereClauses)
	queryOrderBy(&query, orderBy)
	queryLimit(&query, limit)
	queryOffset(&query, offset)

	rows, err := r.db.Query(query, filterValues...)
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

func (r *DbRepo[DTO, IdType]) AllIds(filterValues []interface{}, whereClauses []string, orderBy string, limit int, offset int) ([]IdType, error) {
	filterValues = makeValues(filterValues)
	query := fmt.Sprintf("SELECT id FROM %s", r.table)

	queryWhereClauses(&query, whereClauses)
	queryOrderBy(&query, orderBy)
	queryLimit(&query, limit)
	queryOffset(&query, offset)

	rows, err := r.db.Query(query, filterValues...)
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

func (r *DbRepo[DTO, IdType]) Exists(filterValues []interface{}, whereClauses []string) (bool, error) {
	filterValues = makeValues(filterValues)
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s", r.table)

	queryWhereClauses(&query, whereClauses)

	query += " LIMIT 1)"

	var exists bool
	err := r.db.QueryRow(query, filterValues...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *DbRepo[DTO, IdType]) Insert(dto *DTO, columns []string) error {
	columns = r.getWithoutIdColumns(columns)

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

func (r *DbRepo[DTO, IdType]) Update(dto *DTO, filterValues []interface{}, whereClauses []string, columns []string) error {
	filterValues = makeValues(filterValues)
	// If there are no filters, prevent all records from being updated
	if whereClauses == nil || len(whereClauses) == 0 {
		return fmt.Errorf("filters are required for the UPDATE operation")
	}
	columns = r.getWithoutIdColumns(columns)

	// Prepare the SET part of the query
	setClauses := make([]string, len(columns))
	values := make([]interface{}, len(columns)+len(filterValues))

	var index int
	for i, col := range columns {
		v, e := r.queryMapFun(col, dto)
		if e != nil {
			return e
		}
		setClauses[i] = col + " = ?"
		values[i] = v
		index = i
	}

	for _, value := range filterValues {
		index++
		values[index] = value
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

func (r *DbRepo[DTO, IdType]) Delete(filterValues []interface{}, whereClauses []string) error {
	filterValues = makeValues(filterValues)
	if whereClauses == nil || len(whereClauses) == 0 {
		// If there are no filters, prevent all records from being deleted
		return fmt.Errorf("filters are required for the delete operation")
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s", r.table, strings.Join(whereClauses, " AND "))

	_, err := r.db.Exec(query, filterValues...)
	if err != nil {
		return err
	}

	return nil
}
