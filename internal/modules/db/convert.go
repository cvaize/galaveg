package db

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
)

// ToString converts an interface value (retrieved from the database) to a string.
// Returns (value, error, isNull).
// The isNull flag indicates whether the original value was NULL in the database.
func ToString(value any) (string, error, bool) {
	var v sql.NullString
	e := v.Scan(*value.(*interface{}))
	return v.String, e, !v.Valid
}

// ToInt64 converts an interface value to int64.
// Returns (value, error, isNull).
// Used for large integer fields (e.g., BIGINT).
func ToInt64(value any) (int64, error, bool) {
	var v sql.NullInt64
	e := v.Scan(*value.(*interface{}))
	return v.Int64, e, !v.Valid
}

// ToInt32 converts an interface value to int32.
// Returns (value, error, isNull).
// Used for 32-bit integer fields.
func ToInt32(value any) (int32, error, bool) {
	var v sql.NullInt32
	e := v.Scan(*value.(*interface{}))
	return v.Int32, e, !v.Valid
}

// ToInt16 converts an interface value to int16.
// Returns (value, error, isNull).
// Used for 16-bit integer fields.
func ToInt16(value any) (int16, error, bool) {
	var v sql.NullInt16
	e := v.Scan(*value.(*interface{}))
	return v.Int16, e, !v.Valid
}

// ToBool converts an interface value to boolean.
// Returns (value, error, isNull).
// Used for boolean/bit fields.
func ToBool(value any) (bool, error, bool) {
	var v sql.NullBool
	e := v.Scan(*value.(*interface{}))
	return v.Bool, e, !v.Valid
}

// ToFloat64 converts an interface value to float64.
// Returns (value, error, isNull).
// Used for floating-point/decimal fields.
func ToFloat64(value any) (float64, error, bool) {
	var v sql.NullFloat64
	e := v.Scan(*value.(*interface{}))
	return v.Float64, e, !v.Valid
}

// ToTime converts an interface value to time.Time.
// Returns (value, error, isNull).
// Used for datetime/timestamp fields.
func ToTime(value any) (time.Time, error, bool) {
	var v sql.NullTime
	e := v.Scan(*value.(*interface{}))
	return v.Time, e, !v.Valid
}

// ToByte converts an interface value to byte.
// Returns (value, error, isNull).
// Used for tinyint/byte fields.
func ToByte(value any) (byte, error, bool) {
	var v sql.NullByte
	e := v.Scan(*value.(*interface{}))
	return v.Byte, e, !v.Valid
}

// NilIfEmptyString returns nil if the string is empty or contains only whitespace.
// Otherwise, returns the original string.
// Used to properly handle optional string fields in database operations.
func NilIfEmptyString(value string) interface{} {
	str := strings.TrimSpace(value)
	if str == "" {
		return nil
	}
	return str
}

// NilIfZeroInt64 returns nil if the int64 value is zero.
// Otherwise, returns the original value.
// Used to properly handle optional integer fields in database operations.
func NilIfZeroInt64(value int64) interface{} {
	if value == 0 {
		return nil
	}
	return value
}

// NilIfZeroTime returns nil if the time.Time value is zero (time.IsZero()).
// Otherwise, returns the original value.
// Used to properly handle optional datetime fields in database operations.
func NilIfZeroTime(value time.Time) interface{} {
	if value.IsZero() {
		return nil
	}
	return value
}

// JsonToArrayInt64 parses a JSON array string into a slice of int64.
// The input string should be in format: "[1,2,3]" or "1,2,3".
// Used to convert JSON array fields from the database to Go slices.
func JsonToArrayInt64(value string) ([]int64, error) {
	rows := strings.Split(value, ",")
	ids := make([]int64, len(rows))
	for i, str := range rows {
		str = strings.Trim(str, " \n\t[]")
		i64, i64e := strconv.ParseInt(str, 10, 64)
		if i64e != nil {
			return nil, i64e
		}
		ids[i] = i64
	}
	return ids, nil
}

// JsonToArrayString parses a JSON array string into a slice of strings.
// The input string should be in format: '["a","b","c"]' or '"a","b","c"'.
// Used to convert JSON array fields from the database to Go string slices.
func JsonToArrayString(value string) ([]string, error) {
	result := strings.Split(value, ",")
	for i2, permission := range result {
		result[i2] = strings.Trim(permission, " \n\t[]\"'")
	}
	return result, nil
}
