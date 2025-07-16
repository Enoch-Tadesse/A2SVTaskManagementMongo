package models

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) Before(t time.Time) bool {
	return time.Time(d).Before(t)
}

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	var b []byte
	b = fmt.Appendf(b, "%q", t.Format("2006-01-02"))
	return b, nil
}

func (d Date) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(time.Time(d))
}

func (d *Date) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	var tm time.Time
	if err := bson.UnmarshalValue(t, data, &tm); err != nil {
		return err
	}
	*d = Date(tm)
	return nil
}
