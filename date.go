package date

import (
	"errors"
	"strconv"
	"time"
)

const (
	ISO8601Date = "2006-01-02"
)

// The Date type represents a date, with its underlying type being a
// [time.Time].
// It restricts itself to being [time.UTC] to avoid any time zone issues.
type Date struct {
	// we implement Date as a struct rather than simply as a time.Time to ensure
	// that a time.Time cannot be cast to a date.Date, since that would
	// undermine date.Date's guarantees
	t time.Time
}

// NewDate returns a new [Date] based on year, month and day.
func NewDate(year int, month time.Month, day int) Date {
	return Date{
		t: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

// timeToDate ensures that we only consider the year, month and day.
func timeToDate(t time.Time) Date {
	year, month, day := t.Year(), t.Month(), t.Day()
	return NewDate(year, month, day)
}

// NewDateFromTime returns a [Date] based on a [time.Time].
func NewDateFromTime(t time.Time) Date {
	return timeToDate(t)
}

// Now returns a Date through [time.Now].
func Now() Date {
	return timeToDate(time.Now())
}

// Parse returns a [Date] through [time.Parse].  Note that any time portion of
// the input will be ignored.
func Parse(layout string, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}
	return timeToDate(t), nil
}

// ParseInLocation returns a [Date] through [time.ParseInLocation].
func ParseInLocation(layout string, value string, loc *time.Location) (Date, error) {
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return Date{}, err
	}
	return timeToDate(t), nil
}

// Unix returns a [Date] through [time.Unix].
func Unix(sec int64, nsec int64) Date {
	return timeToDate(time.Unix(sec, nsec))
}

// UnixMicro returns a [Date] through [time.UnixMicro].
func UnixMicro(usec int64) Date {
	return timeToDate(time.UnixMicro(usec))
}

// UnixMilli returns a [Date] through [time.UnixMilli].
func UnixMilli(msec int64) Date {
	return timeToDate(time.UnixMilli(msec))
}

// Add adds a [time.Duration] to the [Date] and returns a new one.
//
// Note that since we restrict ourselves to days, anything less than 24 hours
// will return the same [Date], both for positive and negative durations.
//
// This function is kept for compatibility with [time.Time], it is recommended
// that you use [date.Date.AddDate] instead.
func (d Date) Add(dn time.Duration) Date {
	// to avoid problems calculating from 00:00, we convert the duration added
	// to hours, and then truncate them to an integer, so -24.5 becomes -24
	// then we divide by 24 maintaining our int type, meaning any remainder will
	// be discarded; thus 50 hours become 2 days (48 hours).
	return timeToDate(d.t.AddDate(0, 0, int(dn.Hours())/24))
}

// AddDate returns a new [Date] through [time.Time.AddDate].
func (d Date) AddDate(years int, months int, days int) Date {
	return timeToDate(d.t.AddDate(years, months, days))
}

// After compares itself to another Date through [time.Time.After].
func (d Date) After(e Date) bool {
	return d.t.After(e.t)
}

// AppendFormat calls the underlying [time.Time.AppendFormat].
func (d Date) AppendFormat(b []byte, layout string) []byte {
	return d.t.AppendFormat(b, layout)
}

// Before compares itself to another [Date] through [time.Time.Before].
func (d Date) Before(e Date) bool {
	return d.t.Before(e.t)
}

// Clock returns [time.Time.Clock], returning hours, minutes and seconds.
//
// Since the object is always at midnight, this will always return 0, 0, 0
func (d Date) Clock() (hour int, min int, sec int) {
	return d.t.Clock()
}

// Date returns the year, month and day for the [Date] through [time.Time.Date].
func (d Date) Date() (year int, month time.Month, day int) {
	return d.t.Date()
}

// Day returns the month's day through [time.Time.Day].
func (d Date) Day() int {
	return d.t.Day()
}

// Equal returns true if year, month and day are the same for both [Date].
func (d Date) Equal(e Date) bool {
	dYear, dMonth, dDay := d.Date()
	eYear, eMonth, eDay := e.Date()
	return dYear == eYear && dMonth == eMonth && dDay == eDay
}

// Format returns a representation of the [Date] through [time.Time.Format].
func (d Date) Format(layout string) string {
	return d.t.Format(layout)
}

var longMonthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

// GoString implements the [fmt.GoStringer] interface for [Date].
func (d Date) GoString() string {
	// this code is largely based on time.Time.GoString's code
	toBytes := func(i int) []byte {
		return []byte(strconv.Itoa(i))
	}
	year, month, day := d.Date()
	buf := make([]byte, 0, len("date.NewDate(9999, time.September, 31)"))
	buf = append(buf, "date.NewDate("...)
	buf = append(buf, toBytes(year)...)
	if time.January <= month && month <= time.December {
		buf = append(buf, ", time."...)
		buf = append(buf, longMonthNames[month-1]...)
	} else {
		// It's difficult to construct a date.Date with a date outside the
		// standard range but we might as well try to handle the case.
		buf = append(buf, toBytes(int(month))...)
	}
	buf = append(buf, ", "...)
	buf = append(buf, toBytes(day)...)
	buf = append(buf, ')')
	return string(buf)
}

// GobDecode implements the [gob.GobDecoder] interface.
func (d *Date) GobDecode(data []byte) error {
	return d.UnmarshalBinary(data)
}

// GobEncode implements the [gob.GobEncoder] interface.
func (d Date) GobEncode() ([]byte, error) {
	return d.MarshalBinary()
}

// Hour returns the hour through [time.Time.Hour].
//
// Since the time is always midnight, this will always return 0.
func (d Date) Hour() int {
	return d.t.Hour()
}

// ISOWeek returns the ISO week through [time.Time.ISOWeek].
func (d Date) ISOWeek() (year int, week int) {
	return d.t.ISOWeek()
}

// In is kept here for function compatibility with [time.Time],
// but it does nothing and simply returns the same [Date].
func (d Date) In(loc *time.Location) Date {
	return d
}

// IsDST returns true if in DST through [time.Time.IsDST].
func (d Date) IsDST() bool {
	return d.t.IsDST()
}

// IsZero returns true if a zero Date through [time.Time.IsZero].
func (d Date) IsZero() bool {
	return d.t.IsZero()
}

// Local is kept here for function compatibility with [time.Time],
// but it does nothing and simply returns the same [Date].
func (d Date) Local() Date {
	return d
}

// Location returns the location of the underlying [time.Time], which means it
// should always return [time.UTC].
func (d Date) Location() *time.Location {
	return d.t.Location()
}

// MarshalBinary calls through to [time.Time.MarshalBinary].
func (d Date) MarshalBinary() ([]byte, error) {
	return d.t.MarshalBinary()
}

// padInt can only handle numbers for 2 digits, since that's what it's used for
func padInt(i int) string {
	if i < 10 {
		return "0" + strconv.Itoa(i)
	}
	return strconv.Itoa(i)
}

// MarshalJSON returns a JSON string of the ISO 8601 date format.
func (d Date) MarshalJSON() ([]byte, error) {
	year, month, day := d.Date()
	return []byte(`"` + strconv.Itoa(year) + "-" + padInt(int(month)) + "-" + padInt(day) + `"`), nil
}

// MarshalText returns a string of the ISO 8601 date format.
func (d Date) MarshalText() ([]byte, error) {
	year, month, day := d.Date()
	return []byte(strconv.Itoa(year) + "-" + padInt(int(month)) + "-" + padInt(day)), nil
}

// Minute returns the minute through [time.Time.Minute].
//
// Since the time is always midnight, this will always return 0.
func (d Date) Minute() int {
	return d.t.Minute()
}

// Month returns the month of the Date through [time.Time.Month].
func (d Date) Month() time.Month {
	return d.t.Month()
}

// Nanosecond returns the nanosecond through [time.Time.Nanosecond].
//
// Since the time is always midnight, this will always return 0.
func (d Date) Nanosecond() int {
	return d.t.Nanosecond()
}

// Round rounds the Date according a time.Duration through [time.Time.Round].
//
// Since the [Date] is based on dates, any duration less than 1 day will result
// in the same [Date].
func (d Date) Round(dn time.Duration) Date {
	return timeToDate(d.t.Round(dn))
}

// Second returns the second through [time.Time.Second].
//
// Since the time is always midnight, this will always return 0.
func (d Date) Second() int {
	return d.t.Second()
}

// Returns the Date formatted using the ISO 8601 date format.
func (d Date) String() string {
	return d.t.Format(ISO8601Date)
}

// Sub subtracts another date through [time.Time.Sub].
func (d Date) Sub(e Date) time.Duration {
	return d.t.Sub(e.t)
}

// Time returns the underlying [time.Time] instance.
func (d Date) Time() time.Time {
	return d.t
}

// Truncate truncates the Date to the precision of the time.Duration through
// [time.Time.Truncate].
//
// If the [time.Duration] is less than 24 hours, it will have no impact.
func (d Date) Truncate(dn time.Duration) Date {
	return timeToDate(d.t.Truncate(dn))
}

// UTC is kept here for function compatibility with [time.Time], but otherwise
// does nothing, since the [Date] is always UTC.
func (d Date) UTC() Date {
	return d
}

// Unix returns the UNIX representation through [time.Time.Unix].
func (d Date) Unix() int64 {
	return d.t.Unix()
}

// UnixMicro returns the micro UNIX representation through
// [time.Time.UnixMicro].
func (d Date) UnixMicro() int64 {
	return d.t.UnixMicro()
}

// UnixMilli returns the milli UNIX representation through
// [time.Time.UnixMilli].
func (d Date) UnixMilli() int64 {
	return d.t.UnixMilli()
}

// UnixNano returns the nano UNIX representation through [time.Time.UnixNano].
func (d Date) UnixNano() int64 {
	return d.t.UnixNano()
}

// UnmarshalBinary unmarshals the binary time representation through
// [time.Time.UnmarshalBinary].  This implements the
// [encoding.BinaryUnmarshaler].
//
// Since this uses the underlying [time.Time.UnmarshalBinary], the format has
// not changed, but only year, month and day will be preserved.
func (d *Date) UnmarshalBinary(data []byte) error {
	err := d.t.UnmarshalBinary(data)
	if err != nil {
		return err
	}
	*d = timeToDate(d.t)
	return nil
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.  The date must be
// a quoted string in ISO 8601 date representation.
func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Date.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	t, err := time.Parse(ISO8601Date, string(data))
	if err != nil {
		return err
	}
	*d = timeToDate(t)
	return nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.  The date
// must be in the ISO 8601 date representation.
func (d *Date) UnmarshalText(data []byte) error {
	t, err := time.Parse(ISO8601Date, string(data))
	if err != nil {
		return err
	}
	*d = timeToDate(t)
	return nil
}

// Weekday returns the [time.Weekday] through [time.Time.Weekday].
func (d Date) Weekday() time.Weekday {
	return d.t.Weekday()
}

// Year returns the year through [time.Time.Year].
func (d Date) Year() int {
	return d.t.Year()
}

// YearDay returns the year day through [time.Time.YearDay].
func (d Date) YearDay() int {
	return d.t.YearDay()
}

// Zone returns the string representation of the zone through [time.Time.Zone].
// However, since the time is stored at UTC all the time, it will always be
// the same result.
func (d Date) Zone() (name string, offset int) {
	return d.t.Zone()
}

// ZoneBounds return the zone bounds through [time.Time.ZoneBounds].
func (d Date) ZoneBounds() (start time.Time, end time.Time) {
	return d.t.ZoneBounds()
}
