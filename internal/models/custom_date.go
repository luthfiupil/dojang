package models

import (
	"fmt"
	"strings"
	"time"
)

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) MarshalJSON() ([]byte, error) {
	if cd == nil || cd.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", cd.Format("2006-01-02"))
	return []byte(formatted), nil
}

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	if str == "null" || str == "" {
		cd.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD")
	}
	cd.Time = t
	return nil
}
