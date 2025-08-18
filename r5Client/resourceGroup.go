package r5Client

import r5 "github.com/PotatoEMR/simple-fhir-client/r5"
import "errors"

type ResourceGroup struct {
	Accounts                            []*r5.Account
	ActivityDefinitions                 []*r5.ActivityDefinition
	ActorDefinitions                    []*r5.ActorDefinition
	AdministrableProductDefinitions     []*r5.AdministrableProductDefinition
	AdverseEvents                       []*r5.AdverseEvent
	AllergyIntolerances                 []*r5.AllergyIntolerance
	Appointments                        []*r5.Appointment
	AppointmentResponses                []*r5.AppointmentResponse
	ArtifactAssessments                 []*r5.ArtifactAssessment
	AuditEvents                         []*r5.AuditEvent
	Basics                              []*r5.Basic
	Binarys                             []*r5.Binary
	BiologicallyDerivedProducts         []*r5.BiologicallyDerivedProduct
	BiologicallyDerivedProductDispenses []*r5.BiologicallyDerivedProductDispense
	BodyStructures                      []*r5.BodyStructure
	CanonicalResources                  []*r5.CanonicalResource
	CapabilityStatements                []*r5.CapabilityStatement
	CarePlans                           []*r5.CarePlan
	CareTeams                           []*r5.CareTeam
	ChargeItems                         []*r5.ChargeItem
	ChargeItemDefinitions               []*r5.ChargeItemDefinition
	Citations                           []*r5.Citation
	Claims                              []*r5.Claim
	ClaimResponses                      []*r5.ClaimResponse
	ClinicalImpressions                 []*r5.ClinicalImpression
	ClinicalUseDefinitions              []*r5.ClinicalUseDefinition
	CodeSystems                         []*r5.CodeSystem
	Communications                      []*r5.Communication
	CommunicationRequests               []*r5.CommunicationRequest
	CompartmentDefinitions              []*r5.CompartmentDefinition
	Compositions                        []*r5.Composition
	ConceptMaps                         []*r5.ConceptMap
	Conditions                          []*r5.Condition
	ConditionDefinitions                []*r5.ConditionDefinition
	Consents                            []*r5.Consent
	Contracts                           []*r5.Contract
	Coverages                           []*r5.Coverage
	CoverageEligibilityRequests         []*r5.CoverageEligibilityRequest
	CoverageEligibilityResponses        []*r5.CoverageEligibilityResponse
	DetectedIssues                      []*r5.DetectedIssue
	Devices                             []*r5.Device
	DeviceAssociations                  []*r5.DeviceAssociation
	DeviceDefinitions                   []*r5.DeviceDefinition
	DeviceDispenses                     []*r5.DeviceDispense
	DeviceMetrics                       []*r5.DeviceMetric
	DeviceRequests                      []*r5.DeviceRequest
	DeviceUsages                        []*r5.DeviceUsage
	DiagnosticReports                   []*r5.DiagnosticReport
	DocumentReferences                  []*r5.DocumentReference
	Encounters                          []*r5.Encounter
	EncounterHistorys                   []*r5.EncounterHistory
	Endpoints                           []*r5.Endpoint
	EnrollmentRequests                  []*r5.EnrollmentRequest
	EnrollmentResponses                 []*r5.EnrollmentResponse
	EpisodeOfCares                      []*r5.EpisodeOfCare
	EventDefinitions                    []*r5.EventDefinition
	Evidences                           []*r5.Evidence
	EvidenceReports                     []*r5.EvidenceReport
	EvidenceVariables                   []*r5.EvidenceVariable
	ExampleScenarios                    []*r5.ExampleScenario
	ExplanationOfBenefits               []*r5.ExplanationOfBenefit
	FamilyMemberHistorys                []*r5.FamilyMemberHistory
	Flags                               []*r5.Flag
	FormularyItems                      []*r5.FormularyItem
	GenomicStudys                       []*r5.GenomicStudy
	Goals                               []*r5.Goal
	GraphDefinitions                    []*r5.GraphDefinition
	Groups                              []*r5.Group
	GuidanceResponses                   []*r5.GuidanceResponse
	HealthcareServices                  []*r5.HealthcareService
	ImagingSelections                   []*r5.ImagingSelection
	ImagingStudys                       []*r5.ImagingStudy
	Immunizations                       []*r5.Immunization
	ImmunizationEvaluations             []*r5.ImmunizationEvaluation
	ImmunizationRecommendations         []*r5.ImmunizationRecommendation
	ImplementationGuides                []*r5.ImplementationGuide
	Ingredients                         []*r5.Ingredient
	InsurancePlans                      []*r5.InsurancePlan
	InventoryItems                      []*r5.InventoryItem
	InventoryReports                    []*r5.InventoryReport
	Invoices                            []*r5.Invoice
	Librarys                            []*r5.Library
	Linkages                            []*r5.Linkage
	Lists                               []*r5.List
	Locations                           []*r5.Location
	ManufacturedItemDefinitions         []*r5.ManufacturedItemDefinition
	Measures                            []*r5.Measure
	MeasureReports                      []*r5.MeasureReport
	Medications                         []*r5.Medication
	MedicationAdministrations           []*r5.MedicationAdministration
	MedicationDispenses                 []*r5.MedicationDispense
	MedicationKnowledges                []*r5.MedicationKnowledge
	MedicationRequests                  []*r5.MedicationRequest
	MedicationStatements                []*r5.MedicationStatement
	MedicinalProductDefinitions         []*r5.MedicinalProductDefinition
	MessageDefinitions                  []*r5.MessageDefinition
	MessageHeaders                      []*r5.MessageHeader
	MolecularSequences                  []*r5.MolecularSequence
	NamingSystems                       []*r5.NamingSystem
	NutritionIntakes                    []*r5.NutritionIntake
	NutritionOrders                     []*r5.NutritionOrder
	NutritionProducts                   []*r5.NutritionProduct
	Observations                        []*r5.Observation
	ObservationDefinitions              []*r5.ObservationDefinition
	OperationDefinitions                []*r5.OperationDefinition
	OperationOutcomes                   []*r5.OperationOutcome
	Organizations                       []*r5.Organization
	OrganizationAffiliations            []*r5.OrganizationAffiliation
	PackagedProductDefinitions          []*r5.PackagedProductDefinition
	Patients                            []*r5.Patient
	PaymentNotices                      []*r5.PaymentNotice
	PaymentReconciliations              []*r5.PaymentReconciliation
	Permissions                         []*r5.Permission
	Persons                             []*r5.Person
	PlanDefinitions                     []*r5.PlanDefinition
	Practitioners                       []*r5.Practitioner
	PractitionerRoles                   []*r5.PractitionerRole
	Procedures                          []*r5.Procedure
	Provenances                         []*r5.Provenance
	Questionnaires                      []*r5.Questionnaire
	QuestionnaireResponses              []*r5.QuestionnaireResponse
	RegulatedAuthorizations             []*r5.RegulatedAuthorization
	RelatedPersons                      []*r5.RelatedPerson
	RequestOrchestrations               []*r5.RequestOrchestration
	Requirementss                       []*r5.Requirements
	ResearchStudys                      []*r5.ResearchStudy
	ResearchSubjects                    []*r5.ResearchSubject
	RiskAssessments                     []*r5.RiskAssessment
	Schedules                           []*r5.Schedule
	SearchParameters                    []*r5.SearchParameter
	ServiceRequests                     []*r5.ServiceRequest
	Slots                               []*r5.Slot
	Specimens                           []*r5.Specimen
	SpecimenDefinitions                 []*r5.SpecimenDefinition
	StructureDefinitions                []*r5.StructureDefinition
	StructureMaps                       []*r5.StructureMap
	Subscriptions                       []*r5.Subscription
	SubscriptionStatuss                 []*r5.SubscriptionStatus
	SubscriptionTopics                  []*r5.SubscriptionTopic
	Substances                          []*r5.Substance
	SubstanceDefinitions                []*r5.SubstanceDefinition
	SubstanceNucleicAcids               []*r5.SubstanceNucleicAcid
	SubstancePolymers                   []*r5.SubstancePolymer
	SubstanceProteins                   []*r5.SubstanceProtein
	SubstanceReferenceInformations      []*r5.SubstanceReferenceInformation
	SubstanceSourceMaterials            []*r5.SubstanceSourceMaterial
	SupplyDeliverys                     []*r5.SupplyDelivery
	SupplyRequests                      []*r5.SupplyRequest
	Tasks                               []*r5.Task
	TerminologyCapabilitiess            []*r5.TerminologyCapabilities
	TestPlans                           []*r5.TestPlan
	TestReports                         []*r5.TestReport
	TestScripts                         []*r5.TestScript
	Transports                          []*r5.Transport
	ValueSets                           []*r5.ValueSet
	VerificationResults                 []*r5.VerificationResult
	VisionPrescriptions                 []*r5.VisionPrescription
}

func BundleToGroup(bundle *r5.Bundle) (*ResourceGroup, error) {
	grp := ResourceGroup{}
	for _, e := range bundle.Entry {
		switch res := e.Resource.(type) {
		case *r5.Account:
			grp.Accounts = append(grp.Accounts, res)
		case *r5.ActivityDefinition:
			grp.ActivityDefinitions = append(grp.ActivityDefinitions, res)
		case *r5.ActorDefinition:
			grp.ActorDefinitions = append(grp.ActorDefinitions, res)
		case *r5.AdministrableProductDefinition:
			grp.AdministrableProductDefinitions = append(grp.AdministrableProductDefinitions, res)
		case *r5.AdverseEvent:
			grp.AdverseEvents = append(grp.AdverseEvents, res)
		case *r5.AllergyIntolerance:
			grp.AllergyIntolerances = append(grp.AllergyIntolerances, res)
		case *r5.Appointment:
			grp.Appointments = append(grp.Appointments, res)
		case *r5.AppointmentResponse:
			grp.AppointmentResponses = append(grp.AppointmentResponses, res)
		case *r5.ArtifactAssessment:
			grp.ArtifactAssessments = append(grp.ArtifactAssessments, res)
		case *r5.AuditEvent:
			grp.AuditEvents = append(grp.AuditEvents, res)
		case *r5.Basic:
			grp.Basics = append(grp.Basics, res)
		case *r5.Binary:
			grp.Binarys = append(grp.Binarys, res)
		case *r5.BiologicallyDerivedProduct:
			grp.BiologicallyDerivedProducts = append(grp.BiologicallyDerivedProducts, res)
		case *r5.BiologicallyDerivedProductDispense:
			grp.BiologicallyDerivedProductDispenses = append(grp.BiologicallyDerivedProductDispenses, res)
		case *r5.BodyStructure:
			grp.BodyStructures = append(grp.BodyStructures, res)
		case *r5.CanonicalResource:
			grp.CanonicalResources = append(grp.CanonicalResources, res)
		case *r5.CapabilityStatement:
			grp.CapabilityStatements = append(grp.CapabilityStatements, res)
		case *r5.CarePlan:
			grp.CarePlans = append(grp.CarePlans, res)
		case *r5.CareTeam:
			grp.CareTeams = append(grp.CareTeams, res)
		case *r5.ChargeItem:
			grp.ChargeItems = append(grp.ChargeItems, res)
		case *r5.ChargeItemDefinition:
			grp.ChargeItemDefinitions = append(grp.ChargeItemDefinitions, res)
		case *r5.Citation:
			grp.Citations = append(grp.Citations, res)
		case *r5.Claim:
			grp.Claims = append(grp.Claims, res)
		case *r5.ClaimResponse:
			grp.ClaimResponses = append(grp.ClaimResponses, res)
		case *r5.ClinicalImpression:
			grp.ClinicalImpressions = append(grp.ClinicalImpressions, res)
		case *r5.ClinicalUseDefinition:
			grp.ClinicalUseDefinitions = append(grp.ClinicalUseDefinitions, res)
		case *r5.CodeSystem:
			grp.CodeSystems = append(grp.CodeSystems, res)
		case *r5.Communication:
			grp.Communications = append(grp.Communications, res)
		case *r5.CommunicationRequest:
			grp.CommunicationRequests = append(grp.CommunicationRequests, res)
		case *r5.CompartmentDefinition:
			grp.CompartmentDefinitions = append(grp.CompartmentDefinitions, res)
		case *r5.Composition:
			grp.Compositions = append(grp.Compositions, res)
		case *r5.ConceptMap:
			grp.ConceptMaps = append(grp.ConceptMaps, res)
		case *r5.Condition:
			grp.Conditions = append(grp.Conditions, res)
		case *r5.ConditionDefinition:
			grp.ConditionDefinitions = append(grp.ConditionDefinitions, res)
		case *r5.Consent:
			grp.Consents = append(grp.Consents, res)
		case *r5.Contract:
			grp.Contracts = append(grp.Contracts, res)
		case *r5.Coverage:
			grp.Coverages = append(grp.Coverages, res)
		case *r5.CoverageEligibilityRequest:
			grp.CoverageEligibilityRequests = append(grp.CoverageEligibilityRequests, res)
		case *r5.CoverageEligibilityResponse:
			grp.CoverageEligibilityResponses = append(grp.CoverageEligibilityResponses, res)
		case *r5.DetectedIssue:
			grp.DetectedIssues = append(grp.DetectedIssues, res)
		case *r5.Device:
			grp.Devices = append(grp.Devices, res)
		case *r5.DeviceAssociation:
			grp.DeviceAssociations = append(grp.DeviceAssociations, res)
		case *r5.DeviceDefinition:
			grp.DeviceDefinitions = append(grp.DeviceDefinitions, res)
		case *r5.DeviceDispense:
			grp.DeviceDispenses = append(grp.DeviceDispenses, res)
		case *r5.DeviceMetric:
			grp.DeviceMetrics = append(grp.DeviceMetrics, res)
		case *r5.DeviceRequest:
			grp.DeviceRequests = append(grp.DeviceRequests, res)
		case *r5.DeviceUsage:
			grp.DeviceUsages = append(grp.DeviceUsages, res)
		case *r5.DiagnosticReport:
			grp.DiagnosticReports = append(grp.DiagnosticReports, res)
		case *r5.DocumentReference:
			grp.DocumentReferences = append(grp.DocumentReferences, res)
		case *r5.Encounter:
			grp.Encounters = append(grp.Encounters, res)
		case *r5.EncounterHistory:
			grp.EncounterHistorys = append(grp.EncounterHistorys, res)
		case *r5.Endpoint:
			grp.Endpoints = append(grp.Endpoints, res)
		case *r5.EnrollmentRequest:
			grp.EnrollmentRequests = append(grp.EnrollmentRequests, res)
		case *r5.EnrollmentResponse:
			grp.EnrollmentResponses = append(grp.EnrollmentResponses, res)
		case *r5.EpisodeOfCare:
			grp.EpisodeOfCares = append(grp.EpisodeOfCares, res)
		case *r5.EventDefinition:
			grp.EventDefinitions = append(grp.EventDefinitions, res)
		case *r5.Evidence:
			grp.Evidences = append(grp.Evidences, res)
		case *r5.EvidenceReport:
			grp.EvidenceReports = append(grp.EvidenceReports, res)
		case *r5.EvidenceVariable:
			grp.EvidenceVariables = append(grp.EvidenceVariables, res)
		case *r5.ExampleScenario:
			grp.ExampleScenarios = append(grp.ExampleScenarios, res)
		case *r5.ExplanationOfBenefit:
			grp.ExplanationOfBenefits = append(grp.ExplanationOfBenefits, res)
		case *r5.FamilyMemberHistory:
			grp.FamilyMemberHistorys = append(grp.FamilyMemberHistorys, res)
		case *r5.Flag:
			grp.Flags = append(grp.Flags, res)
		case *r5.FormularyItem:
			grp.FormularyItems = append(grp.FormularyItems, res)
		case *r5.GenomicStudy:
			grp.GenomicStudys = append(grp.GenomicStudys, res)
		case *r5.Goal:
			grp.Goals = append(grp.Goals, res)
		case *r5.GraphDefinition:
			grp.GraphDefinitions = append(grp.GraphDefinitions, res)
		case *r5.Group:
			grp.Groups = append(grp.Groups, res)
		case *r5.GuidanceResponse:
			grp.GuidanceResponses = append(grp.GuidanceResponses, res)
		case *r5.HealthcareService:
			grp.HealthcareServices = append(grp.HealthcareServices, res)
		case *r5.ImagingSelection:
			grp.ImagingSelections = append(grp.ImagingSelections, res)
		case *r5.ImagingStudy:
			grp.ImagingStudys = append(grp.ImagingStudys, res)
		case *r5.Immunization:
			grp.Immunizations = append(grp.Immunizations, res)
		case *r5.ImmunizationEvaluation:
			grp.ImmunizationEvaluations = append(grp.ImmunizationEvaluations, res)
		case *r5.ImmunizationRecommendation:
			grp.ImmunizationRecommendations = append(grp.ImmunizationRecommendations, res)
		case *r5.ImplementationGuide:
			grp.ImplementationGuides = append(grp.ImplementationGuides, res)
		case *r5.Ingredient:
			grp.Ingredients = append(grp.Ingredients, res)
		case *r5.InsurancePlan:
			grp.InsurancePlans = append(grp.InsurancePlans, res)
		case *r5.InventoryItem:
			grp.InventoryItems = append(grp.InventoryItems, res)
		case *r5.InventoryReport:
			grp.InventoryReports = append(grp.InventoryReports, res)
		case *r5.Invoice:
			grp.Invoices = append(grp.Invoices, res)
		case *r5.Library:
			grp.Librarys = append(grp.Librarys, res)
		case *r5.Linkage:
			grp.Linkages = append(grp.Linkages, res)
		case *r5.List:
			grp.Lists = append(grp.Lists, res)
		case *r5.Location:
			grp.Locations = append(grp.Locations, res)
		case *r5.ManufacturedItemDefinition:
			grp.ManufacturedItemDefinitions = append(grp.ManufacturedItemDefinitions, res)
		case *r5.Measure:
			grp.Measures = append(grp.Measures, res)
		case *r5.MeasureReport:
			grp.MeasureReports = append(grp.MeasureReports, res)
		case *r5.Medication:
			grp.Medications = append(grp.Medications, res)
		case *r5.MedicationAdministration:
			grp.MedicationAdministrations = append(grp.MedicationAdministrations, res)
		case *r5.MedicationDispense:
			grp.MedicationDispenses = append(grp.MedicationDispenses, res)
		case *r5.MedicationKnowledge:
			grp.MedicationKnowledges = append(grp.MedicationKnowledges, res)
		case *r5.MedicationRequest:
			grp.MedicationRequests = append(grp.MedicationRequests, res)
		case *r5.MedicationStatement:
			grp.MedicationStatements = append(grp.MedicationStatements, res)
		case *r5.MedicinalProductDefinition:
			grp.MedicinalProductDefinitions = append(grp.MedicinalProductDefinitions, res)
		case *r5.MessageDefinition:
			grp.MessageDefinitions = append(grp.MessageDefinitions, res)
		case *r5.MessageHeader:
			grp.MessageHeaders = append(grp.MessageHeaders, res)
		case *r5.MolecularSequence:
			grp.MolecularSequences = append(grp.MolecularSequences, res)
		case *r5.NamingSystem:
			grp.NamingSystems = append(grp.NamingSystems, res)
		case *r5.NutritionIntake:
			grp.NutritionIntakes = append(grp.NutritionIntakes, res)
		case *r5.NutritionOrder:
			grp.NutritionOrders = append(grp.NutritionOrders, res)
		case *r5.NutritionProduct:
			grp.NutritionProducts = append(grp.NutritionProducts, res)
		case *r5.Observation:
			grp.Observations = append(grp.Observations, res)
		case *r5.ObservationDefinition:
			grp.ObservationDefinitions = append(grp.ObservationDefinitions, res)
		case *r5.OperationDefinition:
			grp.OperationDefinitions = append(grp.OperationDefinitions, res)
		case *r5.OperationOutcome:
			grp.OperationOutcomes = append(grp.OperationOutcomes, res)
		case *r5.Organization:
			grp.Organizations = append(grp.Organizations, res)
		case *r5.OrganizationAffiliation:
			grp.OrganizationAffiliations = append(grp.OrganizationAffiliations, res)
		case *r5.PackagedProductDefinition:
			grp.PackagedProductDefinitions = append(grp.PackagedProductDefinitions, res)
		case *r5.Patient:
			grp.Patients = append(grp.Patients, res)
		case *r5.PaymentNotice:
			grp.PaymentNotices = append(grp.PaymentNotices, res)
		case *r5.PaymentReconciliation:
			grp.PaymentReconciliations = append(grp.PaymentReconciliations, res)
		case *r5.Permission:
			grp.Permissions = append(grp.Permissions, res)
		case *r5.Person:
			grp.Persons = append(grp.Persons, res)
		case *r5.PlanDefinition:
			grp.PlanDefinitions = append(grp.PlanDefinitions, res)
		case *r5.Practitioner:
			grp.Practitioners = append(grp.Practitioners, res)
		case *r5.PractitionerRole:
			grp.PractitionerRoles = append(grp.PractitionerRoles, res)
		case *r5.Procedure:
			grp.Procedures = append(grp.Procedures, res)
		case *r5.Provenance:
			grp.Provenances = append(grp.Provenances, res)
		case *r5.Questionnaire:
			grp.Questionnaires = append(grp.Questionnaires, res)
		case *r5.QuestionnaireResponse:
			grp.QuestionnaireResponses = append(grp.QuestionnaireResponses, res)
		case *r5.RegulatedAuthorization:
			grp.RegulatedAuthorizations = append(grp.RegulatedAuthorizations, res)
		case *r5.RelatedPerson:
			grp.RelatedPersons = append(grp.RelatedPersons, res)
		case *r5.RequestOrchestration:
			grp.RequestOrchestrations = append(grp.RequestOrchestrations, res)
		case *r5.Requirements:
			grp.Requirementss = append(grp.Requirementss, res)
		case *r5.ResearchStudy:
			grp.ResearchStudys = append(grp.ResearchStudys, res)
		case *r5.ResearchSubject:
			grp.ResearchSubjects = append(grp.ResearchSubjects, res)
		case *r5.RiskAssessment:
			grp.RiskAssessments = append(grp.RiskAssessments, res)
		case *r5.Schedule:
			grp.Schedules = append(grp.Schedules, res)
		case *r5.SearchParameter:
			grp.SearchParameters = append(grp.SearchParameters, res)
		case *r5.ServiceRequest:
			grp.ServiceRequests = append(grp.ServiceRequests, res)
		case *r5.Slot:
			grp.Slots = append(grp.Slots, res)
		case *r5.Specimen:
			grp.Specimens = append(grp.Specimens, res)
		case *r5.SpecimenDefinition:
			grp.SpecimenDefinitions = append(grp.SpecimenDefinitions, res)
		case *r5.StructureDefinition:
			grp.StructureDefinitions = append(grp.StructureDefinitions, res)
		case *r5.StructureMap:
			grp.StructureMaps = append(grp.StructureMaps, res)
		case *r5.Subscription:
			grp.Subscriptions = append(grp.Subscriptions, res)
		case *r5.SubscriptionStatus:
			grp.SubscriptionStatuss = append(grp.SubscriptionStatuss, res)
		case *r5.SubscriptionTopic:
			grp.SubscriptionTopics = append(grp.SubscriptionTopics, res)
		case *r5.Substance:
			grp.Substances = append(grp.Substances, res)
		case *r5.SubstanceDefinition:
			grp.SubstanceDefinitions = append(grp.SubstanceDefinitions, res)
		case *r5.SubstanceNucleicAcid:
			grp.SubstanceNucleicAcids = append(grp.SubstanceNucleicAcids, res)
		case *r5.SubstancePolymer:
			grp.SubstancePolymers = append(grp.SubstancePolymers, res)
		case *r5.SubstanceProtein:
			grp.SubstanceProteins = append(grp.SubstanceProteins, res)
		case *r5.SubstanceReferenceInformation:
			grp.SubstanceReferenceInformations = append(grp.SubstanceReferenceInformations, res)
		case *r5.SubstanceSourceMaterial:
			grp.SubstanceSourceMaterials = append(grp.SubstanceSourceMaterials, res)
		case *r5.SupplyDelivery:
			grp.SupplyDeliverys = append(grp.SupplyDeliverys, res)
		case *r5.SupplyRequest:
			grp.SupplyRequests = append(grp.SupplyRequests, res)
		case *r5.Task:
			grp.Tasks = append(grp.Tasks, res)
		case *r5.TerminologyCapabilities:
			grp.TerminologyCapabilitiess = append(grp.TerminologyCapabilitiess, res)
		case *r5.TestPlan:
			grp.TestPlans = append(grp.TestPlans, res)
		case *r5.TestReport:
			grp.TestReports = append(grp.TestReports, res)
		case *r5.TestScript:
			grp.TestScripts = append(grp.TestScripts, res)
		case *r5.Transport:
			grp.Transports = append(grp.Transports, res)
		case *r5.ValueSet:
			grp.ValueSets = append(grp.ValueSets, res)
		case *r5.VerificationResult:
			grp.VerificationResults = append(grp.VerificationResults, res)
		case *r5.VisionPrescription:
			grp.VisionPrescriptions = append(grp.VisionPrescriptions, res)
		default:
			return nil, errors.New("bundle entry not a domain resource, could not put in resourcegroup")
		}
	}
	return &grp, nil
}
