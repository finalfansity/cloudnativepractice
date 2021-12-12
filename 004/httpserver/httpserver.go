package httpserver

import (
	"003/logger"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var FunctionLatency *prometheus.HistogramVec

func HttpServer() {
	serve := http.NewServeMux()
	serve.HandleFunc("/", Index)
	serve.HandleFunc("/healthz", Healthz)
	serve.Handle("/metrics", promhttp.Handler())

	FunctionLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "httpserver",
			Name:      "rand_time_spend",
			Help:      "Time spend",
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		}, []string{"step"},
	)
	err := prometheus.Register(FunctionLatency)
	if err != nil {
		fmt.Println(err)
	}

	serverun := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: serve,
	}
	go func() {
		if err := serverun.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to start server", err)
		}
	}()

	// grace shutdown
	quit := make(chan os.Signal)
	// receive system signal
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // block
	// service will be shut down in 5 seconds, wait for the request to be processed
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serverun.Shutdown(ctx); err != nil {
		logger.Error("shutdown server failed", err)
	}
	logger.Println("server shutdown successfully")
}

func Index(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	statusCode := 200
	now := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())
	delay := rand.Intn(2000)
	fmt.Println(delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	defer FunctionLatency.WithLabelValues("rand time sleep").Observe(time.Now().Sub(now).Seconds())
	w.WriteHeader(statusCode)

	logger.Println("host: " + r.RemoteAddr + " method: " + r.Method + " code: " + strconv.Itoa(statusCode))
	s := viper.GetString("db.host")
	l := viper.GetString("info")
	w.Write([]byte("server info " + l + " db host: " + s))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	statusCode := "200"
	logger.Debug("healthz check ok")
	w.Write([]byte(statusCode))
}
