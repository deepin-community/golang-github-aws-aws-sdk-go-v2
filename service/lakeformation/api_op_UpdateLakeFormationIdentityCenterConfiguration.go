// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the IAM Identity Center connection parameters.
func (c *Client) UpdateLakeFormationIdentityCenterConfiguration(ctx context.Context, params *UpdateLakeFormationIdentityCenterConfigurationInput, optFns ...func(*Options)) (*UpdateLakeFormationIdentityCenterConfigurationOutput, error) {
	if params == nil {
		params = &UpdateLakeFormationIdentityCenterConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLakeFormationIdentityCenterConfiguration", params, optFns, c.addOperationUpdateLakeFormationIdentityCenterConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLakeFormationIdentityCenterConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLakeFormationIdentityCenterConfigurationInput struct {

	// Allows to enable or disable the IAM Identity Center connection.
	ApplicationStatus types.ApplicationStatus

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, view definitions, and other control information to manage
	// your Lake Formation environment.
	CatalogId *string

	// A list of the account IDs of Amazon Web Services accounts of third-party
	// applications that are allowed to access data managed by Lake Formation.
	ExternalFiltering *types.ExternalFilteringConfiguration

	// A list of Amazon Web Services account IDs or Amazon Web Services
	// organization/organizational unit ARNs that are allowed to access to access data
	// managed by Lake Formation.
	//
	// If the ShareRecipients list includes valid values, then the resource share is
	// updated with the principals you want to have access to the resources.
	//
	// If the ShareRecipients value is null, both the list of share recipients and the
	// resource share remain unchanged.
	//
	// If the ShareRecipients value is an empty list, then the existing share
	// recipients list will be cleared, and the resource share will be deleted.
	ShareRecipients []types.DataLakePrincipal

	noSmithyDocumentSerde
}

type UpdateLakeFormationIdentityCenterConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLakeFormationIdentityCenterConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLakeFormationIdentityCenterConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLakeFormationIdentityCenterConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLakeFormationIdentityCenterConfiguration"); err != nil {
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
	if err = addOpUpdateLakeFormationIdentityCenterConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLakeFormationIdentityCenterConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLakeFormationIdentityCenterConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLakeFormationIdentityCenterConfiguration",
	}
}