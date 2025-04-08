package errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SendHttpError(w http.ResponseWriter, err error) {
	st, ok := status.FromError(err)
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	code := st.Code()
	message := st.Message()

	var httpStatus int
	switch code {
	case codes.OK:
		httpStatus = http.StatusOK
	case codes.Canceled:
		httpStatus = http.StatusRequestTimeout
	case codes.Unknown:
		httpStatus = http.StatusInternalServerError
	case codes.InvalidArgument:
		httpStatus = http.StatusBadRequest
	case codes.DeadlineExceeded:
		httpStatus = http.StatusGatewayTimeout
	case codes.NotFound:
		httpStatus = http.StatusNotFound
	case codes.AlreadyExists:
		httpStatus = http.StatusConflict
	case codes.PermissionDenied:
		httpStatus = http.StatusForbidden
	case codes.Unauthenticated:
		httpStatus = http.StatusUnauthorized
	case codes.ResourceExhausted:
		httpStatus = http.StatusTooManyRequests
	case codes.FailedPrecondition:
		httpStatus = http.StatusPreconditionFailed
	case codes.Aborted:
		httpStatus = http.StatusConflict
	case codes.OutOfRange:
		httpStatus = http.StatusRequestedRangeNotSatisfiable
	case codes.Unimplemented:
		httpStatus = http.StatusNotImplemented
	case codes.Internal:
		httpStatus = http.StatusInternalServerError
	case codes.Unavailable:
		httpStatus = http.StatusServiceUnavailable
	case codes.DataLoss:
		httpStatus = http.StatusInternalServerError
	default:
		httpStatus = http.StatusInternalServerError
	}

	http.Error(w, message, httpStatus)
}
