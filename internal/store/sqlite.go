// Copyright (C) 2026 Sidhin S Thomas <thomas.sidhin@outlook.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package store

import (
	"database/sql"
	"fmt"
	"log"
	"strings"


	_ "modernc.org/sqlite"
)

type SqliteStore struct {
	db *sql.DB
}

func CreateSqliteStore(dbPath string) (NoteStore, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Basic check to ensure connection is live
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	s := &SqliteStore{db: db}

	// You would call your migration logic here
	if err := s.migrate(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *SqliteStore) migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type INTEGER,
		title TEXT,
		content TEXT,
		status INTEGER,
		repo TEXT,
		branch TEXT,
		commit_hash TEXT,
		target TEXT,
		due_at DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := s.db.Exec(query)
	return err
}

// CreateNote inserts a new note into the database
func (s *SqliteStore) SaveNote(n Note) error {
	query := `
	INSERT INTO notes (type, title, desc, status, repo, branch, commit_hash, target, due_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	db := s.db
	_, err := db.Exec(query,
		n.Type, n.Title, n.Desc, n.Status, n.Repo,
		n.Branch, n.Commit, n.FollowupTarget, n.DueAt,
	)
	return err
}

func (s *SqliteStore) GetNotesByType(nType NoteType) ([]Note, error) {
	query := createFilterQuery(`type = ?`)
	db := s.db
	rows, err := db.Query(query, nType)
	if err != nil {
		return nil, err
	}

	return getNotesFromRows(rows)
}

func (s  *SqliteStore) GetNotesByRepo(repo string) ([]Note, error) {
	query := createFilterQuery(`repo = ?`)
	db := s.db
	rows, err := db.Query(query, repo)
	if err != nil {
		return nil, err
	}

	return getNotesFromRows(rows)
}

func (s *SqliteStore) Close() error {
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("Error closing Database: %w", err) 
	}
	return nil
}

func createFilterQuery(filter_by string) string {
	select_statement := `SELECT id, type, title, desc, status, created_at, repo, branch, commit_hash, target, due_at`
	from := `FROM notes`
	where_statement := `WHERE`
	order_by := `ORDER BY created_at DESC`

	return strings.Join([]string{select_statement, from, where_statement, filter_by, order_by }, " ")
}

func getNotesFromRows(rows *sql.Rows) ([]Note, error) {
	var notes []Note
	defer rows.Close()

	for rows.Next() {
		var n Note
		err := rows.Scan(
			&n.ID, &n.Type, &n.Title, &n.Desc, &n.Status, &n.CreatedAt,
			&n.Repo, &n.Branch, &n.Commit, &n.FollowupTarget, &n.DueAt,
		)
		if err != nil {
			log.Println("Unable to scan a row due to - ", err)
			continue
		}

		notes = append(notes, n)
	}

	if err := rows.Err(); err != nil {
		// Return partial list and the iteration error
		return notes, fmt.Errorf("Connection lost during read: %w", err)
	}

	return notes, nil
}
