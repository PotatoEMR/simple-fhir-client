package r4b

//generated with command go run ./bultaoreune
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
func (r BodyStructure) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "BodyStructure/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "BodyStructure"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *BodyStructure) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *BodyStructure) T_Morphology(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("morphology", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("morphology", resource.Morphology, optionsValueSet, htmlAttrs)
}
func (resource *BodyStructure) T_Location(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("location", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("location", resource.Location, optionsValueSet, htmlAttrs)
}
func (resource *BodyStructure) T_LocationQualifier(numLocationQualifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocationQualifier >= len(resource.LocationQualifier) {
		return CodeableConceptSelect("locationQualifier["+strconv.Itoa(numLocationQualifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("locationQualifier["+strconv.Itoa(numLocationQualifier)+"]", &resource.LocationQualifier[numLocationQualifier], optionsValueSet, htmlAttrs)
}
func (resource *BodyStructure) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
