// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the step groups in a migration workflow.
func (c *Client) ListWorkflowStepGroups(ctx context.Context, params *ListWorkflowStepGroupsInput, optFns ...func(*Options)) (*ListWorkflowStepGroupsOutput, error) {
	if params == nil {
		params = &ListWorkflowStepGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkflowStepGroups", params, optFns, c.addOperationListWorkflowStepGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkflowStepGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkflowStepGroupsInput struct {

	// The ID of the migration workflow.
	//
	// This member is required.
	WorkflowId *string

	// The maximum number of results that can be returned.
	MaxResults int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkflowStepGroupsOutput struct {

	// The summary of step groups in a migration workflow.
	//
	// This member is required.
	WorkflowStepGroupsSummary []types.WorkflowStepGroupSummary

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkflowStepGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkflowStepGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkflowStepGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkflowStepGroups"); err != nil {
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
	if err = addOpListWorkflowStepGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkflowStepGroups(options.Region), middleware.Before); err != nil {
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

// ListWorkflowStepGroupsPaginatorOptions is the paginator options for
// ListWorkflowStepGroups
type ListWorkflowStepGroupsPaginatorOptions struct {
	// The maximum number of results that can be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkflowStepGroupsPaginator is a paginator for ListWorkflowStepGroups
type ListWorkflowStepGroupsPaginator struct {
	options   ListWorkflowStepGroupsPaginatorOptions
	client    ListWorkflowStepGroupsAPIClient
	params    *ListWorkflowStepGroupsInput
	nextToken *string
	firstPage bool
}

// NewListWorkflowStepGroupsPaginator returns a new ListWorkflowStepGroupsPaginator
func NewListWorkflowStepGroupsPaginator(client ListWorkflowStepGroupsAPIClient, params *ListWorkflowStepGroupsInput, optFns ...func(*ListWorkflowStepGroupsPaginatorOptions)) *ListWorkflowStepGroupsPaginator {
	if params == nil {
		params = &ListWorkflowStepGroupsInput{}
	}

	options := ListWorkflowStepGroupsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkflowStepGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkflowStepGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkflowStepGroups page.
func (p *ListWorkflowStepGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkflowStepGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListWorkflowStepGroups(ctx, &params, optFns...)
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

// ListWorkflowStepGroupsAPIClient is a client that implements the
// ListWorkflowStepGroups operation.
type ListWorkflowStepGroupsAPIClient interface {
	ListWorkflowStepGroups(context.Context, *ListWorkflowStepGroupsInput, ...func(*Options)) (*ListWorkflowStepGroupsOutput, error)
}

var _ ListWorkflowStepGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListWorkflowStepGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkflowStepGroups",
	}
}