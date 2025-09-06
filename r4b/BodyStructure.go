package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/BodyStructure
type BodyStructure struct {
	Id                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []Resource        `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Active            *bool             `json:"active,omitempty"`
	Morphology        *CodeableConcept  `json:"morphology,omitempty"`
	Location          *CodeableConcept  `json:"location,omitempty"`
	LocationQualifier []CodeableConcept `json:"locationQualifier,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Image             []Attachment      `json:"image,omitempty"`
	Patient           Reference         `json:"patient"`
}

type OtherBodyStructure BodyStructure

// on convert struct to json, automatically add resourceType=BodyStructure
func (r BodyStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBodyStructure
		ResourceType string `json:"resourceType"`
	}{
		OtherBodyStructure: OtherBodyStructure(r),
		ResourceType:       "BodyStructure",
	})
}

func (resource *BodyStructure) T_Id() templ.Component {

	if resource == nil {
		return StringInput("BodyStructure.Id", nil)
	}
	return StringInput("BodyStructure.Id", resource.Id)
}
func (resource *BodyStructure) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("BodyStructure.ImplicitRules", nil)
	}
	return StringInput("BodyStructure.ImplicitRules", resource.ImplicitRules)
}
func (resource *BodyStructure) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("BodyStructure.Language", nil, optionsValueSet)
	}
	return CodeSelect("BodyStructure.Language", resource.Language, optionsValueSet)
}
func (resource *BodyStructure) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("BodyStructure.Active", nil)
	}
	return BoolInput("BodyStructure.Active", resource.Active)
}
func (resource *BodyStructure) T_Morphology(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("BodyStructure.Morphology", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BodyStructure.Morphology", resource.Morphology, optionsValueSet)
}
func (resource *BodyStructure) T_Location(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("BodyStructure.Location", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BodyStructure.Location", resource.Location, optionsValueSet)
}
func (resource *BodyStructure) T_LocationQualifier(numLocationQualifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.LocationQualifier) >= numLocationQualifier {
		return CodeableConceptSelect("BodyStructure.LocationQualifier["+strconv.Itoa(numLocationQualifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BodyStructure.LocationQualifier["+strconv.Itoa(numLocationQualifier)+"]", &resource.LocationQualifier[numLocationQualifier], optionsValueSet)
}
func (resource *BodyStructure) T_Description() templ.Component {

	if resource == nil {
		return StringInput("BodyStructure.Description", nil)
	}
	return StringInput("BodyStructure.Description", resource.Description)
}
