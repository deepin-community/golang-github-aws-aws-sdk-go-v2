// Code generated by smithy-go-codegen DO NOT EDIT.

package repostspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/repostspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of AWS re:Post Private private re:Posts in the account with some
// information about each private re:Post.
func (c *Client) ListSpaces(ctx context.Context, params *ListSpacesInput, optFns ...func(*Options)) (*ListSpacesOutput, error) {
	if params == nil {
		params = &ListSpacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSpaces", params, optFns, c.addOperationListSpacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSpacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSpacesInput struct {

	// The maximum number of private re:Posts to include in the results.
	MaxResults *int32

	// The token for the next set of private re:Posts to return. You receive this
	// token from a previous ListSpaces operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSpacesOutput struct {

	// An array of structures that contain some information about the private re:Posts
	// in the account.
	//
	// This member is required.
	Spaces []types.SpaceData

	// The token that you use when you request the next set of private re:Posts.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSpacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSpaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSpaces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSpaces"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSpaces(options.Region), middleware.Before); err != nil {
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

// ListSpacesPaginatorOptions is the paginator options for ListSpaces
type ListSpacesPaginatorOptions struct {
	// The maximum number of private re:Posts to include in the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSpacesPaginator is a paginator for ListSpaces
type ListSpacesPaginator struct {
	options   ListSpacesPaginatorOptions
	client    ListSpacesAPIClient
	params    *ListSpacesInput
	nextToken *string
	firstPage bool
}

// NewListSpacesPaginator returns a new ListSpacesPaginator
func NewListSpacesPaginator(client ListSpacesAPIClient, params *ListSpacesInput, optFns ...func(*ListSpacesPaginatorOptions)) *ListSpacesPaginator {
	if params == nil {
		params = &ListSpacesInput{}
	}

	options := ListSpacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSpacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSpacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSpaces page.
func (p *ListSpacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSpacesOutput, error) {
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
	result, err := p.client.ListSpaces(ctx, &params, optFns...)
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

// ListSpacesAPIClient is a client that implements the ListSpaces operation.
type ListSpacesAPIClient interface {
	ListSpaces(context.Context, *ListSpacesInput, ...func(*Options)) (*ListSpacesOutput, error)
}

var _ ListSpacesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSpaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSpaces",
	}
}