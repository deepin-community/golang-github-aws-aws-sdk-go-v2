// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// There was an issue processing your request.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// We can’t process your request right now because of a server issue. Try again
// later.
type CreateChimeWebhookConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *CreateChimeWebhookConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CreateChimeWebhookConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CreateChimeWebhookConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CreateChimeWebhookConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *CreateChimeWebhookConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type CreateSlackChannelConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *CreateSlackChannelConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CreateSlackChannelConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CreateSlackChannelConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CreateSlackChannelConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *CreateSlackChannelConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type CreateTeamsChannelConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *CreateTeamsChannelConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CreateTeamsChannelConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CreateTeamsChannelConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CreateTeamsChannelConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *CreateTeamsChannelConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DeleteChimeWebhookConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DeleteChimeWebhookConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeleteChimeWebhookConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeleteChimeWebhookConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DeleteChimeWebhookConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *DeleteChimeWebhookConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DeleteMicrosoftTeamsUserIdentityException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DeleteMicrosoftTeamsUserIdentityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeleteMicrosoftTeamsUserIdentityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeleteMicrosoftTeamsUserIdentityException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DeleteMicrosoftTeamsUserIdentityException"
	}
	return *e.ErrorCodeOverride
}
func (e *DeleteMicrosoftTeamsUserIdentityException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DeleteSlackChannelConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DeleteSlackChannelConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeleteSlackChannelConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeleteSlackChannelConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DeleteSlackChannelConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *DeleteSlackChannelConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DeleteSlackUserIdentityException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DeleteSlackUserIdentityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeleteSlackUserIdentityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeleteSlackUserIdentityException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DeleteSlackUserIdentityException"
	}
	return *e.ErrorCodeOverride
}
func (e *DeleteSlackUserIdentityException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// There was an issue deleting your Slack workspace.
type DeleteSlackWorkspaceAuthorizationFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DeleteSlackWorkspaceAuthorizationFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeleteSlackWorkspaceAuthorizationFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeleteSlackWorkspaceAuthorizationFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DeleteSlackWorkspaceAuthorizationFault"
	}
	return *e.ErrorCodeOverride
}
func (e *DeleteSlackWorkspaceAuthorizationFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DeleteTeamsChannelConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DeleteTeamsChannelConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeleteTeamsChannelConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeleteTeamsChannelConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DeleteTeamsChannelConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *DeleteTeamsChannelConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DeleteTeamsConfiguredTeamException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DeleteTeamsConfiguredTeamException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeleteTeamsConfiguredTeamException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeleteTeamsConfiguredTeamException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DeleteTeamsConfiguredTeamException"
	}
	return *e.ErrorCodeOverride
}
func (e *DeleteTeamsConfiguredTeamException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DescribeChimeWebhookConfigurationsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DescribeChimeWebhookConfigurationsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DescribeChimeWebhookConfigurationsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DescribeChimeWebhookConfigurationsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DescribeChimeWebhookConfigurationsException"
	}
	return *e.ErrorCodeOverride
}
func (e *DescribeChimeWebhookConfigurationsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DescribeSlackChannelConfigurationsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DescribeSlackChannelConfigurationsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DescribeSlackChannelConfigurationsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DescribeSlackChannelConfigurationsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DescribeSlackChannelConfigurationsException"
	}
	return *e.ErrorCodeOverride
}
func (e *DescribeSlackChannelConfigurationsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DescribeSlackUserIdentitiesException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DescribeSlackUserIdentitiesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DescribeSlackUserIdentitiesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DescribeSlackUserIdentitiesException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DescribeSlackUserIdentitiesException"
	}
	return *e.ErrorCodeOverride
}
func (e *DescribeSlackUserIdentitiesException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type DescribeSlackWorkspacesException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DescribeSlackWorkspacesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DescribeSlackWorkspacesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DescribeSlackWorkspacesException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DescribeSlackWorkspacesException"
	}
	return *e.ErrorCodeOverride
}
func (e *DescribeSlackWorkspacesException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// We can’t process your request right now because of a server issue. Try again
// later.
type GetAccountPreferencesException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *GetAccountPreferencesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GetAccountPreferencesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GetAccountPreferencesException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "GetAccountPreferencesException"
	}
	return *e.ErrorCodeOverride
}
func (e *GetAccountPreferencesException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// We can’t process your request right now because of a server issue. Try again
// later.
type GetTeamsChannelConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *GetTeamsChannelConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GetTeamsChannelConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GetTeamsChannelConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "GetTeamsChannelConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *GetTeamsChannelConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// Customer/consumer-facing internal service exception.
// https://w.amazon.com/index.php/AWS/API_Standards/Exceptions#InternalServiceError
type InternalServiceError struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServiceError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceError) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServiceError"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServiceError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Your request input doesn't meet the constraints that AWS Chatbot requires.
type InvalidParameterException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your request input doesn't meet the constraints that AWS Chatbot requires.
type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded a service limit for AWS Chatbot.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// We can’t process your request right now because of a server issue. Try again
// later.
type ListMicrosoftTeamsConfiguredTeamsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ListMicrosoftTeamsConfiguredTeamsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ListMicrosoftTeamsConfiguredTeamsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ListMicrosoftTeamsConfiguredTeamsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ListMicrosoftTeamsConfiguredTeamsException"
	}
	return *e.ErrorCodeOverride
}
func (e *ListMicrosoftTeamsConfiguredTeamsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type ListMicrosoftTeamsUserIdentitiesException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ListMicrosoftTeamsUserIdentitiesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ListMicrosoftTeamsUserIdentitiesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ListMicrosoftTeamsUserIdentitiesException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ListMicrosoftTeamsUserIdentitiesException"
	}
	return *e.ErrorCodeOverride
}
func (e *ListMicrosoftTeamsUserIdentitiesException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type ListTeamsChannelConfigurationsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ListTeamsChannelConfigurationsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ListTeamsChannelConfigurationsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ListTeamsChannelConfigurationsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ListTeamsChannelConfigurationsException"
	}
	return *e.ErrorCodeOverride
}
func (e *ListTeamsChannelConfigurationsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We were not able to find the resource for your request.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// We can’t process your request right now because of a server issue. Try again
// later.
type ServiceUnavailableException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The supplied list of tags contains too many tags.
type TooManyTagsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyTagsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// We can’t process your request right now because of a server issue. Try again
// later.
type UpdateAccountPreferencesException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UpdateAccountPreferencesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UpdateAccountPreferencesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UpdateAccountPreferencesException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UpdateAccountPreferencesException"
	}
	return *e.ErrorCodeOverride
}
func (e *UpdateAccountPreferencesException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// We can’t process your request right now because of a server issue. Try again
// later.
type UpdateChimeWebhookConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UpdateChimeWebhookConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UpdateChimeWebhookConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UpdateChimeWebhookConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UpdateChimeWebhookConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *UpdateChimeWebhookConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type UpdateSlackChannelConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UpdateSlackChannelConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UpdateSlackChannelConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UpdateSlackChannelConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UpdateSlackChannelConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *UpdateSlackChannelConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// We can’t process your request right now because of a server issue. Try again
// later.
type UpdateTeamsChannelConfigurationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UpdateTeamsChannelConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UpdateTeamsChannelConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UpdateTeamsChannelConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UpdateTeamsChannelConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *UpdateTeamsChannelConfigurationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}