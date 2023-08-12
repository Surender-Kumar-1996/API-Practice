package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"sync"
)

var (
	listUserRe   = regexp.MustCompile(`^\/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/users\/(\d+)$`) //user/123
	createUserRe = regexp.MustCompile(`^\/users[\/]*$`)
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type dataStore struct {
	m map[string]user
	*sync.RWMutex
}

type userHandler struct {
	store dataStore
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listUserRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createUserRe.MatchString(r.URL.Path):
		h.Post(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
	users := make([]user, 0, len(h.store.m))
	h.store.RLock()
	for _, u := range h.store.m {
		users = append(users, u)
	}
	h.store.Unlock()
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println("Error occured while encoding users json & eror: ", err)
		internalServerErrpr(w, r)
		return
	}
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getUserRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	user, ok := h.store.m[matches[1]]
	if !ok {
		notFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		internalServerErrpr(w, r)
		return
	}
}
func (h *userHandler) Post(w http.ResponseWriter, r *http.Request) {
	u := user{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		badRequest(w, r)
		return
	}
	h.store.Lock()
	h.store.m[u.ID] = u
	h.store.Unlock()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		internalServerErrpr(w, r)
		return
	}

}

//Enpoints

//GET /users
//GET /users/{id}
//POST /users

//JSON based

func main() {
	mux := http.NewServeMux()
	userH := &userHandler{
		store: dataStore{
			m:       map[string]user{"1": user{ID: "1", Name: "bob"}},
			RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/users/", userH)
	mux.Handle("/users", userH)
	log.Println("Starting server at localhost:8088...")
	log.Fatal(http.ListenAndServe("localhost:8088", mux))
	log.Println("Server running at :8088")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"Not Found"}`))
}
func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error":"Bad Request"}`))
}

func internalServerErrpr(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error":"Internal server error"}`))
}
