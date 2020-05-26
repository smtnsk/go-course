## H2

http://terokarvinen.com/2020/go-programming-course-2020-w22/#h2

#### Tehtävänanto:

1. [Käännä ohjelma kolmelle alustalle: Windows, Linux, Mac. Tee staattinen käännös niin, että ohjelma toimii ilman mitään riippuvuuksia tai kirjastoja. Testaa ohjelman toiminta ainakin joillain näistä alustoista ja ota ruutukaappaukset.](#tehtava1)

2. [Kirjastoja kohti. Kirjoita ohjelma, joka käyttää kahta uutta ominaisuutta tai kirjastoa Go by Example-kirjasta. Voit valita mitä vain, helppoja vaihtoehtoja ovat esimerkiksi satunnaisluvut, merkkijonojen käsittely ja aika.](#tehtava2)

3. [Vapaaehtoinen: Patterit mukana. Kirjoita ohjelma, joka käyttää jotain uutta ominaisuutta Go:n standardikirjastoista.](#tehtava3)

---

#### <a id="tehtava1">1. Cross-compilation</a>

Käänsin yksinkertaisen Go-sovelluksen 32- ja 64-bittisenä Macille, Windowsille ja Linuxille. Kone on MacBook Air 2015, käyttöjärjestelmä MacOS Mojave 10.14.6.

[Sovellus itse:](error.go)

![screenshot-1](/assignments/day-2/screenshots/source.png)

Kääntäminen:

![screenshot-2](/assignments/day-2/screenshots/crosscompiling.png)

Ajaminen Macilla:

![screenshot-2](/assignments/day-2/screenshots/execute-darwin.png)

Ajaminen Linuxilla:

![screenshot-2](/assignments/day-2/screenshots/execute-linux.png)

Oho, kaikki sujui ongelmitta. :o
Tarkistetaan vielä, että sovellus toimii tarkoituksenmukaisesti, kun merkkijonossa ei ole neljää kirjainta.

[Lähdekoodi](error.go):

![screenshot-1](/assignments/day-2/screenshots/source-error.png)

Tulos:

![screenshot-2](/assignments/day-2/screenshots/execute-darwin-error.png)

---

#### <a id="tehtava2">2. Go By Example</a>

tulossa

---

#### <a id="tehtava3">3. Standardikirjasto</a>

tulossa
