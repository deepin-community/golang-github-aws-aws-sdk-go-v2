// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationsignals

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/applicationsignals/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of the operations of this service that have been discovered by
// Application Signals. Only the operations that were invoked during the specified
// time range are returned.
func (c *Client) ListServiceOperations(ctx context.Context, params *ListServiceOperationsInput, optFns ...func(*Options)) (*ListServiceOperationsOutput, error) {
	if params == nil {
		params = &ListServiceOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceOperations", params, optFns, c.addOperationListServiceOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceOperationsInput struct {

	// The end of the time period to retrieve information about. When used in a raw
	// HTTP Query API, it is formatted as be epoch time in seconds. For example:
	// 1698778057
	//
	// This member is required.
	EndTime *time.Time

	// Use this field to specify which service you want to retrieve information for.
	// You must specify at least the Type , Name , and Environment attributes.
	//
	// This is a string-to-string map. It can include the following fields.
	//
	//   - Type designates the type of object this is.
	//
	//   - ResourceType specifies the type of the resource. This field is used only
	//   when the value of the Type field is Resource or AWS::Resource .
	//
	//   - Name specifies the name of the object. This is used only if the value of the
	//   Type field is Service , RemoteService , or AWS::Service .
	//
	//   - Identifier identifies the resource objects of this resource. This is used
	//   only if the value of the Type field is Resource or AWS::Resource .
	//
	//   - Environment specifies the location where this object is hosted, or what it
	//   belongs to.
	//
	// This member is required.
	KeyAttributes map[string]string

	// The start of the time period to retrieve information about. When used in a raw
	// HTTP Query API, it is formatted as be epoch time in seconds. For example:
	// 1698778057
	//
	// This member is required.
	StartTime *time.Time

	// The maximum number of results to return in one operation. If you omit this
	// parameter, the default of 50 is used.
	MaxResults *int32

	// Include this value, if it was returned by the previous operation, to get the
	// next set of service operations.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceOperationsOutput struct {

	// The end of the time period that the returned information applies to. When used
	// in a raw HTTP Query API, it is formatted as be epoch time in seconds. For
	// example: 1698778057
	//
	// This member is required.
	EndTime *time.Time

	// An array of structures that each contain information about one operation of
	// this service.
	//
	// This member is required.
	ServiceOperations []types.ServiceOperation

	// The start of the time period that the returned information applies to. When
	// used in a raw HTTP Query API, it is formatted as be epoch time in seconds. For
	// example: 1698778057
	//
	// This member is required.
	StartTime *time.Time

	// Include this value in your next use of this API to get next set of service
	// operations.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListServiceOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListServiceOperations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListServiceOperations"); err != nil {
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
	if err = addOpListServiceOperationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceOperations(options.Region), middleware.Before); err != nil {
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

// ListServiceOperationsPaginatorOptions is the paginator options for
// ListServiceOperations
type ListServiceOperationsPaginatorOptions struct {
	// The maximum number of results to return in one operation. If you omit this
	// parameter, the default of 50 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceOperationsPaginator is a paginator for ListServiceOperations
type ListServiceOperationsPaginator struct {
	options   ListServiceOperationsPaginatorOptions
	client    ListServiceOperationsAPIClient
	params    *ListServiceOperationsInput
	nextToken *string
	firstPage bool
}

// NewListServiceOperationsPaginator returns a new ListServiceOperationsPaginator
func NewListServiceOperationsPaginator(client ListServiceOperationsAPIClient, params *ListServiceOperationsInput, optFns ...func(*ListServiceOperationsPaginatorOptions)) *ListServiceOperationsPaginator {
	if params == nil {
		params = &ListServiceOperationsInput{}
	}

	options := ListServiceOperationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceOperationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceOperations page.
func (p *ListServiceOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceOperationsOutput, error) {
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
	result, err := p.client.ListServiceOperations(ctx, &params, optFns...)
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

// ListServiceOperationsAPIClient is a client that implements the
// ListServiceOperations operation.
type ListServiceOperationsAPIClient interface {
	ListServiceOperations(context.Context, *ListServiceOperationsInput, ...func(*Options)) (*ListServiceOperationsOutput, error)
}

var _ ListServiceOperationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListServiceOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListServiceOperations",
	}
}