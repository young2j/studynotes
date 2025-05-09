// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: jwt/v1/jwt.proto

package jwtv1

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

// Validate checks the field values on GetTokenRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTokenRequestMultiError, or nil if none found.
func (m *GetTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticket

	if len(errors) > 0 {
		return GetTokenRequestMultiError(errors)
	}
	return nil
}

// GetTokenRequestMultiError is an error wrapping multiple validation errors
// returned by GetTokenRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTokenRequestMultiError) AllErrors() []error { return m }

// GetTokenRequestValidationError is the validation error returned by
// GetTokenRequest.Validate if the designated constraints aren't met.
type GetTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTokenRequestValidationError) ErrorName() string { return "GetTokenRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTokenRequestValidationError{}

// Validate checks the field values on GetTokenResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTokenResponseMultiError, or nil if none found.
func (m *GetTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return GetTokenResponseMultiError(errors)
	}
	return nil
}

// GetTokenResponseMultiError is an error wrapping multiple validation errors
// returned by GetTokenResponse.ValidateAll() if the designated constraints
// aren't met.
type GetTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTokenResponseMultiError) AllErrors() []error { return m }

// GetTokenResponseValidationError is the validation error returned by
// GetTokenResponse.Validate if the designated constraints aren't met.
type GetTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTokenResponseValidationError) ErrorName() string { return "GetTokenResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTokenResponseValidationError{}

// Validate checks the field values on RefreshTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RefreshTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RefreshTokenRequestMultiError, or nil if none found.
func (m *RefreshTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return RefreshTokenRequestMultiError(errors)
	}
	return nil
}

// RefreshTokenRequestMultiError is an error wrapping multiple validation
// errors returned by RefreshTokenRequest.ValidateAll() if the designated
// constraints aren't met.
type RefreshTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshTokenRequestMultiError) AllErrors() []error { return m }

// RefreshTokenRequestValidationError is the validation error returned by
// RefreshTokenRequest.Validate if the designated constraints aren't met.
type RefreshTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshTokenRequestValidationError) ErrorName() string {
	return "RefreshTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RefreshTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshTokenRequestValidationError{}

// Validate checks the field values on RefreshTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RefreshTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RefreshTokenResponseMultiError, or nil if none found.
func (m *RefreshTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return RefreshTokenResponseMultiError(errors)
	}
	return nil
}

// RefreshTokenResponseMultiError is an error wrapping multiple validation
// errors returned by RefreshTokenResponse.ValidateAll() if the designated
// constraints aren't met.
type RefreshTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshTokenResponseMultiError) AllErrors() []error { return m }

// RefreshTokenResponseValidationError is the validation error returned by
// RefreshTokenResponse.Validate if the designated constraints aren't met.
type RefreshTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshTokenResponseValidationError) ErrorName() string {
	return "RefreshTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RefreshTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshTokenResponseValidationError{}
