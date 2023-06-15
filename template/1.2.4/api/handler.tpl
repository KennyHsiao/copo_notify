package {{.PkgName}}

import (
	"net/http"
    "com.copo/copo_notify/common/response"
	{{if .After1_1_10}}"github.com/tal-tech/go-zero/rest/httpx"{{end}}
	{{.ImportPackages}}
)

func {{.HandlerName}}(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}

        if err := httpx.ParseJsonBody(r, &req); err != nil {
            response.Json(w, r, response.FAIL, nil, err)
            return
        }

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), ctx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}req{{end}})
		if err != nil {
			response.Json(w, r, err.Error(), nil, err)
		} else {
			{{if .HasResp}}response.Json(w, r, response.SUCCESS, resp, err){{else}}response.Json(w, r, response.SUCCESS, nil, err){{end}}
		}
	}
}
