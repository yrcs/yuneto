// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pagination/pagination.proto

package pagination

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

// Validate checks the field values on PagingRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PagingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PagingRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PagingRequestMultiError, or
// nil if none found.
func (m *PagingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PagingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetQuery()) > 0 {

		{
			sorted_keys := make([]string, len(m.GetQuery()))
			i := 0
			for key := range m.GetQuery() {
				sorted_keys[i] = key
				i++
			}
			sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
			for _, key := range sorted_keys {
				val := m.GetQuery()[key]
				_ = val

				if utf8.RuneCountInString(key) < 1 {
					err := PagingRequestValidationError{
						field:  fmt.Sprintf("Query[%v]", key),
						reason: "value length must be at least 1 runes",
					}
					if !all {
						return err
					}
					errors = append(errors, err)
				}

				if utf8.RuneCountInString(val) < 1 {
					err := PagingRequestValidationError{
						field:  fmt.Sprintf("Query[%v]", key),
						reason: "value length must be at least 1 runes",
					}
					if !all {
						return err
					}
					errors = append(errors, err)
				}

			}
		}

	}

	if len(m.GetOrderBy()) > 0 {

		{
			sorted_keys := make([]string, len(m.GetOrderBy()))
			i := 0
			for key := range m.GetOrderBy() {
				sorted_keys[i] = key
				i++
			}
			sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
			for _, key := range sorted_keys {
				val := m.GetOrderBy()[key]
				_ = val

				if utf8.RuneCountInString(key) < 1 {
					err := PagingRequestValidationError{
						field:  fmt.Sprintf("OrderBy[%v]", key),
						reason: "value length must be at least 1 runes",
					}
					if !all {
						return err
					}
					errors = append(errors, err)
				}

				if _, ok := Order_name[int32(val)]; !ok {
					err := PagingRequestValidationError{
						field:  fmt.Sprintf("OrderBy[%v]", key),
						reason: "value must be one of the defined enum values",
					}
					if !all {
						return err
					}
					errors = append(errors, err)
				}

			}
		}

	}

	if m.Page != nil {

		if m.GetPage() < 1 {
			err := PagingRequestValidationError{
				field:  "Page",
				reason: "value must be greater than or equal to 1",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.PageSize != nil {

		if val := m.GetPageSize(); val < 1 || val > 1000 {
			err := PagingRequestValidationError{
				field:  "PageSize",
				reason: "value must be inside range [1, 1000]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return PagingRequestMultiError(errors)
	}

	return nil
}

// PagingRequestMultiError is an error wrapping multiple validation errors
// returned by PagingRequest.ValidateAll() if the designated constraints
// aren't met.
type PagingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PagingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PagingRequestMultiError) AllErrors() []error { return m }

// PagingRequestValidationError is the validation error returned by
// PagingRequest.Validate if the designated constraints aren't met.
type PagingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PagingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PagingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PagingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PagingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PagingRequestValidationError) ErrorName() string { return "PagingRequestValidationError" }

// Error satisfies the builtin error interface
func (e PagingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PagingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PagingRequestValidationError{}

// Validate checks the field values on PagingReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PagingReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PagingReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PagingReplyMultiError, or
// nil if none found.
func (m *PagingReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PagingReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PagingReplyValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PagingReplyValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PagingReplyValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PagingReplyMultiError(errors)
	}

	return nil
}

// PagingReplyMultiError is an error wrapping multiple validation errors
// returned by PagingReply.ValidateAll() if the designated constraints aren't met.
type PagingReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PagingReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PagingReplyMultiError) AllErrors() []error { return m }

// PagingReplyValidationError is the validation error returned by
// PagingReply.Validate if the designated constraints aren't met.
type PagingReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PagingReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PagingReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PagingReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PagingReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PagingReplyValidationError) ErrorName() string { return "PagingReplyValidationError" }

// Error satisfies the builtin error interface
func (e PagingReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagingReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PagingReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PagingReplyValidationError{}
