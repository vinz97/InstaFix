Progetto InstaFix APL made by Vincenzo PLUCHINO & Philip TAMBE'
https://github.com/vinz97/InstaFix

ISTRUZIONI PER ESEGUIRE IL CODICE

ATTENZIONE: Le seguenti istruzioni sono state testate e verificate su ambiente Windows
Su SO diversi come Mac o Linux potrebbero non funzionare


--- CLIENT C# --- 
Si consiglia di eseguire il codice su Microsoft Visual Studio (non Visual Studio Code)
In altri ambienti o IDE potrebbero sorgere dei problemi per la GUI dei Windows Forms
La cartella da scaricare da Github è "ClientUtenteCS" e il file "ClientUtenteCS.sln"
L'unico accorgimento è installare la libreria Newtonsoft. Per farlo, si osservino i seguenti passi:
--> su Visual Studio 2022, andare in alto nella voce PROGETTO (si trova in mezzo tra GIT e Compilazione)
--> cliccare sulla terzultima voce GESTISCI PACCHETTI NUGET...
--> cercare sulla barra di ricerca NETWONSOFT e cliccare per l'installazione automatica



--- SERVER GO --- 
L'IDE consigliato è GoLand di JetBrains.
Le folder da scaricare da Github sono "ServerGolang" e ".idea"
Per una corretta esecuzione innanzitutto è necessario cambiare in alto a destra su GoLand
le configurazioni di Run/Debug (la stringa accanto ai pulsanti verdi di run e debug): deve essere impostata su "go build *.go"
Riguardo le librerie da installare, si visiti la pagina https://github.com/go-sql-driver/mysql
alla voce installation è indicata la riga di comando da eseguire
N.B. è necessario avere GIT installato 


--- CLIENT PYTHON ---