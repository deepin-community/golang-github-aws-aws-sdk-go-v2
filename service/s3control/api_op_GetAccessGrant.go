// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
	"time"
)

// Get the details of an access grant from your S3 Access Grants instance.
//
// Permissions You must have the s3:GetAccessGrant permission to use this
// operation.
func (c *Client) GetAccessGrant(ctx context.Context, params *GetAccessGrantInput, optFns ...func(*Options)) (*GetAccessGrantOutput, error) {
	if params == nil {
		params = &GetAccessGrantInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessGrant", params, optFns, c.addOperationGetAccessGrantMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessGrantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessGrantInput struct {

	// The ID of the access grant. S3 Access Grants auto-generates this ID when you
	// create the access grant.
	//
	// This member is required.
	AccessGrantId *string

	// The ID of the Amazon Web Services account that is making this request.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

func (in *GetAccessGrantInput) bindEndpointParams(p *EndpointParameters) {

	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type GetAccessGrantOutput struct {

	// The Amazon Resource Name (ARN) of the access grant.
	AccessGrantArn *string

	// The ID of the access grant. S3 Access Grants auto-generates this ID when you
	// create the access grant.
	AccessGrantId *string

	// The configuration options of the grant location. The grant location is the S3
	// path to the data to which you are granting access.
	AccessGrantsLocationConfiguration *types.AccessGrantsLocationConfiguration

	// The ID of the registered location to which you are granting access. S3 Access
	// Grants assigns this ID when you register the location. S3 Access Grants assigns
	// the ID default to the default location s3:// and assigns an auto-generated ID
	// to other locations that you register.
	AccessGrantsLocationId *string

	// The Amazon Resource Name (ARN) of an Amazon Web Services IAM Identity Center
	// application associated with your Identity Center instance. If the grant includes
	// an application ARN, the grantee can only access the S3 data through this
	// application.
	ApplicationArn *string

	// The date and time when you created the access grant.
	CreatedAt *time.Time

	// The S3 path of the data to which you are granting access. It is the result of
	// appending the Subprefix to the location scope.
	GrantScope *string

	// The user, group, or role to which you are granting access. You can grant access
	// to an IAM user or role. If you have added a corporate directory to Amazon Web
	// Services IAM Identity Center and associated this Identity Center instance with
	// the S3 Access Grants instance, the grantee can also be a corporate directory
	// user or group.
	Grantee *types.Grantee

	// The type of permission that was granted in the access grant. Can be one of the
	// following values:
	//
	//   - READ – Grant read-only access to the S3 data.
	//
	//   - WRITE – Grant write-only access to the S3 data.
	//
	//   - READWRITE – Grant both read and write access to the S3 data.
	Permission types.Permission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessGrantMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessGrant{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessGrant{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccessGrant"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetAccessGrantMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAccessGrantValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessGrant(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetAccessGrantUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGetAccessGrantMiddleware struct {
}

func (*endpointPrefix_opGetAccessGrantMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAccessGrantMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*GetAccessGrantInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetAccessGrantMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetAccessGrantMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessGrant(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccessGrant",
	}
}

func copyGetAccessGrantInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetAccessGrantInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetAccessGrantInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *GetAccessGrantInput) copy() interface{} {
	v := *in
	return &v
}
func backFillGetAccessGrantAccountID(input interface{}, v string) error {
	in := input.(*GetAccessGrantInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetAccessGrantUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetAccessGrantInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}