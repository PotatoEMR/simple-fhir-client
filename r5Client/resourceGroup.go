package r5Client

import r5 "github.com/PotatoEMR/simple-fhir-client/r5"
import "errors"

type ResourceGroup struct {
	Account_list                            []*r5.Account
	ActivityDefinition_list                 []*r5.ActivityDefinition
	ActorDefinition_list                    []*r5.ActorDefinition
	AdministrableProductDefinition_list     []*r5.AdministrableProductDefinition
	AdverseEvent_list                       []*r5.AdverseEvent
	AllergyIntolerance_list                 []*r5.AllergyIntolerance
	Appointment_list                        []*r5.Appointment
	AppointmentResponse_list                []*r5.AppointmentResponse
	ArtifactAssessment_list                 []*r5.ArtifactAssessment
	AuditEvent_list                         []*r5.AuditEvent
	Basic_list                              []*r5.Basic
	Binary_list                             []*r5.Binary
	BiologicallyDerivedProduct_list         []*r5.BiologicallyDerivedProduct
	BiologicallyDerivedProductDispense_list []*r5.BiologicallyDerivedProductDispense
	BodyStructure_list                      []*r5.BodyStructure
	CanonicalResource_list                  []*r5.CanonicalResource
	CapabilityStatement_list                []*r5.CapabilityStatement
	CarePlan_list                           []*r5.CarePlan
	CareTeam_list                           []*r5.CareTeam
	ChargeItem_list                         []*r5.ChargeItem
	ChargeItemDefinition_list               []*r5.ChargeItemDefinition
	Citation_list                           []*r5.Citation
	Claim_list                              []*r5.Claim
	ClaimResponse_list                      []*r5.ClaimResponse
	ClinicalImpression_list                 []*r5.ClinicalImpression
	ClinicalUseDefinition_list              []*r5.ClinicalUseDefinition
	CodeSystem_list                         []*r5.CodeSystem
	Communication_list                      []*r5.Communication
	CommunicationRequest_list               []*r5.CommunicationRequest
	CompartmentDefinition_list              []*r5.CompartmentDefinition
	Composition_list                        []*r5.Composition
	ConceptMap_list                         []*r5.ConceptMap
	Condition_list                          []*r5.Condition
	ConditionDefinition_list                []*r5.ConditionDefinition
	Consent_list                            []*r5.Consent
	Contract_list                           []*r5.Contract
	Coverage_list                           []*r5.Coverage
	CoverageEligibilityRequest_list         []*r5.CoverageEligibilityRequest
	CoverageEligibilityResponse_list        []*r5.CoverageEligibilityResponse
	DetectedIssue_list                      []*r5.DetectedIssue
	Device_list                             []*r5.Device
	DeviceAssociation_list                  []*r5.DeviceAssociation
	DeviceDefinition_list                   []*r5.DeviceDefinition
	DeviceDispense_list                     []*r5.DeviceDispense
	DeviceMetric_list                       []*r5.DeviceMetric
	DeviceRequest_list                      []*r5.DeviceRequest
	DeviceUsage_list                        []*r5.DeviceUsage
	DiagnosticReport_list                   []*r5.DiagnosticReport
	DocumentReference_list                  []*r5.DocumentReference
	Encounter_list                          []*r5.Encounter
	EncounterHistory_list                   []*r5.EncounterHistory
	Endpoint_list                           []*r5.Endpoint
	EnrollmentRequest_list                  []*r5.EnrollmentRequest
	EnrollmentResponse_list                 []*r5.EnrollmentResponse
	EpisodeOfCare_list                      []*r5.EpisodeOfCare
	EventDefinition_list                    []*r5.EventDefinition
	Evidence_list                           []*r5.Evidence
	EvidenceReport_list                     []*r5.EvidenceReport
	EvidenceVariable_list                   []*r5.EvidenceVariable
	ExampleScenario_list                    []*r5.ExampleScenario
	ExplanationOfBenefit_list               []*r5.ExplanationOfBenefit
	FamilyMemberHistory_list                []*r5.FamilyMemberHistory
	Flag_list                               []*r5.Flag
	FormularyItem_list                      []*r5.FormularyItem
	GenomicStudy_list                       []*r5.GenomicStudy
	Goal_list                               []*r5.Goal
	GraphDefinition_list                    []*r5.GraphDefinition
	Group_list                              []*r5.Group
	GuidanceResponse_list                   []*r5.GuidanceResponse
	HealthcareService_list                  []*r5.HealthcareService
	ImagingSelection_list                   []*r5.ImagingSelection
	ImagingStudy_list                       []*r5.ImagingStudy
	Immunization_list                       []*r5.Immunization
	ImmunizationEvaluation_list             []*r5.ImmunizationEvaluation
	ImmunizationRecommendation_list         []*r5.ImmunizationRecommendation
	ImplementationGuide_list                []*r5.ImplementationGuide
	Ingredient_list                         []*r5.Ingredient
	InsurancePlan_list                      []*r5.InsurancePlan
	InventoryItem_list                      []*r5.InventoryItem
	InventoryReport_list                    []*r5.InventoryReport
	Invoice_list                            []*r5.Invoice
	Library_list                            []*r5.Library
	Linkage_list                            []*r5.Linkage
	List_list                               []*r5.List
	Location_list                           []*r5.Location
	ManufacturedItemDefinition_list         []*r5.ManufacturedItemDefinition
	Measure_list                            []*r5.Measure
	MeasureReport_list                      []*r5.MeasureReport
	Medication_list                         []*r5.Medication
	MedicationAdministration_list           []*r5.MedicationAdministration
	MedicationDispense_list                 []*r5.MedicationDispense
	MedicationKnowledge_list                []*r5.MedicationKnowledge
	MedicationRequest_list                  []*r5.MedicationRequest
	MedicationStatement_list                []*r5.MedicationStatement
	MedicinalProductDefinition_list         []*r5.MedicinalProductDefinition
	MessageDefinition_list                  []*r5.MessageDefinition
	MessageHeader_list                      []*r5.MessageHeader
	MolecularSequence_list                  []*r5.MolecularSequence
	NamingSystem_list                       []*r5.NamingSystem
	NutritionIntake_list                    []*r5.NutritionIntake
	NutritionOrder_list                     []*r5.NutritionOrder
	NutritionProduct_list                   []*r5.NutritionProduct
	Observation_list                        []*r5.Observation
	ObservationDefinition_list              []*r5.ObservationDefinition
	OperationDefinition_list                []*r5.OperationDefinition
	OperationOutcome_list                   []*r5.OperationOutcome
	Organization_list                       []*r5.Organization
	OrganizationAffiliation_list            []*r5.OrganizationAffiliation
	PackagedProductDefinition_list          []*r5.PackagedProductDefinition
	Patient_list                            []*r5.Patient
	PaymentNotice_list                      []*r5.PaymentNotice
	PaymentReconciliation_list              []*r5.PaymentReconciliation
	Permission_list                         []*r5.Permission
	Person_list                             []*r5.Person
	PlanDefinition_list                     []*r5.PlanDefinition
	Practitioner_list                       []*r5.Practitioner
	PractitionerRole_list                   []*r5.PractitionerRole
	Procedure_list                          []*r5.Procedure
	Provenance_list                         []*r5.Provenance
	Questionnaire_list                      []*r5.Questionnaire
	QuestionnaireResponse_list              []*r5.QuestionnaireResponse
	RegulatedAuthorization_list             []*r5.RegulatedAuthorization
	RelatedPerson_list                      []*r5.RelatedPerson
	RequestOrchestration_list               []*r5.RequestOrchestration
	Requirements_list                       []*r5.Requirements
	ResearchStudy_list                      []*r5.ResearchStudy
	ResearchSubject_list                    []*r5.ResearchSubject
	RiskAssessment_list                     []*r5.RiskAssessment
	Schedule_list                           []*r5.Schedule
	SearchParameter_list                    []*r5.SearchParameter
	ServiceRequest_list                     []*r5.ServiceRequest
	Slot_list                               []*r5.Slot
	Specimen_list                           []*r5.Specimen
	SpecimenDefinition_list                 []*r5.SpecimenDefinition
	StructureDefinition_list                []*r5.StructureDefinition
	StructureMap_list                       []*r5.StructureMap
	Subscription_list                       []*r5.Subscription
	SubscriptionStatus_list                 []*r5.SubscriptionStatus
	SubscriptionTopic_list                  []*r5.SubscriptionTopic
	Substance_list                          []*r5.Substance
	SubstanceDefinition_list                []*r5.SubstanceDefinition
	SubstanceNucleicAcid_list               []*r5.SubstanceNucleicAcid
	SubstancePolymer_list                   []*r5.SubstancePolymer
	SubstanceProtein_list                   []*r5.SubstanceProtein
	SubstanceReferenceInformation_list      []*r5.SubstanceReferenceInformation
	SubstanceSourceMaterial_list            []*r5.SubstanceSourceMaterial
	SupplyDelivery_list                     []*r5.SupplyDelivery
	SupplyRequest_list                      []*r5.SupplyRequest
	Task_list                               []*r5.Task
	TerminologyCapabilities_list            []*r5.TerminologyCapabilities
	TestPlan_list                           []*r5.TestPlan
	TestReport_list                         []*r5.TestReport
	TestScript_list                         []*r5.TestScript
	Transport_list                          []*r5.Transport
	ValueSet_list                           []*r5.ValueSet
	VerificationResult_list                 []*r5.VerificationResult
	VisionPrescription_list                 []*r5.VisionPrescription
}

func BundleToGroup(bundle *r4.Bundle) (*ResourceGroup, error) {
	grp := ResourceGroup{}
	for _, e := range bundle.Entry {
		switch res := e.Resource.(type) {
		case *r5.Account:
			grp.Account_list = append(grp.Account_list, res)
		case *r5.ActivityDefinition:
			grp.ActivityDefinition_list = append(grp.ActivityDefinition_list, res)
		case *r5.ActorDefinition:
			grp.ActorDefinition_list = append(grp.ActorDefinition_list, res)
		case *r5.AdministrableProductDefinition:
			grp.AdministrableProductDefinition_list = append(grp.AdministrableProductDefinition_list, res)
		case *r5.AdverseEvent:
			grp.AdverseEvent_list = append(grp.AdverseEvent_list, res)
		case *r5.AllergyIntolerance:
			grp.AllergyIntolerance_list = append(grp.AllergyIntolerance_list, res)
		case *r5.Appointment:
			grp.Appointment_list = append(grp.Appointment_list, res)
		case *r5.AppointmentResponse:
			grp.AppointmentResponse_list = append(grp.AppointmentResponse_list, res)
		case *r5.ArtifactAssessment:
			grp.ArtifactAssessment_list = append(grp.ArtifactAssessment_list, res)
		case *r5.AuditEvent:
			grp.AuditEvent_list = append(grp.AuditEvent_list, res)
		case *r5.Basic:
			grp.Basic_list = append(grp.Basic_list, res)
		case *r5.Binary:
			grp.Binary_list = append(grp.Binary_list, res)
		case *r5.BiologicallyDerivedProduct:
			grp.BiologicallyDerivedProduct_list = append(grp.BiologicallyDerivedProduct_list, res)
		case *r5.BiologicallyDerivedProductDispense:
			grp.BiologicallyDerivedProductDispense_list = append(grp.BiologicallyDerivedProductDispense_list, res)
		case *r5.BodyStructure:
			grp.BodyStructure_list = append(grp.BodyStructure_list, res)
		case *r5.CanonicalResource:
			grp.CanonicalResource_list = append(grp.CanonicalResource_list, res)
		case *r5.CapabilityStatement:
			grp.CapabilityStatement_list = append(grp.CapabilityStatement_list, res)
		case *r5.CarePlan:
			grp.CarePlan_list = append(grp.CarePlan_list, res)
		case *r5.CareTeam:
			grp.CareTeam_list = append(grp.CareTeam_list, res)
		case *r5.ChargeItem:
			grp.ChargeItem_list = append(grp.ChargeItem_list, res)
		case *r5.ChargeItemDefinition:
			grp.ChargeItemDefinition_list = append(grp.ChargeItemDefinition_list, res)
		case *r5.Citation:
			grp.Citation_list = append(grp.Citation_list, res)
		case *r5.Claim:
			grp.Claim_list = append(grp.Claim_list, res)
		case *r5.ClaimResponse:
			grp.ClaimResponse_list = append(grp.ClaimResponse_list, res)
		case *r5.ClinicalImpression:
			grp.ClinicalImpression_list = append(grp.ClinicalImpression_list, res)
		case *r5.ClinicalUseDefinition:
			grp.ClinicalUseDefinition_list = append(grp.ClinicalUseDefinition_list, res)
		case *r5.CodeSystem:
			grp.CodeSystem_list = append(grp.CodeSystem_list, res)
		case *r5.Communication:
			grp.Communication_list = append(grp.Communication_list, res)
		case *r5.CommunicationRequest:
			grp.CommunicationRequest_list = append(grp.CommunicationRequest_list, res)
		case *r5.CompartmentDefinition:
			grp.CompartmentDefinition_list = append(grp.CompartmentDefinition_list, res)
		case *r5.Composition:
			grp.Composition_list = append(grp.Composition_list, res)
		case *r5.ConceptMap:
			grp.ConceptMap_list = append(grp.ConceptMap_list, res)
		case *r5.Condition:
			grp.Condition_list = append(grp.Condition_list, res)
		case *r5.ConditionDefinition:
			grp.ConditionDefinition_list = append(grp.ConditionDefinition_list, res)
		case *r5.Consent:
			grp.Consent_list = append(grp.Consent_list, res)
		case *r5.Contract:
			grp.Contract_list = append(grp.Contract_list, res)
		case *r5.Coverage:
			grp.Coverage_list = append(grp.Coverage_list, res)
		case *r5.CoverageEligibilityRequest:
			grp.CoverageEligibilityRequest_list = append(grp.CoverageEligibilityRequest_list, res)
		case *r5.CoverageEligibilityResponse:
			grp.CoverageEligibilityResponse_list = append(grp.CoverageEligibilityResponse_list, res)
		case *r5.DetectedIssue:
			grp.DetectedIssue_list = append(grp.DetectedIssue_list, res)
		case *r5.Device:
			grp.Device_list = append(grp.Device_list, res)
		case *r5.DeviceAssociation:
			grp.DeviceAssociation_list = append(grp.DeviceAssociation_list, res)
		case *r5.DeviceDefinition:
			grp.DeviceDefinition_list = append(grp.DeviceDefinition_list, res)
		case *r5.DeviceDispense:
			grp.DeviceDispense_list = append(grp.DeviceDispense_list, res)
		case *r5.DeviceMetric:
			grp.DeviceMetric_list = append(grp.DeviceMetric_list, res)
		case *r5.DeviceRequest:
			grp.DeviceRequest_list = append(grp.DeviceRequest_list, res)
		case *r5.DeviceUsage:
			grp.DeviceUsage_list = append(grp.DeviceUsage_list, res)
		case *r5.DiagnosticReport:
			grp.DiagnosticReport_list = append(grp.DiagnosticReport_list, res)
		case *r5.DocumentReference:
			grp.DocumentReference_list = append(grp.DocumentReference_list, res)
		case *r5.Encounter:
			grp.Encounter_list = append(grp.Encounter_list, res)
		case *r5.EncounterHistory:
			grp.EncounterHistory_list = append(grp.EncounterHistory_list, res)
		case *r5.Endpoint:
			grp.Endpoint_list = append(grp.Endpoint_list, res)
		case *r5.EnrollmentRequest:
			grp.EnrollmentRequest_list = append(grp.EnrollmentRequest_list, res)
		case *r5.EnrollmentResponse:
			grp.EnrollmentResponse_list = append(grp.EnrollmentResponse_list, res)
		case *r5.EpisodeOfCare:
			grp.EpisodeOfCare_list = append(grp.EpisodeOfCare_list, res)
		case *r5.EventDefinition:
			grp.EventDefinition_list = append(grp.EventDefinition_list, res)
		case *r5.Evidence:
			grp.Evidence_list = append(grp.Evidence_list, res)
		case *r5.EvidenceReport:
			grp.EvidenceReport_list = append(grp.EvidenceReport_list, res)
		case *r5.EvidenceVariable:
			grp.EvidenceVariable_list = append(grp.EvidenceVariable_list, res)
		case *r5.ExampleScenario:
			grp.ExampleScenario_list = append(grp.ExampleScenario_list, res)
		case *r5.ExplanationOfBenefit:
			grp.ExplanationOfBenefit_list = append(grp.ExplanationOfBenefit_list, res)
		case *r5.FamilyMemberHistory:
			grp.FamilyMemberHistory_list = append(grp.FamilyMemberHistory_list, res)
		case *r5.Flag:
			grp.Flag_list = append(grp.Flag_list, res)
		case *r5.FormularyItem:
			grp.FormularyItem_list = append(grp.FormularyItem_list, res)
		case *r5.GenomicStudy:
			grp.GenomicStudy_list = append(grp.GenomicStudy_list, res)
		case *r5.Goal:
			grp.Goal_list = append(grp.Goal_list, res)
		case *r5.GraphDefinition:
			grp.GraphDefinition_list = append(grp.GraphDefinition_list, res)
		case *r5.Group:
			grp.Group_list = append(grp.Group_list, res)
		case *r5.GuidanceResponse:
			grp.GuidanceResponse_list = append(grp.GuidanceResponse_list, res)
		case *r5.HealthcareService:
			grp.HealthcareService_list = append(grp.HealthcareService_list, res)
		case *r5.ImagingSelection:
			grp.ImagingSelection_list = append(grp.ImagingSelection_list, res)
		case *r5.ImagingStudy:
			grp.ImagingStudy_list = append(grp.ImagingStudy_list, res)
		case *r5.Immunization:
			grp.Immunization_list = append(grp.Immunization_list, res)
		case *r5.ImmunizationEvaluation:
			grp.ImmunizationEvaluation_list = append(grp.ImmunizationEvaluation_list, res)
		case *r5.ImmunizationRecommendation:
			grp.ImmunizationRecommendation_list = append(grp.ImmunizationRecommendation_list, res)
		case *r5.ImplementationGuide:
			grp.ImplementationGuide_list = append(grp.ImplementationGuide_list, res)
		case *r5.Ingredient:
			grp.Ingredient_list = append(grp.Ingredient_list, res)
		case *r5.InsurancePlan:
			grp.InsurancePlan_list = append(grp.InsurancePlan_list, res)
		case *r5.InventoryItem:
			grp.InventoryItem_list = append(grp.InventoryItem_list, res)
		case *r5.InventoryReport:
			grp.InventoryReport_list = append(grp.InventoryReport_list, res)
		case *r5.Invoice:
			grp.Invoice_list = append(grp.Invoice_list, res)
		case *r5.Library:
			grp.Library_list = append(grp.Library_list, res)
		case *r5.Linkage:
			grp.Linkage_list = append(grp.Linkage_list, res)
		case *r5.List:
			grp.List_list = append(grp.List_list, res)
		case *r5.Location:
			grp.Location_list = append(grp.Location_list, res)
		case *r5.ManufacturedItemDefinition:
			grp.ManufacturedItemDefinition_list = append(grp.ManufacturedItemDefinition_list, res)
		case *r5.Measure:
			grp.Measure_list = append(grp.Measure_list, res)
		case *r5.MeasureReport:
			grp.MeasureReport_list = append(grp.MeasureReport_list, res)
		case *r5.Medication:
			grp.Medication_list = append(grp.Medication_list, res)
		case *r5.MedicationAdministration:
			grp.MedicationAdministration_list = append(grp.MedicationAdministration_list, res)
		case *r5.MedicationDispense:
			grp.MedicationDispense_list = append(grp.MedicationDispense_list, res)
		case *r5.MedicationKnowledge:
			grp.MedicationKnowledge_list = append(grp.MedicationKnowledge_list, res)
		case *r5.MedicationRequest:
			grp.MedicationRequest_list = append(grp.MedicationRequest_list, res)
		case *r5.MedicationStatement:
			grp.MedicationStatement_list = append(grp.MedicationStatement_list, res)
		case *r5.MedicinalProductDefinition:
			grp.MedicinalProductDefinition_list = append(grp.MedicinalProductDefinition_list, res)
		case *r5.MessageDefinition:
			grp.MessageDefinition_list = append(grp.MessageDefinition_list, res)
		case *r5.MessageHeader:
			grp.MessageHeader_list = append(grp.MessageHeader_list, res)
		case *r5.MolecularSequence:
			grp.MolecularSequence_list = append(grp.MolecularSequence_list, res)
		case *r5.NamingSystem:
			grp.NamingSystem_list = append(grp.NamingSystem_list, res)
		case *r5.NutritionIntake:
			grp.NutritionIntake_list = append(grp.NutritionIntake_list, res)
		case *r5.NutritionOrder:
			grp.NutritionOrder_list = append(grp.NutritionOrder_list, res)
		case *r5.NutritionProduct:
			grp.NutritionProduct_list = append(grp.NutritionProduct_list, res)
		case *r5.Observation:
			grp.Observation_list = append(grp.Observation_list, res)
		case *r5.ObservationDefinition:
			grp.ObservationDefinition_list = append(grp.ObservationDefinition_list, res)
		case *r5.OperationDefinition:
			grp.OperationDefinition_list = append(grp.OperationDefinition_list, res)
		case *r5.OperationOutcome:
			grp.OperationOutcome_list = append(grp.OperationOutcome_list, res)
		case *r5.Organization:
			grp.Organization_list = append(grp.Organization_list, res)
		case *r5.OrganizationAffiliation:
			grp.OrganizationAffiliation_list = append(grp.OrganizationAffiliation_list, res)
		case *r5.PackagedProductDefinition:
			grp.PackagedProductDefinition_list = append(grp.PackagedProductDefinition_list, res)
		case *r5.Patient:
			grp.Patient_list = append(grp.Patient_list, res)
		case *r5.PaymentNotice:
			grp.PaymentNotice_list = append(grp.PaymentNotice_list, res)
		case *r5.PaymentReconciliation:
			grp.PaymentReconciliation_list = append(grp.PaymentReconciliation_list, res)
		case *r5.Permission:
			grp.Permission_list = append(grp.Permission_list, res)
		case *r5.Person:
			grp.Person_list = append(grp.Person_list, res)
		case *r5.PlanDefinition:
			grp.PlanDefinition_list = append(grp.PlanDefinition_list, res)
		case *r5.Practitioner:
			grp.Practitioner_list = append(grp.Practitioner_list, res)
		case *r5.PractitionerRole:
			grp.PractitionerRole_list = append(grp.PractitionerRole_list, res)
		case *r5.Procedure:
			grp.Procedure_list = append(grp.Procedure_list, res)
		case *r5.Provenance:
			grp.Provenance_list = append(grp.Provenance_list, res)
		case *r5.Questionnaire:
			grp.Questionnaire_list = append(grp.Questionnaire_list, res)
		case *r5.QuestionnaireResponse:
			grp.QuestionnaireResponse_list = append(grp.QuestionnaireResponse_list, res)
		case *r5.RegulatedAuthorization:
			grp.RegulatedAuthorization_list = append(grp.RegulatedAuthorization_list, res)
		case *r5.RelatedPerson:
			grp.RelatedPerson_list = append(grp.RelatedPerson_list, res)
		case *r5.RequestOrchestration:
			grp.RequestOrchestration_list = append(grp.RequestOrchestration_list, res)
		case *r5.Requirements:
			grp.Requirements_list = append(grp.Requirements_list, res)
		case *r5.ResearchStudy:
			grp.ResearchStudy_list = append(grp.ResearchStudy_list, res)
		case *r5.ResearchSubject:
			grp.ResearchSubject_list = append(grp.ResearchSubject_list, res)
		case *r5.RiskAssessment:
			grp.RiskAssessment_list = append(grp.RiskAssessment_list, res)
		case *r5.Schedule:
			grp.Schedule_list = append(grp.Schedule_list, res)
		case *r5.SearchParameter:
			grp.SearchParameter_list = append(grp.SearchParameter_list, res)
		case *r5.ServiceRequest:
			grp.ServiceRequest_list = append(grp.ServiceRequest_list, res)
		case *r5.Slot:
			grp.Slot_list = append(grp.Slot_list, res)
		case *r5.Specimen:
			grp.Specimen_list = append(grp.Specimen_list, res)
		case *r5.SpecimenDefinition:
			grp.SpecimenDefinition_list = append(grp.SpecimenDefinition_list, res)
		case *r5.StructureDefinition:
			grp.StructureDefinition_list = append(grp.StructureDefinition_list, res)
		case *r5.StructureMap:
			grp.StructureMap_list = append(grp.StructureMap_list, res)
		case *r5.Subscription:
			grp.Subscription_list = append(grp.Subscription_list, res)
		case *r5.SubscriptionStatus:
			grp.SubscriptionStatus_list = append(grp.SubscriptionStatus_list, res)
		case *r5.SubscriptionTopic:
			grp.SubscriptionTopic_list = append(grp.SubscriptionTopic_list, res)
		case *r5.Substance:
			grp.Substance_list = append(grp.Substance_list, res)
		case *r5.SubstanceDefinition:
			grp.SubstanceDefinition_list = append(grp.SubstanceDefinition_list, res)
		case *r5.SubstanceNucleicAcid:
			grp.SubstanceNucleicAcid_list = append(grp.SubstanceNucleicAcid_list, res)
		case *r5.SubstancePolymer:
			grp.SubstancePolymer_list = append(grp.SubstancePolymer_list, res)
		case *r5.SubstanceProtein:
			grp.SubstanceProtein_list = append(grp.SubstanceProtein_list, res)
		case *r5.SubstanceReferenceInformation:
			grp.SubstanceReferenceInformation_list = append(grp.SubstanceReferenceInformation_list, res)
		case *r5.SubstanceSourceMaterial:
			grp.SubstanceSourceMaterial_list = append(grp.SubstanceSourceMaterial_list, res)
		case *r5.SupplyDelivery:
			grp.SupplyDelivery_list = append(grp.SupplyDelivery_list, res)
		case *r5.SupplyRequest:
			grp.SupplyRequest_list = append(grp.SupplyRequest_list, res)
		case *r5.Task:
			grp.Task_list = append(grp.Task_list, res)
		case *r5.TerminologyCapabilities:
			grp.TerminologyCapabilities_list = append(grp.TerminologyCapabilities_list, res)
		case *r5.TestPlan:
			grp.TestPlan_list = append(grp.TestPlan_list, res)
		case *r5.TestReport:
			grp.TestReport_list = append(grp.TestReport_list, res)
		case *r5.TestScript:
			grp.TestScript_list = append(grp.TestScript_list, res)
		case *r5.Transport:
			grp.Transport_list = append(grp.Transport_list, res)
		case *r5.ValueSet:
			grp.ValueSet_list = append(grp.ValueSet_list, res)
		case *r5.VerificationResult:
			grp.VerificationResult_list = append(grp.VerificationResult_list, res)
		case *r5.VisionPrescription:
			grp.VisionPrescription_list = append(grp.VisionPrescription_list, res)
		default:
			return nil, errors.New("bundle entry not a domain resource, could not put in resourcegroup")
		}
	}
	return &grp, nil
}
