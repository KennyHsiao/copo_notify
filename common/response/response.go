package response

import (
	"com.copo/copo_notify/common/errorz"
	"errors"
	"fmt"
	"github.com/gioco-play/easy-i18n/i18n"
	"github.com/tal-tech/go-zero/rest/httpx"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/language"
	"net/http"

	_ "com.copo/copo_notify/locales"
)

type Body struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Trace   string      `json:"trace"`
}

func Json(w http.ResponseWriter, r *http.Request, code string, resp interface{}, err error) {
	var body Body
	span := trace.SpanFromContext(r.Context())
	defer span.End()

	i18n.SetLang(language.Chinese)
	body.Code = code
	body.Message = i18n.Sprintf(code)
	if err != nil {
		if v, ok := err.(*errorz.Err); ok && v.GetMessage() != "" {
			span.RecordError(errors.New(fmt.Sprintf("(%s)%s", code, v.GetMessage())))
		} else {
			span.RecordError(errors.New(fmt.Sprintf("(%s)%s", code, body.Message)))
		}
	} else {
		body.Data = resp
	}
	body.Trace = span.SpanContext().TraceID().String()
	httpx.OkJson(w, body)
}

type State struct {
	Code string
	Desc string
}
