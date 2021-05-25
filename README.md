Aufgabe 3: Nebenläufige Programme mit Go \
In dieser Aufgabe werden Sie ein nebenläufige Go-Anwendung entwickeln, die explizit das "goroutine"-Konzept nutzt.
Weil wir das Problem bereits in mCRL2 analysiert haben und "korrekt" bewiesen haben, werden wir diese Information bei der Umsetzung in Go nutzen.

Aufgabe 3. Nebenläufige Programme mit Go:

Nutzen Sie die letzte Version der mCRL2-Spezifikation aus der vorherigen Aufgabe als Basis für die Implementierung der Ampelschaltung in Go. 
Auch hier soll es wieder vier Instanzen (4 x goroutine) geben, die ohne einer zentralen Steuerung o.ä. miteinander kommunizieren.
Für diese Kommunikation können nur Channels genutzt werden. Die Verwendung von globalen Variablen um zu kommunizieren oder zu synchronisieren ist explizit nicht erlaubt!

Hinweise:
- Bleiben Sie so nah wie möglich an der mCRL2-Spezifikation damit die geprüfte Äquivalenz übertragen werden kann und nutzen Sie z.B. auch die gleichen Bezeichner.
- Nutzen Sie, wegen der Äquivalenz auch weiterhin die synchrone (nicht gepufferte) Kommunikation in Go; d.h.: make(chan <type>) und nicht make(chan <type>, n).
    - Es gibt hierzu allerdings eine Ausnahme. Wenn Sie doch unbedingt ansynchrone Kommunikation in Go nutzen möchten, ist dies erlaubt, wenn Sie die Äquivalenz mittels einer erweiterten mCRL2-Spezifikation beweisen.
- Die Endrekursion aus mCRL2 wird am besten mit einer Schleife implementiert.
- Nutzen Sie unbedingt die Tipps zur Umsetzung der ungerichteten (synchronen) Kommunikation aus der Vorlesung zum Thema: "Go". Stichwort: select (siehe: HelloBonjour-Beispiel)
- Ihre Lösung soll mit maximaler Geschwindigkeit funktionieren und keine Sleep-Timer enthalten, die benutzt werden um Deadlocks und Race Conditions zu umgehen. Hierzu brauchen Sie wirklich das Wissen, das Sie sich in diesem Semester angeeignet haben.
- Nutzen Sie für die Channels "kleine" Datentypen (am besten Enumerations) und keine Zeichenketten. Die mCRL2-Spezifikation nutzt ja auch keine Zeichenketten!
- Nutzen für das quitChannel und für weitere Channel, die nur zur Synchronisation genutzt werden, am besten ein Channel vom Typ: bool.

Allgemeine Hinweise zur Programmierung:

- Vermeiden Sie die Verwendung von "Magic Numbers".
- Vermeiden Sie analog auch die Verwendung von "Magic Strings".
- Vermeiden Sie die Benutzung von goto.
- Vergessen Sie nicht die Kommentare.
- Vergessen Sie nicht die Formatierung, z.B. Operatoren brauchen "Luft".
- Vermeiden Sie sehr lange Funktionen, prüfen Sie, ob Sie Teil-Funktionalität in einer eigenen Funktion auslagern können.
- DRY: Don't Repeat Yourself.
