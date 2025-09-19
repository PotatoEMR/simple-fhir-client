package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                   *FhirDateTime            `json:"date,omitempty"`
	Publisher              *string                  `json:"publisher,omitempty"`
	Contact                []ContactDetail          `json:"contact,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	UseContext             []UsageContext           `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept        `json:"jurisdiction,omitempty"`
	Purpose                *string                  `json:"purpose,omitempty"`
	Copyright              *string                  `json:"copyright,omitempty"`
	CopyrightLabel         *string                  `json:"copyrightLabel,omitempty"`
	ApprovalDate           *FhirDate                `json:"approvalDate,omitempty"`
	LastReviewDate         *FhirDate                `json:"lastReviewDate,omitempty"`
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
	DateAccessed      *FhirDateTime                          `json:"dateAccessed,omitempty"`
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
	ArticleDate           *FhirDateTime                                    `json:"articleDate,omitempty"`
	PublicationDateText   *string                                          `json:"publicationDateText,omitempty"`
	PublicationDateSeason *string                                          `json:"publicationDateSeason,omitempty"`
	LastRevisionDate      *FhirDateTime                                    `json:"lastRevisionDate,omitempty"`
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
	Time              *FhirDateTime   `json:"time,omitempty"`
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
func (resource *Citation) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *Citation) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *Citation) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *Citation) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Citation) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *Citation) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *Citation) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *Citation) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *Citation) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *Citation) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Citation) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *Citation) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *Citation) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *Citation) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *Citation) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Citation) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Citation) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *Citation) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *Citation) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *Citation) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *Citation) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *Citation) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Citation) T_CurrentState(numCurrentState int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCurrentState >= len(resource.CurrentState) {
		return CodeableConceptSelect("currentState["+strconv.Itoa(numCurrentState)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("currentState["+strconv.Itoa(numCurrentState)+"]", &resource.CurrentState[numCurrentState], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *Citation) T_SummaryStyle(numSummary int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSummary >= len(resource.Summary) {
		return CodeableConceptSelect("summary["+strconv.Itoa(numSummary)+"].style", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("summary["+strconv.Itoa(numSummary)+"].style", resource.Summary[numSummary].Style, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_ClassificationType(numClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].type", resource.Classification[numClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_ClassificationClassifier(numClassification int, numClassifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) || numClassifier >= len(resource.Classification[numClassification].Classifier) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].classifier["+strconv.Itoa(numClassifier)+"]", &resource.Classification[numClassification].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_StatusDateActivity(numStatusDate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusDate >= len(resource.StatusDate) {
		return CodeableConceptSelect("statusDate["+strconv.Itoa(numStatusDate)+"].activity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusDate["+strconv.Itoa(numStatusDate)+"].activity", &resource.StatusDate[numStatusDate].Activity, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_StatusDateActual(numStatusDate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusDate >= len(resource.StatusDate) {
		return BoolInput("statusDate["+strconv.Itoa(numStatusDate)+"].actual", nil, htmlAttrs)
	}
	return BoolInput("statusDate["+strconv.Itoa(numStatusDate)+"].actual", resource.StatusDate[numStatusDate].Actual, htmlAttrs)
}
func (resource *Citation) T_StatusDatePeriod(numStatusDate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusDate >= len(resource.StatusDate) {
		return PeriodInput("statusDate["+strconv.Itoa(numStatusDate)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("statusDate["+strconv.Itoa(numStatusDate)+"].period", &resource.StatusDate[numStatusDate].Period, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatedIdentifier(numRelatedIdentifier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedIdentifier >= len(resource.CitedArtifact.RelatedIdentifier) {
		return IdentifierInput("citedArtifact.relatedIdentifier["+strconv.Itoa(numRelatedIdentifier)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("citedArtifact.relatedIdentifier["+strconv.Itoa(numRelatedIdentifier)+"]", &resource.CitedArtifact.RelatedIdentifier[numRelatedIdentifier], htmlAttrs)
}
func (resource *Citation) T_CitedArtifactDateAccessed(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("citedArtifact.dateAccessed", nil, htmlAttrs)
	}
	return FhirDateTimeInput("citedArtifact.dateAccessed", resource.CitedArtifact.DateAccessed, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactCurrentState(numCurrentState int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCurrentState >= len(resource.CitedArtifact.CurrentState) {
		return CodeableConceptSelect("citedArtifact.currentState["+strconv.Itoa(numCurrentState)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.currentState["+strconv.Itoa(numCurrentState)+"]", &resource.CitedArtifact.CurrentState[numCurrentState], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactNote(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.CitedArtifact.Note) {
		return AnnotationTextArea("citedArtifact.note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("citedArtifact.note["+strconv.Itoa(numNote)+"]", &resource.CitedArtifact.Note[numNote], htmlAttrs)
}
func (resource *Citation) T_CitedArtifactVersionValue(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("citedArtifact.version.value", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.version.value", &resource.CitedArtifact.Version.Value, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactVersionBaseCitation(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("citedArtifact.version.baseCitation", nil, htmlAttrs)
	}
	return ReferenceInput("citedArtifact.version.baseCitation", resource.CitedArtifact.Version.BaseCitation, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactStatusDateActivity(numStatusDate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusDate >= len(resource.CitedArtifact.StatusDate) {
		return CodeableConceptSelect("citedArtifact.statusDate["+strconv.Itoa(numStatusDate)+"].activity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.statusDate["+strconv.Itoa(numStatusDate)+"].activity", &resource.CitedArtifact.StatusDate[numStatusDate].Activity, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactStatusDateActual(numStatusDate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusDate >= len(resource.CitedArtifact.StatusDate) {
		return BoolInput("citedArtifact.statusDate["+strconv.Itoa(numStatusDate)+"].actual", nil, htmlAttrs)
	}
	return BoolInput("citedArtifact.statusDate["+strconv.Itoa(numStatusDate)+"].actual", resource.CitedArtifact.StatusDate[numStatusDate].Actual, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactStatusDatePeriod(numStatusDate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusDate >= len(resource.CitedArtifact.StatusDate) {
		return PeriodInput("citedArtifact.statusDate["+strconv.Itoa(numStatusDate)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("citedArtifact.statusDate["+strconv.Itoa(numStatusDate)+"].period", &resource.CitedArtifact.StatusDate[numStatusDate].Period, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactTitleType(numTitle int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTitle >= len(resource.CitedArtifact.Title) || numType >= len(resource.CitedArtifact.Title[numTitle].Type) {
		return CodeableConceptSelect("citedArtifact.title["+strconv.Itoa(numTitle)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.title["+strconv.Itoa(numTitle)+"].type["+strconv.Itoa(numType)+"]", &resource.CitedArtifact.Title[numTitle].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactAbstractType(numAbstract int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAbstract >= len(resource.CitedArtifact.Abstract) {
		return CodeableConceptSelect("citedArtifact.abstract["+strconv.Itoa(numAbstract)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.abstract["+strconv.Itoa(numAbstract)+"].type", resource.CitedArtifact.Abstract[numAbstract].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactAbstractCopyright(numAbstract int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAbstract >= len(resource.CitedArtifact.Abstract) {
		return StringInput("citedArtifact.abstract["+strconv.Itoa(numAbstract)+"].copyright", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.abstract["+strconv.Itoa(numAbstract)+"].copyright", resource.CitedArtifact.Abstract[numAbstract].Copyright, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPartType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("citedArtifact.part.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.part.type", resource.CitedArtifact.Part.Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPartValue(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("citedArtifact.part.value", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.part.value", resource.CitedArtifact.Part.Value, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPartBaseCitation(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("citedArtifact.part.baseCitation", nil, htmlAttrs)
	}
	return ReferenceInput("citedArtifact.part.baseCitation", resource.CitedArtifact.Part.BaseCitation, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToType(numRelatesTo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return CodeSelect("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].type", &resource.CitedArtifact.RelatesTo[numRelatesTo].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToClassifier(numRelatesTo int, numClassifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) || numClassifier >= len(resource.CitedArtifact.RelatesTo[numRelatesTo].Classifier) {
		return CodeableConceptSelect("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].classifier["+strconv.Itoa(numClassifier)+"]", &resource.CitedArtifact.RelatesTo[numRelatesTo].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToLabel(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return StringInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].label", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].label", resource.CitedArtifact.RelatesTo[numRelatesTo].Label, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToDisplay(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return StringInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].display", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].display", resource.CitedArtifact.RelatesTo[numRelatesTo].Display, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToCitation(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return StringInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].citation", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].citation", resource.CitedArtifact.RelatesTo[numRelatesTo].Citation, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToDocument(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return AttachmentInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].document", nil, htmlAttrs)
	}
	return AttachmentInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].document", resource.CitedArtifact.RelatesTo[numRelatesTo].Document, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToResource(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return StringInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].resource", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].resource", resource.CitedArtifact.RelatesTo[numRelatesTo].Resource, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactRelatesToResourceReference(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.CitedArtifact.RelatesTo) {
		return ReferenceInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].resourceReference", nil, htmlAttrs)
	}
	return ReferenceInput("citedArtifact.relatesTo["+strconv.Itoa(numRelatesTo)+"].resourceReference", resource.CitedArtifact.RelatesTo[numRelatesTo].ResourceReference, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormCitedMedium(numPublicationForm int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return CodeableConceptSelect("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].citedMedium", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].citedMedium", resource.CitedArtifact.PublicationForm[numPublicationForm].CitedMedium, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormVolume(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].volume", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].volume", resource.CitedArtifact.PublicationForm[numPublicationForm].Volume, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormIssue(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].issue", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].issue", resource.CitedArtifact.PublicationForm[numPublicationForm].Issue, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormArticleDate(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return FhirDateTimeInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].articleDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].articleDate", resource.CitedArtifact.PublicationForm[numPublicationForm].ArticleDate, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublicationDateText(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publicationDateText", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publicationDateText", resource.CitedArtifact.PublicationForm[numPublicationForm].PublicationDateText, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublicationDateSeason(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publicationDateSeason", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publicationDateSeason", resource.CitedArtifact.PublicationForm[numPublicationForm].PublicationDateSeason, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormLastRevisionDate(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return FhirDateTimeInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].lastRevisionDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].lastRevisionDate", resource.CitedArtifact.PublicationForm[numPublicationForm].LastRevisionDate, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormAccessionNumber(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].accessionNumber", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].accessionNumber", resource.CitedArtifact.PublicationForm[numPublicationForm].AccessionNumber, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPageString(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].pageString", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].pageString", resource.CitedArtifact.PublicationForm[numPublicationForm].PageString, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormFirstPage(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].firstPage", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].firstPage", resource.CitedArtifact.PublicationForm[numPublicationForm].FirstPage, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormLastPage(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].lastPage", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].lastPage", resource.CitedArtifact.PublicationForm[numPublicationForm].LastPage, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPageCount(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].pageCount", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].pageCount", resource.CitedArtifact.PublicationForm[numPublicationForm].PageCount, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormCopyright(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].copyright", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].copyright", resource.CitedArtifact.PublicationForm[numPublicationForm].Copyright, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInType(numPublicationForm int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return CodeableConceptSelect("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publishedIn.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publishedIn.type", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInTitle(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publishedIn.title", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publishedIn.title", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.Title, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInPublisher(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return ReferenceInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publishedIn.publisher", nil, htmlAttrs)
	}
	return ReferenceInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publishedIn.publisher", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.Publisher, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInPublisherLocation(numPublicationForm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPublicationForm >= len(resource.CitedArtifact.PublicationForm) {
		return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publishedIn.publisherLocation", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.publicationForm["+strconv.Itoa(numPublicationForm)+"].publishedIn.publisherLocation", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.PublisherLocation, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactWebLocationClassifier(numWebLocation int, numClassifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numWebLocation >= len(resource.CitedArtifact.WebLocation) || numClassifier >= len(resource.CitedArtifact.WebLocation[numWebLocation].Classifier) {
		return CodeableConceptSelect("citedArtifact.webLocation["+strconv.Itoa(numWebLocation)+"].classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.webLocation["+strconv.Itoa(numWebLocation)+"].classifier["+strconv.Itoa(numClassifier)+"]", &resource.CitedArtifact.WebLocation[numWebLocation].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactWebLocationUrl(numWebLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numWebLocation >= len(resource.CitedArtifact.WebLocation) {
		return StringInput("citedArtifact.webLocation["+strconv.Itoa(numWebLocation)+"].url", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.webLocation["+strconv.Itoa(numWebLocation)+"].url", resource.CitedArtifact.WebLocation[numWebLocation].Url, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactClassificationType(numClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassification >= len(resource.CitedArtifact.Classification) {
		return CodeableConceptSelect("citedArtifact.classification["+strconv.Itoa(numClassification)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.classification["+strconv.Itoa(numClassification)+"].type", resource.CitedArtifact.Classification[numClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactClassificationClassifier(numClassification int, numClassifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassification >= len(resource.CitedArtifact.Classification) || numClassifier >= len(resource.CitedArtifact.Classification[numClassification].Classifier) {
		return CodeableConceptSelect("citedArtifact.classification["+strconv.Itoa(numClassification)+"].classifier["+strconv.Itoa(numClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.classification["+strconv.Itoa(numClassification)+"].classifier["+strconv.Itoa(numClassifier)+"]", &resource.CitedArtifact.Classification[numClassification].Classifier[numClassifier], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactClassificationArtifactAssessment(numClassification int, numArtifactAssessment int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassification >= len(resource.CitedArtifact.Classification) || numArtifactAssessment >= len(resource.CitedArtifact.Classification[numClassification].ArtifactAssessment) {
		return ReferenceInput("citedArtifact.classification["+strconv.Itoa(numClassification)+"].artifactAssessment["+strconv.Itoa(numArtifactAssessment)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("citedArtifact.classification["+strconv.Itoa(numClassification)+"].artifactAssessment["+strconv.Itoa(numArtifactAssessment)+"]", &resource.CitedArtifact.Classification[numClassification].ArtifactAssessment[numArtifactAssessment], htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipComplete(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("citedArtifact.contributorship.complete", nil, htmlAttrs)
	}
	return BoolInput("citedArtifact.contributorship.complete", resource.CitedArtifact.Contributorship.Complete, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributor(numEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) {
		return ReferenceInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].contributor", nil, htmlAttrs)
	}
	return ReferenceInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].contributor", &resource.CitedArtifact.Contributorship.Entry[numEntry].Contributor, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryForenameInitials(numEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) {
		return StringInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].forenameInitials", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].forenameInitials", resource.CitedArtifact.Contributorship.Entry[numEntry].ForenameInitials, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryAffiliation(numEntry int, numAffiliation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) || numAffiliation >= len(resource.CitedArtifact.Contributorship.Entry[numEntry].Affiliation) {
		return ReferenceInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].affiliation["+strconv.Itoa(numAffiliation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].affiliation["+strconv.Itoa(numAffiliation)+"]", &resource.CitedArtifact.Contributorship.Entry[numEntry].Affiliation[numAffiliation], htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionType(numEntry int, numContributionType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) || numContributionType >= len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionType) {
		return CodeableConceptSelect("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].contributionType["+strconv.Itoa(numContributionType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].contributionType["+strconv.Itoa(numContributionType)+"]", &resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionType[numContributionType], optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryRole(numEntry int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) {
		return CodeableConceptSelect("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].role", resource.CitedArtifact.Contributorship.Entry[numEntry].Role, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryCorrespondingContact(numEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) {
		return BoolInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].correspondingContact", nil, htmlAttrs)
	}
	return BoolInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].correspondingContact", resource.CitedArtifact.Contributorship.Entry[numEntry].CorrespondingContact, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryRankingOrder(numEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) {
		return IntInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].rankingOrder", nil, htmlAttrs)
	}
	return IntInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].rankingOrder", resource.CitedArtifact.Contributorship.Entry[numEntry].RankingOrder, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionInstanceType(numEntry int, numContributionInstance int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) || numContributionInstance >= len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance) {
		return CodeableConceptSelect("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].contributionInstance["+strconv.Itoa(numContributionInstance)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].contributionInstance["+strconv.Itoa(numContributionInstance)+"].type", &resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance[numContributionInstance].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionInstanceTime(numEntry int, numContributionInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.CitedArtifact.Contributorship.Entry) || numContributionInstance >= len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance) {
		return FhirDateTimeInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].contributionInstance["+strconv.Itoa(numContributionInstance)+"].time", nil, htmlAttrs)
	}
	return FhirDateTimeInput("citedArtifact.contributorship.entry["+strconv.Itoa(numEntry)+"].contributionInstance["+strconv.Itoa(numContributionInstance)+"].time", resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance[numContributionInstance].Time, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryType(numSummary int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSummary >= len(resource.CitedArtifact.Contributorship.Summary) {
		return CodeableConceptSelect("citedArtifact.contributorship.summary["+strconv.Itoa(numSummary)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.contributorship.summary["+strconv.Itoa(numSummary)+"].type", resource.CitedArtifact.Contributorship.Summary[numSummary].Type, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryStyle(numSummary int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSummary >= len(resource.CitedArtifact.Contributorship.Summary) {
		return CodeableConceptSelect("citedArtifact.contributorship.summary["+strconv.Itoa(numSummary)+"].style", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.contributorship.summary["+strconv.Itoa(numSummary)+"].style", resource.CitedArtifact.Contributorship.Summary[numSummary].Style, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipSummarySource(numSummary int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSummary >= len(resource.CitedArtifact.Contributorship.Summary) {
		return CodeableConceptSelect("citedArtifact.contributorship.summary["+strconv.Itoa(numSummary)+"].source", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("citedArtifact.contributorship.summary["+strconv.Itoa(numSummary)+"].source", resource.CitedArtifact.Contributorship.Summary[numSummary].Source, optionsValueSet, htmlAttrs)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryValue(numSummary int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSummary >= len(resource.CitedArtifact.Contributorship.Summary) {
		return StringInput("citedArtifact.contributorship.summary["+strconv.Itoa(numSummary)+"].value", nil, htmlAttrs)
	}
	return StringInput("citedArtifact.contributorship.summary["+strconv.Itoa(numSummary)+"].value", &resource.CitedArtifact.Contributorship.Summary[numSummary].Value, htmlAttrs)
}
