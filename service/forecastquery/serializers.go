// Code generated by smithy-go-codegen DO NOT EDIT.

package forecastquery

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson11_serializeOpQueryForecast struct {
}

func (*awsAwsjson11_serializeOpQueryForecast) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpQueryForecast) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*QueryForecastInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AmazonForecastRuntime.QueryForecast")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentQueryForecastInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpQueryWhatIfForecast struct {
}

func (*awsAwsjson11_serializeOpQueryWhatIfForecast) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpQueryWhatIfForecast) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*QueryWhatIfForecastInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AmazonForecastRuntime.QueryWhatIfForecast")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentQueryWhatIfForecastInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentFilters(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentQueryForecastInput(v *QueryForecastInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EndDate != nil {
		ok := object.Key("EndDate")
		ok.String(*v.EndDate)
	}

	if v.Filters != nil {
		ok := object.Key("Filters")
		if err := awsAwsjson11_serializeDocumentFilters(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.ForecastArn != nil {
		ok := object.Key("ForecastArn")
		ok.String(*v.ForecastArn)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.StartDate != nil {
		ok := object.Key("StartDate")
		ok.String(*v.StartDate)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentQueryWhatIfForecastInput(v *QueryWhatIfForecastInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EndDate != nil {
		ok := object.Key("EndDate")
		ok.String(*v.EndDate)
	}

	if v.Filters != nil {
		ok := object.Key("Filters")
		if err := awsAwsjson11_serializeDocumentFilters(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.StartDate != nil {
		ok := object.Key("StartDate")
		ok.String(*v.StartDate)
	}

	if v.WhatIfForecastArn != nil {
		ok := object.Key("WhatIfForecastArn")
		ok.String(*v.WhatIfForecastArn)
	}

	return nil
}