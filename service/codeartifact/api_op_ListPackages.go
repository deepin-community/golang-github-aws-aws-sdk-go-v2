// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns a list of [PackageSummary] objects for packages in a repository that match the request
//
// parameters.
//
// [PackageSummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html
func (c *Client) ListPackages(ctx context.Context, params *ListPackagesInput, optFns ...func(*Options)) (*ListPackagesOutput, error) {
	if params == nil {
		params = &ListPackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackages", params, optFns, c.addOperationListPackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackagesInput struct {

	//  The name of the domain that contains the repository that contains the
	// requested packages.
	//
	// This member is required.
	Domain *string

	//  The name of the repository that contains the requested packages.
	//
	// This member is required.
	Repository *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The format used to filter requested packages. Only packages from the provided
	// format will be returned.
	Format types.PackageFormat

	//  The maximum number of results to return per page.
	MaxResults *int32

	// The namespace prefix used to filter requested packages. Only packages with a
	// namespace that starts with the provided string value are returned. Note that
	// although this option is called --namespace and not --namespace-prefix , it has
	// prefix-matching behavior.
	//
	// Each package format uses namespace as follows:
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	//  The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	//  A prefix used to filter requested packages. Only packages with names that
	// start with packagePrefix are returned.
	PackagePrefix *string

	// The value of the Publish package origin control restriction used to filter
	// requested packages. Only packages with the provided restriction are returned.
	// For more information, see [PackageOriginRestrictions].
	//
	// [PackageOriginRestrictions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageOriginRestrictions.html
	Publish types.AllowPublish

	// The value of the Upstream package origin control restriction used to filter
	// requested packages. Only packages with the provided restriction are returned.
	// For more information, see [PackageOriginRestrictions].
	//
	// [PackageOriginRestrictions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageOriginRestrictions.html
	Upstream types.AllowUpstream

	noSmithyDocumentSerde
}

type ListPackagesOutput struct {

	//  If there are additional results, this is the token for the next set of
	// results.
	NextToken *string

	//  The list of returned [PackageSummary] objects.
	//
	// [PackageSummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html
	Packages []types.PackageSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackages{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPackages"); err != nil {
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
	if err = addOpListPackagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackages(options.Region), middleware.Before); err != nil {
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

// ListPackagesPaginatorOptions is the paginator options for ListPackages
type ListPackagesPaginatorOptions struct {
	//  The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackagesPaginator is a paginator for ListPackages
type ListPackagesPaginator struct {
	options   ListPackagesPaginatorOptions
	client    ListPackagesAPIClient
	params    *ListPackagesInput
	nextToken *string
	firstPage bool
}

// NewListPackagesPaginator returns a new ListPackagesPaginator
func NewListPackagesPaginator(client ListPackagesAPIClient, params *ListPackagesInput, optFns ...func(*ListPackagesPaginatorOptions)) *ListPackagesPaginator {
	if params == nil {
		params = &ListPackagesInput{}
	}

	options := ListPackagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPackagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPackages page.
func (p *ListPackagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackagesOutput, error) {
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
	result, err := p.client.ListPackages(ctx, &params, optFns...)
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

// ListPackagesAPIClient is a client that implements the ListPackages operation.
type ListPackagesAPIClient interface {
	ListPackages(context.Context, *ListPackagesInput, ...func(*Options)) (*ListPackagesOutput, error)
}

var _ ListPackagesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPackages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPackages",
	}
}