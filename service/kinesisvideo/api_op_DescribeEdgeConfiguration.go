// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a stream’s edge configuration that was set using the
// StartEdgeConfigurationUpdate API and the latest status of the edge agent's
// recorder and uploader jobs. Use this API to get the status of the configuration
// to determine if the configuration is in sync with the Edge Agent. Use this API
// to evaluate the health of the Edge Agent.
func (c *Client) DescribeEdgeConfiguration(ctx context.Context, params *DescribeEdgeConfigurationInput, optFns ...func(*Options)) (*DescribeEdgeConfigurationOutput, error) {
	if params == nil {
		params = &DescribeEdgeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEdgeConfiguration", params, optFns, c.addOperationDescribeEdgeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEdgeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEdgeConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the stream. Specify either the StreamName or
	// the StreamARN .
	StreamARN *string

	// The name of the stream whose edge configuration you want to update. Specify
	// either the StreamName or the StreamARN .
	StreamName *string

	noSmithyDocumentSerde
}

type DescribeEdgeConfigurationOutput struct {

	// The timestamp at which a stream’s edge configuration was first created.
	CreationTime *time.Time

	// An object that contains the latest status details for an edge agent's recorder
	// and uploader jobs. Use this information to determine the current health of an
	// edge agent.
	EdgeAgentStatus *types.EdgeAgentStatus

	// A description of the stream's edge configuration that will be used to sync with
	// the Edge Agent IoT Greengrass component. The Edge Agent component will run on an
	// IoT Hub Device setup at your premise.
	EdgeConfig *types.EdgeConfig

	// A description of the generated failure status.
	FailedStatusDetails *string

	// The timestamp at which a stream’s edge configuration was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resource Name (ARN) of the stream.
	StreamARN *string

	// The name of the stream from which the edge configuration was updated.
	StreamName *string

	// The latest status of the edge configuration update.
	SyncStatus types.SyncStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEdgeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeEdgeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeEdgeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEdgeConfiguration"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEdgeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEdgeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEdgeConfiguration",
	}
}