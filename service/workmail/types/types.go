// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A rule that controls access to an WorkMail organization.
type AccessControlRule struct {

	// Access protocol actions to include in the rule. Valid values include ActiveSync
	// , AutoDiscover , EWS , IMAP , SMTP , WindowsOutlook , and WebMail .
	Actions []string

	// The date that the rule was created.
	DateCreated *time.Time

	// The date that the rule was modified.
	DateModified *time.Time

	// The rule description.
	Description *string

	// The rule effect.
	Effect AccessControlRuleEffect

	// Impersonation role IDs to include in the rule.
	ImpersonationRoleIds []string

	// IPv4 CIDR ranges to include in the rule.
	IpRanges []string

	// The rule name.
	Name *string

	// Access protocol actions to exclude from the rule. Valid values include
	// ActiveSync , AutoDiscover , EWS , IMAP , SMTP , WindowsOutlook , and WebMail .
	NotActions []string

	// Impersonation role IDs to exclude from the rule.
	NotImpersonationRoleIds []string

	// IPv4 CIDR ranges to exclude from the rule.
	NotIpRanges []string

	// User IDs to exclude from the rule.
	NotUserIds []string

	// User IDs to include in the rule.
	UserIds []string

	noSmithyDocumentSerde
}

// List all the AvailabilityConfiguration 's for the given WorkMail organization.
type AvailabilityConfiguration struct {

	// The date and time at which the availability configuration was created.
	DateCreated *time.Time

	// The date and time at which the availability configuration was last modified.
	DateModified *time.Time

	// Displays the domain to which the provider applies.
	DomainName *string

	// If ProviderType is EWS , then this field contains
	// RedactedEwsAvailabilityProvider . Otherwise, it is not required.
	EwsProvider *RedactedEwsAvailabilityProvider

	// If ProviderType is LAMBDA then this field contains LambdaAvailabilityProvider .
	// Otherwise, it is not required.
	LambdaProvider *LambdaAvailabilityProvider

	// Displays the provider type that applies to this domain.
	ProviderType AvailabilityProviderType

	noSmithyDocumentSerde
}

// At least one delegate must be associated to the resource to disable automatic
// replies from the resource.
type BookingOptions struct {

	// The resource's ability to automatically reply to requests. If disabled,
	// delegates must be associated to the resource.
	AutoAcceptRequests bool

	// The resource's ability to automatically decline any conflicting requests.
	AutoDeclineConflictingRequests bool

	// The resource's ability to automatically decline any recurring requests.
	AutoDeclineRecurringRequests bool

	noSmithyDocumentSerde
}

// The name of the attribute, which is one of the values defined in the
// UserAttribute enumeration.
type Delegate struct {

	// The identifier for the user or group associated as the resource's delegate.
	//
	// This member is required.
	Id *string

	// The type of the delegate: user or group.
	//
	// This member is required.
	Type MemberType

	noSmithyDocumentSerde
}

// A DNS record uploaded to your DNS provider.
type DnsRecord struct {

	// The DNS hostname.- For example, domain.example.com .
	Hostname *string

	// The RFC 1035 record type. Possible values: CNAME , A , MX .
	Type *string

	// The value returned by the DNS for a query to that hostname and record type.
	Value *string

	noSmithyDocumentSerde
}

// The domain to associate with an WorkMail organization.
//
// When you configure a domain hosted in Amazon Route 53 (Route 53), all
// recommended DNS records are added to the organization when you create it. For
// more information, see [Adding a domain]in the WorkMail Administrator Guide.
//
// [Adding a domain]: https://docs.aws.amazon.com/workmail/latest/adminguide/add_domain.html
type Domain struct {

	// The fully qualified domain name.
	//
	// This member is required.
	DomainName *string

	// The hosted zone ID for a domain hosted in Route 53. Required when configuring a
	// domain hosted in Route 53.
	HostedZoneId *string

	noSmithyDocumentSerde
}

// Describes an EWS based availability provider. This is only used as input to the
// service.
type EwsAvailabilityProvider struct {

	// The endpoint of the remote EWS server.
	//
	// This member is required.
	EwsEndpoint *string

	// The password used to authenticate the remote EWS server.
	//
	// This member is required.
	EwsPassword *string

	// The username used to authenticate the remote EWS server.
	//
	// This member is required.
	EwsUsername *string

	noSmithyDocumentSerde
}

// The configuration applied to an organization's folders by its retention policy.
type FolderConfiguration struct {

	// The action to take on the folder contents at the end of the folder
	// configuration period.
	//
	// This member is required.
	Action RetentionAction

	// The folder name.
	//
	// This member is required.
	Name FolderName

	// The number of days for which the folder-configuration action applies.
	Period *int32

	noSmithyDocumentSerde
}

// The representation of an WorkMail group.
type Group struct {

	// The date indicating when the group was disabled from WorkMail use.
	DisabledDate *time.Time

	// The email of the group.
	Email *string

	// The date indicating when the group was enabled for WorkMail use.
	EnabledDate *time.Time

	// The identifier of the group.
	Id *string

	// The name of the group.
	Name *string

	// The state of the group, which can be ENABLED, DISABLED, or DELETED.
	State EntityState

	noSmithyDocumentSerde
}

// The identifier that contains the Group ID and name of a group.
type GroupIdentifier struct {

	// Group ID that matched the group.
	GroupId *string

	// Group name that matched the group.
	GroupName *string

	noSmithyDocumentSerde
}

// The impersonation rule that matched the input.
type ImpersonationMatchedRule struct {

	// The ID of the rule that matched the input
	ImpersonationRuleId *string

	// The name of the rule that matched the input.
	Name *string

	noSmithyDocumentSerde
}

// An impersonation role for the given WorkMail organization.
type ImpersonationRole struct {

	// The date when the impersonation role was created.
	DateCreated *time.Time

	// The date when the impersonation role was last modified.
	DateModified *time.Time

	// The identifier of the impersonation role.
	ImpersonationRoleId *string

	// The impersonation role name.
	Name *string

	// The impersonation role type.
	Type ImpersonationRoleType

	noSmithyDocumentSerde
}

// The rules for the given impersonation role.
type ImpersonationRule struct {

	// The effect of the rule when it matches the input. Allowed effect values are
	// ALLOW or DENY .
	//
	// This member is required.
	Effect AccessEffect

	// The identifier of the rule.
	//
	// This member is required.
	ImpersonationRuleId *string

	// The rule description.
	Description *string

	// The rule name.
	Name *string

	// A list of user IDs that don't match the rule.
	NotTargetUsers []string

	// A list of user IDs that match the rule.
	TargetUsers []string

	noSmithyDocumentSerde
}

// Describes a Lambda based availability provider.
type LambdaAvailabilityProvider struct {

	// The Amazon Resource Name (ARN) of the Lambda that acts as the availability
	// provider.
	//
	// This member is required.
	LambdaArn *string

	noSmithyDocumentSerde
}

//	Filtering options for ListGroups operation. This is only used as input to
//
// Operation.
type ListGroupsFilters struct {

	// Filters only groups with the provided name prefix.
	NamePrefix *string

	// Filters only groups with the provided primary email prefix.
	PrimaryEmailPrefix *string

	// Filters only groups with the provided state.
	State EntityState

	noSmithyDocumentSerde
}

//	Filtering options for ListGroupsForEntity operation. This is only used as
//
// input to Operation.
type ListGroupsForEntityFilters struct {

	// Filters only group names that start with the provided name prefix.
	GroupNamePrefix *string

	noSmithyDocumentSerde
}

// Filtering options for ListResources operation. This is only used as input to
// Operation.
type ListResourcesFilters struct {

	// Filters only resource that start with the entered name prefix .
	NamePrefix *string

	// Filters only resource with the provided primary email prefix.
	PrimaryEmailPrefix *string

	// Filters only resource with the provided state.
	State EntityState

	noSmithyDocumentSerde
}

//	Filtering options for ListUsers operation. This is only used as input to
//
// Operation.
type ListUsersFilters struct {

	// Filters only users with the provided display name prefix.
	DisplayNamePrefix *string

	// Filters only users with the provided email prefix.
	PrimaryEmailPrefix *string

	// Filters only users with the provided state.
	State EntityState

	// Filters only users with the provided username prefix.
	UsernamePrefix *string

	noSmithyDocumentSerde
}

// The details of a mailbox export job, including the user or resource ID
// associated with the mailbox and the S3 bucket that the mailbox contents are
// exported to.
type MailboxExportJob struct {

	// The mailbox export job description.
	Description *string

	// The mailbox export job end timestamp.
	EndTime *time.Time

	// The identifier of the user or resource associated with the mailbox.
	EntityId *string

	// The estimated progress of the mailbox export job, in percentage points.
	EstimatedProgress int32

	// The identifier of the mailbox export job.
	JobId *string

	// The name of the S3 bucket.
	S3BucketName *string

	// The path to the S3 bucket and file that the mailbox export job exports to.
	S3Path *string

	// The mailbox export job start timestamp.
	StartTime *time.Time

	// The state of the mailbox export job.
	State MailboxExportJobState

	noSmithyDocumentSerde
}

// The data for a given domain.
type MailDomainSummary struct {

	// Whether the domain is default or not.
	DefaultDomain bool

	// The domain name.
	DomainName *string

	noSmithyDocumentSerde
}

// The representation of a user or group.
type Member struct {

	// The date indicating when the member was disabled from WorkMail use.
	DisabledDate *time.Time

	// The date indicating when the member was enabled for WorkMail use.
	EnabledDate *time.Time

	// The identifier of the member.
	Id *string

	// The name of the member.
	Name *string

	// The state of the member, which can be ENABLED, DISABLED, or DELETED.
	State EntityState

	// A member can be a user or group.
	Type MemberType

	noSmithyDocumentSerde
}

// The rule that a simulated user matches.
type MobileDeviceAccessMatchedRule struct {

	// Identifier of the rule that a simulated user matches.
	MobileDeviceAccessRuleId *string

	// Name of a rule that a simulated user matches.
	Name *string

	noSmithyDocumentSerde
}

// The override object.
type MobileDeviceAccessOverride struct {

	// The date the override was first created.
	DateCreated *time.Time

	// The date the override was last modified.
	DateModified *time.Time

	// A description of the override.
	Description *string

	// The device to which the override applies.
	DeviceId *string

	// The effect of the override, ALLOW or DENY .
	Effect MobileDeviceAccessRuleEffect

	// The WorkMail user to which the access override applies.
	UserId *string

	noSmithyDocumentSerde
}

// A rule that controls access to mobile devices for an WorkMail group.
type MobileDeviceAccessRule struct {

	// The date and time at which an access rule was created.
	DateCreated *time.Time

	// The date and time at which an access rule was modified.
	DateModified *time.Time

	// The description of a mobile access rule.
	Description *string

	// Device models that a rule will match.
	DeviceModels []string

	// Device operating systems that a rule will match.
	DeviceOperatingSystems []string

	// Device types that a rule will match.
	DeviceTypes []string

	// Device user agents that a rule will match.
	DeviceUserAgents []string

	// The effect of the rule when it matches. Allowed values are ALLOW or DENY .
	Effect MobileDeviceAccessRuleEffect

	// The ID assigned to a mobile access rule.
	MobileDeviceAccessRuleId *string

	// The name of a mobile access rule.
	Name *string

	// Device models that a rule will not match. All other device models will match.
	NotDeviceModels []string

	// Device operating systems that a rule will not match. All other device types
	// will match.
	NotDeviceOperatingSystems []string

	// Device types that a rule will not match. All other device types will match.
	NotDeviceTypes []string

	// Device user agents that a rule will not match. All other device user agents
	// will match.
	NotDeviceUserAgents []string

	noSmithyDocumentSerde
}

// The representation of an organization.
type OrganizationSummary struct {

	// The alias associated with the organization.
	Alias *string

	// The default email domain associated with the organization.
	DefaultMailDomain *string

	// The error message associated with the organization. It is only present if
	// unexpected behavior has occurred with regards to the organization. It provides
	// insight or solutions regarding unexpected behavior.
	ErrorMessage *string

	// The identifier associated with the organization.
	OrganizationId *string

	// The state associated with the organization.
	State *string

	noSmithyDocumentSerde
}

// Permission granted to a user, group, or resource to access a certain aspect of
// another user, group, or resource mailbox.
type Permission struct {

	// The identifier of the user, group, or resource to which the permissions are
	// granted.
	//
	// This member is required.
	GranteeId *string

	// The type of user, group, or resource referred to in GranteeId.
	//
	// This member is required.
	GranteeType MemberType

	// The permissions granted to the grantee. SEND_AS allows the grantee to send
	// email as the owner of the mailbox (the grantee is not mentioned on these
	// emails). SEND_ON_BEHALF allows the grantee to send email on behalf of the owner
	// of the mailbox (the grantee is not mentioned as the physical sender of these
	// emails). FULL_ACCESS allows the grantee full access to the mailbox, irrespective
	// of other folder-level permissions set on the mailbox.
	//
	// This member is required.
	PermissionValues []PermissionType

	noSmithyDocumentSerde
}

// Describes an EWS based availability provider when returned from the service. It
// does not contain the password of the endpoint.
type RedactedEwsAvailabilityProvider struct {

	// The endpoint of the remote EWS server.
	EwsEndpoint *string

	// The username used to authenticate the remote EWS server.
	EwsUsername *string

	noSmithyDocumentSerde
}

// The representation of a resource.
type Resource struct {

	// Resource description.
	Description *string

	// The date indicating when the resource was disabled from WorkMail use.
	DisabledDate *time.Time

	// The email of the resource.
	Email *string

	// The date indicating when the resource was enabled for WorkMail use.
	EnabledDate *time.Time

	// The identifier of the resource.
	Id *string

	// The name of the resource.
	Name *string

	// The state of the resource, which can be ENABLED, DISABLED, or DELETED.
	State EntityState

	// The type of the resource: equipment or room.
	Type ResourceType

	noSmithyDocumentSerde
}

// Describes a tag applied to a resource.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The representation of an WorkMail user.
type User struct {

	// The date indicating when the user was disabled from WorkMail use.
	DisabledDate *time.Time

	// The display name of the user.
	DisplayName *string

	// The email of the user.
	Email *string

	// The date indicating when the user was enabled for WorkMail use.
	EnabledDate *time.Time

	// The identifier of the user.
	Id *string

	// The name of the user.
	Name *string

	// The state of the user, which can be ENABLED, DISABLED, or DELETED.
	State EntityState

	// The role of the user.
	UserRole UserRole

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde