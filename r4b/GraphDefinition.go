package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date              *string               `json:"date,omitempty"`
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
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *GraphDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *GraphDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *GraphDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *GraphDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *GraphDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *GraphDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *GraphDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *GraphDefinition) T_Start(htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil {
		return CodeSelect("start", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("start", &resource.Start, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Profile(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("profile", nil, htmlAttrs)
	}
	return StringInput("profile", resource.Profile, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkPath(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].path", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].path", resource.Link[numLink].Path, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkSliceName(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].sliceName", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].sliceName", resource.Link[numLink].SliceName, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkMin(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return IntInput("link["+strconv.Itoa(numLink)+"].min", nil, htmlAttrs)
	}
	return IntInput("link["+strconv.Itoa(numLink)+"].min", resource.Link[numLink].Min, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkMax(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].max", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].max", resource.Link[numLink].Max, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkDescription(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].description", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].description", resource.Link[numLink].Description, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetType(numLink int, numTarget int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].type", &resource.Link[numLink].Target[numTarget].Type, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetParams(numLink int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) {
		return StringInput("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].params", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].params", resource.Link[numLink].Target[numTarget].Params, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetProfile(numLink int, numTarget int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) {
		return StringInput("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].profile", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].profile", resource.Link[numLink].Target[numTarget].Profile, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentUse(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGraph_compartment_use

	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].use", &resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Use, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentCode(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].code", &resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Code, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentRule(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGraph_compartment_rule

	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].rule", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].rule", &resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Rule, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentExpression(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return StringInput("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].expression", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].expression", resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Expression, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentDescription(numLink int, numTarget int, numCompartment int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numTarget >= len(resource.Link[numLink].Target) || numCompartment >= len(resource.Link[numLink].Target[numTarget].Compartment) {
		return StringInput("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].description", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].target["+strconv.Itoa(numTarget)+"].compartment["+strconv.Itoa(numCompartment)+"].description", resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Description, htmlAttrs)
}
