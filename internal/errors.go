package internal

import "fmt"

type ErrSnapshotCreated struct {
	Name     string
	Contents string
}

type ErrSnapshotUpdated struct {
	Name string
	Diff string
}

type ErrSnapshotMismatch struct {
	Diff string
}

type ErrNoSnapshot struct {
	Name string
}

func (e ErrSnapshotCreated) Error() string {
	return fmt.Sprintf("snapshot created for test %s, contents:\n%s", e.Name, e.Contents)
}

func (e ErrSnapshotUpdated) Error() string {
	return fmt.Sprintf("snapshot %s updated:\n%s", e.Name, e.Diff)
}

func (e ErrSnapshotMismatch) Error() string {
	return fmt.Sprintf("snapshot not equal:\n%s", e.Diff)
}

func (e ErrNoSnapshot) Error() string {
	return fmt.Sprintf("snapshot %s does not exist", e.Name)
}
