// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array that summarizes each privacy budget in a specified
// collaboration. The summary includes the collaboration ARN, creation time,
// creating account, and privacy budget details.
func (c *Client) ListCollaborationPrivacyBudgets(ctx context.Context, params *ListCollaborationPrivacyBudgetsInput, optFns ...func(*Options)) (*ListCollaborationPrivacyBudgetsOutput, error) {
	if params == nil {
		params = &ListCollaborationPrivacyBudgetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCollaborationPrivacyBudgets", params, optFns, c.addOperationListCollaborationPrivacyBudgetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCollaborationPrivacyBudgetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCollaborationPrivacyBudgetsInput struct {

	// A unique identifier for one of your collaborations.
	//
	// This member is required.
	CollaborationIdentifier *string

	// Specifies the type of the privacy budget.
	//
	// This member is required.
	PrivacyBudgetType types.PrivacyBudgetType

	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service may return a nextToken even if the
	// maximum results has not been met.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCollaborationPrivacyBudgetsOutput struct {

	// Summaries of the collaboration privacy budgets.
	//
	// This member is required.
	CollaborationPrivacyBudgetSummaries []types.CollaborationPrivacyBudgetSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCollaborationPrivacyBudgetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCollaborationPrivacyBudgets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCollaborationPrivacyBudgets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCollaborationPrivacyBudgets"); err != nil {
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
	if err = addOpListCollaborationPrivacyBudgetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCollaborationPrivacyBudgets(options.Region), middleware.Before); err != nil {
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

// ListCollaborationPrivacyBudgetsPaginatorOptions is the paginator options for
// ListCollaborationPrivacyBudgets
type ListCollaborationPrivacyBudgetsPaginatorOptions struct {
	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service may return a nextToken even if the
	// maximum results has not been met.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCollaborationPrivacyBudgetsPaginator is a paginator for
// ListCollaborationPrivacyBudgets
type ListCollaborationPrivacyBudgetsPaginator struct {
	options   ListCollaborationPrivacyBudgetsPaginatorOptions
	client    ListCollaborationPrivacyBudgetsAPIClient
	params    *ListCollaborationPrivacyBudgetsInput
	nextToken *string
	firstPage bool
}

// NewListCollaborationPrivacyBudgetsPaginator returns a new
// ListCollaborationPrivacyBudgetsPaginator
func NewListCollaborationPrivacyBudgetsPaginator(client ListCollaborationPrivacyBudgetsAPIClient, params *ListCollaborationPrivacyBudgetsInput, optFns ...func(*ListCollaborationPrivacyBudgetsPaginatorOptions)) *ListCollaborationPrivacyBudgetsPaginator {
	if params == nil {
		params = &ListCollaborationPrivacyBudgetsInput{}
	}

	options := ListCollaborationPrivacyBudgetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCollaborationPrivacyBudgetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCollaborationPrivacyBudgetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCollaborationPrivacyBudgets page.
func (p *ListCollaborationPrivacyBudgetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCollaborationPrivacyBudgetsOutput, error) {
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
	result, err := p.client.ListCollaborationPrivacyBudgets(ctx, &params, optFns...)
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

// ListCollaborationPrivacyBudgetsAPIClient is a client that implements the
// ListCollaborationPrivacyBudgets operation.
type ListCollaborationPrivacyBudgetsAPIClient interface {
	ListCollaborationPrivacyBudgets(context.Context, *ListCollaborationPrivacyBudgetsInput, ...func(*Options)) (*ListCollaborationPrivacyBudgetsOutput, error)
}

var _ ListCollaborationPrivacyBudgetsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCollaborationPrivacyBudgets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCollaborationPrivacyBudgets",
	}
}