package mycrypt

import (
    "testing"
    "reflect"
)

var ALF_OPPG8 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå, ")

// Testen feiler, siden æ ligger tre posisjoner til høyre fra x
// Testen skal passere, hvis man tester for w istedenfor x
func TestKrypter(t *testing.T) {
    wanted := []rune("æ")
    state := Krypter([]rune("x"), ALF_OPPG8, 4)
    if !reflect.DeepEqual(state, wanted) {
         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
    }
}

// Tester posisjon for rune 'æ' i ALF_OPPG8 
// (begynner på 0 teller fra venstre til høyre)
func TestSokIAlfabetet(t *testing.T) {
    wanted := 26
    got := sokIAlfabetet('æ', ALF_OPPG8)
    if got != wanted {
        t.Errorf("Feil, fikk %d, ønsket %d.", got, wanted)
    }
}
