package main

import (
	"ais.com/m/common"
	"ais.com/m/database"
	m "ais.com/m/gun/metrics"
	"ais.com/m/model"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type MetricOutput struct {
	Guns    []model.Gun `json:"guns"`
	Metrics [][]float64 `json:"metrics"`
}

func EuclidHandler(w http.ResponseWriter, r *http.Request) {
	pg := database.GetDB()

	var metricsRes MetricOutput
	pg.Find(&metricsRes.Guns)

	metricsRes.Metrics = make([][]float64, len(metricsRes.Guns))
	weights := []float32{1, 10, 5, 0.05, 15, 15, 2, 10, 15, 1, 15, 5, 0.005, 0.05, 1}

	for i, gun1 := range metricsRes.Guns {
		for j := i; j < len(metricsRes.Guns); j += 1 {
			metric := math.Round(m.Euclidean(gun1, metricsRes.Guns[j], weights))
			metricsRes.Metrics[i] = append(metricsRes.Metrics[i], metric)
			if i != j {
				metricsRes.Metrics[j] = append(metricsRes.Metrics[j], metric)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metricsRes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func L1Handler(w http.ResponseWriter, r *http.Request) {
	pg := database.GetDB()

	var metricsRes MetricOutput
	pg.Find(&metricsRes.Guns)

	metricsRes.Metrics = make([][]float64, len(metricsRes.Guns))
	weights := []float32{1, 10, 5, 0.05, 15, 15, 2, 10, 15, 1, 15, 5, 0.005, 0.05, 1}

	for i, gun1 := range metricsRes.Guns {
		for j := i; j < len(metricsRes.Guns); j += 1 {
			metric := math.Round(m.L1Distance(gun1, metricsRes.Guns[j], weights) * 100) / 100
			metricsRes.Metrics[i] = append(metricsRes.Metrics[i], metric)
			if i != j {
				metricsRes.Metrics[j] = append(metricsRes.Metrics[j], metric)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metricsRes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func PearsonHandler(w http.ResponseWriter, r *http.Request) {
	pg := database.GetDB()

	var metrics MetricOutput
	pg.Find(&metrics.Guns)

	metrics.Metrics = make([][]float64, len(metrics.Guns))
	weights := []float32{1, 10, 5, 0.05, 15, 15, 2, 10, 15, 1, 15, 5, 0.005, 0.05, 1}

	for i, gun1 := range metrics.Guns {
		for j := i; j < len(metrics.Guns); j += 1 {
			metric := math.Round(m.PearsonCorrelation(gun1, metrics.Guns[j], weights) * 100) / 100
			metrics.Metrics[i] = append(metrics.Metrics[i], metric)
			if i != j {
				metrics.Metrics[j] = append(metrics.Metrics[j], metric)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func TreeHandler(w http.ResponseWriter, r *http.Request) {
	pg := database.GetDB()

	var metrics MetricOutput
	pg.Find(&metrics.Guns)

	metrics.Metrics = make([][]float64, len(metrics.Guns))

	for i, gun1 := range metrics.Guns {
		for j := i; j < len(metrics.Guns); j += 1 {
			metric := math.Round(m.TreeMetric(gun1, metrics.Guns[j]) * 100) / 100
			metrics.Metrics[i] = append(metrics.Metrics[i], metric)
			if i != j {
				metrics.Metrics[j] = append(metrics.Metrics[j], metric)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RegisterHTTPEndpoints(router *mux.Router) {
	router.HandleFunc("/api/v1/metric/euclid", EuclidHandler).
		Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/api/v1/metric/city", L1Handler).
		Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/api/v1/metric/pearson", PearsonHandler).
		Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/api/v1/metric/tree", TreeHandler).
		Methods(http.MethodOptions, http.MethodGet)
}

func run(port string) error {
	router := mux.NewRouter()

	RegisterHTTPEndpoints(router)

	router.Use(common.CORSMiddlware)
	router.Use(mux.CORSMethodMiddleware(router))
	httpServer := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return httpServer.Shutdown(ctx)
}

func main() {
	port := "8000"

	if err := run(":" + port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
