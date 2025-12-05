package db

import (
	"database/sql"
	"errors"
	"fmt"
	"galaveg/internal/infrastructures/db"
	"github.com/samber/lo"
	"math"
	"slices"
	"strconv"
	"strings"
)

// DbRepo is a generic database repository that provides CRUD operations for any DTO type.
// It supports pagination, filtering, ordering, and other common database operations.
// Type parameters:
// - DTO: The data transfer object type that maps to database records
// - IdType: The type of the ID field (usually int64)
type DbRepo[DTO any, IdType any] struct {
	db            db.Db                                                      // Database connection
	table         string                                                     // Table name (including prefix)
	selectColumns []string                                                   // Default columns for SELECT queries
	updateColumns []string                                                   // Default columns for INSERT and UPDATE queries
	idColumnKey   string                                                     // Name of the ID column (default: "id")
	dtoMapFun     func(columns []string, values []interface{}) (*DTO, error) // Function to map database rows to DTO
	queryMapFun   func(column string, dto *DTO) (interface{}, error)         // Function to extract values from DTO for queries
}

// DbRepoSettings contains configuration parameters for creating a new DbRepo.
type DbRepoSettings[DTO any, IdType any] struct {
	Db                                     db.Db                                                      // Database connection
	Table                                  string                                                     // Table name (without prefix)
	Prefix                                 string                                                     // Database table prefix
	Columns                                []string                                                   // All table columns
	ColumnsThatShouldNotBeUpdatedByDefault []string                                                   // Columns excluded from UPDATE by default (e.g., created_at)
	IdColumnKey                            string                                                     // Name of the ID column
	DtoMapFun                              func(columns []string, values []interface{}) (*DTO, error) // DTO mapping function
	QueryMapFun                            func(column string, dto *DTO) (interface{}, error)         // Value extraction function
}

// NewDbRepo creates a new generic database repository with the provided settings.
// It validates the settings and initializes the repository structure.
func NewDbRepo[DTO any, IdType any](settings DbRepoSettings[DTO, IdType]) (*DbRepo[DTO, IdType], error) {
	// Validate required parameters
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
		settings.IdColumnKey = "id" // Default ID column name
	}

	// Apply table prefix if provided
	table := settings.Table
	if settings.Prefix != "" && !strings.HasPrefix(table, settings.Prefix) {
		table = settings.Prefix + table
	}

	// Determine which columns should be updated by default
	// Exclude ID column and any explicitly excluded columns
	updateColumns := lo.Filter(settings.Columns, func(s string, _ int) bool {
		return s != settings.IdColumnKey && !slices.Contains(settings.ColumnsThatShouldNotBeUpdatedByDefault, s)
	})

	return &DbRepo[DTO, IdType]{
		settings.Db,
		table,
		settings.Columns,
		updateColumns,
		settings.IdColumnKey,
		settings.DtoMapFun,
		settings.QueryMapFun,
	}, nil
}

// getColumns returns the columns to use for SELECT queries.
// If columns parameter is nil or empty, returns the default select columns.
func (r *DbRepo[DTO, IdType]) getColumns(columns []string) []string {
	if columns == nil || len(columns) == 0 {
		columns = r.selectColumns
	}
	return columns
}

// getUpdateColumns returns the columns to use for INSERT and UPDATE queries.
// If columns parameter is nil or empty, returns the default update columns.
func (r *DbRepo[DTO, IdType]) getUpdateColumns(columns []string) []string {
	if columns == nil || len(columns) == 0 {
		columns = r.updateColumns
	}
	return columns
}

// queryWhereClauses appends WHERE clauses to the SQL query.
// If whereClauses is empty, no WHERE clause is added.
func queryWhereClauses(query *string, whereClauses []string) {
	if whereClauses != nil && len(whereClauses) > 0 {
		*query += " WHERE " + strings.Join(whereClauses, " AND ")
	}
}

// queryOrderBy appends ORDER BY clause to the SQL query.
// If orderBy is empty, no ORDER BY clause is added.
func queryOrderBy(query *string, orderBy string) {
	if len(orderBy) > 0 {
		*query += " ORDER BY " + orderBy
	}
}

// queryLimit appends LIMIT clause to the SQL query.
// If limit is 0 or negative, no LIMIT clause is added.
func queryLimit(query *string, limit int) {
	if limit > 0 {
		*query += " LIMIT " + strconv.Itoa(limit)
	}
}

// queryOffset appends OFFSET clause to the SQL query.
// If offset is 0 or negative, no OFFSET clause is added.
func queryOffset(query *string, offset int) {
	if offset > 0 {
		*query += " OFFSET " + strconv.Itoa(offset)
	}
}

// makeValues ensures whereValues is not nil, returns empty slice if nil.
func makeValues(values []interface{}) []interface{} {
	if values != nil {
		return values
	}
	return make([]interface{}, 0)
}

// First retrieves a single record matching the specified filters.
// Parameters:
// - whereValues: Values for WHERE clause placeholders
// - whereClauses: WHERE conditions (e.g., ["name = ?", "age > ?"])
// - columns: Specific columns to select (nil for all)
// - orderBy: ORDER BY clause
// Returns the first matching DTO or nil if not found.
func (r *DbRepo[DTO, IdType]) First(whereValues []interface{}, whereClauses []string, columns []string, orderBy string) (*DTO, error) {
	whereValues = makeValues(whereValues)
	columns = r.getColumns(columns)

	// Build SELECT query
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), r.table)

	queryWhereClauses(&query, whereClauses)
	queryOrderBy(&query, orderBy)
	query += " LIMIT 1"

	// Prepare slice for scanning results
	vals := make([]interface{}, len(columns))
	for i := range columns {
		var ii interface{}
		vals[i] = &ii
	}

	// Execute query
	err := r.db.QueryRow(query, whereValues...).Scan(vals...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Entry not found
		}
		return nil, err
	}

	// Map database values to DTO
	dto, e := r.dtoMapFun(columns, vals)
	if e != nil {
		return nil, e
	}

	return dto, nil
}

// All retrieves all records matching the specified filters.
// Parameters:
// - whereValues: Values for WHERE clause placeholders
// - whereClauses: WHERE conditions
// - columns: Specific columns to select (nil for all)
// - orderBy: ORDER BY clause
// - limit: Maximum number of records to return (0 for no limit)
// - offset: Number of records to skip
// Returns a slice of matching DTOs.
func (r *DbRepo[DTO, IdType]) All(whereValues []interface{}, whereClauses []string, columns []string, orderBy string, limit int, offset int) ([]*DTO, error) {
	whereValues = makeValues(whereValues)
	columns = r.getColumns(columns)

	// Build SELECT query
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), r.table)

	queryWhereClauses(&query, whereClauses)
	queryOrderBy(&query, orderBy)
	queryLimit(&query, limit)
	queryOffset(&query, offset)

	// Execute query
	rows, err := r.db.Query(query, whereValues...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columnsLen := len(columns)
	var records []*DTO

	// Iterate through result set
	for rows.Next() {
		vals := make([]interface{}, columnsLen)
		for i := 0; i < columnsLen; i++ {
			var ii interface{}
			vals[i] = &ii
		}

		err = rows.Scan(vals...)
		if err != nil {
			return nil, err
		}

		// Map row to DTO
		dto, e := r.dtoMapFun(columns, vals)
		if e != nil {
			return nil, e
		}
		records = append(records, dto)
	}

	return records, nil
}

// Paginate retrieves a paginated set of records.
// Parameters:
// - page: Current page number (1-indexed)
// - perPage: Number of records per page
// - whereValues: Values for WHERE clause placeholders
// - whereClauses: WHERE conditions
// - columns: Specific columns to select (nil for all)
// - orderBy: ORDER BY clause
// Returns:
// - records: Slice of DTOs for the current page
// - totalRecords: Total number of matching records
// - totalPages: Total number of pages
func (r *DbRepo[DTO, IdType]) Paginate(page int, perPage int, whereValues []interface{}, whereClauses []string, columns []string, orderBy string) ([]*DTO, int64, int, error) {
	whereValues = makeValues(whereValues)
	columns = r.getColumns(columns)

	// Build query with window function to get total count
	query := fmt.Sprintf("SELECT %s, COUNT(*) OVER () as total_records FROM %s", strings.Join(columns, ", "), r.table)
	offset := (page - 1) * perPage

	queryWhereClauses(&query, whereClauses)
	queryOrderBy(&query, orderBy)
	queryLimit(&query, perPage)
	queryOffset(&query, offset)

	rows, err := r.db.Query(query, whereValues...)
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	var totalRecords int64
	columnsLen := len(columns)
	valsLen := columnsLen + 1 // +1 for total_records column
	var records []*DTO

	// Process each row
	for rows.Next() {
		vals := make([]interface{}, valsLen)
		for i := 0; i < valsLen; i++ {
			var ii interface{}
			vals[i] = &ii
		}

		err = rows.Scan(vals...)
		if err != nil {
			return nil, 0, 0, err
		}

		// Extract total records from first row
		if totalRecords == 0 {
			count, e, _ := ToInt64(vals[columnsLen:][0])
			if e != nil {
				return nil, 0, 0, e
			}
			totalRecords = count
		}

		// Map row to DTO (exclude the total_records column)
		dto, e := r.dtoMapFun(columns, vals[:columnsLen])
		if e != nil {
			return nil, 0, 0, e
		}
		records = append(records, dto)
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalRecords) / float64(perPage)))
	return records, totalRecords, totalPages, nil
}

// Count returns the number of records matching the specified filters.
// Parameters:
// - whereValues: Values for WHERE clause placeholders
// - whereClauses: WHERE conditions
// Returns the count of matching records.
func (r *DbRepo[DTO, IdType]) Count(whereValues []interface{}, whereClauses []string) (int64, error) {
	whereValues = makeValues(whereValues)

	// Build COUNT query
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", r.table)

	queryWhereClauses(&query, whereClauses)

	var totalRecords int64
	err := r.db.QueryRow(query, whereValues...).Scan(&totalRecords)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil // No matching records
		}
		return 0, err
	}
	return totalRecords, nil
}

// AllIds retrieves IDs of records matching the specified filters.
// Parameters:
// - whereValues: Values for WHERE clause placeholders
// - whereClauses: WHERE conditions
// - orderBy: ORDER BY clause
// - limit: Maximum number of IDs to return
// - offset: Number of IDs to skip
// Returns a slice of IDs.
func (r *DbRepo[DTO, IdType]) AllIds(whereValues []interface{}, whereClauses []string, orderBy string, limit int, offset int) ([]IdType, error) {
	whereValues = makeValues(whereValues)

	// Build ID-only query
	query := fmt.Sprintf("SELECT id FROM %s", r.table)

	queryWhereClauses(&query, whereClauses)
	queryOrderBy(&query, orderBy)
	queryLimit(&query, limit)
	queryOffset(&query, offset)

	rows, err := r.db.Query(query, whereValues...)
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

// Exists checks if any record matches the specified filters.
// Parameters:
// - whereValues: Values for WHERE clause placeholders
// - whereClauses: WHERE conditions
// Returns true if at least one matching record exists.
func (r *DbRepo[DTO, IdType]) Exists(whereValues []interface{}, whereClauses []string) (bool, error) {
	whereValues = makeValues(whereValues)

	// Build EXISTS query
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s", r.table)

	queryWhereClauses(&query, whereClauses)

	query += " LIMIT 1)"

	var exists bool
	err := r.db.QueryRow(query, whereValues...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Insert creates a new record in the database.
// Parameters:
// - dto: The DTO containing data to insert
// - columns: Specific columns to insert
// Returns error if insertion fails.
func (r *DbRepo[DTO, IdType]) Insert(dto *DTO, columns []string) error {
	columns = r.getUpdateColumns(columns)

	// Prepare placeholders and values for INSERT
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

	// Build INSERT query
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		r.table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	// Execute query
	_, err := r.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

// Update modifies existing records matching the specified filters.
// Parameters:
// - dto: The DTO containing updated data
// - whereValues: Values for WHERE clause placeholders
// - whereClauses: WHERE conditions (MUST be provided for safety)
// - columns: Specific columns to update
// Returns error if update fails or if no WHERE clauses are provided.
func (r *DbRepo[DTO, IdType]) Update(dto *DTO, whereValues []interface{}, whereClauses []string, columns []string) error {
	whereValues = makeValues(whereValues)

	// Safety check: prevent accidental updates of all records
	if whereClauses == nil || len(whereClauses) == 0 {
		return fmt.Errorf("filters are required for the UPDATE operation")
	}

	columns = r.getUpdateColumns(columns)

	// Prepare SET clauses and values
	setClauses := make([]string, len(columns))
	values := make([]interface{}, len(columns)+len(whereValues))

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

	// Append filter values
	for _, value := range whereValues {
		index++
		values[index] = value
	}

	// Build UPDATE query
	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		r.table,
		strings.Join(setClauses, ", "),
		strings.Join(whereClauses, " AND "),
	)

	// Execute query
	_, err := r.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes records matching the specified filters.
// Parameters:
// - whereValues: Values for WHERE clause placeholders
// - whereClauses: WHERE conditions (MUST be provided for safety)
// Returns error if deletion fails or if no WHERE clauses are provided.
func (r *DbRepo[DTO, IdType]) Delete(whereValues []interface{}, whereClauses []string) error {
	whereValues = makeValues(whereValues)

	// Safety check: prevent accidental deletion of all records
	if whereClauses == nil || len(whereClauses) == 0 {
		return fmt.Errorf("filters are required for the delete operation")
	}

	// Build DELETE query
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", r.table, strings.Join(whereClauses, " AND "))

	// Execute query
	_, err := r.db.Exec(query, whereValues...)
	if err != nil {
		return err
	}

	return nil
}
