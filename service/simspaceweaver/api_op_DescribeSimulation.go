// Code generated by smithy-go-codegen DO NOT EDIT.

package simspaceweaver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/simspaceweaver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the current state of the given simulation.
func (c *Client) DescribeSimulation(ctx context.Context, params *DescribeSimulationInput, optFns ...func(*Options)) (*DescribeSimulationOutput, error) {
	if params == nil {
		params = &DescribeSimulationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSimulation", params, optFns, c.addOperationDescribeSimulationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSimulationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSimulationInput struct {

	// The name of the simulation.
	//
	// This member is required.
	Simulation *string

	noSmithyDocumentSerde
}

type DescribeSimulationOutput struct {

	// The Amazon Resource Name (ARN) of the simulation. For more information about
	// ARNs, see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	Arn *string

	// The time when the simulation was created, expressed as the number of seconds
	// and milliseconds in UTC since the Unix epoch (0:0:0.000, January 1, 1970).
	CreationTime *time.Time

	// The description of the simulation.
	Description *string

	// A universally unique identifier (UUID) for this simulation.
	ExecutionId *string

	// A collection of additional state information, such as domain and clock
	// configuration.
	LiveSimulationState *types.LiveSimulationState

	// Settings that control how SimSpace Weaver handles your simulation log data.
	LoggingConfiguration *types.LoggingConfiguration

	// The maximum running time of the simulation, specified as a number of minutes (m
	// or M), hours (h or H), or days (d or D). The simulation stops when it reaches
	// this limit. The maximum value is 14D , or its equivalent in the other units. The
	// default value is 14D . A value equivalent to 0 makes the simulation immediately
	// transition to Stopping as soon as it reaches Started .
	MaximumDuration *string

	// The name of the simulation.
	Name *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that the simulation assumes to perform actions. For more information about ARNs,
	// see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference. For more information about
	// IAM roles, see [IAM roles]in the Identity and Access Management User Guide.
	//
	// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	RoleArn *string

	// An error message that SimSpace Weaver returns only if there is a problem with
	// the simulation schema.
	//
	// Deprecated: SchemaError is no longer used, check StartError instead.
	SchemaError *string

	// The location of the simulation schema in Amazon Simple Storage Service (Amazon
	// S3). For more information about Amazon S3, see the [Amazon Simple Storage Service User Guide].
	//
	// [Amazon Simple Storage Service User Guide]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html
	SchemaS3Location *types.S3Location

	// A location in Amazon Simple Storage Service (Amazon S3) where SimSpace Weaver
	// stores simulation data, such as your app .zip files and schema file. For more
	// information about Amazon S3, see the [Amazon Simple Storage Service User Guide].
	//
	// [Amazon Simple Storage Service User Guide]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html
	SnapshotS3Location *types.S3Location

	// An error message that SimSpace Weaver returns only if a problem occurs when the
	// simulation is in the STARTING state.
	StartError *string

	// The current lifecycle state of the simulation.
	Status types.SimulationStatus

	// The desired lifecycle state of the simulation.
	TargetStatus types.SimulationTargetStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSimulationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSimulation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSimulation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSimulation"); err != nil {
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
	if err = addOpDescribeSimulationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSimulation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSimulation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSimulation",
	}
}