package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	VersionAlgorithmString *string                               `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                               `json:"versionAlgorithmCoding"`
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

func (resource *ObservationDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ObservationDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ObservationDefinition) T_Jurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *ObservationDefinition) T_Subject(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("subject", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subject", &resource.Subject[0], optionsValueSet)
}
func (resource *ObservationDefinition) T_PerformerType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("performerType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("performerType", resource.PerformerType, optionsValueSet)
}
func (resource *ObservationDefinition) T_Category(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *ObservationDefinition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet)
}
func (resource *ObservationDefinition) T_PermittedDataType() templ.Component {
	optionsValueSet := VSPermitted_data_type

	if resource == nil {
		return CodeSelect("permittedDataType", nil, optionsValueSet)
	}
	return CodeSelect("permittedDataType", &resource.PermittedDataType[0], optionsValueSet)
}
func (resource *ObservationDefinition) T_BodySite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", resource.BodySite, optionsValueSet)
}
func (resource *ObservationDefinition) T_Method(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("method", resource.Method, optionsValueSet)
}
func (resource *ObservationDefinition) T_PermittedUnit(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("permittedUnit", nil, optionsValueSet)
	}
	return CodingSelect("permittedUnit", &resource.PermittedUnit[0], optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueContext(numQualifiedValue int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.QualifiedValue) >= numQualifiedValue {
		return CodeableConceptSelect("context", nil, optionsValueSet)
	}
	return CodeableConceptSelect("context", resource.QualifiedValue[numQualifiedValue].Context, optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueAppliesTo(numQualifiedValue int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.QualifiedValue) >= numQualifiedValue {
		return CodeableConceptSelect("appliesTo", nil, optionsValueSet)
	}
	return CodeableConceptSelect("appliesTo", &resource.QualifiedValue[numQualifiedValue].AppliesTo[0], optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueGender(numQualifiedValue int) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil && len(resource.QualifiedValue) >= numQualifiedValue {
		return CodeSelect("gender", nil, optionsValueSet)
	}
	return CodeSelect("gender", resource.QualifiedValue[numQualifiedValue].Gender, optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedValueRangeCategory(numQualifiedValue int) templ.Component {
	optionsValueSet := VSObservation_range_category

	if resource == nil && len(resource.QualifiedValue) >= numQualifiedValue {
		return CodeSelect("rangeCategory", nil, optionsValueSet)
	}
	return CodeSelect("rangeCategory", resource.QualifiedValue[numQualifiedValue].RangeCategory, optionsValueSet)
}
func (resource *ObservationDefinition) T_ComponentCode(numComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Component) >= numComponent {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Component[numComponent].Code, optionsValueSet)
}
func (resource *ObservationDefinition) T_ComponentPermittedDataType(numComponent int) templ.Component {
	optionsValueSet := VSPermitted_data_type

	if resource == nil && len(resource.Component) >= numComponent {
		return CodeSelect("permittedDataType", nil, optionsValueSet)
	}
	return CodeSelect("permittedDataType", &resource.Component[numComponent].PermittedDataType[0], optionsValueSet)
}
func (resource *ObservationDefinition) T_ComponentPermittedUnit(numComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Component) >= numComponent {
		return CodingSelect("permittedUnit", nil, optionsValueSet)
	}
	return CodingSelect("permittedUnit", &resource.Component[numComponent].PermittedUnit[0], optionsValueSet)
}
