// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associate the registration with an origination identity such as a phone number
// or sender ID.
func (c *Client) CreateRegistrationAssociation(ctx context.Context, params *CreateRegistrationAssociationInput, optFns ...func(*Options)) (*CreateRegistrationAssociationOutput, error) {
	if params == nil {
		params = &CreateRegistrationAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRegistrationAssociation", params, optFns, c.addOperationCreateRegistrationAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRegistrationAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegistrationAssociationInput struct {

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	// The unique identifier for the origination identity. For example this could be a
	// PhoneNumberId or SenderId.
	//
	// This member is required.
	ResourceId *string

	noSmithyDocumentSerde
}

type CreateRegistrationAssociationOutput struct {

	// The Amazon Resource Name (ARN) for the registration.
	//
	// This member is required.
	RegistrationArn *string

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	// The type of registration form. The list of RegistrationTypes can be found using
	// the DescribeRegistrationTypeDefinitionsaction.
	//
	// This member is required.
	RegistrationType *string

	// The Amazon Resource Name (ARN) of the origination identity that is associated
	// with the registration.
	//
	// This member is required.
	ResourceArn *string

	// The unique identifier for the origination identity. For example this could be a
	// PhoneNumberId or SenderId.
	//
	// This member is required.
	ResourceId *string

	// The registration type or origination identity type.
	//
	// This member is required.
	ResourceType *string

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	IsoCountryCode *string

	// The phone number associated with the registration in E.164 format.
	PhoneNumber *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRegistrationAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateRegistrationAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateRegistrationAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRegistrationAssociation"); err != nil {
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
	if err = addOpCreateRegistrationAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegistrationAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRegistrationAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRegistrationAssociation",
	}
}