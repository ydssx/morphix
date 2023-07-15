// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/sms/v1/sms.proto

package smsv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on SendSMSRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendSMSRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendSMSRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SendSMSRequestMultiError,
// or nil if none found.
func (m *SendSMSRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SendSMSRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MobileNumber

	// no validation rules for Message

	// no validation rules for SenderId

	// no validation rules for TemplateId

	// no validation rules for TemplateParameters

	// no validation rules for Scene

	if len(errors) > 0 {
		return SendSMSRequestMultiError(errors)
	}

	return nil
}

// SendSMSRequestMultiError is an error wrapping multiple validation errors
// returned by SendSMSRequest.ValidateAll() if the designated constraints
// aren't met.
type SendSMSRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendSMSRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendSMSRequestMultiError) AllErrors() []error { return m }

// SendSMSRequestValidationError is the validation error returned by
// SendSMSRequest.Validate if the designated constraints aren't met.
type SendSMSRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendSMSRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendSMSRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendSMSRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendSMSRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendSMSRequestValidationError) ErrorName() string { return "SendSMSRequestValidationError" }

// Error satisfies the builtin error interface
func (e SendSMSRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendSMSRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendSMSRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendSMSRequestValidationError{}

// Validate checks the field values on SendSMSResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SendSMSResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendSMSResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendSMSResponseMultiError, or nil if none found.
func (m *SendSMSResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SendSMSResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for ErrorMessage

	if len(errors) > 0 {
		return SendSMSResponseMultiError(errors)
	}

	return nil
}

// SendSMSResponseMultiError is an error wrapping multiple validation errors
// returned by SendSMSResponse.ValidateAll() if the designated constraints
// aren't met.
type SendSMSResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendSMSResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendSMSResponseMultiError) AllErrors() []error { return m }

// SendSMSResponseValidationError is the validation error returned by
// SendSMSResponse.Validate if the designated constraints aren't met.
type SendSMSResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendSMSResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendSMSResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendSMSResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendSMSResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendSMSResponseValidationError) ErrorName() string { return "SendSMSResponseValidationError" }

// Error satisfies the builtin error interface
func (e SendSMSResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendSMSResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendSMSResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendSMSResponseValidationError{}

// Validate checks the field values on QuerySMSStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuerySMSStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuerySMSStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuerySMSStatusRequestMultiError, or nil if none found.
func (m *QuerySMSStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuerySMSStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SmsId

	// no validation rules for MobileNumber

	// no validation rules for StartTime

	// no validation rules for EndTime

	// no validation rules for SmsCode

	// no validation rules for Scene

	if len(errors) > 0 {
		return QuerySMSStatusRequestMultiError(errors)
	}

	return nil
}

// QuerySMSStatusRequestMultiError is an error wrapping multiple validation
// errors returned by QuerySMSStatusRequest.ValidateAll() if the designated
// constraints aren't met.
type QuerySMSStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuerySMSStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuerySMSStatusRequestMultiError) AllErrors() []error { return m }

// QuerySMSStatusRequestValidationError is the validation error returned by
// QuerySMSStatusRequest.Validate if the designated constraints aren't met.
type QuerySMSStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuerySMSStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuerySMSStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuerySMSStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuerySMSStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuerySMSStatusRequestValidationError) ErrorName() string {
	return "QuerySMSStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QuerySMSStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuerySMSStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuerySMSStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuerySMSStatusRequestValidationError{}

// Validate checks the field values on QuerySMSStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuerySMSStatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuerySMSStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuerySMSStatusResponseMultiError, or nil if none found.
func (m *QuerySMSStatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QuerySMSStatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return QuerySMSStatusResponseMultiError(errors)
	}

	return nil
}

// QuerySMSStatusResponseMultiError is an error wrapping multiple validation
// errors returned by QuerySMSStatusResponse.ValidateAll() if the designated
// constraints aren't met.
type QuerySMSStatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuerySMSStatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuerySMSStatusResponseMultiError) AllErrors() []error { return m }

// QuerySMSStatusResponseValidationError is the validation error returned by
// QuerySMSStatusResponse.Validate if the designated constraints aren't met.
type QuerySMSStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuerySMSStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuerySMSStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuerySMSStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuerySMSStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuerySMSStatusResponseValidationError) ErrorName() string {
	return "QuerySMSStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e QuerySMSStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuerySMSStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuerySMSStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuerySMSStatusResponseValidationError{}

// Validate checks the field values on SMSStatus with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SMSStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SMSStatus with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SMSStatusMultiError, or nil
// if none found.
func (m *SMSStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *SMSStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SmsId

	// no validation rules for MobileNumber

	// no validation rules for Status

	// no validation rules for Timestamp

	if len(errors) > 0 {
		return SMSStatusMultiError(errors)
	}

	return nil
}

// SMSStatusMultiError is an error wrapping multiple validation errors returned
// by SMSStatus.ValidateAll() if the designated constraints aren't met.
type SMSStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SMSStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SMSStatusMultiError) AllErrors() []error { return m }

// SMSStatusValidationError is the validation error returned by
// SMSStatus.Validate if the designated constraints aren't met.
type SMSStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SMSStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SMSStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SMSStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SMSStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SMSStatusValidationError) ErrorName() string { return "SMSStatusValidationError" }

// Error satisfies the builtin error interface
func (e SMSStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSMSStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SMSStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SMSStatusValidationError{}

// Validate checks the field values on SMSTemplate with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SMSTemplate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SMSTemplate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SMSTemplateMultiError, or
// nil if none found.
func (m *SMSTemplate) ValidateAll() error {
	return m.validate(true)
}

func (m *SMSTemplate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TemplateId

	// no validation rules for TemplateContent

	// no validation rules for TemplateName

	// no validation rules for TemplateStatus

	if len(errors) > 0 {
		return SMSTemplateMultiError(errors)
	}

	return nil
}

// SMSTemplateMultiError is an error wrapping multiple validation errors
// returned by SMSTemplate.ValidateAll() if the designated constraints aren't met.
type SMSTemplateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SMSTemplateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SMSTemplateMultiError) AllErrors() []error { return m }

// SMSTemplateValidationError is the validation error returned by
// SMSTemplate.Validate if the designated constraints aren't met.
type SMSTemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SMSTemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SMSTemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SMSTemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SMSTemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SMSTemplateValidationError) ErrorName() string { return "SMSTemplateValidationError" }

// Error satisfies the builtin error interface
func (e SMSTemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSMSTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SMSTemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SMSTemplateValidationError{}

// Validate checks the field values on SMSSignature with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SMSSignature) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SMSSignature with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SMSSignatureMultiError, or
// nil if none found.
func (m *SMSSignature) ValidateAll() error {
	return m.validate(true)
}

func (m *SMSSignature) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SignatureId

	// no validation rules for SignatureContent

	// no validation rules for SignatureType

	// no validation rules for SignatureStatus

	if len(errors) > 0 {
		return SMSSignatureMultiError(errors)
	}

	return nil
}

// SMSSignatureMultiError is an error wrapping multiple validation errors
// returned by SMSSignature.ValidateAll() if the designated constraints aren't met.
type SMSSignatureMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SMSSignatureMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SMSSignatureMultiError) AllErrors() []error { return m }

// SMSSignatureValidationError is the validation error returned by
// SMSSignature.Validate if the designated constraints aren't met.
type SMSSignatureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SMSSignatureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SMSSignatureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SMSSignatureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SMSSignatureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SMSSignatureValidationError) ErrorName() string { return "SMSSignatureValidationError" }

// Error satisfies the builtin error interface
func (e SMSSignatureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSMSSignature.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SMSSignatureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SMSSignatureValidationError{}

// Validate checks the field values on TemplateManagementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TemplateManagementRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TemplateManagementRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TemplateManagementRequestMultiError, or nil if none found.
func (m *TemplateManagementRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TemplateManagementRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTemplate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TemplateManagementRequestValidationError{
					field:  "Template",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TemplateManagementRequestValidationError{
					field:  "Template",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTemplate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TemplateManagementRequestValidationError{
				field:  "Template",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TemplateId

	if len(errors) > 0 {
		return TemplateManagementRequestMultiError(errors)
	}

	return nil
}

// TemplateManagementRequestMultiError is an error wrapping multiple validation
// errors returned by TemplateManagementRequest.ValidateAll() if the
// designated constraints aren't met.
type TemplateManagementRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TemplateManagementRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TemplateManagementRequestMultiError) AllErrors() []error { return m }

// TemplateManagementRequestValidationError is the validation error returned by
// TemplateManagementRequest.Validate if the designated constraints aren't met.
type TemplateManagementRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TemplateManagementRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TemplateManagementRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TemplateManagementRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TemplateManagementRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TemplateManagementRequestValidationError) ErrorName() string {
	return "TemplateManagementRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TemplateManagementRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTemplateManagementRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TemplateManagementRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TemplateManagementRequestValidationError{}

// Validate checks the field values on TemplateManagementResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TemplateManagementResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TemplateManagementResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TemplateManagementResponseMultiError, or nil if none found.
func (m *TemplateManagementResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TemplateManagementResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for ErrorMessage

	for idx, item := range m.GetTemplates() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TemplateManagementResponseValidationError{
						field:  fmt.Sprintf("Templates[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TemplateManagementResponseValidationError{
						field:  fmt.Sprintf("Templates[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TemplateManagementResponseValidationError{
					field:  fmt.Sprintf("Templates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TemplateManagementResponseMultiError(errors)
	}

	return nil
}

// TemplateManagementResponseMultiError is an error wrapping multiple
// validation errors returned by TemplateManagementResponse.ValidateAll() if
// the designated constraints aren't met.
type TemplateManagementResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TemplateManagementResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TemplateManagementResponseMultiError) AllErrors() []error { return m }

// TemplateManagementResponseValidationError is the validation error returned
// by TemplateManagementResponse.Validate if the designated constraints aren't met.
type TemplateManagementResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TemplateManagementResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TemplateManagementResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TemplateManagementResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TemplateManagementResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TemplateManagementResponseValidationError) ErrorName() string {
	return "TemplateManagementResponseValidationError"
}

// Error satisfies the builtin error interface
func (e TemplateManagementResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTemplateManagementResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TemplateManagementResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TemplateManagementResponseValidationError{}

// Validate checks the field values on SignatureManagementRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SignatureManagementRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignatureManagementRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SignatureManagementRequestMultiError, or nil if none found.
func (m *SignatureManagementRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignatureManagementRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSignature()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignatureManagementRequestValidationError{
					field:  "Signature",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignatureManagementRequestValidationError{
					field:  "Signature",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSignature()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignatureManagementRequestValidationError{
				field:  "Signature",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SignatureId

	if len(errors) > 0 {
		return SignatureManagementRequestMultiError(errors)
	}

	return nil
}

// SignatureManagementRequestMultiError is an error wrapping multiple
// validation errors returned by SignatureManagementRequest.ValidateAll() if
// the designated constraints aren't met.
type SignatureManagementRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignatureManagementRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignatureManagementRequestMultiError) AllErrors() []error { return m }

// SignatureManagementRequestValidationError is the validation error returned
// by SignatureManagementRequest.Validate if the designated constraints aren't met.
type SignatureManagementRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignatureManagementRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignatureManagementRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignatureManagementRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignatureManagementRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignatureManagementRequestValidationError) ErrorName() string {
	return "SignatureManagementRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SignatureManagementRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignatureManagementRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignatureManagementRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignatureManagementRequestValidationError{}

// Validate checks the field values on SignatureManagementResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SignatureManagementResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignatureManagementResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SignatureManagementResponseMultiError, or nil if none found.
func (m *SignatureManagementResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SignatureManagementResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for ErrorMessage

	for idx, item := range m.GetSignatures() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SignatureManagementResponseValidationError{
						field:  fmt.Sprintf("Signatures[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SignatureManagementResponseValidationError{
						field:  fmt.Sprintf("Signatures[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SignatureManagementResponseValidationError{
					field:  fmt.Sprintf("Signatures[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SignatureManagementResponseMultiError(errors)
	}

	return nil
}

// SignatureManagementResponseMultiError is an error wrapping multiple
// validation errors returned by SignatureManagementResponse.ValidateAll() if
// the designated constraints aren't met.
type SignatureManagementResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignatureManagementResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignatureManagementResponseMultiError) AllErrors() []error { return m }

// SignatureManagementResponseValidationError is the validation error returned
// by SignatureManagementResponse.Validate if the designated constraints
// aren't met.
type SignatureManagementResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignatureManagementResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignatureManagementResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignatureManagementResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignatureManagementResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignatureManagementResponseValidationError) ErrorName() string {
	return "SignatureManagementResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SignatureManagementResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignatureManagementResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignatureManagementResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignatureManagementResponseValidationError{}
