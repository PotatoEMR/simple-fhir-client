package r4Client

import r4 "github.com/PotatoEMR/simple-fhir-client/r4"
import "errors"

type ResourceGroup struct {
	Account_list                           []*r4.Account
	ActivityDefinition_list                []*r4.ActivityDefinition
	AdverseEvent_list                      []*r4.AdverseEvent
	AllergyIntolerance_list                []*r4.AllergyIntolerance
	Appointment_list                       []*r4.Appointment
	AppointmentResponse_list               []*r4.AppointmentResponse
	AuditEvent_list                        []*r4.AuditEvent
	Basic_list                             []*r4.Basic
	Binary_list                            []*r4.Binary
	BiologicallyDerivedProduct_list        []*r4.BiologicallyDerivedProduct
	BodyStructure_list                     []*r4.BodyStructure
	CapabilityStatement_list               []*r4.CapabilityStatement
	CarePlan_list                          []*r4.CarePlan
	CareTeam_list                          []*r4.CareTeam
	CatalogEntry_list                      []*r4.CatalogEntry
	ChargeItem_list                        []*r4.ChargeItem
	ChargeItemDefinition_list              []*r4.ChargeItemDefinition
	Claim_list                             []*r4.Claim
	ClaimResponse_list                     []*r4.ClaimResponse
	ClinicalImpression_list                []*r4.ClinicalImpression
	CodeSystem_list                        []*r4.CodeSystem
	Communication_list                     []*r4.Communication
	CommunicationRequest_list              []*r4.CommunicationRequest
	CompartmentDefinition_list             []*r4.CompartmentDefinition
	Composition_list                       []*r4.Composition
	ConceptMap_list                        []*r4.ConceptMap
	Condition_list                         []*r4.Condition
	Consent_list                           []*r4.Consent
	Contract_list                          []*r4.Contract
	Coverage_list                          []*r4.Coverage
	CoverageEligibilityRequest_list        []*r4.CoverageEligibilityRequest
	CoverageEligibilityResponse_list       []*r4.CoverageEligibilityResponse
	DetectedIssue_list                     []*r4.DetectedIssue
	Device_list                            []*r4.Device
	DeviceDefinition_list                  []*r4.DeviceDefinition
	DeviceMetric_list                      []*r4.DeviceMetric
	DeviceRequest_list                     []*r4.DeviceRequest
	DeviceUseStatement_list                []*r4.DeviceUseStatement
	DiagnosticReport_list                  []*r4.DiagnosticReport
	DocumentManifest_list                  []*r4.DocumentManifest
	DocumentReference_list                 []*r4.DocumentReference
	EffectEvidenceSynthesis_list           []*r4.EffectEvidenceSynthesis
	Encounter_list                         []*r4.Encounter
	Endpoint_list                          []*r4.Endpoint
	EnrollmentRequest_list                 []*r4.EnrollmentRequest
	EnrollmentResponse_list                []*r4.EnrollmentResponse
	EpisodeOfCare_list                     []*r4.EpisodeOfCare
	EventDefinition_list                   []*r4.EventDefinition
	Evidence_list                          []*r4.Evidence
	EvidenceVariable_list                  []*r4.EvidenceVariable
	ExampleScenario_list                   []*r4.ExampleScenario
	ExplanationOfBenefit_list              []*r4.ExplanationOfBenefit
	FamilyMemberHistory_list               []*r4.FamilyMemberHistory
	Flag_list                              []*r4.Flag
	Goal_list                              []*r4.Goal
	GraphDefinition_list                   []*r4.GraphDefinition
	Group_list                             []*r4.Group
	GuidanceResponse_list                  []*r4.GuidanceResponse
	HealthcareService_list                 []*r4.HealthcareService
	ImagingStudy_list                      []*r4.ImagingStudy
	Immunization_list                      []*r4.Immunization
	ImmunizationEvaluation_list            []*r4.ImmunizationEvaluation
	ImmunizationRecommendation_list        []*r4.ImmunizationRecommendation
	ImplementationGuide_list               []*r4.ImplementationGuide
	InsurancePlan_list                     []*r4.InsurancePlan
	Invoice_list                           []*r4.Invoice
	Library_list                           []*r4.Library
	Linkage_list                           []*r4.Linkage
	List_list                              []*r4.List
	Location_list                          []*r4.Location
	Measure_list                           []*r4.Measure
	MeasureReport_list                     []*r4.MeasureReport
	Media_list                             []*r4.Media
	Medication_list                        []*r4.Medication
	MedicationAdministration_list          []*r4.MedicationAdministration
	MedicationDispense_list                []*r4.MedicationDispense
	MedicationKnowledge_list               []*r4.MedicationKnowledge
	MedicationRequest_list                 []*r4.MedicationRequest
	MedicationStatement_list               []*r4.MedicationStatement
	MedicinalProduct_list                  []*r4.MedicinalProduct
	MedicinalProductAuthorization_list     []*r4.MedicinalProductAuthorization
	MedicinalProductContraindication_list  []*r4.MedicinalProductContraindication
	MedicinalProductIndication_list        []*r4.MedicinalProductIndication
	MedicinalProductIngredient_list        []*r4.MedicinalProductIngredient
	MedicinalProductInteraction_list       []*r4.MedicinalProductInteraction
	MedicinalProductManufactured_list      []*r4.MedicinalProductManufactured
	MedicinalProductPackaged_list          []*r4.MedicinalProductPackaged
	MedicinalProductPharmaceutical_list    []*r4.MedicinalProductPharmaceutical
	MedicinalProductUndesirableEffect_list []*r4.MedicinalProductUndesirableEffect
	MessageDefinition_list                 []*r4.MessageDefinition
	MessageHeader_list                     []*r4.MessageHeader
	MolecularSequence_list                 []*r4.MolecularSequence
	NamingSystem_list                      []*r4.NamingSystem
	NutritionOrder_list                    []*r4.NutritionOrder
	Observation_list                       []*r4.Observation
	ObservationDefinition_list             []*r4.ObservationDefinition
	OperationDefinition_list               []*r4.OperationDefinition
	OperationOutcome_list                  []*r4.OperationOutcome
	Organization_list                      []*r4.Organization
	OrganizationAffiliation_list           []*r4.OrganizationAffiliation
	Patient_list                           []*r4.Patient
	PaymentNotice_list                     []*r4.PaymentNotice
	PaymentReconciliation_list             []*r4.PaymentReconciliation
	Person_list                            []*r4.Person
	PlanDefinition_list                    []*r4.PlanDefinition
	Practitioner_list                      []*r4.Practitioner
	PractitionerRole_list                  []*r4.PractitionerRole
	Procedure_list                         []*r4.Procedure
	Provenance_list                        []*r4.Provenance
	Questionnaire_list                     []*r4.Questionnaire
	QuestionnaireResponse_list             []*r4.QuestionnaireResponse
	RelatedPerson_list                     []*r4.RelatedPerson
	RequestGroup_list                      []*r4.RequestGroup
	ResearchDefinition_list                []*r4.ResearchDefinition
	ResearchElementDefinition_list         []*r4.ResearchElementDefinition
	ResearchStudy_list                     []*r4.ResearchStudy
	ResearchSubject_list                   []*r4.ResearchSubject
	RiskAssessment_list                    []*r4.RiskAssessment
	RiskEvidenceSynthesis_list             []*r4.RiskEvidenceSynthesis
	Schedule_list                          []*r4.Schedule
	SearchParameter_list                   []*r4.SearchParameter
	ServiceRequest_list                    []*r4.ServiceRequest
	Slot_list                              []*r4.Slot
	Specimen_list                          []*r4.Specimen
	SpecimenDefinition_list                []*r4.SpecimenDefinition
	StructureDefinition_list               []*r4.StructureDefinition
	StructureMap_list                      []*r4.StructureMap
	Subscription_list                      []*r4.Subscription
	Substance_list                         []*r4.Substance
	SubstanceNucleicAcid_list              []*r4.SubstanceNucleicAcid
	SubstancePolymer_list                  []*r4.SubstancePolymer
	SubstanceProtein_list                  []*r4.SubstanceProtein
	SubstanceReferenceInformation_list     []*r4.SubstanceReferenceInformation
	SubstanceSourceMaterial_list           []*r4.SubstanceSourceMaterial
	SubstanceSpecification_list            []*r4.SubstanceSpecification
	SupplyDelivery_list                    []*r4.SupplyDelivery
	SupplyRequest_list                     []*r4.SupplyRequest
	Task_list                              []*r4.Task
	TerminologyCapabilities_list           []*r4.TerminologyCapabilities
	TestReport_list                        []*r4.TestReport
	TestScript_list                        []*r4.TestScript
	ValueSet_list                          []*r4.ValueSet
	VerificationResult_list                []*r4.VerificationResult
	VisionPrescription_list                []*r4.VisionPrescription
}

func BundleToGroup(bundle *r4.Bundle) (*ResourceGroup, error) {
	grp := ResourceGroup{}
	for _, e := range bundle.Entry {
		switch res := e.Resource.(type) {
		case *r4.Account:
			grp.Account_list = append(grp.Account_list, res)
		case *r4.ActivityDefinition:
			grp.ActivityDefinition_list = append(grp.ActivityDefinition_list, res)
		case *r4.AdverseEvent:
			grp.AdverseEvent_list = append(grp.AdverseEvent_list, res)
		case *r4.AllergyIntolerance:
			grp.AllergyIntolerance_list = append(grp.AllergyIntolerance_list, res)
		case *r4.Appointment:
			grp.Appointment_list = append(grp.Appointment_list, res)
		case *r4.AppointmentResponse:
			grp.AppointmentResponse_list = append(grp.AppointmentResponse_list, res)
		case *r4.AuditEvent:
			grp.AuditEvent_list = append(grp.AuditEvent_list, res)
		case *r4.Basic:
			grp.Basic_list = append(grp.Basic_list, res)
		case *r4.Binary:
			grp.Binary_list = append(grp.Binary_list, res)
		case *r4.BiologicallyDerivedProduct:
			grp.BiologicallyDerivedProduct_list = append(grp.BiologicallyDerivedProduct_list, res)
		case *r4.BodyStructure:
			grp.BodyStructure_list = append(grp.BodyStructure_list, res)
		case *r4.CapabilityStatement:
			grp.CapabilityStatement_list = append(grp.CapabilityStatement_list, res)
		case *r4.CarePlan:
			grp.CarePlan_list = append(grp.CarePlan_list, res)
		case *r4.CareTeam:
			grp.CareTeam_list = append(grp.CareTeam_list, res)
		case *r4.CatalogEntry:
			grp.CatalogEntry_list = append(grp.CatalogEntry_list, res)
		case *r4.ChargeItem:
			grp.ChargeItem_list = append(grp.ChargeItem_list, res)
		case *r4.ChargeItemDefinition:
			grp.ChargeItemDefinition_list = append(grp.ChargeItemDefinition_list, res)
		case *r4.Claim:
			grp.Claim_list = append(grp.Claim_list, res)
		case *r4.ClaimResponse:
			grp.ClaimResponse_list = append(grp.ClaimResponse_list, res)
		case *r4.ClinicalImpression:
			grp.ClinicalImpression_list = append(grp.ClinicalImpression_list, res)
		case *r4.CodeSystem:
			grp.CodeSystem_list = append(grp.CodeSystem_list, res)
		case *r4.Communication:
			grp.Communication_list = append(grp.Communication_list, res)
		case *r4.CommunicationRequest:
			grp.CommunicationRequest_list = append(grp.CommunicationRequest_list, res)
		case *r4.CompartmentDefinition:
			grp.CompartmentDefinition_list = append(grp.CompartmentDefinition_list, res)
		case *r4.Composition:
			grp.Composition_list = append(grp.Composition_list, res)
		case *r4.ConceptMap:
			grp.ConceptMap_list = append(grp.ConceptMap_list, res)
		case *r4.Condition:
			grp.Condition_list = append(grp.Condition_list, res)
		case *r4.Consent:
			grp.Consent_list = append(grp.Consent_list, res)
		case *r4.Contract:
			grp.Contract_list = append(grp.Contract_list, res)
		case *r4.Coverage:
			grp.Coverage_list = append(grp.Coverage_list, res)
		case *r4.CoverageEligibilityRequest:
			grp.CoverageEligibilityRequest_list = append(grp.CoverageEligibilityRequest_list, res)
		case *r4.CoverageEligibilityResponse:
			grp.CoverageEligibilityResponse_list = append(grp.CoverageEligibilityResponse_list, res)
		case *r4.DetectedIssue:
			grp.DetectedIssue_list = append(grp.DetectedIssue_list, res)
		case *r4.Device:
			grp.Device_list = append(grp.Device_list, res)
		case *r4.DeviceDefinition:
			grp.DeviceDefinition_list = append(grp.DeviceDefinition_list, res)
		case *r4.DeviceMetric:
			grp.DeviceMetric_list = append(grp.DeviceMetric_list, res)
		case *r4.DeviceRequest:
			grp.DeviceRequest_list = append(grp.DeviceRequest_list, res)
		case *r4.DeviceUseStatement:
			grp.DeviceUseStatement_list = append(grp.DeviceUseStatement_list, res)
		case *r4.DiagnosticReport:
			grp.DiagnosticReport_list = append(grp.DiagnosticReport_list, res)
		case *r4.DocumentManifest:
			grp.DocumentManifest_list = append(grp.DocumentManifest_list, res)
		case *r4.DocumentReference:
			grp.DocumentReference_list = append(grp.DocumentReference_list, res)
		case *r4.EffectEvidenceSynthesis:
			grp.EffectEvidenceSynthesis_list = append(grp.EffectEvidenceSynthesis_list, res)
		case *r4.Encounter:
			grp.Encounter_list = append(grp.Encounter_list, res)
		case *r4.Endpoint:
			grp.Endpoint_list = append(grp.Endpoint_list, res)
		case *r4.EnrollmentRequest:
			grp.EnrollmentRequest_list = append(grp.EnrollmentRequest_list, res)
		case *r4.EnrollmentResponse:
			grp.EnrollmentResponse_list = append(grp.EnrollmentResponse_list, res)
		case *r4.EpisodeOfCare:
			grp.EpisodeOfCare_list = append(grp.EpisodeOfCare_list, res)
		case *r4.EventDefinition:
			grp.EventDefinition_list = append(grp.EventDefinition_list, res)
		case *r4.Evidence:
			grp.Evidence_list = append(grp.Evidence_list, res)
		case *r4.EvidenceVariable:
			grp.EvidenceVariable_list = append(grp.EvidenceVariable_list, res)
		case *r4.ExampleScenario:
			grp.ExampleScenario_list = append(grp.ExampleScenario_list, res)
		case *r4.ExplanationOfBenefit:
			grp.ExplanationOfBenefit_list = append(grp.ExplanationOfBenefit_list, res)
		case *r4.FamilyMemberHistory:
			grp.FamilyMemberHistory_list = append(grp.FamilyMemberHistory_list, res)
		case *r4.Flag:
			grp.Flag_list = append(grp.Flag_list, res)
		case *r4.Goal:
			grp.Goal_list = append(grp.Goal_list, res)
		case *r4.GraphDefinition:
			grp.GraphDefinition_list = append(grp.GraphDefinition_list, res)
		case *r4.Group:
			grp.Group_list = append(grp.Group_list, res)
		case *r4.GuidanceResponse:
			grp.GuidanceResponse_list = append(grp.GuidanceResponse_list, res)
		case *r4.HealthcareService:
			grp.HealthcareService_list = append(grp.HealthcareService_list, res)
		case *r4.ImagingStudy:
			grp.ImagingStudy_list = append(grp.ImagingStudy_list, res)
		case *r4.Immunization:
			grp.Immunization_list = append(grp.Immunization_list, res)
		case *r4.ImmunizationEvaluation:
			grp.ImmunizationEvaluation_list = append(grp.ImmunizationEvaluation_list, res)
		case *r4.ImmunizationRecommendation:
			grp.ImmunizationRecommendation_list = append(grp.ImmunizationRecommendation_list, res)
		case *r4.ImplementationGuide:
			grp.ImplementationGuide_list = append(grp.ImplementationGuide_list, res)
		case *r4.InsurancePlan:
			grp.InsurancePlan_list = append(grp.InsurancePlan_list, res)
		case *r4.Invoice:
			grp.Invoice_list = append(grp.Invoice_list, res)
		case *r4.Library:
			grp.Library_list = append(grp.Library_list, res)
		case *r4.Linkage:
			grp.Linkage_list = append(grp.Linkage_list, res)
		case *r4.List:
			grp.List_list = append(grp.List_list, res)
		case *r4.Location:
			grp.Location_list = append(grp.Location_list, res)
		case *r4.Measure:
			grp.Measure_list = append(grp.Measure_list, res)
		case *r4.MeasureReport:
			grp.MeasureReport_list = append(grp.MeasureReport_list, res)
		case *r4.Media:
			grp.Media_list = append(grp.Media_list, res)
		case *r4.Medication:
			grp.Medication_list = append(grp.Medication_list, res)
		case *r4.MedicationAdministration:
			grp.MedicationAdministration_list = append(grp.MedicationAdministration_list, res)
		case *r4.MedicationDispense:
			grp.MedicationDispense_list = append(grp.MedicationDispense_list, res)
		case *r4.MedicationKnowledge:
			grp.MedicationKnowledge_list = append(grp.MedicationKnowledge_list, res)
		case *r4.MedicationRequest:
			grp.MedicationRequest_list = append(grp.MedicationRequest_list, res)
		case *r4.MedicationStatement:
			grp.MedicationStatement_list = append(grp.MedicationStatement_list, res)
		case *r4.MedicinalProduct:
			grp.MedicinalProduct_list = append(grp.MedicinalProduct_list, res)
		case *r4.MedicinalProductAuthorization:
			grp.MedicinalProductAuthorization_list = append(grp.MedicinalProductAuthorization_list, res)
		case *r4.MedicinalProductContraindication:
			grp.MedicinalProductContraindication_list = append(grp.MedicinalProductContraindication_list, res)
		case *r4.MedicinalProductIndication:
			grp.MedicinalProductIndication_list = append(grp.MedicinalProductIndication_list, res)
		case *r4.MedicinalProductIngredient:
			grp.MedicinalProductIngredient_list = append(grp.MedicinalProductIngredient_list, res)
		case *r4.MedicinalProductInteraction:
			grp.MedicinalProductInteraction_list = append(grp.MedicinalProductInteraction_list, res)
		case *r4.MedicinalProductManufactured:
			grp.MedicinalProductManufactured_list = append(grp.MedicinalProductManufactured_list, res)
		case *r4.MedicinalProductPackaged:
			grp.MedicinalProductPackaged_list = append(grp.MedicinalProductPackaged_list, res)
		case *r4.MedicinalProductPharmaceutical:
			grp.MedicinalProductPharmaceutical_list = append(grp.MedicinalProductPharmaceutical_list, res)
		case *r4.MedicinalProductUndesirableEffect:
			grp.MedicinalProductUndesirableEffect_list = append(grp.MedicinalProductUndesirableEffect_list, res)
		case *r4.MessageDefinition:
			grp.MessageDefinition_list = append(grp.MessageDefinition_list, res)
		case *r4.MessageHeader:
			grp.MessageHeader_list = append(grp.MessageHeader_list, res)
		case *r4.MolecularSequence:
			grp.MolecularSequence_list = append(grp.MolecularSequence_list, res)
		case *r4.NamingSystem:
			grp.NamingSystem_list = append(grp.NamingSystem_list, res)
		case *r4.NutritionOrder:
			grp.NutritionOrder_list = append(grp.NutritionOrder_list, res)
		case *r4.Observation:
			grp.Observation_list = append(grp.Observation_list, res)
		case *r4.ObservationDefinition:
			grp.ObservationDefinition_list = append(grp.ObservationDefinition_list, res)
		case *r4.OperationDefinition:
			grp.OperationDefinition_list = append(grp.OperationDefinition_list, res)
		case *r4.OperationOutcome:
			grp.OperationOutcome_list = append(grp.OperationOutcome_list, res)
		case *r4.Organization:
			grp.Organization_list = append(grp.Organization_list, res)
		case *r4.OrganizationAffiliation:
			grp.OrganizationAffiliation_list = append(grp.OrganizationAffiliation_list, res)
		case *r4.Patient:
			grp.Patient_list = append(grp.Patient_list, res)
		case *r4.PaymentNotice:
			grp.PaymentNotice_list = append(grp.PaymentNotice_list, res)
		case *r4.PaymentReconciliation:
			grp.PaymentReconciliation_list = append(grp.PaymentReconciliation_list, res)
		case *r4.Person:
			grp.Person_list = append(grp.Person_list, res)
		case *r4.PlanDefinition:
			grp.PlanDefinition_list = append(grp.PlanDefinition_list, res)
		case *r4.Practitioner:
			grp.Practitioner_list = append(grp.Practitioner_list, res)
		case *r4.PractitionerRole:
			grp.PractitionerRole_list = append(grp.PractitionerRole_list, res)
		case *r4.Procedure:
			grp.Procedure_list = append(grp.Procedure_list, res)
		case *r4.Provenance:
			grp.Provenance_list = append(grp.Provenance_list, res)
		case *r4.Questionnaire:
			grp.Questionnaire_list = append(grp.Questionnaire_list, res)
		case *r4.QuestionnaireResponse:
			grp.QuestionnaireResponse_list = append(grp.QuestionnaireResponse_list, res)
		case *r4.RelatedPerson:
			grp.RelatedPerson_list = append(grp.RelatedPerson_list, res)
		case *r4.RequestGroup:
			grp.RequestGroup_list = append(grp.RequestGroup_list, res)
		case *r4.ResearchDefinition:
			grp.ResearchDefinition_list = append(grp.ResearchDefinition_list, res)
		case *r4.ResearchElementDefinition:
			grp.ResearchElementDefinition_list = append(grp.ResearchElementDefinition_list, res)
		case *r4.ResearchStudy:
			grp.ResearchStudy_list = append(grp.ResearchStudy_list, res)
		case *r4.ResearchSubject:
			grp.ResearchSubject_list = append(grp.ResearchSubject_list, res)
		case *r4.RiskAssessment:
			grp.RiskAssessment_list = append(grp.RiskAssessment_list, res)
		case *r4.RiskEvidenceSynthesis:
			grp.RiskEvidenceSynthesis_list = append(grp.RiskEvidenceSynthesis_list, res)
		case *r4.Schedule:
			grp.Schedule_list = append(grp.Schedule_list, res)
		case *r4.SearchParameter:
			grp.SearchParameter_list = append(grp.SearchParameter_list, res)
		case *r4.ServiceRequest:
			grp.ServiceRequest_list = append(grp.ServiceRequest_list, res)
		case *r4.Slot:
			grp.Slot_list = append(grp.Slot_list, res)
		case *r4.Specimen:
			grp.Specimen_list = append(grp.Specimen_list, res)
		case *r4.SpecimenDefinition:
			grp.SpecimenDefinition_list = append(grp.SpecimenDefinition_list, res)
		case *r4.StructureDefinition:
			grp.StructureDefinition_list = append(grp.StructureDefinition_list, res)
		case *r4.StructureMap:
			grp.StructureMap_list = append(grp.StructureMap_list, res)
		case *r4.Subscription:
			grp.Subscription_list = append(grp.Subscription_list, res)
		case *r4.Substance:
			grp.Substance_list = append(grp.Substance_list, res)
		case *r4.SubstanceNucleicAcid:
			grp.SubstanceNucleicAcid_list = append(grp.SubstanceNucleicAcid_list, res)
		case *r4.SubstancePolymer:
			grp.SubstancePolymer_list = append(grp.SubstancePolymer_list, res)
		case *r4.SubstanceProtein:
			grp.SubstanceProtein_list = append(grp.SubstanceProtein_list, res)
		case *r4.SubstanceReferenceInformation:
			grp.SubstanceReferenceInformation_list = append(grp.SubstanceReferenceInformation_list, res)
		case *r4.SubstanceSourceMaterial:
			grp.SubstanceSourceMaterial_list = append(grp.SubstanceSourceMaterial_list, res)
		case *r4.SubstanceSpecification:
			grp.SubstanceSpecification_list = append(grp.SubstanceSpecification_list, res)
		case *r4.SupplyDelivery:
			grp.SupplyDelivery_list = append(grp.SupplyDelivery_list, res)
		case *r4.SupplyRequest:
			grp.SupplyRequest_list = append(grp.SupplyRequest_list, res)
		case *r4.Task:
			grp.Task_list = append(grp.Task_list, res)
		case *r4.TerminologyCapabilities:
			grp.TerminologyCapabilities_list = append(grp.TerminologyCapabilities_list, res)
		case *r4.TestReport:
			grp.TestReport_list = append(grp.TestReport_list, res)
		case *r4.TestScript:
			grp.TestScript_list = append(grp.TestScript_list, res)
		case *r4.ValueSet:
			grp.ValueSet_list = append(grp.ValueSet_list, res)
		case *r4.VerificationResult:
			grp.VerificationResult_list = append(grp.VerificationResult_list, res)
		case *r4.VisionPrescription:
			grp.VisionPrescription_list = append(grp.VisionPrescription_list, res)
		default:
			return nil, errors.New("bundle entry not a domain resource, could not put in resourcegroup")
		}
	}
	return &grp, nil
}
