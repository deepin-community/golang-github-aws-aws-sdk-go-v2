// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Amazon DataZone environment profiles.
func (c *Client) ListEnvironmentProfiles(ctx context.Context, params *ListEnvironmentProfilesInput, optFns ...func(*Options)) (*ListEnvironmentProfilesOutput, error) {
	if params == nil {
		params = &ListEnvironmentProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironmentProfiles", params, optFns, c.addOperationListEnvironmentProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentProfilesInput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the Amazon Web Services account where you want to list
	// environment profiles.
	AwsAccountId *string

	// The Amazon Web Services region where you want to list environment profiles.
	AwsAccountRegion *string

	// The identifier of the blueprint that was used to create the environment
	// profiles that you want to list.
	EnvironmentBlueprintIdentifier *string

	// The maximum number of environment profiles to return in a single call to
	// ListEnvironmentProfiles . When the number of environment profiles to be listed
	// is greater than the value of MaxResults , the response contains a NextToken
	// value that you can use in a subsequent call to ListEnvironmentProfiles to list
	// the next set of environment profiles.
	MaxResults *int32

	//
	Name *string

	// When the number of environment profiles is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of environment profiles, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListEnvironmentProfiles to list the next set of environment
	// profiles.
	NextToken *string

	// The identifier of the Amazon DataZone project.
	ProjectIdentifier *string

	noSmithyDocumentSerde
}

type ListEnvironmentProfilesOutput struct {

	// The results of the ListEnvironmentProfiles action.
	//
	// This member is required.
	Items []types.EnvironmentProfileSummary

	// When the number of environment profiles is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of environment profiles, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListEnvironmentProfiles to list the next set of environment
	// profiles.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnvironmentProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEnvironmentProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEnvironmentProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEnvironmentProfiles"); err != nil {
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
	if err = addOpListEnvironmentProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnvironmentProfiles(options.Region), middleware.Before); err != nil {
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

// ListEnvironmentProfilesPaginatorOptions is the paginator options for
// ListEnvironmentProfiles
type ListEnvironmentProfilesPaginatorOptions struct {
	// The maximum number of environment profiles to return in a single call to
	// ListEnvironmentProfiles . When the number of environment profiles to be listed
	// is greater than the value of MaxResults , the response contains a NextToken
	// value that you can use in a subsequent call to ListEnvironmentProfiles to list
	// the next set of environment profiles.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentProfilesPaginator is a paginator for ListEnvironmentProfiles
type ListEnvironmentProfilesPaginator struct {
	options   ListEnvironmentProfilesPaginatorOptions
	client    ListEnvironmentProfilesAPIClient
	params    *ListEnvironmentProfilesInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentProfilesPaginator returns a new
// ListEnvironmentProfilesPaginator
func NewListEnvironmentProfilesPaginator(client ListEnvironmentProfilesAPIClient, params *ListEnvironmentProfilesInput, optFns ...func(*ListEnvironmentProfilesPaginatorOptions)) *ListEnvironmentProfilesPaginator {
	if params == nil {
		params = &ListEnvironmentProfilesInput{}
	}

	options := ListEnvironmentProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnvironmentProfiles page.
func (p *ListEnvironmentProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentProfilesOutput, error) {
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
	result, err := p.client.ListEnvironmentProfiles(ctx, &params, optFns...)
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

// ListEnvironmentProfilesAPIClient is a client that implements the
// ListEnvironmentProfiles operation.
type ListEnvironmentProfilesAPIClient interface {
	ListEnvironmentProfiles(context.Context, *ListEnvironmentProfilesInput, ...func(*Options)) (*ListEnvironmentProfilesOutput, error)
}

var _ ListEnvironmentProfilesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListEnvironmentProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEnvironmentProfiles",
	}
}