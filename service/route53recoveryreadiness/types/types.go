// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about a cell.
type CellOutput struct {

	// The Amazon Resource Name (ARN) for the cell.
	//
	// This member is required.
	CellArn *string

	// The name of the cell.
	//
	// This member is required.
	CellName *string

	// A list of cell ARNs.
	//
	// This member is required.
	Cells []string

	// The readiness scope for the cell, which can be a cell Amazon Resource Name
	// (ARN) or a recovery group ARN. This is a list but currently can have only one
	// element.
	//
	// This member is required.
	ParentReadinessScopes []string

	// Tags on the resources.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A component for DNS/routing control readiness checks and architecture checks.
type DNSTargetResource struct {

	// The domain name that acts as an ingress point to a portion of the customer
	// application.
	DomainName *string

	// The hosted zone Amazon Resource Name (ARN) that contains the DNS record with
	// the provided name of the target resource.
	HostedZoneArn *string

	// The Route 53 record set ID that uniquely identifies a DNS record, given a name
	// and a type.
	RecordSetId *string

	// The type of DNS record of the target resource.
	RecordType *string

	// The target resource of the DNS target resource.
	TargetResource *TargetResource

	noSmithyDocumentSerde
}

// Readiness rule information, including the resource type, rule ID, and rule
// description.
type ListRulesOutput struct {

	// The resource type that the readiness rule applies to.
	//
	// This member is required.
	ResourceType *string

	// The description of a readiness rule.
	//
	// This member is required.
	RuleDescription *string

	// The ID for the readiness rule.
	//
	// This member is required.
	RuleId *string

	noSmithyDocumentSerde
}

// Information relating to readiness check status.
type Message struct {

	// The text of a readiness check message.
	MessageText *string

	noSmithyDocumentSerde
}

// The Network Load Balancer resource that a DNS target resource points to.
type NLBResource struct {

	// The Network Load Balancer resource Amazon Resource Name (ARN).
	Arn *string

	noSmithyDocumentSerde
}

// The Route 53 resource that a DNS target resource record points to.
type R53ResourceRecord struct {

	// The DNS target domain name.
	DomainName *string

	// The Route 53 Resource Record Set ID.
	RecordSetId *string

	noSmithyDocumentSerde
}

// A readiness check.
type ReadinessCheckOutput struct {

	// The Amazon Resource Name (ARN) associated with a readiness check.
	//
	// This member is required.
	ReadinessCheckArn *string

	// Name of the resource set to be checked.
	//
	// This member is required.
	ResourceSet *string

	// Name of a readiness check.
	ReadinessCheckName *string

	// A collection of tags associated with a resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary of all readiness check statuses in a recovery group, paginated in
// GetRecoveryGroupReadinessSummary and GetCellReadinessSummary.
type ReadinessCheckSummary struct {

	// The readiness status of this readiness check.
	Readiness Readiness

	// The name of a readiness check.
	ReadinessCheckName *string

	noSmithyDocumentSerde
}

// Recommendations that are provided to make an application more recovery
// resilient.
type Recommendation struct {

	// Text of the recommendations that are provided to make an application more
	// recovery resilient.
	//
	// This member is required.
	RecommendationText *string

	noSmithyDocumentSerde
}

// A representation of the application, typically containing multiple cells.
type RecoveryGroupOutput struct {

	// A list of a cell's Amazon Resource Names (ARNs).
	//
	// This member is required.
	Cells []string

	// The Amazon Resource Name (ARN) for the recovery group.
	//
	// This member is required.
	RecoveryGroupArn *string

	// The name of the recovery group.
	//
	// This member is required.
	RecoveryGroupName *string

	// The tags associated with the recovery group.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The resource element of a resource set.
type Resource struct {

	// The component identifier of the resource, generated when DNS target resource is
	// used.
	ComponentId *string

	// The DNS target resource.
	DnsTargetResource *DNSTargetResource

	// A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this
	// resource is contained within.
	ReadinessScopes []string

	// The Amazon Resource Name (ARN) of the Amazon Web Services resource.
	ResourceArn *string

	noSmithyDocumentSerde
}

// The result of a successful Resource request, with status for an individual
// resource.
type ResourceResult struct {

	// The time (UTC) that the resource was last checked for readiness, in ISO-8601
	// format.
	//
	// This member is required.
	LastCheckedTimestamp *time.Time

	// The readiness of a resource.
	//
	// This member is required.
	Readiness Readiness

	// The component id of the resource.
	ComponentId *string

	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	noSmithyDocumentSerde
}

// A collection of resources of the same type.
type ResourceSetOutput struct {

	// The Amazon Resource Name (ARN) for the resource set.
	//
	// This member is required.
	ResourceSetArn *string

	// The name of the resource set.
	//
	// This member is required.
	ResourceSetName *string

	// The resource type of the resources in the resource set. Enter one of the
	// following values for resource type:
	//
	// AWS::ApiGateway::Stage, AWS::ApiGatewayV2::Stage,
	// AWS::AutoScaling::AutoScalingGroup, AWS::CloudWatch::Alarm,
	// AWS::EC2::CustomerGateway, AWS::DynamoDB::Table, AWS::EC2::Volume,
	// AWS::ElasticLoadBalancing::LoadBalancer,
	// AWS::ElasticLoadBalancingV2::LoadBalancer, AWS::Lambda::Function,
	// AWS::MSK::Cluster, AWS::RDS::DBCluster, AWS::Route53::HealthCheck,
	// AWS::SQS::Queue, AWS::SNS::Topic, AWS::SNS::Subscription, AWS::EC2::VPC,
	// AWS::EC2::VPNConnection, AWS::EC2::VPNGateway,
	// AWS::Route53RecoveryReadiness::DNSTargetResource
	//
	// This member is required.
	ResourceSetType *string

	// A list of resource objects.
	//
	// This member is required.
	Resources []Resource

	// A collection of tags associated with a resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The result of a successful Rule request, with status for an individual rule.
type RuleResult struct {

	// The time the resource was last checked for readiness, in ISO-8601 format, UTC.
	//
	// This member is required.
	LastCheckedTimestamp *time.Time

	// Details about the resource's readiness.
	//
	// This member is required.
	Messages []Message

	// The readiness at rule level.
	//
	// This member is required.
	Readiness Readiness

	// The identifier of the rule.
	//
	// This member is required.
	RuleId *string

	noSmithyDocumentSerde
}

// The target resource that the Route 53 record points to.
type TargetResource struct {

	// The Network Load Balancer Resource.
	NLBResource *NLBResource

	// The Route 53 resource.
	R53Resource *R53ResourceRecord

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde