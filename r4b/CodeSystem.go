package r4b

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/CodeSystem
type CodeSystem struct {
	Id                *string              `json:"id,omitempty"`
	Meta              *Meta                `json:"meta,omitempty"`
	ImplicitRules     *string              `json:"implicitRules,omitempty"`
	Language          *string              `json:"language,omitempty"`
	Text              *Narrative           `json:"text,omitempty"`
	Contained         []Resource           `json:"contained,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Url               *string              `json:"url,omitempty"`
	Identifier        []Identifier         `json:"identifier,omitempty"`
	Version           *string              `json:"version,omitempty"`
	Name              *string              `json:"name,omitempty"`
	Title             *string              `json:"title,omitempty"`
	Status            string               `json:"status"`
	Experimental      *bool                `json:"experimental,omitempty"`
	Date              *string              `json:"date,omitempty"`
	Publisher         *string              `json:"publisher,omitempty"`
	Contact           []ContactDetail      `json:"contact,omitempty"`
	Description       *string              `json:"description,omitempty"`
	UseContext        []UsageContext       `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept    `json:"jurisdiction,omitempty"`
	Purpose           *string              `json:"purpose,omitempty"`
	Copyright         *string              `json:"copyright,omitempty"`
	CaseSensitive     *bool                `json:"caseSensitive,omitempty"`
	ValueSet          *string              `json:"valueSet,omitempty"`
	HierarchyMeaning  *string              `json:"hierarchyMeaning,omitempty"`
	Compositional     *bool                `json:"compositional,omitempty"`
	VersionNeeded     *bool                `json:"versionNeeded,omitempty"`
	Content           string               `json:"content"`
	Supplements       *string              `json:"supplements,omitempty"`
	Count             *int                 `json:"count,omitempty"`
	Filter            []CodeSystemFilter   `json:"filter,omitempty"`
	Property          []CodeSystemProperty `json:"property,omitempty"`
	Concept           []CodeSystemConcept  `json:"concept,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CodeSystem
type CodeSystemFilter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Description       *string     `json:"description,omitempty"`
	Operator          []string    `json:"operator"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CodeSystem
type CodeSystemProperty struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Uri               *string     `json:"uri,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CodeSystem
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

// http://hl7.org/fhir/r4b/StructureDefinition/CodeSystem
type CodeSystemConceptDesignation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CodeSystem
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
	ValueDateTime     string      `json:"valueDateTime"`
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

func (resource *CodeSystem) CodeSystemLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemHierarchyMeaning() templ.Component {
	optionsValueSet := VSCodesystem_hierarchy_meaning
	currentVal := ""
	if resource != nil {
		currentVal = *resource.HierarchyMeaning
	}
	return CodeSelect("hierarchyMeaning", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemContent() templ.Component {
	optionsValueSet := VSCodesystem_content_mode
	currentVal := ""
	if resource != nil {
		currentVal = resource.Content
	}
	return CodeSelect("content", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemFilterCode(numFilter int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Filter) >= numFilter {
		currentVal = resource.Filter[numFilter].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemFilterOperator(numFilter int) templ.Component {
	optionsValueSet := VSFilter_operator
	currentVal := ""
	if resource != nil && len(resource.Filter) >= numFilter {
		currentVal = resource.Filter[numFilter].Operator[0]
	}
	return CodeSelect("operator", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemPropertyCode(numProperty int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Property) >= numProperty {
		currentVal = resource.Property[numProperty].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemPropertyType(numProperty int) templ.Component {
	optionsValueSet := VSConcept_property_type
	currentVal := ""
	if resource != nil && len(resource.Property) >= numProperty {
		currentVal = resource.Property[numProperty].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemConceptCode(numConcept int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Concept) >= numConcept {
		currentVal = resource.Concept[numConcept].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemConceptDesignationLanguage(numConcept int, numDesignation int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Concept[numConcept].Designation) >= numDesignation {
		currentVal = *resource.Concept[numConcept].Designation[numDesignation].Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *CodeSystem) CodeSystemConceptPropertyCode(numConcept int, numProperty int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Concept[numConcept].Property) >= numProperty {
		currentVal = resource.Concept[numConcept].Property[numProperty].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
