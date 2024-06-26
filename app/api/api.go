package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Index(w, r)
}

// Index is the actual handler function
func Index(w http.ResponseWriter, _ *http.Request) {
	logger, _ := zap.NewProduction()
	logger.Info("successfully performed http request")
	logger.Info(quote.Hello())
	logger.Info(quoteV3.HelloV3())
	fmt.Fprintf(w, "Successfully performed HTTP request\n")
}

func RunServer() {
	router := httprouter.New()
	router.GET("/", IndexHandler)
	logger, _ := zap.NewProduction()

	logger.Info("server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
