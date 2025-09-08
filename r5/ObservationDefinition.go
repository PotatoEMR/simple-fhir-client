package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date                   *time.Time                            `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                               `json:"publisher,omitempty"`
	Contact                []ContactDetail                       `json:"contact,omitempty"`
	Description            *string                               `json:"description,omitempty"`
	UseContext             []UsageContext                        `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                     `json:"jurisdiction,omitempty"`
	Purpose                *string                               `json:"purpose,omitempty"`
	Copyright              *string                               `json:"copyright,omitempty"`
	CopyrightLabel         *string                               `json:"copyrightLabel,omitempty"`
	ApprovalDate           *time.Time                            `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time                            `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (r ObservationDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ObservationDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "ObservationDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ObservationDefinition) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *ObservationDefinition) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *ObservationDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ObservationDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("ObservationDefinition.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ObservationDefinition.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *ObservationDefinition) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *ObservationDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ObservationDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ObservationDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("ObservationDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("ObservationDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ObservationDefinition) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("ObservationDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ObservationDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *ObservationDefinition) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ObservationDefinition) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *ObservationDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("ObservationDefinition.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ObservationDefinition) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ObservationDefinition) T_CopyrightLabel(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *ObservationDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("ObservationDefinition.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ObservationDefinition.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ObservationDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("ObservationDefinition.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("ObservationDefinition.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ObservationDefinition) T_DerivedFromCanonical(numDerivedFromCanonical int, htmlAttrs string) templ.Component {

	if resource == nil || numDerivedFromCanonical >= len(resource.DerivedFromCanonical) {
		return StringInput("ObservationDefinition.DerivedFromCanonical."+strconv.Itoa(numDerivedFromCanonical)+".", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.DerivedFromCanonical."+strconv.Itoa(numDerivedFromCanonical)+".", &resource.DerivedFromCanonical[numDerivedFromCanonical], htmlAttrs)
}
func (resource *ObservationDefinition) T_DerivedFromUri(numDerivedFromUri int, htmlAttrs string) templ.Component {

	if resource == nil || numDerivedFromUri >= len(resource.DerivedFromUri) {
		return StringInput("ObservationDefinition.DerivedFromUri."+strconv.Itoa(numDerivedFromUri)+".", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.DerivedFromUri."+strconv.Itoa(numDerivedFromUri)+".", &resource.DerivedFromUri[numDerivedFromUri], htmlAttrs)
}
func (resource *ObservationDefinition) T_Subject(numSubject int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSubject >= len(resource.Subject) {
		return CodeableConceptSelect("ObservationDefinition.Subject."+strconv.Itoa(numSubject)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.Subject."+strconv.Itoa(numSubject)+".", &resource.Subject[numSubject], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_PerformerType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.PerformerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.PerformerType", resource.PerformerType, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("ObservationDefinition.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_PermittedDataType(numPermittedDataType int, htmlAttrs string) templ.Component {
	optionsValueSet := VSPermitted_data_type

	if resource == nil || numPermittedDataType >= len(resource.PermittedDataType) {
		return CodeSelect("ObservationDefinition.PermittedDataType."+strconv.Itoa(numPermittedDataType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ObservationDefinition.PermittedDataType."+strconv.Itoa(numPermittedDataType)+".", &resource.PermittedDataType[numPermittedDataType], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_MultipleResultsAllowed(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("ObservationDefinition.MultipleResultsAllowed", nil, htmlAttrs)
	}
	return BoolInput("ObservationDefinition.MultipleResultsAllowed", resource.MultipleResultsAllowed, htmlAttrs)
}
func (resource *ObservationDefinition) T_BodySite(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.BodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_Method(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.Method", resource.Method, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_PreferredReportName(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.PreferredReportName", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.PreferredReportName", resource.PreferredReportName, htmlAttrs)
}
func (resource *ObservationDefinition) T_PermittedUnit(numPermittedUnit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPermittedUnit >= len(resource.PermittedUnit) {
		return CodingSelect("ObservationDefinition.PermittedUnit."+strconv.Itoa(numPermittedUnit)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ObservationDefinition.PermittedUnit."+strconv.Itoa(numPermittedUnit)+".", &resource.PermittedUnit[numPermittedUnit], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedValueContext(numQualifiedValue int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numQualifiedValue >= len(resource.QualifiedValue) {
		return CodeableConceptSelect("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..Context", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..Context", resource.QualifiedValue[numQualifiedValue].Context, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedValueAppliesTo(numQualifiedValue int, numAppliesTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numQualifiedValue >= len(resource.QualifiedValue) || numAppliesTo >= len(resource.QualifiedValue[numQualifiedValue].AppliesTo) {
		return CodeableConceptSelect("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..AppliesTo."+strconv.Itoa(numAppliesTo)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..AppliesTo."+strconv.Itoa(numAppliesTo)+".", &resource.QualifiedValue[numQualifiedValue].AppliesTo[numAppliesTo], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedValueGender(numQualifiedValue int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil || numQualifiedValue >= len(resource.QualifiedValue) {
		return CodeSelect("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..Gender", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..Gender", resource.QualifiedValue[numQualifiedValue].Gender, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedValueCondition(numQualifiedValue int, htmlAttrs string) templ.Component {

	if resource == nil || numQualifiedValue >= len(resource.QualifiedValue) {
		return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..Condition", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..Condition", resource.QualifiedValue[numQualifiedValue].Condition, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedValueRangeCategory(numQualifiedValue int, htmlAttrs string) templ.Component {
	optionsValueSet := VSObservation_range_category

	if resource == nil || numQualifiedValue >= len(resource.QualifiedValue) {
		return CodeSelect("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..RangeCategory", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..RangeCategory", resource.QualifiedValue[numQualifiedValue].RangeCategory, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedValueValidCodedValueSet(numQualifiedValue int, htmlAttrs string) templ.Component {

	if resource == nil || numQualifiedValue >= len(resource.QualifiedValue) {
		return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..ValidCodedValueSet", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..ValidCodedValueSet", resource.QualifiedValue[numQualifiedValue].ValidCodedValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedValueNormalCodedValueSet(numQualifiedValue int, htmlAttrs string) templ.Component {

	if resource == nil || numQualifiedValue >= len(resource.QualifiedValue) {
		return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..NormalCodedValueSet", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..NormalCodedValueSet", resource.QualifiedValue[numQualifiedValue].NormalCodedValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedValueAbnormalCodedValueSet(numQualifiedValue int, htmlAttrs string) templ.Component {

	if resource == nil || numQualifiedValue >= len(resource.QualifiedValue) {
		return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..AbnormalCodedValueSet", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..AbnormalCodedValueSet", resource.QualifiedValue[numQualifiedValue].AbnormalCodedValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedValueCriticalCodedValueSet(numQualifiedValue int, htmlAttrs string) templ.Component {

	if resource == nil || numQualifiedValue >= len(resource.QualifiedValue) {
		return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..CriticalCodedValueSet", nil, htmlAttrs)
	}
	return StringInput("ObservationDefinition.QualifiedValue."+strconv.Itoa(numQualifiedValue)+"..CriticalCodedValueSet", resource.QualifiedValue[numQualifiedValue].CriticalCodedValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_ComponentCode(numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numComponent >= len(resource.Component) {
		return CodeableConceptSelect("ObservationDefinition.Component."+strconv.Itoa(numComponent)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ObservationDefinition.Component."+strconv.Itoa(numComponent)+"..Code", &resource.Component[numComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_ComponentPermittedDataType(numComponent int, numPermittedDataType int, htmlAttrs string) templ.Component {
	optionsValueSet := VSPermitted_data_type

	if resource == nil || numComponent >= len(resource.Component) || numPermittedDataType >= len(resource.Component[numComponent].PermittedDataType) {
		return CodeSelect("ObservationDefinition.Component."+strconv.Itoa(numComponent)+"..PermittedDataType."+strconv.Itoa(numPermittedDataType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ObservationDefinition.Component."+strconv.Itoa(numComponent)+"..PermittedDataType."+strconv.Itoa(numPermittedDataType)+".", &resource.Component[numComponent].PermittedDataType[numPermittedDataType], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_ComponentPermittedUnit(numComponent int, numPermittedUnit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numComponent >= len(resource.Component) || numPermittedUnit >= len(resource.Component[numComponent].PermittedUnit) {
		return CodingSelect("ObservationDefinition.Component."+strconv.Itoa(numComponent)+"..PermittedUnit."+strconv.Itoa(numPermittedUnit)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ObservationDefinition.Component."+strconv.Itoa(numComponent)+"..PermittedUnit."+strconv.Itoa(numPermittedUnit)+".", &resource.Component[numComponent].PermittedUnit[numPermittedUnit], optionsValueSet, htmlAttrs)
}
