//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/SearchParameter
type SearchParameter struct {
	Id                *string                    `json:"id,omitempty"`
	Meta              *Meta                      `json:"meta,omitempty"`
	ImplicitRules     *string                    `json:"implicitRules,omitempty"`
	Language          *string                    `json:"language,omitempty"`
	Text              *Narrative                 `json:"text,omitempty"`
	Contained         []Resource                 `json:"contained,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Url               string                     `json:"url"`
	Version           *string                    `json:"version,omitempty"`
	Name              string                     `json:"name"`
	DerivedFrom       *string                    `json:"derivedFrom,omitempty"`
	Status            string                     `json:"status"`
	Experimental      *bool                      `json:"experimental,omitempty"`
	Date              *string                    `json:"date,omitempty"`
	Publisher         *string                    `json:"publisher,omitempty"`
	Contact           []ContactDetail            `json:"contact,omitempty"`
	Description       string                     `json:"description"`
	UseContext        []UsageContext             `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept          `json:"jurisdiction,omitempty"`
	Purpose           *string                    `json:"purpose,omitempty"`
	Code              string                     `json:"code"`
	Base              []string                   `json:"base"`
	Type              string                     `json:"type"`
	Expression        *string                    `json:"expression,omitempty"`
	Xpath             *string                    `json:"xpath,omitempty"`
	XpathUsage        *string                    `json:"xpathUsage,omitempty"`
	Target            []string                   `json:"target,omitempty"`
	MultipleOr        *bool                      `json:"multipleOr,omitempty"`
	MultipleAnd       *bool                      `json:"multipleAnd,omitempty"`
	Comparator        []string                   `json:"comparator,omitempty"`
	Modifier          []string                   `json:"modifier,omitempty"`
	Chain             []string                   `json:"chain,omitempty"`
	Component         []SearchParameterComponent `json:"component,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SearchParameter
type SearchParameterComponent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Definition        string      `json:"definition"`
	Expression        string      `json:"expression"`
}

type OtherSearchParameter SearchParameter

// on convert struct to json, automatically add resourceType=SearchParameter
func (r SearchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSearchParameter
		ResourceType string `json:"resourceType"`
	}{
		OtherSearchParameter: OtherSearchParameter(r),
		ResourceType:         "SearchParameter",
	})
}
