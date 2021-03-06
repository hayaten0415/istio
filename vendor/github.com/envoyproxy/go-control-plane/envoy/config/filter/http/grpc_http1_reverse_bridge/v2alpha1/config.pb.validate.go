// Code generated by protoc-gen-validate
// source: envoy/config/filter/http/grpc_http1_reverse_bridge/v2alpha1/config.proto
// DO NOT EDIT!!!

package v2

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on FilterConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FilterConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetContentType()) < 1 {
		return FilterConfigValidationError{
			Field:  "ContentType",
			Reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for WithholdGrpcFrames

	return nil
}

// FilterConfigValidationError is the validation error returned by
// FilterConfig.Validate if the designated constraints aren't met.
type FilterConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FilterConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FilterConfigValidationError{}
