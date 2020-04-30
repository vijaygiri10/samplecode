package log

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"cloud.google.com/go/logging"
)

var (
	production  = "production"
	development = "development"
	uat         = "uat"
	sandbox     = "sandbox"
	logger      *log.Logger
	ProjectID   string
	LogName     string
	env         string
)

//InitializeLogger ...
func InitializeLogger(logPath string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Execption InitializeLogger err: ", err)
		}
	}()

	env = strings.TrimSpace(os.Getenv("JETSEND_ENV"))
	if production != env && uat != env {
		createLogfileObj(logPath)
	}
	fmt.Println("Service JetSend ENV ", env)
}

type basicWriter struct {
	http.ResponseWriter
	WroteHeader bool
	Code        int
	Bytes       int
}

func wrapResponseWriter(w http.ResponseWriter) *basicWriter {
	bw := basicWriter{ResponseWriter: w, WroteHeader: false, Code: 0, Bytes: 0}
	return &bw
}

func (b *basicWriter) WriteHeader(code int) {
	if !b.WroteHeader {
		b.Code = code
		b.WroteHeader = true
		b.ResponseWriter.WriteHeader(code)
	}
}

func (b *basicWriter) Write(buf []byte) (int, error) {
	//b.WriteHeader(http.StatusOK)
	n, err := b.ResponseWriter.Write(buf)
	b.Bytes += n
	return n, err
}

func (b *basicWriter) Status() int {
	return b.Code
}

func (b *basicWriter) IsSuccess() bool {
	return b.Status() >= 200 && 200 <= b.Status()
}

func (b *basicWriter) BytesWritten() int {
	return b.Bytes
}

// Logger ...
func Logger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		lw := wrapResponseWriter(w)
		start := time.Now()

		defer func() {

			if lw.Status() == 0 {
				lw.WriteHeader(http.StatusOK)
			}

			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}

			result := fmt.Sprintf("%v  %v://%v%v %v  from %v - %v %v in %v", r.Method, scheme, r.Host, r.RequestURI, r.Proto, r.RemoteAddr, lw.Status(), lw.BytesWritten(), time.Now().Sub(start))

			if lw.IsSuccess() {
				Info(r.Context(), result)
			} else {
				Error(r.Context(), result)
			}

			recover := recover()
			if recover != nil {
				Error(r.Context(), recover, string(debug.Stack()))
			}

		}()

		h.ServeHTTP(lw, r)

	}

	return http.HandlerFunc(fn)
}

func getClient(ctx context.Context) *logging.Client {
	client, err := logging.NewClient(ctx, ProjectID)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}
	return client
}

func createLogfileObj(dir string) {
	//"/services/rest"

	// pwdDir, _ := os.Getwd()
	// dir := pwdDir + "/log"

	fileName := dir + "/" + LogName + ".log"
	fmt.Println("Log FileName : ", fileName)
	// if err := os.MkdirAll(dir, 0777); err != nil {
	// 	fmt.Println("Unable to create DIR : ", dir, " - err info : ", err)
	// }

	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to Create Log File: ", err)
		return
	}
	// defer log_file.Close()
	logger = log.New(logFile, "", log.Ldate|log.Ltime)
}

// Fatal ...
func Fatal(ctx context.Context, logMSG ...interface{}) {

	if production == env || uat == env {
		// Selects the log to write to.
		client := getClient(ctx)
		defer client.Close()

		logger := client.Logger(LogName)
		logger.StandardLogger(logging.Error).Panicln(logMSG...)
	} else {
		logger.Fatal(logMSG...)
	}

}

// Println log based on Severty level
func Println(ctx context.Context, logMSG ...interface{}) {
	if production == env || uat == env {
		client := getClient(ctx)
		defer client.Close()

		logger := client.Logger(LogName)
		logger.StandardLogger(logging.Debug).Println(logMSG...)
	} else {
		logger.Println(logMSG...)
	}
}

func Debug(ctx context.Context, logMSG ...interface{}) {
	if production == env || uat == env {
		client := getClient(ctx)
		defer client.Close()

		logger := client.Logger(LogName)
		logger.StandardLogger(logging.Debug).Println(logMSG...)
	} else {
		logger.Println(logMSG...)
	}
}

// Error log based on Severty level
func Error(ctx context.Context, logMSG ...interface{}) {
	if production == env || uat == env {
		client := getClient(ctx)
		defer client.Close()

		logger := client.Logger(LogName)
		logger.StandardLogger(logging.Error).Println(logMSG...)
	} else {
		logger.Println(logMSG...)
	}
}

// Info log based on Severty level
func Info(ctx context.Context, logMSG ...interface{}) {
	if production == env || uat == env {
		client := getClient(ctx)
		defer client.Close()

		logger := client.Logger(LogName)
		logger.StandardLogger(logging.Info).Println(logMSG...)
	} else {
		logger.Println(logMSG...)
	}
}

//Warn ...
func Warn(ctx context.Context, logMSG ...interface{}) {
	if production == env || uat == env {
		client := getClient(ctx)
		defer client.Close()

		logger := client.Logger(LogName)
		logger.StandardLogger(logging.Warning).Println(logMSG...)
	} else {
		logger.Println(logMSG...)
	}

}
