
package spec

import (
	"testing"
)

func AssertSame(t *testing.T, a interface{}, b interface{}) {
  if a != b {
    t.Errorf("`%v` and `%v` must be the same.",a,b)
  }
}

func AssertErrorNil(t *testing.T, err error) {
  if err != nil {
    t.Error(err)
  }
}


// func assertSame(t *testing.T, a any, b any) {
//   if a != b {
//     t.Errorf("%s and %s must be the same.",a,b)
//   }
// }

// func assertDifferent(t *testing.T, a any, b any) {
//   if a == b {
//     t.Errorf("%s and %s can't be the same.",a,b)
//   }
// }




