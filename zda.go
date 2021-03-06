package nmea

const (
	// TypeZDA type for ZDA sentences
	TypeZDA = "ZDA"
)

// ZDA represents date & time data.
// http://aprs.gids.nl/nmea/#zda
type ZDA struct {
	BaseSentence
	Time          Time
	Day           int64
	Month         int64
	Year          int64
	OffsetHours   int64 // Local time zone offset from GMT, hours
	OffsetMinutes int64 // Local time zone offset from GMT, minutes
}

func (s ZDA) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"time":           s.Time.String(),
		"day":            s.Day,
		"month":          s.Month,
		"year":           s.Year,
		"offset_hours":   s.OffsetHours,
		"offset_minutes": s.OffsetMinutes,
	}
	bm, err := s.BaseSentence.toMap()
	if err != nil {
		return m, err
	}
	for k, v := range bm {
		m[k] = v
	}
	return m, nil
}

// newZDA constructor
func newZDA(s BaseSentence) (ZDA, error) {
	p := NewParser(s)
	p.AssertType(TypeZDA)
	return ZDA{
		BaseSentence:  s,
		Time:          p.Time(0, "time"),
		Day:           p.Int64(1, "day"),
		Month:         p.Int64(2, "month"),
		Year:          p.Int64(3, "year"),
		OffsetHours:   p.Int64(4, "offset (hours)"),
		OffsetMinutes: p.Int64(5, "offset (minutes)"),
	}, p.Err()
}
