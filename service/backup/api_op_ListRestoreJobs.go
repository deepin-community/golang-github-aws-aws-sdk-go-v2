// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of jobs that Backup initiated to restore a saved resource,
// including details about the recovery process.
func (c *Client) ListRestoreJobs(ctx context.Context, params *ListRestoreJobsInput, optFns ...func(*Options)) (*ListRestoreJobsOutput, error) {
	if params == nil {
		params = &ListRestoreJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRestoreJobs", params, optFns, c.addOperationListRestoreJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRestoreJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRestoreJobsInput struct {

	// The account ID to list the jobs from. Returns only restore jobs associated with
	// the specified account ID.
	ByAccountId *string

	// Returns only copy jobs completed after a date expressed in Unix format and
	// Coordinated Universal Time (UTC).
	ByCompleteAfter *time.Time

	// Returns only copy jobs completed before a date expressed in Unix format and
	// Coordinated Universal Time (UTC).
	ByCompleteBefore *time.Time

	// Returns only restore jobs that were created after the specified date.
	ByCreatedAfter *time.Time

	// Returns only restore jobs that were created before the specified date.
	ByCreatedBefore *time.Time

	// Include this parameter to return only restore jobs for the specified resources:
	//
	//   - Aurora for Amazon Aurora
	//
	//   - CloudFormation for CloudFormation
	//
	//   - DocumentDB for Amazon DocumentDB (with MongoDB compatibility)
	//
	//   - DynamoDB for Amazon DynamoDB
	//
	//   - EBS for Amazon Elastic Block Store
	//
	//   - EC2 for Amazon Elastic Compute Cloud
	//
	//   - EFS for Amazon Elastic File System
	//
	//   - FSx for Amazon FSx
	//
	//   - Neptune for Amazon Neptune
	//
	//   - Redshift for Amazon Redshift
	//
	//   - RDS for Amazon Relational Database Service
	//
	//   - SAP HANA on Amazon EC2 for SAP HANA databases
	//
	//   - Storage Gateway for Storage Gateway
	//
	//   - S3 for Amazon S3
	//
	//   - Timestream for Amazon Timestream
	//
	//   - VirtualMachine for virtual machines
	ByResourceType *string

	// This returns only restore testing jobs that match the specified resource Amazon
	// Resource Name (ARN).
	ByRestoreTestingPlanArn *string

	// Returns only restore jobs associated with the specified job status.
	ByStatus types.RestoreJobStatus

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRestoreJobsOutput struct {

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// An array of objects that contain detailed information about jobs to restore
	// saved resources.
	RestoreJobs []types.RestoreJobsListMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRestoreJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRestoreJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRestoreJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRestoreJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRestoreJobs(options.Region), middleware.Before); err != nil {
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

// ListRestoreJobsPaginatorOptions is the paginator options for ListRestoreJobs
type ListRestoreJobsPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRestoreJobsPaginator is a paginator for ListRestoreJobs
type ListRestoreJobsPaginator struct {
	options   ListRestoreJobsPaginatorOptions
	client    ListRestoreJobsAPIClient
	params    *ListRestoreJobsInput
	nextToken *string
	firstPage bool
}

// NewListRestoreJobsPaginator returns a new ListRestoreJobsPaginator
func NewListRestoreJobsPaginator(client ListRestoreJobsAPIClient, params *ListRestoreJobsInput, optFns ...func(*ListRestoreJobsPaginatorOptions)) *ListRestoreJobsPaginator {
	if params == nil {
		params = &ListRestoreJobsInput{}
	}

	options := ListRestoreJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRestoreJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRestoreJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRestoreJobs page.
func (p *ListRestoreJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRestoreJobsOutput, error) {
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
	result, err := p.client.ListRestoreJobs(ctx, &params, optFns...)
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

// ListRestoreJobsAPIClient is a client that implements the ListRestoreJobs
// operation.
type ListRestoreJobsAPIClient interface {
	ListRestoreJobs(context.Context, *ListRestoreJobsInput, ...func(*Options)) (*ListRestoreJobsOutput, error)
}

var _ ListRestoreJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRestoreJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRestoreJobs",
	}
}