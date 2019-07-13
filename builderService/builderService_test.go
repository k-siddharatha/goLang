package builderService

import (
  "testing"
  "fmt"
)

func TestBuilderService(t *testing.T) {
  bd := NewBuildDirector()
  hpbs := NewHomePageBuilderService()
  bd.SetBuilderService(hpbs)
  pg := bd.Construct()
  fmt.Printf("%v",pg)
}
