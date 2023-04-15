// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/cluster/v1/cluster.proto

package v1

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

// Validate checks the field values on Traefik with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Traefik) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Traefik with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TraefikMultiError, or nil if none found.
func (m *Traefik) ValidateAll() error {
	return m.validate(true)
}

func (m *Traefik) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for HttpNodePort

	// no validation rules for HttpsNodePort

	if len(errors) > 0 {
		return TraefikMultiError(errors)
	}

	return nil
}

// TraefikMultiError is an error wrapping multiple validation errors returned
// by Traefik.ValidateAll() if the designated constraints aren't met.
type TraefikMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TraefikMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TraefikMultiError) AllErrors() []error { return m }

// TraefikValidationError is the validation error returned by Traefik.Validate
// if the designated constraints aren't met.
type TraefikValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TraefikValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TraefikValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TraefikValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TraefikValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TraefikValidationError) ErrorName() string { return "TraefikValidationError" }

// Error satisfies the builtin error interface
func (e TraefikValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTraefik.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TraefikValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TraefikValidationError{}

// Validate checks the field values on Vcluster with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Vcluster) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Vcluster with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VclusterMultiError, or nil
// if none found.
func (m *Vcluster) ValidateAll() error {
	return m.validate(true)
}

func (m *Vcluster) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for HttpsNodePort

	if len(errors) > 0 {
		return VclusterMultiError(errors)
	}

	return nil
}

// VclusterMultiError is an error wrapping multiple validation errors returned
// by Vcluster.ValidateAll() if the designated constraints aren't met.
type VclusterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VclusterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VclusterMultiError) AllErrors() []error { return m }

// VclusterValidationError is the validation error returned by
// Vcluster.Validate if the designated constraints aren't met.
type VclusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VclusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VclusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VclusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VclusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VclusterValidationError) ErrorName() string { return "VclusterValidationError" }

// Error satisfies the builtin error interface
func (e VclusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVcluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VclusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VclusterValidationError{}

// Validate checks the field values on SaveRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SaveRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SaveRequestMultiError, or
// nil if none found.
func (m *SaveRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClusterName

	// no validation rules for InsecureSkipCheck

	if all {
		switch v := interface{}(m.GetBody()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveRequestValidationError{
					field:  "Body",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveRequestValidationError{
					field:  "Body",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBody()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveRequestValidationError{
				field:  "Body",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SaveRequestMultiError(errors)
	}

	return nil
}

// SaveRequestMultiError is an error wrapping multiple validation errors
// returned by SaveRequest.ValidateAll() if the designated constraints aren't met.
type SaveRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveRequestMultiError) AllErrors() []error { return m }

// SaveRequestValidationError is the validation error returned by
// SaveRequest.Validate if the designated constraints aren't met.
type SaveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveRequestValidationError) ErrorName() string { return "SaveRequestValidationError" }

// Error satisfies the builtin error interface
func (e SaveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveRequestValidationError{}

// Validate checks the field values on SaveReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SaveReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SaveReplyMultiError, or nil
// if none found.
func (m *SaveReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	if len(errors) > 0 {
		return SaveReplyMultiError(errors)
	}

	return nil
}

// SaveReplyMultiError is an error wrapping multiple validation errors returned
// by SaveReply.ValidateAll() if the designated constraints aren't met.
type SaveReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveReplyMultiError) AllErrors() []error { return m }

// SaveReplyValidationError is the validation error returned by
// SaveReply.Validate if the designated constraints aren't met.
type SaveReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveReplyValidationError) ErrorName() string { return "SaveReplyValidationError" }

// Error satisfies the builtin error interface
func (e SaveReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveReplyValidationError{}

// Validate checks the field values on DeleteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteRequestMultiError, or
// nil if none found.
func (m *DeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProductName

	// no validation rules for ClusterName

	// no validation rules for InsecureSkipCheck

	if len(errors) > 0 {
		return DeleteRequestMultiError(errors)
	}

	return nil
}

// DeleteRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRequestMultiError) AllErrors() []error { return m }

// DeleteRequestValidationError is the validation error returned by
// DeleteRequest.Validate if the designated constraints aren't met.
type DeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRequestValidationError) ErrorName() string { return "DeleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRequestValidationError{}

// Validate checks the field values on DeleteReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteReplyMultiError, or
// nil if none found.
func (m *DeleteReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	if len(errors) > 0 {
		return DeleteReplyMultiError(errors)
	}

	return nil
}

// DeleteReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteReply.ValidateAll() if the designated constraints aren't met.
type DeleteReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteReplyMultiError) AllErrors() []error { return m }

// DeleteReplyValidationError is the validation error returned by
// DeleteReply.Validate if the designated constraints aren't met.
type DeleteReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteReplyValidationError) ErrorName() string { return "DeleteReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteReplyValidationError{}

// Validate checks the field values on SaveRequest_Body with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SaveRequest_Body) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveRequest_Body with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveRequest_BodyMultiError, or nil if none found.
func (m *SaveRequest_Body) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveRequest_Body) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetApiServer()) < 1 {
		err := SaveRequest_BodyValidationError{
			field:  "ApiServer",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetClusterKind()) < 1 {
		err := SaveRequest_BodyValidationError{
			field:  "ClusterKind",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _SaveRequest_Body_ClusterType_InLookup[m.GetClusterType()]; !ok {
		err := SaveRequest_BodyValidationError{
			field:  "ClusterType",
			reason: "value must be in list [physical virtual]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _SaveRequest_Body_Usage_InLookup[m.GetUsage()]; !ok {
		err := SaveRequest_BodyValidationError{
			field:  "Usage",
			reason: "value must be in list [host worker]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for HostCluster

	// no validation rules for ArgocdHost

	if all {
		switch v := interface{}(m.GetVcluster()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveRequest_BodyValidationError{
					field:  "Vcluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveRequest_BodyValidationError{
					field:  "Vcluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetVcluster()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveRequest_BodyValidationError{
				field:  "Vcluster",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTraefik()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveRequest_BodyValidationError{
					field:  "Traefik",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveRequest_BodyValidationError{
					field:  "Traefik",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTraefik()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveRequest_BodyValidationError{
				field:  "Traefik",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Kubeconfig

	if len(errors) > 0 {
		return SaveRequest_BodyMultiError(errors)
	}

	return nil
}

// SaveRequest_BodyMultiError is an error wrapping multiple validation errors
// returned by SaveRequest_Body.ValidateAll() if the designated constraints
// aren't met.
type SaveRequest_BodyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveRequest_BodyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveRequest_BodyMultiError) AllErrors() []error { return m }

// SaveRequest_BodyValidationError is the validation error returned by
// SaveRequest_Body.Validate if the designated constraints aren't met.
type SaveRequest_BodyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveRequest_BodyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveRequest_BodyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveRequest_BodyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveRequest_BodyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveRequest_BodyValidationError) ErrorName() string { return "SaveRequest_BodyValidationError" }

// Error satisfies the builtin error interface
func (e SaveRequest_BodyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveRequest_Body.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveRequest_BodyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveRequest_BodyValidationError{}

var _SaveRequest_Body_ClusterType_InLookup = map[string]struct{}{
	"physical": {},
	"virtual":  {},
}

var _SaveRequest_Body_Usage_InLookup = map[string]struct{}{
	"host":   {},
	"worker": {},
}
