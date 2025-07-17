package models

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type Date time.Time

// UnmarshalJSON is overriden method for Date type
// used to unmarshal json into concrete value
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// Before method is used to compare if a
// certain time is before the date or not
func (d Date) Before(t time.Time) bool {
	dTime := time.Time(d).Truncate(24 * time.Hour)
	tTime := t.Truncate(24 * time.Hour)
	return dTime.Before(tTime)
}

// MarshalJSON is overriden method for Date type used
// to marshal a date into json returning byte and error
func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	var b []byte
	b = fmt.Appendf(b, "%q", t.Format("2006-01-02"))
	return b, nil
}

// MarshalBSONValue is overriden method for Date type
// used to marshal a date into bson returning byte and error
func (d Date) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(time.Time(d))
}

// UnmarshalBSONValue is overriden method for Date type
// used to unmarshal bson into concrete value
func (d *Date) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	var tm time.Time
	if err := bson.UnmarshalValue(t, data, &tm); err != nil {
		return err
	}
	*d = Date(tm)
	return nil
}
