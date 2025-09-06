package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymer struct {
	Id                    *string                      `json:"id,omitempty"`
	Meta                  *Meta                        `json:"meta,omitempty"`
	ImplicitRules         *string                      `json:"implicitRules,omitempty"`
	Language              *string                      `json:"language,omitempty"`
	Text                  *Narrative                   `json:"text,omitempty"`
	Contained             []Resource                   `json:"contained,omitempty"`
	Extension             []Extension                  `json:"extension,omitempty"`
	ModifierExtension     []Extension                  `json:"modifierExtension,omitempty"`
	Class                 *CodeableConcept             `json:"class,omitempty"`
	Geometry              *CodeableConcept             `json:"geometry,omitempty"`
	CopolymerConnectivity []CodeableConcept            `json:"copolymerConnectivity,omitempty"`
	Modification          []string                     `json:"modification,omitempty"`
	MonomerSet            []SubstancePolymerMonomerSet `json:"monomerSet,omitempty"`
	Repeat                []SubstancePolymerRepeat     `json:"repeat,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerMonomerSet struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	RatioType         *CodeableConcept                             `json:"ratioType,omitempty"`
	StartingMaterial  []SubstancePolymerMonomerSetStartingMaterial `json:"startingMaterial,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerMonomerSetStartingMaterial struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Material          *CodeableConcept `json:"material,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	IsDefining        *bool            `json:"isDefining,omitempty"`
	Amount            *SubstanceAmount `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeat struct {
	Id                      *string                            `json:"id,omitempty"`
	Extension               []Extension                        `json:"extension,omitempty"`
	ModifierExtension       []Extension                        `json:"modifierExtension,omitempty"`
	NumberOfUnits           *int                               `json:"numberOfUnits,omitempty"`
	AverageMolecularFormula *string                            `json:"averageMolecularFormula,omitempty"`
	RepeatUnitAmountType    *CodeableConcept                   `json:"repeatUnitAmountType,omitempty"`
	RepeatUnit              []SubstancePolymerRepeatRepeatUnit `json:"repeatUnit,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnit struct {
	Id                          *string                                                    `json:"id,omitempty"`
	Extension                   []Extension                                                `json:"extension,omitempty"`
	ModifierExtension           []Extension                                                `json:"modifierExtension,omitempty"`
	OrientationOfPolymerisation *CodeableConcept                                           `json:"orientationOfPolymerisation,omitempty"`
	RepeatUnit                  *string                                                    `json:"repeatUnit,omitempty"`
	Amount                      *SubstanceAmount                                           `json:"amount,omitempty"`
	DegreeOfPolymerisation      []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation   `json:"degreeOfPolymerisation,omitempty"`
	StructuralRepresentation    []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `json:"structuralRepresentation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Degree            *CodeableConcept `json:"degree,omitempty"`
	Amount            *SubstanceAmount `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
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

func (resource *SubstancePolymer) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SubstancePolymer.Id", nil)
	}
	return StringInput("SubstancePolymer.Id", resource.Id)
}
func (resource *SubstancePolymer) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SubstancePolymer.ImplicitRules", nil)
	}
	return StringInput("SubstancePolymer.ImplicitRules", resource.ImplicitRules)
}
func (resource *SubstancePolymer) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SubstancePolymer.Language", nil, optionsValueSet)
	}
	return CodeSelect("SubstancePolymer.Language", resource.Language, optionsValueSet)
}
func (resource *SubstancePolymer) T_Class(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstancePolymer.Class", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.Class", resource.Class, optionsValueSet)
}
func (resource *SubstancePolymer) T_Geometry(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstancePolymer.Geometry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.Geometry", resource.Geometry, optionsValueSet)
}
func (resource *SubstancePolymer) T_CopolymerConnectivity(numCopolymerConnectivity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CopolymerConnectivity) >= numCopolymerConnectivity {
		return CodeableConceptSelect("SubstancePolymer.CopolymerConnectivity["+strconv.Itoa(numCopolymerConnectivity)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.CopolymerConnectivity["+strconv.Itoa(numCopolymerConnectivity)+"]", &resource.CopolymerConnectivity[numCopolymerConnectivity], optionsValueSet)
}
func (resource *SubstancePolymer) T_Modification(numModification int) templ.Component {

	if resource == nil || len(resource.Modification) >= numModification {
		return StringInput("SubstancePolymer.Modification["+strconv.Itoa(numModification)+"]", nil)
	}
	return StringInput("SubstancePolymer.Modification["+strconv.Itoa(numModification)+"]", &resource.Modification[numModification])
}
func (resource *SubstancePolymer) T_MonomerSetId(numMonomerSet int) templ.Component {

	if resource == nil || len(resource.MonomerSet) >= numMonomerSet {
		return StringInput("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].Id", nil)
	}
	return StringInput("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].Id", resource.MonomerSet[numMonomerSet].Id)
}
func (resource *SubstancePolymer) T_MonomerSetRatioType(numMonomerSet int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.MonomerSet) >= numMonomerSet {
		return CodeableConceptSelect("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].RatioType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].RatioType", resource.MonomerSet[numMonomerSet].RatioType, optionsValueSet)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialId(numMonomerSet int, numStartingMaterial int) templ.Component {

	if resource == nil || len(resource.MonomerSet) >= numMonomerSet || len(resource.MonomerSet[numMonomerSet].StartingMaterial) >= numStartingMaterial {
		return StringInput("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Id", nil)
	}
	return StringInput("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Id", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Id)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialMaterial(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.MonomerSet) >= numMonomerSet || len(resource.MonomerSet[numMonomerSet].StartingMaterial) >= numStartingMaterial {
		return CodeableConceptSelect("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Material", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Material", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Material, optionsValueSet)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialType(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.MonomerSet) >= numMonomerSet || len(resource.MonomerSet[numMonomerSet].StartingMaterial) >= numStartingMaterial {
		return CodeableConceptSelect("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].Type", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Type, optionsValueSet)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialIsDefining(numMonomerSet int, numStartingMaterial int) templ.Component {

	if resource == nil || len(resource.MonomerSet) >= numMonomerSet || len(resource.MonomerSet[numMonomerSet].StartingMaterial) >= numStartingMaterial {
		return BoolInput("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].IsDefining", nil)
	}
	return BoolInput("SubstancePolymer.MonomerSet["+strconv.Itoa(numMonomerSet)+"].StartingMaterial["+strconv.Itoa(numStartingMaterial)+"].IsDefining", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].IsDefining)
}
func (resource *SubstancePolymer) T_RepeatId(numRepeat int) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat {
		return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].Id", nil)
	}
	return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].Id", resource.Repeat[numRepeat].Id)
}
func (resource *SubstancePolymer) T_RepeatNumberOfUnits(numRepeat int) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat {
		return IntInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].NumberOfUnits", nil)
	}
	return IntInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].NumberOfUnits", resource.Repeat[numRepeat].NumberOfUnits)
}
func (resource *SubstancePolymer) T_RepeatAverageMolecularFormula(numRepeat int) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat {
		return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].AverageMolecularFormula", nil)
	}
	return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].AverageMolecularFormula", resource.Repeat[numRepeat].AverageMolecularFormula)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitAmountType(numRepeat int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat {
		return CodeableConceptSelect("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnitAmountType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnitAmountType", resource.Repeat[numRepeat].RepeatUnitAmountType, optionsValueSet)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitId(numRepeat int, numRepeatUnit int) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat || len(resource.Repeat[numRepeat].RepeatUnit) >= numRepeatUnit {
		return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].Id", nil)
	}
	return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].Id", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].Id)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitOrientationOfPolymerisation(numRepeat int, numRepeatUnit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat || len(resource.Repeat[numRepeat].RepeatUnit) >= numRepeatUnit {
		return CodeableConceptSelect("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].OrientationOfPolymerisation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].OrientationOfPolymerisation", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].OrientationOfPolymerisation, optionsValueSet)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitRepeatUnit(numRepeat int, numRepeatUnit int) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat || len(resource.Repeat[numRepeat].RepeatUnit) >= numRepeatUnit {
		return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].RepeatUnit", nil)
	}
	return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].RepeatUnit", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].RepeatUnit)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationId(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat || len(resource.Repeat[numRepeat].RepeatUnit) >= numRepeatUnit || len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) >= numDegreeOfPolymerisation {
		return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Id", nil)
	}
	return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Id", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Id)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationDegree(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat || len(resource.Repeat[numRepeat].RepeatUnit) >= numRepeatUnit || len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) >= numDegreeOfPolymerisation {
		return CodeableConceptSelect("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Degree", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].DegreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].Degree", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Degree, optionsValueSet)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationId(numRepeat int, numRepeatUnit int, numStructuralRepresentation int) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat || len(resource.Repeat[numRepeat].RepeatUnit) >= numRepeatUnit || len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) >= numStructuralRepresentation {
		return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Id", nil)
	}
	return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Id", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Id)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationType(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat || len(resource.Repeat[numRepeat].RepeatUnit) >= numRepeatUnit || len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) >= numStructuralRepresentation {
		return CodeableConceptSelect("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Type", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Type, optionsValueSet)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationRepresentation(numRepeat int, numRepeatUnit int, numStructuralRepresentation int) templ.Component {

	if resource == nil || len(resource.Repeat) >= numRepeat || len(resource.Repeat[numRepeat].RepeatUnit) >= numRepeatUnit || len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) >= numStructuralRepresentation {
		return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Representation", nil)
	}
	return StringInput("SubstancePolymer.Repeat["+strconv.Itoa(numRepeat)+"].RepeatUnit["+strconv.Itoa(numRepeatUnit)+"].StructuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].Representation", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Representation)
}
