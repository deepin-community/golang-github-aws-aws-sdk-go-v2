// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Activates, deactivates Amazon Inspector deep inspection, or updates custom
// paths for your account.
func (c *Client) UpdateEc2DeepInspectionConfiguration(ctx context.Context, params *UpdateEc2DeepInspectionConfigurationInput, optFns ...func(*Options)) (*UpdateEc2DeepInspectionConfigurationOutput, error) {
	if params == nil {
		params = &UpdateEc2DeepInspectionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEc2DeepInspectionConfiguration", params, optFns, c.addOperationUpdateEc2DeepInspectionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEc2DeepInspectionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEc2DeepInspectionConfigurationInput struct {

	// Specify TRUE to activate Amazon Inspector deep inspection in your account, or
	// FALSE to deactivate. Member accounts in an organization cannot deactivate deep
	// inspection, instead the delegated administrator for the organization can
	// deactivate a member account using [BatchUpdateMemberEc2DeepInspectionStatus].
	//
	// [BatchUpdateMemberEc2DeepInspectionStatus]: https://docs.aws.amazon.com/inspector/v2/APIReference/API_BatchUpdateMemberEc2DeepInspectionStatus.html
	ActivateDeepInspection *bool

	// The Amazon Inspector deep inspection custom paths you are adding for your
	// account.
	PackagePaths []string

	noSmithyDocumentSerde
}

type UpdateEc2DeepInspectionConfigurationOutput struct {

	// An error message explaining why new Amazon Inspector deep inspection custom
	// paths could not be added.
	ErrorMessage *string

	// The current Amazon Inspector deep inspection custom paths for the organization.
	OrgPackagePaths []string

	// The current Amazon Inspector deep inspection custom paths for your account.
	PackagePaths []string

	// The status of Amazon Inspector deep inspection in your account.
	Status types.Ec2DeepInspectionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEc2DeepInspectionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEc2DeepInspectionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEc2DeepInspectionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEc2DeepInspectionConfiguration"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEc2DeepInspectionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEc2DeepInspectionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEc2DeepInspectionConfiguration",
	}
}