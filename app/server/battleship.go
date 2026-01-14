package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

var server *Battleship

// Flashcards Structure
type Battleship struct {
	Database  *mongo.Database
	Router    *gin.Engine
	Version   string
	Port      string
	TokenKey  string
	Origin    string
	LogFormat string
	Mode      string
	DBHost    string
}

func (bs *Battleship) ParseParameters() {
	bs.LogFormat = os.Getenv("LOG_FORMAT")
	bs.Version = os.Getenv("API_VERSION")
	bs.Port = os.Getenv("API_PORT")
	bs.TokenKey = os.Getenv("TOKEN_KEY")
	bs.Origin = os.Getenv("ALLOW_ORIGIN")
	bs.Mode = os.Getenv("MODE")
	bs.DBHost = os.Getenv("DB_HOST")
}

// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
// https://github.com/gin-gonic/gin
func (bs *Battleship) ListenAndServe() error {
	srv := &http.Server{
		Addr:              bs.Port,
		Handler:           bs.Router,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// start
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("Unable to listen and serve: %v", err)
		return err
	}
	return nil
}

// SetServer init mongo database
func SetServer(s *Battleship) {
	server = s
}

// GetServer Flashcards
func GetServer() *Battleship {
	return server
}
