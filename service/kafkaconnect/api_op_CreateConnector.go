// Code generated by smithy-go-codegen DO NOT EDIT.

package kafkaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a connector using the specified properties.
func (c *Client) CreateConnector(ctx context.Context, params *CreateConnectorInput, optFns ...func(*Options)) (*CreateConnectorOutput, error) {
	if params == nil {
		params = &CreateConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConnector", params, optFns, c.addOperationCreateConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectorInput struct {

	// Information about the capacity allocated to the connector. Exactly one of the
	// two properties must be specified.
	//
	// This member is required.
	Capacity *types.Capacity

	// A map of keys to values that represent the configuration for the connector.
	//
	// This member is required.
	ConnectorConfiguration map[string]string

	// The name of the connector.
	//
	// This member is required.
	ConnectorName *string

	// Specifies which Apache Kafka cluster to connect to.
	//
	// This member is required.
	KafkaCluster *types.KafkaCluster

	// Details of the client authentication used by the Apache Kafka cluster.
	//
	// This member is required.
	KafkaClusterClientAuthentication *types.KafkaClusterClientAuthentication

	// Details of encryption in transit to the Apache Kafka cluster.
	//
	// This member is required.
	KafkaClusterEncryptionInTransit *types.KafkaClusterEncryptionInTransit

	// The version of Kafka Connect. It has to be compatible with both the Apache
	// Kafka cluster's version and the plugins.
	//
	// This member is required.
	KafkaConnectVersion *string

	// Amazon MSK Connect does not currently support specifying multiple plugins as a
	// list. To use more than one plugin for your connector, you can create a single
	// custom plugin using a ZIP file that bundles multiple plugins together.
	//
	// Specifies which plugin to use for the connector. You must specify a
	// single-element list containing one customPlugin object.
	//
	// This member is required.
	Plugins []types.Plugin

	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access
	// the Amazon Web Services resources that it needs. The types of resources depends
	// on the logic of the connector. For example, a connector that has Amazon S3 as a
	// destination must have permissions that allow it to write to the S3 destination
	// bucket.
	//
	// This member is required.
	ServiceExecutionRoleArn *string

	// A summary description of the connector.
	ConnectorDescription *string

	// Details about log delivery.
	LogDelivery *types.LogDelivery

	// The tags you want to attach to the connector.
	Tags map[string]string

	// Specifies which worker configuration to use with the connector.
	WorkerConfiguration *types.WorkerConfiguration

	noSmithyDocumentSerde
}

type CreateConnectorOutput struct {

	// The Amazon Resource Name (ARN) that Amazon assigned to the connector.
	ConnectorArn *string

	// The name of the connector.
	ConnectorName *string

	// The state of the connector.
	ConnectorState types.ConnectorState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConnector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateConnector"); err != nil {
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
	if err = addOpCreateConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateConnector",
	}
}