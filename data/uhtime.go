package data

import "time"

//// Instead of a struct, could it just be an ALIAS for time.Time?
//// Would that make omitempty work?
type UhTime struct {
	time.Time
}

const UhFormat = "2006-01-02"

///////////////////////////////////////////////////////////////////////////////////////////////////
func (t UhTime) format() string {
	return t.Time.Format(UhFormat)
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func (t UhTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.format() + `"`), nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func (t UhTime) MarshalText() ([]byte, error) {
	return []byte(t.format()), nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func (t *UhTime) updateIfEarlier(newTime UhTime) {
	if newTime.Before(t.Time) || t.IsZero() {
		*t = newTime
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func (t *UhTime) updateIfLater(newTime UhTime) {
	if newTime.After(t.Time) || t.IsZero() {
		*t = newTime
	}
}
