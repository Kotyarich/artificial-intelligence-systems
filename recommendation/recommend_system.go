package main

import (
	"ais.com/m/common"
	"ais.com/m/database"
	"ais.com/m/gun/recomendation"
	"ais.com/m/model"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func AllGunsHandler(w http.ResponseWriter, r *http.Request) {
	pg := database.GetDB()

	var guns []model.Gun
	pg.Find(&guns)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(guns); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SimilarToOneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["gun"][0]
	var gun model.Gun
	database.GetDB().First(&gun, id)

	comparator := recomendation.InitComparator(nil)
	closest := comparator.ClosestToOne(&gun, 5)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(closest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SimilarToSliceHandler(w http.ResponseWriter, r *http.Request) {
	ids := make([]int, len(r.URL.Query()["guns"]))
	for i, id := range r.URL.Query()["guns"] {
		ids[i], _ = strconv.Atoi(id)
	}
	var gun []*model.Gun
	database.GetDB().Find(&gun, ids)

	comparator := recomendation.InitComparator(nil)
	closest := comparator.ClosestToSlice(gun, 5)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(closest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SimilarHandler(w http.ResponseWriter, r *http.Request) {
	likesIds := make([]int, len(r.URL.Query()["likes"]))
	for i, id := range r.URL.Query()["likes"] {
		likesIds[i], _ = strconv.Atoi(id)
	}

	dislikesIds := make([]int, len(r.URL.Query()["dislike"]))
	for i, id := range r.URL.Query()["dislike"] {
		dislikesIds[i], _ = strconv.Atoi(id)
	}

	var likes []*model.Gun
	var dislikes []*model.Gun
	database.GetDB().Find(&likes, likesIds)
	database.GetDB().Find(&dislikes, dislikesIds)

	comparator := recomendation.InitComparator(nil)
	closest := comparator.Closest(likes, dislikes, 5)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(closest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RegisterHTTPEndpoints(router *mux.Router) {
	router.HandleFunc("/api/v1/guns", AllGunsHandler).
		Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/api/v1/closest/one", SimilarToOneHandler).
		Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/api/v1/closest/several", SimilarToSliceHandler).
		Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/api/v1/closest", SimilarHandler).
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
