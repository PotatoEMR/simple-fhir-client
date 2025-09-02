package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *SubstancePolymer) SubstancePolymerLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerClass(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("class", nil, optionsValueSet)
	}
	return CodeableConceptSelect("class", resource.Class, optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerGeometry(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("geometry", nil, optionsValueSet)
	}
	return CodeableConceptSelect("geometry", resource.Geometry, optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerCopolymerConnectivity(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("copolymerConnectivity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("copolymerConnectivity", &resource.CopolymerConnectivity[0], optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerMonomerSetRatioType(numMonomerSet int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.MonomerSet) >= numMonomerSet {
		return CodeableConceptSelect("ratioType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ratioType", resource.MonomerSet[numMonomerSet].RatioType, optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerMonomerSetStartingMaterialMaterial(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.MonomerSet[numMonomerSet].StartingMaterial) >= numStartingMaterial {
		return CodeableConceptSelect("material", nil, optionsValueSet)
	}
	return CodeableConceptSelect("material", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Material, optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerMonomerSetStartingMaterialType(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.MonomerSet[numMonomerSet].StartingMaterial) >= numStartingMaterial {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Type, optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerRepeatRepeatUnitAmountType(numRepeat int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Repeat) >= numRepeat {
		return CodeableConceptSelect("repeatUnitAmountType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("repeatUnitAmountType", resource.Repeat[numRepeat].RepeatUnitAmountType, optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerRepeatRepeatUnitOrientationOfPolymerisation(numRepeat int, numRepeatUnit int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Repeat[numRepeat].RepeatUnit) >= numRepeatUnit {
		return CodeableConceptSelect("orientationOfPolymerisation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("orientationOfPolymerisation", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].OrientationOfPolymerisation, optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisationDegree(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) >= numDegreeOfPolymerisation {
		return CodeableConceptSelect("degree", nil, optionsValueSet)
	}
	return CodeableConceptSelect("degree", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Degree, optionsValueSet)
}
func (resource *SubstancePolymer) SubstancePolymerRepeatRepeatUnitStructuralRepresentationType(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) >= numStructuralRepresentation {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Type, optionsValueSet)
}
