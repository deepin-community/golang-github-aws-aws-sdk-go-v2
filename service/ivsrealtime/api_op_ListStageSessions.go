// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets all sessions for a specified stage.
func (c *Client) ListStageSessions(ctx context.Context, params *ListStageSessionsInput, optFns ...func(*Options)) (*ListStageSessionsOutput, error) {
	if params == nil {
		params = &ListStageSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStageSessions", params, optFns, c.addOperationListStageSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStageSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStageSessionsInput struct {

	// Stage ARN.
	//
	// This member is required.
	StageArn *string

	// Maximum number of results to return. Default: 50.
	MaxResults *int32

	// The first stage session to retrieve. This is used for pagination; see the
	// nextToken response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStageSessionsOutput struct {

	// List of matching stage sessions.
	//
	// This member is required.
	StageSessions []types.StageSessionSummary

	// If there are more stage sessions than maxResults , use nextToken in the request
	// to get the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStageSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStageSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStageSessions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStageSessions"); err != nil {
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
	if err = addOpListStageSessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStageSessions(options.Region), middleware.Before); err != nil {
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

// ListStageSessionsPaginatorOptions is the paginator options for ListStageSessions
type ListStageSessionsPaginatorOptions struct {
	// Maximum number of results to return. Default: 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStageSessionsPaginator is a paginator for ListStageSessions
type ListStageSessionsPaginator struct {
	options   ListStageSessionsPaginatorOptions
	client    ListStageSessionsAPIClient
	params    *ListStageSessionsInput
	nextToken *string
	firstPage bool
}

// NewListStageSessionsPaginator returns a new ListStageSessionsPaginator
func NewListStageSessionsPaginator(client ListStageSessionsAPIClient, params *ListStageSessionsInput, optFns ...func(*ListStageSessionsPaginatorOptions)) *ListStageSessionsPaginator {
	if params == nil {
		params = &ListStageSessionsInput{}
	}

	options := ListStageSessionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStageSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStageSessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStageSessions page.
func (p *ListStageSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStageSessionsOutput, error) {
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
	result, err := p.client.ListStageSessions(ctx, &params, optFns...)
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

// ListStageSessionsAPIClient is a client that implements the ListStageSessions
// operation.
type ListStageSessionsAPIClient interface {
	ListStageSessions(context.Context, *ListStageSessionsInput, ...func(*Options)) (*ListStageSessionsOutput, error)
}

var _ ListStageSessionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListStageSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStageSessions",
	}
}