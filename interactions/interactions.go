package interactions

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	ia "github.com/bsdlp/discord-interactions-go/interactions"
	"github.com/gorilla/mux"
)

type Service struct {
	PubKey []byte
	r      *mux.Router
}

func New(pubKey string) *Service {
	discordPubkey, err := hex.DecodeString(pubKey)
	if err != nil {
		// handle error
	}
	r := mux.NewRouter()

	s := Service{
		discordPubkey,
		r,
	}

	r.HandleFunc("/", s.InteractionHandler).Methods("POST")

	return &s
}

func (s *Service) Start() error {
	return http.ListenAndServe(":8080", s.r)
}

func (s *Service) InteractionHandler(w http.ResponseWriter, r *http.Request) {
	verified := ia.Verify(r, ed25519.PublicKey(s.PubKey))
	if !verified {
		http.Error(w, "signature mismatch", http.StatusUnauthorized)
		return
	}
	fmt.Println("yo")

	defer r.Body.Close()
	var interaction Interaction
	err := json.NewDecoder(r.Body).Decode(&interaction)
	if err != nil {
		panic("TODO1")
	}
	fmt.Println("yo2")

	// respond to ping
	if interaction.Type == 1 {
		fmt.Println("yo3")
		_, err := w.Write([]byte(`{"type":1}`))
		if err != nil {
			panic("TODO2")
		}
		return
	}
	fmt.Println("yo4")
}

// EncodeJSONResponse uses the json encoder to write an interface to the http response with an optional status code
func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return json.NewEncoder(w).Encode(i)
}
