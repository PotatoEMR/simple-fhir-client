package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystem struct {
	Id                     *string              `json:"id,omitempty"`
	Meta                   *Meta                `json:"meta,omitempty"`
	ImplicitRules          *string              `json:"implicitRules,omitempty"`
	Language               *string              `json:"language,omitempty"`
	Text                   *Narrative           `json:"text,omitempty"`
	Contained              []Resource           `json:"contained,omitempty"`
	Extension              []Extension          `json:"extension,omitempty"`
	ModifierExtension      []Extension          `json:"modifierExtension,omitempty"`
	Url                    *string              `json:"url,omitempty"`
	Identifier             []Identifier         `json:"identifier,omitempty"`
	Version                *string              `json:"version,omitempty"`
	VersionAlgorithmString *string              `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding              `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string              `json:"name,omitempty"`
	Title                  *string              `json:"title,omitempty"`
	Status                 string               `json:"status"`
	Experimental           *bool                `json:"experimental,omitempty"`
	Date                   *FhirDateTime        `json:"date,omitempty"`
	Publisher              *string              `json:"publisher,omitempty"`
	Contact                []ContactDetail      `json:"contact,omitempty"`
	Description            *string              `json:"description,omitempty"`
	UseContext             []UsageContext       `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept    `json:"jurisdiction,omitempty"`
	Purpose                *string              `json:"purpose,omitempty"`
	Copyright              *string              `json:"copyright,omitempty"`
	CopyrightLabel         *string              `json:"copyrightLabel,omitempty"`
	ApprovalDate           *FhirDate            `json:"approvalDate,omitempty"`
	LastReviewDate         *FhirDate            `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period              `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept    `json:"topic,omitempty"`
	Author                 []ContactDetail      `json:"author,omitempty"`
	Editor                 []ContactDetail      `json:"editor,omitempty"`
	Reviewer               []ContactDetail      `json:"reviewer,omitempty"`
	Endorser               []ContactDetail      `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact    `json:"relatedArtifact,omitempty"`
	CaseSensitive          *bool                `json:"caseSensitive,omitempty"`
	ValueSet               *string              `json:"valueSet,omitempty"`
	HierarchyMeaning       *string              `json:"hierarchyMeaning,omitempty"`
	Compositional          *bool                `json:"compositional,omitempty"`
	VersionNeeded          *bool                `json:"versionNeeded,omitempty"`
	Content                string               `json:"content"`
	Supplements            *string              `json:"supplements,omitempty"`
	Count                  *int                 `json:"count,omitempty"`
	Filter                 []CodeSystemFilter   `json:"filter,omitempty"`
	Property               []CodeSystemProperty `json:"property,omitempty"`
	Concept                []CodeSystemConcept  `json:"concept,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Description       *string     `json:"description,omitempty"`
	Operator          []string    `json:"operator"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemProperty struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Uri               *string     `json:"uri,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemConcept struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              string                         `json:"code"`
	Display           *string                        `json:"display,omitempty"`
	Definition        *string                        `json:"definition,omitempty"`
	Designation       []CodeSystemConceptDesignation `json:"designation,omitempty"`
	Property          []CodeSystemConceptProperty    `json:"property,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemConceptDesignation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	AdditionalUse     []Coding    `json:"additionalUse,omitempty"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CodeSystem
type CodeSystemConceptProperty struct {
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Code              string       `json:"code"`
	ValueCode         string       `json:"valueCode"`
	ValueCoding       Coding       `json:"valueCoding"`
	ValueString       string       `json:"valueString"`
	ValueInteger      int          `json:"valueInteger"`
	ValueBoolean      bool         `json:"valueBoolean"`
	ValueDateTime     FhirDateTime `json:"valueDateTime"`
	ValueDecimal      float64      `json:"valueDecimal"`
}

type OtherCodeSystem CodeSystem

// struct -> json, automatically add resourceType=Patient
func (r CodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCodeSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherCodeSystem: OtherCodeSystem(r),
		ResourceType:    "CodeSystem",
	})
}

func (r CodeSystem) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "CodeSystem/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "CodeSystem"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r CodeSystem) ResourceType() string {
	return "CodeSystem"
}

func (resource *CodeSystem) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *CodeSystem) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *CodeSystem) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *CodeSystem) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *CodeSystem) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *CodeSystem) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *CodeSystem) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *CodeSystem) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *CodeSystem) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *CodeSystem) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *CodeSystem) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *CodeSystem) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *CodeSystem) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *CodeSystem) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *CodeSystem) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *CodeSystem) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *CodeSystem) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *CodeSystem) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *CodeSystem) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *CodeSystem) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *CodeSystem) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *CodeSystem) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *CodeSystem) T_CaseSensitive(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("caseSensitive", nil, htmlAttrs)
	}
	return BoolInput("caseSensitive", resource.CaseSensitive, htmlAttrs)
}
func (resource *CodeSystem) T_ValueSet(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("valueSet", nil, htmlAttrs)
	}
	return StringInput("valueSet", resource.ValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_HierarchyMeaning(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCodesystem_hierarchy_meaning

	if resource == nil {
		return CodeSelect("hierarchyMeaning", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("hierarchyMeaning", resource.HierarchyMeaning, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Compositional(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("compositional", nil, htmlAttrs)
	}
	return BoolInput("compositional", resource.Compositional, htmlAttrs)
}
func (resource *CodeSystem) T_VersionNeeded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("versionNeeded", nil, htmlAttrs)
	}
	return BoolInput("versionNeeded", resource.VersionNeeded, htmlAttrs)
}
func (resource *CodeSystem) T_Content(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCodesystem_content_mode

	if resource == nil {
		return CodeSelect("content", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("content", &resource.Content, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Supplements(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("supplements", nil, htmlAttrs)
	}
	return StringInput("supplements", resource.Supplements, htmlAttrs)
}
func (resource *CodeSystem) T_Count(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("count", nil, htmlAttrs)
	}
	return IntInput("count", resource.Count, htmlAttrs)
}
func (resource *CodeSystem) T_FilterCode(numFilter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFilter >= len(resource.Filter) {
		return CodeSelect("filter["+strconv.Itoa(numFilter)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("filter["+strconv.Itoa(numFilter)+"].code", &resource.Filter[numFilter].Code, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_FilterDescription(numFilter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFilter >= len(resource.Filter) {
		return StringInput("filter["+strconv.Itoa(numFilter)+"].description", nil, htmlAttrs)
	}
	return StringInput("filter["+strconv.Itoa(numFilter)+"].description", resource.Filter[numFilter].Description, htmlAttrs)
}
func (resource *CodeSystem) T_FilterOperator(numFilter int, numOperator int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFilter_operator

	if resource == nil || numFilter >= len(resource.Filter) || numOperator >= len(resource.Filter[numFilter].Operator) {
		return CodeSelect("filter["+strconv.Itoa(numFilter)+"].operator["+strconv.Itoa(numOperator)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("filter["+strconv.Itoa(numFilter)+"].operator["+strconv.Itoa(numOperator)+"]", &resource.Filter[numFilter].Operator[numOperator], optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_FilterValue(numFilter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFilter >= len(resource.Filter) {
		return StringInput("filter["+strconv.Itoa(numFilter)+"].value", nil, htmlAttrs)
	}
	return StringInput("filter["+strconv.Itoa(numFilter)+"].value", &resource.Filter[numFilter].Value, htmlAttrs)
}
func (resource *CodeSystem) T_PropertyCode(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeSelect("property["+strconv.Itoa(numProperty)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("property["+strconv.Itoa(numProperty)+"].code", &resource.Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_PropertyUri(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].uri", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].uri", resource.Property[numProperty].Uri, htmlAttrs)
}
func (resource *CodeSystem) T_PropertyDescription(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].description", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].description", resource.Property[numProperty].Description, htmlAttrs)
}
func (resource *CodeSystem) T_PropertyType(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConcept_property_type

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptCode(numConcept int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) {
		return CodeSelect("concept["+strconv.Itoa(numConcept)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("concept["+strconv.Itoa(numConcept)+"].code", &resource.Concept[numConcept].Code, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDisplay(numConcept int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) {
		return StringInput("concept["+strconv.Itoa(numConcept)+"].display", nil, htmlAttrs)
	}
	return StringInput("concept["+strconv.Itoa(numConcept)+"].display", resource.Concept[numConcept].Display, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDefinition(numConcept int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) {
		return StringInput("concept["+strconv.Itoa(numConcept)+"].definition", nil, htmlAttrs)
	}
	return StringInput("concept["+strconv.Itoa(numConcept)+"].definition", resource.Concept[numConcept].Definition, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDesignationUse(numConcept int, numDesignation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numDesignation >= len(resource.Concept[numConcept].Designation) {
		return CodingSelect("concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].use", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].use", resource.Concept[numConcept].Designation[numDesignation].Use, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDesignationAdditionalUse(numConcept int, numDesignation int, numAdditionalUse int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numDesignation >= len(resource.Concept[numConcept].Designation) || numAdditionalUse >= len(resource.Concept[numConcept].Designation[numDesignation].AdditionalUse) {
		return CodingSelect("concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].additionalUse["+strconv.Itoa(numAdditionalUse)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].additionalUse["+strconv.Itoa(numAdditionalUse)+"]", &resource.Concept[numConcept].Designation[numDesignation].AdditionalUse[numAdditionalUse], optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDesignationValue(numConcept int, numDesignation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numDesignation >= len(resource.Concept[numConcept].Designation) {
		return StringInput("concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].value", nil, htmlAttrs)
	}
	return StringInput("concept["+strconv.Itoa(numConcept)+"].designation["+strconv.Itoa(numDesignation)+"].value", &resource.Concept[numConcept].Designation[numDesignation].Value, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyCode(numConcept int, numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return CodeSelect("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].code", &resource.Concept[numConcept].Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueCode(numConcept int, numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return CodeSelect("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueCode", &resource.Concept[numConcept].Property[numProperty].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueCoding(numConcept int, numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return CodingSelect("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueCoding", &resource.Concept[numConcept].Property[numProperty].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueString(numConcept int, numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return StringInput("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueString", &resource.Concept[numConcept].Property[numProperty].ValueString, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueInteger(numConcept int, numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return IntInput("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueInteger", &resource.Concept[numConcept].Property[numProperty].ValueInteger, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueBoolean(numConcept int, numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return BoolInput("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueBoolean", &resource.Concept[numConcept].Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueDateTime(numConcept int, numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return FhirDateTimeInput("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueDateTime", &resource.Concept[numConcept].Property[numProperty].ValueDateTime, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueDecimal(numConcept int, numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return Float64Input("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("concept["+strconv.Itoa(numConcept)+"].property["+strconv.Itoa(numProperty)+"].valueDecimal", &resource.Concept[numConcept].Property[numProperty].ValueDecimal, htmlAttrs)
}
