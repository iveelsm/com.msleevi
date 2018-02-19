type User struct {
  name string,
  email string,
  isAdmin bool,
  password Password,
}

type Password struct {
  encryption string,
  vector string,
  password string,
}
