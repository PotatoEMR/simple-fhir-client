package r4bClient

import r4b "github.com/PotatoEMR/simple-fhir-client/r4b"
import "errors"

type ResourceGroup struct {
	Account_list                        []*r4b.Account
	ActivityDefinition_list             []*r4b.ActivityDefinition
	AdministrableProductDefinition_list []*r4b.AdministrableProductDefinition
	AdverseEvent_list                   []*r4b.AdverseEvent
	AllergyIntolerance_list             []*r4b.AllergyIntolerance
	Appointment_list                    []*r4b.Appointment
	AppointmentResponse_list            []*r4b.AppointmentResponse
	AuditEvent_list                     []*r4b.AuditEvent
	Basic_list                          []*r4b.Basic
	Binary_list                         []*r4b.Binary
	BiologicallyDerivedProduct_list     []*r4b.BiologicallyDerivedProduct
	BodyStructure_list                  []*r4b.BodyStructure
	CapabilityStatement_list            []*r4b.CapabilityStatement
	CarePlan_list                       []*r4b.CarePlan
	CareTeam_list                       []*r4b.CareTeam
	CatalogEntry_list                   []*r4b.CatalogEntry
	ChargeItem_list                     []*r4b.ChargeItem
	ChargeItemDefinition_list           []*r4b.ChargeItemDefinition
	Citation_list                       []*r4b.Citation
	Claim_list                          []*r4b.Claim
	ClaimResponse_list                  []*r4b.ClaimResponse
	ClinicalImpression_list             []*r4b.ClinicalImpression
	ClinicalUseDefinition_list          []*r4b.ClinicalUseDefinition
	CodeSystem_list                     []*r4b.CodeSystem
	Communication_list                  []*r4b.Communication
	CommunicationRequest_list           []*r4b.CommunicationRequest
	CompartmentDefinition_list          []*r4b.CompartmentDefinition
	Composition_list                    []*r4b.Composition
	ConceptMap_list                     []*r4b.ConceptMap
	Condition_list                      []*r4b.Condition
	Consent_list                        []*r4b.Consent
	Contract_list                       []*r4b.Contract
	Coverage_list                       []*r4b.Coverage
	CoverageEligibilityRequest_list     []*r4b.CoverageEligibilityRequest
	CoverageEligibilityResponse_list    []*r4b.CoverageEligibilityResponse
	DetectedIssue_list                  []*r4b.DetectedIssue
	Device_list                         []*r4b.Device
	DeviceDefinition_list               []*r4b.DeviceDefinition
	DeviceMetric_list                   []*r4b.DeviceMetric
	DeviceRequest_list                  []*r4b.DeviceRequest
	DeviceUseStatement_list             []*r4b.DeviceUseStatement
	DiagnosticReport_list               []*r4b.DiagnosticReport
	DocumentManifest_list               []*r4b.DocumentManifest
	DocumentReference_list              []*r4b.DocumentReference
	Encounter_list                      []*r4b.Encounter
	Endpoint_list                       []*r4b.Endpoint
	EnrollmentRequest_list              []*r4b.EnrollmentRequest
	EnrollmentResponse_list             []*r4b.EnrollmentResponse
	EpisodeOfCare_list                  []*r4b.EpisodeOfCare
	EventDefinition_list                []*r4b.EventDefinition
	Evidence_list                       []*r4b.Evidence
	EvidenceReport_list                 []*r4b.EvidenceReport
	EvidenceVariable_list               []*r4b.EvidenceVariable
	ExampleScenario_list                []*r4b.ExampleScenario
	ExplanationOfBenefit_list           []*r4b.ExplanationOfBenefit
	FamilyMemberHistory_list            []*r4b.FamilyMemberHistory
	Flag_list                           []*r4b.Flag
	Goal_list                           []*r4b.Goal
	GraphDefinition_list                []*r4b.GraphDefinition
	Group_list                          []*r4b.Group
	GuidanceResponse_list               []*r4b.GuidanceResponse
	HealthcareService_list              []*r4b.HealthcareService
	ImagingStudy_list                   []*r4b.ImagingStudy
	Immunization_list                   []*r4b.Immunization
	ImmunizationEvaluation_list         []*r4b.ImmunizationEvaluation
	ImmunizationRecommendation_list     []*r4b.ImmunizationRecommendation
	ImplementationGuide_list            []*r4b.ImplementationGuide
	Ingredient_list                     []*r4b.Ingredient
	InsurancePlan_list                  []*r4b.InsurancePlan
	Invoice_list                        []*r4b.Invoice
	Library_list                        []*r4b.Library
	Linkage_list                        []*r4b.Linkage
	List_list                           []*r4b.List
	Location_list                       []*r4b.Location
	ManufacturedItemDefinition_list     []*r4b.ManufacturedItemDefinition
	Measure_list                        []*r4b.Measure
	MeasureReport_list                  []*r4b.MeasureReport
	Media_list                          []*r4b.Media
	Medication_list                     []*r4b.Medication
	MedicationAdministration_list       []*r4b.MedicationAdministration
	MedicationDispense_list             []*r4b.MedicationDispense
	MedicationKnowledge_list            []*r4b.MedicationKnowledge
	MedicationRequest_list              []*r4b.MedicationRequest
	MedicationStatement_list            []*r4b.MedicationStatement
	MedicinalProductDefinition_list     []*r4b.MedicinalProductDefinition
	MessageDefinition_list              []*r4b.MessageDefinition
	MessageHeader_list                  []*r4b.MessageHeader
	MolecularSequence_list              []*r4b.MolecularSequence
	NamingSystem_list                   []*r4b.NamingSystem
	NutritionOrder_list                 []*r4b.NutritionOrder
	NutritionProduct_list               []*r4b.NutritionProduct
	Observation_list                    []*r4b.Observation
	ObservationDefinition_list          []*r4b.ObservationDefinition
	OperationDefinition_list            []*r4b.OperationDefinition
	OperationOutcome_list               []*r4b.OperationOutcome
	Organization_list                   []*r4b.Organization
	OrganizationAffiliation_list        []*r4b.OrganizationAffiliation
	PackagedProductDefinition_list      []*r4b.PackagedProductDefinition
	Patient_list                        []*r4b.Patient
	PaymentNotice_list                  []*r4b.PaymentNotice
	PaymentReconciliation_list          []*r4b.PaymentReconciliation
	Person_list                         []*r4b.Person
	PlanDefinition_list                 []*r4b.PlanDefinition
	Practitioner_list                   []*r4b.Practitioner
	PractitionerRole_list               []*r4b.PractitionerRole
	Procedure_list                      []*r4b.Procedure
	Provenance_list                     []*r4b.Provenance
	Questionnaire_list                  []*r4b.Questionnaire
	QuestionnaireResponse_list          []*r4b.QuestionnaireResponse
	RegulatedAuthorization_list         []*r4b.RegulatedAuthorization
	RelatedPerson_list                  []*r4b.RelatedPerson
	RequestGroup_list                   []*r4b.RequestGroup
	ResearchDefinition_list             []*r4b.ResearchDefinition
	ResearchElementDefinition_list      []*r4b.ResearchElementDefinition
	ResearchStudy_list                  []*r4b.ResearchStudy
	ResearchSubject_list                []*r4b.ResearchSubject
	RiskAssessment_list                 []*r4b.RiskAssessment
	Schedule_list                       []*r4b.Schedule
	SearchParameter_list                []*r4b.SearchParameter
	ServiceRequest_list                 []*r4b.ServiceRequest
	Slot_list                           []*r4b.Slot
	Specimen_list                       []*r4b.Specimen
	SpecimenDefinition_list             []*r4b.SpecimenDefinition
	StructureDefinition_list            []*r4b.StructureDefinition
	StructureMap_list                   []*r4b.StructureMap
	Subscription_list                   []*r4b.Subscription
	SubscriptionStatus_list             []*r4b.SubscriptionStatus
	SubscriptionTopic_list              []*r4b.SubscriptionTopic
	Substance_list                      []*r4b.Substance
	SubstanceDefinition_list            []*r4b.SubstanceDefinition
	SupplyDelivery_list                 []*r4b.SupplyDelivery
	SupplyRequest_list                  []*r4b.SupplyRequest
	Task_list                           []*r4b.Task
	TerminologyCapabilities_list        []*r4b.TerminologyCapabilities
	TestReport_list                     []*r4b.TestReport
	TestScript_list                     []*r4b.TestScript
	ValueSet_list                       []*r4b.ValueSet
	VerificationResult_list             []*r4b.VerificationResult
	VisionPrescription_list             []*r4b.VisionPrescription
}

func BundleToGroup(bundle *r4.Bundle) (*ResourceGroup, error) {
	grp := ResourceGroup{}
	for _, e := range bundle.Entry {
		switch res := e.Resource.(type) {
		case *r4b.Account:
			grp.Account_list = append(grp.Account_list, res)
		case *r4b.ActivityDefinition:
			grp.ActivityDefinition_list = append(grp.ActivityDefinition_list, res)
		case *r4b.AdministrableProductDefinition:
			grp.AdministrableProductDefinition_list = append(grp.AdministrableProductDefinition_list, res)
		case *r4b.AdverseEvent:
			grp.AdverseEvent_list = append(grp.AdverseEvent_list, res)
		case *r4b.AllergyIntolerance:
			grp.AllergyIntolerance_list = append(grp.AllergyIntolerance_list, res)
		case *r4b.Appointment:
			grp.Appointment_list = append(grp.Appointment_list, res)
		case *r4b.AppointmentResponse:
			grp.AppointmentResponse_list = append(grp.AppointmentResponse_list, res)
		case *r4b.AuditEvent:
			grp.AuditEvent_list = append(grp.AuditEvent_list, res)
		case *r4b.Basic:
			grp.Basic_list = append(grp.Basic_list, res)
		case *r4b.Binary:
			grp.Binary_list = append(grp.Binary_list, res)
		case *r4b.BiologicallyDerivedProduct:
			grp.BiologicallyDerivedProduct_list = append(grp.BiologicallyDerivedProduct_list, res)
		case *r4b.BodyStructure:
			grp.BodyStructure_list = append(grp.BodyStructure_list, res)
		case *r4b.CapabilityStatement:
			grp.CapabilityStatement_list = append(grp.CapabilityStatement_list, res)
		case *r4b.CarePlan:
			grp.CarePlan_list = append(grp.CarePlan_list, res)
		case *r4b.CareTeam:
			grp.CareTeam_list = append(grp.CareTeam_list, res)
		case *r4b.CatalogEntry:
			grp.CatalogEntry_list = append(grp.CatalogEntry_list, res)
		case *r4b.ChargeItem:
			grp.ChargeItem_list = append(grp.ChargeItem_list, res)
		case *r4b.ChargeItemDefinition:
			grp.ChargeItemDefinition_list = append(grp.ChargeItemDefinition_list, res)
		case *r4b.Citation:
			grp.Citation_list = append(grp.Citation_list, res)
		case *r4b.Claim:
			grp.Claim_list = append(grp.Claim_list, res)
		case *r4b.ClaimResponse:
			grp.ClaimResponse_list = append(grp.ClaimResponse_list, res)
		case *r4b.ClinicalImpression:
			grp.ClinicalImpression_list = append(grp.ClinicalImpression_list, res)
		case *r4b.ClinicalUseDefinition:
			grp.ClinicalUseDefinition_list = append(grp.ClinicalUseDefinition_list, res)
		case *r4b.CodeSystem:
			grp.CodeSystem_list = append(grp.CodeSystem_list, res)
		case *r4b.Communication:
			grp.Communication_list = append(grp.Communication_list, res)
		case *r4b.CommunicationRequest:
			grp.CommunicationRequest_list = append(grp.CommunicationRequest_list, res)
		case *r4b.CompartmentDefinition:
			grp.CompartmentDefinition_list = append(grp.CompartmentDefinition_list, res)
		case *r4b.Composition:
			grp.Composition_list = append(grp.Composition_list, res)
		case *r4b.ConceptMap:
			grp.ConceptMap_list = append(grp.ConceptMap_list, res)
		case *r4b.Condition:
			grp.Condition_list = append(grp.Condition_list, res)
		case *r4b.Consent:
			grp.Consent_list = append(grp.Consent_list, res)
		case *r4b.Contract:
			grp.Contract_list = append(grp.Contract_list, res)
		case *r4b.Coverage:
			grp.Coverage_list = append(grp.Coverage_list, res)
		case *r4b.CoverageEligibilityRequest:
			grp.CoverageEligibilityRequest_list = append(grp.CoverageEligibilityRequest_list, res)
		case *r4b.CoverageEligibilityResponse:
			grp.CoverageEligibilityResponse_list = append(grp.CoverageEligibilityResponse_list, res)
		case *r4b.DetectedIssue:
			grp.DetectedIssue_list = append(grp.DetectedIssue_list, res)
		case *r4b.Device:
			grp.Device_list = append(grp.Device_list, res)
		case *r4b.DeviceDefinition:
			grp.DeviceDefinition_list = append(grp.DeviceDefinition_list, res)
		case *r4b.DeviceMetric:
			grp.DeviceMetric_list = append(grp.DeviceMetric_list, res)
		case *r4b.DeviceRequest:
			grp.DeviceRequest_list = append(grp.DeviceRequest_list, res)
		case *r4b.DeviceUseStatement:
			grp.DeviceUseStatement_list = append(grp.DeviceUseStatement_list, res)
		case *r4b.DiagnosticReport:
			grp.DiagnosticReport_list = append(grp.DiagnosticReport_list, res)
		case *r4b.DocumentManifest:
			grp.DocumentManifest_list = append(grp.DocumentManifest_list, res)
		case *r4b.DocumentReference:
			grp.DocumentReference_list = append(grp.DocumentReference_list, res)
		case *r4b.Encounter:
			grp.Encounter_list = append(grp.Encounter_list, res)
		case *r4b.Endpoint:
			grp.Endpoint_list = append(grp.Endpoint_list, res)
		case *r4b.EnrollmentRequest:
			grp.EnrollmentRequest_list = append(grp.EnrollmentRequest_list, res)
		case *r4b.EnrollmentResponse:
			grp.EnrollmentResponse_list = append(grp.EnrollmentResponse_list, res)
		case *r4b.EpisodeOfCare:
			grp.EpisodeOfCare_list = append(grp.EpisodeOfCare_list, res)
		case *r4b.EventDefinition:
			grp.EventDefinition_list = append(grp.EventDefinition_list, res)
		case *r4b.Evidence:
			grp.Evidence_list = append(grp.Evidence_list, res)
		case *r4b.EvidenceReport:
			grp.EvidenceReport_list = append(grp.EvidenceReport_list, res)
		case *r4b.EvidenceVariable:
			grp.EvidenceVariable_list = append(grp.EvidenceVariable_list, res)
		case *r4b.ExampleScenario:
			grp.ExampleScenario_list = append(grp.ExampleScenario_list, res)
		case *r4b.ExplanationOfBenefit:
			grp.ExplanationOfBenefit_list = append(grp.ExplanationOfBenefit_list, res)
		case *r4b.FamilyMemberHistory:
			grp.FamilyMemberHistory_list = append(grp.FamilyMemberHistory_list, res)
		case *r4b.Flag:
			grp.Flag_list = append(grp.Flag_list, res)
		case *r4b.Goal:
			grp.Goal_list = append(grp.Goal_list, res)
		case *r4b.GraphDefinition:
			grp.GraphDefinition_list = append(grp.GraphDefinition_list, res)
		case *r4b.Group:
			grp.Group_list = append(grp.Group_list, res)
		case *r4b.GuidanceResponse:
			grp.GuidanceResponse_list = append(grp.GuidanceResponse_list, res)
		case *r4b.HealthcareService:
			grp.HealthcareService_list = append(grp.HealthcareService_list, res)
		case *r4b.ImagingStudy:
			grp.ImagingStudy_list = append(grp.ImagingStudy_list, res)
		case *r4b.Immunization:
			grp.Immunization_list = append(grp.Immunization_list, res)
		case *r4b.ImmunizationEvaluation:
			grp.ImmunizationEvaluation_list = append(grp.ImmunizationEvaluation_list, res)
		case *r4b.ImmunizationRecommendation:
			grp.ImmunizationRecommendation_list = append(grp.ImmunizationRecommendation_list, res)
		case *r4b.ImplementationGuide:
			grp.ImplementationGuide_list = append(grp.ImplementationGuide_list, res)
		case *r4b.Ingredient:
			grp.Ingredient_list = append(grp.Ingredient_list, res)
		case *r4b.InsurancePlan:
			grp.InsurancePlan_list = append(grp.InsurancePlan_list, res)
		case *r4b.Invoice:
			grp.Invoice_list = append(grp.Invoice_list, res)
		case *r4b.Library:
			grp.Library_list = append(grp.Library_list, res)
		case *r4b.Linkage:
			grp.Linkage_list = append(grp.Linkage_list, res)
		case *r4b.List:
			grp.List_list = append(grp.List_list, res)
		case *r4b.Location:
			grp.Location_list = append(grp.Location_list, res)
		case *r4b.ManufacturedItemDefinition:
			grp.ManufacturedItemDefinition_list = append(grp.ManufacturedItemDefinition_list, res)
		case *r4b.Measure:
			grp.Measure_list = append(grp.Measure_list, res)
		case *r4b.MeasureReport:
			grp.MeasureReport_list = append(grp.MeasureReport_list, res)
		case *r4b.Media:
			grp.Media_list = append(grp.Media_list, res)
		case *r4b.Medication:
			grp.Medication_list = append(grp.Medication_list, res)
		case *r4b.MedicationAdministration:
			grp.MedicationAdministration_list = append(grp.MedicationAdministration_list, res)
		case *r4b.MedicationDispense:
			grp.MedicationDispense_list = append(grp.MedicationDispense_list, res)
		case *r4b.MedicationKnowledge:
			grp.MedicationKnowledge_list = append(grp.MedicationKnowledge_list, res)
		case *r4b.MedicationRequest:
			grp.MedicationRequest_list = append(grp.MedicationRequest_list, res)
		case *r4b.MedicationStatement:
			grp.MedicationStatement_list = append(grp.MedicationStatement_list, res)
		case *r4b.MedicinalProductDefinition:
			grp.MedicinalProductDefinition_list = append(grp.MedicinalProductDefinition_list, res)
		case *r4b.MessageDefinition:
			grp.MessageDefinition_list = append(grp.MessageDefinition_list, res)
		case *r4b.MessageHeader:
			grp.MessageHeader_list = append(grp.MessageHeader_list, res)
		case *r4b.MolecularSequence:
			grp.MolecularSequence_list = append(grp.MolecularSequence_list, res)
		case *r4b.NamingSystem:
			grp.NamingSystem_list = append(grp.NamingSystem_list, res)
		case *r4b.NutritionOrder:
			grp.NutritionOrder_list = append(grp.NutritionOrder_list, res)
		case *r4b.NutritionProduct:
			grp.NutritionProduct_list = append(grp.NutritionProduct_list, res)
		case *r4b.Observation:
			grp.Observation_list = append(grp.Observation_list, res)
		case *r4b.ObservationDefinition:
			grp.ObservationDefinition_list = append(grp.ObservationDefinition_list, res)
		case *r4b.OperationDefinition:
			grp.OperationDefinition_list = append(grp.OperationDefinition_list, res)
		case *r4b.OperationOutcome:
			grp.OperationOutcome_list = append(grp.OperationOutcome_list, res)
		case *r4b.Organization:
			grp.Organization_list = append(grp.Organization_list, res)
		case *r4b.OrganizationAffiliation:
			grp.OrganizationAffiliation_list = append(grp.OrganizationAffiliation_list, res)
		case *r4b.PackagedProductDefinition:
			grp.PackagedProductDefinition_list = append(grp.PackagedProductDefinition_list, res)
		case *r4b.Patient:
			grp.Patient_list = append(grp.Patient_list, res)
		case *r4b.PaymentNotice:
			grp.PaymentNotice_list = append(grp.PaymentNotice_list, res)
		case *r4b.PaymentReconciliation:
			grp.PaymentReconciliation_list = append(grp.PaymentReconciliation_list, res)
		case *r4b.Person:
			grp.Person_list = append(grp.Person_list, res)
		case *r4b.PlanDefinition:
			grp.PlanDefinition_list = append(grp.PlanDefinition_list, res)
		case *r4b.Practitioner:
			grp.Practitioner_list = append(grp.Practitioner_list, res)
		case *r4b.PractitionerRole:
			grp.PractitionerRole_list = append(grp.PractitionerRole_list, res)
		case *r4b.Procedure:
			grp.Procedure_list = append(grp.Procedure_list, res)
		case *r4b.Provenance:
			grp.Provenance_list = append(grp.Provenance_list, res)
		case *r4b.Questionnaire:
			grp.Questionnaire_list = append(grp.Questionnaire_list, res)
		case *r4b.QuestionnaireResponse:
			grp.QuestionnaireResponse_list = append(grp.QuestionnaireResponse_list, res)
		case *r4b.RegulatedAuthorization:
			grp.RegulatedAuthorization_list = append(grp.RegulatedAuthorization_list, res)
		case *r4b.RelatedPerson:
			grp.RelatedPerson_list = append(grp.RelatedPerson_list, res)
		case *r4b.RequestGroup:
			grp.RequestGroup_list = append(grp.RequestGroup_list, res)
		case *r4b.ResearchDefinition:
			grp.ResearchDefinition_list = append(grp.ResearchDefinition_list, res)
		case *r4b.ResearchElementDefinition:
			grp.ResearchElementDefinition_list = append(grp.ResearchElementDefinition_list, res)
		case *r4b.ResearchStudy:
			grp.ResearchStudy_list = append(grp.ResearchStudy_list, res)
		case *r4b.ResearchSubject:
			grp.ResearchSubject_list = append(grp.ResearchSubject_list, res)
		case *r4b.RiskAssessment:
			grp.RiskAssessment_list = append(grp.RiskAssessment_list, res)
		case *r4b.Schedule:
			grp.Schedule_list = append(grp.Schedule_list, res)
		case *r4b.SearchParameter:
			grp.SearchParameter_list = append(grp.SearchParameter_list, res)
		case *r4b.ServiceRequest:
			grp.ServiceRequest_list = append(grp.ServiceRequest_list, res)
		case *r4b.Slot:
			grp.Slot_list = append(grp.Slot_list, res)
		case *r4b.Specimen:
			grp.Specimen_list = append(grp.Specimen_list, res)
		case *r4b.SpecimenDefinition:
			grp.SpecimenDefinition_list = append(grp.SpecimenDefinition_list, res)
		case *r4b.StructureDefinition:
			grp.StructureDefinition_list = append(grp.StructureDefinition_list, res)
		case *r4b.StructureMap:
			grp.StructureMap_list = append(grp.StructureMap_list, res)
		case *r4b.Subscription:
			grp.Subscription_list = append(grp.Subscription_list, res)
		case *r4b.SubscriptionStatus:
			grp.SubscriptionStatus_list = append(grp.SubscriptionStatus_list, res)
		case *r4b.SubscriptionTopic:
			grp.SubscriptionTopic_list = append(grp.SubscriptionTopic_list, res)
		case *r4b.Substance:
			grp.Substance_list = append(grp.Substance_list, res)
		case *r4b.SubstanceDefinition:
			grp.SubstanceDefinition_list = append(grp.SubstanceDefinition_list, res)
		case *r4b.SupplyDelivery:
			grp.SupplyDelivery_list = append(grp.SupplyDelivery_list, res)
		case *r4b.SupplyRequest:
			grp.SupplyRequest_list = append(grp.SupplyRequest_list, res)
		case *r4b.Task:
			grp.Task_list = append(grp.Task_list, res)
		case *r4b.TerminologyCapabilities:
			grp.TerminologyCapabilities_list = append(grp.TerminologyCapabilities_list, res)
		case *r4b.TestReport:
			grp.TestReport_list = append(grp.TestReport_list, res)
		case *r4b.TestScript:
			grp.TestScript_list = append(grp.TestScript_list, res)
		case *r4b.ValueSet:
			grp.ValueSet_list = append(grp.ValueSet_list, res)
		case *r4b.VerificationResult:
			grp.VerificationResult_list = append(grp.VerificationResult_list, res)
		case *r4b.VisionPrescription:
			grp.VisionPrescription_list = append(grp.VisionPrescription_list, res)
		default:
			return nil, errors.New("bundle entry not a domain resource, could not put in resourcegroup")
		}
	}
	return &grp, nil
}
