package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinition struct {
	Id                     *string                             `json:"id,omitempty"`
	Meta                   *Meta                               `json:"meta,omitempty"`
	ImplicitRules          *string                             `json:"implicitRules,omitempty"`
	Language               *string                             `json:"language,omitempty"`
	Text                   *Narrative                          `json:"text,omitempty"`
	Contained              []Resource                          `json:"contained,omitempty"`
	Extension              []Extension                         `json:"extension,omitempty"`
	ModifierExtension      []Extension                         `json:"modifierExtension,omitempty"`
	Url                    *string                             `json:"url,omitempty"`
	Identifier             []Identifier                        `json:"identifier,omitempty"`
	Version                *string                             `json:"version,omitempty"`
	VersionAlgorithmString *string                             `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                             `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                             `json:"name,omitempty"`
	Title                  *string                             `json:"title,omitempty"`
	DerivedFromUri         []string                            `json:"derivedFromUri,omitempty"`
	PartOf                 []string                            `json:"partOf,omitempty"`
	Replaces               []string                            `json:"replaces,omitempty"`
	Status                 string                              `json:"status"`
	Experimental           *bool                               `json:"experimental,omitempty"`
	Date                   *FhirDateTime                       `json:"date,omitempty"`
	Publisher              *string                             `json:"publisher,omitempty"`
	Contact                []ContactDetail                     `json:"contact,omitempty"`
	Description            *string                             `json:"description,omitempty"`
	UseContext             []UsageContext                      `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                   `json:"jurisdiction,omitempty"`
	Purpose                *string                             `json:"purpose,omitempty"`
	Copyright              *string                             `json:"copyright,omitempty"`
	CopyrightLabel         *string                             `json:"copyrightLabel,omitempty"`
	ApprovalDate           *FhirDate                           `json:"approvalDate,omitempty"`
	LastReviewDate         *FhirDate                           `json:"lastReviewDate,omitempty"`
	Code                   *CodeableConcept                    `json:"code,omitempty"`
	Instance               []Reference                         `json:"instance,omitempty"`
	Applicability          []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`
	PropertyGroup          []ChargeItemDefinitionPropertyGroup `json:"propertyGroup,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinitionApplicability struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Condition         *Expression      `json:"condition,omitempty"`
	EffectivePeriod   *Period          `json:"effectivePeriod,omitempty"`
	RelatedArtifact   *RelatedArtifact `json:"relatedArtifact,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinitionPropertyGroup struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	PriceComponent    []MonetaryComponent `json:"priceComponent,omitempty"`
}

type OtherChargeItemDefinition ChargeItemDefinition

// on convert struct to json, automatically add resourceType=ChargeItemDefinition
func (r ChargeItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItemDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItemDefinition: OtherChargeItemDefinition(r),
		ResourceType:              "ChargeItemDefinition",
	})
}
func (r ChargeItemDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ChargeItemDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ChargeItemDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ChargeItemDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_DerivedFromUri(numDerivedFromUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFromUri >= len(resource.DerivedFromUri) {
		return StringInput("derivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", nil, htmlAttrs)
	}
	return StringInput("derivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", &resource.DerivedFromUri[numDerivedFromUri], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_PartOf(numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return StringInput("partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return StringInput("partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Replaces(numReplaces int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return StringInput("replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return StringInput("replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Instance(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return ReferenceInput("instance["+strconv.Itoa(numInstance)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("instance["+strconv.Itoa(numInstance)+"]", &resource.Instance[numInstance], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_ApplicabilityCondition(numApplicability int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numApplicability >= len(resource.Applicability) {
		return ExpressionInput("applicability["+strconv.Itoa(numApplicability)+"].condition", nil, htmlAttrs)
	}
	return ExpressionInput("applicability["+strconv.Itoa(numApplicability)+"].condition", resource.Applicability[numApplicability].Condition, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_ApplicabilityEffectivePeriod(numApplicability int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numApplicability >= len(resource.Applicability) {
		return PeriodInput("applicability["+strconv.Itoa(numApplicability)+"].effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("applicability["+strconv.Itoa(numApplicability)+"].effectivePeriod", resource.Applicability[numApplicability].EffectivePeriod, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_ApplicabilityRelatedArtifact(numApplicability int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numApplicability >= len(resource.Applicability) {
		return RelatedArtifactInput("applicability["+strconv.Itoa(numApplicability)+"].relatedArtifact", nil, htmlAttrs)
	}
	return RelatedArtifactInput("applicability["+strconv.Itoa(numApplicability)+"].relatedArtifact", resource.Applicability[numApplicability].RelatedArtifact, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_PropertyGroupPriceComponent(numPropertyGroup int, numPriceComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPropertyGroup >= len(resource.PropertyGroup) || numPriceComponent >= len(resource.PropertyGroup[numPropertyGroup].PriceComponent) {
		return MonetaryComponentInput("propertyGroup["+strconv.Itoa(numPropertyGroup)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"]", nil, htmlAttrs)
	}
	return MonetaryComponentInput("propertyGroup["+strconv.Itoa(numPropertyGroup)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"]", &resource.PropertyGroup[numPropertyGroup].PriceComponent[numPriceComponent], htmlAttrs)
}
