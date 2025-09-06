package r5

//generated with command go run ./bultaoreune -nodownload
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
func (resource *BodyStructure) T_Description() templ.Component {

	if resource == nil {
		return StringInput("BodyStructure.Description", nil)
	}
	return StringInput("BodyStructure.Description", resource.Description)
}
func (resource *BodyStructure) T_IncludedStructureId(numIncludedStructure int) templ.Component {

	if resource == nil || len(resource.IncludedStructure) >= numIncludedStructure {
		return StringInput("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].Id", nil)
	}
	return StringInput("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].Id", resource.IncludedStructure[numIncludedStructure].Id)
}
func (resource *BodyStructure) T_IncludedStructureStructure(numIncludedStructure int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IncludedStructure) >= numIncludedStructure {
		return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].Structure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].Structure", &resource.IncludedStructure[numIncludedStructure].Structure, optionsValueSet)
}
func (resource *BodyStructure) T_IncludedStructureLaterality(numIncludedStructure int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IncludedStructure) >= numIncludedStructure {
		return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].Laterality", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].Laterality", resource.IncludedStructure[numIncludedStructure].Laterality, optionsValueSet)
}
func (resource *BodyStructure) T_IncludedStructureQualifier(numIncludedStructure int, numQualifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IncludedStructure) >= numIncludedStructure || len(resource.IncludedStructure[numIncludedStructure].Qualifier) >= numQualifier {
		return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].Qualifier["+strconv.Itoa(numQualifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].Qualifier["+strconv.Itoa(numQualifier)+"]", &resource.IncludedStructure[numIncludedStructure].Qualifier[numQualifier], optionsValueSet)
}
func (resource *BodyStructure) T_IncludedStructureBodyLandmarkOrientationId(numIncludedStructure int, numBodyLandmarkOrientation int) templ.Component {

	if resource == nil || len(resource.IncludedStructure) >= numIncludedStructure || len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) >= numBodyLandmarkOrientation {
		return StringInput("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].Id", nil)
	}
	return StringInput("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].Id", resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].Id)
}
func (resource *BodyStructure) T_IncludedStructureBodyLandmarkOrientationLandmarkDescription(numIncludedStructure int, numBodyLandmarkOrientation int, numLandmarkDescription int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IncludedStructure) >= numIncludedStructure || len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) >= numBodyLandmarkOrientation || len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].LandmarkDescription) >= numLandmarkDescription {
		return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].LandmarkDescription["+strconv.Itoa(numLandmarkDescription)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].LandmarkDescription["+strconv.Itoa(numLandmarkDescription)+"]", &resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].LandmarkDescription[numLandmarkDescription], optionsValueSet)
}
func (resource *BodyStructure) T_IncludedStructureBodyLandmarkOrientationClockFacePosition(numIncludedStructure int, numBodyLandmarkOrientation int, numClockFacePosition int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IncludedStructure) >= numIncludedStructure || len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) >= numBodyLandmarkOrientation || len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].ClockFacePosition) >= numClockFacePosition {
		return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].ClockFacePosition["+strconv.Itoa(numClockFacePosition)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].ClockFacePosition["+strconv.Itoa(numClockFacePosition)+"]", &resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].ClockFacePosition[numClockFacePosition], optionsValueSet)
}
func (resource *BodyStructure) T_IncludedStructureBodyLandmarkOrientationSurfaceOrientation(numIncludedStructure int, numBodyLandmarkOrientation int, numSurfaceOrientation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IncludedStructure) >= numIncludedStructure || len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) >= numBodyLandmarkOrientation || len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].SurfaceOrientation) >= numSurfaceOrientation {
		return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].SurfaceOrientation["+strconv.Itoa(numSurfaceOrientation)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].SurfaceOrientation["+strconv.Itoa(numSurfaceOrientation)+"]", &resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].SurfaceOrientation[numSurfaceOrientation], optionsValueSet)
}
func (resource *BodyStructure) T_IncludedStructureBodyLandmarkOrientationDistanceFromLandmarkId(numIncludedStructure int, numBodyLandmarkOrientation int, numDistanceFromLandmark int) templ.Component {

	if resource == nil || len(resource.IncludedStructure) >= numIncludedStructure || len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation) >= numBodyLandmarkOrientation || len(resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].DistanceFromLandmark) >= numDistanceFromLandmark {
		return StringInput("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].DistanceFromLandmark["+strconv.Itoa(numDistanceFromLandmark)+"].Id", nil)
	}
	return StringInput("BodyStructure.IncludedStructure["+strconv.Itoa(numIncludedStructure)+"].BodyLandmarkOrientation["+strconv.Itoa(numBodyLandmarkOrientation)+"].DistanceFromLandmark["+strconv.Itoa(numDistanceFromLandmark)+"].Id", resource.IncludedStructure[numIncludedStructure].BodyLandmarkOrientation[numBodyLandmarkOrientation].DistanceFromLandmark[numDistanceFromLandmark].Id)
}
