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

// Creates a namespace in Amazon Redshift Serverless.
func (c *Client) CreateNamespace(ctx context.Context, params *CreateNamespaceInput, optFns ...func(*Options)) (*CreateNamespaceOutput, error) {
	if params == nil {
		params = &CreateNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNamespace", params, optFns, c.addOperationCreateNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNamespaceInput struct {

	// The name of the namespace.
	//
	// This member is required.
	NamespaceName *string

	// The ID of the Key Management Service (KMS) key used to encrypt and store the
	// namespace's admin credentials secret. You can only use this parameter if
	// manageAdminPassword is true.
	AdminPasswordSecretKmsKeyId *string

	// The password of the administrator for the first database created in the
	// namespace.
	//
	// You can't use adminUserPassword if manageAdminPassword is true.
	AdminUserPassword *string

	// The username of the administrator for the first database created in the
	// namespace.
	AdminUsername *string

	// The name of the first database created in the namespace.
	DbName *string

	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the
	// namespace.
	DefaultIamRoleArn *string

	// A list of IAM roles to associate with the namespace.
	IamRoles []string

	// The ID of the Amazon Web Services Key Management Service key used to encrypt
	// your data.
	KmsKeyId *string

	// The types of logs the namespace can export. Available export types are userlog ,
	// connectionlog , and useractivitylog .
	LogExports []types.LogExport

	// If true , Amazon Redshift uses Secrets Manager to manage the namespace's admin
	// credentials. You can't use adminUserPassword if manageAdminPassword is true. If
	// manageAdminPassword is false or not set, Amazon Redshift uses adminUserPassword
	// for the admin user account's password.
	ManageAdminPassword *bool

	// The ARN for the Redshift application that integrates with IAM Identity Center.
	RedshiftIdcApplicationArn *string

	// A list of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateNamespaceOutput struct {

	// The created namespace object.
	Namespace *types.Namespace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNamespace"); err != nil {
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
	if err = addOpCreateNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNamespace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNamespace",
	}
}