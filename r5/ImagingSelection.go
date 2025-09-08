package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (r ImagingSelection) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ImagingSelection/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ImagingSelection"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ImagingSelection) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSImagingselection_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ImagingSelection) T_Issued(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Issued", nil, htmlAttrs)
	}
	return StringInput("Issued", resource.Issued, htmlAttrs)
}
func (resource *ImagingSelection) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ImagingSelection) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ImagingSelection) T_StudyUid(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("StudyUid", nil, htmlAttrs)
	}
	return StringInput("StudyUid", resource.StudyUid, htmlAttrs)
}
func (resource *ImagingSelection) T_SeriesUid(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SeriesUid", nil, htmlAttrs)
	}
	return StringInput("SeriesUid", resource.SeriesUid, htmlAttrs)
}
func (resource *ImagingSelection) T_SeriesNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("SeriesNumber", nil, htmlAttrs)
	}
	return IntInput("SeriesNumber", resource.SeriesNumber, htmlAttrs)
}
func (resource *ImagingSelection) T_FrameOfReferenceUid(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("FrameOfReferenceUid", nil, htmlAttrs)
	}
	return StringInput("FrameOfReferenceUid", resource.FrameOfReferenceUid, htmlAttrs)
}
func (resource *ImagingSelection) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("Performer["+strconv.Itoa(numPerformer)+"]Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Performer["+strconv.Itoa(numPerformer)+"]Function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *ImagingSelection) T_InstanceUid(numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("Instance["+strconv.Itoa(numInstance)+"]Uid", nil, htmlAttrs)
	}
	return StringInput("Instance["+strconv.Itoa(numInstance)+"]Uid", &resource.Instance[numInstance].Uid, htmlAttrs)
}
func (resource *ImagingSelection) T_InstanceNumber(numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return IntInput("Instance["+strconv.Itoa(numInstance)+"]Number", nil, htmlAttrs)
	}
	return IntInput("Instance["+strconv.Itoa(numInstance)+"]Number", resource.Instance[numInstance].Number, htmlAttrs)
}
func (resource *ImagingSelection) T_InstanceSopClass(numInstance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return CodingSelect("Instance["+strconv.Itoa(numInstance)+"]SopClass", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Instance["+strconv.Itoa(numInstance)+"]SopClass", resource.Instance[numInstance].SopClass, optionsValueSet, htmlAttrs)
}
func (resource *ImagingSelection) T_InstanceSubset(numInstance int, numSubset int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numSubset >= len(resource.Instance[numInstance].Subset) {
		return StringInput("Instance["+strconv.Itoa(numInstance)+"]Subset["+strconv.Itoa(numSubset)+"]", nil, htmlAttrs)
	}
	return StringInput("Instance["+strconv.Itoa(numInstance)+"]Subset["+strconv.Itoa(numSubset)+"]", &resource.Instance[numInstance].Subset[numSubset], htmlAttrs)
}
func (resource *ImagingSelection) T_InstanceImageRegion2DRegionType(numInstance int, numImageRegion2D int, htmlAttrs string) templ.Component {
	optionsValueSet := VSImagingselection_2dgraphictype

	if resource == nil || numInstance >= len(resource.Instance) || numImageRegion2D >= len(resource.Instance[numInstance].ImageRegion2D) {
		return CodeSelect("Instance["+strconv.Itoa(numInstance)+"]ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].RegionType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Instance["+strconv.Itoa(numInstance)+"]ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].RegionType", &resource.Instance[numInstance].ImageRegion2D[numImageRegion2D].RegionType, optionsValueSet, htmlAttrs)
}
func (resource *ImagingSelection) T_InstanceImageRegion2DCoordinate(numInstance int, numImageRegion2D int, numCoordinate int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numImageRegion2D >= len(resource.Instance[numInstance].ImageRegion2D) || numCoordinate >= len(resource.Instance[numInstance].ImageRegion2D[numImageRegion2D].Coordinate) {
		return Float64Input("Instance["+strconv.Itoa(numInstance)+"]ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].Coordinate["+strconv.Itoa(numCoordinate)+"]", nil, htmlAttrs)
	}
	return Float64Input("Instance["+strconv.Itoa(numInstance)+"]ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].Coordinate["+strconv.Itoa(numCoordinate)+"]", &resource.Instance[numInstance].ImageRegion2D[numImageRegion2D].Coordinate[numCoordinate], htmlAttrs)
}
func (resource *ImagingSelection) T_InstanceImageRegion3DRegionType(numInstance int, numImageRegion3D int, htmlAttrs string) templ.Component {
	optionsValueSet := VSImagingselection_3dgraphictype

	if resource == nil || numInstance >= len(resource.Instance) || numImageRegion3D >= len(resource.Instance[numInstance].ImageRegion3D) {
		return CodeSelect("Instance["+strconv.Itoa(numInstance)+"]ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].RegionType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Instance["+strconv.Itoa(numInstance)+"]ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].RegionType", &resource.Instance[numInstance].ImageRegion3D[numImageRegion3D].RegionType, optionsValueSet, htmlAttrs)
}
func (resource *ImagingSelection) T_InstanceImageRegion3DCoordinate(numInstance int, numImageRegion3D int, numCoordinate int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numImageRegion3D >= len(resource.Instance[numInstance].ImageRegion3D) || numCoordinate >= len(resource.Instance[numInstance].ImageRegion3D[numImageRegion3D].Coordinate) {
		return Float64Input("Instance["+strconv.Itoa(numInstance)+"]ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].Coordinate["+strconv.Itoa(numCoordinate)+"]", nil, htmlAttrs)
	}
	return Float64Input("Instance["+strconv.Itoa(numInstance)+"]ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].Coordinate["+strconv.Itoa(numCoordinate)+"]", &resource.Instance[numInstance].ImageRegion3D[numImageRegion3D].Coordinate[numCoordinate], htmlAttrs)
}
