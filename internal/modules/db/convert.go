package db

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
)

// ToString return (value, error, isNull)
func ToString(value any) (string, error, bool) {
	var v sql.NullString
	e := v.Scan(*value.(*interface{}))
	return v.String, e, !v.Valid
}

// ToInt64 return (value, error, isNull)
func ToInt64(value any) (int64, error, bool) {
	var v sql.NullInt64
	e := v.Scan(*value.(*interface{}))
	return v.Int64, e, !v.Valid
}

// ToInt32 return (value, error, isNull)
func ToInt32(value any) (int32, error, bool) {
	var v sql.NullInt32
	e := v.Scan(*value.(*interface{}))
	return v.Int32, e, !v.Valid
}

// ToInt16 return (value, error, isNull)
func ToInt16(value any) (int16, error, bool) {
	var v sql.NullInt16
	e := v.Scan(*value.(*interface{}))
	return v.Int16, e, !v.Valid
}

// ToBool return (value, error, isNull)
func ToBool(value any) (bool, error, bool) {
	var v sql.NullBool
	e := v.Scan(*value.(*interface{}))
	return v.Bool, e, !v.Valid
}

// ToFloat64 return (value, error, isNull)
func ToFloat64(value any) (float64, error, bool) {
	var v sql.NullFloat64
	e := v.Scan(*value.(*interface{}))
	return v.Float64, e, !v.Valid
}

// ToTime return (value, error, isNull)
func ToTime(value any) (time.Time, error, bool) {
	var v sql.NullTime
	e := v.Scan(*value.(*interface{}))
	return v.Time, e, !v.Valid
}

// ToByte return (value, error, isNull)
func ToByte(value any) (byte, error, bool) {
	var v sql.NullByte
	e := v.Scan(*value.(*interface{}))
	return v.Byte, e, !v.Valid
}

func NilIfEmptyString(value string) interface{} {
	str := strings.TrimSpace(value)
	if str == "" {
		return nil
	}
	return str
}

func NilIfZeroInt64(value int64) interface{} {
	if value == 0 {
		return nil
	}
	return value
}

func NilIfZeroTime(value time.Time) interface{} {
	if value.IsZero() {
		return nil
	}
	return value
}

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

func JsonToArrayString(value string) ([]string, error) {
	result := strings.Split(value, ",")
	for i2, permission := range result {
		result[i2] = strings.Trim(permission, " \n\t[]\"")
	}
	return result, nil
}
