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

// Creates a new instance.
func (c *Client) CreateDBInstance(ctx context.Context, params *CreateDBInstanceInput, optFns ...func(*Options)) (*CreateDBInstanceOutput, error) {
	if params == nil {
		params = &CreateDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBInstance", params, optFns, c.addOperationCreateDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to CreateDBInstance.
type CreateDBInstanceInput struct {

	// The identifier of the cluster that the instance will belong to.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The compute and memory capacity of the instance; for example, db.r5.large .
	//
	// This member is required.
	DBInstanceClass *string

	// The instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//   - The first character must be a letter.
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: mydbinstance
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The name of the database engine to be used for this instance.
	//
	// Valid value: docdb
	//
	// This member is required.
	Engine *string

	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not
	// perform minor version upgrades regardless of the value set.
	//
	// Default: false
	AutoMinorVersionUpgrade *bool

	// The Amazon EC2 Availability Zone that the instance is created in.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's Amazon Web
	// Services Region.
	//
	// Example: us-east-1d
	AvailabilityZone *string

	// The CA certificate identifier to use for the DB instance's server certificate.
	//
	// For more information, see [Updating Your Amazon DocumentDB TLS Certificates] and [Encrypting Data in Transit] in the Amazon DocumentDB Developer Guide.
	//
	// [Updating Your Amazon DocumentDB TLS Certificates]: https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html
	// [Encrypting Data in Transit]: https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html
	CACertificateIdentifier *string

	// A value that indicates whether to copy tags from the DB instance to snapshots
	// of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool

	// A value that indicates whether to enable Performance Insights for the DB
	// Instance. For more information, see [Using Amazon Performance Insights].
	//
	// [Using Amazon Performance Insights]: https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html
	EnablePerformanceInsights *bool

	// The KMS key identifier for encryption of Performance Insights data.
	//
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the
	// KMS key.
	//
	// If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon
	// DocumentDB uses your default KMS key. There is a default KMS key for your Amazon
	// Web Services account. Your Amazon Web Services account has a different default
	// KMS key for each Amazon Web Services region.
	PerformanceInsightsKMSKeyId *string

	// The time range each week during which system maintenance can occur, in
	// Universal Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block of
	// time for each Amazon Web Services Region, occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// A value that specifies the order in which an Amazon DocumentDB replica is
	// promoted to the primary instance after a failure of the existing primary
	// instance.
	//
	// Default: 1
	//
	// Valid values: 0-15
	PromotionTier *int32

	// The tags to be assigned to the instance. You can assign up to 10 tags to an
	// instance.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDBInstanceOutput struct {

	// Detailed information about an instance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDBInstance"); err != nil {
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
	if err = addOpCreateDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDBInstance",
	}
}