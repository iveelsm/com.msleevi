type Blog struct {
  date string,
  post string,
  author string,
  comments []Comment,
  replies []Comment,
}

type Comment struct {
  name string,
  date string,
  comment string,
}
