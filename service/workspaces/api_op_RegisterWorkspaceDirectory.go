// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers the specified directory. This operation is asynchronous and returns
// before the WorkSpace directory is registered. If this is the first time you are
// registering a directory, you will need to create the workspaces_DefaultRole role
// before you can register a directory. For more information, see [Creating the workspaces_DefaultRole Role].
//
// [Creating the workspaces_DefaultRole Role]: https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-access-control.html#create-default-role
func (c *Client) RegisterWorkspaceDirectory(ctx context.Context, params *RegisterWorkspaceDirectoryInput, optFns ...func(*Options)) (*RegisterWorkspaceDirectoryOutput, error) {
	if params == nil {
		params = &RegisterWorkspaceDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterWorkspaceDirectory", params, optFns, c.addOperationRegisterWorkspaceDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterWorkspaceDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterWorkspaceDirectoryInput struct {

	// The active directory config of the directory.
	ActiveDirectoryConfig *types.ActiveDirectoryConfig

	// The identifier of the directory. You cannot register a directory if it does not
	// have a status of Active. If the directory does not have a status of Active, you
	// will receive an InvalidResourceStateException error. If you have already
	// registered the maximum number of directories that you can register with Amazon
	// WorkSpaces, you will receive a ResourceLimitExceededException error. Deregister
	// directories that you are not using for WorkSpaces, and try again.
	DirectoryId *string

	// Indicates whether self-service capabilities are enabled or disabled.
	EnableSelfService *bool

	// Indicates whether Amazon WorkDocs is enabled or disabled. If you have enabled
	// this parameter and WorkDocs is not available in the Region, you will receive an
	// OperationNotSupportedException error. Set EnableWorkDocs to disabled, and try
	// again.
	EnableWorkDocs *bool

	// The identifiers of the subnets for your virtual private cloud (VPC). Make sure
	// that the subnets are in supported Availability Zones. The subnets must also be
	// in separate Availability Zones. If these conditions are not met, you will
	// receive an OperationNotSupportedException error.
	SubnetIds []string

	// The tags associated with the directory.
	Tags []types.Tag

	// Indicates whether your WorkSpace directory is dedicated or shared. To use Bring
	// Your Own License (BYOL) images, this value must be set to DEDICATED and your
	// Amazon Web Services account must be enabled for BYOL. If your account has not
	// been enabled for BYOL, you will receive an InvalidParameterValuesException
	// error. For more information about BYOL images, see [Bring Your Own Windows Desktop Images].
	//
	// [Bring Your Own Windows Desktop Images]: https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html
	Tenancy types.Tenancy

	// The type of identity management the user is using.
	UserIdentityType types.UserIdentityType

	// Description of the directory to register.
	WorkspaceDirectoryDescription *string

	// The name of the directory to register.
	WorkspaceDirectoryName *string

	// Indicates whether the directory's WorkSpace type is personal or pools.
	WorkspaceType types.WorkspaceType

	noSmithyDocumentSerde
}

type RegisterWorkspaceDirectoryOutput struct {

	// The identifier of the directory.
	DirectoryId *string

	// The registration status of the WorkSpace directory.
	State types.WorkspaceDirectoryState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterWorkspaceDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterWorkspaceDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterWorkspaceDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterWorkspaceDirectory"); err != nil {
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
	if err = addOpRegisterWorkspaceDirectoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterWorkspaceDirectory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterWorkspaceDirectory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterWorkspaceDirectory",
	}
}