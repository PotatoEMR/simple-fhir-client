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
	Date                   *time.Time           `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string              `json:"publisher,omitempty"`
	Contact                []ContactDetail      `json:"contact,omitempty"`
	Description            *string              `json:"description,omitempty"`
	UseContext             []UsageContext       `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept    `json:"jurisdiction,omitempty"`
	Purpose                *string              `json:"purpose,omitempty"`
	Copyright              *string              `json:"copyright,omitempty"`
	CopyrightLabel         *string              `json:"copyrightLabel,omitempty"`
	ApprovalDate           *time.Time           `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time           `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	ValueCode         string      `json:"valueCode"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueString       string      `json:"valueString"`
	ValueInteger      int         `json:"valueInteger"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDateTime     time.Time   `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
	ValueDecimal      float64     `json:"valueDecimal"`
}

type OtherCodeSystem CodeSystem

// on convert struct to json, automatically add resourceType=CodeSystem
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
func (resource *CodeSystem) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *CodeSystem) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *CodeSystem) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *CodeSystem) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *CodeSystem) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *CodeSystem) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *CodeSystem) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *CodeSystem) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *CodeSystem) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *CodeSystem) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *CodeSystem) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *CodeSystem) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *CodeSystem) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *CodeSystem) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *CodeSystem) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_CaseSensitive(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("CaseSensitive", nil, htmlAttrs)
	}
	return BoolInput("CaseSensitive", resource.CaseSensitive, htmlAttrs)
}
func (resource *CodeSystem) T_ValueSet(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ValueSet", nil, htmlAttrs)
	}
	return StringInput("ValueSet", resource.ValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_HierarchyMeaning(htmlAttrs string) templ.Component {
	optionsValueSet := VSCodesystem_hierarchy_meaning

	if resource == nil {
		return CodeSelect("HierarchyMeaning", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("HierarchyMeaning", resource.HierarchyMeaning, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Compositional(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Compositional", nil, htmlAttrs)
	}
	return BoolInput("Compositional", resource.Compositional, htmlAttrs)
}
func (resource *CodeSystem) T_VersionNeeded(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("VersionNeeded", nil, htmlAttrs)
	}
	return BoolInput("VersionNeeded", resource.VersionNeeded, htmlAttrs)
}
func (resource *CodeSystem) T_Content(htmlAttrs string) templ.Component {
	optionsValueSet := VSCodesystem_content_mode

	if resource == nil {
		return CodeSelect("Content", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Content", &resource.Content, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_Supplements(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Supplements", nil, htmlAttrs)
	}
	return StringInput("Supplements", resource.Supplements, htmlAttrs)
}
func (resource *CodeSystem) T_Count(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Count", nil, htmlAttrs)
	}
	return IntInput("Count", resource.Count, htmlAttrs)
}
func (resource *CodeSystem) T_FilterCode(numFilter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFilter >= len(resource.Filter) {
		return CodeSelect("Filter["+strconv.Itoa(numFilter)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Filter["+strconv.Itoa(numFilter)+"]Code", &resource.Filter[numFilter].Code, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_FilterDescription(numFilter int, htmlAttrs string) templ.Component {
	if resource == nil || numFilter >= len(resource.Filter) {
		return StringInput("Filter["+strconv.Itoa(numFilter)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Filter["+strconv.Itoa(numFilter)+"]Description", resource.Filter[numFilter].Description, htmlAttrs)
}
func (resource *CodeSystem) T_FilterOperator(numFilter int, numOperator int, htmlAttrs string) templ.Component {
	optionsValueSet := VSFilter_operator

	if resource == nil || numFilter >= len(resource.Filter) || numOperator >= len(resource.Filter[numFilter].Operator) {
		return CodeSelect("Filter["+strconv.Itoa(numFilter)+"]Operator["+strconv.Itoa(numOperator)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Filter["+strconv.Itoa(numFilter)+"]Operator["+strconv.Itoa(numOperator)+"]", &resource.Filter[numFilter].Operator[numOperator], optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_FilterValue(numFilter int, htmlAttrs string) templ.Component {
	if resource == nil || numFilter >= len(resource.Filter) {
		return StringInput("Filter["+strconv.Itoa(numFilter)+"]Value", nil, htmlAttrs)
	}
	return StringInput("Filter["+strconv.Itoa(numFilter)+"]Value", &resource.Filter[numFilter].Value, htmlAttrs)
}
func (resource *CodeSystem) T_PropertyCode(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeSelect("Property["+strconv.Itoa(numProperty)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Property["+strconv.Itoa(numProperty)+"]Code", &resource.Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_PropertyUri(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("Property["+strconv.Itoa(numProperty)+"]Uri", nil, htmlAttrs)
	}
	return StringInput("Property["+strconv.Itoa(numProperty)+"]Uri", resource.Property[numProperty].Uri, htmlAttrs)
}
func (resource *CodeSystem) T_PropertyDescription(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("Property["+strconv.Itoa(numProperty)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Property["+strconv.Itoa(numProperty)+"]Description", resource.Property[numProperty].Description, htmlAttrs)
}
func (resource *CodeSystem) T_PropertyType(numProperty int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConcept_property_type

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeSelect("Property["+strconv.Itoa(numProperty)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Property["+strconv.Itoa(numProperty)+"]Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptCode(numConcept int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) {
		return CodeSelect("Concept["+strconv.Itoa(numConcept)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Concept["+strconv.Itoa(numConcept)+"]Code", &resource.Concept[numConcept].Code, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDisplay(numConcept int, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) {
		return StringInput("Concept["+strconv.Itoa(numConcept)+"]Display", nil, htmlAttrs)
	}
	return StringInput("Concept["+strconv.Itoa(numConcept)+"]Display", resource.Concept[numConcept].Display, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDefinition(numConcept int, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) {
		return StringInput("Concept["+strconv.Itoa(numConcept)+"]Definition", nil, htmlAttrs)
	}
	return StringInput("Concept["+strconv.Itoa(numConcept)+"]Definition", resource.Concept[numConcept].Definition, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDesignationUse(numConcept int, numDesignation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numDesignation >= len(resource.Concept[numConcept].Designation) {
		return CodingSelect("Concept["+strconv.Itoa(numConcept)+"]Designation["+strconv.Itoa(numDesignation)+"].Use", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Concept["+strconv.Itoa(numConcept)+"]Designation["+strconv.Itoa(numDesignation)+"].Use", resource.Concept[numConcept].Designation[numDesignation].Use, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDesignationAdditionalUse(numConcept int, numDesignation int, numAdditionalUse int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numDesignation >= len(resource.Concept[numConcept].Designation) || numAdditionalUse >= len(resource.Concept[numConcept].Designation[numDesignation].AdditionalUse) {
		return CodingSelect("Concept["+strconv.Itoa(numConcept)+"]Designation["+strconv.Itoa(numDesignation)+"].AdditionalUse["+strconv.Itoa(numAdditionalUse)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Concept["+strconv.Itoa(numConcept)+"]Designation["+strconv.Itoa(numDesignation)+"].AdditionalUse["+strconv.Itoa(numAdditionalUse)+"]", &resource.Concept[numConcept].Designation[numDesignation].AdditionalUse[numAdditionalUse], optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptDesignationValue(numConcept int, numDesignation int, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numDesignation >= len(resource.Concept[numConcept].Designation) {
		return StringInput("Concept["+strconv.Itoa(numConcept)+"]Designation["+strconv.Itoa(numDesignation)+"].Value", nil, htmlAttrs)
	}
	return StringInput("Concept["+strconv.Itoa(numConcept)+"]Designation["+strconv.Itoa(numDesignation)+"].Value", &resource.Concept[numConcept].Designation[numDesignation].Value, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyCode(numConcept int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return CodeSelect("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].Code", &resource.Concept[numConcept].Property[numProperty].Code, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueCode(numConcept int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return CodeSelect("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueCode", &resource.Concept[numConcept].Property[numProperty].ValueCode, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueCoding(numConcept int, numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return CodingSelect("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueCoding", &resource.Concept[numConcept].Property[numProperty].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueString(numConcept int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return StringInput("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueString", &resource.Concept[numConcept].Property[numProperty].ValueString, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueInteger(numConcept int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return IntInput("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueInteger", &resource.Concept[numConcept].Property[numProperty].ValueInteger, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueBoolean(numConcept int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return BoolInput("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueBoolean", &resource.Concept[numConcept].Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueDateTime(numConcept int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return DateTimeInput("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueDateTime", &resource.Concept[numConcept].Property[numProperty].ValueDateTime, htmlAttrs)
}
func (resource *CodeSystem) T_ConceptPropertyValueDecimal(numConcept int, numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numConcept >= len(resource.Concept) || numProperty >= len(resource.Concept[numConcept].Property) {
		return Float64Input("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("Concept["+strconv.Itoa(numConcept)+"]Property["+strconv.Itoa(numProperty)+"].ValueDecimal", &resource.Concept[numConcept].Property[numProperty].ValueDecimal, htmlAttrs)
}
