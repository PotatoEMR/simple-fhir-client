package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/SubstancePolymer
type SubstancePolymer struct {
	Id                    *string                      `json:"id,omitempty"`
	Meta                  *Meta                        `json:"meta,omitempty"`
	ImplicitRules         *string                      `json:"implicitRules,omitempty"`
	Language              *string                      `json:"language,omitempty"`
	Text                  *Narrative                   `json:"text,omitempty"`
	Contained             []Resource                   `json:"contained,omitempty"`
	Extension             []Extension                  `json:"extension,omitempty"`
	ModifierExtension     []Extension                  `json:"modifierExtension,omitempty"`
	Identifier            *Identifier                  `json:"identifier,omitempty"`
	Class                 *CodeableConcept             `json:"class,omitempty"`
	Geometry              *CodeableConcept             `json:"geometry,omitempty"`
	CopolymerConnectivity []CodeableConcept            `json:"copolymerConnectivity,omitempty"`
	Modification          *string                      `json:"modification,omitempty"`
	MonomerSet            []SubstancePolymerMonomerSet `json:"monomerSet,omitempty"`
	Repeat                []SubstancePolymerRepeat     `json:"repeat,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstancePolymer
type SubstancePolymerMonomerSet struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	RatioType         *CodeableConcept                             `json:"ratioType,omitempty"`
	StartingMaterial  []SubstancePolymerMonomerSetStartingMaterial `json:"startingMaterial,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstancePolymer
type SubstancePolymerMonomerSetStartingMaterial struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Category          *CodeableConcept `json:"category,omitempty"`
	IsDefining        *bool            `json:"isDefining,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeat struct {
	Id                      *string                            `json:"id,omitempty"`
	Extension               []Extension                        `json:"extension,omitempty"`
	ModifierExtension       []Extension                        `json:"modifierExtension,omitempty"`
	AverageMolecularFormula *string                            `json:"averageMolecularFormula,omitempty"`
	RepeatUnitAmountType    *CodeableConcept                   `json:"repeatUnitAmountType,omitempty"`
	RepeatUnit              []SubstancePolymerRepeatRepeatUnit `json:"repeatUnit,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnit struct {
	Id                       *string                                                    `json:"id,omitempty"`
	Extension                []Extension                                                `json:"extension,omitempty"`
	ModifierExtension        []Extension                                                `json:"modifierExtension,omitempty"`
	Unit                     *string                                                    `json:"unit,omitempty"`
	Orientation              *CodeableConcept                                           `json:"orientation,omitempty"`
	Amount                   *int                                                       `json:"amount,omitempty"`
	DegreeOfPolymerisation   []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation   `json:"degreeOfPolymerisation,omitempty"`
	StructuralRepresentation []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `json:"structuralRepresentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Average           *int             `json:"average,omitempty"`
	Low               *int             `json:"low,omitempty"`
	High              *int             `json:"high,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
	Format            *CodeableConcept `json:"format,omitempty"`
	Attachment        *Attachment      `json:"attachment,omitempty"`
}

type OtherSubstancePolymer SubstancePolymer

// struct -> json, automatically add resourceType=Patient
func (r SubstancePolymer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstancePolymer
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstancePolymer: OtherSubstancePolymer(r),
		ResourceType:          "SubstancePolymer",
	})
}

func (r SubstancePolymer) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubstancePolymer/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "SubstancePolymer"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r SubstancePolymer) ResourceType() string {
	return "SubstancePolymer"
}

func (resource *SubstancePolymer) T_Class(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("class", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("class", resource.Class, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_Geometry(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("geometry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("geometry", resource.Geometry, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_CopolymerConnectivity(numCopolymerConnectivity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCopolymerConnectivity >= len(resource.CopolymerConnectivity) {
		return CodeableConceptSelect("copolymerConnectivity["+strconv.Itoa(numCopolymerConnectivity)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("copolymerConnectivity["+strconv.Itoa(numCopolymerConnectivity)+"]", &resource.CopolymerConnectivity[numCopolymerConnectivity], optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_Modification(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("modification", nil, htmlAttrs)
	}
	return StringInput("modification", resource.Modification, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetRatioType(numMonomerSet int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) {
		return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].ratioType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].ratioType", resource.MonomerSet[numMonomerSet].RatioType, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialCode(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].code", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Code, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialCategory(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].category", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Category, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialIsDefining(numMonomerSet int, numStartingMaterial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return BoolInput("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].isDefining", nil, htmlAttrs)
	}
	return BoolInput("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].isDefining", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].IsDefining, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialAmount(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return QuantityInput("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].amount", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].amount", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Amount, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatAverageMolecularFormula(numRepeat int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) {
		return StringInput("repeat["+strconv.Itoa(numRepeat)+"].averageMolecularFormula", nil, htmlAttrs)
	}
	return StringInput("repeat["+strconv.Itoa(numRepeat)+"].averageMolecularFormula", resource.Repeat[numRepeat].AverageMolecularFormula, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitAmountType(numRepeat int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) {
		return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnitAmountType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnitAmountType", resource.Repeat[numRepeat].RepeatUnitAmountType, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitUnit(numRepeat int, numRepeatUnit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) {
		return StringInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].unit", nil, htmlAttrs)
	}
	return StringInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].unit", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].Unit, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitOrientation(numRepeat int, numRepeatUnit int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) {
		return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].orientation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].orientation", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].Orientation, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitAmount(numRepeat int, numRepeatUnit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) {
		return IntInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].amount", nil, htmlAttrs)
	}
	return IntInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].amount", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].Amount, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationType(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numDegreeOfPolymerisation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) {
		return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].type", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationAverage(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numDegreeOfPolymerisation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) {
		return IntInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].average", nil, htmlAttrs)
	}
	return IntInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].average", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Average, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationLow(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numDegreeOfPolymerisation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) {
		return IntInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].low", nil, htmlAttrs)
	}
	return IntInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].low", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Low, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationHigh(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numDegreeOfPolymerisation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) {
		return IntInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].high", nil, htmlAttrs)
	}
	return IntInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].high", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].High, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationType(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numStructuralRepresentation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) {
		return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].type", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationRepresentation(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numStructuralRepresentation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) {
		return StringInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].representation", nil, htmlAttrs)
	}
	return StringInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].representation", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Representation, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationFormat(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numStructuralRepresentation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) {
		return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].format", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].format", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Format, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationAttachment(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numStructuralRepresentation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) {
		return AttachmentInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].attachment", nil, htmlAttrs)
	}
	return AttachmentInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].attachment", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Attachment, htmlAttrs)
}
