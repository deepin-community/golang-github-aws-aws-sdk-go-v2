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
// Adds an external destination to your Amazon Kinesis Analytics application.
//
// If you want Amazon Kinesis Analytics to deliver data from an in-application
// stream within your application to an external destination (such as an Amazon
// Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an AWS Lambda
// function), you add the relevant configuration to your application using this
// operation. You can configure one or more outputs for your application. Each
// output configuration maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your
// in-application error stream to an external destination so that you can analyze
// the errors. For more information, see [Understanding Application Output (Destination)].
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the [DescribeApplication]
// operation to find the current application version.
//
// For the limits on the number of application inputs and outputs you can
// configure, see [Limits].
//
// This operation requires permissions to perform the
// kinesisanalytics:AddApplicationOutput action.
//
// [Limits]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html
// [Understanding Application Output (Destination)]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
func (c *Client) AddApplicationOutput(ctx context.Context, params *AddApplicationOutputInput, optFns ...func(*Options)) (*AddApplicationOutputOutput, error) {
	if params == nil {
		params = &AddApplicationOutputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationOutput", params, optFns, c.addOperationAddApplicationOutputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationOutputInput struct {

	// Name of the application to which you want to add the output configuration.
	//
	// This member is required.
	ApplicationName *string

	// Version of the application to which you want to add the output configuration.
	// You can use the [DescribeApplication]operation to get the current application version. If the
	// version specified is not the current version, the
	// ConcurrentModificationException is returned.
	//
	// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// An array of objects, each describing one output configuration. In the output
	// configuration, you specify the name of an in-application stream, a destination
	// (that is, an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream,
	// or an AWS Lambda function), and record the formation to use when writing to the
	// destination.
	//
	// This member is required.
	Output *types.Output

	noSmithyDocumentSerde
}

type AddApplicationOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddApplicationOutputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationOutput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationOutput{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddApplicationOutput"); err != nil {
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
	if err = addOpAddApplicationOutputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationOutput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationOutput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddApplicationOutput",
	}
}