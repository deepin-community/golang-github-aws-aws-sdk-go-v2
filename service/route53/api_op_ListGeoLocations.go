// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of supported geographic locations.
//
// Countries are listed first, and continents are listed last. If Amazon Route 53
// supports subdivisions for a country (for example, states or provinces), the
// subdivisions for that country are listed in alphabetical order immediately after
// the corresponding country.
//
// Route 53 does not perform authorization for this API because it retrieves
// information that is already available to the public.
//
// For a list of supported geolocation codes, see the [GeoLocation] data type.
//
// [GeoLocation]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GeoLocation.html
func (c *Client) ListGeoLocations(ctx context.Context, params *ListGeoLocationsInput, optFns ...func(*Options)) (*ListGeoLocationsOutput, error) {
	if params == nil {
		params = &ListGeoLocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGeoLocations", params, optFns, c.addOperationListGeoLocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGeoLocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get a list of geographic locations that Amazon Route 53 supports
// for geolocation resource record sets.
type ListGeoLocationsInput struct {

	// (Optional) The maximum number of geolocations to be included in the response
	// body for this request. If more than maxitems geolocations remain to be listed,
	// then the value of the IsTruncated element in the response is true .
	MaxItems *int32

	// The code for the continent with which you want to start listing locations that
	// Amazon Route 53 supports for geolocation. If Route 53 has already returned a
	// page or more of results, if IsTruncated is true, and if NextContinentCode from
	// the previous response has a value, enter that value in startcontinentcode to
	// return the next page of results.
	//
	// Include startcontinentcode only if you want to list continents. Don't include
	// startcontinentcode when you're listing countries or countries with their
	// subdivisions.
	StartContinentCode *string

	// The code for the country with which you want to start listing locations that
	// Amazon Route 53 supports for geolocation. If Route 53 has already returned a
	// page or more of results, if IsTruncated is true , and if NextCountryCode from
	// the previous response has a value, enter that value in startcountrycode to
	// return the next page of results.
	StartCountryCode *string

	// The code for the state of the United States with which you want to start
	// listing locations that Amazon Route 53 supports for geolocation. If Route 53 has
	// already returned a page or more of results, if IsTruncated is true , and if
	// NextSubdivisionCode from the previous response has a value, enter that value in
	// startsubdivisioncode to return the next page of results.
	//
	// To list subdivisions (U.S. states), you must include both startcountrycode and
	// startsubdivisioncode .
	StartSubdivisionCode *string

	noSmithyDocumentSerde
}

// A complex type containing the response information for the request.
type ListGeoLocationsOutput struct {

	// A complex type that contains one GeoLocationDetails element for each location
	// that Amazon Route 53 supports for geolocation.
	//
	// This member is required.
	GeoLocationDetailsList []types.GeoLocationDetails

	// A value that indicates whether more locations remain to be listed after the
	// last location in this response. If so, the value of IsTruncated is true . To get
	// more values, submit another request and include the values of NextContinentCode
	// , NextCountryCode , and NextSubdivisionCode in the startcontinentcode ,
	// startcountrycode , and startsubdivisioncode , as applicable.
	//
	// This member is required.
	IsTruncated bool

	// The value that you specified for MaxItems in the request.
	//
	// This member is required.
	MaxItems *int32

	// If IsTruncated is true , you can make a follow-up request to display more
	// locations. Enter the value of NextContinentCode in the startcontinentcode
	// parameter in another ListGeoLocations request.
	NextContinentCode *string

	// If IsTruncated is true , you can make a follow-up request to display more
	// locations. Enter the value of NextCountryCode in the startcountrycode parameter
	// in another ListGeoLocations request.
	NextCountryCode *string

	// If IsTruncated is true , you can make a follow-up request to display more
	// locations. Enter the value of NextSubdivisionCode in the startsubdivisioncode
	// parameter in another ListGeoLocations request.
	NextSubdivisionCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGeoLocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListGeoLocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListGeoLocations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGeoLocations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGeoLocations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListGeoLocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGeoLocations",
	}
}