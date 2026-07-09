package v1

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

type ErrorReasonOption func(*errorReasonOptions)

type errorReasonOptions struct {
	message string
	err     error
}

func WithError(err error) ErrorReasonOption {
	return func(o *errorReasonOptions) {
		o.err = err
	}
}

func WithFmtMsg(format string, args ...interface{}) ErrorReasonOption {
	return func(o *errorReasonOptions) {
		o.message = fmt.Sprintf(format, args...)
	}
}

func buildErrorReasonMessage(defaultMessage string, opts ...ErrorReasonOption) string {
	o := &errorReasonOptions{message: defaultMessage}
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
	if o.err == nil {
		return o.message
	}
	if o.message == "" || o.message == defaultMessage {
		return o.err.Error()
	}
	return fmt.Sprintf("%s: %v", o.message, o.err)
}

func newErrorReason(fn func(string, ...interface{}) *errors.Error, defaultMessage string, opts ...ErrorReasonOption) *errors.Error {
	return fn("%s", buildErrorReasonMessage(defaultMessage, opts...))
}

func ErrorReasonRequestCanceledErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorRequestCanceledErr, "RequestCanceledErr", opts...)
}
func ErrorReasonRequestTimeoutErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorRequestTimeoutErr, "RequestTimeoutErr", opts...)
}
func ErrorReasonRequestFrequentErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorRequestFrequentErr, "RequestFrequentErr", opts...)
}
func ErrorReasonAPIInternalErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorAPIInternalErr, "APIInternalErr", opts...)
}
func ErrorReasonAPIThirdErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorAPIThirdErr, "APIThirdErr", opts...)
}
func ErrorReasonParamError(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorParamError, "ParamError", opts...)
}
func ErrorReasonDataSQLError(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorDataSQLError, "DataSQLError", opts...)
}
func ErrorReasonDataRedisErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorDataRedisErr, "DataRedisErr", opts...)
}
func ErrorReasonDataMQErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorDataMQErr, "DataMQErr", opts...)
}
func ErrorReasonDataFormattingError(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorDataFormattingError, "DataFormattingError", opts...)
}
func ErrorReasonDataProcessingError(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorDataProcessingError, "DataProcessingError", opts...)
}
func ErrorReasonDataRecordNotFound(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorDataRecordNotFound, "DataRecordNotFound", opts...)
}
func ErrorReasonDataDuplicateRecord(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorDataDuplicateRecord, "DataDuplicateRecord", opts...)
}
func ErrorReasonTokenNotRequest(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorTokenNotRequest, "TokenNotRequest", opts...)
}
func ErrorReasonTokenFormatErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorTokenFormatErr, "TokenFormatErr", opts...)
}
func ErrorReasonTokenExpiredErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorTokenExpiredErr, "TokenExpiredErr", opts...)
}
func ErrorReasonTokenInvalidErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorTokenInvalidErr, "TokenInvalidErr", opts...)
}
func ErrorReasonTokenErr(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorTokenErr, "TokenErr", opts...)
}
func ErrorReasonAccountAlreadyExists(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorAccountAlreadyExists, "AccountAlreadyExists", opts...)
}
func ErrorReasonAccountNotFound(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorAccountNotFound, "AccountNotFound", opts...)
}
func ErrorReasonAccountPasswordError(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorAccountPasswordError, "AccountPasswordError", opts...)
}
func ErrorReasonAccountNoDataPermission(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorAccountNoDataPermission, "AccountNoDataPermission", opts...)
}
func ErrorReasonMenuOperationFailed(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorMenuOperationFailed, "MenuOperationFailed", opts...)
}
func ErrorReasonMaterialUploadFailed(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorMaterialUploadFailed, "MaterialUploadFailed", opts...)
}
func ErrorReasonStorageNotFound(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorStorageNotFound, "StorageNotFound", opts...)
}
func ErrorReasonStorageGetConfigFailed(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorStorageGetConfigFailed, "StorageGetConfigFailed", opts...)
}
func ErrorReasonSmsFrequencyLimit(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorSmsFrequencyLimit, "SmsFrequencyLimit", opts...)
}
func ErrorReasonSmsCodeInvalid(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorSmsCodeInvalid, "SmsCodeInvalid", opts...)
}
func ErrorReasonUnauthorized(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorUnauthorized, "Unauthorized", opts...)
}
func ErrorReasonActivationCodeNotFound(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorActivationCodeNotFound, "ActivationCodeNotFound", opts...)
}
func ErrorReasonActivationCodeNotRedeemable(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorActivationCodeNotRedeemable, "ActivationCodeNotRedeemable", opts...)
}
func ErrorReasonActivationCodeProductConfigInvalid(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorActivationCodeProductConfigInvalid, "ActivationCodeProductConfigInvalid", opts...)
}
func ErrorReasonUserMembershipNotFound(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorUserMembershipNotFound, "UserMembershipNotFound", opts...)
}
func ErrorReasonUserMembershipStatusInvalid(opts ...ErrorReasonOption) *errors.Error {
	return newErrorReason(ErrorUserMembershipStatusInvalid, "UserMembershipStatusInvalid", opts...)
}
