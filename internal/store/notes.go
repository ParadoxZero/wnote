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
	"time"
)

type NoteType int

const (
	TodoNote NoteType = iota
	CheckpointNote
	IdeaNote
	FollowupNote
)

type NoteStatus int

const (
	Active NoteStatus = iota
	Blocked
	Completed
	Archived
)

func (n NoteType) String() string {
	return [...]string{"todo", "checkpoint", "idea", "followup"}[n]
}

type Note struct {
	ID        int
	Type      NoteType
	Title     string
	Desc      string
	Status    NoteStatus
	CreatedAt time.Time

	// Context for each notes. Can be optional
	// depending on type of note.
	Repo           *string
	Branch         *string
	Commit         *string
	FollowupTarget *string
	DueAt          *time.Time
}

func (n NoteStatus) String() string {
	return [...]string{"active", "blocked", "completed", "archived"}[n]
}
