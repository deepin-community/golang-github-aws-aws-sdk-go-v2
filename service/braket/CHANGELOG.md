# v1.29.3 (2024-07-10.2)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.29.2 (2024-07-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.29.1 (2024-06-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.29.0 (2024-06-26)

* **Feature**: Support list-of-string endpoint parameter.

# v1.28.1 (2024-06-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.0 (2024-06-18)

* **Feature**: Track usage of various AWS SDK features in user-agent string.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.11 (2024-06-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.10 (2024-06-07)

* **Bug Fix**: Add clock skew correction on all service clients
* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.9 (2024-06-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.8 (2024-05-23)

* No change notes available for this release.

# v1.27.7 (2024-05-16)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.6 (2024-05-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.5 (2024-05-08)

* **Bug Fix**: GoDoc improvement

# v1.27.4 (2024-03-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.3 (2024-03-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.2 (2024-03-07)

* **Bug Fix**: Remove dependency on go-cmp.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.1 (2024-02-23)

* **Bug Fix**: Move all common, SDK-side middleware stack ops into the service client module to prevent cross-module compatibility issues in the future.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.0 (2024-02-22)

* **Feature**: Add middleware stack snapshot tests.

# v1.26.2 (2024-02-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.26.1 (2024-02-20)

* **Bug Fix**: When sourcing values for a service's `EndpointParameters`, the lack of a configured region (i.e. `options.Region == ""`) will now translate to a `nil` value for `EndpointParameters.Region` instead of a pointer to the empty string `""`. This will result in a much more explicit error when calling an operation instead of an obscure hostname lookup failure.

# v1.26.0 (2024-02-13)

* **Feature**: Bump minimum Go version to 1.20 per our language support policy.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.0 (2024-02-09)

* **Feature**: Creating a job will result in DeviceOfflineException when using an offline device, and DeviceRetiredException when using a retired device.

# v1.24.4 (2024-01-04)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.3 (2023-12-08)

* **Bug Fix**: Reinstate presence of default Retryer in functional options, but still respect max attempts set therein.

# v1.24.2 (2023-12-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.1 (2023-12-06)

* **Bug Fix**: Restore pre-refactor auth behavior where all operations could technically be performed anonymously.

# v1.24.0 (2023-12-04)

* **Feature**: This release enhances service support to create quantum tasks and hybrid jobs associated with Braket Direct Reservations.

# v1.23.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.23.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.23.0 (2023-11-29)

* **Feature**: Expose Options() accessor on service clients.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.5 (2023-11-28.2)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.4 (2023-11-28)

* **Bug Fix**: Respect setting RetryMaxAttempts in functional options at client construction.

# v1.22.3 (2023-11-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.2 (2023-11-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.1 (2023-11-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.0 (2023-11-01)

* **Feature**: Adds support for configured endpoints via environment variables and the AWS shared configuration file.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.0 (2023-10-31)

* **Feature**: **BREAKING CHANGE**: Bump minimum go version to 1.19 per the revised [go version support policy](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/).
* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.2 (2023-10-12)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.1 (2023-10-06)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.0 (2023-09-22)

* **Feature**: This release adds support to view the device queue depth (the number of queued quantum tasks and hybrid jobs on a device) and queue position for a quantum task and hybrid job.

# v1.19.0 (2023-09-18)

* **Announcement**: [BREAKFIX] Change in MaxResults datatype from value to pointer type in cognito-sync service.
* **Feature**: Adds several endpoint ruleset changes across all models: smaller rulesets, removed non-unique regional endpoints, fixes FIPS and DualStack endpoints, and make region not required in SDK::Endpoint. Additional breakfix to cognito-sync field.

# v1.18.5 (2023-08-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.4 (2023-08-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.3 (2023-08-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.2 (2023-08-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.1 (2023-08-01)

* No change notes available for this release.

# v1.18.0 (2023-07-31)

* **Feature**: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.14 (2023-07-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.13 (2023-07-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.12 (2023-06-15)

* No change notes available for this release.

# v1.17.11 (2023-06-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.10 (2023-05-04)

* No change notes available for this release.

# v1.17.9 (2023-04-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.8 (2023-04-10)

* No change notes available for this release.

# v1.17.7 (2023-04-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.6 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.5 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.4 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.17.3 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.2 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.17.1 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.16.17 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.16 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.15 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.14 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.13 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.12 (2022-09-14)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.11 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.10 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.9 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.8 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.7 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.6 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.5 (2022-08-01)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.4 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.3 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.2 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.1 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.0 (2022-04-28)

* **Feature**: This release enables Braket Hybrid Jobs with Embedded Simulators to have multiple instances.

# v1.15.4 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Feature**: Updated service client model to latest release.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.13.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.12.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.

# v1.10.2 (2021-12-03)

* **Bug Fix**: Fixed an issue that prevent auto-filling of an API's idempotency parameters when not explictly provided by the caller.

# v1.10.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/aws/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.0 (2021-11-30)

* **Feature**: API client updated

# v1.9.1 (2021-11-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2021-10-21)

* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.2 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2021-09-10)

* **Feature**: API client updated

# v1.6.0 (2021-08-27)

* **Feature**: Updated API model to latest revision.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.3 (2021-08-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.2 (2021-08-04)

* **Bug Fix**: Fixed an issue that caused one or more API operations to fail when attempting to resolve the service endpoint. ([#1349](https://github.com/aws/aws-sdk-go-v2/pull/1349))
* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2021-07-15)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-06-25)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2021-06-04)

* **Feature**: Updated service client to latest API model.

# v1.3.1 (2021-05-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2021-05-14)

* **Feature**: Constant has been added to modules to enable runtime version inspection for reporting.
* **Dependency Update**: Updated to the latest SDK module versions
