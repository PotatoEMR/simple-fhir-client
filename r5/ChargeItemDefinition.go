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
	Date                   *time.Time                          `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                             `json:"publisher,omitempty"`
	Contact                []ContactDetail                     `json:"contact,omitempty"`
	Description            *string                             `json:"description,omitempty"`
	UseContext             []UsageContext                      `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                   `json:"jurisdiction,omitempty"`
	Purpose                *string                             `json:"purpose,omitempty"`
	Copyright              *string                             `json:"copyright,omitempty"`
	CopyrightLabel         *string                             `json:"copyrightLabel,omitempty"`
	ApprovalDate           *time.Time                          `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time                          `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (resource *ChargeItemDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("ChargeItemDefinition.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ChargeItemDefinition.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_DerivedFromUri(numDerivedFromUri int, htmlAttrs string) templ.Component {
	if resource == nil || numDerivedFromUri >= len(resource.DerivedFromUri) {
		return StringInput("ChargeItemDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", &resource.DerivedFromUri[numDerivedFromUri], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_PartOf(numPartOf int, htmlAttrs string) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return StringInput("ChargeItemDefinition.PartOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.PartOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Replaces(numReplaces int, htmlAttrs string) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return StringInput("ChargeItemDefinition.Replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.Replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ChargeItemDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ChargeItemDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ChargeItemDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("ChargeItemDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ChargeItemDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ChargeItemDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("ChargeItemDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ChargeItemDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ChargeItemDefinition.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("ChargeItemDefinition.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ChargeItemDefinition.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ChargeItemDefinition.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ChargeItemDefinition.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("ChargeItemDefinition.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ChargeItemDefinition.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ChargeItemDefinition.Code", resource.Code, optionsValueSet, htmlAttrs)
}
