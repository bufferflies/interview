package domain

type Index struct {
	Entries []Entry
}
type Entry struct {
	hash []byte
	offset int64
}
