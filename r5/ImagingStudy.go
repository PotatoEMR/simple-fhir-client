package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ImagingStudy
type ImagingStudy struct {
	Id                *string              `json:"id,omitempty"`
	Meta              *Meta                `json:"meta,omitempty"`
	ImplicitRules     *string              `json:"implicitRules,omitempty"`
	Language          *string              `json:"language,omitempty"`
	Text              *Narrative           `json:"text,omitempty"`
	Contained         []Resource           `json:"contained,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Identifier        []Identifier         `json:"identifier,omitempty"`
	Status            string               `json:"status"`
	Modality          []CodeableConcept    `json:"modality,omitempty"`
	Subject           Reference            `json:"subject"`
	Encounter         *Reference           `json:"encounter,omitempty"`
	Started           *time.Time           `json:"started,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	BasedOn           []Reference          `json:"basedOn,omitempty"`
	PartOf            []Reference          `json:"partOf,omitempty"`
	Referrer          *Reference           `json:"referrer,omitempty"`
	Endpoint          []Reference          `json:"endpoint,omitempty"`
	NumberOfSeries    *int                 `json:"numberOfSeries,omitempty"`
	NumberOfInstances *int                 `json:"numberOfInstances,omitempty"`
	Procedure         []CodeableReference  `json:"procedure,omitempty"`
	Location          *Reference           `json:"location,omitempty"`
	Reason            []CodeableReference  `json:"reason,omitempty"`
	Note              []Annotation         `json:"note,omitempty"`
	Description       *string              `json:"description,omitempty"`
	Series            []ImagingStudySeries `json:"series,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImagingStudy
type ImagingStudySeries struct {
	Id                *string                       `json:"id,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Uid               string                        `json:"uid"`
	Number            *int                          `json:"number,omitempty"`
	Modality          CodeableConcept               `json:"modality"`
	Description       *string                       `json:"description,omitempty"`
	NumberOfInstances *int                          `json:"numberOfInstances,omitempty"`
	Endpoint          []Reference                   `json:"endpoint,omitempty"`
	BodySite          *CodeableReference            `json:"bodySite,omitempty"`
	Laterality        *CodeableConcept              `json:"laterality,omitempty"`
	Specimen          []Reference                   `json:"specimen,omitempty"`
	Started           *time.Time                    `json:"started,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Performer         []ImagingStudySeriesPerformer `json:"performer,omitempty"`
	Instance          []ImagingStudySeriesInstance  `json:"instance,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImagingStudy
type ImagingStudySeriesPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImagingStudy
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
func (r ImagingStudy) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ImagingStudy/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ImagingStudy"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ImagingStudy) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSImagingstudy_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ImagingStudy) T_Modality(numModality int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numModality >= len(resource.Modality) {
		return CodeableConceptSelect("Modality["+strconv.Itoa(numModality)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Modality["+strconv.Itoa(numModality)+"]", &resource.Modality[numModality], optionsValueSet, htmlAttrs)
}
func (resource *ImagingStudy) T_Started(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Started", nil, htmlAttrs)
	}
	return DateTimeInput("Started", resource.Started, htmlAttrs)
}
func (resource *ImagingStudy) T_NumberOfSeries(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("NumberOfSeries", nil, htmlAttrs)
	}
	return IntInput("NumberOfSeries", resource.NumberOfSeries, htmlAttrs)
}
func (resource *ImagingStudy) T_NumberOfInstances(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("NumberOfInstances", nil, htmlAttrs)
	}
	return IntInput("NumberOfInstances", resource.NumberOfInstances, htmlAttrs)
}
func (resource *ImagingStudy) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ImagingStudy) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesUid(numSeries int, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) {
		return StringInput("Series["+strconv.Itoa(numSeries)+"]Uid", nil, htmlAttrs)
	}
	return StringInput("Series["+strconv.Itoa(numSeries)+"]Uid", &resource.Series[numSeries].Uid, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesNumber(numSeries int, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) {
		return IntInput("Series["+strconv.Itoa(numSeries)+"]Number", nil, htmlAttrs)
	}
	return IntInput("Series["+strconv.Itoa(numSeries)+"]Number", resource.Series[numSeries].Number, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesModality(numSeries int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) {
		return CodeableConceptSelect("Series["+strconv.Itoa(numSeries)+"]Modality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Series["+strconv.Itoa(numSeries)+"]Modality", &resource.Series[numSeries].Modality, optionsValueSet, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesDescription(numSeries int, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) {
		return StringInput("Series["+strconv.Itoa(numSeries)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Series["+strconv.Itoa(numSeries)+"]Description", resource.Series[numSeries].Description, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesNumberOfInstances(numSeries int, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) {
		return IntInput("Series["+strconv.Itoa(numSeries)+"]NumberOfInstances", nil, htmlAttrs)
	}
	return IntInput("Series["+strconv.Itoa(numSeries)+"]NumberOfInstances", resource.Series[numSeries].NumberOfInstances, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesLaterality(numSeries int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) {
		return CodeableConceptSelect("Series["+strconv.Itoa(numSeries)+"]Laterality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Series["+strconv.Itoa(numSeries)+"]Laterality", resource.Series[numSeries].Laterality, optionsValueSet, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesStarted(numSeries int, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) {
		return DateTimeInput("Series["+strconv.Itoa(numSeries)+"]Started", nil, htmlAttrs)
	}
	return DateTimeInput("Series["+strconv.Itoa(numSeries)+"]Started", resource.Series[numSeries].Started, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesPerformerFunction(numSeries int, numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) || numPerformer >= len(resource.Series[numSeries].Performer) {
		return CodeableConceptSelect("Series["+strconv.Itoa(numSeries)+"]Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Series["+strconv.Itoa(numSeries)+"]Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Series[numSeries].Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesInstanceUid(numSeries int, numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) || numInstance >= len(resource.Series[numSeries].Instance) {
		return StringInput("Series["+strconv.Itoa(numSeries)+"]Instance["+strconv.Itoa(numInstance)+"].Uid", nil, htmlAttrs)
	}
	return StringInput("Series["+strconv.Itoa(numSeries)+"]Instance["+strconv.Itoa(numInstance)+"].Uid", &resource.Series[numSeries].Instance[numInstance].Uid, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesInstanceSopClass(numSeries int, numInstance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) || numInstance >= len(resource.Series[numSeries].Instance) {
		return CodingSelect("Series["+strconv.Itoa(numSeries)+"]Instance["+strconv.Itoa(numInstance)+"].SopClass", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Series["+strconv.Itoa(numSeries)+"]Instance["+strconv.Itoa(numInstance)+"].SopClass", &resource.Series[numSeries].Instance[numInstance].SopClass, optionsValueSet, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesInstanceNumber(numSeries int, numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) || numInstance >= len(resource.Series[numSeries].Instance) {
		return IntInput("Series["+strconv.Itoa(numSeries)+"]Instance["+strconv.Itoa(numInstance)+"].Number", nil, htmlAttrs)
	}
	return IntInput("Series["+strconv.Itoa(numSeries)+"]Instance["+strconv.Itoa(numInstance)+"].Number", resource.Series[numSeries].Instance[numInstance].Number, htmlAttrs)
}
func (resource *ImagingStudy) T_SeriesInstanceTitle(numSeries int, numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numSeries >= len(resource.Series) || numInstance >= len(resource.Series[numSeries].Instance) {
		return StringInput("Series["+strconv.Itoa(numSeries)+"]Instance["+strconv.Itoa(numInstance)+"].Title", nil, htmlAttrs)
	}
	return StringInput("Series["+strconv.Itoa(numSeries)+"]Instance["+strconv.Itoa(numInstance)+"].Title", resource.Series[numSeries].Instance[numInstance].Title, htmlAttrs)
}
