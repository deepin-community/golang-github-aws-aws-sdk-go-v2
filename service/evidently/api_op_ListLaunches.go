// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns configuration details about all the launches in the specified project.
func (c *Client) ListLaunches(ctx context.Context, params *ListLaunchesInput, optFns ...func(*Options)) (*ListLaunchesOutput, error) {
	if params == nil {
		params = &ListLaunchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLaunches", params, optFns, c.addOperationListLaunchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLaunchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLaunchesInput struct {

	// The name or ARN of the project to return the launch list from.
	//
	// This member is required.
	Project *string

	// The maximum number of results to include in the response.
	MaxResults *int32

	// The token to use when requesting the next set of results. You received this
	// token from a previous ListLaunches operation.
	NextToken *string

	// Use this optional parameter to limit the returned results to only the launches
	// with the status that you specify here.
	Status types.LaunchStatus

	noSmithyDocumentSerde
}

type ListLaunchesOutput struct {

	// An array of structures that contain the configuration details of the launches
	// in the specified project.
	Launches []types.Launch

	// The token to use in a subsequent ListLaunches operation to return the next set
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLaunchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLaunches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLaunches{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLaunches"); err != nil {
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
	if err = addOpListLaunchesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLaunches(options.Region), middleware.Before); err != nil {
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

// ListLaunchesPaginatorOptions is the paginator options for ListLaunches
type ListLaunchesPaginatorOptions struct {
	// The maximum number of results to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLaunchesPaginator is a paginator for ListLaunches
type ListLaunchesPaginator struct {
	options   ListLaunchesPaginatorOptions
	client    ListLaunchesAPIClient
	params    *ListLaunchesInput
	nextToken *string
	firstPage bool
}

// NewListLaunchesPaginator returns a new ListLaunchesPaginator
func NewListLaunchesPaginator(client ListLaunchesAPIClient, params *ListLaunchesInput, optFns ...func(*ListLaunchesPaginatorOptions)) *ListLaunchesPaginator {
	if params == nil {
		params = &ListLaunchesInput{}
	}

	options := ListLaunchesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLaunchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLaunchesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLaunches page.
func (p *ListLaunchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLaunchesOutput, error) {
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
	result, err := p.client.ListLaunches(ctx, &params, optFns...)
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

// ListLaunchesAPIClient is a client that implements the ListLaunches operation.
type ListLaunchesAPIClient interface {
	ListLaunches(context.Context, *ListLaunchesInput, ...func(*Options)) (*ListLaunchesOutput, error)
}

var _ ListLaunchesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListLaunches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLaunches",
	}
}