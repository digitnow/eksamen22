# Løsningsforslag for Oppgave 8.2 eksamen IS-105
## Mappe- og filstruktur med beskrivelser
* mycrypt - mappe som inneholder pakke for kryptering og dekryptering (Cæser-chiffer i dette tilfelle)
* udp_client.go - fil hvor UDP "KLIENT" er implementert og hvor kryptering blir gjort (hovedsakelig basert på Woodbeck)
  * definerer adressen (ip + port) til "SERVER" (som blir "proxy", dvs. det mellomledd i dette tilfelle, som sender data fra "KLIENT" videre til "KLIENT2")
  * skaper en prosess som "hører" på en adresse (ip + port) på "KLIENT" 
  * definerer alfabetet for språket brukt og krypterer en eksempelmelding (hardkodet i dette tilfelle)
  * client.WriteTo sender den krypterte meldingen til "SERVER"
  * UNØDVENDIG i forhold til Oppgave 8.2: mottar samme melding tilbake fra "SERVER"
  * UNØDVENDIG i forhold til Oppgave 8.2: sjekker om "echo" kommer fra samme node som meldingen ble sendt til
  * UNØDVENDIG i forhold til Oppgave 8.2: sjekker om "echo" inneholder samme antall bytes som meldingen som ble sendt til "SERVER"
* udp_server.go - fil hvor UDP "SERVER" er implementert (blir "proxy" i dette tilfelle)
  * definerer adressen (ip + port) til "KLIENT2" (som skal dekryptere meldingen fra "KLIENT")
  * skaper en prosess som "hører" på en adresse (ip + port) på "SERVER"
  * allokerer plass for data
  * skaper en "evig" løkke, hvor man hører på datastrøm som kommer inn og sender tilbake til noden som datastrøm ble sendt fra ("echo"), samt sender data som kom inn (kan være hvilken som helst node, så her kunne man lagt inn en sjekk for godkjente noder, for eksempel, eller be nodene om å autentisere seg)
* udp_client2.go - fil hvor UDP "KLIENT2" er implementert (der hvor dekryptering blir gjort)
  * definerer adressen (ip + port) til "SERVER" (kilde for data fra "KLIENT2" sitt ståsted); trenger ikke dette for å teste konseptet, men kan være nyttig hvis man trenger å "bevise" hvor data kommer fra, eller eventuelt å implementere autentisering/autorisering
  * skaper en prosess som "hører" på en adresse (ip + port) på "KLIENT2"
  * allokerer plass for data
  * venter på at data skal komme inn 
  * når data kommer, dekrypterer data (må ha kjennskap til nøkkel for å behandle chiffertekst) og skriver de ut til "stdout" (standard ut-strømmen); siden nøkkel blir generert sirkulært kan man bruke samme funksjon Krypter (burde da endret navn?) for å dekryptere (bare gjør en hel rotasjon minus 4)

## Refleksjon
Refaktorering av kode er mulig. 

Man kunne, for eksempel, lagt ut ALF_OPPG8 definisjon til mycrypt.go, og importert den i "KLIENT", "SERVER" og "KLIENT2", slik at man kan endre alfabetet ett sted(kan legge til alle tall 0-9 og flere spesielle tegn).

Man kunne også latt "KLIENT" og "KLIENT2" gå i "evig" løkke og implementere en input-mekanisme på "KLIENT". 

## Testing
* Starte udp_server.go først i et kommandovindu. 
* Starte udp_client2.go i et nytt kommandovindu.
* Starte udp_client.go i et tredje kommandovindu. 

Et mulig resultat: 
* kommandovindu #3
```
$ go run udp_client.go 
Kryptert melding: æcdødskdådqaxiwdmdbpiwyrh
```
