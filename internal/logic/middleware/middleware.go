package middleware

import (
	"Literary_Snowlands/internal/service"
	service2 "Literary_Snowlands/utility"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"strings"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

type JsonResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
	Ok      bool        `json:"ok"`
}

func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	//requestId := uuid.New().String()
	//ctx := context.WithValue(r.Context(), "RequestId", requestId)
	//r.SetCtx(ctx)

	r.Middleware.Next()
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
		ok   = true
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 将sql.ErrNoRows视为无错误状态
			err = nil             // 将错误置为nil
			msg = "No data found" // 提供一个默认的信息
			code = gcode.CodeOK   // 设置正确的状态码
			ok = true             // 没有错误
			res = nil             // 根据情况，你可能需要设置一个默认的返回值
		} else if code == gcode.CodeNil {
			code = gcode.CodeInternalError
			msg = "Internal server error"
			ok = false
		} else {
			msg = err.Error()
			ok = false
		}
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}

	r.Response.WriteJson(JsonResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
		Ok:      ok,
	})
}

//func (s *sMiddleware) JWTMiddleware(r *ghttp.Request) {
//	tokenString := r.Header.Get("Authorization")
//	fmt.Println("拿到的前端的token")
//	if tokenString == "" {
//		r.Response.WriteStatus(http.StatusUnauthorized, "Authorization token required")
//		return
//	}
//	fmt.Println("tokenString", tokenString)
//	token, claims, err := service2.ParseJWT(tokenString)
//	fmt.Println("token", token)
//	if err != nil || token == nil || !token.Valid {
//		r.Response.WriteStatus(http.StatusUnauthorized, "Invalid or expired token")
//		return
//	}
//	userId := claims["userId"].(float64) // JWT中的数字默认是float64
//	ctx := context.WithValue(r.Context(), "userId", int(userId))
//	r.SetCtx(ctx)
//	r.Middleware.Next()
//
//}

func (s *sMiddleware) JWTMiddleware(r *ghttp.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteStatus(http.StatusUnauthorized, "Authorization header is required")
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		r.Response.WriteStatus(http.StatusUnauthorized, "Authorization header must be Bearer token")
		return
	}
	tokenString := parts[1]

	token, claims, err := service2.ParseJWT(tokenString)
	if err != nil || token == nil || !token.Valid {
		r.Response.WriteStatus(http.StatusUnauthorized, "Invalid or expired token")
		return
	}
	// Extract user ID from JWT claims and add it to the context
	userId, ok := claims["userId"].(float64)
	if !ok {
		r.Response.WriteStatus(http.StatusUnauthorized, "Invalid token claims")
		return
	}

	// Convert the userId to an integer and set it in the request's context
	ctx := context.WithValue(r.Context(), "userId", int(userId))
	r.SetCtx(ctx)

	r.Middleware.Next()
}

func (s *sMiddleware) Cors(r *ghttp.Request) {
	r.Response.Header().Set("Access-Control-Allow-Origin", "http://localhost:1024")
	r.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST,PUT,DELETE,OPTIONS")
	r.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	r.Response.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method == "OPTIONS" {
		r.Response.WriteHeader(http.StatusOK)
		return
	}
	r.Middleware.Next()
}
