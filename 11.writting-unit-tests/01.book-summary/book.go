package book

type Book struct {
  Title string
  Author string
}

func (b Book) Summary() string {
  return b.Title + " by " + b.Author
}
