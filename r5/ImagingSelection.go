package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *ImagingSelection) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ImagingSelection.Id", nil)
	}
	return StringInput("ImagingSelection.Id", resource.Id)
}
func (resource *ImagingSelection) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ImagingSelection.ImplicitRules", nil)
	}
	return StringInput("ImagingSelection.ImplicitRules", resource.ImplicitRules)
}
func (resource *ImagingSelection) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ImagingSelection.Language", nil, optionsValueSet)
	}
	return CodeSelect("ImagingSelection.Language", resource.Language, optionsValueSet)
}
func (resource *ImagingSelection) T_Status() templ.Component {
	optionsValueSet := VSImagingselection_status

	if resource == nil {
		return CodeSelect("ImagingSelection.Status", nil, optionsValueSet)
	}
	return CodeSelect("ImagingSelection.Status", &resource.Status, optionsValueSet)
}
func (resource *ImagingSelection) T_Issued() templ.Component {

	if resource == nil {
		return StringInput("ImagingSelection.Issued", nil)
	}
	return StringInput("ImagingSelection.Issued", resource.Issued)
}
func (resource *ImagingSelection) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("ImagingSelection.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImagingSelection.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *ImagingSelection) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ImagingSelection.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImagingSelection.Code", &resource.Code, optionsValueSet)
}
func (resource *ImagingSelection) T_StudyUid() templ.Component {

	if resource == nil {
		return StringInput("ImagingSelection.StudyUid", nil)
	}
	return StringInput("ImagingSelection.StudyUid", resource.StudyUid)
}
func (resource *ImagingSelection) T_SeriesUid() templ.Component {

	if resource == nil {
		return StringInput("ImagingSelection.SeriesUid", nil)
	}
	return StringInput("ImagingSelection.SeriesUid", resource.SeriesUid)
}
func (resource *ImagingSelection) T_SeriesNumber() templ.Component {

	if resource == nil {
		return IntInput("ImagingSelection.SeriesNumber", nil)
	}
	return IntInput("ImagingSelection.SeriesNumber", resource.SeriesNumber)
}
func (resource *ImagingSelection) T_FrameOfReferenceUid() templ.Component {

	if resource == nil {
		return StringInput("ImagingSelection.FrameOfReferenceUid", nil)
	}
	return StringInput("ImagingSelection.FrameOfReferenceUid", resource.FrameOfReferenceUid)
}
func (resource *ImagingSelection) T_PerformerId(numPerformer int) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return StringInput("ImagingSelection.Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("ImagingSelection.Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Performer[numPerformer].Id)
}
func (resource *ImagingSelection) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("ImagingSelection.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImagingSelection.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *ImagingSelection) T_InstanceId(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].Id", nil)
	}
	return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].Id", resource.Instance[numInstance].Id)
}
func (resource *ImagingSelection) T_InstanceUid(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].Uid", nil)
	}
	return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].Uid", &resource.Instance[numInstance].Uid)
}
func (resource *ImagingSelection) T_InstanceNumber(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return IntInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].Number", nil)
	}
	return IntInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].Number", resource.Instance[numInstance].Number)
}
func (resource *ImagingSelection) T_InstanceSopClass(numInstance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return CodingSelect("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].SopClass", nil, optionsValueSet)
	}
	return CodingSelect("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].SopClass", resource.Instance[numInstance].SopClass, optionsValueSet)
}
func (resource *ImagingSelection) T_InstanceSubset(numInstance int, numSubset int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].Subset) >= numSubset {
		return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].Subset["+strconv.Itoa(numSubset)+"]", nil)
	}
	return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].Subset["+strconv.Itoa(numSubset)+"]", &resource.Instance[numInstance].Subset[numSubset])
}
func (resource *ImagingSelection) T_InstanceImageRegion2DId(numInstance int, numImageRegion2D int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].ImageRegion2D) >= numImageRegion2D {
		return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].Id", nil)
	}
	return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].Id", resource.Instance[numInstance].ImageRegion2D[numImageRegion2D].Id)
}
func (resource *ImagingSelection) T_InstanceImageRegion2DRegionType(numInstance int, numImageRegion2D int) templ.Component {
	optionsValueSet := VSImagingselection_2dgraphictype

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].ImageRegion2D) >= numImageRegion2D {
		return CodeSelect("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].RegionType", nil, optionsValueSet)
	}
	return CodeSelect("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].RegionType", &resource.Instance[numInstance].ImageRegion2D[numImageRegion2D].RegionType, optionsValueSet)
}
func (resource *ImagingSelection) T_InstanceImageRegion2DCoordinate(numInstance int, numImageRegion2D int, numCoordinate int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].ImageRegion2D) >= numImageRegion2D || len(resource.Instance[numInstance].ImageRegion2D[numImageRegion2D].Coordinate) >= numCoordinate {
		return Float64Input("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].Coordinate["+strconv.Itoa(numCoordinate)+"]", nil)
	}
	return Float64Input("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion2D["+strconv.Itoa(numImageRegion2D)+"].Coordinate["+strconv.Itoa(numCoordinate)+"]", &resource.Instance[numInstance].ImageRegion2D[numImageRegion2D].Coordinate[numCoordinate])
}
func (resource *ImagingSelection) T_InstanceImageRegion3DId(numInstance int, numImageRegion3D int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].ImageRegion3D) >= numImageRegion3D {
		return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].Id", nil)
	}
	return StringInput("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].Id", resource.Instance[numInstance].ImageRegion3D[numImageRegion3D].Id)
}
func (resource *ImagingSelection) T_InstanceImageRegion3DRegionType(numInstance int, numImageRegion3D int) templ.Component {
	optionsValueSet := VSImagingselection_3dgraphictype

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].ImageRegion3D) >= numImageRegion3D {
		return CodeSelect("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].RegionType", nil, optionsValueSet)
	}
	return CodeSelect("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].RegionType", &resource.Instance[numInstance].ImageRegion3D[numImageRegion3D].RegionType, optionsValueSet)
}
func (resource *ImagingSelection) T_InstanceImageRegion3DCoordinate(numInstance int, numImageRegion3D int, numCoordinate int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].ImageRegion3D) >= numImageRegion3D || len(resource.Instance[numInstance].ImageRegion3D[numImageRegion3D].Coordinate) >= numCoordinate {
		return Float64Input("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].Coordinate["+strconv.Itoa(numCoordinate)+"]", nil)
	}
	return Float64Input("ImagingSelection.Instance["+strconv.Itoa(numInstance)+"].ImageRegion3D["+strconv.Itoa(numImageRegion3D)+"].Coordinate["+strconv.Itoa(numCoordinate)+"]", &resource.Instance[numInstance].ImageRegion3D[numImageRegion3D].Coordinate[numCoordinate])
}
