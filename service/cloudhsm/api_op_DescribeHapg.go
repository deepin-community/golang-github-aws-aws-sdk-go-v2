// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Retrieves information about a high-availability partition group.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
func (c *Client) DescribeHapg(ctx context.Context, params *DescribeHapgInput, optFns ...func(*Options)) (*DescribeHapgOutput, error) {
	if params == nil {
		params = &DescribeHapgInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHapg", params, optFns, c.addOperationDescribeHapgMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHapgOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DescribeHapg action.
type DescribeHapgInput struct {

	// The ARN of the high-availability partition group to describe.
	//
	// This member is required.
	HapgArn *string

	noSmithyDocumentSerde
}

// Contains the output of the DescribeHapg action.
type DescribeHapgOutput struct {

	// The ARN of the high-availability partition group.
	HapgArn *string

	// The serial number of the high-availability partition group.
	HapgSerial *string

	//
	HsmsLastActionFailed []string

	//
	HsmsPendingDeletion []string

	//
	HsmsPendingRegistration []string

	// The label for the high-availability partition group.
	Label *string

	// The date and time the high-availability partition group was last modified.
	LastModifiedTimestamp *string

	// The list of partition serial numbers that belong to the high-availability
	// partition group.
	PartitionSerialList []string

	// The state of the high-availability partition group.
	State types.CloudHsmObjectState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHapgMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHapg{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHapg{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeHapg"); err != nil {
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
	if err = addOpDescribeHapgValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHapg(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHapg(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeHapg",
	}
}