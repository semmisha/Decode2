package main

import (
	"Decoder2/usescases"
	"testing"
)

func TestConvert(t *testing.T) {

	t.Run("Bolid", func(t *testing.T) {

		Realresult := usescases.Bolid("550027DCCF", nil)
		result := "0F00000027DCCF01"
		if result != Realresult {
			t.Errorf("%v != %v", result, Realresult)
		}
	})
	t.Run("Stork", func(t *testing.T) {

		Realresult := usescases.Stork("550027DCCF", nil)
		result := "01CFDC2700550087"
		if result != Realresult {
			t.Errorf("%v != %v", result, Realresult)
		}
	})

}
