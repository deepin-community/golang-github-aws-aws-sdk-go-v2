// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the retriever used by an Amazon Q Business application.
func (c *Client) ListRetrievers(ctx context.Context, params *ListRetrieversInput, optFns ...func(*Options)) (*ListRetrieversOutput, error) {
	if params == nil {
		params = &ListRetrieversInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRetrievers", params, optFns, c.addOperationListRetrieversMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRetrieversOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRetrieversInput struct {

	// The identifier of the Amazon Q Business application using the retriever.
	//
	// This member is required.
	ApplicationId *string

	// The maximum number of retrievers returned.
	MaxResults *int32

	// If the number of retrievers returned exceeds maxResults , Amazon Q Business
	// returns a next token as a pagination token to retrieve the next set of
	// retrievers.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRetrieversOutput struct {

	// If the response is truncated, Amazon Q Business returns this token, which you
	// can use in a later request to list the next set of retrievers.
	NextToken *string

	// An array of summary information for one or more retrievers.
	Retrievers []types.Retriever

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRetrieversMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRetrievers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRetrievers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRetrievers"); err != nil {
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
	if err = addOpListRetrieversValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRetrievers(options.Region), middleware.Before); err != nil {
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

// ListRetrieversPaginatorOptions is the paginator options for ListRetrievers
type ListRetrieversPaginatorOptions struct {
	// The maximum number of retrievers returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRetrieversPaginator is a paginator for ListRetrievers
type ListRetrieversPaginator struct {
	options   ListRetrieversPaginatorOptions
	client    ListRetrieversAPIClient
	params    *ListRetrieversInput
	nextToken *string
	firstPage bool
}

// NewListRetrieversPaginator returns a new ListRetrieversPaginator
func NewListRetrieversPaginator(client ListRetrieversAPIClient, params *ListRetrieversInput, optFns ...func(*ListRetrieversPaginatorOptions)) *ListRetrieversPaginator {
	if params == nil {
		params = &ListRetrieversInput{}
	}

	options := ListRetrieversPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRetrieversPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRetrieversPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRetrievers page.
func (p *ListRetrieversPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRetrieversOutput, error) {
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
	result, err := p.client.ListRetrievers(ctx, &params, optFns...)
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

// ListRetrieversAPIClient is a client that implements the ListRetrievers
// operation.
type ListRetrieversAPIClient interface {
	ListRetrievers(context.Context, *ListRetrieversInput, ...func(*Options)) (*ListRetrieversOutput, error)
}

var _ ListRetrieversAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRetrievers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRetrievers",
	}
}