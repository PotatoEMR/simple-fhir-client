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

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type Citation struct {
	Id                     *string                  `json:"id,omitempty"`
	Meta                   *Meta                    `json:"meta,omitempty"`
	ImplicitRules          *string                  `json:"implicitRules,omitempty"`
	Language               *string                  `json:"language,omitempty"`
	Text                   *Narrative               `json:"text,omitempty"`
	Contained              []Resource               `json:"contained,omitempty"`
	Extension              []Extension              `json:"extension,omitempty"`
	ModifierExtension      []Extension              `json:"modifierExtension,omitempty"`
	Url                    *string                  `json:"url,omitempty"`
	Identifier             []Identifier             `json:"identifier,omitempty"`
	Version                *string                  `json:"version,omitempty"`
	VersionAlgorithmString *string                  `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                  `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                  `json:"name,omitempty"`
	Title                  *string                  `json:"title,omitempty"`
	Status                 string                   `json:"status"`
	Experimental           *bool                    `json:"experimental,omitempty"`
	Date                   *time.Time               `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                  `json:"publisher,omitempty"`
	Contact                []ContactDetail          `json:"contact,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	UseContext             []UsageContext           `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept        `json:"jurisdiction,omitempty"`
	Purpose                *string                  `json:"purpose,omitempty"`
	Copyright              *string                  `json:"copyright,omitempty"`
	CopyrightLabel         *string                  `json:"copyrightLabel,omitempty"`
	ApprovalDate           *time.Time               `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time               `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
	EffectivePeriod        *Period                  `json:"effectivePeriod,omitempty"`
	Author                 []ContactDetail          `json:"author,omitempty"`
	Editor                 []ContactDetail          `json:"editor,omitempty"`
	Reviewer               []ContactDetail          `json:"reviewer,omitempty"`
	Endorser               []ContactDetail          `json:"endorser,omitempty"`
	Summary                []CitationSummary        `json:"summary,omitempty"`
	Classification         []CitationClassification `json:"classification,omitempty"`
	Note                   []Annotation             `json:"note,omitempty"`
	CurrentState           []CodeableConcept        `json:"currentState,omitempty"`
	StatusDate             []CitationStatusDate     `json:"statusDate,omitempty"`
	RelatedArtifact        []RelatedArtifact        `json:"relatedArtifact,omitempty"`
	CitedArtifact          *CitationCitedArtifact   `json:"citedArtifact,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationSummary struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Style             *CodeableConcept `json:"style,omitempty"`
	Text              string           `json:"text"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Classifier        []CodeableConcept `json:"classifier,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationStatusDate struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Activity          CodeableConcept `json:"activity"`
	Actual            *bool           `json:"actual,omitempty"`
	Period            Period          `json:"period"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifact struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                           `json:"identifier,omitempty"`
	RelatedIdentifier []Identifier                           `json:"relatedIdentifier,omitempty"`
	DateAccessed      *time.Time                             `json:"dateAccessed,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactVersion struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Value             string      `json:"value"`
	BaseCitation      *Reference  `json:"baseCitation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactStatusDate struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Activity          CodeableConcept `json:"activity"`
	Actual            *bool           `json:"actual,omitempty"`
	Period            Period          `json:"period"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactTitle struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	Language          *CodeableConcept  `json:"language,omitempty"`
	Text              string            `json:"text"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactAbstract struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Language          *CodeableConcept `json:"language,omitempty"`
	Text              string           `json:"text"`
	Copyright         *string          `json:"copyright,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactPart struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Value             *string          `json:"value,omitempty"`
	BaseCitation      *Reference       `json:"baseCitation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactRelatesTo struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              string            `json:"type"`
	Classifier        []CodeableConcept `json:"classifier,omitempty"`
	Label             *string           `json:"label,omitempty"`
	Display           *string           `json:"display,omitempty"`
	Citation          *string           `json:"citation,omitempty"`
	Document          *Attachment       `json:"document,omitempty"`
	Resource          *string           `json:"resource,omitempty"`
	ResourceReference *Reference        `json:"resourceReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactPublicationForm struct {
	Id                    *string                                          `json:"id,omitempty"`
	Extension             []Extension                                      `json:"extension,omitempty"`
	ModifierExtension     []Extension                                      `json:"modifierExtension,omitempty"`
	PublishedIn           *CitationCitedArtifactPublicationFormPublishedIn `json:"publishedIn,omitempty"`
	CitedMedium           *CodeableConcept                                 `json:"citedMedium,omitempty"`
	Volume                *string                                          `json:"volume,omitempty"`
	Issue                 *string                                          `json:"issue,omitempty"`
	ArticleDate           *time.Time                                       `json:"articleDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	PublicationDateText   *string                                          `json:"publicationDateText,omitempty"`
	PublicationDateSeason *string                                          `json:"publicationDateSeason,omitempty"`
	LastRevisionDate      *time.Time                                       `json:"lastRevisionDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Language              []CodeableConcept                                `json:"language,omitempty"`
	AccessionNumber       *string                                          `json:"accessionNumber,omitempty"`
	PageString            *string                                          `json:"pageString,omitempty"`
	FirstPage             *string                                          `json:"firstPage,omitempty"`
	LastPage              *string                                          `json:"lastPage,omitempty"`
	PageCount             *string                                          `json:"pageCount,omitempty"`
	Copyright             *string                                          `json:"copyright,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
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

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactWebLocation struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Classifier        []CodeableConcept `json:"classifier,omitempty"`
	Url               *string           `json:"url,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactClassification struct {
	Id                 *string           `json:"id,omitempty"`
	Extension          []Extension       `json:"extension,omitempty"`
	ModifierExtension  []Extension       `json:"modifierExtension,omitempty"`
	Type               *CodeableConcept  `json:"type,omitempty"`
	Classifier         []CodeableConcept `json:"classifier,omitempty"`
	ArtifactAssessment []Reference       `json:"artifactAssessment,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactContributorship struct {
	Id                *string                                       `json:"id,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Complete          *bool                                         `json:"complete,omitempty"`
	Entry             []CitationCitedArtifactContributorshipEntry   `json:"entry,omitempty"`
	Summary           []CitationCitedArtifactContributorshipSummary `json:"summary,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactContributorshipEntry struct {
	Id                   *string                                                         `json:"id,omitempty"`
	Extension            []Extension                                                     `json:"extension,omitempty"`
	ModifierExtension    []Extension                                                     `json:"modifierExtension,omitempty"`
	Contributor          Reference                                                       `json:"contributor"`
	ForenameInitials     *string                                                         `json:"forenameInitials,omitempty"`
	Affiliation          []Reference                                                     `json:"affiliation,omitempty"`
	ContributionType     []CodeableConcept                                               `json:"contributionType,omitempty"`
	Role                 *CodeableConcept                                                `json:"role,omitempty"`
	ContributionInstance []CitationCitedArtifactContributorshipEntryContributionInstance `json:"contributionInstance,omitempty"`
	CorrespondingContact *bool                                                           `json:"correspondingContact,omitempty"`
	RankingOrder         *int                                                            `json:"rankingOrder,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
type CitationCitedArtifactContributorshipEntryContributionInstance struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Time              *time.Time      `json:"time,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Citation
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
func (r Citation) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Citation/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Citation"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Citation) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.Url", nil, htmlAttrs)
	}
	return StringInput("Citation.Url", resource.Url, htmlAttrs)
}
func (resource *Citation) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.Version", nil, htmlAttrs)
	}
	return StringInput("Citation.Version", resource.Version, htmlAttrs)
}
func (resource *Citation) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("Citation.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *Citation) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("Citation.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Citation.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.Name", nil, htmlAttrs)
	}
	return StringInput("Citation.Name", resource.Name, htmlAttrs)
}
func (resource *Citation) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.Title", nil, htmlAttrs)
	}
	return StringInput("Citation.Title", resource.Title, htmlAttrs)
}
func (resource *Citation) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Citation.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Citation.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Citation.Experimental", nil, htmlAttrs)
	}
	return BoolInput("Citation.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *Citation) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Citation.Date", nil, htmlAttrs)
	}
	return DateTimeInput("Citation.Date", resource.Date, htmlAttrs)
}
func (resource *Citation) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.Publisher", nil, htmlAttrs)
	}
	return StringInput("Citation.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *Citation) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.Description", nil, htmlAttrs)
	}
	return StringInput("Citation.Description", resource.Description, htmlAttrs)
}
func (resource *Citation) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Citation.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.Purpose", nil, htmlAttrs)
	}
	return StringInput("Citation.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *Citation) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.Copyright", nil, htmlAttrs)
	}
	return StringInput("Citation.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *Citation) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("Citation.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *Citation) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("Citation.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("Citation.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Citation) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("Citation.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("Citation.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Citation) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Citation.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Citation.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Citation) T_CurrentState(numCurrentState int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCurrentState >= len(resource.CurrentState) {
		return CodeableConceptSelect("Citation.CurrentState["+strconv.Itoa(numCurrentState)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CurrentState["+strconv.Itoa(numCurrentState)+"]", &resource.CurrentState[numCurrentState], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_SummaryStyle(numSummary int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSummary >= len(resource.Summary) {
		return CodeableConceptSelect("Citation.Summary["+strconv.Itoa(numSummary)+"].Style", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.Summary["+strconv.Itoa(numSummary)+"].Style", resource.Summary[numSummary].Style, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_SummaryText(numSummary int, htmlAttrs string) templ.Component {
	if resource == nil || numSummary >= len(resource.Summary) {
		return StringInput("Citation.Summary["+strconv.Itoa(numSummary)+"].Text", nil, htmlAttrs)
	}
	return StringInput("Citation.Summary["+strconv.Itoa(numSummary)+"].Text", &resource.Summary[numSummary].Text, htmlAttrs)
}
func (resource *Citation) T_ClassificationType(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("Citation.Classification["+strconv.Itoa(numClassification)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.Classification["+strconv.Itoa(numClassification)+"].Type", resource.Classification[numClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_ClassificationClassifier(numClassification int, numClassifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) || numClassifier >= len(resource.Classification[numClassification].Classifier) {
		return CodeableConceptSelect("Citation.Classification["+strconv.Itoa(numClassification)+"].Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.Classification["+strconv.Itoa(numClassification)+"].Classifier["+strconv.Itoa(numClassifier)+"]", &resource.Classification[numClassification].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_StatusDateActivity(numStatusDate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStatusDate >= len(resource.StatusDate) {
		return CodeableConceptSelect("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Activity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Activity", &resource.StatusDate[numStatusDate].Activity, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_StatusDateActual(numStatusDate int, htmlAttrs string) templ.Component {
	if resource == nil || numStatusDate >= len(resource.StatusDate) {
		return BoolInput("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Actual", nil, htmlAttrs)
	}
	return BoolInput("Citation.StatusDate["+strconv.Itoa(numStatusDate)+"].Actual", resource.StatusDate[numStatusDate].Actual, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactDateAccessed(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Citation.CitedArtifact.DateAccessed", nil, htmlAttrs)
	}
	return DateTimeInput("Citation.CitedArtifact.DateAccessed", resource.CitedArtifact.DateAccessed, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactCurrentState(numCurrentState int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCurrentState >= len(resource.CitedArtifact.CurrentState) {
		return CodeableConceptSelect("Citation.CitedArtifact.CurrentState["+strconv.Itoa(numCurrentState)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.CurrentState["+strconv.Itoa(numCurrentState)+"]", &resource.CitedArtifact.CurrentState[numCurrentState], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactNote(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.CitedArtifact.Note) {
		return AnnotationTextArea("Citation.CitedArtifact.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Citation.CitedArtifact.Note["+strconv.Itoa(numNote)+"]", &resource.CitedArtifact.Note[numNote], htmlAttrs)
}
func (resource *Citation) T_CitedArtifactVersionValue(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.CitedArtifact.Version.Value", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.Version.Value", &resource.CitedArtifact.Version.Value, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactStatusDateActivity(numStatusDate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStatusDate >= len(resource.CitedArtifact.StatusDate) {
		return CodeableConceptSelect("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Activity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Activity", &resource.CitedArtifact.StatusDate[numStatusDate].Activity, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactStatusDateActual(numStatusDate int, htmlAttrs string) templ.Component {
	if resource == nil || numStatusDate >= len(resource.CitedArtifact.StatusDate) {
		return BoolInput("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Actual", nil, htmlAttrs)
	}
	return BoolInput("Citation.CitedArtifact.StatusDate["+strconv.Itoa(numStatusDate)+"].Actual", resource.CitedArtifact.StatusDate[numStatusDate].Actual, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactTitleType(numTitle int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTitle >= len(resource.CitedArtifact.Title) || numType >= len(resource.CitedArtifact.Title[numTitle].Type) {
		return CodeableConceptSelect("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Type["+strconv.Itoa(numType)+"]", &resource.CitedArtifact.Title[numTitle].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactTitleText(numTitle int, htmlAttrs string) templ.Component {
	if resource == nil || numTitle >= len(resource.CitedArtifact.Title) {
		return StringInput("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Text", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.Title["+strconv.Itoa(numTitle)+"].Text", &resource.CitedArtifact.Title[numTitle].Text, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactAbstractType(numAbstract int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAbstract >= len(resource.CitedArtifact.Abstract) {
		return CodeableConceptSelect("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Type", resource.CitedArtifact.Abstract[numAbstract].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactAbstractText(numAbstract int, htmlAttrs string) templ.Component {
	if resource == nil || numAbstract >= len(resource.CitedArtifact.Abstract) {
		return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Text", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Text", &resource.CitedArtifact.Abstract[numAbstract].Text, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactAbstractCopyright(numAbstract int, htmlAttrs string) templ.Component {
	if resource == nil || numAbstract >= len(resource.CitedArtifact.Abstract) {
		return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Copyright", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.Abstract["+strconv.Itoa(numAbstract)+"].Copyright", resource.CitedArtifact.Abstract[numAbstract].Copyright, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPartType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Citation.CitedArtifact.Part.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Part.Type", resource.CitedArtifact.Part.Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPartValue(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Citation.CitedArtifact.Part.Value", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.Part.Value", resource.CitedArtifact.Part.Value, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToType(numRelatesTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return CodeSelect("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Type", &resource.CitedArtifact.RelatesTo[numRelatesTo].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToClassifier(numRelatesTo int, numClassifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) || numClassifier >= len(resource.CitedArtifact.RelatesTo[numRelatesTo].Classifier) {
		return CodeableConceptSelect("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Classifier["+strconv.Itoa(numClassifier)+"]", &resource.CitedArtifact.RelatesTo[numRelatesTo].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToLabel(numRelatesTo int, htmlAttrs string) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Label", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Label", resource.CitedArtifact.RelatesTo[numRelatesTo].Label, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToDisplay(numRelatesTo int, htmlAttrs string) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Display", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Display", resource.CitedArtifact.RelatesTo[numRelatesTo].Display, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToCitation(numRelatesTo int, htmlAttrs string) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Citation", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Citation", resource.CitedArtifact.RelatesTo[numRelatesTo].Citation, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToResource(numRelatesTo int, htmlAttrs string) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Resource", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Resource", resource.CitedArtifact.RelatesTo[numRelatesTo].Resource, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormCitedMedium(numPublicationForm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].CitedMedium", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].CitedMedium", resource.CitedArtifact.PublicationForm[numPublicationForm].CitedMedium, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormVolume(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Volume", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Volume", resource.CitedArtifact.PublicationForm[numPublicationForm].Volume, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormIssue(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Issue", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Issue", resource.CitedArtifact.PublicationForm[numPublicationForm].Issue, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormArticleDate(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return DateTimeInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].ArticleDate", nil, htmlAttrs)
	}
	return DateTimeInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].ArticleDate", resource.CitedArtifact.PublicationForm[numPublicationForm].ArticleDate, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublicationDateText(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublicationDateText", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublicationDateText", resource.CitedArtifact.PublicationForm[numPublicationForm].PublicationDateText, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublicationDateSeason(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublicationDateSeason", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublicationDateSeason", resource.CitedArtifact.PublicationForm[numPublicationForm].PublicationDateSeason, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormLastRevisionDate(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return DateTimeInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].LastRevisionDate", nil, htmlAttrs)
	}
	return DateTimeInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].LastRevisionDate", resource.CitedArtifact.PublicationForm[numPublicationForm].LastRevisionDate, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormAccessionNumber(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].AccessionNumber", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].AccessionNumber", resource.CitedArtifact.PublicationForm[numPublicationForm].AccessionNumber, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPageString(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PageString", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PageString", resource.CitedArtifact.PublicationForm[numPublicationForm].PageString, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormFirstPage(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].FirstPage", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].FirstPage", resource.CitedArtifact.PublicationForm[numPublicationForm].FirstPage, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormLastPage(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].LastPage", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].LastPage", resource.CitedArtifact.PublicationForm[numPublicationForm].LastPage, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPageCount(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PageCount", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PageCount", resource.CitedArtifact.PublicationForm[numPublicationForm].PageCount, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormCopyright(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Copyright", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].Copyright", resource.CitedArtifact.PublicationForm[numPublicationForm].Copyright, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInType(numPublicationForm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Type", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInTitle(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Title", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.Title", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.Title, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInPublisherLocation(numPublicationForm int, htmlAttrs string) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.PublisherLocation", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.PublicationForm["+strconv.Itoa(numPublicationForm)+"].PublishedIn.PublisherLocation", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.PublisherLocation, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactWebLocationClassifier(numWebLocation int, numClassifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numWebLocation >= len(resource.CitedArtifact.WebLocation) || numClassifier >= len(resource.CitedArtifact.WebLocation[numWebLocation].Classifier) {
		return CodeableConceptSelect("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Classifier["+strconv.Itoa(numClassifier)+"]", &resource.CitedArtifact.WebLocation[numWebLocation].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactWebLocationUrl(numWebLocation int, htmlAttrs string) templ.Component {
	if resource == nil || numWebLocation >= len(resource.CitedArtifact.WebLocation) {
		return StringInput("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Url", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.WebLocation["+strconv.Itoa(numWebLocation)+"].Url", resource.CitedArtifact.WebLocation[numWebLocation].Url, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactClassificationType(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.CitedArtifact.Classification) {
		return CodeableConceptSelect("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Type", resource.CitedArtifact.Classification[numClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactClassificationClassifier(numClassification int, numClassifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.CitedArtifact.Classification) || numClassifier >= len(resource.CitedArtifact.Classification[numClassification].Classifier) {
		return CodeableConceptSelect("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Classification["+strconv.Itoa(numClassification)+"].Classifier["+strconv.Itoa(numClassifier)+"]", &resource.CitedArtifact.Classification[numClassification].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipComplete(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Citation.CitedArtifact.Contributorship.Complete", nil, htmlAttrs)
	}
	return BoolInput("Citation.CitedArtifact.Contributorship.Complete", resource.CitedArtifact.Contributorship.Complete, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryForenameInitials(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) {
		return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ForenameInitials", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ForenameInitials", resource.CitedArtifact.Contributorship.Entry[numEntry].ForenameInitials, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionType(numEntry int, numContributionType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) || numContributionType >= len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionType) {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionType["+strconv.Itoa(numContributionType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionType["+strconv.Itoa(numContributionType)+"]", &resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionType[numContributionType], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryRole(numEntry int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].Role", resource.CitedArtifact.Contributorship.Entry[numEntry].Role, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryCorrespondingContact(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) {
		return BoolInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].CorrespondingContact", nil, htmlAttrs)
	}
	return BoolInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].CorrespondingContact", resource.CitedArtifact.Contributorship.Entry[numEntry].CorrespondingContact, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryRankingOrder(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) {
		return IntInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].RankingOrder", nil, htmlAttrs)
	}
	return IntInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].RankingOrder", resource.CitedArtifact.Contributorship.Entry[numEntry].RankingOrder, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionInstanceType(numEntry int, numContributionInstance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) || numContributionInstance >= len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance) {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Type", &resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance[numContributionInstance].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionInstanceTime(numEntry int, numContributionInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) || numContributionInstance >= len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance) {
		return DateTimeInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Time", nil, htmlAttrs)
	}
	return DateTimeInput("Citation.CitedArtifact.Contributorship.Entry["+strconv.Itoa(numEntry)+"].ContributionInstance["+strconv.Itoa(numContributionInstance)+"].Time", resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance[numContributionInstance].Time, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryType(numSummary int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSummary >= len(resource.CitedArtifact.Contributorship.Summary) {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Type", resource.CitedArtifact.Contributorship.Summary[numSummary].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryStyle(numSummary int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSummary >= len(resource.CitedArtifact.Contributorship.Summary) {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Style", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Style", resource.CitedArtifact.Contributorship.Summary[numSummary].Style, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipSummarySource(numSummary int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSummary >= len(resource.CitedArtifact.Contributorship.Summary) {
		return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Source", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Source", resource.CitedArtifact.Contributorship.Summary[numSummary].Source, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryValue(numSummary int, htmlAttrs string) templ.Component {
	if resource == nil || numSummary >= len(resource.CitedArtifact.Contributorship.Summary) {
		return StringInput("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Value", nil, htmlAttrs)
	}
	return StringInput("Citation.CitedArtifact.Contributorship.Summary["+strconv.Itoa(numSummary)+"].Value", &resource.CitedArtifact.Contributorship.Summary[numSummary].Value, htmlAttrs)
}
