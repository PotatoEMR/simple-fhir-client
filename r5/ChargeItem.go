package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItem
type ChargeItem struct {
	Id                     *string               `json:"id,omitempty"`
	Meta                   *Meta                 `json:"meta,omitempty"`
	ImplicitRules          *string               `json:"implicitRules,omitempty"`
	Language               *string               `json:"language,omitempty"`
	Text                   *Narrative            `json:"text,omitempty"`
	Contained              []Resource            `json:"contained,omitempty"`
	Extension              []Extension           `json:"extension,omitempty"`
	ModifierExtension      []Extension           `json:"modifierExtension,omitempty"`
	Identifier             []Identifier          `json:"identifier,omitempty"`
	DefinitionUri          []string              `json:"definitionUri,omitempty"`
	DefinitionCanonical    []string              `json:"definitionCanonical,omitempty"`
	Status                 string                `json:"status"`
	PartOf                 []Reference           `json:"partOf,omitempty"`
	Code                   CodeableConcept       `json:"code"`
	Subject                Reference             `json:"subject"`
	Encounter              *Reference            `json:"encounter,omitempty"`
	OccurrenceDateTime     *string               `json:"occurrenceDateTime"`
	OccurrencePeriod       *Period               `json:"occurrencePeriod"`
	OccurrenceTiming       *Timing               `json:"occurrenceTiming"`
	Performer              []ChargeItemPerformer `json:"performer,omitempty"`
	PerformingOrganization *Reference            `json:"performingOrganization,omitempty"`
	RequestingOrganization *Reference            `json:"requestingOrganization,omitempty"`
	CostCenter             *Reference            `json:"costCenter,omitempty"`
	Quantity               *Quantity             `json:"quantity,omitempty"`
	Bodysite               []CodeableConcept     `json:"bodysite,omitempty"`
	UnitPriceComponent     *MonetaryComponent    `json:"unitPriceComponent,omitempty"`
	TotalPriceComponent    *MonetaryComponent    `json:"totalPriceComponent,omitempty"`
	OverrideReason         *CodeableConcept      `json:"overrideReason,omitempty"`
	Enterer                *Reference            `json:"enterer,omitempty"`
	EnteredDate            *string               `json:"enteredDate,omitempty"`
	Reason                 []CodeableConcept     `json:"reason,omitempty"`
	Service                []CodeableReference   `json:"service,omitempty"`
	Product                []CodeableReference   `json:"product,omitempty"`
	Account                []Reference           `json:"account,omitempty"`
	Note                   []Annotation          `json:"note,omitempty"`
	SupportingInformation  []Reference           `json:"supportingInformation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItem
type ChargeItemPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

type OtherChargeItem ChargeItem

// on convert struct to json, automatically add resourceType=ChargeItem
func (r ChargeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItem
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItem: OtherChargeItem(r),
		ResourceType:    "ChargeItem",
	})
}

func (resource *ChargeItem) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ChargeItem) T_Status() templ.Component {
	optionsValueSet := VSChargeitem_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ChargeItem) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet)
}
func (resource *ChargeItem) T_Bodysite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("bodysite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodysite", &resource.Bodysite[0], optionsValueSet)
}
func (resource *ChargeItem) T_OverrideReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("overrideReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("overrideReason", resource.OverrideReason, optionsValueSet)
}
func (resource *ChargeItem) T_Reason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", &resource.Reason[0], optionsValueSet)
}
func (resource *ChargeItem) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Performer[numPerformer].Function, optionsValueSet)
}
