package r5

//generated with command go run ./bultaoreune -nodownload
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

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *BodyStructure) BodyStructureMorphology(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("morphology", nil, optionsValueSet)
	}
	return CodeableConceptSelect("morphology", resource.Morphology, optionsValueSet)
}
func (resource *BodyStructure) BodyStructureIncludedStructureStructure(numIncludedStructure int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.IncludedStructure) >= numIncludedStructure {
		return CodeableConceptSelect("structure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("structure", &resource.IncludedStructure[numIncludedStructure].Structure, optionsValueSet)
}
func (resource *BodyStructure) BodyStructureIncludedStructureLaterality(numIncludedStructure int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.IncludedStructure) >= numIncludedStructure {
		return CodeableConceptSelect("laterality", nil, optionsValueSet)
	}
	return CodeableConceptSelect("laterality", resource.IncludedStructure[numIncludedStructure].Laterality, optionsValueSet)
}
func (resource *BodyStructure) BodyStructureIncludedStructureQualifier(numIncludedStructure int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.IncludedStructure) >= numIncludedStructure {
		return CodeableConceptSelect("qualifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("qualifier", &resource.IncludedStructure[numIncludedStructure].Qualifier[0], optionsValueSet)
}
func (resource *BodyStructure) BodyStructureIncludedStructureBodyLandmarkOrientationLandmarkDescription(numIncludedStructure int, numBodyLandmarkOrientation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) >= numBodyLandmarkOrientation {
		return CodeableConceptSelect("landmarkDescription", nil, optionsValueSet)
	}
	return CodeableConceptSelect("landmarkDescription", &resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].LandmarkDescription[0], optionsValueSet)
}
func (resource *BodyStructure) BodyStructureIncludedStructureBodyLandmarkOrientationClockFacePosition(numIncludedStructure int, numBodyLandmarkOrientation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) >= numBodyLandmarkOrientation {
		return CodeableConceptSelect("clockFacePosition", nil, optionsValueSet)
	}
	return CodeableConceptSelect("clockFacePosition", &resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].ClockFacePosition[0], optionsValueSet)
}
func (resource *BodyStructure) BodyStructureIncludedStructureBodyLandmarkOrientationSurfaceOrientation(numIncludedStructure int, numBodyLandmarkOrientation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) >= numBodyLandmarkOrientation {
		return CodeableConceptSelect("surfaceOrientation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("surfaceOrientation", &resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].SurfaceOrientation[0], optionsValueSet)
}
