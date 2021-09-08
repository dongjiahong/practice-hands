package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// NewLogger constructs a logger. It's just a regular Go function, without any
// special relationship to Fx.
//
// Since it returns a *log.Logger, Fx will treat NewLogger as the constructor
// function for the standard library's logger. (We'll see how to integrate
// NewLogger into an Fx application in the main function.) Since NewLogger
// doesn't have any parameters, Fx will infer that loggers don't depend on any
// other types - we can create them from thin air.
//
// Fx calls constructors lazily, so NewLogger will only be called only if some
// other function needs a logger. Once instantiated, the logger is cached and
// reused - within the application, it's effectively a singleton.
//
// By default, Fx applications only allow one constructor for each type. See
// the documentation of the In and Out types for ways around this restriction.
// ----- 翻译
// NewLogger 构造一个记录器。它只是一个普通的 Go 函数，没有任何预期声明，
// 找到 NewLogger 由于它返回一个 *log.Logger，Fx 会将 NewLogger 视为标准库记录器的构造函数。
//（我们将在 main 函数中看到如何将 NewLogger 集成到 Fx 应用程序中。）
// 由于 NewLogger 没有任何参数，Fx 将推断记录器不依赖于任何其他类型 - 我们可以凭空创建它们。
// Fx 懒惰地调用构造函数，因此只有在某些其他函数需要记录器时才会调用 NewLogger。一旦实例化，
// 记录器就会被缓存和重用——在应用程序中，它实际上是一个单例。
func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	return logger
}

// NewHandler constructs a simple HTTP handler. Since it returns an
// http.Handler, Fx will treat NewHandler as the constructor for the
// http.Handler type.
//
// Like many Go functions, NewHandler also returns an error. If the error is
// non-nil, Go convention tells the caller to assume that NewHandler failed
// and the other returned values aren't safe to use. Fx understands this
// idiom, and assumes that any function whose last return value is an error
// follows this convention.
//
// Unlike NewLogger, NewHandler has formal parameters. Fx will interpret these
// parameters as dependencies: in order to construct an HTTP handler,
// NewHandler needs a logger. If the application has access to a *log.Logger
// constructor (like NewLogger above), it will use that constructor or its
// cached output and supply a logger to NewHandler. If the application doesn't
// know how to construct a logger and needs an HTTP handler, it will fail to
// start.
//
// Functions may also return multiple objects. For example, we could combine
// NewHandler and NewLogger into a single function:
//
//   func NewHandlerAndLogger() (*log.Logger, http.Handler, error)
//
// Fx also understands this idiom, and would treat NewHandlerAndLogger as the
// constructor for both the *log.Logger and http.Handler types. Just like
// constructors for a single type, NewHandlerAndLogger would be called at most
// once, and both the handler and the logger would be cached and reused as
// necessary.
//
// NewHandler 构造了一个简单的 HTTP 处理程序。因为它返回一个
//http.Handler，Fx 会将 NewHandler 视为构造函数
//http.Handler 类型。
//
//与许多 Go 函数一样，NewHandler 也会返回错误。如果错误是
//非零，Go 约定告诉调用者假设 NewHandler 失败
//并且其他返回值不安全使用。 Fx 明白这一点
//习惯用法，并假设任何最后一个返回值是错误的函数
//遵循这个约定。
//
//与 NewLogger 不同，NewHandler 具有形式参数。 Fx 会解释这些
//参数作为依赖项：为了构造 HTTP 处理程序，
//NewHandler 需要一个记录器。如果应用程序有权访问 *log.Logger
//构造函数（如上面的 NewLogger），它将使用该构造函数或其
//缓存输出并向 NewHandler 提供记录器。如果应用程序没有
//知道如何构建记录器并需要一个 HTTP 处理程序，它将无法
//开始。
//
//函数也可以返回多个对象。例如，我们可以结合
//NewHandler 和 NewLogger 合并为一个函数：
//
//func NewHandlerAndLogger() (*log.Logger, http.Handler, error)
//
//Fx 也理解这个习语，并将 NewHandlerAndLogger 视为
//*log.Logger 和 http.Handler 类型的构造函数。就像
//单一类型的构造函数，最多会调用 NewHandlerAndLogger
//一次，处理程序和记录器都将被缓存并重用为
//必要的。
func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}

// NewMux constructs an HTTP mux. Like NewHandler, it depends on *log.Logger.
// However, it also depends on the Fx-specific Lifecycle interface.
//
// A Lifecycle is available in every Fx application. It lets objects hook into
// the application's start and stop phases. In a non-Fx application, the main
// function often includes blocks like this:
//
//   srv, err := NewServer() // some long-running network server
//   if err != nil {
//     log.Fatalf("failed to construct server: %v", err)
//   }
//   // Construct other objects as necessary.
//   go srv.Start()
//   defer srv.Stop()
//
// In this example, the programmer explicitly constructs a bunch of objects,
// crashing the program if any of the constructors encounter unrecoverable
// errors. Once all the objects are constructed, we start any background
// goroutines and defer cleanup functions.
//
// Fx removes the manual object construction with dependency injection. It
// replaces the inline goroutine spawning and deferred cleanups with the
// Lifecycle type.
//
// Here, NewMux makes an HTTP mux available to other functions. Since
// constructors are called lazily, we know that NewMux won't be called unless
// some other function wants to register a handler. This makes it easy to use
// Fx's Lifecycle to start an HTTP server only if we have handlers registered.
//
//NewMux 构造了一个 HTTP 多路复用器。和 NewHandler 一样，它依赖于 *log.Logger。
//但是，它也取决于特定于 Fx 的 Lifecycle 接口。
//
//每个 Fx 应用程序中都有生命周期。它让对象挂钩
//应用程序的启动和停止阶段。在非 Fx 应用程序中，主要
//函数通常包括这样的块：
//
//   srv, err := NewServer() // some long-running network server
//   if err != nil {
//     log.Fatalf("failed to construct server: %v", err)
//   }
//   // Construct other objects as necessary.
//   go srv.Start()
//   defer srv.Stop()
//
//在这个例子中，程序员显式地构造了一堆对象，
//如果任何构造函数遇到不可恢复的情况，则程序崩溃
//错误。构建完所有对象后，我们开始任何背景
//goroutines 和延迟清理功能。
//
//Fx 移除了依赖注入的手动对象构造。它
//将内联 goroutine 生成和延迟清理替换为
//生命周期类型。
//
//在这里，NewMux 使 HTTP 多路复用器可用于其他功能。自从
//构造函数被懒惰地调用，我们知道除非 NewMux 不会被调用
//其他一些函数想要注册一个处理程序。这使它易于使用
//仅当我们注册了处理程序时，Fx 的生命周期才能启动 HTTP 服务器。
func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Print("Executing NewMux.")
	// First, we construct the mux and server. We don't want to start the server
	// until all handlers are registered.
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// If NewMux is called, we know that another function is using the mux. In
	// that case, we'll use the Lifecycle type to register a Hook that starts
	// and stops our HTTP server.
	//
	// Hooks are executed in dependency order. At startup, NewLogger's hooks run
	// before NewMux's. On shutdown, the order is reversed.
	//
	// Returning an error from OnStart hooks interrupts application startup. Fx
	// immediately runs the OnStop portions of any successfully-executed OnStart
	// hooks (so that types which started cleanly can also shut down cleanly),
	// then exits.
	//
	// Returning an error from OnStop hooks logs a warning, but Fx continues to
	// run the remaining hooks.
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

// Register mounts our HTTP handler on the mux.
//
// Register is a typical top-level application function: it takes a generic
// type like ServeMux, which typically comes from a third-party library, and
// introduces it to a type that contains our application logic. In this case,
// that introduction consists of registering an HTTP handler. Other typical
// examples include registering RPC procedures and starting queue consumers.
//
// Fx calls these functions invocations, and they're treated differently from
// the constructor functions above. Their arguments are still supplied via
// dependency injection and they may still return an error to indicate
// failure, but any other return values are ignored.
//
// Unlike constructors, invocations are called eagerly. See the main function
// below for details.
func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}

func main() {
	app := fx.New(
		// Provide all the constructors we need, which teaches Fx how we'd like to
		// construct the *log.Logger, http.Handler, and *http.ServeMux types.
		// Remember that constructors are called lazily, so this block doesn't do
		// much on its own.
		fx.Provide(
			NewLogger,
			NewHandler,
			NewMux,
		),
		// Since constructors are called lazily, we need some invocations to
		// kick-start our application. In this case, we'll use Register. Since it
		// depends on an http.Handler and *http.ServeMux, calling it requires Fx
		// to build those types using the constructors above. Since we call
		// NewMux, we also register Lifecycle hooks to start and stop an HTTP
		// server.
		fx.Invoke(Register),

		// This is optional. With this, you can control where Fx logs
		// its events. In this case, we're using a NopLogger to keep
		// our test silent. Normally, you'll want to use an
		// fxevent.ZapLogger or an fxevent.ConsoleLogger.
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)

	// In a typical application, we could just use app.Run() here. Since we
	// don't want this example to run forever, we'll use the more-explicit Start
	// and Stop.
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	// Normally, we'd block here with <-app.Done(). Instead, we'll make an HTTP
	// request to demonstrate that our server is running.
	if _, err := http.Get("http://localhost:8080/"); err != nil {
		log.Fatal(err)
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}

}
