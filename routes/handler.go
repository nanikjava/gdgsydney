package routes

import (
	"database/sql"
	"gdgsydney/db"
	"gdgsydney/static"
	"net/http"
)

func AuthenticateUser(d *sql.DB, username string, password string) error {
	var Users struct {
		Username string
	}

	row := d.QueryRow("SELECT u.username FROM users u  WHERE u.username = ? AND u.password = ?",
		username, password)

	err := row.Scan(&Users.Username)
	if err != nil {
		return err
	}

	return nil
}

func SetupRoutes(db *db.Database) {
	apis := API{DB: db}

	http.HandleFunc("/", serveLoginPage)
	http.HandleFunc("/success", serveSuccessPage)
	http.HandleFunc("/api/login", apis.LoginHandler)
}

func serveLoginPage(w http.ResponseWriter, r *http.Request) {
	serveStaticFile(w, "login.html")
}

func serveSuccessPage(w http.ResponseWriter, r *http.Request) {
	serveStaticFile(w, "success.html")
}

// Helper to serve files from the embedded staticFS
func serveStaticFile(w http.ResponseWriter, filename string) {
	w.Header().Set("Content-Type", "text/html")
	data, err := static.StaticFS.ReadFile(filename)
	if err != nil {
		http.Error(w, "Page not found", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
