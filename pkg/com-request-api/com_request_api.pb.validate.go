// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ozonmp/com_request_api/v1/com_request_api.proto

package com_request_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetService()) < 2 {
		return RequestValidationError{
			field:  "Service",
			reason: "value length must be at least 2 runes",
		}
	}

	if utf8.RuneCountInString(m.GetUser()) < 2 {
		return RequestValidationError{
			field:  "User",
			reason: "value length must be at least 2 runes",
		}
	}

	if utf8.RuneCountInString(m.GetText()) < 2 {
		return RequestValidationError{
			field:  "Text",
			reason: "value length must be at least 2 runes",
		}
	}

	return nil
}

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

// Validate checks the field values on DescribeRequestV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRequestV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRequestId() <= 0 {
		return DescribeRequestV1RequestValidationError{
			field:  "RequestId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeRequestV1RequestValidationError is the validation error returned by
// DescribeRequestV1Request.Validate if the designated constraints aren't met.
type DescribeRequestV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRequestV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRequestV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRequestV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRequestV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRequestV1RequestValidationError) ErrorName() string {
	return "DescribeRequestV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRequestV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRequestV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRequestV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRequestV1RequestValidationError{}

// Validate checks the field values on DescribeRequestV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRequestV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeRequestV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeRequestV1ResponseValidationError is the validation error returned by
// DescribeRequestV1Response.Validate if the designated constraints aren't met.
type DescribeRequestV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRequestV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRequestV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRequestV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRequestV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRequestV1ResponseValidationError) ErrorName() string {
	return "DescribeRequestV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRequestV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRequestV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRequestV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRequestV1ResponseValidationError{}

// Validate checks the field values on CreateRequestV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRequestV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRequestV1RequestValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateRequestV1RequestValidationError is the validation error returned by
// CreateRequestV1Request.Validate if the designated constraints aren't met.
type CreateRequestV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestV1RequestValidationError) ErrorName() string {
	return "CreateRequestV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRequestV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequestV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestV1RequestValidationError{}

// Validate checks the field values on CreateRequestV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRequestV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRequestId() <= 0 {
		return CreateRequestV1ResponseValidationError{
			field:  "RequestId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateRequestV1ResponseValidationError is the validation error returned by
// CreateRequestV1Response.Validate if the designated constraints aren't met.
type CreateRequestV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestV1ResponseValidationError) ErrorName() string {
	return "CreateRequestV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRequestV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequestV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestV1ResponseValidationError{}

// Validate checks the field values on ListRequestV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRequestV1Request) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListRequestV1RequestValidationError is the validation error returned by
// ListRequestV1Request.Validate if the designated constraints aren't met.
type ListRequestV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestV1RequestValidationError) ErrorName() string {
	return "ListRequestV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRequestV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequestV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestV1RequestValidationError{}

// Validate checks the field values on ListRequestV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRequestV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRequest() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRequestV1ResponseValidationError{
					field:  fmt.Sprintf("Request[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListRequestV1ResponseValidationError is the validation error returned by
// ListRequestV1Response.Validate if the designated constraints aren't met.
type ListRequestV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestV1ResponseValidationError) ErrorName() string {
	return "ListRequestV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRequestV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequestV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestV1ResponseValidationError{}

// Validate checks the field values on RemoveRequestV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRequestV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRequestId() <= 0 {
		return RemoveRequestV1RequestValidationError{
			field:  "RequestId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveRequestV1RequestValidationError is the validation error returned by
// RemoveRequestV1Request.Validate if the designated constraints aren't met.
type RemoveRequestV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRequestV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRequestV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRequestV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRequestV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRequestV1RequestValidationError) ErrorName() string {
	return "RemoveRequestV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRequestV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRequestV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRequestV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRequestV1RequestValidationError{}

// Validate checks the field values on RemoveRequestV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRequestV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// RemoveRequestV1ResponseValidationError is the validation error returned by
// RemoveRequestV1Response.Validate if the designated constraints aren't met.
type RemoveRequestV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRequestV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRequestV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRequestV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRequestV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRequestV1ResponseValidationError) ErrorName() string {
	return "RemoveRequestV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRequestV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRequestV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRequestV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRequestV1ResponseValidationError{}