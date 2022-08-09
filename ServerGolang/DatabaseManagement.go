package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

// classi presenti nel DB
type Utente struct {
	Id	int
	Nome	string
	Cognome string
	Citta	string
	Indirizzo	string
}

// i nomi dei campi degli struct sono indicati con la lettera iniziale maiuscola per alcuni problemi
// derivati dall'incapsulamento dello struct stesso in json
type Ticket struct {
	Id int        //  `json:"id"`
	Stato string    //  `json:"stato"`
	Categoria string  //  `json:"categoria"`
	Titolo string     //  `json:"titolo"`
	Descrizione string  //  `json:"descrizone"`
}
type Professionisti struct {
	Id int
	Nome string
	Cognome string
	Professione string
	Recensione float32
	Citta string

}
type Preventivi struct {
	Id int
	IdTicket int
	IdProfessionista int
	Descrizione string
	MaterialiRicambi string
	Costo float32
	DataOra string
}


func startDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/instafix")
	//defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
	fmt.Println("DB avviato")
	return db
}

func esempioQuery(datab *sql.DB) {
	res, err := datab.Query("SELECT * FROM utenti")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var utente Utente
		err := res.Scan(&utente.Id, &utente.Nome, &utente.Cognome, &utente.Citta, &utente.Indirizzo)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", utente)
	}

}

func registerQuery(datab *sql.DB, name string, surname string, city string, address string, email string, number string, passw string) string {
	res1, err1 := datab.Query("SELECT * FROM credenziali WHERE email="+ "'" + email + "'")
	defer res1.Close()
	if err1 != nil {
		log.Fatal(err1)
	}
	if res1.Next() == true {
		return "Email esistente"
	}
	query1 := fmt.Sprintf("INSERT INTO utenti (nome, cognome, citta, indirizzo) VALUES ('%s', '%s', '%s', '%s' )",
		name, surname, city, address)
	res2, err2:= datab.Query(query1)
	defer res2.Close()
	query2 := fmt.Sprintf("INSERT INTO credenziali (email, id_utente, psw, telefono) VALUES" +
		"('%s', (SELECT id_utente FROM utenti WHERE nome='%s' AND cognome='%s' AND indirizzo='%s')," +
		" '%s', '%s')", email, name, surname, address, passw, number)
	res3, err3:= datab.Query(query2)
	defer res3.Close()
	if err2 != nil || err3 != nil {
		log.Fatal(err2)
		log.Fatal(err3)
		return "Errore generico"
	}
	return "Account creato"

}

func loginQuery(datab *sql.DB, email string, psw string) string{
	query := fmt.Sprintf("SELECT email FROM credenziali WHERE email= '%s' AND psw='%s'", email, psw)
	result1, err1:= datab.Query(query)
	defer result1.Close()

	if result1.Next() == false {
		return "Credenziali errati"
	}
	if err1 != nil {
		log.Fatal(err1)
		return "Errore generico"
	}
	return "Credenziali corrette"

}

func newticketQuery (datab *sql.DB, title string, category string, desc string, email string) string{
	query1 := fmt.Sprintf("SELECT id_utente FROM utenti WHERE id_utente=" +
		"(SELECT id_utente FROM credenziali WHERE email='%s')", email)
	res1, err1 := datab.Query(query1)
	if err1 != nil {
		log.Fatal(err1)
		return "Errore generico"
	}
	defer res1.Close()
	var idUser int
	// estrapoliamo l'id utente trovato nella query precedente per poterlo utilizzare nella query successiva
	for res1.Next() {
		err := res1.Scan(&idUser)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(idUser)
	}

	if idUser > 0   {
		query2 := fmt.Sprintf("INSERT INTO tickets (id_utente, categoria, titolo, descrizione, stato)" +
			"VALUES ('%d', '%s', '%s', '%s', 'creato')",idUser, category, title, desc)
		result, err := datab.Query(query2)
		if err != nil {
			log.Fatal(err)
			return "Errore generico"
		}
		defer result.Close()
		return "Ticket creato"
	}
	return "Utente non trovato"
}

func getnomeQuery (datab *sql.DB, mail string) string {
	query:= fmt.Sprintf("SELECT nome FROM utenti WHERE id_utente=(SELECT id_utente FROM credenziali " +
		"WHERE email='%s')", mail)
	res, err := datab.Query(query)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
		return "Errore generico"
	}

	var nome string
	for res.Next() {
		err2 := res.Scan(&nome)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	return nome
}

func getTickets (datab *sql.DB, mail string) []Ticket {
	q1 := fmt.Sprintf("SELECT id_utente FROM utenti WHERE id_utente=(SELECT id_utente FROM credenziali " +
	"WHERE email='%s')", mail)
	r1, err1 := datab.Query(q1)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer r1.Close()
	var id int
	for r1.Next() {
		err2 := r1.Scan(&id)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	q2 := fmt.Sprintf("SELECT id_ticket, stato, categoria, titolo, descrizione FROM tickets WHERE " +
		"id_utente='%d'", id)
	r2, err3 := datab.Query(q2)
	defer r2.Close()
	if err3 != nil {
		log.Fatal(err3)
	}

	var richiesta Ticket
	var tickets []Ticket
	// abbiamo bisogno di un array perchè un utente potrebbe avere più ticket creati alla volta
	for r2.Next() {
		err4 := r2.Scan(&richiesta.Id, &richiesta.Stato, &richiesta.Categoria, &richiesta.Titolo, &richiesta.Descrizione)
		if err4 != nil {
			log.Fatal(err4)
		}
		tickets = append (tickets, richiesta)
	}

	return tickets
}

func getProfessionistiQuery(datab *sql.DB, category string) []Professionisti{
	query := fmt.Sprintf("SELECT id_professionista, nome, cognome, professione, recensione, citta FROM professionisti" +
		" WHERE professione='%s' ORDER BY recensione DESC", category)
	resp, err := datab.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Close()


	var lavoratore Professionisti
	var professionisti []Professionisti
	for resp.Next() {
		err2 := resp.Scan(&lavoratore.Id, &lavoratore.Nome, &lavoratore.Cognome, &lavoratore.Professione,
			&lavoratore.Recensione, &lavoratore.Citta)
		if err2 != nil {
			log.Fatal(err2)
		}

		professionisti = append(professionisti, lavoratore)
	}
	return professionisti
}

func selectProfessionistaQuery(datab *sql.DB, mail string, id_professionist string) string {
	query1:= fmt.Sprintf("SELECT id_utente FROM utenti WHERE id_utente=(SELECT id_utente FROM credenziali " +
		"WHERE email='%s')", mail)
	res1, err1 := datab.Query(query1)
	defer res1.Close()
	if err1 != nil {
		log.Fatal(err1)
		return "Errore generico"
	}
	var idUser int
	for res1.Next() {
		err2 := res1.Scan(&idUser)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	idProfessionist, err := strconv.Atoi(id_professionist)
	if err != nil {
		// ... handle error
		panic(err)
	}

	// aggiorna il ticket inserendo l'id del professionista scelto dall'utente
	// viene sempre preso il ticket dell'utente con l'id maggiore in quanto è l'ultimo creato dall'utente
	query2:= fmt.Sprintf("UPDATE tickets SET id_professionista=%d WHERE" +
		" id_ticket= (SELECT MAX(id_ticket) FROM tickets WHERE id_utente=%d)", idProfessionist, idUser)
	resUpdate, errUpdate := datab.Query(query2)
	defer resUpdate.Close()
	if errUpdate != nil {
		log.Fatal(errUpdate)
		return "Errore generico"
	}
	return "ok"
}

func getPreventiviQuery (datab *sql.DB, mail string) []Preventivi{
	q1 := fmt.Sprintf("SELECT id_utente FROM utenti WHERE id_utente=(SELECT id_utente FROM credenziali " +
		"WHERE email='%s')", mail)
	r1, err1 := datab.Query(q1)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer r1.Close()
	var idUser int
	for r1.Next() {
		err2 := r1.Scan(&idUser)
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	q2 := fmt.Sprintf("SELECT * FROM preventivi WHERE id_ticket IN (SELECT id_ticket FROM tickets WHERE " +
		"id_utente= %d AND stato='in attesa')", idUser)
	r2, err3 := datab.Query(q2)
	defer r2.Close()
	if err3 != nil {
		log.Fatal(err3)
	}

	var preventivo Preventivi
	var listaPreventivi []Preventivi

	for r2.Next() {
		err4 := r2.Scan(&preventivo.Id, &preventivo.IdTicket, &preventivo.IdProfessionista, &preventivo.Descrizione, &preventivo.MaterialiRicambi, &preventivo.Costo, &preventivo.DataOra)
		if err4 != nil {
			log.Fatal(err4)
		}
		listaPreventivi = append (listaPreventivi, preventivo)
	}

	return listaPreventivi
}

func getInfoPreventivoQuery(datab *sql.DB, id string) Preventivi {
	idPreventivo, err1 := strconv.Atoi(id)
	if err1 != nil {
		panic(err1)
	}

	query:= fmt.Sprintf("SELECT * FROM preventivi WHERE id_preventivo= %d", idPreventivo)
	response, errore := datab.Query(query)
	defer response.Close()
	if errore != nil {
		log.Fatal(errore)
	}

	var infoPreventivo Preventivi

	for response.Next() {
		err := response.Scan(&infoPreventivo.Id, &infoPreventivo.IdTicket, &infoPreventivo.IdProfessionista, &infoPreventivo.Descrizione, &infoPreventivo.MaterialiRicambi, &infoPreventivo.Costo, &infoPreventivo.DataOra)
		if err != nil {
			log.Fatal(err)
		}
	}

	return infoPreventivo
}

func acceptPreventivoQuery(datab *sql.DB, id string) string {
	idPreventivo, err1 := strconv.Atoi(id)
	if err1 != nil {
		panic(err1)
		return "Errore generico"
	}

	query := fmt.Sprintf("UPDATE TICKETS set stato='in corso' WHERE id_ticket= (SELECT id_ticket FROM" +
		" preventivi WHERE id_preventivo= %d)", idPreventivo)

	response, errore := datab.Query(query)
	defer response.Close()
	if errore != nil {
		log.Fatal(errore)
		return "Errore generico"
	}

	return "Preventivo accettato"
}

func denyPreventivoQuery(datab *sql.DB, id string) string {
	idPreventivo, err1 := strconv.Atoi(id)
	if err1 != nil {
		panic(err1)
		return "Errore generico"
	}

	query := fmt.Sprintf("UPDATE TICKETS set stato='rifiutato' WHERE id_ticket= (SELECT id_ticket FROM" +
		" preventivi WHERE id_preventivo= %d)", idPreventivo)

	response, errore := datab.Query(query)
	defer response.Close()
	if errore != nil {
		log.Fatal(errore)
		return "Errore generico"
	}

	return "Preventivo rifiutato"
}

func downloadFatturaQuery(datab *sql.DB, mail string, id_ticket string) string {
	idTicket, err1 := strconv.Atoi(id_ticket)
	if err1 != nil {
		panic(err1)
		return "Errore generico"
	}

	q1 := fmt.Sprintf("SELECT id_utente FROM utenti WHERE id_utente=(SELECT id_utente FROM credenziali " +
		"WHERE email='%s')", mail)
	r1, err1 := datab.Query(q1)
	if err1 != nil {
		log.Fatal(err1)
		return "Errore generico"
	}
	defer r1.Close()
	var idUser int
	for r1.Next() {
		err2 := r1.Scan(&idUser)
		if err2 != nil {
			log.Fatal(err2)
			return "Errore generico"
		}
	}
	// si controlla se la fattura che l'utente vuole scaricare si riferisce a un ticket aperto da lui
	// e anche se il ticket è stato completato (stato finito)
	q2 := fmt.Sprintf("SELECT path_fattura FROM fatture WHERE id_ticket= (SELECT id_ticket FROM tickets" +
		" WHERE id_ticket= %d AND id_utente= %d AND stato='finito')", idTicket, idUser)
	r2, err2 := datab.Query(q2)
	if err2 != nil {
		log.Fatal(err2)
		return "Errore generico"
	}
	defer r2.Close()
	var pathFile string
	for r2.Next() {
		err3 := r2.Scan(&pathFile)
		if err3 != nil {
			log.Fatal(err3)
			return "Errore generico"
		}
	}

	if pathFile == "" {
		return "Non disponibile"
	}

	// le fatture vengono salvate nella directory "fatture" del server locale
	// ES http://localhost/fatture/fatt_31.pdf

	return pathFile
}