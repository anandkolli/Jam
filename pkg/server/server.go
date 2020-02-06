package server

import (
        "log"
        "net/http"

        "time"

        "jam/pkg/handlers"
        "jam/pkg/pg"
        "github.com/gorilla/mux"
)

const (
        writeTimeout = 15
        readTimeout  = 15
        baseURL = "/api/v1/activity"
        timeSpentPerPerson = baseURL + "/persons"
	personsSpentAnyTime = timeSpentPerPerson + "/count"
	sessions = baseURL + "/sessions"
	funnel = baseURL + "/funnel"
	vip = baseURL + "/vip"
)

// Server holds the address on which it binds and an instance of DB
type Server struct {
        addr   string
        db     *pg.PgClient
        router *mux.Router
        server *http.Server
}

// NewServer creates http server and stores db instance info and address on which sever binds and listen
func NewServer(addr string, db *pg.PgClient) *Server {
        r := mux.NewRouter()

        return &Server{
                addr:   addr,
                db:     db,
                router: r,
                server: &http.Server{
                        Handler:      r,
                        Addr:         addr,
                        WriteTimeout: writeTimeout * time.Second,
                        ReadTimeout:  readTimeout * time.Second,
                },
        }
}

// setHandlers setup the router to corrsponding function handler
func (s *Server) setHandlers() {
        s.router.HandleFunc(baseURL, func(w http.ResponseWriter, r *http.Request) {
                handlers.GetAllActivity(s.db, w, r)
        }).Methods("GET")

        s.router.HandleFunc(timeSpentPerPerson, func(w http.ResponseWriter, r *http.Request) {
                handlers.GetTimeSpentPerPerson(s.db, w, r)
        }).Methods("GET")

        s.router.HandleFunc(personsSpentAnyTime, func(w http.ResponseWriter, r *http.Request) {
                handlers.GetPersonsSpentAnyTime(s.db, w, r)
        }).Methods("GET")

        s.router.HandleFunc(sessions, func(w http.ResponseWriter, r *http.Request) {
                handlers.GetSessions(s.db, w, r)
        }).Methods("GET")

        s.router.HandleFunc(funnel, func(w http.ResponseWriter, r *http.Request) {
                handlers.GetFunnel(s.db, w, r)
        }).Methods("GET")

        s.router.HandleFunc(vip, func(w http.ResponseWriter, r *http.Request) {
                handlers.GetVip(s.db, w, r)
        }).Methods("GET")
}

// Run sets the router handlers and listen on host:port
func (s *Server) Run() {
        s.setHandlers()
        log.Println("Server is up and running on", s.addr)
        log.Fatal(s.server.ListenAndServe())
}
