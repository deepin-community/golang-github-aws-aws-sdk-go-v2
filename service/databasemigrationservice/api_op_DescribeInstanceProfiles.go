// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of instance profiles for your account in the current
// region.
func (c *Client) DescribeInstanceProfiles(ctx context.Context, params *DescribeInstanceProfilesInput, optFns ...func(*Options)) (*DescribeInstanceProfilesOutput, error) {
	if params == nil {
		params = &DescribeInstanceProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceProfiles", params, optFns, c.addOperationDescribeInstanceProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceProfilesInput struct {

	// Filters applied to the instance profiles described in the form of key-value
	// pairs.
	Filters []types.Filter

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords .
	//
	// If Marker is returned by a previous response, there are more results available.
	// The value of Marker is a unique pagination token for each page. To retrieve the
	// next page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, DMS includes a pagination token in the
	// response so that you can retrieve the remaining results.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeInstanceProfilesOutput struct {

	// A description of instance profiles.
	InstanceProfiles []types.InstanceProfile

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords .
	//
	// If Marker is returned by a previous response, there are more results available.
	// The value of Marker is a unique pagination token for each page. To retrieve the
	// next page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstanceProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstanceProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInstanceProfiles"); err != nil {
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
	if err = addOpDescribeInstanceProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceProfiles(options.Region), middleware.Before); err != nil {
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

// DescribeInstanceProfilesPaginatorOptions is the paginator options for
// DescribeInstanceProfiles
type DescribeInstanceProfilesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, DMS includes a pagination token in the
	// response so that you can retrieve the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstanceProfilesPaginator is a paginator for DescribeInstanceProfiles
type DescribeInstanceProfilesPaginator struct {
	options   DescribeInstanceProfilesPaginatorOptions
	client    DescribeInstanceProfilesAPIClient
	params    *DescribeInstanceProfilesInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstanceProfilesPaginator returns a new
// DescribeInstanceProfilesPaginator
func NewDescribeInstanceProfilesPaginator(client DescribeInstanceProfilesAPIClient, params *DescribeInstanceProfilesInput, optFns ...func(*DescribeInstanceProfilesPaginatorOptions)) *DescribeInstanceProfilesPaginator {
	if params == nil {
		params = &DescribeInstanceProfilesInput{}
	}

	options := DescribeInstanceProfilesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstanceProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstanceProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstanceProfiles page.
func (p *DescribeInstanceProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstanceProfilesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeInstanceProfiles(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeInstanceProfilesAPIClient is a client that implements the
// DescribeInstanceProfiles operation.
type DescribeInstanceProfilesAPIClient interface {
	DescribeInstanceProfiles(context.Context, *DescribeInstanceProfilesInput, ...func(*Options)) (*DescribeInstanceProfilesOutput, error)
}

var _ DescribeInstanceProfilesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeInstanceProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInstanceProfiles",
	}
}