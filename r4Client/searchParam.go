package r4Client

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "net/url"

type SearchParam interface {
	SpEncode(baseURL *string) (*string, error)
} //search params for Account
type SpAccount struct {
	Owner      string
	Identifier string
	Period     string
	Subject    string
	Patient    string
	Name       string
	Type       string
	Status     string
}

func (p SpAccount) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"owner", p.Owner},
		{"identifier", p.Identifier},
		{"period", p.Period},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"name", p.Name},
		{"type", p.Type},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Account")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ActivityDefinition
type SpActivityDefinition struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ComposedOf          string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	Topic               string
	ContextTypeQuantity string
	Status              string
}

func (p SpActivityDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"successor", p.Successor},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"composedof", p.ComposedOf},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ActivityDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for AdverseEvent
type SpAdverseEvent struct {
	Date               string
	Severity           string
	Recorder           string
	Study              string
	Actuality          string
	Seriousness        string
	Subject            string
	Resultingcondition string
	Substance          string
	Location           string
	Category           string
	Event              string
}

func (p SpAdverseEvent) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"severity", p.Severity},
		{"recorder", p.Recorder},
		{"study", p.Study},
		{"actuality", p.Actuality},
		{"seriousness", p.Seriousness},
		{"subject", p.Subject},
		{"resultingcondition", p.Resultingcondition},
		{"substance", p.Substance},
		{"location", p.Location},
		{"category", p.Category},
		{"event", p.Event},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "AdverseEvent")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for AllergyIntolerance
type SpAllergyIntolerance struct {
	Severity           string
	Date               string
	Identifier         string
	Manifestation      string
	Recorder           string
	Code               string
	VerificationStatus string
	Criticality        string
	ClinicalStatus     string
	Type               string
	Onset              string
	Route              string
	Asserter           string
	Patient            string
	Category           string
	LastDate           string
}

func (p SpAllergyIntolerance) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"severity", p.Severity},
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"manifestation", p.Manifestation},
		{"recorder", p.Recorder},
		{"code", p.Code},
		{"verificationstatus", p.VerificationStatus},
		{"criticality", p.Criticality},
		{"clinicalstatus", p.ClinicalStatus},
		{"type", p.Type},
		{"onset", p.Onset},
		{"route", p.Route},
		{"asserter", p.Asserter},
		{"patient", p.Patient},
		{"category", p.Category},
		{"lastdate", p.LastDate},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "AllergyIntolerance")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Appointment
type SpAppointment struct {
	Date            string
	Identifier      string
	Specialty       string
	ServiceCategory string
	Practitioner    string
	PartStatus      string
	AppointmentType string
	ServiceType     string
	Slot            string
	ReasonCode      string
	Actor           string
	BasedOn         string
	Patient         string
	ReasonReference string
	SupportingInfo  string
	Location        string
	Status          string
}

func (p SpAppointment) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"specialty", p.Specialty},
		{"servicecategory", p.ServiceCategory},
		{"practitioner", p.Practitioner},
		{"partstatus", p.PartStatus},
		{"appointmenttype", p.AppointmentType},
		{"servicetype", p.ServiceType},
		{"slot", p.Slot},
		{"reasoncode", p.ReasonCode},
		{"actor", p.Actor},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"reasonreference", p.ReasonReference},
		{"supportinginfo", p.SupportingInfo},
		{"location", p.Location},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Appointment")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for AppointmentResponse
type SpAppointmentResponse struct {
	Actor        string
	Identifier   string
	Practitioner string
	PartStatus   string
	Patient      string
	Appointment  string
	Location     string
}

func (p SpAppointmentResponse) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"actor", p.Actor},
		{"identifier", p.Identifier},
		{"practitioner", p.Practitioner},
		{"partstatus", p.PartStatus},
		{"patient", p.Patient},
		{"appointment", p.Appointment},
		{"location", p.Location},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "AppointmentResponse")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for AuditEvent
type SpAuditEvent struct {
	Date       string
	EntityType string
	Agent      string
	Address    string
	EntityRole string
	Source     string
	Type       string
	Altid      string
	Site       string
	AgentName  string
	EntityName string
	Subtype    string
	Patient    string
	Action     string
	AgentRole  string
	Entity     string
	Outcome    string
	Policy     string
}

func (p SpAuditEvent) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"entitytype", p.EntityType},
		{"agent", p.Agent},
		{"address", p.Address},
		{"entityrole", p.EntityRole},
		{"source", p.Source},
		{"type", p.Type},
		{"altid", p.Altid},
		{"site", p.Site},
		{"agentname", p.AgentName},
		{"entityname", p.EntityName},
		{"subtype", p.Subtype},
		{"patient", p.Patient},
		{"action", p.Action},
		{"agentrole", p.AgentRole},
		{"entity", p.Entity},
		{"outcome", p.Outcome},
		{"policy", p.Policy},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "AuditEvent")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Basic
type SpBasic struct {
	Identifier string
	Code       string
	Subject    string
	Created    string
	Patient    string
	Author     string
}

func (p SpBasic) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"subject", p.Subject},
		{"created", p.Created},
		{"patient", p.Patient},
		{"author", p.Author},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Basic")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Binary
type SpBinary struct {
}

func (p SpBinary) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Binary")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for BiologicallyDerivedProduct
type SpBiologicallyDerivedProduct struct {
}

func (p SpBiologicallyDerivedProduct) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "BiologicallyDerivedProduct")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for BodyStructure
type SpBodyStructure struct {
	Identifier string
	Morphology string
	Patient    string
	Location   string
}

func (p SpBodyStructure) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"morphology", p.Morphology},
		{"patient", p.Patient},
		{"location", p.Location},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "BodyStructure")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Bundle
type SpBundle struct {
	Identifier  string
	Composition string
	Type        string
	Message     string
	Timestamp   string
}

func (p SpBundle) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"composition", p.Composition},
		{"type", p.Type},
		{"message", p.Message},
		{"timestamp", p.Timestamp},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Bundle")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for CapabilityStatement
type SpCapabilityStatement struct {
	Date                string
	ResourceProfile     string
	ContextTypeValue    string
	Software            string
	Resource            string
	Jurisdiction        string
	Format              string
	Description         string
	ContextType         string
	Title               string
	Fhirversion         string
	Version             string
	Url                 string
	SupportedProfile    string
	Mode                string
	ContextQuantity     string
	SecurityService     string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Guide               string
	Status              string
}

func (p SpCapabilityStatement) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"resourceprofile", p.ResourceProfile},
		{"contexttypevalue", p.ContextTypeValue},
		{"software", p.Software},
		{"resource", p.Resource},
		{"jurisdiction", p.Jurisdiction},
		{"format", p.Format},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"fhirversion", p.Fhirversion},
		{"version", p.Version},
		{"url", p.Url},
		{"supportedprofile", p.SupportedProfile},
		{"mode", p.Mode},
		{"contextquantity", p.ContextQuantity},
		{"securityservice", p.SecurityService},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"guide", p.Guide},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "CapabilityStatement")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for CarePlan
type SpCarePlan struct {
	Date                  string
	CareTeam              string
	Identifier            string
	Performer             string
	Goal                  string
	Subject               string
	Replaces              string
	InstantiatesCanonical string
	PartOf                string
	Encounter             string
	Intent                string
	ActivityReference     string
	Condition             string
	BasedOn               string
	Patient               string
	ActivityDate          string
	InstantiatesUri       string
	Category              string
	ActivityCode          string
	Status                string
}

func (p SpCarePlan) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"careteam", p.CareTeam},
		{"identifier", p.Identifier},
		{"performer", p.Performer},
		{"goal", p.Goal},
		{"subject", p.Subject},
		{"replaces", p.Replaces},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"partof", p.PartOf},
		{"encounter", p.Encounter},
		{"intent", p.Intent},
		{"activityreference", p.ActivityReference},
		{"condition", p.Condition},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"activitydate", p.ActivityDate},
		{"instantiatesuri", p.InstantiatesUri},
		{"category", p.Category},
		{"activitycode", p.ActivityCode},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "CarePlan")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for CareTeam
type SpCareTeam struct {
	Date        string
	Identifier  string
	Patient     string
	Subject     string
	Encounter   string
	Category    string
	Participant string
	Status      string
}

func (p SpCareTeam) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"encounter", p.Encounter},
		{"category", p.Category},
		{"participant", p.Participant},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "CareTeam")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for CatalogEntry
type SpCatalogEntry struct {
}

func (p SpCatalogEntry) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "CatalogEntry")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ChargeItem
type SpChargeItem struct {
	Identifier             string
	PerformingOrganization string
	Code                   string
	Quantity               string
	Subject                string
	Occurrence             string
	EnteredDate            string
	PerformerFunction      string
	Patient                string
	FactorOverride         string
	Service                string
	PriceOverride          string
	Context                string
	Enterer                string
	PerformerActor         string
	Account                string
	RequestingOrganization string
}

func (p SpChargeItem) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"performingorganization", p.PerformingOrganization},
		{"code", p.Code},
		{"quantity", p.Quantity},
		{"subject", p.Subject},
		{"occurrence", p.Occurrence},
		{"entereddate", p.EnteredDate},
		{"performerfunction", p.PerformerFunction},
		{"patient", p.Patient},
		{"factoroverride", p.FactorOverride},
		{"service", p.Service},
		{"priceoverride", p.PriceOverride},
		{"context", p.Context},
		{"enterer", p.Enterer},
		{"performeractor", p.PerformerActor},
		{"account", p.Account},
		{"requestingorganization", p.RequestingOrganization},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ChargeItem")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ChargeItemDefinition
type SpChargeItemDefinition struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpChargeItemDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ChargeItemDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Claim
type SpClaim struct {
	CareTeam     string
	Identifier   string
	Use          string
	Created      string
	Encounter    string
	Priority     string
	Payee        string
	Provider     string
	Patient      string
	Insurer      string
	DetailUdi    string
	Enterer      string
	ProcedureUdi string
	SubdetailUdi string
	Facility     string
	ItemUdi      string
	Status       string
}

func (p SpClaim) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"careteam", p.CareTeam},
		{"identifier", p.Identifier},
		{"use", p.Use},
		{"created", p.Created},
		{"encounter", p.Encounter},
		{"priority", p.Priority},
		{"payee", p.Payee},
		{"provider", p.Provider},
		{"patient", p.Patient},
		{"insurer", p.Insurer},
		{"detailudi", p.DetailUdi},
		{"enterer", p.Enterer},
		{"procedureudi", p.ProcedureUdi},
		{"subdetailudi", p.SubdetailUdi},
		{"facility", p.Facility},
		{"itemudi", p.ItemUdi},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Claim")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ClaimResponse
type SpClaimResponse struct {
	Identifier  string
	Request     string
	Disposition string
	Insurer     string
	Created     string
	Patient     string
	Use         string
	PaymentDate string
	Outcome     string
	Requestor   string
	Status      string
}

func (p SpClaimResponse) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"request", p.Request},
		{"disposition", p.Disposition},
		{"insurer", p.Insurer},
		{"created", p.Created},
		{"patient", p.Patient},
		{"use", p.Use},
		{"paymentdate", p.PaymentDate},
		{"outcome", p.Outcome},
		{"requestor", p.Requestor},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ClaimResponse")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ClinicalImpression
type SpClinicalImpression struct {
	Date           string
	Identifier     string
	Previous       string
	FindingCode    string
	Assessor       string
	Subject        string
	Encounter      string
	FindingRef     string
	Problem        string
	Patient        string
	SupportingInfo string
	Investigation  string
	Status         string
}

func (p SpClinicalImpression) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"previous", p.Previous},
		{"findingcode", p.FindingCode},
		{"assessor", p.Assessor},
		{"subject", p.Subject},
		{"encounter", p.Encounter},
		{"findingref", p.FindingRef},
		{"problem", p.Problem},
		{"patient", p.Patient},
		{"supportinginfo", p.SupportingInfo},
		{"investigation", p.Investigation},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ClinicalImpression")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for CodeSystem
type SpCodeSystem struct {
	Date                string
	Identifier          string
	Code                string
	ContextTypeValue    string
	ContentMode         string
	Jurisdiction        string
	Description         string
	ContextType         string
	Language            string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Supplements         string
	System              string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpCodeSystem) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"contexttypevalue", p.ContextTypeValue},
		{"contentmode", p.ContentMode},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"language", p.Language},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"supplements", p.Supplements},
		{"system", p.System},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "CodeSystem")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Communication
type SpCommunication struct {
	Identifier            string
	Subject               string
	InstantiatesCanonical string
	Received              string
	PartOf                string
	Medium                string
	Encounter             string
	Sent                  string
	BasedOn               string
	Sender                string
	Patient               string
	Recipient             string
	InstantiatesUri       string
	Category              string
	Status                string
}

func (p SpCommunication) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"subject", p.Subject},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"received", p.Received},
		{"partof", p.PartOf},
		{"medium", p.Medium},
		{"encounter", p.Encounter},
		{"sent", p.Sent},
		{"basedon", p.BasedOn},
		{"sender", p.Sender},
		{"patient", p.Patient},
		{"recipient", p.Recipient},
		{"instantiatesuri", p.InstantiatesUri},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Communication")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for CommunicationRequest
type SpCommunicationRequest struct {
	Requester       string
	Authored        string
	Identifier      string
	Subject         string
	Replaces        string
	Medium          string
	Encounter       string
	Occurrence      string
	Priority        string
	GroupIdentifier string
	BasedOn         string
	Sender          string
	Patient         string
	Recipient       string
	Category        string
	Status          string
}

func (p SpCommunicationRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"requester", p.Requester},
		{"authored", p.Authored},
		{"identifier", p.Identifier},
		{"subject", p.Subject},
		{"replaces", p.Replaces},
		{"medium", p.Medium},
		{"encounter", p.Encounter},
		{"occurrence", p.Occurrence},
		{"priority", p.Priority},
		{"groupidentifier", p.GroupIdentifier},
		{"basedon", p.BasedOn},
		{"sender", p.Sender},
		{"patient", p.Patient},
		{"recipient", p.Recipient},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "CommunicationRequest")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for CompartmentDefinition
type SpCompartmentDefinition struct {
	Date                string
	Code                string
	ContextTypeValue    string
	Resource            string
	Description         string
	ContextType         string
	Version             string
	Url                 string
	ContextQuantity     string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpCompartmentDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"code", p.Code},
		{"contexttypevalue", p.ContextTypeValue},
		{"resource", p.Resource},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "CompartmentDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Composition
type SpComposition struct {
	Date            string
	Identifier      string
	Period          string
	RelatedId       string
	Subject         string
	Author          string
	Confidentiality string
	Section         string
	Encounter       string
	Type            string
	Title           string
	Attester        string
	Entry           string
	RelatedRef      string
	Patient         string
	Context         string
	Category        string
	Status          string
}

func (p SpComposition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"period", p.Period},
		{"relatedid", p.RelatedId},
		{"subject", p.Subject},
		{"author", p.Author},
		{"confidentiality", p.Confidentiality},
		{"section", p.Section},
		{"encounter", p.Encounter},
		{"type", p.Type},
		{"title", p.Title},
		{"attester", p.Attester},
		{"entry", p.Entry},
		{"relatedref", p.RelatedRef},
		{"patient", p.Patient},
		{"context", p.Context},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Composition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ConceptMap
type SpConceptMap struct {
	Date                string
	Other               string
	ContextTypeValue    string
	TargetSystem        string
	Dependson           string
	Jurisdiction        string
	Description         string
	ContextType         string
	Source              string
	Title               string
	ContextQuantity     string
	SourceUri           string
	Context             string
	ContextTypeQuantity string
	SourceSystem        string
	TargetCode          string
	TargetUri           string
	Identifier          string
	Product             string
	Version             string
	Url                 string
	Target              string
	SourceCode          string
	Name                string
	Publisher           string
	Status              string
}

func (p SpConceptMap) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"other", p.Other},
		{"contexttypevalue", p.ContextTypeValue},
		{"targetsystem", p.TargetSystem},
		{"dependson", p.Dependson},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"source", p.Source},
		{"title", p.Title},
		{"contextquantity", p.ContextQuantity},
		{"sourceuri", p.SourceUri},
		{"context", p.Context},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"sourcesystem", p.SourceSystem},
		{"targetcode", p.TargetCode},
		{"targeturi", p.TargetUri},
		{"identifier", p.Identifier},
		{"product", p.Product},
		{"version", p.Version},
		{"url", p.Url},
		{"target", p.Target},
		{"sourcecode", p.SourceCode},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ConceptMap")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Condition
type SpCondition struct {
	Severity           string
	EvidenceDetail     string
	Identifier         string
	OnsetInfo          string
	RecordedDate       string
	Code               string
	Evidence           string
	Subject            string
	VerificationStatus string
	ClinicalStatus     string
	Encounter          string
	OnsetDate          string
	AbatementDate      string
	Asserter           string
	Stage              string
	AbatementString    string
	Patient            string
	OnsetAge           string
	AbatementAge       string
	Category           string
	BodySite           string
}

func (p SpCondition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"severity", p.Severity},
		{"evidencedetail", p.EvidenceDetail},
		{"identifier", p.Identifier},
		{"onsetinfo", p.OnsetInfo},
		{"recordeddate", p.RecordedDate},
		{"code", p.Code},
		{"evidence", p.Evidence},
		{"subject", p.Subject},
		{"verificationstatus", p.VerificationStatus},
		{"clinicalstatus", p.ClinicalStatus},
		{"encounter", p.Encounter},
		{"onsetdate", p.OnsetDate},
		{"abatementdate", p.AbatementDate},
		{"asserter", p.Asserter},
		{"stage", p.Stage},
		{"abatementstring", p.AbatementString},
		{"patient", p.Patient},
		{"onsetage", p.OnsetAge},
		{"abatementage", p.AbatementAge},
		{"category", p.Category},
		{"bodysite", p.BodySite},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Condition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Consent
type SpConsent struct {
	Date            string
	Identifier      string
	Period          string
	Data            string
	Purpose         string
	SourceReference string
	Actor           string
	SecurityLabel   string
	Patient         string
	Organization    string
	Scope           string
	Action          string
	Consentor       string
	Category        string
	Status          string
}

func (p SpConsent) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"period", p.Period},
		{"data", p.Data},
		{"purpose", p.Purpose},
		{"sourcereference", p.SourceReference},
		{"actor", p.Actor},
		{"securitylabel", p.SecurityLabel},
		{"patient", p.Patient},
		{"organization", p.Organization},
		{"scope", p.Scope},
		{"action", p.Action},
		{"consentor", p.Consentor},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Consent")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Contract
type SpContract struct {
	Identifier   string
	Instantiates string
	Patient      string
	Subject      string
	Authority    string
	Domain       string
	Issued       string
	Url          string
	Signer       string
	Status       string
}

func (p SpContract) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"instantiates", p.Instantiates},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"authority", p.Authority},
		{"domain", p.Domain},
		{"issued", p.Issued},
		{"url", p.Url},
		{"signer", p.Signer},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Contract")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Coverage
type SpCoverage struct {
	Identifier   string
	Payor        string
	Subscriber   string
	Beneficiary  string
	Patient      string
	ClassValue   string
	Type         string
	Dependent    string
	ClassType    string
	PolicyHolder string
	Status       string
}

func (p SpCoverage) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"payor", p.Payor},
		{"subscriber", p.Subscriber},
		{"beneficiary", p.Beneficiary},
		{"patient", p.Patient},
		{"classvalue", p.ClassValue},
		{"type", p.Type},
		{"dependent", p.Dependent},
		{"classtype", p.ClassType},
		{"policyholder", p.PolicyHolder},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Coverage")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for CoverageEligibilityRequest
type SpCoverageEligibilityRequest struct {
	Identifier string
	Provider   string
	Patient    string
	Created    string
	Enterer    string
	Facility   string
	Status     string
}

func (p SpCoverageEligibilityRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"provider", p.Provider},
		{"patient", p.Patient},
		{"created", p.Created},
		{"enterer", p.Enterer},
		{"facility", p.Facility},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "CoverageEligibilityRequest")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for CoverageEligibilityResponse
type SpCoverageEligibilityResponse struct {
	Identifier  string
	Request     string
	Disposition string
	Patient     string
	Insurer     string
	Created     string
	Outcome     string
	Requestor   string
	Status      string
}

func (p SpCoverageEligibilityResponse) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"request", p.Request},
		{"disposition", p.Disposition},
		{"patient", p.Patient},
		{"insurer", p.Insurer},
		{"created", p.Created},
		{"outcome", p.Outcome},
		{"requestor", p.Requestor},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "CoverageEligibilityResponse")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for DetectedIssue
type SpDetectedIssue struct {
	Identifier string
	Code       string
	Identified string
	Patient    string
	Author     string
	Implicated string
}

func (p SpDetectedIssue) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"identified", p.Identified},
		{"patient", p.Patient},
		{"author", p.Author},
		{"implicated", p.Implicated},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DetectedIssue")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Device
type SpDevice struct {
	UdiDi        string
	Identifier   string
	UdiCarrier   string
	DeviceName   string
	Patient      string
	Organization string
	Model        string
	Location     string
	Type         string
	Url          string
	Manufacturer string
	Status       string
}

func (p SpDevice) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"udidi", p.UdiDi},
		{"identifier", p.Identifier},
		{"udicarrier", p.UdiCarrier},
		{"devicename", p.DeviceName},
		{"patient", p.Patient},
		{"organization", p.Organization},
		{"model", p.Model},
		{"location", p.Location},
		{"type", p.Type},
		{"url", p.Url},
		{"manufacturer", p.Manufacturer},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Device")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for DeviceDefinition
type SpDeviceDefinition struct {
	Parent     string
	Identifier string
	Type       string
}

func (p SpDeviceDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"parent", p.Parent},
		{"identifier", p.Identifier},
		{"type", p.Type},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DeviceDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for DeviceMetric
type SpDeviceMetric struct {
	Parent     string
	Identifier string
	Source     string
	Type       string
	Category   string
}

func (p SpDeviceMetric) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"parent", p.Parent},
		{"identifier", p.Identifier},
		{"source", p.Source},
		{"type", p.Type},
		{"category", p.Category},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DeviceMetric")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for DeviceRequest
type SpDeviceRequest struct {
	Requester             string
	Insurance             string
	Identifier            string
	Code                  string
	Performer             string
	EventDate             string
	Subject               string
	InstantiatesCanonical string
	Encounter             string
	AuthoredOn            string
	Intent                string
	GroupIdentifier       string
	BasedOn               string
	Patient               string
	InstantiatesUri       string
	PriorRequest          string
	Device                string
	Status                string
}

func (p SpDeviceRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"requester", p.Requester},
		{"insurance", p.Insurance},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"performer", p.Performer},
		{"eventdate", p.EventDate},
		{"subject", p.Subject},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"encounter", p.Encounter},
		{"authoredon", p.AuthoredOn},
		{"intent", p.Intent},
		{"groupidentifier", p.GroupIdentifier},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"instantiatesuri", p.InstantiatesUri},
		{"priorrequest", p.PriorRequest},
		{"device", p.Device},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DeviceRequest")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for DeviceUseStatement
type SpDeviceUseStatement struct {
	Identifier string
	Subject    string
	Patient    string
	Device     string
}

func (p SpDeviceUseStatement) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"device", p.Device},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DeviceUseStatement")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for DiagnosticReport
type SpDiagnosticReport struct {
	Date               string
	Identifier         string
	Performer          string
	Code               string
	Subject            string
	Media              string
	Encounter          string
	Result             string
	Conclusion         string
	BasedOn            string
	Patient            string
	Specimen           string
	Issued             string
	Category           string
	ResultsInterpreter string
	Status             string
}

func (p SpDiagnosticReport) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"performer", p.Performer},
		{"code", p.Code},
		{"subject", p.Subject},
		{"media", p.Media},
		{"encounter", p.Encounter},
		{"result", p.Result},
		{"conclusion", p.Conclusion},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"specimen", p.Specimen},
		{"issued", p.Issued},
		{"category", p.Category},
		{"resultsinterpreter", p.ResultsInterpreter},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DiagnosticReport")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for DocumentManifest
type SpDocumentManifest struct {
	Identifier  string
	Item        string
	RelatedId   string
	Subject     string
	Author      string
	Created     string
	Description string
	Source      string
	Type        string
	RelatedRef  string
	Patient     string
	Recipient   string
	Status      string
}

func (p SpDocumentManifest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"item", p.Item},
		{"relatedid", p.RelatedId},
		{"subject", p.Subject},
		{"author", p.Author},
		{"created", p.Created},
		{"description", p.Description},
		{"source", p.Source},
		{"type", p.Type},
		{"relatedref", p.RelatedRef},
		{"patient", p.Patient},
		{"recipient", p.Recipient},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DocumentManifest")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for DocumentReference
type SpDocumentReference struct {
	Date          string
	Subject       string
	Description   string
	Language      string
	Type          string
	Relation      string
	Setting       string
	Related       string
	Patient       string
	Relationship  string
	Event         string
	Authenticator string
	Identifier    string
	Period        string
	Custodian     string
	Author        string
	Format        string
	Encounter     string
	Contenttype   string
	SecurityLabel string
	Location      string
	Category      string
	Relatesto     string
	Facility      string
	Status        string
}

func (p SpDocumentReference) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"subject", p.Subject},
		{"description", p.Description},
		{"language", p.Language},
		{"type", p.Type},
		{"relation", p.Relation},
		{"setting", p.Setting},
		{"related", p.Related},
		{"patient", p.Patient},
		{"relationship", p.Relationship},
		{"event", p.Event},
		{"authenticator", p.Authenticator},
		{"identifier", p.Identifier},
		{"period", p.Period},
		{"custodian", p.Custodian},
		{"author", p.Author},
		{"format", p.Format},
		{"encounter", p.Encounter},
		{"contenttype", p.Contenttype},
		{"securitylabel", p.SecurityLabel},
		{"location", p.Location},
		{"category", p.Category},
		{"relatesto", p.Relatesto},
		{"facility", p.Facility},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DocumentReference")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for EffectEvidenceSynthesis
type SpEffectEvidenceSynthesis struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpEffectEvidenceSynthesis) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "EffectEvidenceSynthesis")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Encounter
type SpEncounter struct {
	Date               string
	Identifier         string
	ParticipantType    string
	Practitioner       string
	Subject            string
	Length             string
	EpisodeOfCare      string
	Diagnosis          string
	Appointment        string
	PartOf             string
	Type               string
	ReasonCode         string
	Participant        string
	BasedOn            string
	Patient            string
	ReasonReference    string
	LocationPeriod     string
	Location           string
	ServiceProvider    string
	SpecialArrangement string
	Class              string
	Account            string
	Status             string
}

func (p SpEncounter) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"participanttype", p.ParticipantType},
		{"practitioner", p.Practitioner},
		{"subject", p.Subject},
		{"length", p.Length},
		{"episodeofcare", p.EpisodeOfCare},
		{"diagnosis", p.Diagnosis},
		{"appointment", p.Appointment},
		{"partof", p.PartOf},
		{"type", p.Type},
		{"reasoncode", p.ReasonCode},
		{"participant", p.Participant},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"reasonreference", p.ReasonReference},
		{"locationperiod", p.LocationPeriod},
		{"location", p.Location},
		{"serviceprovider", p.ServiceProvider},
		{"specialarrangement", p.SpecialArrangement},
		{"class", p.Class},
		{"account", p.Account},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Encounter")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Endpoint
type SpEndpoint struct {
	PayloadType    string
	Identifier     string
	Organization   string
	ConnectionType string
	Name           string
	Status         string
}

func (p SpEndpoint) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"payloadtype", p.PayloadType},
		{"identifier", p.Identifier},
		{"organization", p.Organization},
		{"connectiontype", p.ConnectionType},
		{"name", p.Name},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Endpoint")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for EnrollmentRequest
type SpEnrollmentRequest struct {
	Identifier string
	Subject    string
	Patient    string
	Status     string
}

func (p SpEnrollmentRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "EnrollmentRequest")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for EnrollmentResponse
type SpEnrollmentResponse struct {
	Identifier string
	Request    string
	Status     string
}

func (p SpEnrollmentResponse) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"request", p.Request},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "EnrollmentResponse")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for EpisodeOfCare
type SpEpisodeOfCare struct {
	Date             string
	Identifier       string
	Condition        string
	Patient          string
	Organization     string
	Type             string
	CareManager      string
	Status           string
	IncomingReferral string
}

func (p SpEpisodeOfCare) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"condition", p.Condition},
		{"patient", p.Patient},
		{"organization", p.Organization},
		{"type", p.Type},
		{"caremanager", p.CareManager},
		{"status", p.Status},
		{"incomingreferral", p.IncomingReferral},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "EpisodeOfCare")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for EventDefinition
type SpEventDefinition struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ComposedOf          string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	Topic               string
	ContextTypeQuantity string
	Status              string
}

func (p SpEventDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"successor", p.Successor},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"composedof", p.ComposedOf},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "EventDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Evidence
type SpEvidence struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ComposedOf          string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	Topic               string
	ContextTypeQuantity string
	Status              string
}

func (p SpEvidence) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"successor", p.Successor},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"composedof", p.ComposedOf},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Evidence")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for EvidenceVariable
type SpEvidenceVariable struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ComposedOf          string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	Topic               string
	ContextTypeQuantity string
	Status              string
}

func (p SpEvidenceVariable) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"successor", p.Successor},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"composedof", p.ComposedOf},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "EvidenceVariable")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ExampleScenario
type SpExampleScenario struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	ContextType         string
	Version             string
	Url                 string
	ContextQuantity     string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpExampleScenario) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"contexttype", p.ContextType},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ExampleScenario")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ExplanationOfBenefit
type SpExplanationOfBenefit struct {
	Coverage     string
	CareTeam     string
	Identifier   string
	Created      string
	Encounter    string
	Payee        string
	Disposition  string
	Provider     string
	Patient      string
	DetailUdi    string
	Claim        string
	Enterer      string
	ProcedureUdi string
	SubdetailUdi string
	Facility     string
	ItemUdi      string
	Status       string
}

func (p SpExplanationOfBenefit) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"coverage", p.Coverage},
		{"careteam", p.CareTeam},
		{"identifier", p.Identifier},
		{"created", p.Created},
		{"encounter", p.Encounter},
		{"payee", p.Payee},
		{"disposition", p.Disposition},
		{"provider", p.Provider},
		{"patient", p.Patient},
		{"detailudi", p.DetailUdi},
		{"claim", p.Claim},
		{"enterer", p.Enterer},
		{"procedureudi", p.ProcedureUdi},
		{"subdetailudi", p.SubdetailUdi},
		{"facility", p.Facility},
		{"itemudi", p.ItemUdi},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ExplanationOfBenefit")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for FamilyMemberHistory
type SpFamilyMemberHistory struct {
	Date                  string
	Identifier            string
	Code                  string
	Patient               string
	Sex                   string
	InstantiatesCanonical string
	InstantiatesUri       string
	Relationship          string
	Status                string
}

func (p SpFamilyMemberHistory) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"patient", p.Patient},
		{"sex", p.Sex},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"instantiatesuri", p.InstantiatesUri},
		{"relationship", p.Relationship},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "FamilyMemberHistory")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Flag
type SpFlag struct {
	Date       string
	Identifier string
	Subject    string
	Patient    string
	Author     string
	Encounter  string
}

func (p SpFlag) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"author", p.Author},
		{"encounter", p.Encounter},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Flag")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Goal
type SpGoal struct {
	Identifier        string
	LifecycleStatus   string
	AchievementStatus string
	Patient           string
	Subject           string
	StartDate         string
	Category          string
	TargetDate        string
}

func (p SpGoal) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"lifecyclestatus", p.LifecycleStatus},
		{"achievementstatus", p.AchievementStatus},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"startdate", p.StartDate},
		{"category", p.Category},
		{"targetdate", p.TargetDate},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Goal")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for GraphDefinition
type SpGraphDefinition struct {
	Date                string
	ContextTypeValue    string
	Jurisdiction        string
	Start               string
	Description         string
	ContextType         string
	Version             string
	Url                 string
	ContextQuantity     string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpGraphDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"start", p.Start},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "GraphDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Group
type SpGroup struct {
	Actual              string
	Identifier          string
	CharacteristicValue string
	ManagingEntity      string
	Code                string
	Member              string
	Exclude             string
	Type                string
	Value               string
	Characteristic      string
}

func (p SpGroup) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"actual", p.Actual},
		{"identifier", p.Identifier},
		{"characteristicvalue", p.CharacteristicValue},
		{"managingentity", p.ManagingEntity},
		{"code", p.Code},
		{"member", p.Member},
		{"exclude", p.Exclude},
		{"type", p.Type},
		{"value", p.Value},
		{"characteristic", p.Characteristic},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Group")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for GuidanceResponse
type SpGuidanceResponse struct {
	Request    string
	Identifier string
	Patient    string
	Subject    string
}

func (p SpGuidanceResponse) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"request", p.Request},
		{"identifier", p.Identifier},
		{"patient", p.Patient},
		{"subject", p.Subject},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "GuidanceResponse")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for HealthcareService
type SpHealthcareService struct {
	Identifier      string
	Specialty       string
	Endpoint        string
	ServiceCategory string
	CoverageArea    string
	ServiceType     string
	Organization    string
	Name            string
	Active          string
	Location        string
	Program         string
	Characteristic  string
}

func (p SpHealthcareService) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"specialty", p.Specialty},
		{"endpoint", p.Endpoint},
		{"servicecategory", p.ServiceCategory},
		{"coveragearea", p.CoverageArea},
		{"servicetype", p.ServiceType},
		{"organization", p.Organization},
		{"name", p.Name},
		{"active", p.Active},
		{"location", p.Location},
		{"program", p.Program},
		{"characteristic", p.Characteristic},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "HealthcareService")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ImagingStudy
type SpImagingStudy struct {
	Identifier  string
	Reason      string
	DicomClass  string
	Modality    string
	Bodysite    string
	Instance    string
	Performer   string
	Subject     string
	Started     string
	Interpreter string
	Encounter   string
	Referrer    string
	Endpoint    string
	Patient     string
	Series      string
	Basedon     string
	Status      string
}

func (p SpImagingStudy) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"reason", p.Reason},
		{"dicomclass", p.DicomClass},
		{"modality", p.Modality},
		{"bodysite", p.Bodysite},
		{"instance", p.Instance},
		{"performer", p.Performer},
		{"subject", p.Subject},
		{"started", p.Started},
		{"interpreter", p.Interpreter},
		{"encounter", p.Encounter},
		{"referrer", p.Referrer},
		{"endpoint", p.Endpoint},
		{"patient", p.Patient},
		{"series", p.Series},
		{"basedon", p.Basedon},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ImagingStudy")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Immunization
type SpImmunization struct {
	Date            string
	Identifier      string
	Performer       string
	Reaction        string
	LotNumber       string
	StatusReason    string
	ReasonCode      string
	Manufacturer    string
	TargetDisease   string
	Patient         string
	Series          string
	VaccineCode     string
	ReasonReference string
	Location        string
	Status          string
	ReactionDate    string
}

func (p SpImmunization) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"performer", p.Performer},
		{"reaction", p.Reaction},
		{"lotnumber", p.LotNumber},
		{"statusreason", p.StatusReason},
		{"reasoncode", p.ReasonCode},
		{"manufacturer", p.Manufacturer},
		{"targetdisease", p.TargetDisease},
		{"patient", p.Patient},
		{"series", p.Series},
		{"vaccinecode", p.VaccineCode},
		{"reasonreference", p.ReasonReference},
		{"location", p.Location},
		{"status", p.Status},
		{"reactiondate", p.ReactionDate},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Immunization")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ImmunizationEvaluation
type SpImmunizationEvaluation struct {
	Date              string
	Identifier        string
	TargetDisease     string
	Patient           string
	DoseStatus        string
	ImmunizationEvent string
	Status            string
}

func (p SpImmunizationEvaluation) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"targetdisease", p.TargetDisease},
		{"patient", p.Patient},
		{"dosestatus", p.DoseStatus},
		{"immunizationevent", p.ImmunizationEvent},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ImmunizationEvaluation")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ImmunizationRecommendation
type SpImmunizationRecommendation struct {
	Date          string
	Identifier    string
	TargetDisease string
	Patient       string
	VaccineType   string
	Information   string
	Support       string
	Status        string
}

func (p SpImmunizationRecommendation) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"targetdisease", p.TargetDisease},
		{"patient", p.Patient},
		{"vaccinetype", p.VaccineType},
		{"information", p.Information},
		{"support", p.Support},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ImmunizationRecommendation")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ImplementationGuide
type SpImplementationGuide struct {
	Date                string
	ContextTypeValue    string
	Resource            string
	Jurisdiction        string
	Description         string
	ContextType         string
	Experimental        string
	Global              string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpImplementationGuide) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"contexttypevalue", p.ContextTypeValue},
		{"resource", p.Resource},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"experimental", p.Experimental},
		{"global", p.Global},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ImplementationGuide")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for InsurancePlan
type SpInsurancePlan struct {
	Identifier        string
	Address           string
	AddressState      string
	OwnedBy           string
	Type              string
	AddressPostalcode string
	AdministeredBy    string
	AddressCountry    string
	Endpoint          string
	Phonetic          string
	Name              string
	AddressUse        string
	AddressCity       string
	Status            string
}

func (p SpInsurancePlan) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"address", p.Address},
		{"addressstate", p.AddressState},
		{"ownedby", p.OwnedBy},
		{"type", p.Type},
		{"addresspostalcode", p.AddressPostalcode},
		{"administeredby", p.AdministeredBy},
		{"addresscountry", p.AddressCountry},
		{"endpoint", p.Endpoint},
		{"phonetic", p.Phonetic},
		{"name", p.Name},
		{"addressuse", p.AddressUse},
		{"addresscity", p.AddressCity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "InsurancePlan")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Invoice
type SpInvoice struct {
	Date            string
	Identifier      string
	Totalgross      string
	Subject         string
	ParticipantRole string
	Type            string
	Issuer          string
	Participant     string
	Totalnet        string
	Patient         string
	Recipient       string
	Account         string
	Status          string
}

func (p SpInvoice) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"totalgross", p.Totalgross},
		{"subject", p.Subject},
		{"participantrole", p.ParticipantRole},
		{"type", p.Type},
		{"issuer", p.Issuer},
		{"participant", p.Participant},
		{"totalnet", p.Totalnet},
		{"patient", p.Patient},
		{"recipient", p.Recipient},
		{"account", p.Account},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Invoice")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Library
type SpLibrary struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ComposedOf          string
	Type                string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	Topic               string
	ContentType         string
	ContextTypeQuantity string
	Status              string
}

func (p SpLibrary) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"successor", p.Successor},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"composedof", p.ComposedOf},
		{"type", p.Type},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
		{"contenttype", p.ContentType},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Library")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Linkage
type SpLinkage struct {
	Item   string
	Author string
	Source string
}

func (p SpLinkage) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"item", p.Item},
		{"author", p.Author},
		{"source", p.Source},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Linkage")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for List
type SpList struct {
	Date        string
	Identifier  string
	Item        string
	EmptyReason string
	Code        string
	Notes       string
	Subject     string
	Patient     string
	Source      string
	Encounter   string
	Title       string
	Status      string
}

func (p SpList) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"item", p.Item},
		{"emptyreason", p.EmptyReason},
		{"code", p.Code},
		{"notes", p.Notes},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"source", p.Source},
		{"encounter", p.Encounter},
		{"title", p.Title},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "List")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Location
type SpLocation struct {
	Identifier        string
	Partof            string
	Address           string
	AddressState      string
	OperationalStatus string
	Type              string
	AddressPostalcode string
	AddressCountry    string
	Endpoint          string
	Organization      string
	Name              string
	AddressUse        string
	Near              string
	AddressCity       string
	Status            string
}

func (p SpLocation) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"partof", p.Partof},
		{"address", p.Address},
		{"addressstate", p.AddressState},
		{"operationalstatus", p.OperationalStatus},
		{"type", p.Type},
		{"addresspostalcode", p.AddressPostalcode},
		{"addresscountry", p.AddressCountry},
		{"endpoint", p.Endpoint},
		{"organization", p.Organization},
		{"name", p.Name},
		{"addressuse", p.AddressUse},
		{"near", p.Near},
		{"addresscity", p.AddressCity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Location")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Measure
type SpMeasure struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ComposedOf          string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	Topic               string
	ContextTypeQuantity string
	Status              string
}

func (p SpMeasure) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"successor", p.Successor},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"composedof", p.ComposedOf},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Measure")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MeasureReport
type SpMeasureReport struct {
	Date              string
	Identifier        string
	Period            string
	Measure           string
	Patient           string
	Subject           string
	Reporter          string
	Status            string
	EvaluatedResource string
}

func (p SpMeasureReport) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"period", p.Period},
		{"measure", p.Measure},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"reporter", p.Reporter},
		{"status", p.Status},
		{"evaluatedresource", p.EvaluatedResource},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MeasureReport")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Media
type SpMedia struct {
	Identifier string
	Modality   string
	Subject    string
	Created    string
	Encounter  string
	Type       string
	Operator   string
	View       string
	Site       string
	BasedOn    string
	Patient    string
	Device     string
	Status     string
}

func (p SpMedia) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"modality", p.Modality},
		{"subject", p.Subject},
		{"created", p.Created},
		{"encounter", p.Encounter},
		{"type", p.Type},
		{"operator", p.Operator},
		{"view", p.View},
		{"site", p.Site},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"device", p.Device},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Media")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Medication
type SpMedication struct {
	IngredientCode string
	Identifier     string
	Code           string
	Ingredient     string
	Form           string
	LotNumber      string
	ExpirationDate string
	Manufacturer   string
	Status         string
}

func (p SpMedication) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"ingredientcode", p.IngredientCode},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"ingredient", p.Ingredient},
		{"form", p.Form},
		{"lotnumber", p.LotNumber},
		{"expirationdate", p.ExpirationDate},
		{"manufacturer", p.Manufacturer},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Medication")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicationAdministration
type SpMedicationAdministration struct {
	Identifier     string
	Request        string
	Code           string
	Performer      string
	Subject        string
	Medication     string
	ReasonGiven    string
	Patient        string
	EffectiveTime  string
	Context        string
	ReasonNotGiven string
	Device         string
	Status         string
}

func (p SpMedicationAdministration) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"request", p.Request},
		{"code", p.Code},
		{"performer", p.Performer},
		{"subject", p.Subject},
		{"medication", p.Medication},
		{"reasongiven", p.ReasonGiven},
		{"patient", p.Patient},
		{"effectivetime", p.EffectiveTime},
		{"context", p.Context},
		{"reasonnotgiven", p.ReasonNotGiven},
		{"device", p.Device},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicationAdministration")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicationDispense
type SpMedicationDispense struct {
	Identifier       string
	Performer        string
	Code             string
	Receiver         string
	Subject          string
	Destination      string
	Medication       string
	Responsibleparty string
	Type             string
	Whenhandedover   string
	Whenprepared     string
	Prescription     string
	Patient          string
	Context          string
	Status           string
}

func (p SpMedicationDispense) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"performer", p.Performer},
		{"code", p.Code},
		{"receiver", p.Receiver},
		{"subject", p.Subject},
		{"destination", p.Destination},
		{"medication", p.Medication},
		{"responsibleparty", p.Responsibleparty},
		{"type", p.Type},
		{"whenhandedover", p.Whenhandedover},
		{"whenprepared", p.Whenprepared},
		{"prescription", p.Prescription},
		{"patient", p.Patient},
		{"context", p.Context},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicationDispense")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicationKnowledge
type SpMedicationKnowledge struct {
	Code                  string
	Ingredient            string
	Doseform              string
	ClassificationType    string
	MonographType         string
	Classification        string
	Manufacturer          string
	IngredientCode        string
	SourceCost            string
	Monograph             string
	MonitoringProgramName string
	MonitoringProgramType string
	Status                string
}

func (p SpMedicationKnowledge) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"code", p.Code},
		{"ingredient", p.Ingredient},
		{"doseform", p.Doseform},
		{"classificationtype", p.ClassificationType},
		{"monographtype", p.MonographType},
		{"classification", p.Classification},
		{"manufacturer", p.Manufacturer},
		{"ingredientcode", p.IngredientCode},
		{"sourcecost", p.SourceCost},
		{"monograph", p.Monograph},
		{"monitoringprogramname", p.MonitoringProgramName},
		{"monitoringprogramtype", p.MonitoringProgramType},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicationKnowledge")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicationRequest
type SpMedicationRequest struct {
	Requester             string
	Date                  string
	Identifier            string
	IntendedDispenser     string
	Authoredon            string
	Code                  string
	Subject               string
	Medication            string
	Encounter             string
	Priority              string
	Intent                string
	Patient               string
	IntendedPerformer     string
	IntendedPerformertype string
	Category              string
	Status                string
}

func (p SpMedicationRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"requester", p.Requester},
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"intendeddispenser", p.IntendedDispenser},
		{"authoredon", p.Authoredon},
		{"code", p.Code},
		{"subject", p.Subject},
		{"medication", p.Medication},
		{"encounter", p.Encounter},
		{"priority", p.Priority},
		{"intent", p.Intent},
		{"patient", p.Patient},
		{"intendedperformer", p.IntendedPerformer},
		{"intendedperformertype", p.IntendedPerformertype},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicationRequest")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicationStatement
type SpMedicationStatement struct {
	Identifier string
	Effective  string
	Code       string
	Subject    string
	Patient    string
	Context    string
	Medication string
	PartOf     string
	Source     string
	Category   string
	Status     string
}

func (p SpMedicationStatement) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"effective", p.Effective},
		{"code", p.Code},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"context", p.Context},
		{"medication", p.Medication},
		{"partof", p.PartOf},
		{"source", p.Source},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicationStatement")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProduct
type SpMedicinalProduct struct {
	Identifier   string
	Name         string
	NameLanguage string
}

func (p SpMedicinalProduct) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"name", p.Name},
		{"namelanguage", p.NameLanguage},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProduct")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProductAuthorization
type SpMedicinalProductAuthorization struct {
	Identifier string
	Country    string
	Subject    string
	Holder     string
	Status     string
}

func (p SpMedicinalProductAuthorization) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"country", p.Country},
		{"subject", p.Subject},
		{"holder", p.Holder},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductAuthorization")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProductContraindication
type SpMedicinalProductContraindication struct {
	Subject string
}

func (p SpMedicinalProductContraindication) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"subject", p.Subject},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductContraindication")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProductIndication
type SpMedicinalProductIndication struct {
	Subject string
}

func (p SpMedicinalProductIndication) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"subject", p.Subject},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductIndication")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProductIngredient
type SpMedicinalProductIngredient struct {
}

func (p SpMedicinalProductIngredient) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductIngredient")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProductInteraction
type SpMedicinalProductInteraction struct {
	Subject string
}

func (p SpMedicinalProductInteraction) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"subject", p.Subject},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductInteraction")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProductManufactured
type SpMedicinalProductManufactured struct {
}

func (p SpMedicinalProductManufactured) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductManufactured")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProductPackaged
type SpMedicinalProductPackaged struct {
	Identifier string
	Subject    string
}

func (p SpMedicinalProductPackaged) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"subject", p.Subject},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductPackaged")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProductPharmaceutical
type SpMedicinalProductPharmaceutical struct {
	Identifier    string
	Route         string
	TargetSpecies string
}

func (p SpMedicinalProductPharmaceutical) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"route", p.Route},
		{"targetspecies", p.TargetSpecies},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductPharmaceutical")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MedicinalProductUndesirableEffect
type SpMedicinalProductUndesirableEffect struct {
	Subject string
}

func (p SpMedicinalProductUndesirableEffect) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"subject", p.Subject},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductUndesirableEffect")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MessageDefinition
type SpMessageDefinition struct {
	Date                string
	Identifier          string
	Parent              string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	Focus               string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Name                string
	Context             string
	Publisher           string
	Event               string
	Category            string
	ContextTypeQuantity string
	Status              string
}

func (p SpMessageDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"parent", p.Parent},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"focus", p.Focus},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"event", p.Event},
		{"category", p.Category},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MessageDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MessageHeader
type SpMessageHeader struct {
	Code           string
	Receiver       string
	Author         string
	Destination    string
	Focus          string
	Source         string
	Target         string
	DestinationUri string
	SourceUri      string
	Sender         string
	Responsible    string
	Enterer        string
	ResponseId     string
	Event          string
}

func (p SpMessageHeader) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"code", p.Code},
		{"receiver", p.Receiver},
		{"author", p.Author},
		{"destination", p.Destination},
		{"focus", p.Focus},
		{"source", p.Source},
		{"target", p.Target},
		{"destinationuri", p.DestinationUri},
		{"sourceuri", p.SourceUri},
		{"sender", p.Sender},
		{"responsible", p.Responsible},
		{"enterer", p.Enterer},
		{"responseid", p.ResponseId},
		{"event", p.Event},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MessageHeader")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for MolecularSequence
type SpMolecularSequence struct {
	Identifier                      string
	ReferenceseqidVariantCoordinate string
	Chromosome                      string
	WindowEnd                       string
	Type                            string
	WindowStart                     string
	VariantEnd                      string
	ChromosomeVariantCoordinate     string
	Patient                         string
	VariantStart                    string
	ChromosomeWindowCoordinate      string
	ReferenceseqidWindowCoordinate  string
	Referenceseqid                  string
}

func (p SpMolecularSequence) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"referenceseqidvariantcoordinate", p.ReferenceseqidVariantCoordinate},
		{"chromosome", p.Chromosome},
		{"windowend", p.WindowEnd},
		{"type", p.Type},
		{"windowstart", p.WindowStart},
		{"variantend", p.VariantEnd},
		{"chromosomevariantcoordinate", p.ChromosomeVariantCoordinate},
		{"patient", p.Patient},
		{"variantstart", p.VariantStart},
		{"chromosomewindowcoordinate", p.ChromosomeWindowCoordinate},
		{"referenceseqidwindowcoordinate", p.ReferenceseqidWindowCoordinate},
		{"referenceseqid", p.Referenceseqid},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MolecularSequence")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for NamingSystem
type SpNamingSystem struct {
	Date                string
	Period              string
	ContextTypeValue    string
	Kind                string
	Jurisdiction        string
	Description         string
	ContextType         string
	Type                string
	IdType              string
	ContextQuantity     string
	Responsible         string
	Contact             string
	Name                string
	Context             string
	Publisher           string
	Telecom             string
	Value               string
	ContextTypeQuantity string
	Status              string
}

func (p SpNamingSystem) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"period", p.Period},
		{"contexttypevalue", p.ContextTypeValue},
		{"kind", p.Kind},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"type", p.Type},
		{"idtype", p.IdType},
		{"contextquantity", p.ContextQuantity},
		{"responsible", p.Responsible},
		{"contact", p.Contact},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"telecom", p.Telecom},
		{"value", p.Value},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "NamingSystem")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for NutritionOrder
type SpNutritionOrder struct {
	Identifier            string
	Datetime              string
	Provider              string
	Patient               string
	Supplement            string
	Formula               string
	InstantiatesCanonical string
	InstantiatesUri       string
	Encounter             string
	Oraldiet              string
	Status                string
	Additive              string
}

func (p SpNutritionOrder) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"datetime", p.Datetime},
		{"provider", p.Provider},
		{"patient", p.Patient},
		{"supplement", p.Supplement},
		{"formula", p.Formula},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"instantiatesuri", p.InstantiatesUri},
		{"encounter", p.Encounter},
		{"oraldiet", p.Oraldiet},
		{"status", p.Status},
		{"additive", p.Additive},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "NutritionOrder")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Observation
type SpObservation struct {
	Date                       string
	ComboDataAbsentReason      string
	Code                       string
	ComboCodeValueQuantity     string
	Subject                    string
	ComponentDataAbsentReason  string
	ValueConcept               string
	ValueDate                  string
	Focus                      string
	DerivedFrom                string
	PartOf                     string
	HasMember                  string
	CodeValueString            string
	ComponentCodeValueQuantity string
	BasedOn                    string
	CodeValueDate              string
	Patient                    string
	Specimen                   string
	ComponentCode              string
	CodeValueQuantity          string
	ComboCodeValueConcept      string
	ValueString                string
	Identifier                 string
	Performer                  string
	ComboCode                  string
	Method                     string
	ValueQuantity              string
	ComponentValueQuantity     string
	DataAbsentReason           string
	ComboValueQuantity         string
	Encounter                  string
	CodeValueConcept           string
	ComponentCodeValueConcept  string
	ComponentValueConcept      string
	Category                   string
	Device                     string
	ComboValueConcept          string
	Status                     string
}

func (p SpObservation) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"combodataabsentreason", p.ComboDataAbsentReason},
		{"code", p.Code},
		{"combocodevaluequantity", p.ComboCodeValueQuantity},
		{"subject", p.Subject},
		{"componentdataabsentreason", p.ComponentDataAbsentReason},
		{"valueconcept", p.ValueConcept},
		{"valuedate", p.ValueDate},
		{"focus", p.Focus},
		{"derivedfrom", p.DerivedFrom},
		{"partof", p.PartOf},
		{"hasmember", p.HasMember},
		{"codevaluestring", p.CodeValueString},
		{"componentcodevaluequantity", p.ComponentCodeValueQuantity},
		{"basedon", p.BasedOn},
		{"codevaluedate", p.CodeValueDate},
		{"patient", p.Patient},
		{"specimen", p.Specimen},
		{"componentcode", p.ComponentCode},
		{"codevaluequantity", p.CodeValueQuantity},
		{"combocodevalueconcept", p.ComboCodeValueConcept},
		{"valuestring", p.ValueString},
		{"identifier", p.Identifier},
		{"performer", p.Performer},
		{"combocode", p.ComboCode},
		{"method", p.Method},
		{"valuequantity", p.ValueQuantity},
		{"componentvaluequantity", p.ComponentValueQuantity},
		{"dataabsentreason", p.DataAbsentReason},
		{"combovaluequantity", p.ComboValueQuantity},
		{"encounter", p.Encounter},
		{"codevalueconcept", p.CodeValueConcept},
		{"componentcodevalueconcept", p.ComponentCodeValueConcept},
		{"componentvalueconcept", p.ComponentValueConcept},
		{"category", p.Category},
		{"device", p.Device},
		{"combovalueconcept", p.ComboValueConcept},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Observation")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ObservationDefinition
type SpObservationDefinition struct {
}

func (p SpObservationDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ObservationDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for OperationDefinition
type SpOperationDefinition struct {
	Date                string
	Code                string
	Instance            string
	ContextTypeValue    string
	Kind                string
	Jurisdiction        string
	Description         string
	ContextType         string
	Title               string
	Type                string
	Version             string
	Url                 string
	ContextQuantity     string
	InputProfile        string
	OutputProfile       string
	System              string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
	Base                string
}

func (p SpOperationDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"code", p.Code},
		{"instance", p.Instance},
		{"contexttypevalue", p.ContextTypeValue},
		{"kind", p.Kind},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"type", p.Type},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"inputprofile", p.InputProfile},
		{"outputprofile", p.OutputProfile},
		{"system", p.System},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
		{"base", p.Base},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "OperationDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for OperationOutcome
type SpOperationOutcome struct {
}

func (p SpOperationOutcome) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "OperationOutcome")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Organization
type SpOrganization struct {
	Identifier        string
	Partof            string
	Address           string
	AddressState      string
	Active            string
	Type              string
	AddressPostalcode string
	AddressCountry    string
	Endpoint          string
	Phonetic          string
	Name              string
	AddressUse        string
	AddressCity       string
}

func (p SpOrganization) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"partof", p.Partof},
		{"address", p.Address},
		{"addressstate", p.AddressState},
		{"active", p.Active},
		{"type", p.Type},
		{"addresspostalcode", p.AddressPostalcode},
		{"addresscountry", p.AddressCountry},
		{"endpoint", p.Endpoint},
		{"phonetic", p.Phonetic},
		{"name", p.Name},
		{"addressuse", p.AddressUse},
		{"addresscity", p.AddressCity},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Organization")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for OrganizationAffiliation
type SpOrganizationAffiliation struct {
	Date                      string
	Identifier                string
	Specialty                 string
	Role                      string
	Active                    string
	PrimaryOrganization       string
	Network                   string
	Endpoint                  string
	Phone                     string
	Service                   string
	ParticipatingOrganization string
	Telecom                   string
	Location                  string
	Email                     string
}

func (p SpOrganizationAffiliation) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"specialty", p.Specialty},
		{"role", p.Role},
		{"active", p.Active},
		{"primaryorganization", p.PrimaryOrganization},
		{"network", p.Network},
		{"endpoint", p.Endpoint},
		{"phone", p.Phone},
		{"service", p.Service},
		{"participatingorganization", p.ParticipatingOrganization},
		{"telecom", p.Telecom},
		{"location", p.Location},
		{"email", p.Email},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "OrganizationAffiliation")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Patient
type SpPatient struct {
	Identifier          string
	Given               string
	Address             string
	Birthdate           string
	Deceased            string
	AddressState        string
	Gender              string
	GeneralPractitioner string
	Link                string
	Active              string
	Language            string
	AddressPostalcode   string
	AddressCountry      string
	DeathDate           string
	Phonetic            string
	Phone               string
	Organization        string
	Name                string
	AddressUse          string
	Telecom             string
	Family              string
	AddressCity         string
	Email               string
}

func (p SpPatient) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"given", p.Given},
		{"address", p.Address},
		{"birthdate", p.Birthdate},
		{"deceased", p.Deceased},
		{"addressstate", p.AddressState},
		{"gender", p.Gender},
		{"generalpractitioner", p.GeneralPractitioner},
		{"link", p.Link},
		{"active", p.Active},
		{"language", p.Language},
		{"addresspostalcode", p.AddressPostalcode},
		{"addresscountry", p.AddressCountry},
		{"deathdate", p.DeathDate},
		{"phonetic", p.Phonetic},
		{"phone", p.Phone},
		{"organization", p.Organization},
		{"name", p.Name},
		{"addressuse", p.AddressUse},
		{"telecom", p.Telecom},
		{"family", p.Family},
		{"addresscity", p.AddressCity},
		{"email", p.Email},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Patient")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for PaymentNotice
type SpPaymentNotice struct {
	Identifier    string
	Request       string
	Provider      string
	Created       string
	Response      string
	PaymentStatus string
	Status        string
}

func (p SpPaymentNotice) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"request", p.Request},
		{"provider", p.Provider},
		{"created", p.Created},
		{"response", p.Response},
		{"paymentstatus", p.PaymentStatus},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "PaymentNotice")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for PaymentReconciliation
type SpPaymentReconciliation struct {
	Identifier    string
	Request       string
	Disposition   string
	Created       string
	PaymentIssuer string
	Outcome       string
	Requestor     string
	Status        string
}

func (p SpPaymentReconciliation) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"request", p.Request},
		{"disposition", p.Disposition},
		{"created", p.Created},
		{"paymentissuer", p.PaymentIssuer},
		{"outcome", p.Outcome},
		{"requestor", p.Requestor},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "PaymentReconciliation")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Person
type SpPerson struct {
	Identifier        string
	Address           string
	Birthdate         string
	AddressState      string
	Gender            string
	Practitioner      string
	Link              string
	Relatedperson     string
	AddressPostalcode string
	AddressCountry    string
	Phonetic          string
	Phone             string
	Patient           string
	Organization      string
	Name              string
	AddressUse        string
	Telecom           string
	AddressCity       string
	Email             string
}

func (p SpPerson) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"address", p.Address},
		{"birthdate", p.Birthdate},
		{"addressstate", p.AddressState},
		{"gender", p.Gender},
		{"practitioner", p.Practitioner},
		{"link", p.Link},
		{"relatedperson", p.Relatedperson},
		{"addresspostalcode", p.AddressPostalcode},
		{"addresscountry", p.AddressCountry},
		{"phonetic", p.Phonetic},
		{"phone", p.Phone},
		{"patient", p.Patient},
		{"organization", p.Organization},
		{"name", p.Name},
		{"addressuse", p.AddressUse},
		{"telecom", p.Telecom},
		{"addresscity", p.AddressCity},
		{"email", p.Email},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Person")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for PlanDefinition
type SpPlanDefinition struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ComposedOf          string
	Type                string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	Topic               string
	Definition          string
	ContextTypeQuantity string
	Status              string
}

func (p SpPlanDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"successor", p.Successor},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"composedof", p.ComposedOf},
		{"type", p.Type},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
		{"definition", p.Definition},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "PlanDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Practitioner
type SpPractitioner struct {
	Identifier        string
	Given             string
	Address           string
	AddressState      string
	Gender            string
	Active            string
	AddressPostalcode string
	AddressCountry    string
	Phonetic          string
	Phone             string
	Name              string
	AddressUse        string
	Telecom           string
	Family            string
	AddressCity       string
	Communication     string
	Email             string
}

func (p SpPractitioner) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"given", p.Given},
		{"address", p.Address},
		{"addressstate", p.AddressState},
		{"gender", p.Gender},
		{"active", p.Active},
		{"addresspostalcode", p.AddressPostalcode},
		{"addresscountry", p.AddressCountry},
		{"phonetic", p.Phonetic},
		{"phone", p.Phone},
		{"name", p.Name},
		{"addressuse", p.AddressUse},
		{"telecom", p.Telecom},
		{"family", p.Family},
		{"addresscity", p.AddressCity},
		{"communication", p.Communication},
		{"email", p.Email},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Practitioner")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for PractitionerRole
type SpPractitionerRole struct {
	Date         string
	Identifier   string
	Specialty    string
	Role         string
	Practitioner string
	Active       string
	Endpoint     string
	Phone        string
	Service      string
	Organization string
	Telecom      string
	Location     string
	Email        string
}

func (p SpPractitionerRole) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"specialty", p.Specialty},
		{"role", p.Role},
		{"practitioner", p.Practitioner},
		{"active", p.Active},
		{"endpoint", p.Endpoint},
		{"phone", p.Phone},
		{"service", p.Service},
		{"organization", p.Organization},
		{"telecom", p.Telecom},
		{"location", p.Location},
		{"email", p.Email},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "PractitionerRole")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Procedure
type SpProcedure struct {
	Date                  string
	Identifier            string
	Code                  string
	Performer             string
	Subject               string
	InstantiatesCanonical string
	PartOf                string
	Encounter             string
	ReasonCode            string
	BasedOn               string
	Patient               string
	ReasonReference       string
	Location              string
	InstantiatesUri       string
	Category              string
	Status                string
}

func (p SpProcedure) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"performer", p.Performer},
		{"subject", p.Subject},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"partof", p.PartOf},
		{"encounter", p.Encounter},
		{"reasoncode", p.ReasonCode},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"reasonreference", p.ReasonReference},
		{"location", p.Location},
		{"instantiatesuri", p.InstantiatesUri},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Procedure")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Provenance
type SpProvenance struct {
	AgentType     string
	Agent         string
	SignatureType string
	Patient       string
	Location      string
	Recorded      string
	AgentRole     string
	When          string
	Entity        string
	Target        string
}

func (p SpProvenance) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"agenttype", p.AgentType},
		{"agent", p.Agent},
		{"signaturetype", p.SignatureType},
		{"patient", p.Patient},
		{"location", p.Location},
		{"recorded", p.Recorded},
		{"agentrole", p.AgentRole},
		{"when", p.When},
		{"entity", p.Entity},
		{"target", p.Target},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Provenance")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Questionnaire
type SpQuestionnaire struct {
	Date                string
	Identifier          string
	Code                string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	SubjectType         string
	Name                string
	Context             string
	Publisher           string
	Definition          string
	ContextTypeQuantity string
	Status              string
}

func (p SpQuestionnaire) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"subjecttype", p.SubjectType},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"definition", p.Definition},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Questionnaire")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for QuestionnaireResponse
type SpQuestionnaireResponse struct {
	Authored      string
	Identifier    string
	Questionnaire string
	BasedOn       string
	Subject       string
	Author        string
	Patient       string
	PartOf        string
	Encounter     string
	Source        string
	Status        string
}

func (p SpQuestionnaireResponse) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"authored", p.Authored},
		{"identifier", p.Identifier},
		{"questionnaire", p.Questionnaire},
		{"basedon", p.BasedOn},
		{"subject", p.Subject},
		{"author", p.Author},
		{"patient", p.Patient},
		{"partof", p.PartOf},
		{"encounter", p.Encounter},
		{"source", p.Source},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "QuestionnaireResponse")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for RelatedPerson
type SpRelatedPerson struct {
	Identifier        string
	Address           string
	Birthdate         string
	AddressState      string
	Gender            string
	Active            string
	AddressPostalcode string
	AddressCountry    string
	Phonetic          string
	Phone             string
	Patient           string
	Name              string
	AddressUse        string
	Telecom           string
	AddressCity       string
	Relationship      string
	Email             string
}

func (p SpRelatedPerson) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"address", p.Address},
		{"birthdate", p.Birthdate},
		{"addressstate", p.AddressState},
		{"gender", p.Gender},
		{"active", p.Active},
		{"addresspostalcode", p.AddressPostalcode},
		{"addresscountry", p.AddressCountry},
		{"phonetic", p.Phonetic},
		{"phone", p.Phone},
		{"patient", p.Patient},
		{"name", p.Name},
		{"addressuse", p.AddressUse},
		{"telecom", p.Telecom},
		{"addresscity", p.AddressCity},
		{"relationship", p.Relationship},
		{"email", p.Email},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "RelatedPerson")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for RequestGroup
type SpRequestGroup struct {
	Authored              string
	Identifier            string
	Code                  string
	Subject               string
	Author                string
	InstantiatesCanonical string
	Encounter             string
	Priority              string
	Intent                string
	Participant           string
	GroupIdentifier       string
	Patient               string
	InstantiatesUri       string
	Status                string
}

func (p SpRequestGroup) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"authored", p.Authored},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"subject", p.Subject},
		{"author", p.Author},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"encounter", p.Encounter},
		{"priority", p.Priority},
		{"intent", p.Intent},
		{"participant", p.Participant},
		{"groupidentifier", p.GroupIdentifier},
		{"patient", p.Patient},
		{"instantiatesuri", p.InstantiatesUri},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "RequestGroup")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ResearchDefinition
type SpResearchDefinition struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ComposedOf          string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	Topic               string
	ContextTypeQuantity string
	Status              string
}

func (p SpResearchDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"successor", p.Successor},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"composedof", p.ComposedOf},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ResearchDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ResearchElementDefinition
type SpResearchElementDefinition struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ComposedOf          string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	DependsOn           string
	Name                string
	Context             string
	Publisher           string
	Topic               string
	ContextTypeQuantity string
	Status              string
}

func (p SpResearchElementDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"successor", p.Successor},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"composedof", p.ComposedOf},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"dependson", p.DependsOn},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ResearchElementDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ResearchStudy
type SpResearchStudy struct {
	Date                  string
	Identifier            string
	Partof                string
	Sponsor               string
	Focus                 string
	Principalinvestigator string
	Title                 string
	Protocol              string
	Site                  string
	Location              string
	Category              string
	Keyword               string
	Status                string
}

func (p SpResearchStudy) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"partof", p.Partof},
		{"sponsor", p.Sponsor},
		{"focus", p.Focus},
		{"principalinvestigator", p.Principalinvestigator},
		{"title", p.Title},
		{"protocol", p.Protocol},
		{"site", p.Site},
		{"location", p.Location},
		{"category", p.Category},
		{"keyword", p.Keyword},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ResearchStudy")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ResearchSubject
type SpResearchSubject struct {
	Date       string
	Identifier string
	Study      string
	Individual string
	Patient    string
	Status     string
}

func (p SpResearchSubject) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"study", p.Study},
		{"individual", p.Individual},
		{"patient", p.Patient},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ResearchSubject")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for RiskAssessment
type SpRiskAssessment struct {
	Date        string
	Identifier  string
	Condition   string
	Performer   string
	Method      string
	Subject     string
	Patient     string
	Probability string
	Risk        string
	Encounter   string
}

func (p SpRiskAssessment) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"condition", p.Condition},
		{"performer", p.Performer},
		{"method", p.Method},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"probability", p.Probability},
		{"risk", p.Risk},
		{"encounter", p.Encounter},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "RiskAssessment")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for RiskEvidenceSynthesis
type SpRiskEvidenceSynthesis struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpRiskEvidenceSynthesis) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "RiskEvidenceSynthesis")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Schedule
type SpSchedule struct {
	Actor           string
	Date            string
	Identifier      string
	Specialty       string
	ServiceCategory string
	ServiceType     string
	Active          string
}

func (p SpSchedule) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"actor", p.Actor},
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"specialty", p.Specialty},
		{"servicecategory", p.ServiceCategory},
		{"servicetype", p.ServiceType},
		{"active", p.Active},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Schedule")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SearchParameter
type SpSearchParameter struct {
	Date                string
	Code                string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Type                string
	Version             string
	Url                 string
	Target              string
	ContextQuantity     string
	Component           string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
	Base                string
}

func (p SpSearchParameter) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"code", p.Code},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"type", p.Type},
		{"version", p.Version},
		{"url", p.Url},
		{"target", p.Target},
		{"contextquantity", p.ContextQuantity},
		{"component", p.Component},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
		{"base", p.Base},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SearchParameter")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ServiceRequest
type SpServiceRequest struct {
	Authored              string
	Requester             string
	Identifier            string
	Code                  string
	Performer             string
	Requisition           string
	Replaces              string
	Subject               string
	InstantiatesCanonical string
	Encounter             string
	Occurrence            string
	Priority              string
	Intent                string
	PerformerType         string
	BasedOn               string
	Patient               string
	Specimen              string
	InstantiatesUri       string
	BodySite              string
	Category              string
	Status                string
}

func (p SpServiceRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"authored", p.Authored},
		{"requester", p.Requester},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"performer", p.Performer},
		{"requisition", p.Requisition},
		{"replaces", p.Replaces},
		{"subject", p.Subject},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"encounter", p.Encounter},
		{"occurrence", p.Occurrence},
		{"priority", p.Priority},
		{"intent", p.Intent},
		{"performertype", p.PerformerType},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"specimen", p.Specimen},
		{"instantiatesuri", p.InstantiatesUri},
		{"bodysite", p.BodySite},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ServiceRequest")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Slot
type SpSlot struct {
	Schedule        string
	Identifier      string
	Specialty       string
	ServiceCategory string
	AppointmentType string
	ServiceType     string
	Start           string
	Status          string
}

func (p SpSlot) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"schedule", p.Schedule},
		{"identifier", p.Identifier},
		{"specialty", p.Specialty},
		{"servicecategory", p.ServiceCategory},
		{"appointmenttype", p.AppointmentType},
		{"servicetype", p.ServiceType},
		{"start", p.Start},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Slot")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Specimen
type SpSpecimen struct {
	Container   string
	Identifier  string
	Parent      string
	ContainerId string
	Bodysite    string
	Subject     string
	Patient     string
	Collected   string
	Accession   string
	Type        string
	Collector   string
	Status      string
}

func (p SpSpecimen) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"container", p.Container},
		{"identifier", p.Identifier},
		{"parent", p.Parent},
		{"containerid", p.ContainerId},
		{"bodysite", p.Bodysite},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"collected", p.Collected},
		{"accession", p.Accession},
		{"type", p.Type},
		{"collector", p.Collector},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Specimen")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SpecimenDefinition
type SpSpecimenDefinition struct {
	Container  string
	Identifier string
	Type       string
}

func (p SpSpecimenDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"container", p.Container},
		{"identifier", p.Identifier},
		{"type", p.Type},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SpecimenDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for StructureDefinition
type SpStructureDefinition struct {
	Date                string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ContextType         string
	Experimental        string
	Title               string
	Type                string
	ContextQuantity     string
	Path                string
	Context             string
	BasePath            string
	Keyword             string
	ContextTypeQuantity string
	Identifier          string
	Valueset            string
	Kind                string
	Abstract            string
	Version             string
	Url                 string
	ExtContext          string
	Name                string
	Publisher           string
	Derivation          string
	Status              string
	Base                string
}

func (p SpStructureDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"experimental", p.Experimental},
		{"title", p.Title},
		{"type", p.Type},
		{"contextquantity", p.ContextQuantity},
		{"path", p.Path},
		{"context", p.Context},
		{"basepath", p.BasePath},
		{"keyword", p.Keyword},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"identifier", p.Identifier},
		{"valueset", p.Valueset},
		{"kind", p.Kind},
		{"abstract", p.Abstract},
		{"version", p.Version},
		{"url", p.Url},
		{"extcontext", p.ExtContext},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"derivation", p.Derivation},
		{"status", p.Status},
		{"base", p.Base},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "StructureDefinition")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for StructureMap
type SpStructureMap struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpStructureMap) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "StructureMap")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Subscription
type SpSubscription struct {
	Payload  string
	Criteria string
	Contact  string
	Type     string
	Url      string
	Status   string
}

func (p SpSubscription) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"payload", p.Payload},
		{"criteria", p.Criteria},
		{"contact", p.Contact},
		{"type", p.Type},
		{"url", p.Url},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Subscription")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Substance
type SpSubstance struct {
	Identifier          string
	ContainerIdentifier string
	Code                string
	Quantity            string
	SubstanceReference  string
	Expiry              string
	Category            string
	Status              string
}

func (p SpSubstance) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"containeridentifier", p.ContainerIdentifier},
		{"code", p.Code},
		{"quantity", p.Quantity},
		{"substancereference", p.SubstanceReference},
		{"expiry", p.Expiry},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Substance")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SubstanceNucleicAcid
type SpSubstanceNucleicAcid struct {
}

func (p SpSubstanceNucleicAcid) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SubstanceNucleicAcid")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SubstancePolymer
type SpSubstancePolymer struct {
}

func (p SpSubstancePolymer) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SubstancePolymer")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SubstanceProtein
type SpSubstanceProtein struct {
}

func (p SpSubstanceProtein) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SubstanceProtein")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SubstanceReferenceInformation
type SpSubstanceReferenceInformation struct {
}

func (p SpSubstanceReferenceInformation) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SubstanceReferenceInformation")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SubstanceSourceMaterial
type SpSubstanceSourceMaterial struct {
}

func (p SpSubstanceSourceMaterial) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SubstanceSourceMaterial")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SubstanceSpecification
type SpSubstanceSpecification struct {
	Code string
}

func (p SpSubstanceSpecification) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"code", p.Code},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SubstanceSpecification")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SupplyDelivery
type SpSupplyDelivery struct {
	Identifier string
	Receiver   string
	Patient    string
	Supplier   string
	Status     string
}

func (p SpSupplyDelivery) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"receiver", p.Receiver},
		{"patient", p.Patient},
		{"supplier", p.Supplier},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SupplyDelivery")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for SupplyRequest
type SpSupplyRequest struct {
	Requester  string
	Date       string
	Identifier string
	Subject    string
	Supplier   string
	Category   string
	Status     string
}

func (p SpSupplyRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"requester", p.Requester},
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"subject", p.Subject},
		{"supplier", p.Supplier},
		{"category", p.Category},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SupplyRequest")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for Task
type SpTask struct {
	Owner           string
	Requester       string
	Identifier      string
	BusinessStatus  string
	Period          string
	Code            string
	Performer       string
	Subject         string
	Focus           string
	PartOf          string
	Encounter       string
	Priority        string
	AuthoredOn      string
	Intent          string
	GroupIdentifier string
	BasedOn         string
	Patient         string
	Modified        string
	Status          string
}

func (p SpTask) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"owner", p.Owner},
		{"requester", p.Requester},
		{"identifier", p.Identifier},
		{"businessstatus", p.BusinessStatus},
		{"period", p.Period},
		{"code", p.Code},
		{"performer", p.Performer},
		{"subject", p.Subject},
		{"focus", p.Focus},
		{"partof", p.PartOf},
		{"encounter", p.Encounter},
		{"priority", p.Priority},
		{"authoredon", p.AuthoredOn},
		{"intent", p.Intent},
		{"groupidentifier", p.GroupIdentifier},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"modified", p.Modified},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Task")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for TerminologyCapabilities
type SpTerminologyCapabilities struct {
	Date                string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpTerminologyCapabilities) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "TerminologyCapabilities")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for TestReport
type SpTestReport struct {
	Result      string
	Identifier  string
	Tester      string
	Testscript  string
	Issued      string
	Participant string
}

func (p SpTestReport) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"result", p.Result},
		{"identifier", p.Identifier},
		{"tester", p.Tester},
		{"testscript", p.Testscript},
		{"issued", p.Issued},
		{"participant", p.Participant},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "TestReport")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for TestScript
type SpTestScript struct {
	Date                 string
	Identifier           string
	ContextTypeValue     string
	Jurisdiction         string
	Description          string
	TestscriptCapability string
	ContextType          string
	Title                string
	Version              string
	Url                  string
	ContextQuantity      string
	Name                 string
	Context              string
	Publisher            string
	ContextTypeQuantity  string
	Status               string
}

func (p SpTestScript) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"testscriptcapability", p.TestscriptCapability},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "TestScript")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for ValueSet
type SpValueSet struct {
	Date                string
	Identifier          string
	Code                string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	Expansion           string
	Reference           string
	ContextQuantity     string
	Name                string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpValueSet) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"expansion", p.Expansion},
		{"reference", p.Reference},
		{"contextquantity", p.ContextQuantity},
		{"name", p.Name},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ValueSet")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for VerificationResult
type SpVerificationResult struct {
	Target string
}

func (p SpVerificationResult) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"target", p.Target},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "VerificationResult")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}

// search params for VisionPrescription
type SpVisionPrescription struct {
	Prescriber  string
	Identifier  string
	Patient     string
	Datewritten string
	Encounter   string
	Status      string
}

func (p SpVisionPrescription) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"prescriber", p.Prescriber},
		{"identifier", p.Identifier},
		{"patient", p.Patient},
		{"datewritten", p.Datewritten},
		{"encounter", p.Encounter},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "VisionPrescription")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = v.Encode()
	s := u.String()
	return &s, nil
}
