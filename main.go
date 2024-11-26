package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /run", handleRun)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}

type runBodyT struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func handleRun(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	runBody := new(runBodyT)
	err = json.Unmarshal(body, runBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd := exec.Command(runBody.Command, runBody.Args...)
	output, err := cmd.Output()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
