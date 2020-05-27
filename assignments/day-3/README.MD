 ## H3

http://terokarvinen.com/2020/go-programming-course-2020-w22/#h3

#### Tehtävänanto:

1. [Kirjastosta päivää. Tee ohjelma, joka käyttää ainakin kahta uutta kirjastoa. Jos ohjelmallasi on jokin käyttötarkoitus, se on mainiota. Muista ensin kokeilla kutakin kirjastoa erikseen. Kirjastoja löytyy GoByExample ja Go:n virallisesta dokumentaatiosta. Käytä jotain uutta kirjastoa, eli ei pelkästään rand, fmt, time eikä string.](#tehtava1)

2. [Lähteet. Katso, että olet viitannut kaikissa tehtävävastauksissasi kaikkiin lähteisiin, joita olet käyttänyt. Kurssiin, tehtävänantoihin, käyttämääsi materiaaliin (GoByExample tms), muiden koodeihin, StackOverflown vastauksiin ja kaikkiin muihinkin lähteisiin.](#tehtava2)

3. [Kaikki tai jotain. Valitse aihe tai pari vaihtoehtoa kurssin loppuprojektille. Palastele: mieti, miten voit ensin toteuttaa pienen, toimivan kokonaisuuden. Ja sen päälle vaihtoehtoja laajentaa.](#tehtava3)

4. [Vapaaehtoinen, helppo ja suositeltava: Asenna Go omaan isäntäjärjestelmääsi. Senhän saa Linuxille, Windowsille ja Macille. https://golang.org/ Jos Windowsin virustorjuntaohjelmasi poistelee mielivaltaisesti tiedostoja, kerro, kuinka annoit itsellesi virustorjunnan asetuksista luvan laittaa binäärejä johonkin hakemistoon.](#tehtava4)

5. [Vapaaehtoinen: Grep. Tee ohjelma, joka näyttää tekstitiedostosta vain rivit, joilla mainitaan hakusana. Kun ohjelmasi toimii ja haluat lisähaastetta, voit tehdä lisäominaisuuksia: case insensitive, regular expressions, rekursio. Jos kaipaat vielä lisähaastetta, voit kokeilla, toimiiko ohjelmasi isolla (yli 50 GB) tiedostoilla, esim. tietokantadumpeilla?](#tehtava5)

---

#### <a id="tehtava1">1. Uudet kirjastot</a>

(tulossa)

---

#### <a id="tehtava2">2. Lähteet</a>

[Lähteet lisätty ensimmäisen päivän rot13-tehtävälle](https://github.com/smtnsk/go-course/blob/master/assignments/day-1/README.MD). Flag-tehtävässä en käyttänyt kuin opettajan näyttämiä esimerkkejä.\
Toisen päivän raportissa on jo lähteet linkitetty.

---

#### <a id="tehtava3">3. Idea loppuprojektille</a>

Teen kuvankäsittelyohjelman. Sain tänään toimivaan kuntoon koodinpätkän, jolla voi manipuloida JPEG-formaatissa tallennettujen kuvien värikanavia. Laajennan sitä projektia varten.

Mahdollisia laajennuksia:
* Efektejä (blur, edge-detection...)
* Kuvan suurennus ja pienennys
* Kuvan leikkaus
* Useampi kuvaformaatti (PNG, GIF...)
* Monipuolisempi värien muokkaus (per värikanava + alpha) 
* Muuttaminen yhdestä väriavaruudesta toiseen

En varmasti saa näitä kaikkia tehtyä, mutta koska runkona on jo toimivaa koodia, voin huoletta syventyä ainakin yhden tai kahden lisätoiminnon luomiseen ilman kauheaa kiirettä. Lisäksi voin käyttää efektien tekemiseen mallina aiemmin kirjoittamaani Java- ja C-koodia.

---

#### <a id="tehtava4">4. Go:n asennus</a>

En käyttänyt virtuaalikonetta alunperinkään, vaan Go on asennettuna natiivisti Macille ja Alpine Linuxille:\
![screenshot-](/assignments/day-3/screenshots/environment.png)

---

#### <a id="tehtava5">5. Köyhän miehen grep</a>

(teen jos ehdin, huonolta näyttää)

