# v1.57.3 (2024-07-10.2)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.57.2 (2024-07-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.57.1 (2024-06-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.57.0 (2024-06-26)

* **Feature**: Support list-of-string endpoint parameter.

# v1.56.1 (2024-06-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.56.0 (2024-06-18)

* **Feature**: Track usage of various AWS SDK features in user-agent string.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.55.0 (2024-06-17)

* **Feature**: This release includes support for creating I-frame only video segments for DASH trick play.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.54.0 (2024-06-14)

* **Feature**: This release adds the ability to search for historical job records within the management console using a search box and/or via the SDK/CLI with partial string matching search on input file name.

# v1.53.7 (2024-06-07)

* **Bug Fix**: Add clock skew correction on all service clients
* **Dependency Update**: Updated to the latest SDK module versions

# v1.53.6 (2024-06-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.53.5 (2024-05-23)

* No change notes available for this release.

# v1.53.4 (2024-05-16)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.53.3 (2024-05-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.53.2 (2024-05-08)

* **Bug Fix**: GoDoc improvement

# v1.53.1 (2024-04-29)

* No change notes available for this release.

# v1.53.0 (2024-04-09)

* **Feature**: This release includes support for bringing your own fonts to use for burn-in or DVB-Sub captioning workflows.

# v1.52.5 (2024-03-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.52.4 (2024-03-27)

* No change notes available for this release.

# v1.52.3 (2024-03-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.52.2 (2024-03-07)

* **Bug Fix**: Remove dependency on go-cmp.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.52.1 (2024-02-23)

* **Bug Fix**: Move all common, SDK-side middleware stack ops into the service client module to prevent cross-module compatibility issues in the future.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.52.0 (2024-02-22)

* **Feature**: Add middleware stack snapshot tests.

# v1.51.2 (2024-02-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.51.1 (2024-02-20)

* **Bug Fix**: When sourcing values for a service's `EndpointParameters`, the lack of a configured region (i.e. `options.Region == ""`) will now translate to a `nil` value for `EndpointParameters.Region` instead of a pointer to the empty string `""`. This will result in a much more explicit error when calling an operation instead of an obscure hostname lookup failure.

# v1.51.0 (2024-02-13)

* **Feature**: Bump minimum Go version to 1.20 per our language support policy.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.50.0 (2024-02-01)

* **Feature**: This release includes support for broadcast-mixed audio description tracks.

# v1.49.1 (2024-01-04)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.49.0 (2024-01-03)

* **Feature**: This release includes video engine updates including HEVC improvements, support for ingesting VP9 encoded video in MP4 containers, and support for user-specified 3D LUTs.

# v1.48.5 (2023-12-08)

* **Bug Fix**: Reinstate presence of default Retryer in functional options, but still respect max attempts set therein.

# v1.48.4 (2023-12-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.48.3 (2023-12-06)

* **Bug Fix**: Restore pre-refactor auth behavior where all operations could technically be performed anonymously.

# v1.48.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.48.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.48.0 (2023-11-29)

* **Feature**: Expose Options() accessor on service clients.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.47.4 (2023-11-28.2)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.47.3 (2023-11-28)

* **Bug Fix**: Respect setting RetryMaxAttempts in functional options at client construction.

# v1.47.2 (2023-11-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.47.1 (2023-11-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.47.0 (2023-11-10)

* **Feature**: This release includes the ability to specify any input source as the primary input for corresponding follow modes, and allows users to specify fit and fill behaviors without resizing content.

# v1.46.1 (2023-11-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.46.0 (2023-11-01)

* **Feature**: Adds support for configured endpoints via environment variables and the AWS shared configuration file.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.45.0 (2023-10-31)

* **Feature**: **BREAKING CHANGE**: Bump minimum go version to 1.19 per the revised [go version support policy](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/).
* **Dependency Update**: Updated to the latest SDK module versions

# v1.44.0 (2023-10-24)

* **Feature**: **BREAKFIX**: Correct nullability and default value representation of various input fields across a large number of services. Calling code that references one or more of the affected fields will need to update usage accordingly. See [2162](https://github.com/aws/aws-sdk-go-v2/issues/2162).
* **Feature**: **BREAKFIX**: Correct nullability representation of APIGateway-based services.

# v1.43.2 (2023-10-12)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.43.1 (2023-10-06)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.43.0 (2023-10-03)

* **Feature**: This release adds the ability to replace video frames without modifying the audio essence.

# v1.42.0 (2023-09-22)

* **Feature**: This release supports the creation of of audio-only tracks in CMAF output groups.

# v1.41.0 (2023-08-24)

* **Feature**: This release includes additional audio channel tags in Quicktime outputs, support for film grain synthesis for AV1 outputs, ability to create audio-only FLAC outputs, and ability to specify Amazon S3 destination storage class.

# v1.40.5 (2023-08-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.40.4 (2023-08-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.40.3 (2023-08-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.40.2 (2023-08-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.40.1 (2023-08-01)

* No change notes available for this release.

# v1.40.0 (2023-07-31)

* **Feature**: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.39.2 (2023-07-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.39.1 (2023-07-26)

* **Documentation**: This release includes general updates to user documentation.

# v1.39.0 (2023-07-21)

* **Feature**: This release includes improvements to Preserve 444 handling, compatibility of HEVC sources without frame rates, and general improvements to MP4 outputs.

# v1.38.2 (2023-07-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.38.1 (2023-06-30)

* **Documentation**: This release includes improved color handling of overlays and general updates to user documentation.

# v1.38.0 (2023-06-21)

* **Feature**: This release introduces the bandwidth reduction filter for the HEVC encoder, increases the limits of outputs per job, and updates support for the Nagra SDK to version 1.14.7.

# v1.37.2 (2023-06-15)

* No change notes available for this release.

# v1.37.1 (2023-06-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.37.0 (2023-05-18)

* **Feature**: This release introduces a new MXF Profile for XDCAM which is strictly compliant with the SMPTE RDD 9 standard and improved handling of output name modifiers.

# v1.36.1 (2023-05-04)

* No change notes available for this release.

# v1.36.0 (2023-04-24)

* **Feature**: This release introduces a noise reduction pre-filter, linear interpolation deinterlace mode, video pass-through, updated default job settings, and expanded LC-AAC Stereo audio bitrate ranges.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.35.0 (2023-04-10)

* **Feature**: AWS Elemental MediaConvert SDK now supports conversion of 608 paint-on captions to pop-on captions for SCC sources.

# v1.34.1 (2023-04-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.34.0 (2023-03-23)

* **Feature**: AWS Elemental MediaConvert SDK now supports passthrough of ID3v2 tags for audio inputs to audio-only HLS outputs.

# v1.33.2 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.33.1 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.33.0 (2023-03-03)

* **Feature**: The AWS Elemental MediaConvert SDK has improved handling for different input and output color space combinations.

# v1.32.0 (2023-02-27)

* **Feature**: The AWS Elemental MediaConvert SDK has added support for HDR10 to SDR tone mapping, and animated GIF video input sources.

# v1.31.3 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.31.2 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.31.1 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.31.0 (2023-02-06)

* **Feature**: The AWS Elemental MediaConvert SDK has added improved scene change detection capabilities and a bandwidth reduction filter, along with video quality enhancements, to the AVC encoder.

# v1.30.1 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.30.0 (2023-01-12)

* **Feature**: The AWS Elemental MediaConvert SDK has added support for compact DASH manifest generation, audio normalization using TruePeak measurements, and the ability to clip the sample range in the color corrector.

# v1.29.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.28.2 (2022-12-21)

* No change notes available for this release.

# v1.28.1 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.0 (2022-12-02)

* **Feature**: The AWS Elemental MediaConvert SDK has added support for configurable ID3 eMSG box attributes and the ability to signal them with InbandEventStream tags in DASH and CMAF outputs.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.0 (2022-11-07)

* **Feature**: The AWS Elemental MediaConvert SDK has added support for setting the SDR reference white point for HDR conversions and conversion of HDR10 to DolbyVision without mastering metadata.

# v1.26.2 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.26.1 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.26.0 (2022-10-14)

* **Feature**: MediaConvert now supports specifying the minimum percentage of the HRD buffer available at the end of each encoded video segment.

# v1.25.12 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.11 (2022-09-14)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.10 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.9 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.8 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.7 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.6 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.5 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.4 (2022-08-01)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.3 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.2 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.1 (2022-06-23)

* **Documentation**: AWS Elemental MediaConvert SDK has released support for automatic DolbyVision metadata generation when converting HDR10 to DolbyVision.

# v1.25.0 (2022-06-14)

* **Feature**: AWS Elemental MediaConvert SDK has added support for rules that constrain Automatic-ABR rendition selection when generating ABR package ladders.

# v1.24.1 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.0 (2022-05-24)

* **Feature**: AWS Elemental MediaConvert SDK has added support for rules that constrain Automatic-ABR rendition selection when generating ABR package ladders.

# v1.23.1 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.23.0 (2022-04-29)

* **Feature**: AWS Elemental MediaConvert SDK nows supports creation of Dolby Vision profile 8.1, the ability to generate black frames of video, and introduces audio-only DASH and CMAF support.

# v1.22.1 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.0 (2022-04-08)

* **Feature**: AWS Elemental MediaConvert SDK has added support for the pass-through of WebVTT styling to WebVTT outputs, pass-through of KLV metadata to supported formats, and improved filter support for processing 444/RGB content.

# v1.21.4 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.3 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.2 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.1 (2022-03-08.3)

* No change notes available for this release.

# v1.21.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.0 (2022-01-28)

* **Feature**: Updated to latest API model.

# v1.18.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.0 (2022-01-07)

* **Feature**: API client updated
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.

# v1.15.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/aws/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2021-11-19)

* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.13.0 (2021-10-21)

* **Feature**: API client updated
* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.12.0 (2021-10-11)

* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.0 (2021-09-24)

* **Feature**: API client updated

# v1.10.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.0 (2021-08-27)

* **Feature**: Updated API model to latest revision.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.1 (2021-08-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2021-08-04)

* **Feature**: Updated to latest API model.
* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2021-07-15)

* **Feature**: Updated service model to latest version.
* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2021-07-01)

* **Feature**: API client updated

# v1.6.0 (2021-06-25)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2021-05-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-05-14)

* **Feature**: Constant has been added to modules to enable runtime version inspection for reporting.
* **Feature**: Updated to latest service API model.
* **Dependency Update**: Updated to the latest SDK module versions
