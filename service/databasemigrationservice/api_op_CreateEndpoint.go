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

// Creates an endpoint using the provided settings.
//
// For a MySQL source or target endpoint, don't explicitly specify the database
// using the DatabaseName request parameter on the CreateEndpoint API call.
// Specifying DatabaseName when you create a MySQL endpoint replicates all the
// task tables to this single database. For MySQL endpoints, you specify the
// database only when you specify the schema in the table-mapping rules of the DMS
// task.
func (c *Client) CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) {
	if params == nil {
		params = &CreateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEndpoint", params, optFns, c.addOperationCreateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEndpointInput struct {

	// The database endpoint identifier. Identifiers must begin with a letter and must
	// contain only ASCII letters, digits, and hyphens. They can't end with a hyphen,
	// or contain two consecutive hyphens.
	//
	// This member is required.
	EndpointIdentifier *string

	// The type of endpoint. Valid values are source and target .
	//
	// This member is required.
	EndpointType types.ReplicationEndpointTypeValue

	// The type of engine for the endpoint. Valid values, depending on the EndpointType
	// value, include "mysql" , "oracle" , "postgres" , "mariadb" , "aurora" ,
	// "aurora-postgresql" , "opensearch" , "redshift" , "s3" , "db2" , "db2-zos" ,
	// "azuredb" , "sybase" , "dynamodb" , "mongodb" , "kinesis" , "kafka" ,
	// "elasticsearch" , "docdb" , "sqlserver" , "neptune" , and "babelfish" .
	//
	// This member is required.
	EngineName *string

	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn *string

	// The name of the endpoint database. For a MySQL source or target endpoint, do
	// not specify DatabaseName. To migrate to a specific database, use this setting
	// and targetDbType .
	DatabaseName *string

	// The settings in JSON format for the DMS transfer type of source endpoint.
	//
	// Possible settings include the following:
	//
	//   - ServiceAccessRoleArn - The Amazon Resource Name (ARN) used by the service
	//   access IAM role. The role must allow the iam:PassRole action.
	//
	//   - BucketName - The name of the S3 bucket to use.
	//
	// Shorthand syntax for these settings is as follows:
	// ServiceAccessRoleArn=string,BucketName=string
	//
	// JSON syntax for these settings is as follows: { "ServiceAccessRoleArn":
	// "string", "BucketName": "string", }
	DmsTransferSettings *types.DmsTransferSettings

	// Provides information that defines a DocumentDB endpoint.
	DocDbSettings *types.DocDbSettings

	// Settings in JSON format for the target Amazon DynamoDB endpoint. For
	// information about other available settings, see [Using Object Mapping to Migrate Data to DynamoDB]in the Database Migration
	// Service User Guide.
	//
	// [Using Object Mapping to Migrate Data to DynamoDB]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping
	DynamoDbSettings *types.DynamoDbSettings

	// Settings in JSON format for the target OpenSearch endpoint. For more
	// information about the available settings, see [Extra Connection Attributes When Using OpenSearch as a Target for DMS]in the Database Migration Service
	// User Guide.
	//
	// [Extra Connection Attributes When Using OpenSearch as a Target for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration
	ElasticsearchSettings *types.ElasticsearchSettings

	// The external table definition.
	ExternalTableDefinition *string

	// Additional attributes associated with the connection. Each attribute is
	// specified as a name-value pair associated by an equal sign (=). Multiple
	// attributes are separated by a semicolon (;) with no additional white space. For
	// information on the attributes available for connecting your source or target
	// endpoint, see [Working with DMS Endpoints]in the Database Migration Service User Guide.
	//
	// [Working with DMS Endpoints]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Endpoints.html
	ExtraConnectionAttributes *string

	// Settings in JSON format for the source GCP MySQL endpoint.
	GcpMySQLSettings *types.GcpMySQLSettings

	// Settings in JSON format for the source IBM Db2 LUW endpoint. For information
	// about other available settings, see [Extra connection attributes when using Db2 LUW as a source for DMS]in the Database Migration Service User
	// Guide.
	//
	// [Extra connection attributes when using Db2 LUW as a source for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib
	IBMDb2Settings *types.IBMDb2Settings

	// Settings in JSON format for the target Apache Kafka endpoint. For more
	// information about the available settings, see [Using object mapping to migrate data to a Kafka topic]in the Database Migration Service
	// User Guide.
	//
	// [Using object mapping to migrate data to a Kafka topic]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping
	KafkaSettings *types.KafkaSettings

	// Settings in JSON format for the target endpoint for Amazon Kinesis Data
	// Streams. For more information about the available settings, see [Using object mapping to migrate data to a Kinesis data stream]in the Database
	// Migration Service User Guide.
	//
	// [Using object mapping to migrate data to a Kinesis data stream]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping
	KinesisSettings *types.KinesisSettings

	// An KMS key identifier that is used to encrypt the connection parameters for the
	// endpoint.
	//
	// If you don't specify a value for the KmsKeyId parameter, then DMS uses your
	// default encryption key.
	//
	// KMS creates the default encryption key for your Amazon Web Services account.
	// Your Amazon Web Services account has a different default encryption key for each
	// Amazon Web Services Region.
	KmsKeyId *string

	// Settings in JSON format for the source and target Microsoft SQL Server
	// endpoint. For information about other available settings, see [Extra connection attributes when using SQL Server as a source for DMS]and [Extra connection attributes when using SQL Server as a target for DMS] in the
	// Database Migration Service User Guide.
	//
	// [Extra connection attributes when using SQL Server as a source for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib
	// [Extra connection attributes when using SQL Server as a target for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib
	MicrosoftSQLServerSettings *types.MicrosoftSQLServerSettings

	// Settings in JSON format for the source MongoDB endpoint. For more information
	// about the available settings, see [Endpoint configuration settings when using MongoDB as a source for Database Migration Service]in the Database Migration Service User Guide.
	//
	// [Endpoint configuration settings when using MongoDB as a source for Database Migration Service]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration
	MongoDbSettings *types.MongoDbSettings

	// Settings in JSON format for the source and target MySQL endpoint. For
	// information about other available settings, see [Extra connection attributes when using MySQL as a source for DMS]and [Extra connection attributes when using a MySQL-compatible database as a target for DMS] in the Database Migration
	// Service User Guide.
	//
	// [Extra connection attributes when using MySQL as a source for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib
	// [Extra connection attributes when using a MySQL-compatible database as a target for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib
	MySQLSettings *types.MySQLSettings

	// Settings in JSON format for the target Amazon Neptune endpoint. For more
	// information about the available settings, see [Specifying graph-mapping rules using Gremlin and R2RML for Amazon Neptune as a target]in the Database Migration Service
	// User Guide.
	//
	// [Specifying graph-mapping rules using Gremlin and R2RML for Amazon Neptune as a target]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings
	NeptuneSettings *types.NeptuneSettings

	// Settings in JSON format for the source and target Oracle endpoint. For
	// information about other available settings, see [Extra connection attributes when using Oracle as a source for DMS]and [Extra connection attributes when using Oracle as a target for DMS] in the Database Migration
	// Service User Guide.
	//
	// [Extra connection attributes when using Oracle as a target for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib
	// [Extra connection attributes when using Oracle as a source for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib
	OracleSettings *types.OracleSettings

	// The password to be used to log in to the endpoint database.
	Password *string

	// The port used by the endpoint database.
	Port *int32

	// Settings in JSON format for the source and target PostgreSQL endpoint. For
	// information about other available settings, see [Extra connection attributes when using PostgreSQL as a source for DMS]and [Extra connection attributes when using PostgreSQL as a target for DMS] in the Database Migration
	// Service User Guide.
	//
	// [Extra connection attributes when using PostgreSQL as a source for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib
	// [Extra connection attributes when using PostgreSQL as a target for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html#CHAP_Target.PostgreSQL.ConnectionAttrib
	PostgreSQLSettings *types.PostgreSQLSettings

	// Settings in JSON format for the target Redis endpoint.
	RedisSettings *types.RedisSettings

	// Provides information that defines an Amazon Redshift endpoint.
	RedshiftSettings *types.RedshiftSettings

	// A friendly name for the resource identifier at the end of the EndpointArn
	// response parameter that is returned in the created Endpoint object. The value
	// for this parameter can have up to 31 characters. It can contain only ASCII
	// letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain
	// two consecutive hyphens, and can only begin with a letter, such as
	// Example-App-ARN1 . For example, this value might result in the EndpointArn
	// value arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1 . If you don't
	// specify a ResourceIdentifier value, DMS generates a default identifier value
	// for the end of EndpointArn .
	ResourceIdentifier *string

	// Settings in JSON format for the target Amazon S3 endpoint. For more information
	// about the available settings, see [Extra Connection Attributes When Using Amazon S3 as a Target for DMS]in the Database Migration Service User Guide.
	//
	// [Extra Connection Attributes When Using Amazon S3 as a Target for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring
	S3Settings *types.S3Settings

	// The name of the server where the endpoint database resides.
	ServerName *string

	//  The Amazon Resource Name (ARN) for the service access role that you want to
	// use to create the endpoint. The role must allow the iam:PassRole action.
	ServiceAccessRoleArn *string

	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default
	// is none
	SslMode types.DmsSslModeValue

	// Settings in JSON format for the source and target SAP ASE endpoint. For
	// information about other available settings, see [Extra connection attributes when using SAP ASE as a source for DMS]and [Extra connection attributes when using SAP ASE as a target for DMS] in the Database Migration
	// Service User Guide.
	//
	// [Extra connection attributes when using SAP ASE as a source for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html#CHAP_Source.SAP.ConnectionAttrib
	// [Extra connection attributes when using SAP ASE as a target for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html#CHAP_Target.SAP.ConnectionAttrib
	SybaseSettings *types.SybaseSettings

	// One or more tags to be assigned to the endpoint.
	Tags []types.Tag

	// Settings in JSON format for the target Amazon Timestream endpoint.
	TimestreamSettings *types.TimestreamSettings

	// The user name to be used to log in to the endpoint database.
	Username *string

	noSmithyDocumentSerde
}

type CreateEndpointOutput struct {

	// The endpoint that was created.
	Endpoint *types.Endpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEndpoint"); err != nil {
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
	if err = addOpCreateEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEndpoint",
	}
}