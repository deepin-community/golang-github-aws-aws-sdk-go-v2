// Code generated by smithy-go-codegen DO NOT EDIT.

package supplychain

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateBillOfMaterialsImportJob struct {
}

func (*awsRestjson1_serializeOpCreateBillOfMaterialsImportJob) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateBillOfMaterialsImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateBillOfMaterialsImportJobInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/api/configuration/instances/{instanceId}/bill-of-materials-import-jobs")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsCreateBillOfMaterialsImportJobInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateBillOfMaterialsImportJobInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCreateBillOfMaterialsImportJobInput(v *CreateBillOfMaterialsImportJobInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.InstanceId == nil || len(*v.InstanceId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member instanceId must not be empty")}
	}
	if v.InstanceId != nil {
		if err := encoder.SetURI("instanceId").String(*v.InstanceId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateBillOfMaterialsImportJobInput(v *CreateBillOfMaterialsImportJobInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("clientToken")
		ok.String(*v.ClientToken)
	}

	if v.S3uri != nil {
		ok := object.Key("s3uri")
		ok.String(*v.S3uri)
	}

	return nil
}

type awsRestjson1_serializeOpGetBillOfMaterialsImportJob struct {
}

func (*awsRestjson1_serializeOpGetBillOfMaterialsImportJob) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetBillOfMaterialsImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetBillOfMaterialsImportJobInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/api/configuration/instances/{instanceId}/bill-of-materials-import-jobs/{jobId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetBillOfMaterialsImportJobInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetBillOfMaterialsImportJobInput(v *GetBillOfMaterialsImportJobInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.InstanceId == nil || len(*v.InstanceId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member instanceId must not be empty")}
	}
	if v.InstanceId != nil {
		if err := encoder.SetURI("instanceId").String(*v.InstanceId); err != nil {
			return err
		}
	}

	if v.JobId == nil || len(*v.JobId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member jobId must not be empty")}
	}
	if v.JobId != nil {
		if err := encoder.SetURI("jobId").String(*v.JobId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpSendDataIntegrationEvent struct {
}

func (*awsRestjson1_serializeOpSendDataIntegrationEvent) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSendDataIntegrationEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SendDataIntegrationEventInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/api-data/data-integration/instance/{instanceId}/data-integration-events")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsSendDataIntegrationEventInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentSendDataIntegrationEventInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsSendDataIntegrationEventInput(v *SendDataIntegrationEventInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.InstanceId == nil || len(*v.InstanceId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member instanceId must not be empty")}
	}
	if v.InstanceId != nil {
		if err := encoder.SetURI("instanceId").String(*v.InstanceId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentSendDataIntegrationEventInput(v *SendDataIntegrationEventInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("clientToken")
		ok.String(*v.ClientToken)
	}

	if v.Data != nil {
		ok := object.Key("data")
		ok.String(*v.Data)
	}

	if v.EventGroupId != nil {
		ok := object.Key("eventGroupId")
		ok.String(*v.EventGroupId)
	}

	if v.EventTimestamp != nil {
		ok := object.Key("eventTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.EventTimestamp))
	}

	if len(v.EventType) > 0 {
		ok := object.Key("eventType")
		ok.String(string(v.EventType))
	}

	return nil
}