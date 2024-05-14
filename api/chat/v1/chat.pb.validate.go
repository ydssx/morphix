// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/chat/v1/chat.proto

package chat

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

// Validate checks the field values on ClientMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClientMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClientMessageMultiError, or
// nil if none found.
func (m *ClientMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for MessageText

	// no validation rules for ImgUrl

	if len(errors) > 0 {
		return ClientMessageMultiError(errors)
	}

	return nil
}

// ClientMessageMultiError is an error wrapping multiple validation errors
// returned by ClientMessage.ValidateAll() if the designated constraints
// aren't met.
type ClientMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientMessageMultiError) AllErrors() []error { return m }

// ClientMessageValidationError is the validation error returned by
// ClientMessage.Validate if the designated constraints aren't met.
type ClientMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientMessageValidationError) ErrorName() string { return "ClientMessageValidationError" }

// Error satisfies the builtin error interface
func (e ClientMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientMessageValidationError{}

// Validate checks the field values on ServerMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServerMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServerMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServerMessageMultiError, or
// nil if none found.
func (m *ServerMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ServerMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SenderId

	// no validation rules for MessageText

	if len(errors) > 0 {
		return ServerMessageMultiError(errors)
	}

	return nil
}

// ServerMessageMultiError is an error wrapping multiple validation errors
// returned by ServerMessage.ValidateAll() if the designated constraints
// aren't met.
type ServerMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerMessageMultiError) AllErrors() []error { return m }

// ServerMessageValidationError is the validation error returned by
// ServerMessage.Validate if the designated constraints aren't met.
type ServerMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerMessageValidationError) ErrorName() string { return "ServerMessageValidationError" }

// Error satisfies the builtin error interface
func (e ServerMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerMessageValidationError{}

// Validate checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ChatMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ChatMessageMultiError, or
// nil if none found.
func (m *ChatMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for MessageText

	if len(errors) > 0 {
		return ChatMessageMultiError(errors)
	}

	return nil
}

// ChatMessageMultiError is an error wrapping multiple validation errors
// returned by ChatMessage.ValidateAll() if the designated constraints aren't met.
type ChatMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatMessageMultiError) AllErrors() []error { return m }

// ChatMessageValidationError is the validation error returned by
// ChatMessage.Validate if the designated constraints aren't met.
type ChatMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatMessageValidationError) ErrorName() string { return "ChatMessageValidationError" }

// Error satisfies the builtin error interface
func (e ChatMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatMessageValidationError{}
