package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/NamingSystem
type NamingSystem struct {
	Id                     *string                `json:"id,omitempty"`
	Meta                   *Meta                  `json:"meta,omitempty"`
	ImplicitRules          *string                `json:"implicitRules,omitempty"`
	Language               *string                `json:"language,omitempty"`
	Text                   *Narrative             `json:"text,omitempty"`
	Contained              []Resource             `json:"contained,omitempty"`
	Extension              []Extension            `json:"extension,omitempty"`
	ModifierExtension      []Extension            `json:"modifierExtension,omitempty"`
	Url                    *string                `json:"url,omitempty"`
	Identifier             []Identifier           `json:"identifier,omitempty"`
	Version                *string                `json:"version,omitempty"`
	VersionAlgorithmString *string                `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                `json:"versionAlgorithmCoding,omitempty"`
	Name                   string                 `json:"name"`
	Title                  *string                `json:"title,omitempty"`
	Status                 string                 `json:"status"`
	Kind                   string                 `json:"kind"`
	Experimental           *bool                  `json:"experimental,omitempty"`
	Date                   FhirDateTime           `json:"date"`
	Publisher              *string                `json:"publisher,omitempty"`
	Contact                []ContactDetail        `json:"contact,omitempty"`
	Responsible            *string                `json:"responsible,omitempty"`
	Type                   *CodeableConcept       `json:"type,omitempty"`
	Description            *string                `json:"description,omitempty"`
	UseContext             []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept      `json:"jurisdiction,omitempty"`
	Purpose                *string                `json:"purpose,omitempty"`
	Copyright              *string                `json:"copyright,omitempty"`
	CopyrightLabel         *string                `json:"copyrightLabel,omitempty"`
	ApprovalDate           *FhirDate              `json:"approvalDate,omitempty"`
	LastReviewDate         *FhirDate              `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept      `json:"topic,omitempty"`
	Author                 []ContactDetail        `json:"author,omitempty"`
	Editor                 []ContactDetail        `json:"editor,omitempty"`
	Reviewer               []ContactDetail        `json:"reviewer,omitempty"`
	Endorser               []ContactDetail        `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact      `json:"relatedArtifact,omitempty"`
	Usage                  *string                `json:"usage,omitempty"`
	UniqueId               []NamingSystemUniqueId `json:"uniqueId"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NamingSystem
type NamingSystemUniqueId struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Value             string      `json:"value"`
	Preferred         *bool       `json:"preferred,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
	Period            *Period     `json:"period,omitempty"`
	Authoritative     *bool       `json:"authoritative,omitempty"`
}

type OtherNamingSystem NamingSystem

// on convert struct to json, automatically add resourceType=NamingSystem
func (r NamingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNamingSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherNamingSystem: OtherNamingSystem(r),
		ResourceType:      "NamingSystem",
	})
}
func (r NamingSystem) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "NamingSystem/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "NamingSystem"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *NamingSystem) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *NamingSystem) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *NamingSystem) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *NamingSystem) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *NamingSystem) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *NamingSystem) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Kind(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSNamingsystem_type

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *NamingSystem) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", &resource.Date, htmlAttrs)
}
func (resource *NamingSystem) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *NamingSystem) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *NamingSystem) T_Responsible(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("responsible", nil, htmlAttrs)
	}
	return StringInput("responsible", resource.Responsible, htmlAttrs)
}
func (resource *NamingSystem) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *NamingSystem) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *NamingSystem) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *NamingSystem) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *NamingSystem) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *NamingSystem) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *NamingSystem) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *NamingSystem) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *NamingSystem) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *NamingSystem) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *NamingSystem) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *NamingSystem) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *NamingSystem) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *NamingSystem) T_Usage(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("usage", nil, htmlAttrs)
	}
	return StringInput("usage", resource.Usage, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdType(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSNamingsystem_identifier_type

	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return CodeSelect("uniqueId["+strconv.Itoa(numUniqueId)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("uniqueId["+strconv.Itoa(numUniqueId)+"].type", &resource.UniqueId[numUniqueId].Type, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdValue(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return StringInput("uniqueId["+strconv.Itoa(numUniqueId)+"].value", nil, htmlAttrs)
	}
	return StringInput("uniqueId["+strconv.Itoa(numUniqueId)+"].value", &resource.UniqueId[numUniqueId].Value, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdPreferred(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return BoolInput("uniqueId["+strconv.Itoa(numUniqueId)+"].preferred", nil, htmlAttrs)
	}
	return BoolInput("uniqueId["+strconv.Itoa(numUniqueId)+"].preferred", resource.UniqueId[numUniqueId].Preferred, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdComment(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return StringInput("uniqueId["+strconv.Itoa(numUniqueId)+"].comment", nil, htmlAttrs)
	}
	return StringInput("uniqueId["+strconv.Itoa(numUniqueId)+"].comment", resource.UniqueId[numUniqueId].Comment, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdPeriod(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return PeriodInput("uniqueId["+strconv.Itoa(numUniqueId)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("uniqueId["+strconv.Itoa(numUniqueId)+"].period", resource.UniqueId[numUniqueId].Period, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdAuthoritative(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return BoolInput("uniqueId["+strconv.Itoa(numUniqueId)+"].authoritative", nil, htmlAttrs)
	}
	return BoolInput("uniqueId["+strconv.Itoa(numUniqueId)+"].authoritative", resource.UniqueId[numUniqueId].Authoritative, htmlAttrs)
}
