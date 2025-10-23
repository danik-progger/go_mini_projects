package proxy

import (
	"context"
	"load_balancer/balancer"
	server "load_balancer/cmd"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func SetUpProxy(serverUrl *url.URL, serverPool server.ServerPool) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(serverUrl)
	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		log.Printf("[%s] %s\n", serverUrl.Host, e.Error())
		retries := balancer.GetRetryFromContext(request)
		if retries < 3 {
			<-time.After(10 * time.Millisecond)
			ctx := context.WithValue(request.Context(), balancer.Retry, retries+1)
			proxy.ServeHTTP(writer, request.WithContext(ctx))
			return
		}

		// after 3 retries, mark this backend as down
		serverPool.MarkBackendStatus(serverUrl, false)

		// if the same request routing for few attempts with different backends, increase the count
		attempts := balancer.GetAttemptsFromContext(request)
		log.Printf("%s(%s) Attempting retry %d\n", request.RemoteAddr, request.URL.Path, attempts)
		ctx := context.WithValue(request.Context(), balancer.Attempts, attempts+1)
		balancer.LB(writer, request.WithContext(ctx), &serverPool)
	}

	return proxy
}
