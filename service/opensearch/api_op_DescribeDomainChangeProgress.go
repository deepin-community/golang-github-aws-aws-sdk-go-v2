// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the current blue/green deployment happening on an
// Amazon OpenSearch Service domain. For more information, see [Making configuration changes in Amazon OpenSearch Service].
//
// [Making configuration changes in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-configuration-changes.html
func (c *Client) DescribeDomainChangeProgress(ctx context.Context, params *DescribeDomainChangeProgressInput, optFns ...func(*Options)) (*DescribeDomainChangeProgressOutput, error) {
	if params == nil {
		params = &DescribeDomainChangeProgressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomainChangeProgress", params, optFns, c.addOperationDescribeDomainChangeProgressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainChangeProgressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeDomainChangeProgress operation.
type DescribeDomainChangeProgressInput struct {

	// The name of the domain to get progress information for.
	//
	// This member is required.
	DomainName *string

	// The specific change ID for which you want to get progress information. If
	// omitted, the request returns information about the most recent configuration
	// change.
	ChangeId *string

	noSmithyDocumentSerde
}

// The result of a DescribeDomainChangeProgress request. Contains progress
// information for the requested domain change.
type DescribeDomainChangeProgressOutput struct {

	// Container for information about the stages of a configuration change happening
	// on a domain.
	ChangeProgressStatus *types.ChangeProgressStatusDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDomainChangeProgressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDomainChangeProgress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDomainChangeProgress{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDomainChangeProgress"); err != nil {
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
	if err = addOpDescribeDomainChangeProgressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomainChangeProgress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomainChangeProgress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDomainChangeProgress",
	}
}