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
	Date                   time.Time              `json:"date,format:'2006-01-02T15:04:05Z07:00'"`
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
	ApprovalDate           *time.Time             `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time             `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (resource *NamingSystem) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *NamingSystem) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *NamingSystem) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *NamingSystem) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", &resource.Name, htmlAttrs)
}
func (resource *NamingSystem) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *NamingSystem) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Kind(htmlAttrs string) templ.Component {
	optionsValueSet := VSNamingsystem_type

	if resource == nil {
		return CodeSelect("Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *NamingSystem) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", &resource.Date, htmlAttrs)
}
func (resource *NamingSystem) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *NamingSystem) T_Responsible(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Responsible", nil, htmlAttrs)
	}
	return StringInput("Responsible", resource.Responsible, htmlAttrs)
}
func (resource *NamingSystem) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *NamingSystem) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *NamingSystem) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *NamingSystem) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *NamingSystem) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *NamingSystem) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *NamingSystem) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Usage", nil, htmlAttrs)
	}
	return StringInput("Usage", resource.Usage, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdType(numUniqueId int, htmlAttrs string) templ.Component {
	optionsValueSet := VSNamingsystem_identifier_type

	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return CodeSelect("UniqueId["+strconv.Itoa(numUniqueId)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("UniqueId["+strconv.Itoa(numUniqueId)+"]Type", &resource.UniqueId[numUniqueId].Type, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdValue(numUniqueId int, htmlAttrs string) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return StringInput("UniqueId["+strconv.Itoa(numUniqueId)+"]Value", nil, htmlAttrs)
	}
	return StringInput("UniqueId["+strconv.Itoa(numUniqueId)+"]Value", &resource.UniqueId[numUniqueId].Value, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdPreferred(numUniqueId int, htmlAttrs string) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return BoolInput("UniqueId["+strconv.Itoa(numUniqueId)+"]Preferred", nil, htmlAttrs)
	}
	return BoolInput("UniqueId["+strconv.Itoa(numUniqueId)+"]Preferred", resource.UniqueId[numUniqueId].Preferred, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdComment(numUniqueId int, htmlAttrs string) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return StringInput("UniqueId["+strconv.Itoa(numUniqueId)+"]Comment", nil, htmlAttrs)
	}
	return StringInput("UniqueId["+strconv.Itoa(numUniqueId)+"]Comment", resource.UniqueId[numUniqueId].Comment, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdAuthoritative(numUniqueId int, htmlAttrs string) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return BoolInput("UniqueId["+strconv.Itoa(numUniqueId)+"]Authoritative", nil, htmlAttrs)
	}
	return BoolInput("UniqueId["+strconv.Itoa(numUniqueId)+"]Authoritative", resource.UniqueId[numUniqueId].Authoritative, htmlAttrs)
}
