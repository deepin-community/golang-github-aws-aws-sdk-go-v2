// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the specified application.
func (c *Client) GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) {
	if params == nil {
		params = &GetApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApplication", params, optFns, c.addOperationGetApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApplicationInput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string

	// The semantic version of the application to get.
	SemanticVersion *string

	noSmithyDocumentSerde
}

type GetApplicationOutput struct {

	// The application Amazon Resource Name (ARN).
	ApplicationId *string

	// The name of the author publishing the app.
	//
	// Minimum length=1. Maximum length=127.
	//
	// Pattern "^[a-z0-9](([a-z0-9]|-(?!-))*[a-z0-9])?$";
	Author *string

	// The date and time this resource was created.
	CreationTime *string

	// The description of the application.
	//
	// Minimum length=1. Maximum length=256
	Description *string

	// A URL with more information about the application, for example the location of
	// your GitHub repository for the application.
	HomePageUrl *string

	// Whether the author of this application has been verified. This means means that
	// AWS has made a good faith review, as a reasonable and prudent service provider,
	// of the information provided by the requester and has confirmed that the
	// requester's identity is as claimed.
	IsVerifiedAuthor *bool

	// Labels to improve discovery of apps in search results.
	//
	// Minimum length=1. Maximum length=127. Maximum number of labels: 10
	//
	// Pattern: "^[a-zA-Z0-9+\\-_:\\/@]+$";
	Labels []string

	// A link to a license file of the app that matches the spdxLicenseID value of
	// your application.
	//
	// Maximum size 5 MB
	LicenseUrl *string

	// The name of the application.
	//
	// Minimum length=1. Maximum length=140
	//
	// Pattern: "[a-zA-Z0-9\\-]+";
	Name *string

	// A link to the readme file in Markdown language that contains a more detailed
	// description of the application and how it works.
	//
	// Maximum size 5 MB
	ReadmeUrl *string

	// A valid identifier from https://spdx.org/licenses/.
	SpdxLicenseId *string

	// The URL to the public profile of a verified author. This URL is submitted by
	// the author.
	VerifiedAuthorUrl *string

	// Version information about the application.
	Version *types.Version

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetApplication"); err != nil {
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
	if err = addOpGetApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetApplication",
	}
}