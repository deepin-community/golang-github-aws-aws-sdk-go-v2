// Code generated by smithy-go-codegen DO NOT EDIT.

package inspectorscan

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/inspectorscan/document"
	internaldocument "github.com/aws/aws-sdk-go-v2/service/inspectorscan/internal/document"
	"github.com/aws/aws-sdk-go-v2/service/inspectorscan/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strconv"
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

type awsRestjson1_deserializeOpScanSbom struct {
}

func (*awsRestjson1_deserializeOpScanSbom) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpScanSbom) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsRestjson1_deserializeOpErrorScanSbom(response, &metadata)
	}
	output := &ScanSbomOutput{}
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

	err = awsRestjson1_deserializeOpDocumentScanSbomOutput(&output, shape)
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

func awsRestjson1_deserializeOpErrorScanSbom(response *smithyhttp.Response, metadata *middleware.Metadata) error {
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

	case strings.EqualFold("InternalServerException", errorCode):
		return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsRestjson1_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentScanSbomOutput(v **ScanSbomOutput, value interface{}) error {
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

	var sv *ScanSbomOutput
	if *v == nil {
		sv = &ScanSbomOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "sbom":
			if err := awsRestjson1_deserializeDocumentSbom(&sv.Sbom, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeOpHttpBindingsInternalServerException(v *types.InternalServerException, response *smithyhttp.Response) error {
	if v == nil {
		return fmt.Errorf("unsupported deserialization for nil %T", v)
	}

	if headerValues := response.Header.Values("Retry-After"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		vv, err := strconv.ParseInt(headerValues[0], 0, 32)
		if err != nil {
			return err
		}
		v.RetryAfterSeconds = ptr.Int32(int32(vv))
	}

	return nil
}
func awsRestjson1_deserializeOpHttpBindingsThrottlingException(v *types.ThrottlingException, response *smithyhttp.Response) error {
	if v == nil {
		return fmt.Errorf("unsupported deserialization for nil %T", v)
	}

	if headerValues := response.Header.Values("Retry-After"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		vv, err := strconv.ParseInt(headerValues[0], 0, 32)
		if err != nil {
			return err
		}
		v.RetryAfterSeconds = ptr.Int32(int32(vv))
	}

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

	if err := awsRestjson1_deserializeOpHttpBindingsInternalServerException(output, response); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response error with invalid HTTP bindings, %w", err)}
	}

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

	if err := awsRestjson1_deserializeOpHttpBindingsThrottlingException(output, response); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response error with invalid HTTP bindings, %w", err)}
	}

	return output
}

func awsRestjson1_deserializeErrorValidationException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ValidationException{}
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

	err := awsRestjson1_deserializeDocumentValidationException(&output, shape)

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

		case "reason":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected InternalServerExceptionReason to be of type string, got %T instead", value)
				}
				sv.Reason = types.InternalServerExceptionReason(jtv)
			}

		case "retryAfterSeconds":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Integer to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.RetryAfterSeconds = ptr.Int32(int32(i64))
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentSbom(v *document.Interface, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	*v = internaldocument.NewDocumentUnmarshaler(value)
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

		case "retryAfterSeconds":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Integer to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.RetryAfterSeconds = ptr.Int32(int32(i64))
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentValidationException(v **types.ValidationException, value interface{}) error {
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

	var sv *types.ValidationException
	if *v == nil {
		sv = &types.ValidationException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "fields":
			if err := awsRestjson1_deserializeDocumentValidationExceptionFields(&sv.Fields, value); err != nil {
				return err
			}

		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		case "reason":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ValidationExceptionReason to be of type string, got %T instead", value)
				}
				sv.Reason = types.ValidationExceptionReason(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentValidationExceptionField(v **types.ValidationExceptionField, value interface{}) error {
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

	var sv *types.ValidationExceptionField
	if *v == nil {
		sv = &types.ValidationExceptionField{}
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

		case "name":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Name = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentValidationExceptionFields(v *[]types.ValidationExceptionField, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.ValidationExceptionField
	if *v == nil {
		cv = []types.ValidationExceptionField{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.ValidationExceptionField
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentValidationExceptionField(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}