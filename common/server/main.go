package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"gocache/common/log"
	"gocache/common/response"
	"gocache/common/router"
	"gocache/config"
	"gocache/utils"
	"gocache/utils/errors"
)

func ListenAndServe(listenAddr string, router router.Router) {
	env := config.GetEnvironment()

	nextRequestID := func() string {
		return utils.GetMD5Hash(strconv.FormatInt(time.Now().UnixNano(), 10))
	}

	mux := http.NewServeMux()
	for _, list := range *(router.GetMethods()) {
		for prefix, handlerFunc := range list {
			mux.HandleFunc(prefix, handlerFunc)
		}
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(utils.ContentType, utils.ApplicationJSON)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(response.JSONResponse(nil, fmt.Errorf(errors.NotFound)))
		return
	})

	server := http.Server{
		Addr:         listenAddr,
		Handler:      tracing(nextRequestID)(logging(mux)),
		ErrorLog:     log.GetApiLog(),
		ReadTimeout:  env.Timeout,
		WriteTimeout: env.Timeout,
		IdleTimeout:  env.Timeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(fmt.Errorf(errors.ServerNotListen, listenAddr, err))
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (kill -2)
	<-stop
	log.Http(fmt.Errorf(errors.ServerShutDown))

	ctx, cancel := context.WithTimeout(context.Background(), env.Timeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Errorf(errors.ServerNotShutdown, listenAddr, err))
	}
	// Wait for ListenAndServe goroutine to close.
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			requestID, ok := r.Context().Value(utils.RequestID).(string)
			if !ok {
				requestID = utils.Unknown
			}

			byt, _ := io.ReadAll(r.Body)

			log.Http(fmt.Errorf(utils.HttpLogFormat,
				requestID, r.Method, r.RequestURI, string(byt), r.RemoteAddr, r.UserAgent()))
		}()
		next.ServeHTTP(w, r)
	})
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(utils.RequestID)
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), utils.RequestID, requestID)
			w.Header().Set(utils.RequestID, requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
