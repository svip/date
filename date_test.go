package date_test

import (
	"encoding/json"
	"github.com/svip/date"
	"testing"
	"time"
)

// Mostly only tested the functions that are re-implemented.

func TestNewDate(t *testing.T) {
	d := date.NewDate(2024, time.June, 5)
	year, month, day := d.Date()
	if year != 2024 {
		t.Fatalf("Expected year to be 2024, was %d", year)
	}
	if month != time.June {
		t.Fatalf("Expected month to be June, was %v", month)
	}
	if day != 5 {
		t.Fatalf("Expected day to be 5, was %d", day)
	}
	ut := d.Time()
	year, month, day = ut.Date()
	if year != 2024 {
		t.Fatalf("Expected year to be 2024, was %d", year)
	}
	if month != time.June {
		t.Fatalf("Expected month to be June, was %v", month)
	}
	if day != 5 {
		t.Fatalf("Expected day to be 5, was %d", day)
	}
}

func TestNewDateFromTime(t *testing.T) {
	ot := time.Date(2024, time.June, 5, 15, 6, 7, 0, time.UTC)
	d := date.NewDateFromTime(ot)
	year, month, day := d.Date()
	if year != 2024 {
		t.Fatalf("Expected year to be 2024, was %d", year)
	}
	if month != time.June {
		t.Fatalf("Expected month to be June, was %v", month)
	}
	if day != 5 {
		t.Fatalf("Expected day to be 5, was %d", day)
	}
	ut := d.Time()
	year, month, day = ut.Date()
	if year != 2024 {
		t.Fatalf("Expected year to be 2024, was %d", year)
	}
	if month != time.June {
		t.Fatalf("Expected month to be June, was %v", month)
	}
	if day != 5 {
		t.Fatalf("Expected day to be 5, was %d", day)
	}
	hours, min, sec := ut.Clock()
	if hours != 0 {
		t.Fatalf("Expected hours to be 0, was %d", hours)
	}
	if min != 0 {
		t.Fatalf("Expected min to be 0, was %d", min)
	}
	if sec != 0 {
		t.Fatalf("Expected sec to be 0, was %d", sec)
	}
}

func TestDateEqual(t *testing.T) {
	d1 := date.NewDate(2024, time.June, 5)
	t2 := time.Date(2024, time.June, 5, 15, 2, 3, 0, time.UTC)
	d2 := date.NewDateFromTime(t2)
	test := d1.Equal(d2)
	if !test {
		t.Fatal("The two dates should be equal, but was not")
	}

	d1 = date.NewDate(2024, time.June, 5)
	d2 = d1.AddDate(0, 0, 1)
	test = d1.Equal(d2)
	if test {
		t.Fatal("The two dates should not be equal, but was")
	}
}

func TestDateGoString(t *testing.T) {
	d := date.NewDate(2024, time.June, 5)
	test := d.GoString()
	if test != "date.NewDate(2024, time.June, 5)" {
		t.Fatalf("Wrong GoString result, got %v", test)
	}
}

func TestMarshalJSON(t *testing.T) {
	data := struct {
		Date date.Date `json:"date"`
	}{
		Date: date.NewDate(2024, time.June, 5),
	}
	b, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}
	if string(b) != `{"date":"2024-06-05"}` {
		t.Fatalf("Wrong JSON result, got %v", string(b))
	}
}

func TestMarshalText(t *testing.T) {
	d := date.NewDate(2024, time.June, 5)
	b, err := d.MarshalText()
	if err != nil {
		t.Error(err)
	}
	if string(b) != `2024-06-05` {
		t.Fatalf("Wrong text result, got %v", string(b))
	}
}

func TestUnmarshalJSON(t *testing.T) {
	input := []byte(`{"date":"2024-06-05"}`)
	var data struct {
		Date date.Date `json:"date"`
	}
	err := json.Unmarshal(input, &data)
	if err != nil {
		t.Error(err)
	}
	year, month, day := data.Date.Date()
	if year != 2024 {
		t.Fatalf("Expected year to be 2024, was %d", year)
	}
	if month != time.June {
		t.Fatalf("Expected month to be June, was %v", month)
	}
	if day != 5 {
		t.Fatalf("Expected day to be 5, was %d", day)
	}
}

func TestUnmarshalText(t *testing.T) {
	input := []byte(`2024-06-05`)
	var d date.Date
	err := d.UnmarshalText(input)
	if err != nil {
		t.Error(err)
	}
	year, month, day := d.Date()
	if year != 2024 {
		t.Fatalf("Expected year to be 2024, was %d", year)
	}
	if month != time.June {
		t.Fatalf("Expected month to be June, was %v", month)
	}
	if day != 5 {
		t.Fatalf("Expected day to be 5, was %d", day)
	}
}
