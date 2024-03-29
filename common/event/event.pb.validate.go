// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/event/event.proto

package event

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

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Event) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EventMultiError, or nil if none found.
func (m *Event) ValidateAll() error {
	return m.validate(true)
}

func (m *Event) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Subject

	switch v := m.Payload.(type) {
	case *Event_PaymentCompleted:
		if v == nil {
			err := EventValidationError{
				field:  "Payload",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetPaymentCompleted()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "PaymentCompleted",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "PaymentCompleted",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPaymentCompleted()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					field:  "PaymentCompleted",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Event_CancelPayment:
		if v == nil {
			err := EventValidationError{
				field:  "Payload",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetCancelPayment()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "CancelPayment",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "CancelPayment",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCancelPayment()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					field:  "CancelPayment",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return EventMultiError(errors)
	}

	return nil
}

// EventMultiError is an error wrapping multiple validation errors returned by
// Event.ValidateAll() if the designated constraints aren't met.
type EventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventMultiError) AllErrors() []error { return m }

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}

// Validate checks the field values on PayloadPaymentCompleted with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PayloadPaymentCompleted) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayloadPaymentCompleted with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PayloadPaymentCompletedMultiError, or nil if none found.
func (m *PayloadPaymentCompleted) ValidateAll() error {
	return m.validate(true)
}

func (m *PayloadPaymentCompleted) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for Amount

	// no validation rules for OrderId

	// no validation rules for Subject

	if len(errors) > 0 {
		return PayloadPaymentCompletedMultiError(errors)
	}

	return nil
}

// PayloadPaymentCompletedMultiError is an error wrapping multiple validation
// errors returned by PayloadPaymentCompleted.ValidateAll() if the designated
// constraints aren't met.
type PayloadPaymentCompletedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayloadPaymentCompletedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayloadPaymentCompletedMultiError) AllErrors() []error { return m }

// PayloadPaymentCompletedValidationError is the validation error returned by
// PayloadPaymentCompleted.Validate if the designated constraints aren't met.
type PayloadPaymentCompletedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayloadPaymentCompletedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayloadPaymentCompletedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayloadPaymentCompletedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayloadPaymentCompletedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayloadPaymentCompletedValidationError) ErrorName() string {
	return "PayloadPaymentCompletedValidationError"
}

// Error satisfies the builtin error interface
func (e PayloadPaymentCompletedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayloadPaymentCompleted.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayloadPaymentCompletedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayloadPaymentCompletedValidationError{}

// Validate checks the field values on PayloadCancelPayment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PayloadCancelPayment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayloadCancelPayment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PayloadCancelPaymentMultiError, or nil if none found.
func (m *PayloadCancelPayment) ValidateAll() error {
	return m.validate(true)
}

func (m *PayloadCancelPayment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return PayloadCancelPaymentMultiError(errors)
	}

	return nil
}

// PayloadCancelPaymentMultiError is an error wrapping multiple validation
// errors returned by PayloadCancelPayment.ValidateAll() if the designated
// constraints aren't met.
type PayloadCancelPaymentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayloadCancelPaymentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayloadCancelPaymentMultiError) AllErrors() []error { return m }

// PayloadCancelPaymentValidationError is the validation error returned by
// PayloadCancelPayment.Validate if the designated constraints aren't met.
type PayloadCancelPaymentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayloadCancelPaymentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayloadCancelPaymentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayloadCancelPaymentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayloadCancelPaymentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayloadCancelPaymentValidationError) ErrorName() string {
	return "PayloadCancelPaymentValidationError"
}

// Error satisfies the builtin error interface
func (e PayloadCancelPaymentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayloadCancelPayment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayloadCancelPaymentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayloadCancelPaymentValidationError{}
