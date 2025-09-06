package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type Citation struct {
	Id                *string                  `json:"id,omitempty"`
	Meta              *Meta                    `json:"meta,omitempty"`
	ImplicitRules     *string                  `json:"implicitRules,omitempty"`
	Language          *string                  `json:"language,omitempty"`
	Text              *Narrative               `json:"text,omitempty"`
	Contained         []Resource               `json:"contained,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Url               *string                  `json:"url,omitempty"`
	Identifier        []Identifier             `json:"identifier,omitempty"`
	Version           *string                  `json:"version,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	Title             *string                  `json:"title,omitempty"`
	Status            string                   `json:"status"`
	Experimental      *bool                    `json:"experimental,omitempty"`
	Date              *string                  `json:"date,omitempty"`
	Publisher         *string                  `json:"publisher,omitempty"`
	Contact           []ContactDetail          `json:"contact,omitempty"`
	Description       *string                  `json:"description,omitempty"`
	UseContext        []UsageContext           `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept        `json:"jurisdiction,omitempty"`
	Purpose           *string                  `json:"purpose,omitempty"`
	Copyright         *string                  `json:"copyright,omitempty"`
	ApprovalDate      *string                  `json:"approvalDate,omitempty"`
	LastReviewDate    *string                  `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                  `json:"effectivePeriod,omitempty"`
	Author            []ContactDetail          `json:"author,omitempty"`
	Editor            []ContactDetail          `json:"editor,omitempty"`
	Reviewer          []ContactDetail          `json:"reviewer,omitempty"`
	Endorser          []ContactDetail          `json:"endorser,omitempty"`
	Summary           []CitationSummary        `json:"summary,omitempty"`
	Classification    []CitationClassification `json:"classification,omitempty"`
	Note              []Annotation             `json:"note,omitempty"`
	CurrentState      []CodeableConcept        `json:"currentState,omitempty"`
	StatusDate        []CitationStatusDate     `json:"statusDate,omitempty"`
	RelatesTo         []CitationRelatesTo      `json:"relatesTo,omitempty"`
	CitedArtifact     *CitationCitedArtifact   `json:"citedArtifact,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationSummary struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Style             *CodeableConcept `json:"style,omitempty"`
	Text              string           `json:"text"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Classifier        []CodeableConcept `json:"classifier,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationStatusDate struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Activity          CodeableConcept `json:"activity"`
	Actual            *bool           `json:"actual,omitempty"`
	Period            Period          `json:"period"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationRelatesTo struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	RelationshipType  CodeableConcept   `json:"relationshipType"`
	TargetClassifier  []CodeableConcept `json:"targetClassifier,omitempty"`
	TargetUri         string            `json:"targetUri"`
	TargetIdentifier  Identifier        `json:"targetIdentifier"`
	TargetReference   Reference         `json:"targetReference"`
	TargetAttachment  Attachment        `json:"targetAttachment"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifact struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                           `json:"identifier,omitempty"`
	RelatedIdentifier []Identifier                           `json:"relatedIdentifier,omitempty"`
	DateAccessed      *string                                `json:"dateAccessed,omitempty"`
	Version           *CitationCitedArtifactVersion          `json:"version,omitempty"`
	CurrentState      []CodeableConcept                      `json:"currentState,omitempty"`
	StatusDate        []CitationCitedArtifactStatusDate      `json:"statusDate,omitempty"`
	Title             []CitationCitedArtifactTitle           `json:"title,omitempty"`
	Abstract          []CitationCitedArtifactAbstract        `json:"abstract,omitempty"`
	Part              *CitationCitedArtifactPart             `json:"part,omitempty"`
	RelatesTo         []CitationCitedArtifactRelatesTo       `json:"relatesTo,omitempty"`
	PublicationForm   []CitationCitedArtifactPublicationForm `json:"publicationForm,omitempty"`
	WebLocation       []CitationCitedArtifactWebLocation     `json:"webLocation,omitempty"`
	Classification    []CitationCitedArtifactClassification  `json:"classification,omitempty"`
	Contributorship   *CitationCitedArtifactContributorship  `json:"contributorship,omitempty"`
	Note              []Annotation                           `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactVersion struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Value             string      `json:"value"`
	BaseCitation      *Reference  `json:"baseCitation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactStatusDate struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Activity          CodeableConcept `json:"activity"`
	Actual            *bool           `json:"actual,omitempty"`
	Period            Period          `json:"period"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactTitle struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	Language          *CodeableConcept  `json:"language,omitempty"`
	Text              string            `json:"text"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactAbstract struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Language          *CodeableConcept `json:"language,omitempty"`
	Text              string           `json:"text"`
	Copyright         *string          `json:"copyright,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactPart struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Value             *string          `json:"value,omitempty"`
	BaseCitation      *Reference       `json:"baseCitation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactRelatesTo struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	RelationshipType  CodeableConcept   `json:"relationshipType"`
	TargetClassifier  []CodeableConcept `json:"targetClassifier,omitempty"`
	TargetUri         string            `json:"targetUri"`
	TargetIdentifier  Identifier        `json:"targetIdentifier"`
	TargetReference   Reference         `json:"targetReference"`
	TargetAttachment  Attachment        `json:"targetAttachment"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactPublicationForm struct {
	Id                *string                                              `json:"id,omitempty"`
	Extension         []Extension                                          `json:"extension,omitempty"`
	ModifierExtension []Extension                                          `json:"modifierExtension,omitempty"`
	PublishedIn       *CitationCitedArtifactPublicationFormPublishedIn     `json:"publishedIn,omitempty"`
	PeriodicRelease   *CitationCitedArtifactPublicationFormPeriodicRelease `json:"periodicRelease,omitempty"`
	ArticleDate       *string                                              `json:"articleDate,omitempty"`
	LastRevisionDate  *string                                              `json:"lastRevisionDate,omitempty"`
	Language          []CodeableConcept                                    `json:"language,omitempty"`
	AccessionNumber   *string                                              `json:"accessionNumber,omitempty"`
	PageString        *string                                              `json:"pageString,omitempty"`
	FirstPage         *string                                              `json:"firstPage,omitempty"`
	LastPage          *string                                              `json:"lastPage,omitempty"`
	PageCount         *string                                              `json:"pageCount,omitempty"`
	Copyright         *string                                              `json:"copyright,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactPublicationFormPublishedIn struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Identifier        []Identifier     `json:"identifier,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Publisher         *Reference       `json:"publisher,omitempty"`
	PublisherLocation *string          `json:"publisherLocation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactPublicationFormPeriodicRelease struct {
	Id                *string                                                               `json:"id,omitempty"`
	Extension         []Extension                                                           `json:"extension,omitempty"`
	ModifierExtension []Extension                                                           `json:"modifierExtension,omitempty"`
	CitedMedium       *CodeableConcept                                                      `json:"citedMedium,omitempty"`
	Volume            *string                                                               `json:"volume,omitempty"`
	Issue             *string                                                               `json:"issue,omitempty"`
	DateOfPublication *CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication `json:"dateOfPublication,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Date              *string     `json:"date,omitempty"`
	Year              *string     `json:"year,omitempty"`
	Month             *string     `json:"month,omitempty"`
	Day               *string     `json:"day,omitempty"`
	Season            *string     `json:"season,omitempty"`
	Text              *string     `json:"text,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactWebLocation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Url               *string          `json:"url,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactClassification struct {
	Id                *string                                           `json:"id,omitempty"`
	Extension         []Extension                                       `json:"extension,omitempty"`
	ModifierExtension []Extension                                       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept                                  `json:"type,omitempty"`
	Classifier        []CodeableConcept                                 `json:"classifier,omitempty"`
	WhoClassified     *CitationCitedArtifactClassificationWhoClassified `json:"whoClassified,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactClassificationWhoClassified struct {
	Id                  *string     `json:"id,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Person              *Reference  `json:"person,omitempty"`
	Organization        *Reference  `json:"organization,omitempty"`
	Publisher           *Reference  `json:"publisher,omitempty"`
	ClassifierCopyright *string     `json:"classifierCopyright,omitempty"`
	FreeToShare         *bool       `json:"freeToShare,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactContributorship struct {
	Id                *string                                       `json:"id,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Complete          *bool                                         `json:"complete,omitempty"`
	Entry             []CitationCitedArtifactContributorshipEntry   `json:"entry,omitempty"`
	Summary           []CitationCitedArtifactContributorshipSummary `json:"summary,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactContributorshipEntry struct {
	Id                   *string                                                         `json:"id,omitempty"`
	Extension            []Extension                                                     `json:"extension,omitempty"`
	ModifierExtension    []Extension                                                     `json:"modifierExtension,omitempty"`
	Name                 *HumanName                                                      `json:"name,omitempty"`
	Initials             *string                                                         `json:"initials,omitempty"`
	CollectiveName       *string                                                         `json:"collectiveName,omitempty"`
	Identifier           []Identifier                                                    `json:"identifier,omitempty"`
	AffiliationInfo      []CitationCitedArtifactContributorshipEntryAffiliationInfo      `json:"affiliationInfo,omitempty"`
	Address              []Address                                                       `json:"address,omitempty"`
	Telecom              []ContactPoint                                                  `json:"telecom,omitempty"`
	ContributionType     []CodeableConcept                                               `json:"contributionType,omitempty"`
	Role                 *CodeableConcept                                                `json:"role,omitempty"`
	ContributionInstance []CitationCitedArtifactContributorshipEntryContributionInstance `json:"contributionInstance,omitempty"`
	CorrespondingContact *bool                                                           `json:"correspondingContact,omitempty"`
	ListOrder            *int                                                            `json:"listOrder,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactContributorshipEntryAffiliationInfo struct {
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Affiliation       *string      `json:"affiliation,omitempty"`
	Role              *string      `json:"role,omitempty"`
	Identifier        []Identifier `json:"identifier,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactContributorshipEntryContributionInstance struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Time              *string         `json:"time,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Citation
type CitationCitedArtifactContributorshipSummary struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Style             *CodeableConcept `json:"style,omitempty"`
	Source            *CodeableConcept `json:"source,omitempty"`
	Value             string           `json:"value"`
}

type OtherCitation Citation

// on convert struct to json, automatically add resourceType=Citation
func (r Citation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCitation
		ResourceType string `json:"resourceType"`
	}{
		OtherCitation: OtherCitation(r),
		ResourceType:  "Citation",
	})
}

func (resource *Citation) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Citation.Id", nil)
	}
	return StringInput("Citation.Id", resource.Id)
}
func (resource *Citation) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Citation.ImplicitRules", nil)
	}
	return StringInput("Citation.ImplicitRules", resource.ImplicitRules)
}
func (resource *Citation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Citation.Language", nil, optionsValueSet)
	}
	return CodeSelect("Citation.Language", resource.Language, optionsValueSet)
}
func (resource *Citation) T_Url() templ.Component {

	if resource == nil {
		return StringInput("Citation.Url", nil)
	}
	return StringInput("Citation.Url", resource.Url)
}
func (resource *Citation) T_Version() templ.Component {

	if resource == nil {
		return StringInput("Citation.Version", nil)
	}
	return StringInput("Citation.Version", resource.Version)
}
func (resource *Citation) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Citation.Name", nil)
	}
	return StringInput("Citation.Name", resource.Name)
}
func (resource *Citation) T_Title() templ.Component {

	if resource == nil {
		return StringInput("Citation.Title", nil)
	}
	return StringInput("Citation.Title", resource.Title)
}
func (resource *Citation) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Citation.Status", nil, optionsValueSet)
	}
	return CodeSelect("Citation.Status", &resource.Status, optionsValueSet)
}
func (resource *Citation) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("Citation.Experimental", nil)
	}
	return BoolInput("Citation.Experimental", resource.Experimental)
}
func (resource *Citation) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Citation.Date", nil)
	}
	return StringInput("Citation.Date", resource.Date)
}
func (resource *Citation) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("Citation.Publisher", nil)
	}
	return StringInput("Citation.Publisher", resource.Publisher)
}
func (resource *Citation) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Citation.Description", nil)
	}
	return StringInput("Citation.Description", resource.Description)
}
func (resource *Citation) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("Citation.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *Citation) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("Citation.Purpose", nil)
	}
	return StringInput("Citation.Purpose", resource.Purpose)
}
func (resource *Citation) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("Citation.Copyright", nil)
	}
	return StringInput("Citation.Copyright", resource.Copyright)
}
func (resource *Citation) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("Citation.ApprovalDate", nil)
	}
	return StringInput("Citation.ApprovalDate", resource.ApprovalDate)
}
func (resource *Citation) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("Citation.LastReviewDate", nil)
	}
	return StringInput("Citation.LastReviewDate", resource.LastReviewDate)
}
func (resource *Citation) T_CurrentState(numCurrentState int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CurrentState) >= numCurrentState {
		return CodeableConceptSelect("Citation.CurrentState["+strconv.Itoa(numCurrentState)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CurrentState["+strconv.Itoa(numCurrentState)+"]", &resource.CurrentState[numCurrentState], optionsValueSet)
}
func (resource *Citation) T_SummaryId(numSummary int) templ.Component {

	if resource == nil || len(resource.Summary) >= numSummary {
		return StringInput("Citation.Summary["+strconv.Itoa(numSummary)+"].Id", nil)
	}
	return StringInput("Citation.Summary["+strconv.Itoa(numSummary)+"].Id", resource.Summary[numSummary].Id)
}
func (resource *Citation) T_SummaryStyle(numSummary int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Summary) >= numSummary {
		return CodeableConceptSelect("Citation.Summary["+strconv.Itoa(numSummary)+"].Style", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.Summary["+strconv.Itoa(numSummary)+"].Style", resource.Summary[numSummary].Style, optionsValueSet)
}
func (resource *Citation) T_SummaryText(numSummary int) templ.Component {

	if resource == nil || len(resource.Summary) >= numSummary {
		return StringInput("Citation.Summary["+strconv.Itoa(numSummary)+"].Text", nil)
	}
	return StringInput("Citation.Summary["+strconv.Itoa(numSummary)+"].Text", &resource.Summary[numSummary].Text)
}
func (resource *Citation) T_ClassificationId(numClassification int) templ.Component {

	if resource == nil || len(resource.Classification) >= numClassification {
		return StringInput("Citation.Classification["+strconv.Itoa(numClassification)+"].Id", nil)
	}
	return StringInput("Citation.Classification["+strconv.Itoa(numClassification)+"].Id", resource.Classification[numClassification].Id)
}
func (resource *Citation) T_ClassificationType(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("Citation.Classification["+strconv.Itoa(numClassification)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.Classification["+strconv.Itoa(numClassification)+"].Type", resource.Classification[numClassification].Type, optionsValueSet)
}
func (resource *Citation) T_ClassificationClassifier(numClassification int, numClassifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Classification) >= numClassification || len(resource.Classification[numClassification].Classifier) >= numClassifier {
		return CodeableConceptSelect("Citation.Classification["+strconv.Itoa(numClassification)+"].Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.Classification["+strconv.Itoa(numClassification)+"].Classifier["+strconv.Itoa(numClassifier)+"]", &resource.Classification[numClassification].Classifier[numClassifier], optionsValueSet)
}
func (resource *Citation) T_StatusDateId(numStatusDate int) templ.Component {

	if resource == nil || len(resource.StatusDate) >= numStatusDate {
		return StringInput("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Id", nil)
	}
	return StringInput("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Id", resource.StatusDate[numStatusDate].Id)
}
func (resource *Citation) T_StatusDateActivity(numStatusDate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.StatusDate) >= numStatusDate {
		return CodeableConceptSelect("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Activity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Activity", &resource.StatusDate[numStatusDate].Activity, optionsValueSet)
}
func (resource *Citation) T_StatusDateActual(numStatusDate int) templ.Component {

	if resource == nil || len(resource.StatusDate) >= numStatusDate {
		return BoolInput("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Actual", nil)
	}
	return BoolInput("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Actual", resource.StatusDate[numStatusDate].Actual)
}
func (resource *Citation) T_RelatesToId(numRelatesTo int) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return StringInput("Citation.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", nil)
	}
	return StringInput("Citation.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", resource.RelatesTo[numRelatesTo].Id)
}
func (resource *Citation) T_RelatesToRelationshipType(numRelatesTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return CodeableConceptSelect("Citation.RelatesTo["+strconv.Itoa(numRelatesTo)+"].RelationshipType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.RelatesTo["+strconv.Itoa(numRelatesTo)+"].RelationshipType", &resource.RelatesTo[numRelatesTo].RelationshipType, optionsValueSet)
}
func (resource *Citation) T_RelatesToTargetClassifier(numRelatesTo int, numTargetClassifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo || len(resource.RelatesTo[numRelatesTo].TargetClassifier) >= numTargetClassifier {
		return CodeableConceptSelect("Citation.RelatesTo["+strconv.Itoa(numRelatesTo)+"].TargetClassifier["+strconv.Itoa(numTargetClassifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.RelatesTo["+strconv.Itoa(numRelatesTo)+"].TargetClassifier["+strconv.Itoa(numTargetClassifier)+"]", &resource.RelatesTo[numRelatesTo].TargetClassifier[numTargetClassifier], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactId() templ.Component {

	if resource == nil {
		return StringInput("Citation.CitedArtifact.Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Id", resource.CitedArtifact.Id)
}
func (resource *Citation) T_CitedArtifactDateAccessed() templ.Component {

	if resource == nil {
		return StringInput("Citation.CitedArtifact.DateAccessed", nil)
	}
	return StringInput("Citation.CitedArtifact.DateAccessed", resource.CitedArtifact.DateAccessed)
}
func (resource *Citation) T_CitedArtifactCurrentState(numCurrentState int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.CurrentState) >= numCurrentState {
		return CodeableConceptSelect("Citation.CitedArtifact.CurrentState["+strconv.Itoa(numCurrentState)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.CurrentState["+strconv.Itoa(numCurrentState)+"]", &resource.CitedArtifact.CurrentState[numCurrentState], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactVersionId() templ.Component {

	if resource == nil {
		return StringInput("Citation.CitedArtifact.Version.Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Version.Id", resource.CitedArtifact.Version.Id)
}
func (resource *Citation) T_CitedArtifactVersionValue() templ.Component {

	if resource == nil {
		return StringInput("Citation.CitedArtifact.Version.Value", nil)
	}
	return StringInput("Citation.CitedArtifact.Version.Value", &resource.CitedArtifact.Version.Value)
}
func (resource *Citation) T_CitedArtifactStatusDateId(numStatusDate int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.StatusDate) >= numStatusDate {
		return StringInput("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Id", resource.CitedArtifact.StatusDate[numStatusDate].Id)
}
func (resource *Citation) T_CitedArtifactStatusDateActivity(numStatusDate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.StatusDate) >= numStatusDate {
		return CodeableConceptSelect("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Activity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Activity", &resource.CitedArtifact.StatusDate[numStatusDate].Activity, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactStatusDateActual(numStatusDate int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.StatusDate) >= numStatusDate {
		return BoolInput("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Actual", nil)
	}
	return BoolInput("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Actual", resource.CitedArtifact.StatusDate[numStatusDate].Actual)
}
func (resource *Citation) T_CitedArtifactTitleId(numTitle int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Title) >= numTitle {
		return StringInput("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Id", resource.CitedArtifact.Title[numTitle].Id)
}
func (resource *Citation) T_CitedArtifactTitleType(numTitle int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Title) >= numTitle || len(resource.CitedArtifact.Title[numTitle].Type) >= numType {
		return CodeableConceptSelect("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Type["+strconv.Itoa(numType)+"]", &resource.CitedArtifact.Title[numTitle].Type[numType], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactTitleLanguage(numTitle int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Title) >= numTitle {
		return CodeableConceptSelect("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Language", resource.CitedArtifact.Title[numTitle].Language, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactTitleText(numTitle int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Title) >= numTitle {
		return StringInput("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Text", nil)
	}
	return StringInput("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Text", &resource.CitedArtifact.Title[numTitle].Text)
}
func (resource *Citation) T_CitedArtifactAbstractId(numAbstract int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Abstract) >= numAbstract {
		return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Id", resource.CitedArtifact.Abstract[numAbstract].Id)
}
func (resource *Citation) T_CitedArtifactAbstractType(numAbstract int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Abstract) >= numAbstract {
		return CodeableConceptSelect("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Type", resource.CitedArtifact.Abstract[numAbstract].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactAbstractLanguage(numAbstract int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Abstract) >= numAbstract {
		return CodeableConceptSelect("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Language", resource.CitedArtifact.Abstract[numAbstract].Language, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactAbstractText(numAbstract int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Abstract) >= numAbstract {
		return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Text", nil)
	}
	return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Text", &resource.CitedArtifact.Abstract[numAbstract].Text)
}
func (resource *Citation) T_CitedArtifactAbstractCopyright(numAbstract int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Abstract) >= numAbstract {
		return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Copyright", nil)
	}
	return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Copyright", resource.CitedArtifact.Abstract[numAbstract].Copyright)
}
func (resource *Citation) T_CitedArtifactPartId() templ.Component {

	if resource == nil {
		return StringInput("Citation.CitedArtifact.Part.Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Part.Id", resource.CitedArtifact.Part.Id)
}
func (resource *Citation) T_CitedArtifactPartType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Citation.CitedArtifact.Part.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Part.Type", resource.CitedArtifact.Part.Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactPartValue() templ.Component {

	if resource == nil {
		return StringInput("Citation.CitedArtifact.Part.Value", nil)
	}
	return StringInput("Citation.CitedArtifact.Part.Value", resource.CitedArtifact.Part.Value)
}
func (resource *Citation) T_CitedArtifactRelatesToId(numRelatesTo int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.RelatesTo) >= numRelatesTo {
		return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", resource.CitedArtifact.RelatesTo[numRelatesTo].Id)
}
func (resource *Citation) T_CitedArtifactRelatesToRelationshipType(numRelatesTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.RelatesTo) >= numRelatesTo {
		return CodeableConceptSelect("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].RelationshipType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].RelationshipType", &resource.CitedArtifact.RelatesTo[numRelatesTo].RelationshipType, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactRelatesToTargetClassifier(numRelatesTo int, numTargetClassifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.RelatesTo) >= numRelatesTo || len(resource.CitedArtifact.RelatesTo[numRelatesTo].TargetClassifier) >= numTargetClassifier {
		return CodeableConceptSelect("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].TargetClassifier["+strconv.Itoa(numTargetClassifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].TargetClassifier["+strconv.Itoa(numTargetClassifier)+"]", &resource.CitedArtifact.RelatesTo[numRelatesTo].TargetClassifier[numTargetClassifier], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactPublicationFormId(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Id", resource.CitedArtifact.PublicationForm[numPublicationForm].Id)
}
func (resource *Citation) T_CitedArtifactPublicationFormArticleDate(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].ArticleDate", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].ArticleDate", resource.CitedArtifact.PublicationForm[numPublicationForm].ArticleDate)
}
func (resource *Citation) T_CitedArtifactPublicationFormLastRevisionDate(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].LastRevisionDate", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].LastRevisionDate", resource.CitedArtifact.PublicationForm[numPublicationForm].LastRevisionDate)
}
func (resource *Citation) T_CitedArtifactPublicationFormLanguage(numPublicationForm int, numLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm || len(resource.CitedArtifact.PublicationForm[numPublicationForm].Language) >= numLanguage {
		return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Language["+strconv.Itoa(numLanguage)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Language["+strconv.Itoa(numLanguage)+"]", &resource.CitedArtifact.PublicationForm[numPublicationForm].Language[numLanguage], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactPublicationFormAccessionNumber(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].AccessionNumber", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].AccessionNumber", resource.CitedArtifact.PublicationForm[numPublicationForm].AccessionNumber)
}
func (resource *Citation) T_CitedArtifactPublicationFormPageString(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PageString", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PageString", resource.CitedArtifact.PublicationForm[numPublicationForm].PageString)
}
func (resource *Citation) T_CitedArtifactPublicationFormFirstPage(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].FirstPage", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].FirstPage", resource.CitedArtifact.PublicationForm[numPublicationForm].FirstPage)
}
func (resource *Citation) T_CitedArtifactPublicationFormLastPage(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].LastPage", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].LastPage", resource.CitedArtifact.PublicationForm[numPublicationForm].LastPage)
}
func (resource *Citation) T_CitedArtifactPublicationFormPageCount(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PageCount", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PageCount", resource.CitedArtifact.PublicationForm[numPublicationForm].PageCount)
}
func (resource *Citation) T_CitedArtifactPublicationFormCopyright(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Copyright", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Copyright", resource.CitedArtifact.PublicationForm[numPublicationForm].Copyright)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInId(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Id", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Id", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.Id)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInType(numPublicationForm int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Type", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInTitle(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Title", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Title", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.Title)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInPublisherLocation(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.PublisherLocation", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.PublisherLocation", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.PublisherLocation)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseId(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.Id", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.Id", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.Id)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseCitedMedium(numPublicationForm int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.CitedMedium", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.CitedMedium", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.CitedMedium, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseVolume(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.Volume", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.Volume", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.Volume)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseIssue(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.Issue", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.Issue", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.Issue)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseDateOfPublicationId(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Id", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Id", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.DateOfPublication.Id)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseDateOfPublicationDate(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Date", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Date", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.DateOfPublication.Date)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseDateOfPublicationYear(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Year", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Year", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.DateOfPublication.Year)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseDateOfPublicationMonth(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Month", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Month", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.DateOfPublication.Month)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseDateOfPublicationDay(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Day", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Day", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.DateOfPublication.Day)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseDateOfPublicationSeason(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Season", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Season", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.DateOfPublication.Season)
}
func (resource *Citation) T_CitedArtifactPublicationFormPeriodicReleaseDateOfPublicationText(numPublicationForm int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Text", nil)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PeriodicRelease.DateOfPublication.Text", resource.CitedArtifact.PublicationForm[numPublicationForm].PeriodicRelease.DateOfPublication.Text)
}
func (resource *Citation) T_CitedArtifactWebLocationId(numWebLocation int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.WebLocation) >= numWebLocation {
		return StringInput("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Id", resource.CitedArtifact.WebLocation[numWebLocation].Id)
}
func (resource *Citation) T_CitedArtifactWebLocationType(numWebLocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.WebLocation) >= numWebLocation {
		return CodeableConceptSelect("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Type", resource.CitedArtifact.WebLocation[numWebLocation].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactWebLocationUrl(numWebLocation int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.WebLocation) >= numWebLocation {
		return StringInput("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Url", nil)
	}
	return StringInput("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Url", resource.CitedArtifact.WebLocation[numWebLocation].Url)
}
func (resource *Citation) T_CitedArtifactClassificationId(numClassification int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Classification) >= numClassification {
		return StringInput("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Id", resource.CitedArtifact.Classification[numClassification].Id)
}
func (resource *Citation) T_CitedArtifactClassificationType(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Classification) >= numClassification {
		return CodeableConceptSelect("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Type", resource.CitedArtifact.Classification[numClassification].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactClassificationClassifier(numClassification int, numClassifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Classification) >= numClassification || len(resource.CitedArtifact.Classification[numClassification].Classifier) >= numClassifier {
		return CodeableConceptSelect("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Classifier["+strconv.Itoa(numClassifier)+"]", &resource.CitedArtifact.Classification[numClassification].Classifier[numClassifier], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactClassificationWhoClassifiedId(numClassification int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Classification) >= numClassification {
		return StringInput("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].WhoClassified.Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].WhoClassified.Id", resource.CitedArtifact.Classification[numClassification].WhoClassified.Id)
}
func (resource *Citation) T_CitedArtifactClassificationWhoClassifiedClassifierCopyright(numClassification int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Classification) >= numClassification {
		return StringInput("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].WhoClassified.ClassifierCopyright", nil)
	}
	return StringInput("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].WhoClassified.ClassifierCopyright", resource.CitedArtifact.Classification[numClassification].WhoClassified.ClassifierCopyright)
}
func (resource *Citation) T_CitedArtifactClassificationWhoClassifiedFreeToShare(numClassification int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Classification) >= numClassification {
		return BoolInput("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].WhoClassified.FreeToShare", nil)
	}
	return BoolInput("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].WhoClassified.FreeToShare", resource.CitedArtifact.Classification[numClassification].WhoClassified.FreeToShare)
}
func (resource *Citation) T_CitedArtifactContributorshipId() templ.Component {

	if resource == nil {
		return StringInput("Citation.CitedArtifact.Contributorship.Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Id", resource.CitedArtifact.Contributorship.Id)
}
func (resource *Citation) T_CitedArtifactContributorshipComplete() templ.Component {

	if resource == nil {
		return BoolInput("Citation.CitedArtifact.Contributorship.Complete", nil)
	}
	return BoolInput("Citation.CitedArtifact.Contributorship.Complete", resource.CitedArtifact.Contributorship.Complete)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryId(numEntry int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry {
		return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].Id", resource.CitedArtifact.Contributorship.Entry[numEntry].Id)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryInitials(numEntry int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry {
		return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].Initials", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].Initials", resource.CitedArtifact.Contributorship.Entry[numEntry].Initials)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryCollectiveName(numEntry int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry {
		return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].CollectiveName", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].CollectiveName", resource.CitedArtifact.Contributorship.Entry[numEntry].CollectiveName)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionType(numEntry int, numContributionType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry || len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionType) >= numContributionType {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionType["+strconv.Itoa(numContributionType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionType["+strconv.Itoa(numContributionType)+"]", &resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionType[numContributionType], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryRole(numEntry int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].Role", resource.CitedArtifact.Contributorship.Entry[numEntry].Role, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryCorrespondingContact(numEntry int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry {
		return BoolInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].CorrespondingContact", nil)
	}
	return BoolInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].CorrespondingContact", resource.CitedArtifact.Contributorship.Entry[numEntry].CorrespondingContact)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryListOrder(numEntry int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry {
		return IntInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ListOrder", nil)
	}
	return IntInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ListOrder", resource.CitedArtifact.Contributorship.Entry[numEntry].ListOrder)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryAffiliationInfoId(numEntry int, numAffiliationInfo int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry || len(resource.CitedArtifact.Contributorship.Entry[numEntry].AffiliationInfo) >= numAffiliationInfo {
		return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].AffiliationInfo["+strconv.Itoa(numAffiliationInfo)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].AffiliationInfo["+strconv.Itoa(numAffiliationInfo)+"].Id", resource.CitedArtifact.Contributorship.Entry[numEntry].AffiliationInfo[numAffiliationInfo].Id)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryAffiliationInfoAffiliation(numEntry int, numAffiliationInfo int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry || len(resource.CitedArtifact.Contributorship.Entry[numEntry].AffiliationInfo) >= numAffiliationInfo {
		return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].AffiliationInfo["+strconv.Itoa(numAffiliationInfo)+"].Affiliation", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].AffiliationInfo["+strconv.Itoa(numAffiliationInfo)+"].Affiliation", resource.CitedArtifact.Contributorship.Entry[numEntry].AffiliationInfo[numAffiliationInfo].Affiliation)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryAffiliationInfoRole(numEntry int, numAffiliationInfo int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry || len(resource.CitedArtifact.Contributorship.Entry[numEntry].AffiliationInfo) >= numAffiliationInfo {
		return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].AffiliationInfo["+strconv.Itoa(numAffiliationInfo)+"].Role", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].AffiliationInfo["+strconv.Itoa(numAffiliationInfo)+"].Role", resource.CitedArtifact.Contributorship.Entry[numEntry].AffiliationInfo[numAffiliationInfo].Role)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionInstanceId(numEntry int, numContributionInstance int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry || len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance) >= numContributionInstance {
		return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Id", resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance[numContributionInstance].Id)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionInstanceType(numEntry int, numContributionInstance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry || len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance) >= numContributionInstance {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Type", &resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance[numContributionInstance].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionInstanceTime(numEntry int, numContributionInstance int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Entry) >= numEntry || len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance) >= numContributionInstance {
		return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Time", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Time", resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance[numContributionInstance].Time)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryId(numSummary int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Summary) >= numSummary {
		return StringInput("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Id", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Id", resource.CitedArtifact.Contributorship.Summary[numSummary].Id)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryType(numSummary int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Summary) >= numSummary {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Type", resource.CitedArtifact.Contributorship.Summary[numSummary].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryStyle(numSummary int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Summary) >= numSummary {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Style", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Style", resource.CitedArtifact.Contributorship.Summary[numSummary].Style, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipSummarySource(numSummary int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Summary) >= numSummary {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Source", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Source", resource.CitedArtifact.Contributorship.Summary[numSummary].Source, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryValue(numSummary int) templ.Component {

	if resource == nil || len(resource.CitedArtifact.Contributorship.Summary) >= numSummary {
		return StringInput("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Value", nil)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Value", &resource.CitedArtifact.Contributorship.Summary[numSummary].Value)
}
