package r5Client

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "net/url"

type SearchParam interface {
	SpEncode(baseURL *string) (*string, error)
} //search params for Account
type SpAccount struct {
	Owner          string
	Identifier     string
	Period         string
	Patient        string
	Subject        string
	Name           string
	Guarantor      string
	Type           string
	Relatedaccount string
	Status         string
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
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"name", p.Name},
		{"guarantor", p.Guarantor},
		{"type", p.Type},
		{"relatedaccount", p.Relatedaccount},
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

//search params for ActivityDefinition
type SpActivityDefinition struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Kind                string
	Jurisdiction        string
	DerivedFrom         string
	Description         string
	ContextType         string
	Predecessor         string
	ComposedOf          string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	DependsOn           string
	Effective           string
	Context             string
	Name                string
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
		{"kind", p.Kind},
		{"jurisdiction", p.Jurisdiction},
		{"derivedfrom", p.DerivedFrom},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"composedof", p.ComposedOf},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"dependson", p.DependsOn},
		{"effective", p.Effective},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for ActorDefinition
type SpActorDefinition struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ContextType         string
	Title               string
	Type                string
	Version             string
	Url                 string
	ContextQuantity     string
	Context             string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpActorDefinition) SpEncode(baseURL *string) (*string, error) {
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
		{"type", p.Type},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
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

	resUrl, err := url.JoinPath(*baseURL, "ActorDefinition")
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

//search params for AdministrableProductDefinition
type SpAdministrableProductDefinition struct {
	Identifier       string
	ManufacturedItem string
	Ingredient       string
	Route            string
	DoseForm         string
	Device           string
	FormOf           string
	TargetSpecies    string
	Status           string
}

func (p SpAdministrableProductDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"manufactureditem", p.ManufacturedItem},
		{"ingredient", p.Ingredient},
		{"route", p.Route},
		{"doseform", p.DoseForm},
		{"device", p.Device},
		{"formof", p.FormOf},
		{"targetspecies", p.TargetSpecies},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "AdministrableProductDefinition")
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

//search params for AdverseEvent
type SpAdverseEvent struct {
	Date            string
	Identifier      string
	Recorder        string
	Study           string
	Code            string
	Actuality       string
	Subject         string
	Substance       string
	Patient         string
	Resultingeffect string
	Seriousness     string
	Location        string
	Category        string
	Status          string
}

func (p SpAdverseEvent) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"recorder", p.Recorder},
		{"study", p.Study},
		{"code", p.Code},
		{"actuality", p.Actuality},
		{"subject", p.Subject},
		{"substance", p.Substance},
		{"patient", p.Patient},
		{"resultingeffect", p.Resultingeffect},
		{"seriousness", p.Seriousness},
		{"location", p.Location},
		{"category", p.Category},
		{"status", p.Status},
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

//search params for AllergyIntolerance
type SpAllergyIntolerance struct {
	Date                   string
	Severity               string
	Identifier             string
	Code                   string
	VerificationStatus     string
	Criticality            string
	ManifestationReference string
	ClinicalStatus         string
	Type                   string
	Participant            string
	ManifestationCode      string
	Route                  string
	Patient                string
	Category               string
	LastDate               string
}

func (p SpAllergyIntolerance) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"severity", p.Severity},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"verificationstatus", p.VerificationStatus},
		{"criticality", p.Criticality},
		{"manifestationreference", p.ManifestationReference},
		{"clinicalstatus", p.ClinicalStatus},
		{"type", p.Type},
		{"participant", p.Participant},
		{"manifestationcode", p.ManifestationCode},
		{"route", p.Route},
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

//search params for Appointment
type SpAppointment struct {
	Date                 string
	Identifier           string
	Specialty            string
	ServiceCategory      string
	Practitioner         string
	AppointmentType      string
	PartStatus           string
	Subject              string
	ServiceType          string
	Slot                 string
	ReasonCode           string
	Actor                string
	BasedOn              string
	Patient              string
	ReasonReference      string
	SupportingInfo       string
	RequestedPeriod      string
	Location             string
	Group                string
	ServiceTypeReference string
	Status               string
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
		{"appointmenttype", p.AppointmentType},
		{"partstatus", p.PartStatus},
		{"subject", p.Subject},
		{"servicetype", p.ServiceType},
		{"slot", p.Slot},
		{"reasoncode", p.ReasonCode},
		{"actor", p.Actor},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"reasonreference", p.ReasonReference},
		{"supportinginfo", p.SupportingInfo},
		{"requestedperiod", p.RequestedPeriod},
		{"location", p.Location},
		{"group", p.Group},
		{"servicetypereference", p.ServiceTypeReference},
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

//search params for AppointmentResponse
type SpAppointmentResponse struct {
	Actor        string
	Identifier   string
	Practitioner string
	PartStatus   string
	Patient      string
	Appointment  string
	Location     string
	Group        string
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
		{"group", p.Group},
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

//search params for ArtifactAssessment
type SpArtifactAssessment struct {
	Date       string
	Identifier string
}

func (p SpArtifactAssessment) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ArtifactAssessment")
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

//search params for AuditEvent
type SpAuditEvent struct {
	Date       string
	Agent      string
	EntityRole string
	Code       string
	Purpose    string
	Encounter  string
	Source     string
	BasedOn    string
	Patient    string
	Action     string
	AgentRole  string
	Category   string
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
		{"agent", p.Agent},
		{"entityrole", p.EntityRole},
		{"code", p.Code},
		{"purpose", p.Purpose},
		{"encounter", p.Encounter},
		{"source", p.Source},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"action", p.Action},
		{"agentrole", p.AgentRole},
		{"category", p.Category},
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

//search params for Basic
type SpBasic struct {
	Identifier string
	Code       string
	Author     string
	Created    string
	Patient    string
	Subject    string
}

func (p SpBasic) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"author", p.Author},
		{"created", p.Created},
		{"patient", p.Patient},
		{"subject", p.Subject},
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

//search params for Binary
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

//search params for BiologicallyDerivedProduct
type SpBiologicallyDerivedProduct struct {
	Identifier            string
	Request               string
	Code                  string
	ProductStatus         string
	SerialNumber          string
	BiologicalSourceEvent string
	ProductCategory       string
	Collector             string
}

func (p SpBiologicallyDerivedProduct) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"request", p.Request},
		{"code", p.Code},
		{"productstatus", p.ProductStatus},
		{"serialnumber", p.SerialNumber},
		{"biologicalsourceevent", p.BiologicalSourceEvent},
		{"productcategory", p.ProductCategory},
		{"collector", p.Collector},
	}
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

//search params for BiologicallyDerivedProductDispense
type SpBiologicallyDerivedProductDispense struct {
	Identifier string
	Product    string
	Performer  string
	Patient    string
	Status     string
}

func (p SpBiologicallyDerivedProductDispense) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"product", p.Product},
		{"performer", p.Performer},
		{"patient", p.Patient},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "BiologicallyDerivedProductDispense")
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

//search params for BodyStructure
type SpBodyStructure struct {
	Identifier         string
	Included_structure string
	Excluded_structure string
	Morphology         string
	Patient            string
}

func (p SpBodyStructure) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"included_structure", p.Included_structure},
		{"excluded_structure", p.Excluded_structure},
		{"morphology", p.Morphology},
		{"patient", p.Patient},
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

//search params for Bundle
type SpBundle struct {
	Identifier  string
	Composition string
	Message     string
	Type        string
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
		{"message", p.Message},
		{"type", p.Type},
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

//search params for CapabilityStatement
type SpCapabilityStatement struct {
	Date                string
	Identifier          string
	ResourceProfile     string
	ContextTypeValue    string
	Software            string
	Resource            string
	Jurisdiction        string
	Format              string
	Description         string
	ContextType         string
	Fhirversion         string
	Title               string
	Version             string
	SupportedProfile    string
	Url                 string
	Mode                string
	ContextQuantity     string
	SecurityService     string
	Context             string
	Name                string
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
		{"identifier", p.Identifier},
		{"resourceprofile", p.ResourceProfile},
		{"contexttypevalue", p.ContextTypeValue},
		{"software", p.Software},
		{"resource", p.Resource},
		{"jurisdiction", p.Jurisdiction},
		{"format", p.Format},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"fhirversion", p.Fhirversion},
		{"title", p.Title},
		{"version", p.Version},
		{"supportedprofile", p.SupportedProfile},
		{"url", p.Url},
		{"mode", p.Mode},
		{"contextquantity", p.ContextQuantity},
		{"securityservice", p.SecurityService},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for CarePlan
type SpCarePlan struct {
	CareTeam              string
	Date                  string
	Identifier            string
	Goal                  string
	Custodian             string
	Replaces              string
	Subject               string
	InstantiatesCanonical string
	PartOf                string
	Encounter             string
	Intent                string
	ActivityReference     string
	Condition             string
	BasedOn               string
	Patient               string
	InstantiatesUri       string
	Category              string
	Status                string
}

func (p SpCarePlan) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"careteam", p.CareTeam},
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"goal", p.Goal},
		{"custodian", p.Custodian},
		{"replaces", p.Replaces},
		{"subject", p.Subject},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"partof", p.PartOf},
		{"encounter", p.Encounter},
		{"intent", p.Intent},
		{"activityreference", p.ActivityReference},
		{"condition", p.Condition},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"instantiatesuri", p.InstantiatesUri},
		{"category", p.Category},
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

//search params for CareTeam
type SpCareTeam struct {
	Date        string
	Identifier  string
	Patient     string
	Subject     string
	Name        string
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
		{"name", p.Name},
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

//search params for ChargeItem
type SpChargeItem struct {
	Identifier             string
	PerformingOrganization string
	Code                   string
	Quantity               string
	Subject                string
	Encounter              string
	Occurrence             string
	EnteredDate            string
	PerformerFunction      string
	FactorOverride         string
	Patient                string
	Service                string
	PriceOverride          string
	Enterer                string
	PerformerActor         string
	Account                string
	RequestingOrganization string
	Status                 string
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
		{"encounter", p.Encounter},
		{"occurrence", p.Occurrence},
		{"entereddate", p.EnteredDate},
		{"performerfunction", p.PerformerFunction},
		{"factoroverride", p.FactorOverride},
		{"patient", p.Patient},
		{"service", p.Service},
		{"priceoverride", p.PriceOverride},
		{"enterer", p.Enterer},
		{"performeractor", p.PerformerActor},
		{"account", p.Account},
		{"requestingorganization", p.RequestingOrganization},
		{"status", p.Status},
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

//search params for ChargeItemDefinition
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

//search params for Citation
type SpCitation struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	ClassificationType  string
	ContextType         string
	Title               string
	Classification      string
	Version             string
	Url                 string
	ContextQuantity     string
	Effective           string
	Context             string
	Name                string
	Classifier          string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpCitation) SpEncode(baseURL *string) (*string, error) {
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
		{"classificationtype", p.ClassificationType},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"classification", p.Classification},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"context", p.Context},
		{"name", p.Name},
		{"classifier", p.Classifier},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Citation")
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

//search params for Claim
type SpClaim struct {
	CareTeam     string
	Identifier   string
	Created      string
	Use          string
	Encounter    string
	Priority     string
	Payee        string
	Provider     string
	Insurer      string
	Patient      string
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
		{"created", p.Created},
		{"use", p.Use},
		{"encounter", p.Encounter},
		{"priority", p.Priority},
		{"payee", p.Payee},
		{"provider", p.Provider},
		{"insurer", p.Insurer},
		{"patient", p.Patient},
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

//search params for ClaimResponse
type SpClaimResponse struct {
	Identifier  string
	Request     string
	Disposition string
	Created     string
	Insurer     string
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
		{"created", p.Created},
		{"insurer", p.Insurer},
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

//search params for ClinicalImpression
type SpClinicalImpression struct {
	Date           string
	Identifier     string
	Performer      string
	Problem        string
	Previous       string
	FindingCode    string
	Patient        string
	Subject        string
	SupportingInfo string
	Encounter      string
	FindingRef     string
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
		{"performer", p.Performer},
		{"problem", p.Problem},
		{"previous", p.Previous},
		{"findingcode", p.FindingCode},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"supportinginfo", p.SupportingInfo},
		{"encounter", p.Encounter},
		{"findingref", p.FindingRef},
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

//search params for ClinicalUseDefinition
type SpClinicalUseDefinition struct {
	ContraindicationReference string
	Identifier                string
	IndicationReference       string
	Product                   string
	Subject                   string
	Effect                    string
	Interaction               string
	Indication                string
	Type                      string
	Contraindication          string
	EffectReference           string
	Status                    string
}

func (p SpClinicalUseDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"contraindicationreference", p.ContraindicationReference},
		{"identifier", p.Identifier},
		{"indicationreference", p.IndicationReference},
		{"product", p.Product},
		{"subject", p.Subject},
		{"effect", p.Effect},
		{"interaction", p.Interaction},
		{"indication", p.Indication},
		{"type", p.Type},
		{"contraindication", p.Contraindication},
		{"effectreference", p.EffectReference},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ClinicalUseDefinition")
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

//search params for CodeSystem
type SpCodeSystem struct {
	Date                string
	Identifier          string
	Code                string
	ContextTypeValue    string
	ContentMode         string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Language            string
	Predecessor         string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Supplements         string
	Effective           string
	System              string
	Context             string
	Name                string
	Publisher           string
	Topic               string
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
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"language", p.Language},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"supplements", p.Supplements},
		{"effective", p.Effective},
		{"system", p.System},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for Communication
type SpCommunication struct {
	Identifier            string
	Subject               string
	InstantiatesCanonical string
	PartOf                string
	Received              string
	Encounter             string
	Medium                string
	Sent                  string
	BasedOn               string
	Sender                string
	Patient               string
	Recipient             string
	Topic                 string
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
		{"partof", p.PartOf},
		{"received", p.Received},
		{"encounter", p.Encounter},
		{"medium", p.Medium},
		{"sent", p.Sent},
		{"basedon", p.BasedOn},
		{"sender", p.Sender},
		{"patient", p.Patient},
		{"recipient", p.Recipient},
		{"topic", p.Topic},
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

//search params for CommunicationRequest
type SpCommunicationRequest struct {
	Authored            string
	Requester           string
	Identifier          string
	Replaces            string
	Subject             string
	Encounter           string
	Medium              string
	Occurrence          string
	Priority            string
	GroupIdentifier     string
	BasedOn             string
	Patient             string
	Recipient           string
	InformationProvider string
	Category            string
	Status              string
}

func (p SpCommunicationRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"authored", p.Authored},
		{"requester", p.Requester},
		{"identifier", p.Identifier},
		{"replaces", p.Replaces},
		{"subject", p.Subject},
		{"encounter", p.Encounter},
		{"medium", p.Medium},
		{"occurrence", p.Occurrence},
		{"priority", p.Priority},
		{"groupidentifier", p.GroupIdentifier},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"recipient", p.Recipient},
		{"informationprovider", p.InformationProvider},
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

//search params for CompartmentDefinition
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
	Context             string
	Name                string
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
		{"context", p.Context},
		{"name", p.Name},
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

//search params for Composition
type SpComposition struct {
	Date            string
	Identifier      string
	Period          string
	EventCode       string
	Author          string
	Subject         string
	Section         string
	Encounter       string
	Title           string
	Type            string
	Version         string
	Attester        string
	Url             string
	EventReference  string
	SectionText     string
	Entry           string
	Related         string
	Patient         string
	Category        string
	SectionCodeText string
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
		{"eventcode", p.EventCode},
		{"author", p.Author},
		{"subject", p.Subject},
		{"section", p.Section},
		{"encounter", p.Encounter},
		{"title", p.Title},
		{"type", p.Type},
		{"version", p.Version},
		{"attester", p.Attester},
		{"url", p.Url},
		{"eventreference", p.EventReference},
		{"sectiontext", p.SectionText},
		{"entry", p.Entry},
		{"related", p.Related},
		{"patient", p.Patient},
		{"category", p.Category},
		{"sectioncodetext", p.SectionCodeText},
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

//search params for ConceptMap
type SpConceptMap struct {
	Date                string
	TargetScope         string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	TargetGroupSystem   string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	ContextQuantity     string
	Effective           string
	SourceScope         string
	Context             string
	ContextTypeQuantity string
	TargetCode          string
	Identifier          string
	SourceScopeUri      string
	SourceGroupSystem   string
	MappingProperty     string
	OtherMap            string
	Version             string
	Url                 string
	SourceCode          string
	TargetScopeUri      string
	Name                string
	Publisher           string
	Topic               string
	Status              string
}

func (p SpConceptMap) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"targetscope", p.TargetScope},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"targetgroupsystem", p.TargetGroupSystem},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"sourcescope", p.SourceScope},
		{"context", p.Context},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"targetcode", p.TargetCode},
		{"identifier", p.Identifier},
		{"sourcescopeuri", p.SourceScopeUri},
		{"sourcegroupsystem", p.SourceGroupSystem},
		{"mappingproperty", p.MappingProperty},
		{"othermap", p.OtherMap},
		{"version", p.Version},
		{"url", p.Url},
		{"sourcecode", p.SourceCode},
		{"targetscopeuri", p.TargetScopeUri},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
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

//search params for Condition
type SpCondition struct {
	EvidenceDetail      string
	Severity            string
	Identifier          string
	OnsetInfo           string
	RecordedDate        string
	Code                string
	Evidence            string
	ParticipantFunction string
	Subject             string
	ParticipantActor    string
	VerificationStatus  string
	ClinicalStatus      string
	Encounter           string
	OnsetDate           string
	AbatementDate       string
	Stage               string
	AbatementString     string
	Patient             string
	AbatementAge        string
	OnsetAge            string
	BodySite            string
	Category            string
}

func (p SpCondition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"evidencedetail", p.EvidenceDetail},
		{"severity", p.Severity},
		{"identifier", p.Identifier},
		{"onsetinfo", p.OnsetInfo},
		{"recordeddate", p.RecordedDate},
		{"code", p.Code},
		{"evidence", p.Evidence},
		{"participantfunction", p.ParticipantFunction},
		{"subject", p.Subject},
		{"participantactor", p.ParticipantActor},
		{"verificationstatus", p.VerificationStatus},
		{"clinicalstatus", p.ClinicalStatus},
		{"encounter", p.Encounter},
		{"onsetdate", p.OnsetDate},
		{"abatementdate", p.AbatementDate},
		{"stage", p.Stage},
		{"abatementstring", p.AbatementString},
		{"patient", p.Patient},
		{"abatementage", p.AbatementAge},
		{"onsetage", p.OnsetAge},
		{"bodysite", p.BodySite},
		{"category", p.Category},
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

//search params for ConditionDefinition
type SpConditionDefinition struct {
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
	Context             string
	Name                string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpConditionDefinition) SpEncode(baseURL *string) (*string, error) {
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
		{"context", p.Context},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ConditionDefinition")
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

//search params for Consent
type SpConsent struct {
	Date            string
	Identifier      string
	Controller      string
	Period          string
	Data            string
	Manager         string
	Purpose         string
	Subject         string
	VerifiedDate    string
	Grantee         string
	SourceReference string
	Verified        string
	Actor           string
	SecurityLabel   string
	Patient         string
	Action          string
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
		{"controller", p.Controller},
		{"period", p.Period},
		{"data", p.Data},
		{"manager", p.Manager},
		{"purpose", p.Purpose},
		{"subject", p.Subject},
		{"verifieddate", p.VerifiedDate},
		{"grantee", p.Grantee},
		{"sourcereference", p.SourceReference},
		{"verified", p.Verified},
		{"actor", p.Actor},
		{"securitylabel", p.SecurityLabel},
		{"patient", p.Patient},
		{"action", p.Action},
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

//search params for Contract
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

//search params for Coverage
type SpCoverage struct {
	Identifier     string
	Subscriber     string
	Subscriberid   string
	Type           string
	Beneficiary    string
	Patient        string
	Insurer        string
	ClassValue     string
	PaymentbyParty string
	ClassType      string
	Dependent      string
	PolicyHolder   string
	Status         string
}

func (p SpCoverage) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"subscriber", p.Subscriber},
		{"subscriberid", p.Subscriberid},
		{"type", p.Type},
		{"beneficiary", p.Beneficiary},
		{"patient", p.Patient},
		{"insurer", p.Insurer},
		{"classvalue", p.ClassValue},
		{"paymentbyparty", p.PaymentbyParty},
		{"classtype", p.ClassType},
		{"dependent", p.Dependent},
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

//search params for CoverageEligibilityRequest
type SpCoverageEligibilityRequest struct {
	Identifier string
	Provider   string
	Created    string
	Patient    string
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
		{"created", p.Created},
		{"patient", p.Patient},
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

//search params for CoverageEligibilityResponse
type SpCoverageEligibilityResponse struct {
	Identifier  string
	Request     string
	Disposition string
	Created     string
	Insurer     string
	Patient     string
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
		{"created", p.Created},
		{"insurer", p.Insurer},
		{"patient", p.Patient},
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

//search params for DetectedIssue
type SpDetectedIssue struct {
	Identifier string
	Code       string
	Identified string
	Author     string
	Subject    string
	Patient    string
	Implicated string
	Category   string
	Status     string
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
		{"author", p.Author},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"implicated", p.Implicated},
		{"category", p.Category},
		{"status", p.Status},
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

//search params for Device
type SpDevice struct {
	UdiDi                 string
	Identifier            string
	Parent                string
	ManufactureDate       string
	UdiCarrier            string
	Code                  string
	DeviceName            string
	LotNumber             string
	SerialNumber          string
	Specification         string
	Type                  string
	Version               string
	Url                   string
	Manufacturer          string
	CodeValueConcept      string
	Organization          string
	BiologicalSourceEvent string
	Definition            string
	Location              string
	Model                 string
	ExpirationDate        string
	SpecificationVersion  string
	Status                string
}

func (p SpDevice) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"udidi", p.UdiDi},
		{"identifier", p.Identifier},
		{"parent", p.Parent},
		{"manufacturedate", p.ManufactureDate},
		{"udicarrier", p.UdiCarrier},
		{"code", p.Code},
		{"devicename", p.DeviceName},
		{"lotnumber", p.LotNumber},
		{"serialnumber", p.SerialNumber},
		{"specification", p.Specification},
		{"type", p.Type},
		{"version", p.Version},
		{"url", p.Url},
		{"manufacturer", p.Manufacturer},
		{"codevalueconcept", p.CodeValueConcept},
		{"organization", p.Organization},
		{"biologicalsourceevent", p.BiologicalSourceEvent},
		{"definition", p.Definition},
		{"location", p.Location},
		{"model", p.Model},
		{"expirationdate", p.ExpirationDate},
		{"specificationversion", p.SpecificationVersion},
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

//search params for DeviceAssociation
type SpDeviceAssociation struct {
	Identifier string
	Subject    string
	Patient    string
	Device     string
	Operator   string
	Status     string
}

func (p SpDeviceAssociation) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"device", p.Device},
		{"operator", p.Operator},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DeviceAssociation")
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

//search params for DeviceDefinition
type SpDeviceDefinition struct {
	Identifier           string
	DeviceName           string
	Organization         string
	Specification        string
	Type                 string
	SpecificationVersion string
	Manufacturer         string
}

func (p SpDeviceDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"devicename", p.DeviceName},
		{"organization", p.Organization},
		{"specification", p.Specification},
		{"type", p.Type},
		{"specificationversion", p.SpecificationVersion},
		{"manufacturer", p.Manufacturer},
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

//search params for DeviceDispense
type SpDeviceDispense struct {
	Identifier string
	Code       string
	Subject    string
	Patient    string
	Status     string
}

func (p SpDeviceDispense) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DeviceDispense")
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

//search params for DeviceMetric
type SpDeviceMetric struct {
	Identifier string
	Category   string
	Type       string
	Device     string
}

func (p SpDeviceMetric) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"category", p.Category},
		{"type", p.Type},
		{"device", p.Device},
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

//search params for DeviceRequest
type SpDeviceRequest struct {
	Insurance             string
	PerformerCode         string
	Requester             string
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
	Device                string
	PriorRequest          string
	Status                string
}

func (p SpDeviceRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"insurance", p.Insurance},
		{"performercode", p.PerformerCode},
		{"requester", p.Requester},
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
		{"device", p.Device},
		{"priorrequest", p.PriorRequest},
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

//search params for DeviceUsage
type SpDeviceUsage struct {
	Identifier string
	Patient    string
	Device     string
	Status     string
}

func (p SpDeviceUsage) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"patient", p.Patient},
		{"device", p.Device},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "DeviceUsage")
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

//search params for DiagnosticReport
type SpDiagnosticReport struct {
	Date               string
	Identifier         string
	Study              string
	Code               string
	Performer          string
	Subject            string
	Encounter          string
	Media              string
	Conclusion         string
	Result             string
	BasedOn            string
	Patient            string
	Specimen           string
	Category           string
	Issued             string
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
		{"study", p.Study},
		{"code", p.Code},
		{"performer", p.Performer},
		{"subject", p.Subject},
		{"encounter", p.Encounter},
		{"media", p.Media},
		{"conclusion", p.Conclusion},
		{"result", p.Result},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"specimen", p.Specimen},
		{"category", p.Category},
		{"issued", p.Issued},
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

//search params for DocumentReference
type SpDocumentReference struct {
	Date              string
	Modality          string
	Subject           string
	Description       string
	Language          string
	Type              string
	Relation          string
	Setting           string
	DocStatus         string
	BasedOn           string
	FormatCanonical   string
	Patient           string
	Context           string
	Relationship      string
	Creation          string
	Identifier        string
	Period            string
	EventCode         string
	Bodysite          string
	Custodian         string
	Author            string
	FormatCode        string
	BodysiteReference string
	FormatUri         string
	Version           string
	Attester          string
	Contenttype       string
	EventReference    string
	SecurityLabel     string
	Location          string
	Category          string
	Relatesto         string
	Facility          string
	Status            string
}

func (p SpDocumentReference) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"modality", p.Modality},
		{"subject", p.Subject},
		{"description", p.Description},
		{"language", p.Language},
		{"type", p.Type},
		{"relation", p.Relation},
		{"setting", p.Setting},
		{"docstatus", p.DocStatus},
		{"basedon", p.BasedOn},
		{"formatcanonical", p.FormatCanonical},
		{"patient", p.Patient},
		{"context", p.Context},
		{"relationship", p.Relationship},
		{"creation", p.Creation},
		{"identifier", p.Identifier},
		{"period", p.Period},
		{"eventcode", p.EventCode},
		{"bodysite", p.Bodysite},
		{"custodian", p.Custodian},
		{"author", p.Author},
		{"formatcode", p.FormatCode},
		{"bodysitereference", p.BodysiteReference},
		{"formaturi", p.FormatUri},
		{"version", p.Version},
		{"attester", p.Attester},
		{"contenttype", p.Contenttype},
		{"eventreference", p.EventReference},
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

//search params for Encounter
type SpEncounter struct {
	Date               string
	ParticipantType    string
	Subject            string
	SubjectStatus      string
	Appointment        string
	PartOf             string
	Type               string
	Participant        string
	ReasonCode         string
	BasedOn            string
	DateStart          string
	Patient            string
	LocationPeriod     string
	SpecialArrangement string
	Class              string
	Identifier         string
	DiagnosisCode      string
	Practitioner       string
	EpisodeOfCare      string
	Length             string
	Careteam           string
	EndDate            string
	DiagnosisReference string
	ReasonReference    string
	Location           string
	ServiceProvider    string
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
		{"participanttype", p.ParticipantType},
		{"subject", p.Subject},
		{"subjectstatus", p.SubjectStatus},
		{"appointment", p.Appointment},
		{"partof", p.PartOf},
		{"type", p.Type},
		{"participant", p.Participant},
		{"reasoncode", p.ReasonCode},
		{"basedon", p.BasedOn},
		{"datestart", p.DateStart},
		{"patient", p.Patient},
		{"locationperiod", p.LocationPeriod},
		{"specialarrangement", p.SpecialArrangement},
		{"class", p.Class},
		{"identifier", p.Identifier},
		{"diagnosiscode", p.DiagnosisCode},
		{"practitioner", p.Practitioner},
		{"episodeofcare", p.EpisodeOfCare},
		{"length", p.Length},
		{"careteam", p.Careteam},
		{"enddate", p.EndDate},
		{"diagnosisreference", p.DiagnosisReference},
		{"reasonreference", p.ReasonReference},
		{"location", p.Location},
		{"serviceprovider", p.ServiceProvider},
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

//search params for EncounterHistory
type SpEncounterHistory struct {
	Identifier string
	Patient    string
	Subject    string
	Encounter  string
	Status     string
}

func (p SpEncounterHistory) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"encounter", p.Encounter},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "EncounterHistory")
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

//search params for Endpoint
type SpEndpoint struct {
	PayloadType    string
	Identifier     string
	ConnectionType string
	Organization   string
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
		{"connectiontype", p.ConnectionType},
		{"organization", p.Organization},
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

//search params for EnrollmentRequest
type SpEnrollmentRequest struct {
	Identifier string
	Patient    string
	Subject    string
	Status     string
}

func (p SpEnrollmentRequest) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"patient", p.Patient},
		{"subject", p.Subject},
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

//search params for EnrollmentResponse
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

//search params for EpisodeOfCare
type SpEpisodeOfCare struct {
	Date               string
	Identifier         string
	DiagnosisCode      string
	DiagnosisReference string
	Patient            string
	Organization       string
	ReasonReference    string
	Type               string
	CareManager        string
	ReasonCode         string
	IncomingReferral   string
	Status             string
}

func (p SpEpisodeOfCare) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"diagnosiscode", p.DiagnosisCode},
		{"diagnosisreference", p.DiagnosisReference},
		{"patient", p.Patient},
		{"organization", p.Organization},
		{"reasonreference", p.ReasonReference},
		{"type", p.Type},
		{"caremanager", p.CareManager},
		{"reasoncode", p.ReasonCode},
		{"incomingreferral", p.IncomingReferral},
		{"status", p.Status},
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

//search params for EventDefinition
type SpEventDefinition struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	DerivedFrom         string
	Description         string
	ContextType         string
	Predecessor         string
	ComposedOf          string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	DependsOn           string
	Effective           string
	Context             string
	Name                string
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
		{"derivedfrom", p.DerivedFrom},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"composedof", p.ComposedOf},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"dependson", p.DependsOn},
		{"effective", p.Effective},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for Evidence
type SpEvidence struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Description         string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	Context             string
	Publisher           string
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
		{"contexttypevalue", p.ContextTypeValue},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
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

//search params for EvidenceReport
type SpEvidenceReport struct {
	ContextQuantity     string
	Identifier          string
	ContextTypeValue    string
	Context             string
	Publisher           string
	ContextType         string
	ContextTypeQuantity string
	Url                 string
	Status              string
}

func (p SpEvidenceReport) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"contextquantity", p.ContextQuantity},
		{"identifier", p.Identifier},
		{"contexttypevalue", p.ContextTypeValue},
		{"context", p.Context},
		{"publisher", p.Publisher},
		{"contexttype", p.ContextType},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"url", p.Url},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "EvidenceReport")
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

//search params for EvidenceVariable
type SpEvidenceVariable struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	DerivedFrom         string
	Description         string
	ContextType         string
	Predecessor         string
	ComposedOf          string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	DependsOn           string
	Context             string
	Name                string
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
		{"derivedfrom", p.DerivedFrom},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"composedof", p.ComposedOf},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"dependson", p.DependsOn},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for ExampleScenario
type SpExampleScenario struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	ContextType         string
	Version             string
	Url                 string
	ContextQuantity     string
	Context             string
	Name                string
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
		{"context", p.Context},
		{"name", p.Name},
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

//search params for ExplanationOfBenefit
type SpExplanationOfBenefit struct {
	CareTeam     string
	Coverage     string
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
		{"careteam", p.CareTeam},
		{"coverage", p.Coverage},
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

//search params for FamilyMemberHistory
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

//search params for Flag
type SpFlag struct {
	Date       string
	Identifier string
	Author     string
	Patient    string
	Subject    string
	Encounter  string
	Category   string
	Status     string
}

func (p SpFlag) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"author", p.Author},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"encounter", p.Encounter},
		{"category", p.Category},
		{"status", p.Status},
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

//search params for FormularyItem
type SpFormularyItem struct {
	Identifier string
	Code       string
}

func (p SpFormularyItem) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "FormularyItem")
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

//search params for GenomicStudy
type SpGenomicStudy struct {
	Identifier string
	Patient    string
	Subject    string
	Focus      string
	Status     string
}

func (p SpGenomicStudy) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"focus", p.Focus},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "GenomicStudy")
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

//search params for Goal
type SpGoal struct {
	TargetMeasure     string
	Identifier        string
	Addresses         string
	LifecycleStatus   string
	AchievementStatus string
	Patient           string
	Subject           string
	Description       string
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
		{"targetmeasure", p.TargetMeasure},
		{"identifier", p.Identifier},
		{"addresses", p.Addresses},
		{"lifecyclestatus", p.LifecycleStatus},
		{"achievementstatus", p.AchievementStatus},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"description", p.Description},
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

//search params for GraphDefinition
type SpGraphDefinition struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	Start               string
	Description         string
	ContextType         string
	Version             string
	Url                 string
	ContextQuantity     string
	Context             string
	Name                string
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
		{"identifier", p.Identifier},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"start", p.Start},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for Group
type SpGroup struct {
	Identifier              string
	CharacteristicValue     string
	ManagingEntity          string
	Code                    string
	Member                  string
	Name                    string
	Exclude                 string
	Membership              string
	Type                    string
	CharacteristicReference string
	Value                   string
	Characteristic          string
}

func (p SpGroup) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"characteristicvalue", p.CharacteristicValue},
		{"managingentity", p.ManagingEntity},
		{"code", p.Code},
		{"member", p.Member},
		{"name", p.Name},
		{"exclude", p.Exclude},
		{"membership", p.Membership},
		{"type", p.Type},
		{"characteristicreference", p.CharacteristicReference},
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

//search params for GuidanceResponse
type SpGuidanceResponse struct {
	Identifier string
	Request    string
	Patient    string
	Subject    string
	Status     string
}

func (p SpGuidanceResponse) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"request", p.Request},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"status", p.Status},
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

//search params for HealthcareService
type SpHealthcareService struct {
	Identifier      string
	Specialty       string
	ServiceCategory string
	ServiceType     string
	Active          string
	Eligibility     string
	Program         string
	Characteristic  string
	Endpoint        string
	CoverageArea    string
	Organization    string
	OfferedIn       string
	Name            string
	Location        string
	Communication   string
}

func (p SpHealthcareService) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"specialty", p.Specialty},
		{"servicecategory", p.ServiceCategory},
		{"servicetype", p.ServiceType},
		{"active", p.Active},
		{"eligibility", p.Eligibility},
		{"program", p.Program},
		{"characteristic", p.Characteristic},
		{"endpoint", p.Endpoint},
		{"coveragearea", p.CoverageArea},
		{"organization", p.Organization},
		{"offeredin", p.OfferedIn},
		{"name", p.Name},
		{"location", p.Location},
		{"communication", p.Communication},
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

//search params for ImagingSelection
type SpImagingSelection struct {
	Identifier    string
	BodyStructure string
	BasedOn       string
	Code          string
	Subject       string
	Patient       string
	DerivedFrom   string
	Issued        string
	BodySite      string
	StudyUid      string
	Status        string
}

func (p SpImagingSelection) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"bodystructure", p.BodyStructure},
		{"basedon", p.BasedOn},
		{"code", p.Code},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"derivedfrom", p.DerivedFrom},
		{"issued", p.Issued},
		{"bodysite", p.BodySite},
		{"studyuid", p.StudyUid},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ImagingSelection")
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

//search params for ImagingStudy
type SpImagingStudy struct {
	Identifier    string
	Reason        string
	DicomClass    string
	Instance      string
	Modality      string
	Performer     string
	Subject       string
	Started       string
	Encounter     string
	Referrer      string
	BodyStructure string
	Endpoint      string
	BasedOn       string
	Patient       string
	Series        string
	BodySite      string
	Status        string
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
		{"instance", p.Instance},
		{"modality", p.Modality},
		{"performer", p.Performer},
		{"subject", p.Subject},
		{"started", p.Started},
		{"encounter", p.Encounter},
		{"referrer", p.Referrer},
		{"bodystructure", p.BodyStructure},
		{"endpoint", p.Endpoint},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"series", p.Series},
		{"bodysite", p.BodySite},
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

//search params for Immunization
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
	ReactionDate    string
	Status          string
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
		{"reactiondate", p.ReactionDate},
		{"status", p.Status},
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

//search params for ImmunizationEvaluation
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

//search params for ImmunizationRecommendation
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

//search params for ImplementationGuide
type SpImplementationGuide struct {
	Date                string
	Identifier          string
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
	Context             string
	Name                string
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
		{"identifier", p.Identifier},
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
		{"context", p.Context},
		{"name", p.Name},
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

//search params for Ingredient
type SpIngredient struct {
	Identifier                    string
	Role                          string
	Substance                     string
	StrengthConcentrationRatio    string
	For                           string
	SubstanceCode                 string
	StrengthConcentrationQuantity string
	Manufacturer                  string
	SubstanceDefinition           string
	Function                      string
	StrengthPresentationRatio     string
	StrengthPresentationQuantity  string
	Status                        string
}

func (p SpIngredient) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"role", p.Role},
		{"substance", p.Substance},
		{"strengthconcentrationratio", p.StrengthConcentrationRatio},
		{"for", p.For},
		{"substancecode", p.SubstanceCode},
		{"strengthconcentrationquantity", p.StrengthConcentrationQuantity},
		{"manufacturer", p.Manufacturer},
		{"substancedefinition", p.SubstanceDefinition},
		{"function", p.Function},
		{"strengthpresentationratio", p.StrengthPresentationRatio},
		{"strengthpresentationquantity", p.StrengthPresentationQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Ingredient")
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

//search params for InsurancePlan
type SpInsurancePlan struct {
	Identifier        string
	Address           string
	AddressState      string
	OwnedBy           string
	Type              string
	AddressPostalcode string
	AddressCountry    string
	AdministeredBy    string
	Endpoint          string
	Phonetic          string
	AddressUse        string
	Name              string
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
		{"addresscountry", p.AddressCountry},
		{"administeredby", p.AdministeredBy},
		{"endpoint", p.Endpoint},
		{"phonetic", p.Phonetic},
		{"addressuse", p.AddressUse},
		{"name", p.Name},
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

//search params for InventoryItem
type SpInventoryItem struct {
	Identifier string
	Code       string
	Subject    string
	Status     string
}

func (p SpInventoryItem) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"subject", p.Subject},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "InventoryItem")
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

//search params for InventoryReport
type SpInventoryReport struct {
	ItemReference string
	Identifier    string
	Item          string
	Status        string
}

func (p SpInventoryReport) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"itemreference", p.ItemReference},
		{"identifier", p.Identifier},
		{"item", p.Item},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "InventoryReport")
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

//search params for Invoice
type SpInvoice struct {
	Date            string
	Identifier      string
	Totalgross      string
	ParticipantRole string
	Subject         string
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
		{"participantrole", p.ParticipantRole},
		{"subject", p.Subject},
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

//search params for Library
type SpLibrary struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	DerivedFrom         string
	Description         string
	ContextType         string
	Predecessor         string
	ComposedOf          string
	Title               string
	Type                string
	Version             string
	Url                 string
	ContextQuantity     string
	DependsOn           string
	Effective           string
	Context             string
	Name                string
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
		{"derivedfrom", p.DerivedFrom},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"composedof", p.ComposedOf},
		{"title", p.Title},
		{"type", p.Type},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"dependson", p.DependsOn},
		{"effective", p.Effective},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for Linkage
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

//search params for List
type SpList struct {
	Date        string
	Identifier  string
	EmptyReason string
	Item        string
	Code        string
	Notes       string
	Patient     string
	Subject     string
	Encounter   string
	Source      string
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
		{"emptyreason", p.EmptyReason},
		{"item", p.Item},
		{"code", p.Code},
		{"notes", p.Notes},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"encounter", p.Encounter},
		{"source", p.Source},
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

//search params for Location
type SpLocation struct {
	Identifier        string
	Partof            string
	Address           string
	AddressState      string
	OperationalStatus string
	Type              string
	AddressPostalcode string
	Characteristic    string
	AddressCountry    string
	Endpoint          string
	Contains          string
	Organization      string
	AddressUse        string
	Name              string
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
		{"characteristic", p.Characteristic},
		{"addresscountry", p.AddressCountry},
		{"endpoint", p.Endpoint},
		{"contains", p.Contains},
		{"organization", p.Organization},
		{"addressuse", p.AddressUse},
		{"name", p.Name},
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

//search params for ManufacturedItemDefinition
type SpManufacturedItemDefinition struct {
	Identifier string
	Ingredient string
	Name       string
	DoseForm   string
	Status     string
}

func (p SpManufacturedItemDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"ingredient", p.Ingredient},
		{"name", p.Name},
		{"doseform", p.DoseForm},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "ManufacturedItemDefinition")
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

//search params for Measure
type SpMeasure struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	DerivedFrom         string
	Description         string
	ContextType         string
	Predecessor         string
	ComposedOf          string
	Title               string
	Version             string
	Url                 string
	ContextQuantity     string
	DependsOn           string
	Effective           string
	Context             string
	Name                string
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
		{"derivedfrom", p.DerivedFrom},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"composedof", p.ComposedOf},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"dependson", p.DependsOn},
		{"effective", p.Effective},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for MeasureReport
type SpMeasureReport struct {
	Date              string
	Identifier        string
	Period            string
	Measure           string
	Patient           string
	Subject           string
	Reporter          string
	Location          string
	EvaluatedResource string
	Status            string
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
		{"location", p.Location},
		{"evaluatedresource", p.EvaluatedResource},
		{"status", p.Status},
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

//search params for Medication
type SpMedication struct {
	IngredientCode               string
	Identifier                   string
	Code                         string
	Ingredient                   string
	Form                         string
	LotNumber                    string
	SerialNumber                 string
	ExpirationDate               string
	Marketingauthorizationholder string
	Status                       string
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
		{"serialnumber", p.SerialNumber},
		{"expirationdate", p.ExpirationDate},
		{"marketingauthorizationholder", p.Marketingauthorizationholder},
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

//search params for MedicationAdministration
type SpMedicationAdministration struct {
	Date                string
	Identifier          string
	Request             string
	Code                string
	Performer           string
	PerformerDeviceCode string
	Subject             string
	Medication          string
	ReasonGiven         string
	Encounter           string
	ReasonGivenCode     string
	Patient             string
	ReasonNotGiven      string
	Device              string
	Status              string
}

func (p SpMedicationAdministration) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"request", p.Request},
		{"code", p.Code},
		{"performer", p.Performer},
		{"performerdevicecode", p.PerformerDeviceCode},
		{"subject", p.Subject},
		{"medication", p.Medication},
		{"reasongiven", p.ReasonGiven},
		{"encounter", p.Encounter},
		{"reasongivencode", p.ReasonGivenCode},
		{"patient", p.Patient},
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

//search params for MedicationDispense
type SpMedicationDispense struct {
	Identifier       string
	Code             string
	Performer        string
	Receiver         string
	Subject          string
	Destination      string
	Medication       string
	Responsibleparty string
	Encounter        string
	Type             string
	Recorded         string
	Whenhandedover   string
	Whenprepared     string
	Prescription     string
	Patient          string
	Location         string
	Status           string
}

func (p SpMedicationDispense) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"performer", p.Performer},
		{"receiver", p.Receiver},
		{"subject", p.Subject},
		{"destination", p.Destination},
		{"medication", p.Medication},
		{"responsibleparty", p.Responsibleparty},
		{"encounter", p.Encounter},
		{"type", p.Type},
		{"recorded", p.Recorded},
		{"whenhandedover", p.Whenhandedover},
		{"whenprepared", p.Whenprepared},
		{"prescription", p.Prescription},
		{"patient", p.Patient},
		{"location", p.Location},
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

//search params for MedicationKnowledge
type SpMedicationKnowledge struct {
	ProductType           string
	Identifier            string
	Code                  string
	Ingredient            string
	Doseform              string
	ClassificationType    string
	MonographType         string
	Classification        string
	IngredientCode        string
	PackagingCostConcept  string
	SourceCost            string
	MonitoringProgramName string
	Monograph             string
	MonitoringProgramType string
	PackagingCost         string
	Status                string
}

func (p SpMedicationKnowledge) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"producttype", p.ProductType},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"ingredient", p.Ingredient},
		{"doseform", p.Doseform},
		{"classificationtype", p.ClassificationType},
		{"monographtype", p.MonographType},
		{"classification", p.Classification},
		{"ingredientcode", p.IngredientCode},
		{"packagingcostconcept", p.PackagingCostConcept},
		{"sourcecost", p.SourceCost},
		{"monitoringprogramname", p.MonitoringProgramName},
		{"monograph", p.Monograph},
		{"monitoringprogramtype", p.MonitoringProgramType},
		{"packagingcost", p.PackagingCost},
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

//search params for MedicationRequest
type SpMedicationRequest struct {
	Requester             string
	Identifier            string
	IntendedDispenser     string
	Authoredon            string
	Code                  string
	ComboDate             string
	Subject               string
	Medication            string
	Encounter             string
	Priority              string
	Intent                string
	GroupIdentifier       string
	IntendedPerformer     string
	Patient               string
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
		{"identifier", p.Identifier},
		{"intendeddispenser", p.IntendedDispenser},
		{"authoredon", p.Authoredon},
		{"code", p.Code},
		{"combodate", p.ComboDate},
		{"subject", p.Subject},
		{"medication", p.Medication},
		{"encounter", p.Encounter},
		{"priority", p.Priority},
		{"intent", p.Intent},
		{"groupidentifier", p.GroupIdentifier},
		{"intendedperformer", p.IntendedPerformer},
		{"patient", p.Patient},
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

//search params for MedicationStatement
type SpMedicationStatement struct {
	Effective  string
	Identifier string
	Code       string
	Adherence  string
	Patient    string
	Subject    string
	Medication string
	Encounter  string
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
		{"effective", p.Effective},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"adherence", p.Adherence},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"medication", p.Medication},
		{"encounter", p.Encounter},
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

//search params for MedicinalProductDefinition
type SpMedicinalProductDefinition struct {
	Identifier            string
	Ingredient            string
	MasterFile            string
	Contact               string
	Domain                string
	Name                  string
	NameLanguage          string
	Type                  string
	Characteristic        string
	CharacteristicType    string
	ProductClassification string
	Status                string
}

func (p SpMedicinalProductDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"ingredient", p.Ingredient},
		{"masterfile", p.MasterFile},
		{"contact", p.Contact},
		{"domain", p.Domain},
		{"name", p.Name},
		{"namelanguage", p.NameLanguage},
		{"type", p.Type},
		{"characteristic", p.Characteristic},
		{"characteristictype", p.CharacteristicType},
		{"productclassification", p.ProductClassification},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "MedicinalProductDefinition")
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

//search params for MessageDefinition
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
	Context             string
	Name                string
	Publisher           string
	Category            string
	Event               string
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
		{"context", p.Context},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"category", p.Category},
		{"event", p.Event},
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

//search params for MessageHeader
type SpMessageHeader struct {
	Code        string
	Receiver    string
	Sender      string
	Author      string
	Responsible string
	Destination string
	Focus       string
	ResponseId  string
	Source      string
	Event       string
	Target      string
}

func (p SpMessageHeader) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"code", p.Code},
		{"receiver", p.Receiver},
		{"sender", p.Sender},
		{"author", p.Author},
		{"responsible", p.Responsible},
		{"destination", p.Destination},
		{"focus", p.Focus},
		{"responseid", p.ResponseId},
		{"source", p.Source},
		{"event", p.Event},
		{"target", p.Target},
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

//search params for MolecularSequence
type SpMolecularSequence struct {
	Identifier string
	Subject    string
	Patient    string
	Focus      string
	Type       string
}

func (p SpMolecularSequence) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"focus", p.Focus},
		{"type", p.Type},
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

//search params for NamingSystem
type SpNamingSystem struct {
	Date                string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Type                string
	ContextQuantity     string
	Effective           string
	Contact             string
	Responsible         string
	Context             string
	Telecom             string
	Value               string
	ContextTypeQuantity string
	Identifier          string
	Period              string
	Kind                string
	Version             string
	Url                 string
	IdType              string
	Name                string
	Publisher           string
	Topic               string
	Status              string
}

func (p SpNamingSystem) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"type", p.Type},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"contact", p.Contact},
		{"responsible", p.Responsible},
		{"context", p.Context},
		{"telecom", p.Telecom},
		{"value", p.Value},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"identifier", p.Identifier},
		{"period", p.Period},
		{"kind", p.Kind},
		{"version", p.Version},
		{"url", p.Url},
		{"idtype", p.IdType},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"topic", p.Topic},
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

//search params for NutritionIntake
type SpNutritionIntake struct {
	Date       string
	Identifier string
	Nutrition  string
	Code       string
	Patient    string
	Subject    string
	Encounter  string
	Source     string
	Status     string
}

func (p SpNutritionIntake) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"nutrition", p.Nutrition},
		{"code", p.Code},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"encounter", p.Encounter},
		{"source", p.Source},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "NutritionIntake")
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

//search params for NutritionOrder
type SpNutritionOrder struct {
	Identifier      string
	GroupIdentifier string
	Datetime        string
	Provider        string
	Subject         string
	Patient         string
	Supplement      string
	Formula         string
	Encounter       string
	Oraldiet        string
	Additive        string
	Status          string
}

func (p SpNutritionOrder) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"groupidentifier", p.GroupIdentifier},
		{"datetime", p.Datetime},
		{"provider", p.Provider},
		{"subject", p.Subject},
		{"patient", p.Patient},
		{"supplement", p.Supplement},
		{"formula", p.Formula},
		{"encounter", p.Encounter},
		{"oraldiet", p.Oraldiet},
		{"additive", p.Additive},
		{"status", p.Status},
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

//search params for NutritionProduct
type SpNutritionProduct struct {
	Identifier   string
	Code         string
	LotNumber    string
	SerialNumber string
	Status       string
}

func (p SpNutritionProduct) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"lotnumber", p.LotNumber},
		{"serialnumber", p.SerialNumber},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "NutritionProduct")
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

//search params for Observation
type SpObservation struct {
	Date                       string
	ComboDataAbsentReason      string
	Code                       string
	ComboCodeValueQuantity     string
	ComponentDataAbsentReason  string
	Subject                    string
	ValueConcept               string
	ValueDate                  string
	DerivedFrom                string
	Focus                      string
	PartOf                     string
	ComponentValueCanonical    string
	HasMember                  string
	ValueReference             string
	CodeValueString            string
	ComponentCodeValueQuantity string
	BasedOn                    string
	CodeValueDate              string
	Patient                    string
	Specimen                   string
	CodeValueQuantity          string
	ComponentCode              string
	ValueMarkdown              string
	ComboCodeValueConcept      string
	Identifier                 string
	ComponentValueReference    string
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
	ValueCanonical             string
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
		{"componentdataabsentreason", p.ComponentDataAbsentReason},
		{"subject", p.Subject},
		{"valueconcept", p.ValueConcept},
		{"valuedate", p.ValueDate},
		{"derivedfrom", p.DerivedFrom},
		{"focus", p.Focus},
		{"partof", p.PartOf},
		{"componentvaluecanonical", p.ComponentValueCanonical},
		{"hasmember", p.HasMember},
		{"valuereference", p.ValueReference},
		{"codevaluestring", p.CodeValueString},
		{"componentcodevaluequantity", p.ComponentCodeValueQuantity},
		{"basedon", p.BasedOn},
		{"codevaluedate", p.CodeValueDate},
		{"patient", p.Patient},
		{"specimen", p.Specimen},
		{"codevaluequantity", p.CodeValueQuantity},
		{"componentcode", p.ComponentCode},
		{"valuemarkdown", p.ValueMarkdown},
		{"combocodevalueconcept", p.ComboCodeValueConcept},
		{"identifier", p.Identifier},
		{"componentvaluereference", p.ComponentValueReference},
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
		{"valuecanonical", p.ValueCanonical},
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

//search params for ObservationDefinition
type SpObservationDefinition struct {
	Identifier   string
	Code         string
	Method       string
	Experimental string
	Category     string
	Title        string
	Url          string
	Status       string
}

func (p SpObservationDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"method", p.Method},
		{"experimental", p.Experimental},
		{"category", p.Category},
		{"title", p.Title},
		{"url", p.Url},
		{"status", p.Status},
	}
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

//search params for OperationDefinition
type SpOperationDefinition struct {
	Date                string
	Identifier          string
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
	Context             string
	Name                string
	Publisher           string
	ContextTypeQuantity string
	Base                string
	Status              string
}

func (p SpOperationDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
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
		{"context", p.Context},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"base", p.Base},
		{"status", p.Status},
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

//search params for OperationOutcome
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

//search params for Organization
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
	AddressUse        string
	Name              string
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
		{"addressuse", p.AddressUse},
		{"name", p.Name},
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

//search params for OrganizationAffiliation
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
	Location                  string
	Telecom                   string
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
		{"location", p.Location},
		{"telecom", p.Telecom},
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

//search params for PackagedProductDefinition
type SpPackagedProductDefinition struct {
	Identifier       string
	ManufacturedItem string
	Nutrition        string
	Package          string
	Name             string
	Biological       string
	PackageFor       string
	ContainedItem    string
	Medication       string
	Device           string
	Status           string
}

func (p SpPackagedProductDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"manufactureditem", p.ManufacturedItem},
		{"nutrition", p.Nutrition},
		{"package", p.Package},
		{"name", p.Name},
		{"biological", p.Biological},
		{"packagefor", p.PackageFor},
		{"containeditem", p.ContainedItem},
		{"medication", p.Medication},
		{"device", p.Device},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "PackagedProductDefinition")
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

//search params for Patient
type SpPatient struct {
	Given               string
	Identifier          string
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
	AddressUse          string
	Name                string
	Telecom             string
	AddressCity         string
	Family              string
	Email               string
}

func (p SpPatient) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"given", p.Given},
		{"identifier", p.Identifier},
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
		{"addressuse", p.AddressUse},
		{"name", p.Name},
		{"telecom", p.Telecom},
		{"addresscity", p.AddressCity},
		{"family", p.Family},
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

//search params for PaymentNotice
type SpPaymentNotice struct {
	Identifier    string
	Request       string
	Created       string
	Response      string
	Reporter      string
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
		{"created", p.Created},
		{"response", p.Response},
		{"reporter", p.Reporter},
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

//search params for PaymentReconciliation
type SpPaymentReconciliation struct {
	Identifier          string
	Request             string
	Disposition         string
	Created             string
	AllocationEncounter string
	AllocationAccount   string
	Outcome             string
	PaymentIssuer       string
	Requestor           string
	Status              string
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
		{"allocationencounter", p.AllocationEncounter},
		{"allocationaccount", p.AllocationAccount},
		{"outcome", p.Outcome},
		{"paymentissuer", p.PaymentIssuer},
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

//search params for Permission
type SpPermission struct {
	Status string
}

func (p SpPermission) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Permission")
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

//search params for Person
type SpPerson struct {
	Identifier        string
	Given             string
	Address           string
	Birthdate         string
	Deceased          string
	AddressState      string
	Gender            string
	Practitioner      string
	Link              string
	Relatedperson     string
	AddressPostalcode string
	AddressCountry    string
	DeathDate         string
	Phonetic          string
	Phone             string
	Patient           string
	Organization      string
	AddressUse        string
	Name              string
	Telecom           string
	AddressCity       string
	Family            string
	Email             string
}

func (p SpPerson) SpEncode(baseURL *string) (*string, error) {
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
		{"practitioner", p.Practitioner},
		{"link", p.Link},
		{"relatedperson", p.Relatedperson},
		{"addresspostalcode", p.AddressPostalcode},
		{"addresscountry", p.AddressCountry},
		{"deathdate", p.DeathDate},
		{"phonetic", p.Phonetic},
		{"phone", p.Phone},
		{"patient", p.Patient},
		{"organization", p.Organization},
		{"addressuse", p.AddressUse},
		{"name", p.Name},
		{"telecom", p.Telecom},
		{"addresscity", p.AddressCity},
		{"family", p.Family},
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

//search params for PlanDefinition
type SpPlanDefinition struct {
	Date                string
	Identifier          string
	Successor           string
	ContextTypeValue    string
	Jurisdiction        string
	DerivedFrom         string
	Description         string
	ContextType         string
	Predecessor         string
	ComposedOf          string
	Title               string
	Type                string
	Version             string
	Url                 string
	ContextQuantity     string
	DependsOn           string
	Effective           string
	Context             string
	Name                string
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
		{"derivedfrom", p.DerivedFrom},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"composedof", p.ComposedOf},
		{"title", p.Title},
		{"type", p.Type},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"dependson", p.DependsOn},
		{"effective", p.Effective},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for Practitioner
type SpPractitioner struct {
	Given               string
	Identifier          string
	Address             string
	Deceased            string
	AddressState        string
	Gender              string
	QualificationPeriod string
	Active              string
	AddressPostalcode   string
	AddressCountry      string
	DeathDate           string
	Phonetic            string
	Phone               string
	AddressUse          string
	Name                string
	Telecom             string
	AddressCity         string
	Communication       string
	Family              string
	Email               string
}

func (p SpPractitioner) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"given", p.Given},
		{"identifier", p.Identifier},
		{"address", p.Address},
		{"deceased", p.Deceased},
		{"addressstate", p.AddressState},
		{"gender", p.Gender},
		{"qualificationperiod", p.QualificationPeriod},
		{"active", p.Active},
		{"addresspostalcode", p.AddressPostalcode},
		{"addresscountry", p.AddressCountry},
		{"deathdate", p.DeathDate},
		{"phonetic", p.Phonetic},
		{"phone", p.Phone},
		{"addressuse", p.AddressUse},
		{"name", p.Name},
		{"telecom", p.Telecom},
		{"addresscity", p.AddressCity},
		{"communication", p.Communication},
		{"family", p.Family},
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

//search params for PractitionerRole
type SpPractitionerRole struct {
	Date           string
	Identifier     string
	Specialty      string
	Role           string
	Practitioner   string
	Active         string
	Characteristic string
	Endpoint       string
	Phone          string
	Service        string
	Organization   string
	Location       string
	Telecom        string
	Communication  string
	Email          string
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
		{"characteristic", p.Characteristic},
		{"endpoint", p.Endpoint},
		{"phone", p.Phone},
		{"service", p.Service},
		{"organization", p.Organization},
		{"location", p.Location},
		{"telecom", p.Telecom},
		{"communication", p.Communication},
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

//search params for Procedure
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
	Report                string
	InstantiatesUri       string
	Location              string
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
		{"report", p.Report},
		{"instantiatesuri", p.InstantiatesUri},
		{"location", p.Location},
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

//search params for Provenance
type SpProvenance struct {
	AgentType     string
	Agent         string
	SignatureType string
	Activity      string
	Encounter     string
	Recorded      string
	When          string
	Target        string
	BasedOn       string
	Patient       string
	Location      string
	AgentRole     string
	Entity        string
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
		{"activity", p.Activity},
		{"encounter", p.Encounter},
		{"recorded", p.Recorded},
		{"when", p.When},
		{"target", p.Target},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"location", p.Location},
		{"agentrole", p.AgentRole},
		{"entity", p.Entity},
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

//search params for Questionnaire
type SpQuestionnaire struct {
	Date                string
	Identifier          string
	ComboCode           string
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
	Context             string
	Name                string
	Publisher           string
	QuestionnaireCode   string
	Definition          string
	ContextTypeQuantity string
	ItemCode            string
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
		{"combocode", p.ComboCode},
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
		{"context", p.Context},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"questionnairecode", p.QuestionnaireCode},
		{"definition", p.Definition},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"itemcode", p.ItemCode},
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

//search params for QuestionnaireResponse
type SpQuestionnaireResponse struct {
	Authored      string
	Identifier    string
	Questionnaire string
	BasedOn       string
	Author        string
	Patient       string
	Subject       string
	PartOf        string
	Encounter     string
	Source        string
	ItemSubject   string
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
		{"author", p.Author},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"partof", p.PartOf},
		{"encounter", p.Encounter},
		{"source", p.Source},
		{"itemsubject", p.ItemSubject},
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

//search params for RegulatedAuthorization
type SpRegulatedAuthorization struct {
	Identifier string
	Subject    string
	CaseType   string
	Holder     string
	Region     string
	Case       string
	Status     string
}

func (p SpRegulatedAuthorization) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"subject", p.Subject},
		{"casetype", p.CaseType},
		{"holder", p.Holder},
		{"region", p.Region},
		{"case", p.Case},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "RegulatedAuthorization")
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

//search params for RelatedPerson
type SpRelatedPerson struct {
	Identifier        string
	Given             string
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
	AddressUse        string
	Name              string
	Telecom           string
	AddressCity       string
	Family            string
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
		{"given", p.Given},
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
		{"addressuse", p.AddressUse},
		{"name", p.Name},
		{"telecom", p.Telecom},
		{"addresscity", p.AddressCity},
		{"family", p.Family},
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

//search params for RequestOrchestration
type SpRequestOrchestration struct {
	Authored              string
	Identifier            string
	Code                  string
	Author                string
	Subject               string
	InstantiatesCanonical string
	Encounter             string
	Priority              string
	Intent                string
	Participant           string
	GroupIdentifier       string
	BasedOn               string
	Patient               string
	InstantiatesUri       string
	Status                string
}

func (p SpRequestOrchestration) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"authored", p.Authored},
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"author", p.Author},
		{"subject", p.Subject},
		{"instantiatescanonical", p.InstantiatesCanonical},
		{"encounter", p.Encounter},
		{"priority", p.Priority},
		{"intent", p.Intent},
		{"participant", p.Participant},
		{"groupidentifier", p.GroupIdentifier},
		{"basedon", p.BasedOn},
		{"patient", p.Patient},
		{"instantiatesuri", p.InstantiatesUri},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "RequestOrchestration")
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

//search params for Requirements
type SpRequirements struct {
	Date                string
	Identifier          string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Title               string
	Version             string
	Url                 string
	Actor               string
	ContextQuantity     string
	Context             string
	Name                string
	Publisher           string
	ContextTypeQuantity string
	Status              string
}

func (p SpRequirements) SpEncode(baseURL *string) (*string, error) {
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
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"actor", p.Actor},
		{"contextquantity", p.ContextQuantity},
		{"context", p.Context},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Requirements")
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

//search params for ResearchStudy
type SpResearchStudy struct {
	Date                            string
	ObjectiveType                   string
	StudyDesign                     string
	Description                     string
	Eligibility                     string
	PartOf                          string
	Title                           string
	ProgressStatusStatePeriodActual string
	RecruitmentTarget               string
	Protocol                        string
	Classifier                      string
	Keyword                         string
	FocusCode                       string
	Phase                           string
	Identifier                      string
	ProgressStatusStateActual       string
	FocusReference                  string
	ObjectiveDescription            string
	ProgressStatusStatePeriod       string
	Condition                       string
	Site                            string
	Name                            string
	RecruitmentActual               string
	Region                          string
	Status                          string
}

func (p SpResearchStudy) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"objectivetype", p.ObjectiveType},
		{"studydesign", p.StudyDesign},
		{"description", p.Description},
		{"eligibility", p.Eligibility},
		{"partof", p.PartOf},
		{"title", p.Title},
		{"progressstatusstateperiodactual", p.ProgressStatusStatePeriodActual},
		{"recruitmenttarget", p.RecruitmentTarget},
		{"protocol", p.Protocol},
		{"classifier", p.Classifier},
		{"keyword", p.Keyword},
		{"focuscode", p.FocusCode},
		{"phase", p.Phase},
		{"identifier", p.Identifier},
		{"progressstatusstateactual", p.ProgressStatusStateActual},
		{"focusreference", p.FocusReference},
		{"objectivedescription", p.ObjectiveDescription},
		{"progressstatusstateperiod", p.ProgressStatusStatePeriod},
		{"condition", p.Condition},
		{"site", p.Site},
		{"name", p.Name},
		{"recruitmentactual", p.RecruitmentActual},
		{"region", p.Region},
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

//search params for ResearchSubject
type SpResearchSubject struct {
	Date          string
	Identifier    string
	Subject_state string
	Study         string
	Subject       string
	Patient       string
	Status        string
}

func (p SpResearchSubject) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"identifier", p.Identifier},
		{"subject_state", p.Subject_state},
		{"study", p.Study},
		{"subject", p.Subject},
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

//search params for RiskAssessment
type SpRiskAssessment struct {
	Date        string
	Identifier  string
	Condition   string
	Performer   string
	Method      string
	Patient     string
	Probability string
	Subject     string
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
		{"patient", p.Patient},
		{"probability", p.Probability},
		{"subject", p.Subject},
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

//search params for Schedule
type SpSchedule struct {
	Actor                string
	Date                 string
	Identifier           string
	Specialty            string
	ServiceCategory      string
	ServiceType          string
	Name                 string
	Active               string
	ServiceTypeReference string
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
		{"name", p.Name},
		{"active", p.Active},
		{"servicetypereference", p.ServiceTypeReference},
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

//search params for SearchParameter
type SpSearchParameter struct {
	Date                string
	Identifier          string
	Code                string
	ContextTypeValue    string
	Jurisdiction        string
	DerivedFrom         string
	Description         string
	ContextType         string
	Type                string
	Version             string
	Url                 string
	Target              string
	ContextQuantity     string
	Component           string
	Context             string
	Name                string
	Publisher           string
	ContextTypeQuantity string
	Base                string
	Status              string
}

func (p SpSearchParameter) SpEncode(baseURL *string) (*string, error) {
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
		{"derivedfrom", p.DerivedFrom},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"type", p.Type},
		{"version", p.Version},
		{"url", p.Url},
		{"target", p.Target},
		{"contextquantity", p.ContextQuantity},
		{"component", p.Component},
		{"context", p.Context},
		{"name", p.Name},
		{"publisher", p.Publisher},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"base", p.Base},
		{"status", p.Status},
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

//search params for ServiceRequest
type SpServiceRequest struct {
	Authored              string
	Requester             string
	Identifier            string
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
	BodyStructure         string
	BasedOn               string
	CodeReference         string
	Patient               string
	Specimen              string
	CodeConcept           string
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
		{"bodystructure", p.BodyStructure},
		{"basedon", p.BasedOn},
		{"codereference", p.CodeReference},
		{"patient", p.Patient},
		{"specimen", p.Specimen},
		{"codeconcept", p.CodeConcept},
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

//search params for Slot
type SpSlot struct {
	Identifier           string
	Schedule             string
	Specialty            string
	ServiceCategory      string
	AppointmentType      string
	ServiceType          string
	Start                string
	ServiceTypeReference string
	Status               string
}

func (p SpSlot) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"schedule", p.Schedule},
		{"specialty", p.Specialty},
		{"servicecategory", p.ServiceCategory},
		{"appointmenttype", p.AppointmentType},
		{"servicetype", p.ServiceType},
		{"start", p.Start},
		{"servicetypereference", p.ServiceTypeReference},
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

//search params for Specimen
type SpSpecimen struct {
	Identifier      string
	Parent          string
	Bodysite        string
	Patient         string
	Subject         string
	Collected       string
	Accession       string
	Procedure       string
	Type            string
	Collector       string
	ContainerDevice string
	Status          string
}

func (p SpSpecimen) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"parent", p.Parent},
		{"bodysite", p.Bodysite},
		{"patient", p.Patient},
		{"subject", p.Subject},
		{"collected", p.Collected},
		{"accession", p.Accession},
		{"procedure", p.Procedure},
		{"type", p.Type},
		{"collector", p.Collector},
		{"containerdevice", p.ContainerDevice},
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

//search params for SpecimenDefinition
type SpSpecimenDefinition struct {
	Container    string
	Identifier   string
	IsDerived    string
	Experimental string
	TypeTested   string
	Title        string
	Type         string
	Url          string
	Status       string
}

func (p SpSpecimenDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"container", p.Container},
		{"identifier", p.Identifier},
		{"isderived", p.IsDerived},
		{"experimental", p.Experimental},
		{"typetested", p.TypeTested},
		{"title", p.Title},
		{"type", p.Type},
		{"url", p.Url},
		{"status", p.Status},
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

//search params for StructureDefinition
type SpStructureDefinition struct {
	Date                 string
	ContextTypeValue     string
	ExtContextType       string
	Jurisdiction         string
	Description          string
	ContextType          string
	Experimental         string
	Title                string
	Type                 string
	ContextQuantity      string
	Path                 string
	BasePath             string
	Context              string
	Keyword              string
	ContextTypeQuantity  string
	ExtContextExpression string
	Identifier           string
	Valueset             string
	Kind                 string
	Abstract             string
	Version              string
	Url                  string
	ExtContext           string
	Name                 string
	Publisher            string
	Derivation           string
	Base                 string
	Status               string
}

func (p SpStructureDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"contexttypevalue", p.ContextTypeValue},
		{"extcontexttype", p.ExtContextType},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"experimental", p.Experimental},
		{"title", p.Title},
		{"type", p.Type},
		{"contextquantity", p.ContextQuantity},
		{"path", p.Path},
		{"basepath", p.BasePath},
		{"context", p.Context},
		{"keyword", p.Keyword},
		{"contexttypequantity", p.ContextTypeQuantity},
		{"extcontextexpression", p.ExtContextExpression},
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
		{"base", p.Base},
		{"status", p.Status},
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

//search params for StructureMap
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
	Context             string
	Name                string
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
		{"context", p.Context},
		{"name", p.Name},
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

//search params for Subscription
type SpSubscription struct {
	Owner        string
	Identifier   string
	Payload      string
	Contact      string
	Name         string
	Topic        string
	FilterValue  string
	Type         string
	ContentLevel string
	Url          string
	Status       string
}

func (p SpSubscription) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"owner", p.Owner},
		{"identifier", p.Identifier},
		{"payload", p.Payload},
		{"contact", p.Contact},
		{"name", p.Name},
		{"topic", p.Topic},
		{"filtervalue", p.FilterValue},
		{"type", p.Type},
		{"contentlevel", p.ContentLevel},
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

//search params for SubscriptionStatus
type SpSubscriptionStatus struct {
}

func (p SpSubscriptionStatus) SpEncode(baseURL *string) (*string, error) {
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

	resUrl, err := url.JoinPath(*baseURL, "SubscriptionStatus")
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

//search params for SubscriptionTopic
type SpSubscriptionTopic struct {
	Date               string
	Effective          string
	Identifier         string
	Resource           string
	DerivedOrSelf      string
	Event              string
	Title              string
	Version            string
	Url                string
	Status             string
	TriggerDescription string
}

func (p SpSubscriptionTopic) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"date", p.Date},
		{"effective", p.Effective},
		{"identifier", p.Identifier},
		{"resource", p.Resource},
		{"derivedorself", p.DerivedOrSelf},
		{"event", p.Event},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"status", p.Status},
		{"triggerdescription", p.TriggerDescription},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SubscriptionTopic")
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

//search params for Substance
type SpSubstance struct {
	Identifier         string
	Code               string
	CodeReference      string
	Quantity           string
	SubstanceReference string
	Expiry             string
	Category           string
	Status             string
}

func (p SpSubstance) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"codereference", p.CodeReference},
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

//search params for SubstanceDefinition
type SpSubstanceDefinition struct {
	Identifier     string
	Code           string
	Domain         string
	Name           string
	Classification string
}

func (p SpSubstanceDefinition) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"code", p.Code},
		{"domain", p.Domain},
		{"name", p.Name},
		{"classification", p.Classification},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "SubstanceDefinition")
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

//search params for SubstanceNucleicAcid
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

//search params for SubstancePolymer
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

//search params for SubstanceProtein
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

//search params for SubstanceReferenceInformation
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

//search params for SubstanceSourceMaterial
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

//search params for SupplyDelivery
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

//search params for SupplyRequest
type SpSupplyRequest struct {
	Date       string
	Requester  string
	Identifier string
	Patient    string
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
		{"date", p.Date},
		{"requester", p.Requester},
		{"identifier", p.Identifier},
		{"patient", p.Patient},
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

//search params for Task
type SpTask struct {
	Owner                       string
	RequestedperformerReference string
	Requester                   string
	BusinessStatus              string
	Identifier                  string
	Period                      string
	Code                        string
	Performer                   string
	Subject                     string
	Focus                       string
	PartOf                      string
	Encounter                   string
	AuthoredOn                  string
	Priority                    string
	Intent                      string
	Output                      string
	Actor                       string
	GroupIdentifier             string
	BasedOn                     string
	Patient                     string
	Modified                    string
	Status                      string
}

func (p SpTask) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"owner", p.Owner},
		{"requestedperformerreference", p.RequestedperformerReference},
		{"requester", p.Requester},
		{"businessstatus", p.BusinessStatus},
		{"identifier", p.Identifier},
		{"period", p.Period},
		{"code", p.Code},
		{"performer", p.Performer},
		{"subject", p.Subject},
		{"focus", p.Focus},
		{"partof", p.PartOf},
		{"encounter", p.Encounter},
		{"authoredon", p.AuthoredOn},
		{"priority", p.Priority},
		{"intent", p.Intent},
		{"output", p.Output},
		{"actor", p.Actor},
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

//search params for TerminologyCapabilities
type SpTerminologyCapabilities struct {
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
	Context             string
	Name                string
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
		{"identifier", p.Identifier},
		{"contexttypevalue", p.ContextTypeValue},
		{"jurisdiction", p.Jurisdiction},
		{"description", p.Description},
		{"contexttype", p.ContextType},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for TestPlan
type SpTestPlan struct {
	Identifier string
	Scope      string
	Url        string
	Status     string
}

func (p SpTestPlan) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"scope", p.Scope},
		{"url", p.Url},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "TestPlan")
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

//search params for TestReport
type SpTestReport struct {
	Result      string
	Identifier  string
	Tester      string
	Testscript  string
	Issued      string
	Participant string
	Status      string
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
		{"status", p.Status},
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

//search params for TestScript
type SpTestScript struct {
	Date                     string
	Identifier               string
	ContextTypeValue         string
	Jurisdiction             string
	Description              string
	TestscriptCapability     string
	ContextType              string
	ScopeArtifactPhase       string
	Title                    string
	ScopeArtifactConformance string
	Version                  string
	ScopeArtifact            string
	Url                      string
	ContextQuantity          string
	Context                  string
	Name                     string
	Publisher                string
	ContextTypeQuantity      string
	Status                   string
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
		{"scopeartifactphase", p.ScopeArtifactPhase},
		{"title", p.Title},
		{"scopeartifactconformance", p.ScopeArtifactConformance},
		{"version", p.Version},
		{"scopeartifact", p.ScopeArtifact},
		{"url", p.Url},
		{"contextquantity", p.ContextQuantity},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for Transport
type SpTransport struct {
	Identifier string
	Status     string
}

func (p SpTransport) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"identifier", p.Identifier},
		{"status", p.Status},
	}
	for _, kv := range params {
		if kv.value != "" {
			v.Set(kv.key, kv.value)
		}
	}

	resUrl, err := url.JoinPath(*baseURL, "Transport")
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

//search params for ValueSet
type SpValueSet struct {
	Date                string
	Identifier          string
	Code                string
	ContextTypeValue    string
	Jurisdiction        string
	Description         string
	DerivedFrom         string
	ContextType         string
	Predecessor         string
	Title               string
	Version             string
	Url                 string
	Expansion           string
	Reference           string
	ContextQuantity     string
	Effective           string
	Context             string
	Name                string
	Publisher           string
	Topic               string
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
		{"derivedfrom", p.DerivedFrom},
		{"contexttype", p.ContextType},
		{"predecessor", p.Predecessor},
		{"title", p.Title},
		{"version", p.Version},
		{"url", p.Url},
		{"expansion", p.Expansion},
		{"reference", p.Reference},
		{"contextquantity", p.ContextQuantity},
		{"effective", p.Effective},
		{"context", p.Context},
		{"name", p.Name},
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

//search params for VerificationResult
type SpVerificationResult struct {
	StatusDate            string
	PrimarysourceWho      string
	PrimarysourceDate     string
	ValidatorOrganization string
	AttestationMethod     string
	AttestationOnbehalfof string
	Target                string
	AttestationWho        string
	PrimarysourceType     string
	Status                string
}

func (p SpVerificationResult) SpEncode(baseURL *string) (*string, error) {
	v := url.Values{}

	params := []struct {
		key   string
		value string
	}{
		{"statusdate", p.StatusDate},
		{"primarysourcewho", p.PrimarysourceWho},
		{"primarysourcedate", p.PrimarysourceDate},
		{"validatororganization", p.ValidatorOrganization},
		{"attestationmethod", p.AttestationMethod},
		{"attestationonbehalfof", p.AttestationOnbehalfof},
		{"target", p.Target},
		{"attestationwho", p.AttestationWho},
		{"primarysourcetype", p.PrimarysourceType},
		{"status", p.Status},
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

//search params for VisionPrescription
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
