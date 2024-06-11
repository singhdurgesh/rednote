package datatypes

import (
	"encoding/json"
	"time"

	gorm_datatypes "gorm.io/datatypes"
)

type Date struct {
	gorm_datatypes.Date
}

const dateFormat = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) error {
	// Parse the JSON string into a time.Time
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return err
	}

	d.Date = gorm_datatypes.Date(t)

	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if time.Time(d.Date).IsZero() {
		return json.Marshal("")
	}

	return json.Marshal(time.Time(d.Date).Format(dateFormat))
}

func (d *Date) String() string {
	if time.Time(d.Date).IsZero() {
		return ""
	}
	return time.Time(d.Date).Format(dateFormat)
}
