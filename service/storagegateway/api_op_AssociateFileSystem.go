// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associate an Amazon FSx file system with the FSx File Gateway. After the
// association process is complete, the file shares on the Amazon FSx file system
// are available for access through the gateway. This operation only supports the
// FSx File Gateway type.
func (c *Client) AssociateFileSystem(ctx context.Context, params *AssociateFileSystemInput, optFns ...func(*Options)) (*AssociateFileSystemOutput, error) {
	if params == nil {
		params = &AssociateFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateFileSystem", params, optFns, c.addOperationAssociateFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateFileSystemInput struct {

	// A unique string value that you supply that is used by the FSx File Gateway to
	// ensure idempotent file system association creation.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with
	// the FSx File Gateway.
	//
	// This member is required.
	LocationARN *string

	// The password of the user credential.
	//
	// This member is required.
	Password *string

	// The user name of the user credential that has permission to access the root
	// share D$ of the Amazon FSx file system. The user account must belong to the
	// Amazon FSx delegated admin user group.
	//
	// This member is required.
	UserName *string

	// The Amazon Resource Name (ARN) of the storage used for the audit logs.
	AuditDestinationARN *string

	// The refresh cache information for the file share or FSx file systems.
	CacheAttributes *types.CacheAttributes

	// Specifies the network configuration information for the gateway associated with
	// the Amazon FSx file system.
	//
	// If multiple file systems are associated with this gateway, this parameter's
	// IpAddresses field is required.
	EndpointNetworkConfiguration *types.EndpointNetworkConfiguration

	// A list of up to 50 tags that can be assigned to the file system association.
	// Each tag is a key-value pair.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type AssociateFileSystemOutput struct {

	// The ARN of the newly created file system association.
	FileSystemAssociationARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateFileSystem"); err != nil {
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
	if err = addOpAssociateFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateFileSystem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateFileSystem",
	}
}