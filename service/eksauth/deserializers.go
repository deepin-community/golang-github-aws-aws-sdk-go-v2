// Code generated by smithy-go-codegen DO NOT EDIT.

package eksauth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/eksauth/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strings"
	"time"
)

func deserializeS3Expires(v string) (*time.Time, error) {
	t, err := smithytime.ParseHTTPDate(v)
	if err != nil {
		return nil, nil
	}
	return &t, nil
}

type awsRestjson1_deserializeOpAssumeRoleForPodIdentity struct {
}

func (*awsRestjson1_deserializeOpAssumeRoleForPodIdentity) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpAssumeRoleForPodIdentity) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorAssumeRoleForPodIdentity(response, &metadata)
	}
	output := &AssumeRoleForPodIdentityOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentAssumeRoleForPodIdentityOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorAssumeRoleForPodIdentity(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("AccessDeniedException", errorCode):
		return awsRestjson1_deserializeErrorAccessDeniedException(response, errorBody)

	case strings.EqualFold("ExpiredTokenException", errorCode):
		return awsRestjson1_deserializeErrorExpiredTokenException(response, errorBody)

	case strings.EqualFold("InternalServerException", errorCode):
		return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)

	case strings.EqualFold("InvalidParameterException", errorCode):
		return awsRestjson1_deserializeErrorInvalidParameterException(response, errorBody)

	case strings.EqualFold("InvalidRequestException", errorCode):
		return awsRestjson1_deserializeErrorInvalidRequestException(response, errorBody)

	case strings.EqualFold("InvalidTokenException", errorCode):
		return awsRestjson1_deserializeErrorInvalidTokenException(response, errorBody)

	case strings.EqualFold("ResourceNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorResourceNotFoundException(response, errorBody)

	case strings.EqualFold("ServiceUnavailableException", errorCode):
		return awsRestjson1_deserializeErrorServiceUnavailableException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentAssumeRoleForPodIdentityOutput(v **AssumeRoleForPodIdentityOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *AssumeRoleForPodIdentityOutput
	if *v == nil {
		sv = &AssumeRoleForPodIdentityOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "assumedRoleUser":
			if err := awsRestjson1_deserializeDocumentAssumedRoleUser(&sv.AssumedRoleUser, value); err != nil {
				return err
			}

		case "audience":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Audience = ptr.String(jtv)
			}

		case "credentials":
			if err := awsRestjson1_deserializeDocumentCredentials(&sv.Credentials, value); err != nil {
				return err
			}

		case "podIdentityAssociation":
			if err := awsRestjson1_deserializeDocumentPodIdentityAssociation(&sv.PodIdentityAssociation, value); err != nil {
				return err
			}

		case "subject":
			if err := awsRestjson1_deserializeDocumentSubject(&sv.Subject, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeErrorAccessDeniedException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.AccessDeniedException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentAccessDeniedException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorExpiredTokenException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ExpiredTokenException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentExpiredTokenException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInternalServerException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InternalServerException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInternalServerException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInvalidParameterException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InvalidParameterException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInvalidParameterException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInvalidRequestException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InvalidRequestException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInvalidRequestException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInvalidTokenException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InvalidTokenException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInvalidTokenException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorResourceNotFoundException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ResourceNotFoundException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentResourceNotFoundException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorServiceUnavailableException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ServiceUnavailableException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentServiceUnavailableException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorThrottlingException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ThrottlingException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentThrottlingException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocumentAccessDeniedException(v **types.AccessDeniedException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.AccessDeniedException
	if *v == nil {
		sv = &types.AccessDeniedException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentAssumedRoleUser(v **types.AssumedRoleUser, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.AssumedRoleUser
	if *v == nil {
		sv = &types.AssumedRoleUser{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "arn":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Arn = ptr.String(jtv)
			}

		case "assumeRoleId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.AssumeRoleId = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentCredentials(v **types.Credentials, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Credentials
	if *v == nil {
		sv = &types.Credentials{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "accessKeyId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.AccessKeyId = ptr.String(jtv)
			}

		case "expiration":
			if value != nil {
				switch jtv := value.(type) {
				case json.Number:
					f64, err := jtv.Float64()
					if err != nil {
						return err
					}
					sv.Expiration = ptr.Time(smithytime.ParseEpochSeconds(f64))

				default:
					return fmt.Errorf("expected Timestamp to be a JSON Number, got %T instead", value)

				}
			}

		case "secretAccessKey":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.SecretAccessKey = ptr.String(jtv)
			}

		case "sessionToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.SessionToken = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentExpiredTokenException(v **types.ExpiredTokenException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ExpiredTokenException
	if *v == nil {
		sv = &types.ExpiredTokenException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInternalServerException(v **types.InternalServerException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalServerException
	if *v == nil {
		sv = &types.InternalServerException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInvalidParameterException(v **types.InvalidParameterException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InvalidParameterException
	if *v == nil {
		sv = &types.InvalidParameterException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInvalidRequestException(v **types.InvalidRequestException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InvalidRequestException
	if *v == nil {
		sv = &types.InvalidRequestException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInvalidTokenException(v **types.InvalidTokenException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InvalidTokenException
	if *v == nil {
		sv = &types.InvalidTokenException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentPodIdentityAssociation(v **types.PodIdentityAssociation, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.PodIdentityAssociation
	if *v == nil {
		sv = &types.PodIdentityAssociation{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "associationArn":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.AssociationArn = ptr.String(jtv)
			}

		case "associationId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.AssociationId = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentResourceNotFoundException(v **types.ResourceNotFoundException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ResourceNotFoundException
	if *v == nil {
		sv = &types.ResourceNotFoundException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentServiceUnavailableException(v **types.ServiceUnavailableException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ServiceUnavailableException
	if *v == nil {
		sv = &types.ServiceUnavailableException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentSubject(v **types.Subject, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Subject
	if *v == nil {
		sv = &types.Subject{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "namespace":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Namespace = ptr.String(jtv)
			}

		case "serviceAccount":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.ServiceAccount = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentThrottlingException(v **types.ThrottlingException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ThrottlingException
	if *v == nil {
		sv = &types.ThrottlingException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}