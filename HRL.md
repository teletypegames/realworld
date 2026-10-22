# HRL a Personal Tape-ekhez: kutatás és javaslat

Állapot: a javaslat A–C lépése és a D lépés egy első, kérésre működő változata megvalósult a `hrl` branch-en (inkwell, realworld, wiki-pages); lásd a 8. fejezetet. Készült: 2026-10-22.

A kérdés: vezérelhetné-e a realworld (inkwell-alapú játék) Personal Tape-jeit
(Dexet és a többieket) egy olyan MI-technika, mint a **hierarchikus
megerősítéses tanulás (HRL)**? Ez a dokumentum rögzíti, mit néztünk meg, mire
jutottunk, és egy konkrét, megvalósítható tervet.

**Röviden:** a *hierarchia* nagyon jól illik a tape-ekhez, a *tanulás* nem. A
javaslat egy háromszintű tape-„agy”, amely az inkwell deklaratív tartalmán
végzett szimbolikus keresésre épül, tanítás nélkül. Az első szintje (hol tart
Paul a puzzle-gráfban?) már ma megvalósítható. Az inkwell ugyanis reflexióval
már most is be tud járni minden scene-t, hotspotot, feltételt és actiont, a
realworld pedig szinte egyáltalán nem használ átláthatatlan kódot.

---

## 1. Háttér

### 1.1 Az inkwell

Az `engines/inkwell` egy point-and-click adventure keretrendszer Go-ban,
Ebitengine fölött. Egy játék egyszerű struct literálokból áll, amelyek
entitásonkénti `*Manager` regiszterekbe kerülnek: scene-ek, hotspotok, exitek,
triggerek, tárgyak, karakterek, dialógusok, scriptek, verbek. A viselkedést ezek
írják le:

- **Actionök** (`Action` / `Runner`): `Seq`, `Par`, `If`, `Say`, `Give`,
  `TakeAway`, `RequireItem`, `SetFlag`, `SetVar`, `GoTo`, `RunDialog`,
  `RunScript` stb., valamint a két kiskapu, a `Custom` és a `Reply` (closure-ök).
- **Feltételek** (`Condition.Eval(ctx) bool`): `Flag`, `HasItem`,
  `SelectedItem`, `InScene`, `Knows`, `VarEq`, `And`/`Or`/`Not`, és a
  `Given{Flag, NotFlag, Item, NotItem, Var, Is}` literál.
- **Világállapot** (`State`): flagek, varok, inventory, `Visited`, `Known`,
  `Talked`.

A closure-öket leszámítva minden adat.

### 1.2 A Personal Tape-ek a realworldben

A `README.md`-ből („The rule everything follows”):

> **The hand holds the world. The voice speaks to the tape.**
> `WORLD`: egér, a négy verb, inventory, vagyis Paul cselekszik.
> `TAPE`: a `talk` verb és a tape-kommentár, vagyis a világban semmi nem történik.

A tape-ek tehát **tanácsadó vagy társ-NPC-k, nem cselekvő NPC-k**. Soha nem
nyúlnak a világhoz. Egyetlen döntésük az, hogy *mit mondanak, mikor, és
mondanak-e egyáltalán valamit*.

A tape-ek nem egyformák:

- **Dex** (`tape/tape.dex.go`, `speaker/speaker.dex.go`, `dialog/dialog.dex_talk.go`)
  élő személyiség az egyes foglalatban: „Száraz, hű, és minden ellen, amit Paul
  épp tenni készül”. A wiki szerint folyamatosan kommentál. Ő a fő jelölt az
  adaptív viselkedésre.
- **Tartalom-tape-ek**: a `tape.commentary.go` szó szerint egy felvétel
  („Recording. If you're hearing this, Paul…”), mellette a mystery, server farm,
  support és sumphor. Ezek megírt tartalmat hordoznak, és többnyire statikusak.
  A mystery, server farm, support és sumphor tape-eket nem néztük meg
  részletesen.

A kánonban kész képességkorlát is van: operátori licenc nélkül a Personal Tape
**restricted mode**-ban fut, és Paulnak nincs licence.

---

## 2. HRL röviden, és hol jöhetne szóba

A HRL szintekre bontja a policy-t. A felső szint **subgoalokat vagy opciókat**
választ („szerezd meg a kulcsot”, „beszélj X-szel”, „juss el B scene-be”), az
alsó szint primitív akciókkal hajtja végre őket. Egy adventure játékban a
hierarchia természetes, mert a puzzle-függőségi gráf *maga* a subgoal-hierarchia.

Lehetséges felhasználások az inkwellben:

| Felhasználás | Illeszkedés | Megjegyzés |
|---|---|---|
| Automatikus playtesting, megoldhatóság | jó | softlockot, elérhetetlen állapotot, sorrendtörést talál. A sima RL elvérzik a ritka jutalmon és a verb × hotspot × item akciótéren, ezen segít a hierarchia |
| Hint rendszer | jó | ha egy ágens tudja a következő subgoalt, az *maga* a tipp |
| Tape-kommentár (Dex) | jó a struktúra | lásd 3. fejezet |
| Cselekvő NPC-k | gyenge | az adventure NPC-k viselkedése megírt; a behavior tree vagy a GOAP a determinisztikus hierarchikus alternatíva |

### 2.1 Miért nem a *tanult* HRL a jó eszköz

- **Az inkwell állapottere diszkrét és deklaratív.** A puzzle-gráf kinyerhető és
  egzaktul bejárható (BFS/A*, PDDL-szerű tervezés). Ehhez nem kell tanítás, GPU
  vagy jutalomfüggvény-hangolás.
- **Egy tape-nek nincs jutalomjele.** Egy kommentáló policy jutalma a
  játékosélmény lenne: nem unatkozik, nem akad el, nem érzi, hogy a játék
  megoldja helyette. Ezt offline nem lehet mérni, és nincs játékosadat.
- **A sorok megírtak és kétnyelvűek.** Minden sor `Text{en, hu}`. Egy tanult
  policy csak választhatna közülük, és ilyen kis készletből választani
  szabályokkal is jól megoldható.
- **Determinizmus és mentés.** A tape állapota `State` varokban él, és
  reprodukálhatóan túl kell élnie a mentést. Egy betanított modell ezt
  bonyolítja, és minden tartalomváltozás után újra kellene tanítani.
- **Eszközkészlet.** Az ML-ökoszisztéma Python, a játék Go. Tanult megközelítés
  előtt egy headless léptető API és egy gym-szerű híd (JSON stdio-n vagy gRPC)
  kellene.

Tanult HRL-nek csak akkor van értelme, ha a tartalom túl nagy a teljes
bejáráshoz, vagy *emberszerű* felfedezést akarunk szimulálni. Érdekes jövőbeli
változat **egy LLM a felső szinten**. Elolvassa a leírásokat és dialógusokat,
úgy javasol subgoalokat, ahogy egy játékos tenné, a végrehajtást pedig egy kis
kereső végzi. Ez azt teszteli, hogy a puzzle-ök *kitalálhatók*-e, nem csak azt,
hogy megoldhatók-e.

---

## 3. Javaslat: háromszintű tape-agy, tanítás nélkül

A HRL formáját megtartjuk, a tanulást keresésre és megírt adatra cseréljük.

```
1. szint  HOL TART PAUL?      szimbolikus: frontier + távolság a puzzle-gráfban
             │
2. szint  MIÉRT SZÓLJON?      utility: hallgat / figyelmeztet / tipp (finom|direkt) / lore / beszólás
             │                tape-enkénti súlyok = személyiség; cooldownok
3. szint  MELYIK SOR?         címkézett, megírt Text{en, hu} sorok
```

- **1. szint: subgoal-becslés.** Az aktuális `State`-ből kiszámolja:
  - a **frontiert**: mely szabályok teljesíthetők most, és melyek vezetnek új
    állapotba;
  - a **távolságot** a következő mérföldkőig (például `FlagNormanFound`);
  - az **elakadást**: mióta nem változott érdemben az állapot.
- **2. szint: szándék.** Utility-függvény az 1. szint kimenetén. A személyiség
  ugyanaz a struktúra más súlyokkal: Dex többet ellenkezik és kevésbé direkt
  tippet ad, egy másik tape bőbeszédűbb vagy pontosabb lehet. A restricted mode
  (licenc nélkül) korlátozza, mit mondhat egy tape, például nem adhat direkt
  tippet. A cooldownok miatt Dex nem beszéli túl a játékot.
- **3. szint: a sor.** Szándék és subgoal szerint címkézett sorok a meglévő
  `Text{en, hu}` formában, a meglévő `speaker.Dex.Say`-en keresztül.

Ez illik a fikcióhoz. A tape „figyeli” Pault és értelmezi a helyzetet, ahogy a
kánon leírja (a Sklo Socius sajtócikk: az eszköz *értelmezi* a szalag
személyiségét).

**LLM-használat:** csak **íráskor**, jelölt Dex-sorok generálására
(szándék, subgoal) párokhoz, amelyeket aztán kézzel válogatunk. Futás közben
nem, mert sérülne az i18n, a hang konzisztenciája és a determinizmus.

---

## 4. Megvalósíthatósági vizsgálat: megépíthető-e az 1. szint?

Igen. A kódból kiderült:

### 4.1 Az inkwell már most mindent lát

Az `engines/inkwell/explore.go` (a debug explorer, README 11.9) reflexióval
bejárja az egész `Game`-et: managerek → entitások → literálok → actionök →
gyerekek, **a nem exportált mezőket is**, amelyeket egy ugyanarra a memóriára
mutató friss pointeren át olvas (`unsafe`). Belelát a `Say`, `If` és `Seq`
belsejébe, és csak a closure-öknél áll meg (`Reply`, `Custom`, `Derived`). Egy
szabálykinyerő pontosan ezt az utat használhatja újra.

A feltétel- és action-típusok kis structok nem exportált mezőkkel, például
`flagCond{name}`, `andCond{cs}`, `varEqCond{name, v}`, `setFlagAction{name}`,
`seqAction{children}`, `ifAction{…}` (`inkbase/condition/*.go`,
`inkbase/action/*.go`). Az inkwell csomagon belülről ezért triviálisan
értelmezhetők.

### 4.2 A realworld szinte teljesen deklaratív

Megszámolva a `games/realworld/inc` alatt:

| Konstrukció | Darab | Gond? |
|---|---|---|
| `Custom(...)` | 0 | nincs |
| `Reply(func…)` | 2 | nem: `scene.street.go` OnLook, csak szöveg, állapotot nem változtat |
| `Derived` | 1 | nem: `scene.selector.go`, a `Validate` előtt feloldódik, utána sima exit-lista |
| saját `Eval` | 1 | nem: `stage.manager.go` `fixed`, konstans |
| `SetFlag(` | 40 | deklaratív |
| `Give(` / `TakeAway(` | 15 / 5 | deklaratív |
| `Given{` | 54 | deklaratív |
| `Show:` / `When:` | 47 / 35 | deklaratív |
| `OnUseWith` | 27 | deklaratív (`map[*Item]Action`) |

Az állapot kicsi és nevesített:
- 37 flag az `inc`-ben (`FlagPoliceTip`, `FlagHasTape`, … `FlagServerFarmMapped`);
- varok: `armilla.slot2`, `armilla.strip`, `note`, `decay`;
- kb. 15 tárgy és 38 scene.

A történet nagyrészt lineáris, így az elérhető állapotok halmaza a 2^37-hez
képest elenyésző, a BFS olcsó.

### 4.3 Mit kell lefednie a modellnek

1. **Interakciók:** hotspot × verb (`OnLook`, `OnUse`, `OnTalk`, `OnTake`,
   `OnGive`, `OnVerb`), `OnUseWith[item]`, `OnUseWithAny`, valamint exitek
   `Exit.Needs` (flag) és `Exit.Blocked` mezővel.
2. **Automatikus effektek:** `Trigger.When` → `Do` (`Once`-szal) és
   `Scene.OnEnter`/`OnLeave`. Minden lépés után fixpontig futtatandók.
3. **Dialógusok:** choice-ok `Show` feltétellel, `Once` (a `Talked`-on át), és
   `GotoNode`-hurkok (Dex root node-ja önmagára mutat vissza).
4. **Szekvencia-szemantika:** egy `Seq`-en belül az `If` elágazik, a
   `RequireItem` pedig megszakítja a sor hátralevő részét. A szabályokat ezért
   útvonalanként, útfeltételekkel kell kinyerni, nem „az action összes effektje”
   lapos halmazaként.

A `Hotspot`-nak nincs `When` mezője. Az elágazás az actionökön belül történik
(`If`-fel), ami a modellnek kedvez.

---

## 5. Megvalósítási terv

### A lépés: szabálykinyerő az inkwellben

```go
// az inkwell csomagban, hogy a nem exportált action/feltétel típusok látszódjanak
type Rule struct {
    Scene   *Scene
    Trigger string        // "hotspot:<név>/use", "usewith:<hotspot>/<item>", "exit:<scene>", "trigger:<név>", "dialog:<d>/<node>/<choice>"
    Guard   Condition     // útfeltétel: And mindenből, ami az úton van
    Effects []Effect      // SetFlag, ClearFlag, Give, TakeAway, SetVar, GoTo, StartDialog, End
}

func Rules(game *Game) []Rule
```

- A managereket úgy kell bejárni, ahogy az `explore.go` teszi.
- A `Seq`-et ki kell lapítani, az `If`-nél el kell ágazni (then / else
  `Not(cond)`-dal), a `RequireItem` kudarcánál meg kell állni, a
  `RunScript` / `RunDialog` tartalmát pedig be kell emelni.
- A closure-ök (`Reply`, `Custom`) átláthatatlanok. Az átláthatatlan
  *effektekről* validációs figyelmeztetés kell, hogy aki állapotot módosító
  `Custom`-ot ír, tudja: az elemző nem látja.
- Kell még egy könnyű `Ctx` / szimulált `State`, hogy a feltételek futó `Game`
  nélkül is kiértékelhetők legyenek. A feltételek csak a `ctx.State`-et
  olvassák, így ez egyszerű.

Becslés: 400–600 sor Go.

### B lépés: megoldhatósági teszt a realworldben

`go test`, BFS a kezdőállapotból a befejezésig (`ShowEnd`, vagy az
`awakening_finale` script). Jelenti:
- ha megoldhatatlan: a legmélyebb elért állapotokat és az ottani frontiert;
- a **softlockokat**: elérhető állapotokat, amelyekből a befejezés elérhetetlen;
- a beállított, de soha nem olvasott, illetve olvasott, de soha nem beállított
  flageket;
- a soha el nem érhető hotspotokat és dialógus-choice-okat.

Ez önmagában is hasznos, és a realworldben ma még egyetlen `_test.go` sincs.
**Ezzel érdemes kezdeni**, mert ez bizonyítja, hogy a kinyerés működik.

### C lépés: Dex 1. szintje futásidőben

Minden érdemi állapotváltozáskor (nem minden frame-ben) BFS az aktuális
`State`-ből az előre kiszámolt szabálytáblán, és ennek kimenete:

```go
type Situation struct {
    Frontier  []Rule   // mit tud Paul most értelmesen csinálni
    Next      []Rule   // az első lépések a következő mérföldkőhöz vezető legrövidebb úton
    Distance  int
    StuckFor  float64  // másodpercek az utolsó érdemi változás óta
}
```

### D lépés: a 2. és 3. szint

Utility-alapú szándékválasztó tape-enkénti súlyokkal és cooldownokkal, valamint
tape-enkénti címkézett sortáblák. A restricted mode a szándékok korlátja. Oda
kell bekötni, ahol a tape-kommentár már ma is történik: triggerek,
`dialog.dex_talk.go`, és a csatorna `ChatLog`-ja a `tape/tape.channel.go`-ban.

---

## 6. Kockázatok és nyitott kérdések

- **Útvonal-robbanás** mélyen egymásba ágyazott `If`/`Seq` esetén: ekkora
  tartalomnál nem valószínű, de a szabályokat deduplikálni kell.
- **A dialógusok `GotoNode`-ciklusaihoz** látogatott halmaz kell; a naiv
  rekurzív kibontás nem állna meg.
- **A csak időzítést érintő actionöket** (`Wait`, `Say`, `Walk`, zene) a modell
  figyelmen kívül hagyja. A megoldhatósághoz ez rendben van, de egy *időzítésen*
  múló puzzle láthatatlan maradna.
- **A scene actorok és hotspotok** állapotfüggő megjelenését nem ellenőriztük
  teljes körűen.
- **Hol legyen a kód:** a kinyerő az inkwellbe való (általános, minden
  inkwell-játék használhatja), a teszt és a tape-agy a realworldbe.
- **Nem ellenőrzött:** a realworld pontos befejezési feltétele és a többi tape
  dialógusainak belseje.

## 7. Ajánlás

1. Az **A + B** lépés prototípusként egy branch-en. Megoldhatósági és
   softlock-tesztként azonnal hasznos.
2. Ha beválik, jöhet a **C**, és egy első, egyszerű Dex 2. szint: ha Paul N
   perce elakadt, száraz, indirekt lökés a `Next` felé.
3. A tanuláshoz (LLM mint felső szintű „játékos”) csak akkor érdemes
   visszatérni, ha az lesz a cél, hogy azt teszteljük, *kitalálhatók*-e a
   puzzle-ök, és nem csak azt, hogy megoldhatók-e.

---

## 8. Megvalósítás (`hrl` branch)

A terv a megvalósítás közben három ponton változott. Mindhárom változás oka
az volt, amit a valódi játék mutatott.

### 8.1 Statikus szabálykinyerés helyett szimuláció

Az 5. fejezet egy `Rules(game)` kinyerőt javasolt, amely reflexióval
értelmezi az actionöket. Helyette `engines/inkwell/core.solve.go` a játék
**valódi actionjeit futtatja** egy ablak nélküli `*Game`-másolaton: néma
`AudioPlayer`, `MemoryStore`, előretekert óra, a motor saját útjai
(`InvokeVerb`, triggerciklus, jelenetváltás). Így nem kell újraírni a
`Seq`/`If`/`RequireItem`/`Given`/`Once` szemantikáját, a `Reply`-closure-ök és
a `HotspotPanel`-en ülő tape-foglalatok is működnek, és ami a szimulátorban
igaz, az a játékban is.

### 8.2 Teljes BFS helyett mérföldkőről mérföldkőre

A teljes szélességi bejárás 1 millió állapot után sem ért véget: a játék
nyitott, sok egymástól független puzzle-lel, és minden sorrend külön állapot.
A megoldás a 3. fejezet 1. szintje maga: **szélességi keresés a legközelebbi
új flagig** (mérföldkő).

- `Game.Situation` ezt egyszer teszi meg az aktuális állapotból (Dexnek):
  ezredmásodpercek.
- `Game.Solve` mérföldkőről mérföldkőre halad a végéig, és ha zsákutcába
  fut, visszalép és a következő legközelebbi mérföldkövet próbálja
  (mélységi bejárás, egy pillanatba legfeljebb egyszer lép be). A realworldöt
  kb. 95 másodperc alatt, kb. 415 ezer állapotból, 113 lépésben oldja meg.
- `SolveOptions.Exhaustive` kis játékhoz megtartja a teljes bejárást, amely
  minden zsákutcát megszámol.

### 8.3 `go test` helyett `-solve` kapcsoló

A realworld `CLAUDE.md` tiltja a tesztfájlokat. A megoldhatósági ellenőrzés
ezért a `-explore` mintájára futtatási kapcsoló: `go run . -solve`, amely
nullától eltérő kóddal lép ki, ha a játék nem végigjátszható.

### 8.4 Amit a solver első futása talált

1. **Javítva (realworld):** a `armilla.slot2` a kazetta *nevét* tárolja, de
   hét helyen `*Item` pointerrel hasonlították össze (`VarEq(inc.VarSlot2,
   item.SupportTape)`), így a 2. foglalatba tett kazetta semmit nem nyitott ki.
2. **Javítva (inkwell):** a `Once` választások kulcsa `<node>:<hely>` volt,
   dialógusnév nélkül; egy másik `root` node-ban elhasznált választás eltüntette
   Vasquez alkuját. Most `<dialog>/<node>:<hely>` (a régi mentések elfelejtik,
   melyik választást használták el).
3. **Nincs javítva, döntés kell:** a 2. foglalatba tett új kazetta megsemmisíti
   az ott lévőt, a support tape pedig `Once` választással vehető meg. Aki a
   support tape után, de a security room előtt betesz egy másik kazettát, nem
   tudja kijátszani az ACP-t. A `-solve` ezt jelenti első zsákutcaként, és
   megtalálja a kerülőutat.

Gyanú, nem ellenőrizve futó játékban: a `DialogBox` a választás actionjeit a
sorba (`queueAction`) teszi, a `RunDialog` viszont addig foglalja a
script-helyet, amíg a párbeszéd nyitva van. Egy ablak nélküli próbában a
választás actionjei így soha nem futottak le. A szimulátor a választást azonnal
futtatja, ahogy a README leírja.

### 8.5 Dex

A Dex-dialógus új választása: „Innen merre tovább?” / „Where do I even go from
here?”. A válasz `Reply`-jal, kimondáskor dől el: `System.Situation()` →
a legközelebbi mérföldkőhöz vezető út utolsó nem-dialógus lépése → a cél
(hotspot, tárgy, kijárat) és a helyszín Dex hangján. Csak irányt mutat, a
lépéseket nem sorolja fel (restricted mode).

**Ami hátravan:** a 2. szint időzítése, vagyis hogy Dex magától szóljon, ha a
játékos N perce nem halad; a tape-enkénti személyiség-súlyok; és hogy a
`Situation` ne javasoljon zsákutcába vezető lépést (ma a legközelebbi
mérföldkövet adja, akkor is, ha az odavezető út eldob egy később kellő
kazettát).
