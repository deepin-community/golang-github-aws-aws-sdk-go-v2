// Code generated by smithy-go-codegen DO NOT EDIT.

package docdbelastic

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdbelastic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon DocumentDB elastic cluster and returns its cluster
// structure.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	if params == nil {
		params = &CreateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCluster", params, optFns, c.addOperationCreateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// The name of the Amazon DocumentDB elastic clusters administrator.
	//
	// Constraints:
	//
	//   - Must be from 1 to 63 letters or numbers.
	//
	//   - The first character must be a letter.
	//
	//   - Cannot be a reserved word.
	//
	// This member is required.
	AdminUserName *string

	// The password for the Amazon DocumentDB elastic clusters administrator. The
	// password can contain any printable ASCII characters.
	//
	// Constraints:
	//
	//   - Must contain from 8 to 100 characters.
	//
	//   - Cannot contain a forward slash (/), double quote ("), or the "at" symbol
	//   (@).
	//
	// This member is required.
	AdminUserPassword *string

	// The authentication type used to determine where to fetch the password used for
	// accessing the elastic cluster. Valid types are PLAIN_TEXT or SECRET_ARN .
	//
	// This member is required.
	AuthType types.Auth

	// The name of the new elastic cluster. This parameter is stored as a lowercase
	// string.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//   - The first character must be a letter.
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster
	//
	// This member is required.
	ClusterName *string

	// The number of vCPUs assigned to each elastic cluster shard. Maximum is 64.
	// Allowed values are 2, 4, 8, 16, 32, 64.
	//
	// This member is required.
	ShardCapacity *int32

	// The number of shards assigned to the elastic cluster. Maximum is 32.
	//
	// This member is required.
	ShardCount *int32

	// The number of days for which automatic snapshots are retained.
	BackupRetentionPeriod *int32

	// The client token for the elastic cluster.
	ClientToken *string

	// The KMS key identifier to use to encrypt the new elastic cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are creating a cluster using the same Amazon account that owns this
	// KMS encryption key, you can use the KMS key alias instead of the ARN as the KMS
	// encryption key.
	//
	// If an encryption key is not specified, Amazon DocumentDB uses the default
	// encryption key that KMS creates for your account. Your account has a different
	// default encryption key for each Amazon Region.
	KmsKeyId *string

	// The daily time range during which automated backups are created if automated
	// backups are enabled, as determined by the backupRetentionPeriod .
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// Default: a 30-minute window selected at random from an 8-hour block of time for
	// each Amazon Web Services Region, occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// The number of replica instances applying to all shards in the elastic cluster.
	// A shardInstanceCount value of 1 means there is one writer instance, and any
	// additional instances are replicas that can be used for reads and to improve
	// availability.
	ShardInstanceCount *int32

	// The Amazon EC2 subnet IDs for the new elastic cluster.
	SubnetIds []string

	// The tags to be assigned to the new elastic cluster.
	Tags map[string]string

	// A list of EC2 VPC security groups to associate with the new elastic cluster.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type CreateClusterOutput struct {

	// The new elastic cluster that has been created.
	//
	// This member is required.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCluster"); err != nil {
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
	if err = addIdempotencyToken_opCreateClusterMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCluster struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCluster) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateClusterInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateClusterMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCluster{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCluster",
	}
}