models

r4, r4b, r5

client

create, read, update, delete, patch, search

#r4

	note := r4.Annotation{Text: "it's burning up"}
	id := "abc"
	a := r4.AllergyIntolerance{Id: &id, Note: []r4.Annotation{note}}
	note2 := r4.Annotation{Text: "fhir!"}
	a.Note = append(a.Note, note2)
	j, err := json.Marshal(a)
	if err != nil {
		t.Error("r4 marshal allergy into json")
	}
	var rt rtTest
	err = json.Unmarshal(j, &rt)
	if err != nil {
		t.Error("r4 unmarshal allergy json into resourceType")
	}
	if rt.ResourceType != "AllergyIntolerance" {
		t.Error("r4 set resource type on json.marshal")
	}
	var a2 r4.AllergyIntolerance
	err = json.Unmarshal(j, &a2)
	if err != nil {
		t.Error("r4 unmarshal allergy json into new allergy")
	}
	if *a2.Id != "abc" {
		t.Error("r4 get original id after marshal/unmarshal")
	}
	if a2.Note[0].Text != "it's burning up" || a2.Note[1].Text != "fhir!" {
		t.Error("r4 get original note text after marshal/unmarshal")
	}

#client

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
	bday := "1849-07-12"
	newPat.BirthDate = &bday
	_, err = client.UpdatePatient(newPat)
	if err != nil {
		t.Error("Error updating patient", err)
	}
	patAgain, err := client.ReadPatient(*newPat.Id)
	if err != nil {
		t.Error("Error reading patient again", err)
	}
	if *patAgain.BirthDate != bday {
		t.Error("Birthday doesnt match")
	}
	fmt.Println(patAgain)
	oo, err := client.DeletePatient(newPat)
	if err != nil {
		t.Error("Error deleting patient", err)
	}
	fmt.Println("deleted?", oo)
	_, err = client.ReadPatient(*newPat.Id)
	if err == nil {
		t.Error("SHOULD get an error here because we're reading deleted patient")
	}
	fmt.Println("note: this SHOULD say deleted because we're getting patient we just deleted")
	fmt.Println(err)

#search

	client := fhirClient.New("hapi.fhir.org/baseR4/")
	allPatients, err := fhirClient.Search(fhirClient.SpPatient{Name: "a"}, client)
	if err != nil {
		t.Error("Search", err)
	} else {
		for _, e := range allPatients.Entry {
			switch pat := e.Resource.(type) {
			case *r4.Patient:
				fmt.Println("pat is", pat)
			default:
				t.Error("got resource of different type, not patient")
			}
		}
	}

	resGroup, err := fhirClient.SearchGrouped(fhirClient.SpPatient{Name: "a"}, client)
	if err != nil {
		t.Error("SearchGrouped", err)
	}
	pats := resGroup.Patient_list
	for _, p := range pats {
		fmt.Println(p)
	}

generate

go run ./bultaoreune
