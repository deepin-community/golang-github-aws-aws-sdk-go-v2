// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about all of the operations that return an operation ID and
// that have ever been performed on domains that were registered by the current
// account.
//
// This command runs only in the us-east-1 Region.
func (c *Client) ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) {
	if params == nil {
		params = &ListOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOperations", params, optFns, c.addOperationListOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListOperations request includes the following elements.
type ListOperationsInput struct {

	// For an initial request for a list of operations, omit this element. If the
	// number of operations that are not yet complete is greater than the value that
	// you specified for MaxItems , you can use Marker to return additional
	// operations. Get the value of NextPageMarker from the previous response, and
	// submit another request that includes the value of NextPageMarker in the Marker
	// element.
	Marker *string

	// Number of domains to be returned.
	//
	// Default: 20
	MaxItems *int32

	//  The sort type for returned values.
	SortBy types.ListOperationsSortAttributeName

	//  The sort order for returned values, either ascending or descending.
	SortOrder types.SortOrder

	//  The status of the operations.
	Status []types.OperationStatus

	// An optional parameter that lets you get information about all the operations
	// that you submitted after a specified date and time. Specify the date and time in
	// Unix time format and Coordinated Universal time (UTC).
	SubmittedSince *time.Time

	//  An arrays of the domains operation types.
	Type []types.OperationType

	noSmithyDocumentSerde
}

// The ListOperations response includes the following elements.
type ListOperationsOutput struct {

	// If there are more operations than you specified for MaxItems in the request,
	// submit another request and include the value of NextPageMarker in the value of
	// Marker .
	NextPageMarker *string

	// Lists summaries of the operations.
	Operations []types.OperationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOperations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOperations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOperations(options.Region), middleware.Before); err != nil {
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

// ListOperationsPaginatorOptions is the paginator options for ListOperations
type ListOperationsPaginatorOptions struct {
	// Number of domains to be returned.
	//
	// Default: 20
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOperationsPaginator is a paginator for ListOperations
type ListOperationsPaginator struct {
	options   ListOperationsPaginatorOptions
	client    ListOperationsAPIClient
	params    *ListOperationsInput
	nextToken *string
	firstPage bool
}

// NewListOperationsPaginator returns a new ListOperationsPaginator
func NewListOperationsPaginator(client ListOperationsAPIClient, params *ListOperationsInput, optFns ...func(*ListOperationsPaginatorOptions)) *ListOperationsPaginator {
	if params == nil {
		params = &ListOperationsInput{}
	}

	options := ListOperationsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOperationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOperations page.
func (p *ListOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOperationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListOperations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListOperationsAPIClient is a client that implements the ListOperations
// operation.
type ListOperationsAPIClient interface {
	ListOperations(context.Context, *ListOperationsInput, ...func(*Options)) (*ListOperationsOutput, error)
}

var _ ListOperationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOperations",
	}
}