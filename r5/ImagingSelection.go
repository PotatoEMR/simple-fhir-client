//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/ImagingSelection
type ImagingSelection struct {
	Id                  *string                     `json:"id,omitempty"`
	Meta                *Meta                       `json:"meta,omitempty"`
	ImplicitRules       *string                     `json:"implicitRules,omitempty"`
	Language            *string                     `json:"language,omitempty"`
	Text                *Narrative                  `json:"text,omitempty"`
	Contained           []Resource                  `json:"contained,omitempty"`
	Extension           []Extension                 `json:"extension,omitempty"`
	ModifierExtension   []Extension                 `json:"modifierExtension,omitempty"`
	Identifier          []Identifier                `json:"identifier,omitempty"`
	Status              string                      `json:"status"`
	Subject             *Reference                  `json:"subject,omitempty"`
	Issued              *string                     `json:"issued,omitempty"`
	Performer           []ImagingSelectionPerformer `json:"performer,omitempty"`
	BasedOn             []Reference                 `json:"basedOn,omitempty"`
	Category            []CodeableConcept           `json:"category,omitempty"`
	Code                CodeableConcept             `json:"code"`
	StudyUid            *string                     `json:"studyUid,omitempty"`
	DerivedFrom         []Reference                 `json:"derivedFrom,omitempty"`
	Endpoint            []Reference                 `json:"endpoint,omitempty"`
	SeriesUid           *string                     `json:"seriesUid,omitempty"`
	SeriesNumber        *int                        `json:"seriesNumber,omitempty"`
	FrameOfReferenceUid *string                     `json:"frameOfReferenceUid,omitempty"`
	BodySite            *CodeableReference          `json:"bodySite,omitempty"`
	Focus               []Reference                 `json:"focus,omitempty"`
	Instance            []ImagingSelectionInstance  `json:"instance,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImagingSelection
type ImagingSelectionPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             *Reference       `json:"actor,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImagingSelection
type ImagingSelectionInstance struct {
	Id                *string                                 `json:"id,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Uid               string                                  `json:"uid"`
	Number            *int                                    `json:"number,omitempty"`
	SopClass          *Coding                                 `json:"sopClass,omitempty"`
	Subset            []string                                `json:"subset,omitempty"`
	ImageRegion2D     []ImagingSelectionInstanceImageRegion2D `json:"imageRegion2D,omitempty"`
	ImageRegion3D     []ImagingSelectionInstanceImageRegion3D `json:"imageRegion3D,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImagingSelection
type ImagingSelectionInstanceImageRegion2D struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	RegionType        string      `json:"regionType"`
	Coordinate        []float64   `json:"coordinate"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImagingSelection
type ImagingSelectionInstanceImageRegion3D struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	RegionType        string      `json:"regionType"`
	Coordinate        []float64   `json:"coordinate"`
}

type OtherImagingSelection ImagingSelection

// on convert struct to json, automatically add resourceType=ImagingSelection
func (r ImagingSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImagingSelection
		ResourceType string `json:"resourceType"`
	}{
		OtherImagingSelection: OtherImagingSelection(r),
		ResourceType:          "ImagingSelection",
	})
}
