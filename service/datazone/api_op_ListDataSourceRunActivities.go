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

// Lists data source run activities.
func (c *Client) ListDataSourceRunActivities(ctx context.Context, params *ListDataSourceRunActivitiesInput, optFns ...func(*Options)) (*ListDataSourceRunActivitiesOutput, error) {
	if params == nil {
		params = &ListDataSourceRunActivitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSourceRunActivities", params, optFns, c.addOperationListDataSourceRunActivitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSourceRunActivitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSourceRunActivitiesInput struct {

	// The identifier of the Amazon DataZone domain in which to list data source run
	// activities.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the data source run.
	//
	// This member is required.
	Identifier *string

	// The maximum number of activities to return in a single call to
	// ListDataSourceRunActivities . When the number of activities to be listed is
	// greater than the value of MaxResults , the response contains a NextToken value
	// that you can use in a subsequent call to ListDataSourceRunActivities to list
	// the next set of activities.
	MaxResults *int32

	// When the number of activities is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of activities, the response includes a pagination token
	// named NextToken . You can specify this NextToken value in a subsequent call to
	// ListDataSourceRunActivities to list the next set of activities.
	NextToken *string

	// The status of the data source run.
	Status types.DataAssetActivityStatus

	noSmithyDocumentSerde
}

type ListDataSourceRunActivitiesOutput struct {

	// The results of the ListDataSourceRunActivities action.
	//
	// This member is required.
	Items []types.DataSourceRunActivity

	// When the number of activities is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of activities, the response includes a pagination token
	// named NextToken . You can specify this NextToken value in a subsequent call to
	// ListDataSourceRunActivities to list the next set of activities.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataSourceRunActivitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataSourceRunActivities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSourceRunActivities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataSourceRunActivities"); err != nil {
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
	if err = addOpListDataSourceRunActivitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSourceRunActivities(options.Region), middleware.Before); err != nil {
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

// ListDataSourceRunActivitiesPaginatorOptions is the paginator options for
// ListDataSourceRunActivities
type ListDataSourceRunActivitiesPaginatorOptions struct {
	// The maximum number of activities to return in a single call to
	// ListDataSourceRunActivities . When the number of activities to be listed is
	// greater than the value of MaxResults , the response contains a NextToken value
	// that you can use in a subsequent call to ListDataSourceRunActivities to list
	// the next set of activities.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataSourceRunActivitiesPaginator is a paginator for
// ListDataSourceRunActivities
type ListDataSourceRunActivitiesPaginator struct {
	options   ListDataSourceRunActivitiesPaginatorOptions
	client    ListDataSourceRunActivitiesAPIClient
	params    *ListDataSourceRunActivitiesInput
	nextToken *string
	firstPage bool
}

// NewListDataSourceRunActivitiesPaginator returns a new
// ListDataSourceRunActivitiesPaginator
func NewListDataSourceRunActivitiesPaginator(client ListDataSourceRunActivitiesAPIClient, params *ListDataSourceRunActivitiesInput, optFns ...func(*ListDataSourceRunActivitiesPaginatorOptions)) *ListDataSourceRunActivitiesPaginator {
	if params == nil {
		params = &ListDataSourceRunActivitiesInput{}
	}

	options := ListDataSourceRunActivitiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataSourceRunActivitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataSourceRunActivitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataSourceRunActivities page.
func (p *ListDataSourceRunActivitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataSourceRunActivitiesOutput, error) {
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
	result, err := p.client.ListDataSourceRunActivities(ctx, &params, optFns...)
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

// ListDataSourceRunActivitiesAPIClient is a client that implements the
// ListDataSourceRunActivities operation.
type ListDataSourceRunActivitiesAPIClient interface {
	ListDataSourceRunActivities(context.Context, *ListDataSourceRunActivitiesInput, ...func(*Options)) (*ListDataSourceRunActivitiesOutput, error)
}

var _ ListDataSourceRunActivitiesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDataSourceRunActivities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataSourceRunActivities",
	}
}