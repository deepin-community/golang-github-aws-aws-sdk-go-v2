// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of controls from Audit Manager.
func (c *Client) ListControls(ctx context.Context, params *ListControlsInput, optFns ...func(*Options)) (*ListControlsOutput, error) {
	if params == nil {
		params = &ListControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListControls", params, optFns, c.addOperationListControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListControlsInput struct {

	// A filter that narrows the list of controls to a specific type.
	//
	// This member is required.
	ControlType types.ControlType

	// A filter that narrows the list of controls to a specific resource from the
	// Amazon Web Services Control Catalog.
	//
	// To use this parameter, specify the ARN of the Control Catalog resource. You can
	// specify either a control domain, a control objective, or a common control. For
	// information about how to find the ARNs for these resources, see [ListDomains]ListDomains , [ListObjectives]
	// ListObjectives , and [ListCommonControls]ListCommonControls .
	//
	// You can only filter by one Control Catalog resource at a time. Specifying
	// multiple resource ARNs isn’t currently supported. If you want to filter by more
	// than one ARN, we recommend that you run the ListControls operation separately
	// for each ARN.
	//
	// Alternatively, specify UNCATEGORIZED to list controls that aren't mapped to a
	// Control Catalog resource. For example, this operation might return a list of
	// custom controls that don't belong to any control domain or control objective.
	//
	// [ListCommonControls]: https://docs.aws.amazon.com/controlcatalog/latest/APIReference/API_ListCommonControls.html
	// [ListDomains]: https://docs.aws.amazon.com/controlcatalog/latest/APIReference/API_ListDomains.html
	// [ListObjectives]: https://docs.aws.amazon.com/controlcatalog/latest/APIReference/API_ListObjectives.html
	ControlCatalogId *string

	// The maximum number of results on a page or for an API request call.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListControlsOutput struct {

	//  A list of metadata that the ListControls API returns for each control.
	ControlMetadataList []types.ControlMetadata

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListControls{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListControls"); err != nil {
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
	if err = addOpListControlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListControls(options.Region), middleware.Before); err != nil {
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

// ListControlsPaginatorOptions is the paginator options for ListControls
type ListControlsPaginatorOptions struct {
	// The maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListControlsPaginator is a paginator for ListControls
type ListControlsPaginator struct {
	options   ListControlsPaginatorOptions
	client    ListControlsAPIClient
	params    *ListControlsInput
	nextToken *string
	firstPage bool
}

// NewListControlsPaginator returns a new ListControlsPaginator
func NewListControlsPaginator(client ListControlsAPIClient, params *ListControlsInput, optFns ...func(*ListControlsPaginatorOptions)) *ListControlsPaginator {
	if params == nil {
		params = &ListControlsInput{}
	}

	options := ListControlsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListControlsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListControlsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListControls page.
func (p *ListControlsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListControlsOutput, error) {
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
	result, err := p.client.ListControls(ctx, &params, optFns...)
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

// ListControlsAPIClient is a client that implements the ListControls operation.
type ListControlsAPIClient interface {
	ListControls(context.Context, *ListControlsInput, ...func(*Options)) (*ListControlsOutput, error)
}

var _ ListControlsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListControls",
	}
}