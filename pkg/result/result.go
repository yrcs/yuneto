package result

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	thttp "github.com/go-kratos/kratos/v2/transport/http"
	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const jsonContentType = "application/json; charset=utf-8"

// httpError is an HTTP error.
type httpError struct {
	Code      int32             `json:"-"`
	Requestid string            `json:"requestid"`
	Reason    string            `json:"reason"`
	Message   string            `json:"message"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	Cause     string            `json:"cause,omitempty"`
}

// fromError try to convert an error to *HTTPError.
func fromError(err error) *httpError {
	if err == nil {
		return nil
	}
	requestid := strings.ToUpper(uuid.NewString())
	if se := new(errors.Error); errors.As(err, &se) {
		cause := ""
		if causeErr := se.Unwrap(); causeErr != nil {
			cause = causeErr.Error()
		}
		return &httpError{
			Code:      se.Code,
			Requestid: requestid,
			Reason:    se.Reason,
			Message:   se.Message,
			Metadata:  se.Metadata,
			Cause:     cause,
		}
	}
	gs, ok := status.FromError(err)
	if !ok {
		return &httpError{
			Code:      errors.UnknownCode,
			Requestid: requestid,
			Reason:    errors.UnknownReason,
			Message:   err.Error(),
		}
	}
	ret := errors.New(
		httpstatus.FromGRPCCode(gs.Code()),
		errors.UnknownReason,
		gs.Message(),
	)
	for _, detail := range gs.Details() {
		switch d := detail.(type) {
		case *errdetails.ErrorInfo:
			ret.Reason = d.Reason
			ret = ret.WithMetadata(d.Metadata)
		}
	}
	return &httpError{
		Code:      ret.Code,
		Requestid: requestid,
		Reason:    ret.Reason,
		Message:   ret.Message,
		Metadata:  ret.Metadata,
	}
}

type errorRender struct {
	body        []byte
	contentType string
}

// Render (JSON) writes data with custom ContentType.
func (er *errorRender) Render(w http.ResponseWriter) error {
	er.WriteContentType(w)
	_, err := w.Write(er.body)
	return err
}

// WriteContentType (JSON) writes JSON ContentType.
func (er *errorRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", er.contentType)
}

// Error encodes the object to the HTTP response.
func Error(ctx *gin.Context, err error) {
	if err == nil {
		ctx.Status(http.StatusOK)
		return
	}
	codec, _ := thttp.CodecForRequest(ctx.Request, "Accept")
	se := fromError(err)
	body, err := codec.Marshal(se)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &httpError{
			Requestid: uuid.NewString(),
			Reason:    "JSON_MARSHAL_FAILED",
			Message:   "JSON 序列化错误",
			Cause:     err.Error(),
		})
		return
	}
	code := int(se.Code)
	ctx.Render(code, &errorRender{body, jsonContentType})
}

func Result(ctx *gin.Context, obj any) {
	if _, ok := obj.(*emptypb.Empty); ok {
		return
	}
	codec, _ := thttp.CodecForRequest(ctx.Request, "Accept")
	body, err := codec.Marshal(obj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &httpError{
			Requestid: uuid.NewString(),
			Reason:    "JSON_MARSHAL_FAILED",
			Message:   "JSON 序列化错误",
			Cause:     err.Error(),
		})
		return
	}
	ctx.Writer.Header().Set("Content-Type", jsonContentType)
	ctx.Writer.Write(body)
}
