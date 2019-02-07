package zgoserver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

const connectionString = "mongodb://localhost:27017"
const dbName = ""

// App is used to construct the server.
type App struct {
	Router *mux.Router
}

// // mongoModel is used to construct the mongo database connection.
type Zmodel struct {
	DbHost string
	DbName string
	DbUser string
	DbPass string
	Do     *mongo.Database
}

// Initialize test
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
}

func (m *Zmodel) Connect() error {

	var err error

	mongouri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		m.DbUser,
		m.DbPass,
		m.DbHost,
		m.DbName,
	)

	db, err := mongo.NewClient(mongouri)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	}

	err = db.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}

	m.Do = db.Database(m.DbName)
	return nil
}

// Run
func (a *App) Run(addr string) {

	allowedHeaders := handlers.AllowedHeaders([]string{"Origin, Content-Type,X-Key, Authorization"})
	// allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3400"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	err := http.ListenAndServe(addr, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(a.Router))
	if err != nil {
		log.Fatal(err)
	}
}
