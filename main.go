package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func applyTc(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	time := ps.ByName("time")
	args := strings.Split(fmt.Sprintf("qdisc add dev eth0 root netem delay %s", time), " ")
	cmd := exec.Command("tc", args...)
	log.Printf("tc command: %s", cmd.String())
	w.Write([]byte(fmt.Sprintf("{\"tc command\":\"%s\"}\n", cmd.String())))

	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("error: %v, tc command: %s, output: %s", err, cmd.String(), string(out))
		return
	}
}

func main() {
	router := httprouter.New()
	router.GET("/latency/:time", applyTc)

	http.ListenAndServe(":2332", router)
}
