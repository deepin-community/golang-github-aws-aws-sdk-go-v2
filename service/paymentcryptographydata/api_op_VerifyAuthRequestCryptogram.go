// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptographydata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Verifies Authorization Request Cryptogram (ARQC) for a EMV chip payment card
// authorization. For more information, see [Verify auth request cryptogram]in the Amazon Web Services Payment
// Cryptography User Guide.
//
// ARQC generation is done outside of Amazon Web Services Payment Cryptography and
// is typically generated on a point of sale terminal for an EMV chip card to
// obtain payment authorization during transaction time. For ARQC verification, you
// must first import the ARQC generated outside of Amazon Web Services Payment
// Cryptography by calling [ImportKey]. This operation uses the imported ARQC and an major
// encryption key (DUKPT) created by calling [CreateKey]to either provide a boolean ARQC
// verification result or provide an APRC (Authorization Response Cryptogram)
// response using Method 1 or Method 2. The ARPC_METHOD_1 uses AuthResponseCode to
// generate ARPC and ARPC_METHOD_2 uses CardStatusUpdate to generate ARPC.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # VerifyCardValidationData
//
// # VerifyPinData
//
// [Verify auth request cryptogram]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/data-operations.verifyauthrequestcryptogram.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
func (c *Client) VerifyAuthRequestCryptogram(ctx context.Context, params *VerifyAuthRequestCryptogramInput, optFns ...func(*Options)) (*VerifyAuthRequestCryptogramOutput, error) {
	if params == nil {
		params = &VerifyAuthRequestCryptogramInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifyAuthRequestCryptogram", params, optFns, c.addOperationVerifyAuthRequestCryptogramMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyAuthRequestCryptogramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type VerifyAuthRequestCryptogramInput struct {

	// The auth request cryptogram imported into Amazon Web Services Payment
	// Cryptography for ARQC verification using a major encryption key and transaction
	// data.
	//
	// This member is required.
	AuthRequestCryptogram *string

	// The keyARN of the major encryption key that Amazon Web Services Payment
	// Cryptography uses for ARQC verification.
	//
	// This member is required.
	KeyIdentifier *string

	// The method to use when deriving the major encryption key for ARQC verification
	// within Amazon Web Services Payment Cryptography. The same key derivation mode
	// was used for ARQC generation outside of Amazon Web Services Payment
	// Cryptography.
	//
	// This member is required.
	MajorKeyDerivationMode types.MajorKeyDerivationMode

	// The attributes and values to use for deriving a session key for ARQC
	// verification within Amazon Web Services Payment Cryptography. The same
	// attributes were used for ARQC generation outside of Amazon Web Services Payment
	// Cryptography.
	//
	// This member is required.
	SessionKeyDerivationAttributes types.SessionKeyDerivation

	// The transaction data that Amazon Web Services Payment Cryptography uses for
	// ARQC verification. The same transaction is used for ARQC generation outside of
	// Amazon Web Services Payment Cryptography.
	//
	// This member is required.
	TransactionData *string

	// The attributes and values for auth request cryptogram verification. These
	// parameters are required in case using ARPC Method 1 or Method 2 for ARQC
	// verification.
	AuthResponseAttributes types.CryptogramAuthResponse

	noSmithyDocumentSerde
}

type VerifyAuthRequestCryptogramOutput struct {

	// The keyARN of the major encryption key that Amazon Web Services Payment
	// Cryptography uses for ARQC verification.
	//
	// This member is required.
	KeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed.
	//
	// Amazon Web Services Payment Cryptography computes the KCV according to the CMAC
	// specification.
	//
	// This member is required.
	KeyCheckValue *string

	// The result for ARQC verification or ARPC generation within Amazon Web Services
	// Payment Cryptography.
	AuthResponseValue *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifyAuthRequestCryptogramMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpVerifyAuthRequestCryptogram{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpVerifyAuthRequestCryptogram{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "VerifyAuthRequestCryptogram"); err != nil {
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
	if err = addOpVerifyAuthRequestCryptogramValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyAuthRequestCryptogram(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifyAuthRequestCryptogram(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "VerifyAuthRequestCryptogram",
	}
}