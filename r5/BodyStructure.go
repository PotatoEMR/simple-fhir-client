package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *BodyStructure) BodyStructureLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
