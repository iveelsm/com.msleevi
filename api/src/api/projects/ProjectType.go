package projects

import (
  "time"
)

type Project struct {
  name          string
  date          time.Time
  url           string
  image         string
  synopsis      string
  description   string
}
