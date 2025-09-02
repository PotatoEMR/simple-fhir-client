package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	VersionAlgorithmString *string                  `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                  `json:"versionAlgorithmCoding"`
	Name                   *string                  `json:"name,omitempty"`
	Title                  *string                  `json:"title,omitempty"`
	Status                 string                   `json:"status"`
	Experimental           *bool                    `json:"experimental,omitempty"`
	Date                   *string                  `json:"date,omitempty"`
	Publisher              *string                  `json:"publisher,omitempty"`
	Contact                []ContactDetail          `json:"contact,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	UseContext             []UsageContext           `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept        `json:"jurisdiction,omitempty"`
	Purpose                *string                  `json:"purpose,omitempty"`
	Copyright              *string                  `json:"copyright,omitempty"`
	CopyrightLabel         *string                  `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                  `json:"approvalDate,omitempty"`
	LastReviewDate         *string                  `json:"lastReviewDate,omitempty"`
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
	ArticleDate           *string                                          `json:"articleDate,omitempty"`
	PublicationDateText   *string                                          `json:"publicationDateText,omitempty"`
	PublicationDateSeason *string                                          `json:"publicationDateSeason,omitempty"`
	LastRevisionDate      *string                                          `json:"lastRevisionDate,omitempty"`
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
	Time              *string         `json:"time,omitempty"`
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

func (resource *Citation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Citation) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Citation) T_Jurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *Citation) T_CurrentState(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("currentState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("currentState", &resource.CurrentState[0], optionsValueSet)
}
func (resource *Citation) T_SummaryStyle(numSummary int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Summary) >= numSummary {
		return CodeableConceptSelect("style", nil, optionsValueSet)
	}
	return CodeableConceptSelect("style", resource.Summary[numSummary].Style, optionsValueSet)
}
func (resource *Citation) T_ClassificationType(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Classification[numClassification].Type, optionsValueSet)
}
func (resource *Citation) T_ClassificationClassifier(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("classifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classifier", &resource.Classification[numClassification].Classifier[0], optionsValueSet)
}
func (resource *Citation) T_StatusDateActivity(numStatusDate int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.StatusDate) >= numStatusDate {
		return CodeableConceptSelect("activity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("activity", &resource.StatusDate[numStatusDate].Activity, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactCurrentState(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("currentState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("currentState", &resource.CitedArtifact.CurrentState[0], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactStatusDateActivity(numStatusDate int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.StatusDate) >= numStatusDate {
		return CodeableConceptSelect("activity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("activity", &resource.CitedArtifact.StatusDate[numStatusDate].Activity, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactTitleType(numTitle int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Title) >= numTitle {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.CitedArtifact.Title[numTitle].Type[0], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactTitleLanguage(numTitle int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Title) >= numTitle {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", resource.CitedArtifact.Title[numTitle].Language, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactAbstractType(numAbstract int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Abstract) >= numAbstract {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.CitedArtifact.Abstract[numAbstract].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactAbstractLanguage(numAbstract int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Abstract) >= numAbstract {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", resource.CitedArtifact.Abstract[numAbstract].Language, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactPartType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.CitedArtifact.Part.Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactRelatesToType(numRelatesTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.RelatesTo) >= numRelatesTo {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.CitedArtifact.RelatesTo[numRelatesTo].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactRelatesToClassifier(numRelatesTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.RelatesTo) >= numRelatesTo {
		return CodeableConceptSelect("classifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classifier", &resource.CitedArtifact.RelatesTo[numRelatesTo].Classifier[0], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactPublicationFormCitedMedium(numPublicationForm int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return CodeableConceptSelect("citedMedium", nil, optionsValueSet)
	}
	return CodeableConceptSelect("citedMedium", resource.CitedArtifact.PublicationForm[numPublicationForm].CitedMedium, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactPublicationFormLanguage(numPublicationForm int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", &resource.CitedArtifact.PublicationForm[numPublicationForm].Language[0], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactPublicationFormPublishedInType(numPublicationForm int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.PublicationForm) >= numPublicationForm {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.CitedArtifact.PublicationForm[numPublicationForm].PublishedIn.Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactWebLocationClassifier(numWebLocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.WebLocation) >= numWebLocation {
		return CodeableConceptSelect("classifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classifier", &resource.CitedArtifact.WebLocation[numWebLocation].Classifier[0], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactClassificationType(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Classification) >= numClassification {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.CitedArtifact.Classification[numClassification].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactClassificationClassifier(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Classification) >= numClassification {
		return CodeableConceptSelect("classifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classifier", &resource.CitedArtifact.Classification[numClassification].Classifier[0], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionType(numEntry int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Contributorship.Entry) >= numEntry {
		return CodeableConceptSelect("contributionType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("contributionType", &resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionType[0], optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryRole(numEntry int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Contributorship.Entry) >= numEntry {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.CitedArtifact.Contributorship.Entry[numEntry].Role, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipEntryContributionInstanceType(numEntry int, numContributionInstance int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance) >= numContributionInstance {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.CitedArtifact.Contributorship.Entry[numEntry].ContributionInstance[numContributionInstance].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryType(numSummary int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Contributorship.Summary) >= numSummary {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.CitedArtifact.Contributorship.Summary[numSummary].Type, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipSummaryStyle(numSummary int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Contributorship.Summary) >= numSummary {
		return CodeableConceptSelect("style", nil, optionsValueSet)
	}
	return CodeableConceptSelect("style", resource.CitedArtifact.Contributorship.Summary[numSummary].Style, optionsValueSet)
}
func (resource *Citation) T_CitedArtifactContributorshipSummarySource(numSummary int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CitedArtifact.Contributorship.Summary) >= numSummary {
		return CodeableConceptSelect("source", nil, optionsValueSet)
	}
	return CodeableConceptSelect("source", resource.CitedArtifact.Contributorship.Summary[numSummary].Source, optionsValueSet)
}
