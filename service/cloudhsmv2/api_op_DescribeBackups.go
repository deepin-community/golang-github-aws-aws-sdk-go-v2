// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about backups of CloudHSM clusters. Lists either the backups
// you own or the backups shared with you when the Shared parameter is true.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the backups. When the response contains only a subset of
// backups, it includes a NextToken value. Use this value in a subsequent
// DescribeBackups request to get more backups. When you receive a response with no
// NextToken (or an empty or null value), that means there are no more backups to
// get.
//
// Cross-account use: Yes. Customers can describe backups in other Amazon Web
// Services accounts that are shared with them.
func (c *Client) DescribeBackups(ctx context.Context, params *DescribeBackupsInput, optFns ...func(*Options)) (*DescribeBackupsOutput, error) {
	if params == nil {
		params = &DescribeBackupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBackups", params, optFns, c.addOperationDescribeBackupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupsInput struct {

	// One or more filters to limit the items returned in the response.
	//
	// Use the backupIds filter to return only the specified backups. Specify backups
	// by their backup identifier (ID).
	//
	// Use the sourceBackupIds filter to return only the backups created from a source
	// backup. The sourceBackupID of a source backup is returned by the CopyBackupToRegion operation.
	//
	// Use the clusterIds filter to return only the backups for the specified
	// clusters. Specify clusters by their cluster identifier (ID).
	//
	// Use the states filter to return only backups that match the specified state.
	//
	// Use the neverExpires filter to return backups filtered by the value in the
	// neverExpires parameter. True returns all backups exempt from the backup
	// retention policy. False returns all backups with a backup retention policy
	// defined at the cluster.
	Filters map[string][]string

	// The maximum number of backups to return in the response. When there are more
	// backups than the number you specify, the response contains a NextToken value.
	MaxResults *int32

	// The NextToken value that you received in the previous response. Use this value
	// to get more backups.
	NextToken *string

	// Describe backups that are shared with you.
	//
	// By default when using this option, the command returns backups that have been
	// shared using a standard Resource Access Manager resource share. In order for a
	// backup that was shared using the PutResourcePolicy command to be returned, the
	// share must be promoted to a standard resource share using the RAM [PromoteResourceShareCreatedFromPolicy]API
	// operation.
	//
	// For more information about sharing backups, see [Working with shared backups] in the CloudHSM User Guide.
	//
	// [Working with shared backups]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/sharing.html
	// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/cli/latest/reference/ram/promote-resource-share-created-from-policy.html
	Shared *bool

	// Designates whether or not to sort the return backups by ascending chronological
	// order of generation.
	SortAscending *bool

	noSmithyDocumentSerde
}

type DescribeBackupsOutput struct {

	// A list of backups.
	Backups []types.Backup

	// An opaque string that indicates that the response contains only a subset of
	// backups. Use this value in a subsequent DescribeBackups request to get more
	// backups.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBackupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBackups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBackups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBackups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackups(options.Region), middleware.Before); err != nil {
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

// DescribeBackupsPaginatorOptions is the paginator options for DescribeBackups
type DescribeBackupsPaginatorOptions struct {
	// The maximum number of backups to return in the response. When there are more
	// backups than the number you specify, the response contains a NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBackupsPaginator is a paginator for DescribeBackups
type DescribeBackupsPaginator struct {
	options   DescribeBackupsPaginatorOptions
	client    DescribeBackupsAPIClient
	params    *DescribeBackupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeBackupsPaginator returns a new DescribeBackupsPaginator
func NewDescribeBackupsPaginator(client DescribeBackupsAPIClient, params *DescribeBackupsInput, optFns ...func(*DescribeBackupsPaginatorOptions)) *DescribeBackupsPaginator {
	if params == nil {
		params = &DescribeBackupsInput{}
	}

	options := DescribeBackupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBackupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBackupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBackups page.
func (p *DescribeBackupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBackupsOutput, error) {
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
	result, err := p.client.DescribeBackups(ctx, &params, optFns...)
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

// DescribeBackupsAPIClient is a client that implements the DescribeBackups
// operation.
type DescribeBackupsAPIClient interface {
	DescribeBackups(context.Context, *DescribeBackupsInput, ...func(*Options)) (*DescribeBackupsOutput, error)
}

var _ DescribeBackupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeBackups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBackups",
	}
}