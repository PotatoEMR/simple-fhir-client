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

// http://hl7.org/fhir/r5/StructureDefinition/NutritionIntake
type NutritionIntake struct {
	Id                    *string                          `json:"id,omitempty"`
	Meta                  *Meta                            `json:"meta,omitempty"`
	ImplicitRules         *string                          `json:"implicitRules,omitempty"`
	Language              *string                          `json:"language,omitempty"`
	Text                  *Narrative                       `json:"text,omitempty"`
	Contained             []Resource                       `json:"contained,omitempty"`
	Extension             []Extension                      `json:"extension,omitempty"`
	ModifierExtension     []Extension                      `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                     `json:"identifier,omitempty"`
	InstantiatesCanonical []string                         `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                         `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference                      `json:"basedOn,omitempty"`
	PartOf                []Reference                      `json:"partOf,omitempty"`
	Status                string                           `json:"status"`
	StatusReason          []CodeableConcept                `json:"statusReason,omitempty"`
	Code                  *CodeableConcept                 `json:"code,omitempty"`
	Subject               Reference                        `json:"subject"`
	Encounter             *Reference                       `json:"encounter,omitempty"`
	OccurrenceDateTime    *time.Time                       `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrencePeriod      *Period                          `json:"occurrencePeriod,omitempty"`
	Recorded              *time.Time                       `json:"recorded,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	ReportedBoolean       *bool                            `json:"reportedBoolean,omitempty"`
	ReportedReference     *Reference                       `json:"reportedReference,omitempty"`
	ConsumedItem          []NutritionIntakeConsumedItem    `json:"consumedItem"`
	IngredientLabel       []NutritionIntakeIngredientLabel `json:"ingredientLabel,omitempty"`
	Performer             []NutritionIntakePerformer       `json:"performer,omitempty"`
	Location              *Reference                       `json:"location,omitempty"`
	DerivedFrom           []Reference                      `json:"derivedFrom,omitempty"`
	Reason                []CodeableReference              `json:"reason,omitempty"`
	Note                  []Annotation                     `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionIntake
type NutritionIntakeConsumedItem struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	NutritionProduct  CodeableReference `json:"nutritionProduct"`
	Schedule          *Timing           `json:"schedule,omitempty"`
	Amount            *Quantity         `json:"amount,omitempty"`
	Rate              *Quantity         `json:"rate,omitempty"`
	NotConsumed       *bool             `json:"notConsumed,omitempty"`
	NotConsumedReason *CodeableConcept  `json:"notConsumedReason,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionIntake
type NutritionIntakeIngredientLabel struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Nutrient          CodeableReference `json:"nutrient"`
	Amount            Quantity          `json:"amount"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionIntake
type NutritionIntakePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

type OtherNutritionIntake NutritionIntake

// on convert struct to json, automatically add resourceType=NutritionIntake
func (r NutritionIntake) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionIntake
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionIntake: OtherNutritionIntake(r),
		ResourceType:         "NutritionIntake",
	})
}
func (r NutritionIntake) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "NutritionIntake/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "NutritionIntake"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *NutritionIntake) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *NutritionIntake) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *NutritionIntake) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NutritionIntake) T_StatusReason(numStatusReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStatusReason >= len(resource.StatusReason) {
		return CodeableConceptSelect("StatusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StatusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *NutritionIntake) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *NutritionIntake) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *NutritionIntake) T_Recorded(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Recorded", nil, htmlAttrs)
	}
	return DateTimeInput("Recorded", resource.Recorded, htmlAttrs)
}
func (resource *NutritionIntake) T_ReportedBoolean(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ReportedBoolean", nil, htmlAttrs)
	}
	return BoolInput("ReportedBoolean", resource.ReportedBoolean, htmlAttrs)
}
func (resource *NutritionIntake) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *NutritionIntake) T_ConsumedItemType(numConsumedItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConsumedItem >= len(resource.ConsumedItem) {
		return CodeableConceptSelect("ConsumedItem["+strconv.Itoa(numConsumedItem)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConsumedItem["+strconv.Itoa(numConsumedItem)+"]Type", &resource.ConsumedItem[numConsumedItem].Type, optionsValueSet, htmlAttrs)
}
func (resource *NutritionIntake) T_ConsumedItemNotConsumed(numConsumedItem int, htmlAttrs string) templ.Component {
	if resource == nil || numConsumedItem >= len(resource.ConsumedItem) {
		return BoolInput("ConsumedItem["+strconv.Itoa(numConsumedItem)+"]NotConsumed", nil, htmlAttrs)
	}
	return BoolInput("ConsumedItem["+strconv.Itoa(numConsumedItem)+"]NotConsumed", resource.ConsumedItem[numConsumedItem].NotConsumed, htmlAttrs)
}
func (resource *NutritionIntake) T_ConsumedItemNotConsumedReason(numConsumedItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConsumedItem >= len(resource.ConsumedItem) {
		return CodeableConceptSelect("ConsumedItem["+strconv.Itoa(numConsumedItem)+"]NotConsumedReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConsumedItem["+strconv.Itoa(numConsumedItem)+"]NotConsumedReason", resource.ConsumedItem[numConsumedItem].NotConsumedReason, optionsValueSet, htmlAttrs)
}
func (resource *NutritionIntake) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("Performer["+strconv.Itoa(numPerformer)+"]Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Performer["+strconv.Itoa(numPerformer)+"]Function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
