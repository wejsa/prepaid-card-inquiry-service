package pkg

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"net/http/httputil"
	"os"
	"strings"
	"sync"
	"time"
)

var ResponseCompressMinSize = 1024
var ResponseCompressEnable = false

func init() {

	hostname, _ := os.Hostname()

	logFileName := "pis-" + hostname + ".log"

	var fileLogger = &lumberjack.Logger{
		Filename:   "./" + logFileName,
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 100,
		LocalTime:  true,
	}
	log.SetOutput(fileLogger)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	SetLogLevel("INFO")
}

func ReadEnvResponseCompress(isEnable bool, minSize int) {
	if minSize > 0 {
		ResponseCompressMinSize = minSize
	}
	ResponseCompressEnable = isEnable
}

var logLevel string

func LogLevel() (level string) {
	return logLevel
}

func SetLogLevel(level string) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	logLevel = level
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// Logger sets the logging format
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		corrId := c.Request.Header.Get("X-KM-Correlation-Id")
		userId := c.Request.Header.Get("X-Km-Userid")
		caller := c.Request.Header.Get("X-KM-CALLER")

		if caller == "" {
			caller = "CLIENT"
		}

		// debug
		dumpRequest(c)

		if c.Request.Method == "POST" {
			if body, err := io.ReadAll(c.Request.Body); err == nil {
				log.Printf("INFO T[%s] U[%s] - [%s-REQ] %s %s       %s", corrId, userId, caller, c.Request.Method, c.Request.URL.Path, string(body))

				//set a new body, which will simulate the same data we read:
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			} else {
				log.Printf("INFO T[%s] U[%s] - [%s-REQ] %s %s", corrId, userId, caller, c.Request.Method, c.Request.URL.Path)
			}

		} else {
			if c.Request.URL.RawQuery == "" {
				log.Printf("INFO T[%s] U[%s] - [%s-REQ] %s %s", corrId, userId, caller, c.Request.Method, c.Request.URL.Path)
			} else {
				log.Printf("INFO T[%s] U[%s] - [%s-REQ] %s %s?%s", corrId, userId, caller, c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery)
			}
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// before request
		c.Next()

		// after request
		elapsed := time.Since(start)

		// access the status we are sending
		status := c.Writer.Status()

		//w.Header().Del("X-KM-TEMP")
		c.Writer.Header().Del("X-KM-TEMP")

		delayed := ""

		if elapsed.Seconds() > 0.1 {
			switch {
			case elapsed.Seconds() > 1:
				delayed = "_DELAYED10_"
			case elapsed.Seconds() > 0.5:
				delayed = "_DELAYED05_"
			case elapsed.Seconds() > 0.1:
				delayed = "_DELAYED01_"
			}
		}

		if delayed != "" {
			log.Printf("WARN T[%s] U[%s] - [%s-RES] %s %s %d %s (%s)     %s", corrId, userId, caller, c.Request.Method, c.Request.URL.Path, status, delayed, elapsed, blw.body.String())
		} else {
			log.Printf("INFO T[%s] U[%s] - [%s-RES] %s %s %d (%s)     %s", corrId, userId, caller, c.Request.Method, c.Request.URL.Path, status, elapsed, blw.body.String())
		}
	}
}

// dumpRequest is logging request header info
func dumpRequest(c *gin.Context) {
	corrId := c.Request.Header.Get("X-KM-Correlation-Id")
	userId := c.Request.Header.Get("X-Km-Userid")

	requestDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Println("Dump Error : ", err.Error())
	}
	for _, line := range strings.Split(strings.TrimRight(string(requestDump), "\n"), "\n") {
		log.Printf("INFO T[%s] U[%s] http-outgoing >> %s", corrId, userId, line)
	}
}
