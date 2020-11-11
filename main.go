package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	flags := ParseFlags()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] remote address: %s host: %s - %s %s", flags.Name, r.RemoteAddr, r.Host, r.Proto, r.Method)
		fmt.Fprintf(w, "%d - OK", http.StatusOK)
	})
	log.Fatalf("listen and serve: %v", http.ListenAndServe(":8080", nil))
}

type Flags struct {
	Name string
}

func ParseFlags() Flags {

	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	name := f.String("name", getStringEnv("HAT_NAME", "haproxy-test"),
		"unique instance name to be printed in the logs")
	f.Parse(os.Args[1:])

	flags := Flags{
		Name: stringValue(name),
	}
	return flags
}

func getStringEnv(envName string, defaultValue string) string {

	env, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}
	return env
}

func stringValue(v *string) string {

	if v == nil {
		return ""
	}
	return *v
}
