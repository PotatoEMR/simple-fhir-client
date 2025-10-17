package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	r4 "github.com/PotatoEMR/simple-fhir-client/r4"
	fhirClient "github.com/PotatoEMR/simple-fhir-client/r4Client"
	r4b "github.com/PotatoEMR/simple-fhir-client/r4b"
	r5 "github.com/PotatoEMR/simple-fhir-client/r5"
	"github.com/a-h/templ"
)

type rtTest struct {
	ResourceType string `json:"resourceType"`
}

func TestAllergy_r4(t *testing.T) {
	note := r4.Annotation{Text: "it's burning up"}
	id := "abc"
	a := r4.AllergyIntolerance{Id: &id, Note: []r4.Annotation{note}}
	note2 := r4.Annotation{Text: "fhir!"}
	a.Note = append(a.Note, note2)
	a.Criticality = r4.VSAllergy_intolerance_criticality[0].Code
	onsetTime := r4.FhirDateTime{Time: time.Now()}
	a.OnsetDateTime = &onsetTime
	j, err := json.Marshal(a)
	if err != nil {
		t.Error("r4 marshal allergy into json", err.Error())
	}
	fmt.Println(string(j))
	var rt rtTest
	err = json.Unmarshal(j, &rt)
	if err != nil {
		t.Error("r4 unmarshal allergy json into resourceType", err.Error())
	}
	if rt.ResourceType != "AllergyIntolerance" {
		t.Error("r4 set resource type on json.marshal")
	}
	var a2 r4.AllergyIntolerance
	err = json.Unmarshal(j, &a2)
	if err != nil {
		t.Error("r4 unmarshal allergy json into new allergy", err.Error())
	}
	if a2.Id == nil || *a2.Id != "abc" {
		t.Error("r4 get original id after marshal/unmarshal")
	}
	if a2.OnsetDateTime == nil || (*a2.OnsetDateTime).Format(r4.FhirDateTimeFormat) != onsetTime.Format(r4.FhirDateTimeFormat) {
		t.Error("r4 onset ", *a2.OnsetDateTime, " different from old onset", onsetTime)
	}
	if a2.Criticality == nil || *a2.Criticality != "low" {
		t.Error("r4 get original criticality after marshal/unmarshal...array of codings maybe awkward")
	}
	if a2.Note[0].Text != "it's burning up" || a2.Note[1].Text != "fhir!" {
		t.Error("r4 get original note text after marshal/unmarshal")
	}
	fmt.Println("does criticality look right? selected should be low")
	allergyCritField := a2.T_Criticality(nil)
	allergyCritField.Render(context.Background(), os.Stdout)
	fmt.Println("does reaction[0].severity look right? it was never set, should be no selected")
	allergyReactionSeverityField := a2.T_ReactionSeverity(0, templ.Attributes{"class": "bg-red-400 m-4 p-2", "style": "border-color: yellow"})
	allergyReactionSeverityField.Render(context.Background(), os.Stdout)
}

func TestAllergy_r4b(t *testing.T) {
	note := r4b.Annotation{Text: "it's burning up"}
	id := "abc"
	a := r4b.AllergyIntolerance{Id: &id, Note: []r4b.Annotation{note}}
	note2 := r4b.Annotation{Text: "fhir!"}
	a.Note = append(a.Note, note2)
	j, err := json.Marshal(a)
	if err != nil {
		t.Error("r4b marshal allergy into json")
	}
	var rt rtTest
	err = json.Unmarshal(j, &rt)
	if err != nil {
		t.Error("r4b unmarshal allergy json into resourceType")
	}
	if rt.ResourceType != "AllergyIntolerance" {
		t.Error("r4b set resource type on json.marshal")
	}
	var a2 r4b.AllergyIntolerance
	err = json.Unmarshal(j, &a2)
	if err != nil {
		t.Error("r4b unmarshal allergy json into new allergy")
	}
	if *a2.Id != "abc" {
		t.Error("r4b get original id after marshal/unmarshal")
	}
	if a2.Note[0].Text != "it's burning up" || a2.Note[1].Text != "fhir!" {
		t.Error("r4b get original note text after marshal/unmarshal")
	}
}

func TestAllergy_r5(t *testing.T) {
	note := r5.Annotation{Text: "it's burning up"}
	id := "abc"
	a := r5.AllergyIntolerance{Id: &id, Note: []r5.Annotation{note}}
	note2 := r5.Annotation{Text: "fhir!"}
	a.Note = append(a.Note, note2)
	j, err := json.Marshal(a)
	if err != nil {
		t.Error("r5 marshal allergy into json")
	}
	var rt rtTest
	err = json.Unmarshal(j, &rt)
	if err != nil {
		t.Error("r5 unmarshal allergy json into resourceType")
	}
	if rt.ResourceType != "AllergyIntolerance" {
		t.Error("r5 set resource type on json.marshal")
	}
	var a2 r5.AllergyIntolerance
	err = json.Unmarshal(j, &a2)
	if err != nil {
		t.Error("r5 unmarshal allergy json into new allergy")
	}
	if *a2.Id != "abc" {
		t.Error("r5 get original id after marshal/unmarshal")
	}
	if a2.Note[0].Text != "it's burning up" || a2.Note[1].Text != "fhir!" {
		t.Error("r5 get original note text after marshal/unmarshal")
	}
}

func TestClient(t *testing.T) {
	client := fhirClient.New("https://r4.smarthealthit.org")
	given := []string{"William"}
	family := "Osler"
	name := r4.HumanName{Given: given, Family: &family}
	pat := r4.Patient{Name: []r4.HumanName{name}}
	newPat, err := client.CreatePatient(&pat)
	if err != nil {
		t.Error("TestClient CreatePatient error", err)
	}
	if newPat.Id == nil {
		t.Error("Server did not set newpat id?")
	}
	fmt.Println(*newPat.Id)
	bday := r4.FhirDate{Time: time.Date(1849, 7, 12, 0, 0, 0, 0, time.UTC)}
	newPat.BirthDate = &bday
	_, err = client.UpdatePatient(newPat)
	if err != nil {
		t.Error("Error updating patient", err)
	}
	patAgain, err := client.ReadPatient(*newPat.Id)
	if err != nil {
		t.Error("Error reading patient again", err)
	}
	if patAgain.BirthDate == nil {
		t.Error("Birthday nil, did time set properly or did it get 400 OperationOutcome?")
	} else {
		bdAgain := *patAgain.BirthDate
		if bdAgain != bday {
			fmt.Println("Birthday error", bdAgain, "does not match", bday)
		}
	}
	oo, err := client.DeletePatient(newPat)
	if err != nil {
		t.Error("Error deleting patient", err)
	}
	fmt.Println("deleted?", oo)
	idkpat, err := client.ReadPatient(*newPat.Id)
	if err == nil {
		fmt.Println(idkpat.String())
		t.Error("SHOULD get an error here because we're reading deleted patient")
	} else {
		fmt.Println("note: this SHOULD say deleted because we're getting patient we just deleted")
		fmt.Println(err)
	}
}

func TestSearch(t *testing.T) {
	client := fhirClient.New("https://r4.smarthealthit.org")
	allPatients, err := client.SearchBundled(fhirClient.SpPatient{Name: "a"})
	if err != nil {
		t.Error("Search", err)
	} else {
		for i, e := range allPatients.Entry {
			if i < 5 {
				switch pat := e.Resource.(type) {
				case *r4.Patient:
					fmt.Println("pat is", pat)
				default:
					t.Error("got resource of different type, not patient")
				}
			}
		}
	}

	resGroup, err := client.SearchGrouped(fhirClient.SpPatient{Name: "a"})
	if err != nil {
		t.Error("SearchGrouped", err)
		return
	}

	pats := resGroup.Patients
	for i, p := range pats {
		if i < 3 {
			fmt.Println(p)
		}
	}

	firstId := pats[0].Id
	pEverything, _ := client.PatientEverythingBundled(*firstId)
	for i, e := range pEverything.Entry {
		if i < 3 {
			fmt.Printf("type %T\n", e.Resource)
		}
	}

	PEgrouped, _ := client.PatientEverythingGrouped(*firstId)
	for i, o := range PEgrouped.Observations {
		if i < 3 {
			fmt.Println("observation:", *o.Subject, *o.Id)
		}
	}
}

type testTimeStruct struct {
	BirthdayDate     r4.FhirDate
	BirthdayDateTime r4.FhirDateTime
}

func TestTime(t *testing.T) {
	// Valid FhirDate: full date (YYYY-MM-DD), no timezone
	jsonData1 := `{"birthdayDate": "1990-05-15"}`
	var result1 testTimeStruct
	err := json.Unmarshal([]byte(jsonData1), &result1)
	if err != nil {
		t.Errorf("Failed to unmarshal full date: %v", err)
	}
	if result1.BirthdayDate.Year() != 1990 || result1.BirthdayDate.Month() != 5 || result1.BirthdayDate.Day() != 15 {
		t.Errorf("Expected 1990-05-15, got %v", result1.BirthdayDate)
	}

	// Valid FhirDate: year-month (YYYY-MM), no timezone
	jsonData2 := `{"birthdayDate": "1990-05"}`
	var result2 testTimeStruct
	err = json.Unmarshal([]byte(jsonData2), &result2)
	if err != nil {
		t.Errorf("Failed to unmarshal year-month date: %v", err)
	}
	if result2.BirthdayDate.Year() != 1990 || result2.BirthdayDate.Month() != 5 {
		t.Errorf("Expected 1990-05, got %v", result2.BirthdayDate)
	}

	// Valid FhirDate: year only (YYYY), no timezone
	jsonData3 := `{"birthdayDate": "1990"}`
	var result3 testTimeStruct
	err = json.Unmarshal([]byte(jsonData3), &result3)
	if err != nil {
		t.Errorf("Failed to unmarshal year only date: %v", err)
	}
	if result3.BirthdayDate.Year() != 1990 {
		t.Errorf("Expected 1990, got %v", result3.BirthdayDate)
	}

	// Valid FhirDateTime: full datetime with timezone
	jsonData4 := `{"birthdayDateTime": "2015-02-07T13:28:17-05:00"}`
	var result4 testTimeStruct
	err = json.Unmarshal([]byte(jsonData4), &result4)
	if err != nil {
		t.Errorf("Failed to unmarshal full datetime: %v", err)
	}
	if result4.BirthdayDateTime.Year() != 2015 || result4.BirthdayDateTime.Month() != 2 || result4.BirthdayDateTime.Day() != 7 || result4.BirthdayDateTime.Hour() != 13 {
		t.Errorf("Expected 2015-02-07T13:28:17-05:00, got %v", result4.BirthdayDateTime)
	}

	// Valid FhirDateTime: datetime with Z timezone
	jsonData5 := `{"birthdayDateTime": "2017-01-01T00:00:00.000Z"}`
	var result5 testTimeStruct
	err = json.Unmarshal([]byte(jsonData5), &result5)
	if err != nil {
		t.Errorf("Failed to unmarshal datetime with Z: %v", err)
	}
	if result5.BirthdayDateTime.Year() != 2017 || result5.BirthdayDateTime.Month() != 1 || result5.BirthdayDateTime.Day() != 1 {
		t.Errorf("Expected 2017-01-01T00:00:00.000Z, got %v", result5.BirthdayDateTime)
	}

	// Valid FhirDateTime: just date (YYYY-MM-DD)
	jsonData6 := `{"birthdayDateTime": "1905-08-23"}`
	var result6 testTimeStruct
	err = json.Unmarshal([]byte(jsonData6), &result6)
	if err != nil {
		t.Errorf("Failed to unmarshal datetime as date: %v", err)
	}
	if result6.BirthdayDateTime.Year() != 1905 || result6.BirthdayDateTime.Month() != 8 || result6.BirthdayDateTime.Day() != 23 {
		t.Errorf("Expected 1905-08-23, got %v", result6.BirthdayDateTime)
	}

	// Valid FhirDateTime: year-month
	jsonData7 := `{"birthdayDateTime": "1973-06"}`
	var result7 testTimeStruct
	err = json.Unmarshal([]byte(jsonData7), &result7)
	if err != nil {
		t.Errorf("Failed to unmarshal datetime as year-month: %v", err)
	}
	if result7.BirthdayDateTime.Year() != 1973 || result7.BirthdayDateTime.Month() != 6 {
		t.Errorf("Expected 1973-06, got %v", result7.BirthdayDateTime)
	}

	// Valid FhirDateTime: year only
	jsonData8 := `{"birthdayDateTime": "2018"}`
	var result8 testTimeStruct
	err = json.Unmarshal([]byte(jsonData8), &result8)
	if err != nil {
		t.Errorf("Failed to unmarshal datetime as year: %v", err)
	}
	if result8.BirthdayDateTime.Year() != 2018 {
		t.Errorf("Expected 2018, got %v", result8.BirthdayDateTime)
	}

	// Valid: both fields with different precisions
	jsonData9 := `{"birthdayDate": "1990", "birthdayDateTime": "1990-12-25T14:30:15+02:00"}`
	var result9 testTimeStruct
	err = json.Unmarshal([]byte(jsonData9), &result9)
	if err != nil {
		t.Errorf("Failed to unmarshal mixed precisions: %v", err)
	}
	if result9.BirthdayDate.Year() != 1990 || result9.BirthdayDateTime.Year() != 1990 || result9.BirthdayDateTime.Month() != 12 || result9.BirthdayDateTime.Day() != 25 {
		t.Errorf("Expected 1990 and 1990-12-25T14:30:15+02:00, got %v and %v", result9.BirthdayDate, result9.BirthdayDateTime)
	}

	// Valid: leap year date
	jsonData10 := `{"birthdayDate": "2000-02-29", "birthdayDateTime": "2000-02-29T12:00:00Z"}`
	var result10 testTimeStruct
	err = json.Unmarshal([]byte(jsonData10), &result10)
	if err != nil {
		t.Errorf("Failed to unmarshal leap year: %v", err)
	}
	if result10.BirthdayDate.Year() != 2000 || result10.BirthdayDate.Month() != 2 || result10.BirthdayDate.Day() != 29 {
		t.Errorf("Expected 2000-02-29, got %v", result10.BirthdayDate)
	}

	// Invalid FhirDate: with timezone (should fail)
	jsonData11 := `{"birthdayDate": "1990-05-15T14:30:00Z"}`
	var result11 testTimeStruct
	err = json.Unmarshal([]byte(jsonData11), &result11)
	if err == nil {
		t.Errorf("Expected error for date with timezone, but got none")
	}

	// Invalid FhirDateTime: time without timezone when time is specified
	jsonData12 := `{"birthdayDateTime": "2015-02-07T13:28:17"}`
	var result12 testTimeStruct
	err = json.Unmarshal([]byte(jsonData12), &result12)
	if err != nil {
		t.Errorf(`FHIR says "If hours and minutes are specified, a timezone offset SHALL be populated" but we'll allow it`)
	}
	if result12.BirthdayDateTime.Year() != 2015 || result12.BirthdayDateTime.Month() != 2 || result12.BirthdayDateTime.Day() != 7 || result12.BirthdayDateTime.Hour() != 13 || result12.BirthdayDateTime.Minute() != 28 || result12.BirthdayDateTime.Second() != 17 {
		t.Errorf("Expected 2015-02-07T13:28:17, got %v", result12.BirthdayDateTime)
	}

	// Invalid: 24:00 time (should fail)
	jsonData13 := `{"birthdayDateTime": "2015-02-07T24:00:00Z"}`
	var result13 testTimeStruct
	err = json.Unmarshal([]byte(jsonData13), &result13)
	if err == nil {
		t.Errorf("Expected error for 24:00 time, but got none")
	}

	// Invalid: invalid date
	jsonData14 := `{"birthdayDate": "2023-02-30"}`
	var result14 testTimeStruct
	err = json.Unmarshal([]byte(jsonData14), &result14)
	if err == nil {
		t.Errorf("Expected error for invalid date, but got none")
	}

	// Invalid FhirDate: missing leading zero in month
	jsonData15 := `{"birthdayDate": "1990-5-15"}`
	var result15 testTimeStruct
	err = json.Unmarshal([]byte(jsonData15), &result15)
	if err == nil {
		t.Errorf("Expected error for month without leading zero, but got none")
	}

	// Invalid FhirDate: missing leading zero in day
	jsonData16 := `{"birthdayDate": "1990-05-5"}`
	var result16 testTimeStruct
	err = json.Unmarshal([]byte(jsonData16), &result16)
	if err == nil {
		t.Errorf("Expected error for day without leading zero, but got none")
	}

	// Invalid FhirDate: extra characters after valid date
	jsonData17 := `{"birthdayDate": "1990-05-15abc"}`
	var result17 testTimeStruct
	err = json.Unmarshal([]byte(jsonData17), &result17)
	if err == nil {
		t.Errorf("Expected error for extra characters after date, but got none")
	}

	// Invalid FhirDate: not enough digits in year (3-digit year)
	jsonData18 := `{"birthdayDate": "990"}`
	var result18 testTimeStruct
	err = json.Unmarshal([]byte(jsonData18), &result18)
	if err == nil {
		t.Errorf("Expected error for 3-digit year, but got none")
	}

	// Invalid FhirDate: too many digits in year (5-digit year)
	jsonData19 := `{"birthdayDate": "10000-01-01"}`
	var result19 testTimeStruct
	err = json.Unmarshal([]byte(jsonData19), &result19)
	if err == nil {
		t.Errorf("Expected error for 5-digit year, but got none")
	}

	// Invalid FhirDate: invalid separator
	jsonData20 := `{"birthdayDate": "1990/05/15"}`
	var result20 testTimeStruct
	err = json.Unmarshal([]byte(jsonData20), &result20)
	if err == nil {
		t.Errorf("Expected error for wrong separator, but got none")
	}

	// Invalid FhirDate: completely wrong format
	jsonData21 := `{"birthdayDate": "May 15, 1990"}`
	var result21 testTimeStruct
	err = json.Unmarshal([]byte(jsonData21), &result21)
	if err == nil {
		t.Errorf("Expected error for natural language date, but got none")
	}
}
