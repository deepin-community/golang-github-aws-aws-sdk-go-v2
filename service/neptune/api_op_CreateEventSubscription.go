// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an event notification subscription. This action requires a topic ARN
// (Amazon Resource Name) created by either the Neptune console, the SNS console,
// or the SNS API. To obtain an ARN with SNS, you must create a topic in Amazon SNS
// and subscribe to the topic. The ARN is displayed in the SNS console.
//
// You can specify the type of source (SourceType) you want to be notified of,
// provide a list of Neptune sources (SourceIds) that triggers the events, and
// provide a list of event categories (EventCategories) for events you want to be
// notified of. For example, you can specify SourceType = db-instance, SourceIds =
// mydbinstance1, mydbinstance2 and EventCategories = Availability, Backup.
//
// If you specify both the SourceType and SourceIds, such as SourceType =
// db-instance and SourceIdentifier = myDBInstance1, you are notified of all the
// db-instance events for the specified source. If you specify a SourceType but do
// not specify a SourceIdentifier, you receive notice of the events for that source
// type for all your Neptune sources. If you do not specify either the SourceType
// nor the SourceIdentifier, you are notified of events generated from all Neptune
// sources belonging to your customer account.
func (c *Client) CreateEventSubscription(ctx context.Context, params *CreateEventSubscriptionInput, optFns ...func(*Options)) (*CreateEventSubscriptionOutput, error) {
	if params == nil {
		params = &CreateEventSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventSubscription", params, optFns, c.addOperationCreateEventSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventSubscriptionInput struct {

	// The Amazon Resource Name (ARN) of the SNS topic created for event notification.
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	//
	// This member is required.
	SnsTopicArn *string

	// The name of the subscription.
	//
	// Constraints: The name must be less than 255 characters.
	//
	// This member is required.
	SubscriptionName *string

	//  A Boolean value; set to true to activate the subscription, set to false to
	// create the subscription but not active it.
	Enabled *bool

	//  A list of event categories for a SourceType that you want to subscribe to. You
	// can see a list of the categories for a given SourceType by using the
	// DescribeEventCategories action.
	EventCategories []string

	// The list of identifiers of the event sources for which events are returned. If
	// not specified, then all sources are included in the response. An identifier must
	// begin with a letter and must contain only ASCII letters, digits, and hyphens; it
	// can't end with a hyphen or contain two consecutive hyphens.
	//
	// Constraints:
	//
	//   - If SourceIds are supplied, SourceType must also be provided.
	//
	//   - If the source type is a DB instance, then a DBInstanceIdentifier must be
	//   supplied.
	//
	//   - If the source type is a DB security group, a DBSecurityGroupName must be
	//   supplied.
	//
	//   - If the source type is a DB parameter group, a DBParameterGroupName must be
	//   supplied.
	//
	//   - If the source type is a DB snapshot, a DBSnapshotIdentifier must be supplied.
	SourceIds []string

	// The type of source that is generating the events. For example, if you want to
	// be notified of events generated by a DB instance, you would set this parameter
	// to db-instance. if this value is not specified, all events are returned.
	//
	// Valid values: db-instance | db-cluster | db-parameter-group | db-security-group
	// | db-snapshot | db-cluster-snapshot
	SourceType *string

	// The tags to be applied to the new event subscription.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEventSubscriptionOutput struct {

	// Contains the results of a successful invocation of the DescribeEventSubscriptions action.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateEventSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateEventSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEventSubscription"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEventSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventSubscription(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateEventSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEventSubscription",
	}
}