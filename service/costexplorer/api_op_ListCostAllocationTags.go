// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a list of cost allocation tags. All inputs in the API are optional and
// serve as filters. By default, all cost allocation tags are returned.
func (c *Client) ListCostAllocationTags(ctx context.Context, params *ListCostAllocationTagsInput, optFns ...func(*Options)) (*ListCostAllocationTagsOutput, error) {
	if params == nil {
		params = &ListCostAllocationTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCostAllocationTags", params, optFns, c.addOperationListCostAllocationTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCostAllocationTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCostAllocationTagsInput struct {

	// The maximum number of objects that are returned for this request. By default,
	// the request returns 100 results.
	MaxResults *int32

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	// The status of cost allocation tag keys that are returned for this request.
	Status types.CostAllocationTagStatus

	// The list of cost allocation tag keys that are returned for this request.
	TagKeys []string

	// The type of CostAllocationTag object that are returned for this request. The
	// AWSGenerated type tags are tags that Amazon Web Services defines and applies to
	// support Amazon Web Services resources for cost allocation purposes. The
	// UserDefined type tags are tags that you define, create, and apply to resources.
	Type types.CostAllocationTagType

	noSmithyDocumentSerde
}

type ListCostAllocationTagsOutput struct {

	// A list of cost allocation tags that includes the detailed metadata for each
	// one.
	CostAllocationTags []types.CostAllocationTag

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCostAllocationTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCostAllocationTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCostAllocationTags{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCostAllocationTags"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCostAllocationTags(options.Region), middleware.Before); err != nil {
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

// ListCostAllocationTagsPaginatorOptions is the paginator options for
// ListCostAllocationTags
type ListCostAllocationTagsPaginatorOptions struct {
	// The maximum number of objects that are returned for this request. By default,
	// the request returns 100 results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCostAllocationTagsPaginator is a paginator for ListCostAllocationTags
type ListCostAllocationTagsPaginator struct {
	options   ListCostAllocationTagsPaginatorOptions
	client    ListCostAllocationTagsAPIClient
	params    *ListCostAllocationTagsInput
	nextToken *string
	firstPage bool
}

// NewListCostAllocationTagsPaginator returns a new ListCostAllocationTagsPaginator
func NewListCostAllocationTagsPaginator(client ListCostAllocationTagsAPIClient, params *ListCostAllocationTagsInput, optFns ...func(*ListCostAllocationTagsPaginatorOptions)) *ListCostAllocationTagsPaginator {
	if params == nil {
		params = &ListCostAllocationTagsInput{}
	}

	options := ListCostAllocationTagsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCostAllocationTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCostAllocationTagsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCostAllocationTags page.
func (p *ListCostAllocationTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCostAllocationTagsOutput, error) {
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
	result, err := p.client.ListCostAllocationTags(ctx, &params, optFns...)
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

// ListCostAllocationTagsAPIClient is a client that implements the
// ListCostAllocationTags operation.
type ListCostAllocationTagsAPIClient interface {
	ListCostAllocationTags(context.Context, *ListCostAllocationTagsInput, ...func(*Options)) (*ListCostAllocationTagsOutput, error)
}

var _ ListCostAllocationTagsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCostAllocationTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCostAllocationTags",
	}
}