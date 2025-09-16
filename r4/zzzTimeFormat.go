package r4

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// https://build.fhir.org/datatypes.html#date
// does not support sending json with reduced precision eg "2006-01"
const FhirDateFormat string = "2006-01-02"

// https://build.fhir.org/datatypes.html#dateTime
// fhir spec allows many ways to format a dateTime, unfortunately
// this format chosen for no particular reason
const FhirDateTimeFormat string = "2006-01-02T15:04:05-07:00"

type FhirDate struct {
	time.Time
}

func (v FhirDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Format(FhirDateFormat))
}

//2015-02-07T13:28:17-05:00
//2017-01-01T00:00:00.000Z

func (v *FhirDate) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) || bytes.Equal(b, []byte("\"\"")) {
		return nil
		//fhir server should not send "" or "null" rather than just nothing (omitempty) but in case you try
	}
	acceptedDateFormats := [6]string{"2006",
		"2006-01",
		FhirDateFormat,
		"15:04",            // HTML time input
		"15:04:05",         // HTML time input with seconds
		"2006-01-02T15:04", // HTML datetime-local
	}
	parseErrors := [6]string{} //not sure if storing errors like this makes parsing slower but will look at it later
	trimmedString := strings.Trim(string(b), "\"")
	for i, l := range acceptedDateFormats {
		t, err := time.Parse(l, trimmedString)
		if err == nil {
			*v = FhirDate(FhirDateTime{Time: t})
			return nil
		} else {
			parseErrors[i] = err.Error()
		}
	}
	//if you're seeing this error, you probably try to unmarshal a json with date not in an accepted fhir format
	return fmt.Errorf("%#v date invalid, r4/zzzTimeFormat UnmarshalJSON could not parse into accepted date format:\n%s ", string(b), strings.Join(parseErrors[:], "\n"))
}

type FhirDateTime struct {
	time.Time
}

func (v FhirDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Format(FhirDateTimeFormat))
}

func (v *FhirDateTime) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) || bytes.Equal(b, []byte("\"\"")) {
		return nil
		//fhir server should not send "" or "null" rather than just nothing (omitempty) but in case you try
	}
	acceptedDateTimeFormats := [12]string{"2006",
		"2006-01",
		FhirDateFormat,
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z",
		FhirDateTimeFormat,
		"2006-01-02T15:04:05.999999999",
		"2006-01-02T15:04:05.999999999Z",
		"2006-01-02T15:04:05.999999999-07:00",
		"15:04",            // HTML time input
		"15:04:05",         // HTML time input with seconds
		"2006-01-02T15:04", // HTML datetime-local
	}
	parseErrors := [12]string{} //not sure if storing errors like this makes parsing slower but will look at it later
	for i, l := range acceptedDateTimeFormats {
		t, err := time.Parse(l, strings.Trim(string(b), "\""))
		if err == nil {
			*v = FhirDateTime{Time: t}
			return nil
		} else {
			parseErrors[i] = err.Error()
		}
	}
	//if you're seeing this error, you probably try to unmarshal a json with dateTime not in an accepted fhir format
	return fmt.Errorf("%#v dateTime invalid, r4/zzzTimeFormat UnmarshalJSON could not parse into accepted dateTime format:\n%s ", string(b), strings.Join(parseErrors[:], "\n"))
}
