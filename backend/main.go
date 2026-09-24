package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

//go:embed dist
var frontendFS embed.FS

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is missing")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v\n", err)
	}
	fmt.Println("Successfully connected to the database")

	dao := NewRecipeDaoPg(db)

	handler := NewRecipeHandler(dao)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /recipes", handler.GetAll)
	mux.HandleFunc("POST /recipes", handler.Create)
	mux.HandleFunc("POST /recipes/{id}/ingredients/", handler.AddIngredient)
	mux.HandleFunc("DELETE /recipes/{id}", handler.DeleteRecipe)
	mux.HandleFunc("DELETE /ingredients/{id}", handler.DeleteIngredient)

	distFS, err := fs.Sub(frontendFS, "dist")
	if err != nil {
		log.Fatalf("Failed to load frontend static files: %v", err)
	}

	fmt.Println("--- EMBEDDED FILES ---")
	fs.WalkDir(distFS, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println("-", path)
		return nil
	})
	fmt.Println("----------------------")

	mux.Handle("/", http.FileServer(http.FS(distFS)))

	port := ":8080"
	fmt.Printf("Running on http://localhost%s\n", port)

	// 4. Wrap the router in our CORS middleware and start the server
	err = http.ListenAndServe(port, enableCORS(mux))
	if err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
