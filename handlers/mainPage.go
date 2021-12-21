package handlers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"strconv"
	"time"
)
type Reader struct{
	Id uint16
	Name string `json:"name"`
	Birthdate string `json:"birthdate"`
	Adress string `json:"adress"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Debt uint16 `json:"debt"`
}

type Book struct{
	Book_id  uint16 `json:"book_id"`
	BookName string `json:"book_name"`
	Genre    string `json:"genre"`
	Year uint16 `json:"year"`
	Quantity uint16 `json:"quantity"`
	Available uint16 `json:"available"`
	Registration string `json:"registration"`
	Price uint16 `json:"book_price"`
}

type Document struct{
	DocId uint16 `json:"doc_id"`
	ReaderSurname string `json:"reader_surname"`
	BookName string `json:"book"`
	Date string `json:"date"`
	Price float64 `json:"price"`
	QuantityBook uint16 `json:"quant"`
}

type Client struct{
	Reader
	Book
	Penny float64
}

var readers []Reader
var books []Book
var clients []Client
var documents []Document
var surname string
var bookName string
var penny float64

func Home(w http.ResponseWriter, r *http.Request){

	t, err := template.ParseFiles("templates/index.html","templates/footer.html","templates/header.html")

	if err != nil{
		fmt.Fprintf(w,err.Error())
	}

	db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
	if err != nil {
		panic(err)
	}


	defer db.Close()

	res, err := db.Query(fmt.Sprintf("Select * from `books` ORDER BY book_name"))
	if err!=nil{
		panic(err)
	}

	books = []Book{}
	for res.Next() {
		var book Book
		err = res.Scan(&book.Book_id,&book.BookName, &book.Genre, &book.Year, &book.Quantity, &book.Available,&book.Registration, &book.Price)
		if err!=nil{
			panic(err)
		}
		books = append(books, book)
	}
	t.ExecuteTemplate(w,"index",books)

}

func NewReader(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/new_reader.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	t.ExecuteTemplate(w,"new_reader",nil)
}

func SaveReader(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	birthdate := r.FormValue("birthdate")
	email := r.FormValue("email")
	adress := r.FormValue("adress")

	if name == "" || surname == "" || birthdate == "" || email == "" || adress == "" {
		fmt.Fprintf(w, "Не все поля заполнены")
	} else {

		db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `readers` (`name`,`surname`,`birthdate`,`email`,`adress`) VALUES ('%s','%s','%s','%s','%s')", name, surname, birthdate, email, adress))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ShowReaders(w http.ResponseWriter, r *http.Request){

	t, err := template.ParseFiles("templates/readers.html","templates/footer.html","templates/header.html")

	if err != nil{
		fmt.Fprintf(w,err.Error())
	}


	db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
	if err != nil {
		panic(err)
	}

	defer db.Close()


	res, err := db.Query(fmt.Sprintf("Select * from `readers` ORDER BY name"))
	if err!=nil{
		panic(err)
	}

	readers = []Reader{}
	for res.Next() {
		var reader Reader
		err = res.Scan(&reader.Id,&reader.Name, &reader.Surname, &reader.Birthdate, &reader.Adress, &reader.Email, &reader.Debt)
		if err!=nil{
			panic(err)
		}
		readers = append(readers,reader)
	}


	t.ExecuteTemplate(w,"readers" ,readers)
}

func AllBooks(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/books.html","templates/footer.html","templates/header.html")

	if err != nil{
		fmt.Fprintf(w,err.Error())
	}


	db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
	if err != nil {
		panic(err)
	}

	defer db.Close()


	res, err := db.Query(fmt.Sprintf("Select * from `books` ORDER BY book_name, available DESC"))
	if err!=nil{
		panic(err)
	}

	books = []Book{}
	for res.Next() {
		var book Book
		err = res.Scan(&book.Book_id,&book.BookName, &book.Genre, &book.Year, &book.Quantity, &book.Available, &book.Registration, &book.Price)
		if err!=nil{
			panic(err)
		}
		books = append(books, book)
	}


	t.ExecuteTemplate(w,"books" ,books)
}

func SaveBook(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("book_name")
	genre := r.FormValue("genre")
	year := r.FormValue("year")
	quantity := r.FormValue("quantity")
	registration := time.Now()
	reg := registration.Format("2006-01-02")
	price := r.FormValue("book_price")
	available := quantity

	if name == "" || genre == "" || year == "" || quantity == "" {
		fmt.Fprintf(w, "Не все поля заполнены")
	} else {

		db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `books` (`book_name`,`genre`,`year`,`quantity`, `registration`, `book_price`, `available`) VALUES ('%s','%s','%s','%s','%s','%s','%s')", name, genre, year, quantity, reg, price,available))
		if err != nil {
			panic(err)
		}
		defer insert.Close()


		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func NewBook(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/new_book.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	t.ExecuteTemplate(w,"new_book",nil)
}

func GiveBook(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/give.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	t.ExecuteTemplate(w,"give",nil)
}

func SearchReader(w http.ResponseWriter, r *http.Request) {
	surname = r.FormValue("surname")


	if surname == "" {
		fmt.Fprintf(w, "Не все поля заполнены")
	} else {

		db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}

		defer db.Close()

			res := db.QueryRow("Select * from `readers` WHERE surname = ?",surname)


		reader := new(Reader)
		err = res.Scan(&reader.Id, &reader.Name, &reader.Surname, &reader.Birthdate, &reader.Adress, &reader.Email, &reader.Debt)
		if err == sql.ErrNoRows {
			http.Redirect(w, r, "/reader_reg", http.StatusSeeOther)
		} else if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		if reader.Debt > 5{
			fmt.Fprintf(w,"У вас задолженность по книгам. Пожалуйста, погасите её, чтобы взять новую")
		} else {
			http.Redirect(w, r, "/book_search", http.StatusSeeOther)
		}
	}


	}

func ReaderReg(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/reader_reg.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	t.ExecuteTemplate(w,"reader_reg",nil)
}

func BookPage(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/book_search.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	t.ExecuteTemplate(w,"book_search",nil)
}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")


	if name == "" {
		fmt.Fprintf(w, "Не все поля заполнены")
	} else {

		db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		res := db.QueryRow("Select * from `books` WHERE name = ?",name)


		book := new(Book)
		err = res.Scan(&book.Book_id,&book.BookName, &book.Genre, &book.Year, &book.Quantity, &book.Available, &book.Registration, &book.Price)
		if err == sql.ErrNoRows {
			fmt.Fprintf(w,"Книга не найдена")
			return
		} else if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		fmt.Fprintf(w,"Книга найдена, будете брать?")
	}


}

func GiveTest(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/search.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	t.ExecuteTemplate(w,"search",nil)
}

func ReaderWithBook(w http.ResponseWriter, r *http.Request) {
	surname = r.FormValue("surname")

	if surname == "" {
		fmt.Fprintf(w, "Не все поля заполнены")
	} else {

		db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}

		defer db.Close()


		resReader := db.QueryRow("Select * from `readers` WHERE surname = ?", surname)

		reader := new(Reader)
		errReader := resReader.Scan(&reader.Id, &reader.Name, &reader.Birthdate, &reader.Adress, &reader.Surname, &reader.Email, &reader.Debt)
		if errReader == sql.ErrNoRows {
			http.Redirect(w, r, "/reg_offer", http.StatusSeeOther)
		} else if errReader != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		if reader.Debt != 0 {
			fmt.Fprintf(w, "Чтобы взять новую книгу, необходимо вернуть старые")
			return
		}
		http.Redirect(w, r, "/choose", http.StatusSeeOther)

	}
}

func ChooseForm(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/choose.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	t.ExecuteTemplate(w,"choose",nil)
}

func AllDocuments(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/all_documents.html","templates/footer.html","templates/header.html")

	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("Select * from `documents`"))
	if err!=nil{
		panic(err)
	}

	documents = []Document{}
	for res.Next() {
		var document Document
		err = res.Scan(&document.DocId,&document.ReaderSurname,&document.BookName, &document.Date, &document.Price, &document.QuantityBook)
		if err!=nil{
			panic(err)
		}
		documents = append(documents, document)
	}


	t.ExecuteTemplate(w,"all_documents" ,documents)
}

func Choose(w http.ResponseWriter, r *http.Request) {
	bookName = r.FormValue("book_name")
	quant := r.FormValue("quant")
	if bookName == "" || quant == "" {
		fmt.Fprintf(w, "Не все поля заполнены")
	} else {
		db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}

		defer db.Close()
		book := new(Book)

		number, _ := strconv.Atoi(quant)

		resBook := db.QueryRow("Select * from `books` WHERE book_name = ?", bookName)
		errBook := resBook.Scan(&book.Book_id, &book.BookName, &book.Genre, &book.Year, &book.Quantity, &book.Available, &book.Registration, &book.Price)
		if errBook == sql.ErrNoRows {
			fmt.Fprintf(w, "Книга не найдена")
			return
		}
		if number > int(book.Quantity) {
			fmt.Fprintf(w,"Столько книг нет в наличии")
			return
		}
		if number > 5{
			fmt.Fprintf(w,"Можно взять не более 5 книг")
			return
		}
		updDebt := db.QueryRow("UPDATE `readers` set debt = debt+1 WHERE surname = ?", surname)

		updDebt.Err()

		if book.Available > 0 {
			updBook := db.QueryRow("UPDATE `books` SET available=available-? where book_name = ?", number, bookName)
			updBook.Err()
		}

		registration := time.Now().Add(720 * time.Hour).Format("2006-01-02")


		price := book.Price
		if number == 2 || number == 3 {
			price -= price * 10 / 100
		} else if number >= 4 {
			price -= price * 15/100
		}

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `documents` (`reader_surname`,`book`,`date`,`price`,`quant`) VALUES ('%s','%s','%s','%d','%s')", surname, bookName, registration, price, quant))
		if err != nil {
			panic(err)
		}
		defer insert.Close()
		http.Redirect(w, r, "/all_documents", http.StatusSeeOther)
	}
}

func RefundController(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/refund_controller.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	t.ExecuteTemplate(w,"refund_controller",nil)
}

func Refund(w http.ResponseWriter, r *http.Request) {
	surname = r.FormValue("surname")

	if surname == "" {
		fmt.Fprintf(w, "Не все поля заполнены")
	} else {

		db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
		if err != nil {
			panic(err)
		}

		defer db.Close()


		resReader := db.QueryRow("Select * from `readers` WHERE surname = ?", surname)

		reader := new(Reader)
		errReader := resReader.Scan(&reader.Id, &reader.Name, &reader.Birthdate, &reader.Adress, &reader.Surname, &reader.Email, &reader.Debt)
		if errReader == sql.ErrNoRows {
			http.Redirect(w, r, "/reg_offer", http.StatusSeeOther)
		} else if errReader != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		document := new(Document)
		resDoc := db.QueryRow("Select * from `documents` WHERE reader_surname = ?", surname)
		errDoc := resDoc.Scan(&document.DocId,&document.ReaderSurname,&document.BookName, &document.Date, &document.Price, &document.QuantityBook)
		if errDoc == sql.ErrNoRows {
			fmt.Fprintf(w,"У данного пользователя нет книг для возврата")
			return
		} else if errDoc != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		t1 := document.Date
		dt1,err := time.Parse("2006-01-02",t1)
		t2 := time.Now()
		if t2.After(dt1){
			days :=(t2.Sub(dt1).Hours())/24
			penny = document.Price * (0.01) * (days)
		} else if t2.Before(dt1){
			penny = 0
		}

		quant := document.QuantityBook
		bookName = document.BookName
		updDebt := db.QueryRow("UPDATE `readers` set debt = debt-1 WHERE surname = ?", surname)

		updDebt.Err()

		updBook := db.QueryRow("UPDATE `books` SET available=available+? where book_name = ?", quant, bookName)

		updBook.Err()

		updDoc := db.QueryRow("DELETE  from `documents` where reader_surname = ?",surname)

		updDoc.Err()

		http.Redirect(w, r, "/refund_book", http.StatusSeeOther)
	}
}

func RedirectController(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/reg_offer.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	t.ExecuteTemplate(w,"reg_offer",nil)
}

func RefundBook(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/refund_book.html","templates/footer.html","templates/header.html")
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	reader := new(Reader)
	resRead := db.QueryRow("Select * from `readers` WHERE surname = ?", surname)
	errRead := resRead.Scan(&reader.Id, &reader.Name, &reader.Birthdate, &reader.Adress, &reader.Surname, &reader.Email, &reader.Debt)
	if errRead == sql.ErrNoRows {
		fmt.Fprintf(w,"Пользователь не найден")
		return
	} else if errRead != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	client := new(Client)
	client.BookName = bookName
	client.Surname = surname
	client.Penny = penny
	clients = []Client{}
	clients = append(clients,*client)
	t.ExecuteTemplate(w,"refund_book",clients)
}