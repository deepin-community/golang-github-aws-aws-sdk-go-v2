// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// In an environment account, update an environment account connection to use a
// new IAM role.
//
// For more information, see [Environment account connections] in the Proton User guide.
//
// [Environment account connections]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html
func (c *Client) UpdateEnvironmentAccountConnection(ctx context.Context, params *UpdateEnvironmentAccountConnectionInput, optFns ...func(*Options)) (*UpdateEnvironmentAccountConnectionOutput, error) {
	if params == nil {
		params = &UpdateEnvironmentAccountConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnvironmentAccountConnection", params, optFns, c.addOperationUpdateEnvironmentAccountConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnvironmentAccountConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnvironmentAccountConnectionInput struct {

	// The ID of the environment account connection to update.
	//
	// This member is required.
	Id *string

	// The Amazon Resource Name (ARN) of an IAM service role in the environment
	// account. Proton uses this role to provision infrastructure resources using
	// CodeBuild-based provisioning in the associated environment account.
	CodebuildRoleArn *string

	// The Amazon Resource Name (ARN) of the IAM service role that Proton uses when
	// provisioning directly defined components in the associated environment account.
	// It determines the scope of infrastructure that a component can provision in the
	// account.
	//
	// The environment account connection must have a componentRoleArn to allow
	// directly defined components to be associated with any environments running in
	// the account.
	//
	// For more information about components, see [Proton components] in the Proton User Guide.
	//
	// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
	ComponentRoleArn *string

	// The Amazon Resource Name (ARN) of the IAM service role that's associated with
	// the environment account connection to update.
	RoleArn *string

	noSmithyDocumentSerde
}

type UpdateEnvironmentAccountConnectionOutput struct {

	// The environment account connection detail data that's returned by Proton.
	//
	// This member is required.
	EnvironmentAccountConnection *types.EnvironmentAccountConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEnvironmentAccountConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateEnvironmentAccountConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateEnvironmentAccountConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEnvironmentAccountConnection"); err != nil {
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
	if err = addOpUpdateEnvironmentAccountConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnvironmentAccountConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEnvironmentAccountConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEnvironmentAccountConnection",
	}
}