// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a team member's attributes in an AWS CodeStar project. For example, you
// can change a team member's role in the project, or change whether they have
// remote access to project resources.
func (c *Client) UpdateTeamMember(ctx context.Context, params *UpdateTeamMemberInput, optFns ...func(*Options)) (*UpdateTeamMemberOutput, error) {
	if params == nil {
		params = &UpdateTeamMemberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTeamMember", params, optFns, c.addOperationUpdateTeamMemberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTeamMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTeamMemberInput struct {

	// The ID of the project.
	//
	// This member is required.
	ProjectId *string

	// The Amazon Resource Name (ARN) of the user for whom you want to change team
	// membership attributes.
	//
	// This member is required.
	UserArn *string

	// The role assigned to the user in the project. Project roles have different
	// levels of access. For more information, see [Working with Teams]in the AWS CodeStar User Guide.
	//
	// [Working with Teams]: http://docs.aws.amazon.com/codestar/latest/userguide/working-with-teams.html
	ProjectRole *string

	// Whether a team member is allowed to remotely access project resources using the
	// SSH public key associated with the user's profile. Even if this is set to True,
	// the user must associate a public key with their profile before the user can
	// access resources.
	RemoteAccessAllowed *bool

	noSmithyDocumentSerde
}

type UpdateTeamMemberOutput struct {

	// The project role granted to the user.
	ProjectRole *string

	// Whether a team member is allowed to remotely access project resources using the
	// SSH public key associated with the user's profile.
	RemoteAccessAllowed *bool

	// The Amazon Resource Name (ARN) of the user whose team membership attributes
	// were updated.
	UserArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTeamMemberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTeamMember{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTeamMember{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTeamMember"); err != nil {
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
	if err = addOpUpdateTeamMemberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTeamMember(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTeamMember(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTeamMember",
	}
}