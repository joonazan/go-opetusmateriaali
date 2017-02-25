# Vikkelä viiva

$$$liike/main.go$$$

En liittänyt kuvaa tästä ohjelmasta, sillä se näyttää eri aikoihin erilaiselta.

Ainoa merkittävä ero aiempaan viivoja piirtävään ohjelmaan on tämä rivi:

```Go
	gl.Vertex2d(aika-1, 0)
```

Siinä luodun pisteen x-koordinaatti riippuu ajasta. Kun aika on nolla, se on `-1`. Ajan kasvaessa x-koordinaattikin kasvaa. Mutta mistä tämä aika tulee?

Aiemmin kirjoitin:

> Runin ottaman funktion on pakko ottaa yksi `float64`-tyyppinen argumentti, eikä se saa palauttaa mitään.

Run-funktio lupaa kutsua sille annettua funktiota ajalla ohjelman käynnistymisestä sekunteina. Piirrä-funktion ensimmäiselle argumentille on annettu nimi `aika` määrittelyn alussa:

```Go
func piirrä(aika float64) {
```

### Sykli

Kokeile vaihtaa viivan alkupisteeksi `(0, 0)` ja loppupisteeksi `(0.5*math.Cos(aika), 0.5*math.Sin(aika))`.

Käytän koodissa siniä(`math.Sin`) aallon pisteiden korkeuden laskemiseen. Sini muuttaa luvun luvuksi välillä [-1, 1] yksitoikkoisen aaltoilevasti. Kosini(`math.Cos`) on muuten samanlainen, mutta se on hieman siniä jäljessä. Tässä hyödynnetään niiden ominaisuuksia ympyrän kaaren laskemiseen, mutta tavallisempaa on, että niitä ripotellaan kaikkialle minne halutaan pehmeitä edestakaisia muutoksia.

### Tehtäviä

 - Tee viivaa monimutkaisempi kuvio, joka venyy.
 - Laita värit muuttumaan.
 - Piirrä monta pistettä ja kokeile vaihtaa `gl.LINES` tilalle `gl.LINE_STRIP` tai `gl.TRIANGLE_FAN`.
 - Tee kuvio, joka liikkuu edestakaisin.
 - Opettele for-silmukka ja piirrä sillä ympyrä.
 - Piirrä täyttyvä latausympyrä.