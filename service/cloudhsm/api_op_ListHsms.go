// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Retrieves the identifiers of all of the HSMs provisioned for the current
// customer.
//
// This operation supports pagination with the use of the NextToken member. If
// more results are available, the NextToken member of the response contains a
// token that you pass in the next call to ListHsms to retrieve the next set of
// items.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
func (c *Client) ListHsms(ctx context.Context, params *ListHsmsInput, optFns ...func(*Options)) (*ListHsmsOutput, error) {
	if params == nil {
		params = &ListHsmsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHsms", params, optFns, c.addOperationListHsmsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHsmsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHsmsInput struct {

	// The NextToken value from a previous call to ListHsms . Pass null if this is the
	// first call.
	NextToken *string

	noSmithyDocumentSerde
}

// Contains the output of the ListHsms operation.
type ListHsmsOutput struct {

	// The list of ARNs that identify the HSMs.
	HsmList []string

	// If not null, more results are available. Pass this value to ListHsms to
	// retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHsmsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHsms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHsms{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListHsms"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHsms(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListHsms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListHsms",
	}
}