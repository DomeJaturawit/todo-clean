package util

import (
	"database/sql/driver"
	"time"
)

type TestMatchTime struct{}

// Match satisfies sqlmock.Argument interface
func (t TestMatchTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
