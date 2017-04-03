package ginpprof

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

// Wrap adds several routes from package `net/http/pprof` to *gin.Engine object
func Wrap(router *gin.Engine) {
	router.GET("/debug/pprof/", IndexHandler())
	router.GET("/debug/pprof/heap", HeapHandler())
	router.GET("/debug/pprof/goroutine", GoroutineHandler())
	router.GET("/debug/pprof/block", BlockHandler())
	router.GET("/debug/pprof/threadcreate", ThreadCreateHandler())
	router.GET("/debug/pprof/cmdline", CmdlineHandler())
	router.GET("/debug/pprof/profile", ProfileHandler())
	router.GET("/debug/pprof/symbol", SymbolHandler())
	router.GET("/debug/pprof/trace", TraceHandler())
}

// Wrapper make sure we are backward compatible
var Wrapper = Wrap

// IndexHandler will pass the call from /debug/pprof to pprof
func IndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Index(ctx.Writer, ctx.Request)
	}
}

// HeapHandler will pass the call from /debug/pprof/heap to pprof
func HeapHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("heap").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// GoroutineHandler will pass the call from /debug/pprof/goroutine to pprof
func GoroutineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("goroutine").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// BlockHandler will pass the call from /debug/pprof/block to pprof
func BlockHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("block").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// ThreadCreateHandler will pass the call from /debug/pprof/threadcreate to pprof
func ThreadCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// CmdlineHandler will pass the call from /debug/pprof/cmdline to pprof
func CmdlineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Cmdline(ctx.Writer, ctx.Request)
	}
}

// ProfileHandler will pass the call from /debug/pprof/profile to pprof
func ProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Profile(ctx.Writer, ctx.Request)
	}
}

// SymbolHandler will pass the call from /debug/pprof/symbol to pprof
func SymbolHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Symbol(ctx.Writer, ctx.Request)
	}
}

// TraceHandler will pass the call from /debug/pprof/trace to pprof
func TraceHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Trace(ctx.Writer, ctx.Request)
	}
}
