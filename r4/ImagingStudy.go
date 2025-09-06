package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ImagingStudy
type ImagingStudy struct {
	Id                 *string              `json:"id,omitempty"`
	Meta               *Meta                `json:"meta,omitempty"`
	ImplicitRules      *string              `json:"implicitRules,omitempty"`
	Language           *string              `json:"language,omitempty"`
	Text               *Narrative           `json:"text,omitempty"`
	Contained          []Resource           `json:"contained,omitempty"`
	Extension          []Extension          `json:"extension,omitempty"`
	ModifierExtension  []Extension          `json:"modifierExtension,omitempty"`
	Identifier         []Identifier         `json:"identifier,omitempty"`
	Status             string               `json:"status"`
	Modality           []Coding             `json:"modality,omitempty"`
	Subject            Reference            `json:"subject"`
	Encounter          *Reference           `json:"encounter,omitempty"`
	Started            *string              `json:"started,omitempty"`
	BasedOn            []Reference          `json:"basedOn,omitempty"`
	Referrer           *Reference           `json:"referrer,omitempty"`
	Interpreter        []Reference          `json:"interpreter,omitempty"`
	Endpoint           []Reference          `json:"endpoint,omitempty"`
	NumberOfSeries     *int                 `json:"numberOfSeries,omitempty"`
	NumberOfInstances  *int                 `json:"numberOfInstances,omitempty"`
	ProcedureReference *Reference           `json:"procedureReference,omitempty"`
	ProcedureCode      []CodeableConcept    `json:"procedureCode,omitempty"`
	Location           *Reference           `json:"location,omitempty"`
	ReasonCode         []CodeableConcept    `json:"reasonCode,omitempty"`
	ReasonReference    []Reference          `json:"reasonReference,omitempty"`
	Note               []Annotation         `json:"note,omitempty"`
	Description        *string              `json:"description,omitempty"`
	Series             []ImagingStudySeries `json:"series,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImagingStudy
type ImagingStudySeries struct {
	Id                *string                       `json:"id,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Uid               string                        `json:"uid"`
	Number            *int                          `json:"number,omitempty"`
	Modality          Coding                        `json:"modality"`
	Description       *string                       `json:"description,omitempty"`
	NumberOfInstances *int                          `json:"numberOfInstances,omitempty"`
	Endpoint          []Reference                   `json:"endpoint,omitempty"`
	BodySite          *Coding                       `json:"bodySite,omitempty"`
	Laterality        *Coding                       `json:"laterality,omitempty"`
	Specimen          []Reference                   `json:"specimen,omitempty"`
	Started           *string                       `json:"started,omitempty"`
	Performer         []ImagingStudySeriesPerformer `json:"performer,omitempty"`
	Instance          []ImagingStudySeriesInstance  `json:"instance,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImagingStudy
type ImagingStudySeriesPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImagingStudy
type ImagingStudySeriesInstance struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Uid               string      `json:"uid"`
	SopClass          Coding      `json:"sopClass"`
	Number            *int        `json:"number,omitempty"`
	Title             *string     `json:"title,omitempty"`
}

type OtherImagingStudy ImagingStudy

// on convert struct to json, automatically add resourceType=ImagingStudy
func (r ImagingStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImagingStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherImagingStudy: OtherImagingStudy(r),
		ResourceType:      "ImagingStudy",
	})
}

func (resource *ImagingStudy) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ImagingStudy.Id", nil)
	}
	return StringInput("ImagingStudy.Id", resource.Id)
}
func (resource *ImagingStudy) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ImagingStudy.ImplicitRules", nil)
	}
	return StringInput("ImagingStudy.ImplicitRules", resource.ImplicitRules)
}
func (resource *ImagingStudy) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ImagingStudy.Language", nil, optionsValueSet)
	}
	return CodeSelect("ImagingStudy.Language", resource.Language, optionsValueSet)
}
func (resource *ImagingStudy) T_Status() templ.Component {
	optionsValueSet := VSImagingstudy_status

	if resource == nil {
		return CodeSelect("ImagingStudy.Status", nil, optionsValueSet)
	}
	return CodeSelect("ImagingStudy.Status", &resource.Status, optionsValueSet)
}
func (resource *ImagingStudy) T_Modality(numModality int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Modality) >= numModality {
		return CodingSelect("ImagingStudy.Modality["+strconv.Itoa(numModality)+"]", nil, optionsValueSet)
	}
	return CodingSelect("ImagingStudy.Modality["+strconv.Itoa(numModality)+"]", &resource.Modality[numModality], optionsValueSet)
}
func (resource *ImagingStudy) T_Started() templ.Component {

	if resource == nil {
		return StringInput("ImagingStudy.Started", nil)
	}
	return StringInput("ImagingStudy.Started", resource.Started)
}
func (resource *ImagingStudy) T_NumberOfSeries() templ.Component {

	if resource == nil {
		return IntInput("ImagingStudy.NumberOfSeries", nil)
	}
	return IntInput("ImagingStudy.NumberOfSeries", resource.NumberOfSeries)
}
func (resource *ImagingStudy) T_NumberOfInstances() templ.Component {

	if resource == nil {
		return IntInput("ImagingStudy.NumberOfInstances", nil)
	}
	return IntInput("ImagingStudy.NumberOfInstances", resource.NumberOfInstances)
}
func (resource *ImagingStudy) T_ProcedureCode(numProcedureCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProcedureCode) >= numProcedureCode {
		return CodeableConceptSelect("ImagingStudy.ProcedureCode["+strconv.Itoa(numProcedureCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImagingStudy.ProcedureCode["+strconv.Itoa(numProcedureCode)+"]", &resource.ProcedureCode[numProcedureCode], optionsValueSet)
}
func (resource *ImagingStudy) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("ImagingStudy.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImagingStudy.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *ImagingStudy) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ImagingStudy.Description", nil)
	}
	return StringInput("ImagingStudy.Description", resource.Description)
}
func (resource *ImagingStudy) T_SeriesId(numSeries int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries {
		return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Id", nil)
	}
	return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Id", resource.Series[numSeries].Id)
}
func (resource *ImagingStudy) T_SeriesUid(numSeries int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries {
		return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Uid", nil)
	}
	return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Uid", &resource.Series[numSeries].Uid)
}
func (resource *ImagingStudy) T_SeriesNumber(numSeries int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries {
		return IntInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Number", nil)
	}
	return IntInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Number", resource.Series[numSeries].Number)
}
func (resource *ImagingStudy) T_SeriesModality(numSeries int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries {
		return CodingSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Modality", nil, optionsValueSet)
	}
	return CodingSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Modality", &resource.Series[numSeries].Modality, optionsValueSet)
}
func (resource *ImagingStudy) T_SeriesDescription(numSeries int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries {
		return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Description", nil)
	}
	return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Description", resource.Series[numSeries].Description)
}
func (resource *ImagingStudy) T_SeriesNumberOfInstances(numSeries int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries {
		return IntInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].NumberOfInstances", nil)
	}
	return IntInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].NumberOfInstances", resource.Series[numSeries].NumberOfInstances)
}
func (resource *ImagingStudy) T_SeriesBodySite(numSeries int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries {
		return CodingSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].BodySite", nil, optionsValueSet)
	}
	return CodingSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].BodySite", resource.Series[numSeries].BodySite, optionsValueSet)
}
func (resource *ImagingStudy) T_SeriesLaterality(numSeries int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries {
		return CodingSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Laterality", nil, optionsValueSet)
	}
	return CodingSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Laterality", resource.Series[numSeries].Laterality, optionsValueSet)
}
func (resource *ImagingStudy) T_SeriesStarted(numSeries int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries {
		return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Started", nil)
	}
	return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Started", resource.Series[numSeries].Started)
}
func (resource *ImagingStudy) T_SeriesPerformerId(numSeries int, numPerformer int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries || len(resource.Series[numSeries].Performer) >= numPerformer {
		return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Series[numSeries].Performer[numPerformer].Id)
}
func (resource *ImagingStudy) T_SeriesPerformerFunction(numSeries int, numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries || len(resource.Series[numSeries].Performer) >= numPerformer {
		return CodeableConceptSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Series[numSeries].Performer[numPerformer].Function, optionsValueSet)
}
func (resource *ImagingStudy) T_SeriesInstanceId(numSeries int, numInstance int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries || len(resource.Series[numSeries].Instance) >= numInstance {
		return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].Id", nil)
	}
	return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].Id", resource.Series[numSeries].Instance[numInstance].Id)
}
func (resource *ImagingStudy) T_SeriesInstanceUid(numSeries int, numInstance int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries || len(resource.Series[numSeries].Instance) >= numInstance {
		return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].Uid", nil)
	}
	return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].Uid", &resource.Series[numSeries].Instance[numInstance].Uid)
}
func (resource *ImagingStudy) T_SeriesInstanceSopClass(numSeries int, numInstance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries || len(resource.Series[numSeries].Instance) >= numInstance {
		return CodingSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].SopClass", nil, optionsValueSet)
	}
	return CodingSelect("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].SopClass", &resource.Series[numSeries].Instance[numInstance].SopClass, optionsValueSet)
}
func (resource *ImagingStudy) T_SeriesInstanceNumber(numSeries int, numInstance int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries || len(resource.Series[numSeries].Instance) >= numInstance {
		return IntInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].Number", nil)
	}
	return IntInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].Number", resource.Series[numSeries].Instance[numInstance].Number)
}
func (resource *ImagingStudy) T_SeriesInstanceTitle(numSeries int, numInstance int) templ.Component {

	if resource == nil || len(resource.Series) >= numSeries || len(resource.Series[numSeries].Instance) >= numInstance {
		return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].Title", nil)
	}
	return StringInput("ImagingStudy.Series["+strconv.Itoa(numSeries)+"].Instance["+strconv.Itoa(numInstance)+"].Title", resource.Series[numSeries].Instance[numInstance].Title)
}
