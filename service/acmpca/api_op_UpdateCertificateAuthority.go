// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the status or configuration of a private certificate authority (CA).
// Your private CA must be in the ACTIVE or DISABLED state before you can update
// it. You can disable a private CA that is in the ACTIVE state or make a CA that
// is in the DISABLED state active again.
//
// Both Amazon Web Services Private CA and the IAM principal must have permission
// to write to the S3 bucket that you specify. If the IAM principal making the call
// does not have permission to write to the bucket, then an exception is thrown.
// For more information, see [Access policies for CRLs in Amazon S3].
//
// [Access policies for CRLs in Amazon S3]: https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies
func (c *Client) UpdateCertificateAuthority(ctx context.Context, params *UpdateCertificateAuthorityInput, optFns ...func(*Options)) (*UpdateCertificateAuthorityOutput, error) {
	if params == nil {
		params = &UpdateCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCertificateAuthority", params, optFns, c.addOperationUpdateCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCertificateAuthorityInput struct {

	// Amazon Resource Name (ARN) of the private CA that issued the certificate to be
	// revoked. This must be of the form:
	//
	//     arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	// Contains information to enable Online Certificate Status Protocol (OCSP)
	// support, to enable a certificate revocation list (CRL), to enable both, or to
	// enable neither. If this parameter is not supplied, existing capibilites remain
	// unchanged. For more information, see the [OcspConfiguration]and [CrlConfiguration] types.
	//
	// The following requirements apply to revocation configurations.
	//
	//   - A configuration disabling CRLs or OCSP must contain only the Enabled=False
	//   parameter, and will fail if other parameters such as CustomCname or
	//   ExpirationInDays are included.
	//
	//   - In a CRL configuration, the S3BucketName parameter must conform to [Amazon S3 bucket naming rules].
	//
	//   - A configuration containing a custom Canonical Name (CNAME) parameter for
	//   CRLs or OCSP must conform to [RFC2396]restrictions on the use of special characters in
	//   a CNAME.
	//
	//   - In a CRL or OCSP configuration, the value of a CNAME parameter must not
	//   include a protocol prefix such as "http://" or "https://".
	//
	// [RFC2396]: https://www.ietf.org/rfc/rfc2396.txt
	// [OcspConfiguration]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_OcspConfiguration.html
	// [Amazon S3 bucket naming rules]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html
	// [CrlConfiguration]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CrlConfiguration.html
	RevocationConfiguration *types.RevocationConfiguration

	// Status of your private CA.
	Status types.CertificateAuthorityStatus

	noSmithyDocumentSerde
}

type UpdateCertificateAuthorityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCertificateAuthority"); err != nil {
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
	if err = addOpUpdateCertificateAuthorityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCertificateAuthority(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCertificateAuthority(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCertificateAuthority",
	}
}