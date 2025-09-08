package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/BodyStructure
type BodyStructure struct {
	Id                *string                          `json:"id,omitempty"`
	Meta              *Meta                            `json:"meta,omitempty"`
	ImplicitRules     *string                          `json:"implicitRules,omitempty"`
	Language          *string                          `json:"language,omitempty"`
	Text              *Narrative                       `json:"text,omitempty"`
	Contained         []Resource                       `json:"contained,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                     `json:"identifier,omitempty"`
	Active            *bool                            `json:"active,omitempty"`
	Morphology        *CodeableConcept                 `json:"morphology,omitempty"`
	IncludedStructure []BodyStructureIncludedStructure `json:"includedStructure"`
	Description       *string                          `json:"description,omitempty"`
	Image             []Attachment                     `json:"image,omitempty"`
	Patient           Reference                        `json:"patient"`
}

// http://hl7.org/fhir/r5/StructureDefinition/BodyStructure
type BodyStructureIncludedStructure struct {
	Id                      *string                                                 `json:"id,omitempty"`
	Extension               []Extension                                             `json:"extension,omitempty"`
	ModifierExtension       []Extension                                             `json:"modifierExtension,omitempty"`
	Structure               CodeableConcept                                         `json:"structure"`
	Laterality              *CodeableConcept                                        `json:"laterality,omitempty"`
	BodyLandmarkOrientation []BodyStructureIncludedStructureBodyLandmarkOrientation `json:"bodyLandmarkOrientation,omitempty"`
	SpatialReference        []Reference                                             `json:"spatialReference,omitempty"`
	Qualifier               []CodeableConcept                                       `json:"qualifier,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/BodyStructure
type BodyStructureIncludedStructureBodyLandmarkOrientation struct {
	Id                   *string                                                                     `json:"id,omitempty"`
	Extension            []Extension                                                                 `json:"extension,omitempty"`
	ModifierExtension    []Extension                                                                 `json:"modifierExtension,omitempty"`
	LandmarkDescription  []CodeableConcept                                                           `json:"landmarkDescription,omitempty"`
	ClockFacePosition    []CodeableConcept                                                           `json:"clockFacePosition,omitempty"`
	DistanceFromLandmark []BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark `json:"distanceFromLandmark,omitempty"`
	SurfaceOrientation   []CodeableConcept                                                           `json:"surfaceOrientation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/BodyStructure
type BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Device            []CodeableReference `json:"device,omitempty"`
	Value             []Quantity          `json:"value,omitempty"`
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
func (resource *BodyStructure) T_Active(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Active", nil, htmlAttrs)
	}
	return BoolInput("Active", resource.Active, htmlAttrs)
}
func (resource *BodyStructure) T_Morphology(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Morphology", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Morphology", resource.Morphology, optionsValueSet, htmlAttrs)
}
func (resource *BodyStructure) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *BodyStructure) T_IncludedStructureStructure(numIncludedStructure int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIncludedStructure >= len(resource.IncludedStructure) {
		return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]Structure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]Structure", &resource.IncludedStructure[numIncludedStructure].Structure, optionsValueSet, htmlAttrs)
}
func (resource *BodyStructure) T_IncludedStructureLaterality(numIncludedStructure int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIncludedStructure >= len(resource.IncludedStructure) {
		return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]Laterality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]Laterality", resource.IncludedStructure[numIncludedStructure].Laterality, optionsValueSet, htmlAttrs)
}
func (resource *BodyStructure) T_IncludedStructureQualifier(numIncludedStructure int, numQualifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIncludedStructure >= len(resource.IncludedStructure) || numQualifier >= len(resource.IncludedStructure[numIncludedStructure].Qualifier) {
		return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]Qualifier["+strconv.Itoa(numQualifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]Qualifier["+strconv.Itoa(numQualifier)+"]", &resource.IncludedStructure[numIncludedStructure].Qualifier[numQualifier], optionsValueSet, htmlAttrs)
}
func (resource *BodyStructure) T_IncludedStructureBodyLandmarkOrientationLandmarkDescription(numIncludedStructure int, numBodyLandmarkOrientation int, numLandmarkDescription int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIncludedStructure >= len(resource.IncludedStructure) || numBodyLandmarkOrientation >= len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) || numLandmarkDescription >= len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].LandmarkDescription) {
		return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].LandmarkDescription["+strconv.Itoa(numLandmarkDescription)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].LandmarkDescription["+strconv.Itoa(numLandmarkDescription)+"]", &resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].LandmarkDescription[numLandmarkDescription], optionsValueSet, htmlAttrs)
}
func (resource *BodyStructure) T_IncludedStructureBodyLandmarkOrientationClockFacePosition(numIncludedStructure int, numBodyLandmarkOrientation int, numClockFacePosition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIncludedStructure >= len(resource.IncludedStructure) || numBodyLandmarkOrientation >= len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) || numClockFacePosition >= len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].ClockFacePosition) {
		return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].ClockFacePosition["+strconv.Itoa(numClockFacePosition)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].ClockFacePosition["+strconv.Itoa(numClockFacePosition)+"]", &resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].ClockFacePosition[numClockFacePosition], optionsValueSet, htmlAttrs)
}
func (resource *BodyStructure) T_IncludedStructureBodyLandmarkOrientationSurfaceOrientation(numIncludedStructure int, numBodyLandmarkOrientation int, numSurfaceOrientation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIncludedStructure >= len(resource.IncludedStructure) || numBodyLandmarkOrientation >= len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) || numSurfaceOrientation >= len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].SurfaceOrientation) {
		return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].SurfaceOrientation["+strconv.Itoa(numSurfaceOrientation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("IncludedStructure["+strconv.Itoa(numIncludedStructure)+"]BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].SurfaceOrientation["+strconv.Itoa(numSurfaceOrientation)+"]", &resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].SurfaceOrientation[numSurfaceOrientation], optionsValueSet, htmlAttrs)
}
