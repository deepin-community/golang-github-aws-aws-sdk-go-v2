// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the hours of operation for the specified Amazon
// Connect instance.
//
// For more information about hours of operation, see [Set the Hours of Operation for a Queue] in the Amazon Connect
// Administrator Guide.
//
// [Set the Hours of Operation for a Queue]: https://docs.aws.amazon.com/connect/latest/adminguide/set-hours-operation.html
func (c *Client) ListHoursOfOperations(ctx context.Context, params *ListHoursOfOperationsInput, optFns ...func(*Options)) (*ListHoursOfOperationsOutput, error) {
	if params == nil {
		params = &ListHoursOfOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHoursOfOperations", params, optFns, c.addOperationListHoursOfOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHoursOfOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHoursOfOperationsInput struct {

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListHoursOfOperationsOutput struct {

	// Information about the hours of operation.
	HoursOfOperationSummaryList []types.HoursOfOperationSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHoursOfOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListHoursOfOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListHoursOfOperations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListHoursOfOperations"); err != nil {
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
	if err = addOpListHoursOfOperationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHoursOfOperations(options.Region), middleware.Before); err != nil {
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

// ListHoursOfOperationsPaginatorOptions is the paginator options for
// ListHoursOfOperations
type ListHoursOfOperationsPaginatorOptions struct {
	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHoursOfOperationsPaginator is a paginator for ListHoursOfOperations
type ListHoursOfOperationsPaginator struct {
	options   ListHoursOfOperationsPaginatorOptions
	client    ListHoursOfOperationsAPIClient
	params    *ListHoursOfOperationsInput
	nextToken *string
	firstPage bool
}

// NewListHoursOfOperationsPaginator returns a new ListHoursOfOperationsPaginator
func NewListHoursOfOperationsPaginator(client ListHoursOfOperationsAPIClient, params *ListHoursOfOperationsInput, optFns ...func(*ListHoursOfOperationsPaginatorOptions)) *ListHoursOfOperationsPaginator {
	if params == nil {
		params = &ListHoursOfOperationsInput{}
	}

	options := ListHoursOfOperationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHoursOfOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHoursOfOperationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHoursOfOperations page.
func (p *ListHoursOfOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHoursOfOperationsOutput, error) {
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
	result, err := p.client.ListHoursOfOperations(ctx, &params, optFns...)
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

// ListHoursOfOperationsAPIClient is a client that implements the
// ListHoursOfOperations operation.
type ListHoursOfOperationsAPIClient interface {
	ListHoursOfOperations(context.Context, *ListHoursOfOperationsInput, ...func(*Options)) (*ListHoursOfOperationsOutput, error)
}

var _ ListHoursOfOperationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListHoursOfOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListHoursOfOperations",
	}
}