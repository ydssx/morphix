// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/product/v1/product.proto

package productv1

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

// Validate checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Product) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProductMultiError, or nil if none found.
func (m *Product) ValidateAll() error {
	return m.validate(true)
}

func (m *Product) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Price

	// no validation rules for Stock

	if len(errors) > 0 {
		return ProductMultiError(errors)
	}

	return nil
}

// ProductMultiError is an error wrapping multiple validation errors returned
// by Product.ValidateAll() if the designated constraints aren't met.
type ProductMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductMultiError) AllErrors() []error { return m }

// ProductValidationError is the validation error returned by Product.Validate
// if the designated constraints aren't met.
type ProductValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductValidationError) ErrorName() string { return "ProductValidationError" }

// Error satisfies the builtin error interface
func (e ProductValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProduct.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductValidationError{}

// Validate checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductRequestMultiError, or nil if none found.
func (m *CreateProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Price

	// no validation rules for Stock

	if len(errors) > 0 {
		return CreateProductRequestMultiError(errors)
	}

	return nil
}

// CreateProductRequestMultiError is an error wrapping multiple validation
// errors returned by CreateProductRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductRequestMultiError) AllErrors() []error { return m }

// CreateProductRequestValidationError is the validation error returned by
// CreateProductRequest.Validate if the designated constraints aren't met.
type CreateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductRequestValidationError) ErrorName() string {
	return "CreateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductRequestValidationError{}

// Validate checks the field values on CreateProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProductResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductResponseMultiError, or nil if none found.
func (m *CreateProductResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateProductResponseMultiError(errors)
	}

	return nil
}

// CreateProductResponseMultiError is an error wrapping multiple validation
// errors returned by CreateProductResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateProductResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductResponseMultiError) AllErrors() []error { return m }

// CreateProductResponseValidationError is the validation error returned by
// CreateProductResponse.Validate if the designated constraints aren't met.
type CreateProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductResponseValidationError) ErrorName() string {
	return "CreateProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductResponseValidationError{}

// Validate checks the field values on GetProductsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProductsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductsRequestMultiError, or nil if none found.
func (m *GetProductsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetProductsRequestMultiError(errors)
	}

	return nil
}

// GetProductsRequestMultiError is an error wrapping multiple validation errors
// returned by GetProductsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProductsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductsRequestMultiError) AllErrors() []error { return m }

// GetProductsRequestValidationError is the validation error returned by
// GetProductsRequest.Validate if the designated constraints aren't met.
type GetProductsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductsRequestValidationError) ErrorName() string {
	return "GetProductsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductsRequestValidationError{}

// Validate checks the field values on GetProductsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProductsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductsResponseMultiError, or nil if none found.
func (m *GetProductsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetProducts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetProductsResponseValidationError{
						field:  fmt.Sprintf("Products[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetProductsResponseValidationError{
						field:  fmt.Sprintf("Products[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetProductsResponseValidationError{
					field:  fmt.Sprintf("Products[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetProductsResponseMultiError(errors)
	}

	return nil
}

// GetProductsResponseMultiError is an error wrapping multiple validation
// errors returned by GetProductsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetProductsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductsResponseMultiError) AllErrors() []error { return m }

// GetProductsResponseValidationError is the validation error returned by
// GetProductsResponse.Validate if the designated constraints aren't met.
type GetProductsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductsResponseValidationError) ErrorName() string {
	return "GetProductsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductsResponseValidationError{}

// Validate checks the field values on GetProductRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductRequestMultiError, or nil if none found.
func (m *GetProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetProductRequestMultiError(errors)
	}

	return nil
}

// GetProductRequestMultiError is an error wrapping multiple validation errors
// returned by GetProductRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductRequestMultiError) AllErrors() []error { return m }

// GetProductRequestValidationError is the validation error returned by
// GetProductRequest.Validate if the designated constraints aren't met.
type GetProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductRequestValidationError) ErrorName() string {
	return "GetProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductRequestValidationError{}

// Validate checks the field values on UpdateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProductRequestMultiError, or nil if none found.
func (m *UpdateProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Price

	// no validation rules for Stock

	if len(errors) > 0 {
		return UpdateProductRequestMultiError(errors)
	}

	return nil
}

// UpdateProductRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateProductRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProductRequestMultiError) AllErrors() []error { return m }

// UpdateProductRequestValidationError is the validation error returned by
// UpdateProductRequest.Validate if the designated constraints aren't met.
type UpdateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProductRequestValidationError) ErrorName() string {
	return "UpdateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProductRequestValidationError{}

// Validate checks the field values on UpdateProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProductResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProductResponseMultiError, or nil if none found.
func (m *UpdateProductResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProductResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateProductResponseMultiError(errors)
	}

	return nil
}

// UpdateProductResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateProductResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateProductResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProductResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProductResponseMultiError) AllErrors() []error { return m }

// UpdateProductResponseValidationError is the validation error returned by
// UpdateProductResponse.Validate if the designated constraints aren't met.
type UpdateProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProductResponseValidationError) ErrorName() string {
	return "UpdateProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProductResponseValidationError{}

// Validate checks the field values on DeleteProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProductRequestMultiError, or nil if none found.
func (m *DeleteProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteProductRequestMultiError(errors)
	}

	return nil
}

// DeleteProductRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteProductRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProductRequestMultiError) AllErrors() []error { return m }

// DeleteProductRequestValidationError is the validation error returned by
// DeleteProductRequest.Validate if the designated constraints aren't met.
type DeleteProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProductRequestValidationError) ErrorName() string {
	return "DeleteProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProductRequestValidationError{}

// Validate checks the field values on DeleteProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProductResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProductResponseMultiError, or nil if none found.
func (m *DeleteProductResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProductResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteProductResponseMultiError(errors)
	}

	return nil
}

// DeleteProductResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteProductResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteProductResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProductResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProductResponseMultiError) AllErrors() []error { return m }

// DeleteProductResponseValidationError is the validation error returned by
// DeleteProductResponse.Validate if the designated constraints aren't met.
type DeleteProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProductResponseValidationError) ErrorName() string {
	return "DeleteProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProductResponseValidationError{}
