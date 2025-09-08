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

// on convert struct to json, automatically add resourceType=SubstancePolymer
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
func (resource *SubstancePolymer) T_Class(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Class", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Class", resource.Class, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_Geometry(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Geometry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Geometry", resource.Geometry, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_CopolymerConnectivity(numCopolymerConnectivity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCopolymerConnectivity >= len(resource.CopolymerConnectivity) {
		return CodeableConceptSelect("CopolymerConnectivity["+strconv.Itoa(numCopolymerConnectivity)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CopolymerConnectivity["+strconv.Itoa(numCopolymerConnectivity)+"]", &resource.CopolymerConnectivity[numCopolymerConnectivity], optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_Modification(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Modification", nil, htmlAttrs)
	}
	return StringInput("Modification", resource.Modification, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetRatioType(numMonomerSet int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) {
		return CodeableConceptSelect("MonomerSet["+strconv.Itoa(numMonomerSet)+"]RatioType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MonomerSet["+strconv.Itoa(numMonomerSet)+"]RatioType", resource.MonomerSet[numMonomerSet].RatioType, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialCode(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return CodeableConceptSelect("MonomerSet["+strconv.Itoa(numMonomerSet)+"]StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MonomerSet["+strconv.Itoa(numMonomerSet)+"]StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Code", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Code, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialCategory(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return CodeableConceptSelect("MonomerSet["+strconv.Itoa(numMonomerSet)+"]StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MonomerSet["+strconv.Itoa(numMonomerSet)+"]StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Category", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Category, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialIsDefining(numMonomerSet int, numStartingMaterial int, htmlAttrs string) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return BoolInput("MonomerSet["+strconv.Itoa(numMonomerSet)+"]StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].IsDefining", nil, htmlAttrs)
	}
	return BoolInput("MonomerSet["+strconv.Itoa(numMonomerSet)+"]StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].IsDefining", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].IsDefining, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatAverageMolecularFormula(numRepeat int, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) {
		return StringInput("Repeat["+strconv.Itoa(numRepeat)+"]AverageMolecularFormula", nil, htmlAttrs)
	}
	return StringInput("Repeat["+strconv.Itoa(numRepeat)+"]AverageMolecularFormula", resource.Repeat[numRepeat].AverageMolecularFormula, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitAmountType(numRepeat int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) {
		return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnitAmountType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnitAmountType", resource.Repeat[numRepeat].RepeatUnitAmountType, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitUnit(numRepeat int, numRepeatUnit int, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) {
		return StringInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].Unit", nil, htmlAttrs)
	}
	return StringInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].Unit", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].Unit, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitOrientation(numRepeat int, numRepeatUnit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) {
		return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].Orientation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].Orientation", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].Orientation, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitAmount(numRepeat int, numRepeatUnit int, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) {
		return IntInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].Amount", nil, htmlAttrs)
	}
	return IntInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].Amount", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].Amount, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationType(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numDegreeOfPolymerisation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) {
		return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Type", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationAverage(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numDegreeOfPolymerisation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) {
		return IntInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Average", nil, htmlAttrs)
	}
	return IntInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Average", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Average, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationLow(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numDegreeOfPolymerisation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) {
		return IntInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Low", nil, htmlAttrs)
	}
	return IntInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Low", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Low, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationHigh(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numDegreeOfPolymerisation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) {
		return IntInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].High", nil, htmlAttrs)
	}
	return IntInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].High", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].High, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationType(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numStructuralRepresentation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) {
		return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Type", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationRepresentation(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numStructuralRepresentation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) {
		return StringInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Representation", nil, htmlAttrs)
	}
	return StringInput("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Representation", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Representation, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationFormat(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numStructuralRepresentation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) {
		return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Format", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Repeat["+strconv.Itoa(numRepeat)+"]RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Format", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Format, optionsValueSet, htmlAttrs)
}
