// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new VPC connection.
func (c *Client) CreateVPCConnection(ctx context.Context, params *CreateVPCConnectionInput, optFns ...func(*Options)) (*CreateVPCConnectionOutput, error) {
	if params == nil {
		params = &CreateVPCConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVPCConnection", params, optFns, c.addOperationCreateVPCConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVPCConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVPCConnectionInput struct {

	// The Amazon Web Services account ID of the account where you want to create a
	// new VPC connection.
	//
	// This member is required.
	AwsAccountId *string

	// The display name for the VPC connection.
	//
	// This member is required.
	Name *string

	// The IAM role to associate with the VPC connection.
	//
	// This member is required.
	RoleArn *string

	// A list of security group IDs for the VPC connection.
	//
	// This member is required.
	SecurityGroupIds []string

	// A list of subnet IDs for the VPC connection.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the VPC connection that you're creating. This ID is a unique
	// identifier for each Amazon Web Services Region in an Amazon Web Services
	// account.
	//
	// This member is required.
	VPCConnectionId *string

	// A list of IP addresses of DNS resolver endpoints for the VPC connection.
	DnsResolvers []string

	// A map of the key-value pairs for the resource tag or tags assigned to the VPC
	// connection.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateVPCConnectionOutput struct {

	// The Amazon Resource Name (ARN) of the VPC connection.
	Arn *string

	// The availability status of the VPC connection.
	AvailabilityStatus types.VPCConnectionAvailabilityStatus

	// The status of the creation of the VPC connection.
	CreationStatus types.VPCConnectionResourceStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The ID for the VPC connection that you're creating. This ID is unique per
	// Amazon Web Services Region for each Amazon Web Services account.
	VPCConnectionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVPCConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVPCConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVPCConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVPCConnection"); err != nil {
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
	if err = addOpCreateVPCConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVPCConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVPCConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVPCConnection",
	}
}