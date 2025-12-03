package db

import (
	"database/sql"
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
