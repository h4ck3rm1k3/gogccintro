package tests

import (
	"testing"
)

func TestReadGob(b* testing.T)  {
	DoUnmarshal("scope.gob")
}
