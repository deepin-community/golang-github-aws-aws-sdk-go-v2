// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ExclusionStatus string

// Enum values for ExclusionStatus
const (
	ExclusionStatusExcluded ExclusionStatus = "excluded"
	ExclusionStatusIncluded ExclusionStatus = "included"
)

// Values returns all known values for ExclusionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExclusionStatus) Values() []ExclusionStatus {
	return []ExclusionStatus{
		"excluded",
		"included",
	}
}

type RecommendationLanguage string

// Enum values for RecommendationLanguage
const (
	RecommendationLanguageEnglish             RecommendationLanguage = "en"
	RecommendationLanguageJapanese            RecommendationLanguage = "ja"
	RecommendationLanguageChinese             RecommendationLanguage = "zh"
	RecommendationLanguageFrench              RecommendationLanguage = "fr"
	RecommendationLanguageGerman              RecommendationLanguage = "de"
	RecommendationLanguageKorean              RecommendationLanguage = "ko"
	RecommendationLanguageTraditionalChinese  RecommendationLanguage = "zh_TW"
	RecommendationLanguageItalian             RecommendationLanguage = "it"
	RecommendationLanguageSpanish             RecommendationLanguage = "es"
	RecommendationLanguageBrazilianPortuguese RecommendationLanguage = "pt_BR"
	RecommendationLanguageBahasaIndonesia     RecommendationLanguage = "id"
)

// Values returns all known values for RecommendationLanguage. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationLanguage) Values() []RecommendationLanguage {
	return []RecommendationLanguage{
		"en",
		"ja",
		"zh",
		"fr",
		"de",
		"ko",
		"zh_TW",
		"it",
		"es",
		"pt_BR",
		"id",
	}
}

type RecommendationLifecycleStage string

// Enum values for RecommendationLifecycleStage
const (
	RecommendationLifecycleStageInProgress      RecommendationLifecycleStage = "in_progress"
	RecommendationLifecycleStagePendingResponse RecommendationLifecycleStage = "pending_response"
	RecommendationLifecycleStageDismissed       RecommendationLifecycleStage = "dismissed"
	RecommendationLifecycleStageResolved        RecommendationLifecycleStage = "resolved"
)

// Values returns all known values for RecommendationLifecycleStage. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationLifecycleStage) Values() []RecommendationLifecycleStage {
	return []RecommendationLifecycleStage{
		"in_progress",
		"pending_response",
		"dismissed",
		"resolved",
	}
}

type RecommendationPillar string

// Enum values for RecommendationPillar
const (
	RecommendationPillarCostOptimizing        RecommendationPillar = "cost_optimizing"
	RecommendationPillarPerformance           RecommendationPillar = "performance"
	RecommendationPillarSecurity              RecommendationPillar = "security"
	RecommendationPillarServiceLimits         RecommendationPillar = "service_limits"
	RecommendationPillarFaultTolerance        RecommendationPillar = "fault_tolerance"
	RecommendationPillarOperationalExcellence RecommendationPillar = "operational_excellence"
)

// Values returns all known values for RecommendationPillar. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationPillar) Values() []RecommendationPillar {
	return []RecommendationPillar{
		"cost_optimizing",
		"performance",
		"security",
		"service_limits",
		"fault_tolerance",
		"operational_excellence",
	}
}

type RecommendationSource string

// Enum values for RecommendationSource
const (
	RecommendationSourceAwsConfig        RecommendationSource = "aws_config"
	RecommendationSourceComputeOptimizer RecommendationSource = "compute_optimizer"
	RecommendationSourceCostExplorer     RecommendationSource = "cost_explorer"
	RecommendationSourceLse              RecommendationSource = "lse"
	RecommendationSourceManual           RecommendationSource = "manual"
	RecommendationSourcePse              RecommendationSource = "pse"
	RecommendationSourceRds              RecommendationSource = "rds"
	RecommendationSourceResilience       RecommendationSource = "resilience"
	RecommendationSourceResilienceHub    RecommendationSource = "resilience_hub"
	RecommendationSourceSecurityHub      RecommendationSource = "security_hub"
	RecommendationSourceStir             RecommendationSource = "stir"
	RecommendationSourceTaCheck          RecommendationSource = "ta_check"
	RecommendationSourceWellArchitected  RecommendationSource = "well_architected"
)

// Values returns all known values for RecommendationSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationSource) Values() []RecommendationSource {
	return []RecommendationSource{
		"aws_config",
		"compute_optimizer",
		"cost_explorer",
		"lse",
		"manual",
		"pse",
		"rds",
		"resilience",
		"resilience_hub",
		"security_hub",
		"stir",
		"ta_check",
		"well_architected",
	}
}

type RecommendationStatus string

// Enum values for RecommendationStatus
const (
	RecommendationStatusOk      RecommendationStatus = "ok"
	RecommendationStatusWarning RecommendationStatus = "warning"
	RecommendationStatusError   RecommendationStatus = "error"
)

// Values returns all known values for RecommendationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationStatus) Values() []RecommendationStatus {
	return []RecommendationStatus{
		"ok",
		"warning",
		"error",
	}
}

type RecommendationType string

// Enum values for RecommendationType
const (
	RecommendationTypeStandard RecommendationType = "standard"
	RecommendationTypePriority RecommendationType = "priority"
)

// Values returns all known values for RecommendationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationType) Values() []RecommendationType {
	return []RecommendationType{
		"standard",
		"priority",
	}
}

type ResourceStatus string

// Enum values for ResourceStatus
const (
	ResourceStatusOk      ResourceStatus = "ok"
	ResourceStatusWarning ResourceStatus = "warning"
	ResourceStatusError   ResourceStatus = "error"
)

// Values returns all known values for ResourceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceStatus) Values() []ResourceStatus {
	return []ResourceStatus{
		"ok",
		"warning",
		"error",
	}
}

type UpdateRecommendationLifecycleStage string

// Enum values for UpdateRecommendationLifecycleStage
const (
	UpdateRecommendationLifecycleStagePendingResponse UpdateRecommendationLifecycleStage = "pending_response"
	UpdateRecommendationLifecycleStageInProgress      UpdateRecommendationLifecycleStage = "in_progress"
	UpdateRecommendationLifecycleStageDismissed       UpdateRecommendationLifecycleStage = "dismissed"
	UpdateRecommendationLifecycleStageResolved        UpdateRecommendationLifecycleStage = "resolved"
)

// Values returns all known values for UpdateRecommendationLifecycleStage. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UpdateRecommendationLifecycleStage) Values() []UpdateRecommendationLifecycleStage {
	return []UpdateRecommendationLifecycleStage{
		"pending_response",
		"in_progress",
		"dismissed",
		"resolved",
	}
}

type UpdateRecommendationLifecycleStageReasonCode string

// Enum values for UpdateRecommendationLifecycleStageReasonCode
const (
	UpdateRecommendationLifecycleStageReasonCodeNonCriticalAccount    UpdateRecommendationLifecycleStageReasonCode = "non_critical_account"
	UpdateRecommendationLifecycleStageReasonCodeTemporaryAccount      UpdateRecommendationLifecycleStageReasonCode = "temporary_account"
	UpdateRecommendationLifecycleStageReasonCodeValidBusinessCase     UpdateRecommendationLifecycleStageReasonCode = "valid_business_case"
	UpdateRecommendationLifecycleStageReasonCodeOtherMethodsAvailable UpdateRecommendationLifecycleStageReasonCode = "other_methods_available"
	UpdateRecommendationLifecycleStageReasonCodeLowPriority           UpdateRecommendationLifecycleStageReasonCode = "low_priority"
	UpdateRecommendationLifecycleStageReasonCodeNotApplicable         UpdateRecommendationLifecycleStageReasonCode = "not_applicable"
	UpdateRecommendationLifecycleStageReasonCodeOther                 UpdateRecommendationLifecycleStageReasonCode = "other"
)

// Values returns all known values for
// UpdateRecommendationLifecycleStageReasonCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UpdateRecommendationLifecycleStageReasonCode) Values() []UpdateRecommendationLifecycleStageReasonCode {
	return []UpdateRecommendationLifecycleStageReasonCode{
		"non_critical_account",
		"temporary_account",
		"valid_business_case",
		"other_methods_available",
		"low_priority",
		"not_applicable",
		"other",
	}
}