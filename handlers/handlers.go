package handlers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func Handlers(){
	r:= mux.NewRouter()
	r.HandleFunc("/",Home).Methods("GET")
	r.HandleFunc("/new_reader",NewReader).Methods("GET")
	r.HandleFunc("/new_book",NewBook).Methods("GET")
	r.HandleFunc("/give",GiveBook).Methods("GET")
	r.HandleFunc("/save_reader",SaveReader).Methods("POST")
	r.HandleFunc("/readers/",ShowReaders).Methods("GET")
	r.HandleFunc("/books",AllBooks).Methods("GET")
	r.HandleFunc("/save_book", SaveBook).Methods("POST")
	r.HandleFunc("/search_reader", SearchReader).Methods("GET")
	r.HandleFunc("/reader_reg", ReaderReg).Methods("GET")
	r.HandleFunc("/search_book", SearchBook).Methods("GET")
	r.HandleFunc("/book_search", BookPage).Methods("GET")
	r.HandleFunc("/search",GiveTest).Methods("GET")
	r.HandleFunc("/reader_with_book",ReaderWithBook).Methods("GET")
	r.HandleFunc("/all_documents",AllDocuments).Methods("GET")
	r.HandleFunc("/choose_form",Choose).Methods("GET")
	r.HandleFunc("/choose",ChooseForm).Methods("GET")
	r.HandleFunc("/refund_controller",RefundController).Methods("GET")
	r.HandleFunc("/reg_offer",RedirectController).Methods("GET")
	r.HandleFunc("/refund",Refund).Methods("GET")
	r.HandleFunc("/reg_offer",RedirectController).Methods("GET")
	r.HandleFunc("/refund_book",RefundBook).Methods("GET")



	http.Handle("/",r)
	http.ListenAndServe(":8080",nil)


}
