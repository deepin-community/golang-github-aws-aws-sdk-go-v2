// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation sets and then enacts a data retrieval policy in the region
// specified in the PUT request. You can set one policy per region for an AWS
// account. The policy is enacted within a few minutes of a successful PUT
// operation.
//
// The set policy operation does not affect retrieval jobs that were in progress
// before the policy was enacted. For more information about data retrieval
// policies, see [Amazon Glacier Data Retrieval Policies].
//
// [Amazon Glacier Data Retrieval Policies]: https://docs.aws.amazon.com/amazonglacier/latest/dev/data-retrieval-policy.html
func (c *Client) SetDataRetrievalPolicy(ctx context.Context, params *SetDataRetrievalPolicyInput, optFns ...func(*Options)) (*SetDataRetrievalPolicyOutput, error) {
	if params == nil {
		params = &SetDataRetrievalPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetDataRetrievalPolicy", params, optFns, c.addOperationSetDataRetrievalPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetDataRetrievalPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// SetDataRetrievalPolicy input.
type SetDataRetrievalPolicyInput struct {

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You can
	// either specify an AWS account ID or optionally a single ' - ' (hyphen), in which
	// case Amazon Glacier uses the AWS account ID associated with the credentials used
	// to sign the request. If you specify your account ID, do not include any hyphens
	// ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The data retrieval policy in JSON format.
	Policy *types.DataRetrievalPolicy

	noSmithyDocumentSerde
}

type SetDataRetrievalPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetDataRetrievalPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSetDataRetrievalPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSetDataRetrievalPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetDataRetrievalPolicy"); err != nil {
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
	if err = addOpSetDataRetrievalPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetDataRetrievalPolicy(options.Region), middleware.Before); err != nil {
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
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
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

func newServiceMetadataMiddleware_opSetDataRetrievalPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetDataRetrievalPolicy",
	}
}