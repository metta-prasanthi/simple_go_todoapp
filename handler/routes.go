package handler

import (
    "net/http"

    "../repositories"
)

func SetUpRouting() *http.ServeMux {
    todoHandler := &todoHandler{
        samples:  &repositories.Sample{},
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/samples", todoHandler.GetSamples)

    return mux
}