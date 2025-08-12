// Package handlers contains HTTP handlers for the web app.
package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"web_app/internal/models"
)

// HealthHandler returns a simple OK response.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// ItemHandler provides item-related endpoints and holds a DB connection.
type ItemHandler struct {
	db *sql.DB
}

// NewItemHandler constructs an ItemHandler.
func NewItemHandler(db *sql.DB) *ItemHandler { return &ItemHandler{db: db} }

// ListItems returns all items ordered by id asc.
func (h *ItemHandler) ListItems(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query(`SELECT id, name, created_at FROM items ORDER BY id ASC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	items := make([]models.Item, 0)
	for rows.Next() {
		var it models.Item
		var created string
		if err := rows.Scan(&it.ID, &it.Name, &created); err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}
		// SQLite returns text timestamp; parse to time.Time for consistency
		if t, err := time.Parse(time.RFC3339Nano, created); err == nil {
			it.CreatedAt = t
		}
		items = append(items, it)
	}
	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, items)
}

// CreateItem creates a new item from JSON body {"name":"..."}
func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusBadRequest, errors.New("name is required"))
		return
	}

	res, err := h.db.Exec(`INSERT INTO items(name) VALUES (?)`, req.Name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	id, _ := res.LastInsertId()

	// Fetch the created item including timestamp
	var it models.Item
	var created string
	if err := h.db.QueryRow(`SELECT id, name, created_at FROM items WHERE id = ?`, id).
		Scan(&it.ID, &it.Name, &created); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	if t, err := time.Parse(time.RFC3339Nano, created); err == nil {
		it.CreatedAt = t
	}

	writeJSON(w, http.StatusCreated, it)
}

// Helpers
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{
		"error": err.Error(),
	})
}
