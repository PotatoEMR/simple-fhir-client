package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ObservationDefinition
type ObservationDefinition struct {
	Id                     *string                               `json:"id,omitempty"`
	Meta                   *Meta                                 `json:"meta,omitempty"`
	ImplicitRules          *string                               `json:"implicitRules,omitempty"`
	Language               *string                               `json:"language,omitempty"`
	Text                   *Narrative                            `json:"text,omitempty"`
	Contained              []Resource                            `json:"contained,omitempty"`
	Extension              []Extension                           `json:"extension,omitempty"`
	ModifierExtension      []Extension                           `json:"modifierExtension,omitempty"`
	Url                    *string                               `json:"url,omitempty"`
	Identifier             *Identifier                           `json:"identifier,omitempty"`
	Version                *string                               `json:"version,omitempty"`
	VersionAlgorithmString *string                               `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                               `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                               `json:"name,omitempty"`
	Title                  *string                               `json:"title,omitempty"`
	Status                 string                                `json:"status"`
	Experimental           *bool                                 `json:"experimental,omitempty"`
	Date                   *string                               `json:"date,omitempty"`
	Publisher              *string                               `json:"publisher,omitempty"`
	Contact                []ContactDetail                       `json:"contact,omitempty"`
	Description            *string                               `json:"description,omitempty"`
	UseContext             []UsageContext                        `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                     `json:"jurisdiction,omitempty"`
	Purpose                *string                               `json:"purpose,omitempty"`
	Copyright              *string                               `json:"copyright,omitempty"`
	CopyrightLabel         *string                               `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                               `json:"approvalDate,omitempty"`
	LastReviewDate         *string                               `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                               `json:"effectivePeriod,omitempty"`
	DerivedFromCanonical   []string                              `json:"derivedFromCanonical,omitempty"`
	DerivedFromUri         []string                              `json:"derivedFromUri,omitempty"`
	Subject                []CodeableConcept                     `json:"subject,omitempty"`
	PerformerType          *CodeableConcept                      `json:"performerType,omitempty"`
	Category               []CodeableConcept                     `json:"category,omitempty"`
	Code                   CodeableConcept                       `json:"code"`
	PermittedDataType      []string                              `json:"permittedDataType,omitempty"`
	MultipleResultsAllowed *bool                                 `json:"multipleResultsAllowed,omitempty"`
	BodySite               *CodeableConcept                      `json:"bodySite,omitempty"`
	Method                 *CodeableConcept                      `json:"method,omitempty"`
	Specimen               []Reference                           `json:"specimen,omitempty"`
	Device                 []Reference                           `json:"device,omitempty"`
	PreferredReportName    *string                               `json:"preferredReportName,omitempty"`
	PermittedUnit          []Coding                              `json:"permittedUnit,omitempty"`
	QualifiedValue         []ObservationDefinitionQualifiedValue `json:"qualifiedValue,omitempty"`
	HasMember              []Reference                           `json:"hasMember,omitempty"`
	Component              []ObservationDefinitionComponent      `json:"component,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ObservationDefinition
type ObservationDefinitionQualifiedValue struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Context               *CodeableConcept  `json:"context,omitempty"`
	AppliesTo             []CodeableConcept `json:"appliesTo,omitempty"`
	Gender                *string           `json:"gender,omitempty"`
	Age                   *Range            `json:"age,omitempty"`
	GestationalAge        *Range            `json:"gestationalAge,omitempty"`
	Condition             *string           `json:"condition,omitempty"`
	RangeCategory         *string           `json:"rangeCategory,omitempty"`
	Range                 *Range            `json:"range,omitempty"`
	ValidCodedValueSet    *string           `json:"validCodedValueSet,omitempty"`
	NormalCodedValueSet   *string           `json:"normalCodedValueSet,omitempty"`
	AbnormalCodedValueSet *string           `json:"abnormalCodedValueSet,omitempty"`
	CriticalCodedValueSet *string           `json:"criticalCodedValueSet,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ObservationDefinition
type ObservationDefinitionComponent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	PermittedDataType []string        `json:"permittedDataType,omitempty"`
	PermittedUnit     []Coding        `json:"permittedUnit,omitempty"`
}

type OtherObservationDefinition ObservationDefinition

// on convert struct to json, automatically add resourceType=ObservationDefinition
func (r ObservationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservationDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherObservationDefinition: OtherObservationDefinition(r),
		ResourceType:               "ObservationDefinition",
	})
}

func (resource *ObservationDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Id", nil)
	}
	return StringInput("ObservationDefinition.Id", resource.Id)
}
func (resource *ObservationDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.ImplicitRules", nil)
	}
	return StringInput("ObservationDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ObservationDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ObservationDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ObservationDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Url", nil)
	}
	return StringInput("ObservationDefinition.Url", resource.Url)
}
func (resource *ObservationDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Version", nil)
	}
	return StringInput("ObservationDefinition.Version", resource.Version)
}
func (resource *ObservationDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Name", nil)
	}
	return StringInput("ObservationDefinition.Name", resource.Name)
}
func (resource *ObservationDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Title", nil)
	}
	return StringInput("ObservationDefinition.Title", resource.Title)
}
func (resource *ObservationDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ObservationDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *ObservationDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ObservationDefinition.Experimental", nil)
	}
	return BoolInput("ObservationDefinition.Experimental", resource.Experimental)
}
func (resource *ObservationDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Date", nil)
	}
	return StringInput("ObservationDefinition.Date", resource.Date)
}
func (resource *ObservationDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Publisher", nil)
	}
	return StringInput("ObservationDefinition.Publisher", resource.Publisher)
}
func (resource *ObservationDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Description", nil)
	}
	return StringInput("ObservationDefinition.Description", resource.Description)
}
func (resource *ObservationDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ObservationDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ObservationDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Purpose", nil)
	}
	return StringInput("ObservationDefinition.Purpose", resource.Purpose)
}
func (resource *ObservationDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Copyright", nil)
	}
	return StringInput("ObservationDefinition.Copyright", resource.Copyright)
}
func (resource *ObservationDefinition) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.CopyrightLabel", nil)
	}
	return StringInput("ObservationDefinition.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *ObservationDefinition) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.ApprovalDate", nil)
	}
	return StringInput("ObservationDefinition.ApprovalDate", resource.ApprovalDate)
}
func (resource *ObservationDefinition) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.LastReviewDate", nil)
	}
	return StringInput("ObservationDefinition.LastReviewDate", resource.LastReviewDate)
}
func (resource *ObservationDefinition) T_DerivedFromCanonical(numDerivedFromCanonical int) templ.Component {

	if resource == nil || len(resource.DerivedFromCanonical) >= numDerivedFromCanonical {
		return StringInput("ObservationDefinition.DerivedFromCanonical["+strconv.Itoa(numDerivedFromCanonical)+"]", nil)
	}
	return StringInput("ObservationDefinition.DerivedFromCanonical["+strconv.Itoa(numDerivedFromCanonical)+"]", &resource.DerivedFromCanonical[numDerivedFromCanonical])
}
func (resource *ObservationDefinition) T_DerivedFromUri(numDerivedFromUri int) templ.Component {

	if resource == nil || len(resource.DerivedFromUri) >= numDerivedFromUri {
		return StringInput("ObservationDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", nil)
	}
	return StringInput("ObservationDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", &resource.DerivedFromUri[numDerivedFromUri])
}
func (resource *ObservationDefinition) T_Subject(numSubject int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Subject) >= numSubject {
		return CodeableConceptSelect("ObservationDefinition.Subject["+strconv.Itoa(numSubject)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.Subject["+strconv.Itoa(numSubject)+"]", &resource.Subject[numSubject], optionsValueSet)
}
func (resource *ObservationDefinition) T_PerformerType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.PerformerType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.PerformerType", resource.PerformerType, optionsValueSet)
}
func (resource *ObservationDefinition) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("ObservationDefinition.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *ObservationDefinition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.Code", &resource.Code, optionsValueSet)
}
func (resource *ObservationDefinition) T_PermittedDataType(numPermittedDataType int) templ.Component {
	optionsValueSet := VSPermitted_data_type

	if resource == nil || len(resource.PermittedDataType) >= numPermittedDataType {
		return CodeSelect("ObservationDefinition.PermittedDataType["+strconv.Itoa(numPermittedDataType)+"]", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.PermittedDataType["+strconv.Itoa(numPermittedDataType)+"]", &resource.PermittedDataType[numPermittedDataType], optionsValueSet)
}
func (resource *ObservationDefinition) T_MultipleResultsAllowed() templ.Component {

	if resource == nil {
		return BoolInput("ObservationDefinition.MultipleResultsAllowed", nil)
	}
	return BoolInput("ObservationDefinition.MultipleResultsAllowed", resource.MultipleResultsAllowed)
}
func (resource *ObservationDefinition) T_BodySite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.BodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.BodySite", resource.BodySite, optionsValueSet)
}
func (resource *ObservationDefinition) T_Method(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.Method", resource.Method, optionsValueSet)
}
func (resource *ObservationDefinition) T_PreferredReportName() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.PreferredReportName", nil)
	}
	return StringInput("ObservationDefinition.PreferredReportName", resource.PreferredReportName)
}
func (resource *ObservationDefinition) T_PermittedUnit(numPermittedUnit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PermittedUnit) >= numPermittedUnit {
		return CodingSelect("ObservationDefinition.PermittedUnit["+strconv.Itoa(numPermittedUnit)+"]", nil, optionsValueSet)
	}
	return CodingSelect("ObservationDefinition.PermittedUnit["+strconv.Itoa(numPermittedUnit)+"]", &resource.PermittedUnit[numPermittedUnit], optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueId(numQualifiedValue int) templ.Component {

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue {
		return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].Id", nil)
	}
	return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].Id", resource.QualifiedValue[numQualifiedValue].Id)
}
func (resource *ObservationDefinition) T_QualifiedValueContext(numQualifiedValue int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue {
		return CodeableConceptSelect("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].Context", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].Context", resource.QualifiedValue[numQualifiedValue].Context, optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueAppliesTo(numQualifiedValue int, numAppliesTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue || len(resource.QualifiedValue[numQualifiedValue].AppliesTo) >= numAppliesTo {
		return CodeableConceptSelect("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].AppliesTo["+strconv.Itoa(numAppliesTo)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].AppliesTo["+strconv.Itoa(numAppliesTo)+"]", &resource.QualifiedValue[numQualifiedValue].AppliesTo[numAppliesTo], optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueGender(numQualifiedValue int) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue {
		return CodeSelect("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].Gender", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].Gender", resource.QualifiedValue[numQualifiedValue].Gender, optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueCondition(numQualifiedValue int) templ.Component {

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue {
		return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].Condition", nil)
	}
	return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].Condition", resource.QualifiedValue[numQualifiedValue].Condition)
}
func (resource *ObservationDefinition) T_QualifiedValueRangeCategory(numQualifiedValue int) templ.Component {
	optionsValueSet := VSObservation_range_category

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue {
		return CodeSelect("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].RangeCategory", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].RangeCategory", resource.QualifiedValue[numQualifiedValue].RangeCategory, optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueValidCodedValueSet(numQualifiedValue int) templ.Component {

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue {
		return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].ValidCodedValueSet", nil)
	}
	return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].ValidCodedValueSet", resource.QualifiedValue[numQualifiedValue].ValidCodedValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueNormalCodedValueSet(numQualifiedValue int) templ.Component {

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue {
		return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].NormalCodedValueSet", nil)
	}
	return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].NormalCodedValueSet", resource.QualifiedValue[numQualifiedValue].NormalCodedValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueAbnormalCodedValueSet(numQualifiedValue int) templ.Component {

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue {
		return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].AbnormalCodedValueSet", nil)
	}
	return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].AbnormalCodedValueSet", resource.QualifiedValue[numQualifiedValue].AbnormalCodedValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueCriticalCodedValueSet(numQualifiedValue int) templ.Component {

	if resource == nil || len(resource.QualifiedValue) >= numQualifiedValue {
		return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].CriticalCodedValueSet", nil)
	}
	return StringInput("ObservationDefinition.QualifiedValue["+strconv.Itoa(numQualifiedValue)+"].CriticalCodedValueSet", resource.QualifiedValue[numQualifiedValue].CriticalCodedValueSet)
}
func (resource *ObservationDefinition) T_ComponentId(numComponent int) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return StringInput("ObservationDefinition.Component["+strconv.Itoa(numComponent)+"].Id", nil)
	}
	return StringInput("ObservationDefinition.Component["+strconv.Itoa(numComponent)+"].Id", resource.Component[numComponent].Id)
}
func (resource *ObservationDefinition) T_ComponentCode(numComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return CodeableConceptSelect("ObservationDefinition.Component["+strconv.Itoa(numComponent)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.Component["+strconv.Itoa(numComponent)+"].Code", &resource.Component[numComponent].Code, optionsValueSet)
}
func (resource *ObservationDefinition) T_ComponentPermittedDataType(numComponent int, numPermittedDataType int) templ.Component {
	optionsValueSet := VSPermitted_data_type

	if resource == nil || len(resource.Component) >= numComponent || len(resource.Component[numComponent].PermittedDataType) >= numPermittedDataType {
		return CodeSelect("ObservationDefinition.Component["+strconv.Itoa(numComponent)+"].PermittedDataType["+strconv.Itoa(numPermittedDataType)+"]", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.Component["+strconv.Itoa(numComponent)+"].PermittedDataType["+strconv.Itoa(numPermittedDataType)+"]", &resource.Component[numComponent].PermittedDataType[numPermittedDataType], optionsValueSet)
}
func (resource *ObservationDefinition) T_ComponentPermittedUnit(numComponent int, numPermittedUnit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent || len(resource.Component[numComponent].PermittedUnit) >= numPermittedUnit {
		return CodingSelect("ObservationDefinition.Component["+strconv.Itoa(numComponent)+"].PermittedUnit["+strconv.Itoa(numPermittedUnit)+"]", nil, optionsValueSet)
	}
	return CodingSelect("ObservationDefinition.Component["+strconv.Itoa(numComponent)+"].PermittedUnit["+strconv.Itoa(numPermittedUnit)+"]", &resource.Component[numComponent].PermittedUnit[numPermittedUnit], optionsValueSet)
}
