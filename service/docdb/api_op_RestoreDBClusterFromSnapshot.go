// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new cluster from a snapshot or cluster snapshot.
//
// If a snapshot is specified, the target cluster is created from the source DB
// snapshot with a default configuration and default security group.
//
// If a cluster snapshot is specified, the target cluster is created from the
// source cluster restore point with the same configuration as the original source
// DB cluster, except that the new cluster is created with the default security
// group.
func (c *Client) RestoreDBClusterFromSnapshot(ctx context.Context, params *RestoreDBClusterFromSnapshotInput, optFns ...func(*Options)) (*RestoreDBClusterFromSnapshotOutput, error) {
	if params == nil {
		params = &RestoreDBClusterFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBClusterFromSnapshot", params, optFns, c.addOperationRestoreDBClusterFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBClusterFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to RestoreDBClusterFromSnapshot.
type RestoreDBClusterFromSnapshotInput struct {

	// The name of the cluster to create from the snapshot or cluster snapshot. This
	// parameter isn't case sensitive.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//   - The first character must be a letter.
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-snapshot-id
	//
	// This member is required.
	DBClusterIdentifier *string

	// The database engine to use for the new cluster.
	//
	// Default: The same as source.
	//
	// Constraint: Must be compatible with the engine of the source.
	//
	// This member is required.
	Engine *string

	// The identifier for the snapshot or cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a
	// cluster snapshot. However, you can use only the ARN to specify a snapshot.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing snapshot.
	//
	// This member is required.
	SnapshotIdentifier *string

	// Provides the list of Amazon EC2 Availability Zones that instances in the
	// restored DB cluster can be created in.
	AvailabilityZones []string

	// The name of the DB cluster parameter group to associate with this DB cluster.
	//
	// Type: String. Required: No.
	//
	// If this argument is omitted, the default DB cluster parameter group is used. If
	// supplied, must match the name of an existing default DB cluster parameter group.
	// The string must consist of from 1 to 255 letters, numbers or hyphens. Its first
	// character must be a letter, and it cannot end with a hyphen or contain two
	// consecutive hyphens.
	DBClusterParameterGroupName *string

	// The name of the subnet group to use for the new cluster.
	//
	// Constraints: If provided, must match the name of an existing DBSubnetGroup .
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string

	// Specifies whether this cluster can be deleted. If DeletionProtection is
	// enabled, the cluster cannot be deleted unless it is modified and
	// DeletionProtection is disabled. DeletionProtection protects clusters from being
	// accidentally deleted.
	DeletionProtection *bool

	// A list of log types that must be enabled for exporting to Amazon CloudWatch
	// Logs.
	EnableCloudwatchLogsExports []string

	// The version of the database engine to use for the new cluster.
	EngineVersion *string

	// The KMS key identifier to use when restoring an encrypted cluster from a DB
	// snapshot or cluster snapshot.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are restoring a cluster with the same Amazon Web Services account
	// that owns the KMS encryption key used to encrypt the new cluster, then you can
	// use the KMS key alias instead of the ARN for the KMS encryption key.
	//
	// If you do not specify a value for the KmsKeyId parameter, then the following
	// occurs:
	//
	//   - If the snapshot or cluster snapshot in SnapshotIdentifier is encrypted, then
	//   the restored cluster is encrypted using the KMS key that was used to encrypt the
	//   snapshot or the cluster snapshot.
	//
	//   - If the snapshot or the cluster snapshot in SnapshotIdentifier is not
	//   encrypted, then the restored DB cluster is not encrypted.
	KmsKeyId *string

	// The port number on which the new cluster accepts connections.
	//
	// Constraints: Must be a value from 1150 to 65535 .
	//
	// Default: The same port as the original cluster.
	Port *int32

	// The storage type to associate with the DB cluster.
	//
	// For information on storage types for Amazon DocumentDB clusters, see Cluster
	// storage configurations in the Amazon DocumentDB Developer Guide.
	//
	// Valid values for storage type - standard | iopt1
	//
	// Default value is standard
	StorageType *string

	// The tags to be assigned to the restored cluster.
	Tags []types.Tag

	// A list of virtual private cloud (VPC) security groups that the new cluster will
	// belong to.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBClusterFromSnapshotOutput struct {

	// Detailed information about a cluster.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBClusterFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreDBClusterFromSnapshot"); err != nil {
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
	if err = addOpRestoreDBClusterFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterFromSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBClusterFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreDBClusterFromSnapshot",
	}
}