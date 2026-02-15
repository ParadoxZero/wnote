package store

type NoteStore interface {
	SaveNote(n Note) error
	GetNotesByType(nType NoteType) ([]Note, error)
	GetNotesByRepo(repo string) ([]Note, error)
	Close() error
}

