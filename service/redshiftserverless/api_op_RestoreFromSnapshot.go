// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Restores a namespace from a snapshot.
func (c *Client) RestoreFromSnapshot(ctx context.Context, params *RestoreFromSnapshotInput, optFns ...func(*Options)) (*RestoreFromSnapshotOutput, error) {
	if params == nil {
		params = &RestoreFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreFromSnapshot", params, optFns, c.addOperationRestoreFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreFromSnapshotInput struct {

	// The name of the namespace to restore the snapshot to.
	//
	// This member is required.
	NamespaceName *string

	// The name of the workgroup used to restore the snapshot.
	//
	// This member is required.
	WorkgroupName *string

	// The ID of the Key Management Service (KMS) key used to encrypt and store the
	// namespace's admin credentials secret.
	AdminPasswordSecretKmsKeyId *string

	// If true , Amazon Redshift uses Secrets Manager to manage the restored snapshot's
	// admin credentials. If MmanageAdminPassword is false or not set, Amazon Redshift
	// uses the admin credentials that the namespace or cluster had at the time the
	// snapshot was taken.
	ManageAdminPassword *bool

	// The Amazon Web Services account that owns the snapshot.
	OwnerAccount *string

	// The Amazon Resource Name (ARN) of the snapshot to restore from. Required if
	// restoring from Amazon Redshift Serverless to a provisioned cluster. Must not be
	// specified at the same time as snapshotName .
	//
	// The format of the ARN is
	// arn:aws:redshift:<region>:<account_id>:snapshot:<cluster_identifier>/<snapshot_identifier>.
	SnapshotArn *string

	// The name of the snapshot to restore from. Must not be specified at the same
	// time as snapshotArn .
	SnapshotName *string

	noSmithyDocumentSerde
}

type RestoreFromSnapshotOutput struct {

	// A collection of database objects and users.
	Namespace *types.Namespace

	// The owner Amazon Web Services; account of the snapshot that was restored.
	OwnerAccount *string

	// The name of the snapshot used to restore the namespace.
	SnapshotName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreFromSnapshot"); err != nil {
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
	if err = addOpRestoreFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreFromSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreFromSnapshot",
	}
}