package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/GraphDefinition
type GraphDefinition struct {
	Id                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []Resource            `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Url               *string               `json:"url,omitempty"`
	Version           *string               `json:"version,omitempty"`
	Name              string                `json:"name"`
	Status            string                `json:"status"`
	Experimental      *bool                 `json:"experimental,omitempty"`
	Date              *time.Time            `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher         *string               `json:"publisher,omitempty"`
	Contact           []ContactDetail       `json:"contact,omitempty"`
	Description       *string               `json:"description,omitempty"`
	UseContext        []UsageContext        `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept     `json:"jurisdiction,omitempty"`
	Purpose           *string               `json:"purpose,omitempty"`
	Start             string                `json:"start"`
	Profile           *string               `json:"profile,omitempty"`
	Link              []GraphDefinitionLink `json:"link,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/GraphDefinition
type GraphDefinitionLink struct {
	Id                *string                     `json:"id,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Path              *string                     `json:"path,omitempty"`
	SliceName         *string                     `json:"sliceName,omitempty"`
	Min               *int                        `json:"min,omitempty"`
	Max               *string                     `json:"max,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	Target            []GraphDefinitionLinkTarget `json:"target,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/GraphDefinition
type GraphDefinitionLinkTarget struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Type              string                                 `json:"type"`
	Params            *string                                `json:"params,omitempty"`
	Profile           *string                                `json:"profile,omitempty"`
	Compartment       []GraphDefinitionLinkTargetCompartment `json:"compartment,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/GraphDefinition
type GraphDefinitionLinkTargetCompartment struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Use               string      `json:"use"`
	Code              string      `json:"code"`
	Rule              string      `json:"rule"`
	Expression        *string     `json:"expression,omitempty"`
	Description       *string     `json:"description,omitempty"`
}

type OtherGraphDefinition GraphDefinition

// on convert struct to json, automatically add resourceType=GraphDefinition
func (r GraphDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGraphDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherGraphDefinition: OtherGraphDefinition(r),
		ResourceType:         "GraphDefinition",
	})
}
func (r GraphDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "GraphDefinition/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "GraphDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *GraphDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("GraphDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *GraphDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("GraphDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *GraphDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("GraphDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Name", &resource.Name, htmlAttrs)
}
func (resource *GraphDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("GraphDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("GraphDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("GraphDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *GraphDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("GraphDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("GraphDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *GraphDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("GraphDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *GraphDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("GraphDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *GraphDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("GraphDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("GraphDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("GraphDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *GraphDefinition) T_Start(htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil {
		return CodeSelect("GraphDefinition.Start", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Start", &resource.Start, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Profile(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("GraphDefinition.Profile", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Profile", resource.Profile, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkPath(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Path", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Path", resource.Link[numLink].Path, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkSliceName(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].SliceName", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].SliceName", resource.Link[numLink].SliceName, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkMin(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return IntInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Min", nil, htmlAttrs)
	}
	return IntInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Min", resource.Link[numLink].Min, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkMax(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Max", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Max", resource.Link[numLink].Max, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkDescription(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Description", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Description", resource.Link[numLink].Description, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetType(numLink int, numTarget int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Type", &resource.Link[numLink].Target[numTarget].Type, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetParams(numLink int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Params", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Params", resource.Link[numLink].Target[numTarget].Params, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetProfile(numLink int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Profile", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Profile", resource.Link[numLink].Target[numTarget].Profile, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentUse(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGraph_compartment_use

	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Use", &resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Use, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentCode(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Code", &resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Code, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentRule(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGraph_compartment_rule

	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Rule", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Rule", &resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Rule, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentExpression(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Expression", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Expression", resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Expression, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentDescription(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Description", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Description", resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Description, htmlAttrs)
}
