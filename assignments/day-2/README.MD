## H2

http://terokarvinen.com/2020/go-programming-course-2020-w22/#h2

#### Tehtävänanto:

1. [Käännä ohjelma kolmelle alustalle: Windows, Linux, Mac. Tee staattinen käännös niin, että ohjelma toimii ilman mitään riippuvuuksia tai kirjastoja. Testaa ohjelman toiminta ainakin joillain näistä alustoista ja ota ruutukaappaukset.](#tehtava1)

2. [Kirjastoja kohti. Kirjoita ohjelma, joka käyttää kahta uutta ominaisuutta tai kirjastoa Go by Example-kirjasta. Voit valita mitä vain, helppoja vaihtoehtoja ovat esimerkiksi satunnaisluvut, merkkijonojen käsittely ja aika.](#tehtava2)

3. [Vapaaehtoinen: Patterit mukana. Kirjoita ohjelma, joka käyttää jotain uutta ominaisuutta Go:n standardikirjastoista.](#tehtava3)

---

#### <a id="tehtava1">1. Cross-compilation</a>

Käänsin yksinkertaisen Go-sovelluksen 32- ja 64-bittisenä Macille, Windowsille ja Linuxille. Kone on MacBook Air 2015, käyttöjärjestelmä MacOS Mojave 10.14.6.

[Sovellus itse:](error.go)\
![screenshot-1](/assignments/day-2/screenshots/source.png)

Kääntäminen:\
![screenshot-2](/assignments/day-2/screenshots/crosscompiling.png)

Ajaminen Macilla:\
![screenshot-3](/assignments/day-2/screenshots/execute-darwin.png)

Ajaminen Linuxilla:\
![screenshot-4](/assignments/day-2/screenshots/execute-linux.png)

Oho, kaikki sujui ongelmitta. :o\
Tarkistetaan vielä, että sovellus toimii tarkoituksenmukaisesti, kun merkkijonossa ei ole neljää kirjainta.

[Lähdekoodi](error.go):\
![screenshot-5](/assignments/day-2/screenshots/source-error.png)

Tulos:\
![screenshot-6](/assignments/day-2/screenshots/execute-darwin-error.png)

---

#### <a id="tehtava2">2. Go By Example</a>

Kokeilin [variadisia funktioita](https://gobyexample.com/variadic-functions), [gorutiineja](https://gobyexample.com/goroutines) ja lambdalausekkeita.\
Variadinen lambdagorutiini? Anonyymi variadinen gorutiini? Mene ja tiedä.

[Lähdekoodi:](variadic-routines.go)\
![screenshot-7](/assignments/day-2/screenshots/variadic-routines-source.png)

Jos merkkijonot tulostettaisiin järjestyksessä, näkisimme siis "You are not seeing these strings in the proper order."

Tulos:\
![screenshot-8](/assignments/day-2/screenshots/variadic-routines.png)

---

#### <a id="tehtava3">3. Standardikirjasto</a>

Kokeilin [sync-kirjastosta](https://golang.org/pkg/sync/) [WaitGrouppia](https://golang.org/pkg/sync/#WaitGroup) soveltamalla sitä edelliseen tehtävään.

[Lähdekoodi:](waitgroup.go)\
![screenshot-9](/assignments/day-2/screenshots/waitgroup-source.png)

Jos merkkijonot tulostettaisiin järjestyksessä, näkisimme tällä kertaa "You are not seeing these strings in the proper order either."

Tulos:\
![screenshot-10](/assignments/day-2/screenshots/waitgroup.png)
