package database

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	_ "github.com/lib/pq"
)

type SlateNullString sql.NullString
type SlateNullInt64 sql.NullInt64
type SlateNullTime sql.NullTime

func (ns *SlateNullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

func (ns *SlateNullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}

func (ns *SlateNullString) Scan(value interface{}) error {
	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	if bv, err := driver.String.ConvertValue(value); err == nil {
		// if this is a string type
		if v, ok := bv.(string); ok {
			// set the value of the pointer yne to YesNoEnum(v)
			ns.String = v
			return nil
		}
	}

	return errors.New("failed to scan SlateNullString")
}

func (n *SlateNullInt64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Int64)
	}
	return json.Marshal(nil)
}

func (v *SlateNullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

func (nt *SlateNullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (nt *SlateNullTime) Scan(value interface{}) error {
	var originNt = sql.NullTime{}
	if err := originNt.Scan(value); err == nil {
		nt.Time = originNt.Time
		nt.Valid = originNt.Valid
		return nil
	}
	nt.Valid = false
	return errors.New("failed to scan SlateNullTime")
}
