// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of publishing destinations associated with the specified
// detectorId .
func (c *Client) ListPublishingDestinations(ctx context.Context, params *ListPublishingDestinationsInput, optFns ...func(*Options)) (*ListPublishingDestinationsOutput, error) {
	if params == nil {
		params = &ListPublishingDestinationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPublishingDestinations", params, optFns, c.addOperationListPublishingDestinationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPublishingDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPublishingDestinationsInput struct {

	// The ID of the detector to retrieve publishing destinations for.
	//
	// This member is required.
	DetectorId *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token to use for paginating results that are returned in the response. Set
	// the value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPublishingDestinationsOutput struct {

	// A Destinations object that includes information about each publishing
	// destination returned.
	//
	// This member is required.
	Destinations []types.Destination

	// A token to use for paginating results that are returned in the response. Set
	// the value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPublishingDestinationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPublishingDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPublishingDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPublishingDestinations"); err != nil {
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
	if err = addOpListPublishingDestinationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPublishingDestinations(options.Region), middleware.Before); err != nil {
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

// ListPublishingDestinationsPaginatorOptions is the paginator options for
// ListPublishingDestinations
type ListPublishingDestinationsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPublishingDestinationsPaginator is a paginator for
// ListPublishingDestinations
type ListPublishingDestinationsPaginator struct {
	options   ListPublishingDestinationsPaginatorOptions
	client    ListPublishingDestinationsAPIClient
	params    *ListPublishingDestinationsInput
	nextToken *string
	firstPage bool
}

// NewListPublishingDestinationsPaginator returns a new
// ListPublishingDestinationsPaginator
func NewListPublishingDestinationsPaginator(client ListPublishingDestinationsAPIClient, params *ListPublishingDestinationsInput, optFns ...func(*ListPublishingDestinationsPaginatorOptions)) *ListPublishingDestinationsPaginator {
	if params == nil {
		params = &ListPublishingDestinationsInput{}
	}

	options := ListPublishingDestinationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPublishingDestinationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPublishingDestinationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPublishingDestinations page.
func (p *ListPublishingDestinationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPublishingDestinationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListPublishingDestinations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListPublishingDestinationsAPIClient is a client that implements the
// ListPublishingDestinations operation.
type ListPublishingDestinationsAPIClient interface {
	ListPublishingDestinations(context.Context, *ListPublishingDestinationsInput, ...func(*Options)) (*ListPublishingDestinationsOutput, error)
}

var _ ListPublishingDestinationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPublishingDestinations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPublishingDestinations",
	}
}