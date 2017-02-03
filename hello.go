package hello

import "fmt"

func hello(name string) string {
  return fmt.Sprintf("Hello %q from Go", name)
}

