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

// Lists information about requested configured audience model associations.
func (c *Client) ListConfiguredAudienceModelAssociations(ctx context.Context, params *ListConfiguredAudienceModelAssociationsInput, optFns ...func(*Options)) (*ListConfiguredAudienceModelAssociationsOutput, error) {
	if params == nil {
		params = &ListConfiguredAudienceModelAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfiguredAudienceModelAssociations", params, optFns, c.addOperationListConfiguredAudienceModelAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfiguredAudienceModelAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfiguredAudienceModelAssociationsInput struct {

	// A unique identifier for a membership that contains the configured audience
	// model associations that you want to retrieve.
	//
	// This member is required.
	MembershipIdentifier *string

	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service may return a nextToken even if the
	// maximum results has not been met.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConfiguredAudienceModelAssociationsOutput struct {

	// Summaries of the configured audience model associations that you requested.
	//
	// This member is required.
	ConfiguredAudienceModelAssociationSummaries []types.ConfiguredAudienceModelAssociationSummary

	// The token value provided to access the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConfiguredAudienceModelAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfiguredAudienceModelAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfiguredAudienceModelAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConfiguredAudienceModelAssociations"); err != nil {
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
	if err = addOpListConfiguredAudienceModelAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfiguredAudienceModelAssociations(options.Region), middleware.Before); err != nil {
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

// ListConfiguredAudienceModelAssociationsPaginatorOptions is the paginator
// options for ListConfiguredAudienceModelAssociations
type ListConfiguredAudienceModelAssociationsPaginatorOptions struct {
	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service may return a nextToken even if the
	// maximum results has not been met.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConfiguredAudienceModelAssociationsPaginator is a paginator for
// ListConfiguredAudienceModelAssociations
type ListConfiguredAudienceModelAssociationsPaginator struct {
	options   ListConfiguredAudienceModelAssociationsPaginatorOptions
	client    ListConfiguredAudienceModelAssociationsAPIClient
	params    *ListConfiguredAudienceModelAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListConfiguredAudienceModelAssociationsPaginator returns a new
// ListConfiguredAudienceModelAssociationsPaginator
func NewListConfiguredAudienceModelAssociationsPaginator(client ListConfiguredAudienceModelAssociationsAPIClient, params *ListConfiguredAudienceModelAssociationsInput, optFns ...func(*ListConfiguredAudienceModelAssociationsPaginatorOptions)) *ListConfiguredAudienceModelAssociationsPaginator {
	if params == nil {
		params = &ListConfiguredAudienceModelAssociationsInput{}
	}

	options := ListConfiguredAudienceModelAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConfiguredAudienceModelAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConfiguredAudienceModelAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConfiguredAudienceModelAssociations page.
func (p *ListConfiguredAudienceModelAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConfiguredAudienceModelAssociationsOutput, error) {
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
	result, err := p.client.ListConfiguredAudienceModelAssociations(ctx, &params, optFns...)
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

// ListConfiguredAudienceModelAssociationsAPIClient is a client that implements
// the ListConfiguredAudienceModelAssociations operation.
type ListConfiguredAudienceModelAssociationsAPIClient interface {
	ListConfiguredAudienceModelAssociations(context.Context, *ListConfiguredAudienceModelAssociationsInput, ...func(*Options)) (*ListConfiguredAudienceModelAssociationsOutput, error)
}

var _ ListConfiguredAudienceModelAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListConfiguredAudienceModelAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConfiguredAudienceModelAssociations",
	}
}