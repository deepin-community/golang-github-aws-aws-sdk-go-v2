// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds an [InputProcessingConfiguration] to an application. An input processor preprocesses records on the
// input stream before the application's SQL code executes. Currently, the only
// input processor available is [AWS Lambda].
//
// [AWS Lambda]: https://docs.aws.amazon.com/lambda/
// [InputProcessingConfiguration]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html
func (c *Client) AddApplicationInputProcessingConfiguration(ctx context.Context, params *AddApplicationInputProcessingConfigurationInput, optFns ...func(*Options)) (*AddApplicationInputProcessingConfigurationOutput, error) {
	if params == nil {
		params = &AddApplicationInputProcessingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationInputProcessingConfiguration", params, optFns, c.addOperationAddApplicationInputProcessingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationInputProcessingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationInputProcessingConfigurationInput struct {

	// Name of the application to which you want to add the input processing
	// configuration.
	//
	// This member is required.
	ApplicationName *string

	// Version of the application to which you want to add the input processing
	// configuration. You can use the [DescribeApplication]operation to get the current application
	// version. If the version specified is not the current version, the
	// ConcurrentModificationException is returned.
	//
	// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The ID of the input configuration to add the input processing configuration to.
	// You can get a list of the input IDs for an application using the [DescribeApplication]operation.
	//
	// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
	//
	// This member is required.
	InputId *string

	// The [InputProcessingConfiguration] to add to the application.
	//
	// [InputProcessingConfiguration]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html
	//
	// This member is required.
	InputProcessingConfiguration *types.InputProcessingConfiguration

	noSmithyDocumentSerde
}

type AddApplicationInputProcessingConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddApplicationInputProcessingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationInputProcessingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationInputProcessingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddApplicationInputProcessingConfiguration"); err != nil {
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
	if err = addOpAddApplicationInputProcessingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationInputProcessingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationInputProcessingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddApplicationInputProcessingConfiguration",
	}
}