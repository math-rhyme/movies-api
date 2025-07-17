package main
import (
 	"flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

// later will be generated automatically at build time
const version = "1.0.0"

// configuration settings (will be read from CL flags)
// for now: port to listen to and operating env for app (develop, staging etc)
type config struct {
	port int
	env string
}

// app struct to hold dependencies for HTTP handlers, helpers, middlware
type application struct {
	config config
	logger *log.Logger
}


func main() {
	
	//declaring an instance of the config struct
	var cfg config

	// reading CL flags into config
	flag.IntVar(&cfg.port, "port", 4000, "API sserver port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|produstion)")
	flag.Parse()

	// initialize a new logger
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	// declare an instance of the app struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// declare a new servemux and add a /v1/healthcheck route which dispatches requests
	// 		to the healthcheckHandler method 
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// declare an HTTP server with timeout
	// srv := &http.Server{
	// 	Addr: 			fmt.Sprintf(":%d", cfg.port),
	// 	Handler:		mux,
	// 	IdleTimeout: 	time.Minute,
	// 	ReadTimeout:	10 * time.Second,
	// 	WriteTimeout:	30 * time.Second,
	// }

    srv := &http.Server{
        Addr:         fmt.Sprintf(":%d", cfg.port),
        Handler:      app.routes(),
        IdleTimeout:  time.Minute,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

	// start the HTTP server
	fmt.Println("hello world!")
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}