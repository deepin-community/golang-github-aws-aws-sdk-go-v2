// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of organization conformance packs.
//
// When you specify the limit and the next token, you receive a paginated
// response.
//
// Limit and next token are not applicable if you specify organization conformance
// packs names. They are only applicable, when you request all the organization
// conformance packs.
//
// # For accounts within an organization
//
// If you deploy an organizational rule or conformance pack in an organization
// administrator account, and then establish a delegated administrator and deploy
// an organizational rule or conformance pack in the delegated administrator
// account, you won't be able to see the organizational rule or conformance pack in
// the organization administrator account from the delegated administrator account
// or see the organizational rule or conformance pack in the delegated
// administrator account from organization administrator account. The
// DescribeOrganizationConfigRules and DescribeOrganizationConformancePacks APIs
// can only see and interact with the organization-related resource that were
// deployed from within the account calling those APIs.
func (c *Client) DescribeOrganizationConformancePacks(ctx context.Context, params *DescribeOrganizationConformancePacksInput, optFns ...func(*Options)) (*DescribeOrganizationConformancePacksOutput, error) {
	if params == nil {
		params = &DescribeOrganizationConformancePacksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationConformancePacks", params, optFns, c.addOperationDescribeOrganizationConformancePacksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationConformancePacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConformancePacksInput struct {

	// The maximum number of organization config packs returned on each page. If you
	// do no specify a number, Config uses the default. The default is 100.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The name that you assign to an organization conformance pack.
	OrganizationConformancePackNames []string

	noSmithyDocumentSerde
}

type DescribeOrganizationConformancePacksOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Returns a list of OrganizationConformancePacks objects.
	OrganizationConformancePacks []types.OrganizationConformancePack

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationConformancePacksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationConformancePacks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationConformancePacks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeOrganizationConformancePacks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationConformancePacks(options.Region), middleware.Before); err != nil {
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

// DescribeOrganizationConformancePacksPaginatorOptions is the paginator options
// for DescribeOrganizationConformancePacks
type DescribeOrganizationConformancePacksPaginatorOptions struct {
	// The maximum number of organization config packs returned on each page. If you
	// do no specify a number, Config uses the default. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOrganizationConformancePacksPaginator is a paginator for
// DescribeOrganizationConformancePacks
type DescribeOrganizationConformancePacksPaginator struct {
	options   DescribeOrganizationConformancePacksPaginatorOptions
	client    DescribeOrganizationConformancePacksAPIClient
	params    *DescribeOrganizationConformancePacksInput
	nextToken *string
	firstPage bool
}

// NewDescribeOrganizationConformancePacksPaginator returns a new
// DescribeOrganizationConformancePacksPaginator
func NewDescribeOrganizationConformancePacksPaginator(client DescribeOrganizationConformancePacksAPIClient, params *DescribeOrganizationConformancePacksInput, optFns ...func(*DescribeOrganizationConformancePacksPaginatorOptions)) *DescribeOrganizationConformancePacksPaginator {
	if params == nil {
		params = &DescribeOrganizationConformancePacksInput{}
	}

	options := DescribeOrganizationConformancePacksPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOrganizationConformancePacksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrganizationConformancePacksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOrganizationConformancePacks page.
func (p *DescribeOrganizationConformancePacksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrganizationConformancePacksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeOrganizationConformancePacks(ctx, &params, optFns...)
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

// DescribeOrganizationConformancePacksAPIClient is a client that implements the
// DescribeOrganizationConformancePacks operation.
type DescribeOrganizationConformancePacksAPIClient interface {
	DescribeOrganizationConformancePacks(context.Context, *DescribeOrganizationConformancePacksInput, ...func(*Options)) (*DescribeOrganizationConformancePacksOutput, error)
}

var _ DescribeOrganizationConformancePacksAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeOrganizationConformancePacks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeOrganizationConformancePacks",
	}
}