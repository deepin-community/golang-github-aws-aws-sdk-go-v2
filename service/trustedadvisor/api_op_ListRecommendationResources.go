// Code generated by smithy-go-codegen DO NOT EDIT.

package trustedadvisor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/trustedadvisor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List Resources of a Recommendation
func (c *Client) ListRecommendationResources(ctx context.Context, params *ListRecommendationResourcesInput, optFns ...func(*Options)) (*ListRecommendationResourcesOutput, error) {
	if params == nil {
		params = &ListRecommendationResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecommendationResources", params, optFns, c.addOperationListRecommendationResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecommendationResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecommendationResourcesInput struct {

	// The Recommendation identifier
	//
	// This member is required.
	RecommendationIdentifier *string

	// The exclusion status of the resource
	ExclusionStatus types.ExclusionStatus

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The AWS Region code of the resource
	RegionCode *string

	// The status of the resource
	Status types.ResourceStatus

	noSmithyDocumentSerde
}

type ListRecommendationResourcesOutput struct {

	// A list of Recommendation Resources
	//
	// This member is required.
	RecommendationResourceSummaries []types.RecommendationResourceSummary

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecommendationResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecommendationResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecommendationResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRecommendationResources"); err != nil {
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
	if err = addOpListRecommendationResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecommendationResources(options.Region), middleware.Before); err != nil {
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

// ListRecommendationResourcesPaginatorOptions is the paginator options for
// ListRecommendationResources
type ListRecommendationResourcesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecommendationResourcesPaginator is a paginator for
// ListRecommendationResources
type ListRecommendationResourcesPaginator struct {
	options   ListRecommendationResourcesPaginatorOptions
	client    ListRecommendationResourcesAPIClient
	params    *ListRecommendationResourcesInput
	nextToken *string
	firstPage bool
}

// NewListRecommendationResourcesPaginator returns a new
// ListRecommendationResourcesPaginator
func NewListRecommendationResourcesPaginator(client ListRecommendationResourcesAPIClient, params *ListRecommendationResourcesInput, optFns ...func(*ListRecommendationResourcesPaginatorOptions)) *ListRecommendationResourcesPaginator {
	if params == nil {
		params = &ListRecommendationResourcesInput{}
	}

	options := ListRecommendationResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecommendationResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecommendationResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecommendationResources page.
func (p *ListRecommendationResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecommendationResourcesOutput, error) {
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
	result, err := p.client.ListRecommendationResources(ctx, &params, optFns...)
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

// ListRecommendationResourcesAPIClient is a client that implements the
// ListRecommendationResources operation.
type ListRecommendationResourcesAPIClient interface {
	ListRecommendationResources(context.Context, *ListRecommendationResourcesInput, ...func(*Options)) (*ListRecommendationResourcesOutput, error)
}

var _ ListRecommendationResourcesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRecommendationResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRecommendationResources",
	}
}