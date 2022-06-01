package mycrypt

func Krypter(melding []rune, chiffer []rune, caesarForsyvning int) []rune {
    kryptertMelding := make([]rune, len(melding))
    for i := 0; i < len(melding); i++ {
        indeks := sokIAlfabetet(melding[i], chiffer)
        if indeks + caesarForsyvning >= len(chiffer) {
            kryptertMelding[i] = chiffer[indeks+caesarForsyvning-len(chiffer)]
        } else {
            kryptertMelding[i] = chiffer[indeks+caesarForsyvning]
        }
    }
    return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
    for i := 0; i < len(alfabet); i++ {
        if symbol == alfabet[i] {
            return i 
            break
        }
    }
    return -1
}
