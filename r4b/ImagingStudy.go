package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/ImagingStudy
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

// http://hl7.org/fhir/r4b/StructureDefinition/ImagingStudy
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

// http://hl7.org/fhir/r4b/StructureDefinition/ImagingStudy
type ImagingStudySeriesPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImagingStudy
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

func (resource *ImagingStudy) ImagingStudyLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ImagingStudy) ImagingStudyStatus() templ.Component {
	optionsValueSet := VSImagingstudy_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ImagingStudy) ImagingStudyModality(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("modality", nil, optionsValueSet)
	}
	return CodingSelect("modality", &resource.Modality[0], optionsValueSet)
}
func (resource *ImagingStudy) ImagingStudyProcedureCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("procedureCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("procedureCode", &resource.ProcedureCode[0], optionsValueSet)
}
func (resource *ImagingStudy) ImagingStudyReasonCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *ImagingStudy) ImagingStudySeriesModality(numSeries int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Series) >= numSeries {
		return CodingSelect("modality", nil, optionsValueSet)
	}
	return CodingSelect("modality", &resource.Series[numSeries].Modality, optionsValueSet)
}
func (resource *ImagingStudy) ImagingStudySeriesBodySite(numSeries int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Series) >= numSeries {
		return CodingSelect("bodySite", nil, optionsValueSet)
	}
	return CodingSelect("bodySite", resource.Series[numSeries].BodySite, optionsValueSet)
}
func (resource *ImagingStudy) ImagingStudySeriesLaterality(numSeries int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Series) >= numSeries {
		return CodingSelect("laterality", nil, optionsValueSet)
	}
	return CodingSelect("laterality", resource.Series[numSeries].Laterality, optionsValueSet)
}
func (resource *ImagingStudy) ImagingStudySeriesPerformerFunction(numSeries int, numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Series[numSeries].Performer) >= numPerformer {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Series[numSeries].Performer[numPerformer].Function, optionsValueSet)
}
func (resource *ImagingStudy) ImagingStudySeriesInstanceSopClass(numSeries int, numInstance int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Series[numSeries].Instance) >= numInstance {
		return CodingSelect("sopClass", nil, optionsValueSet)
	}
	return CodingSelect("sopClass", &resource.Series[numSeries].Instance[numInstance].SopClass, optionsValueSet)
}
