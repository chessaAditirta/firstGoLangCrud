package main
//prepare your import
import (

	"log"
	"net/http"
	"text/template"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

//definition your table and field as same as on your database
type Person struct {
	Id int
	Identity string
	Fullname string
	Address string
	Telnum string
	Email string
}

//connection database method
func dbConn() (db *sql.DB)  {
	dbDriver := "mysql"		//your database driver
	dbUser   := "root"		//username db
	dbPass   := "root"		//pass db
	dbName   := "firstgo"	//database name
	db, error := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if error != nil {
		panic(error.Error())
	}
	return db

}

//make directory form for your all tmpl file and add var tmpl
var tmpl  = template.Must(template.ParseGlob("form/*"))

//add method Index for implementation Index or Home UI
func Index(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	selectDB, error := db.Query("SELECT * FROM  person ORDER BY id ASC ")
	if error != nil {
		panic(error.Error())
	}

	person	:= Person{}
	res		:= []Person{}

	for selectDB.Next() {
		var id int
		var identity, fullname, address, telnum, email string

		error = selectDB.Scan(&id, &identity, &fullname,
			&address, &telnum, &email)

		if error != nil {
			panic(error.Error())
		}

		person.Id = id
		person.Identity = identity
		person.Fullname = fullname
		person.Address = address
		person.Telnum = telnum
		person.Email = email

		res = append(res, person)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//get person data by id
func Show(w http.ResponseWriter, r *http.Request)  {
	db := dbConn()
	nPersonId := r.URL.Query().Get("id")
	selectDB, error := db.Query("SELECT * FROM person WHERE id=?", nPersonId)
	if error != nil {
		panic(error.Error())
	}
	person := Person{}
	for selectDB.Next() {
		var id int
		var identity, fullname, address,
		telnum, email string

		error = selectDB.Scan(&id, &identity, &fullname,
			&address, &telnum, &email)

		if error != nil {
			panic(error.Error())
		}
		person.Id = id
		person.Identity = identity
		person.Fullname = fullname
		person.Address = address
		person.Telnum = telnum
		person.Email = email
	}
	tmpl.ExecuteTemplate(w, "Show", person)
	defer db.Close()
}

//add or post person entity
func Create(w http.ResponseWriter, r *http.Request)  {
	tmpl.ExecuteTemplate(w, "Create", nil)
}

//edit person entity by id
func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nPersonId := r.URL.Query().Get("id")
	selectDB, error := db.Query("SELECT * FROM person WHERE id=?", nPersonId)
	if error != nil {
		panic(error.Error())
	}
	person := Person{}
	for selectDB.Next() {
		var id int
		var identity, fullname, address,
		telnum, email string

		error = selectDB.Scan(&id, &identity, &fullname,
			&address, &telnum, &email)

		if error != nil {
			panic(error.Error())
		}
		person.Id = id
		person.Identity = identity
		person.Fullname = fullname
		person.Address = address
		person.Telnum = telnum
		person.Email = email
	}
	tmpl.ExecuteTemplate(w, "Edit", person)
	defer db.Close()
}

//insert data to person table from firs_go database
func Insert(w http.ResponseWriter,  r *http.Request)  {
	db := dbConn()
	if r.Method == "POST" {
		identity := r.FormValue("identity")
		fullname := r.FormValue("fullname")
		address := r.FormValue("address")
		telnum := r.FormValue("telnum")
		email := r.FormValue("email")
		insertForm, error := db.Prepare("INSERT INTO person(identity, fullname, address, telnum, email) VALUES(?,?,?,?,?)")
		if error != nil {
			panic(error.Error())
		}
		//insertForm.Exec(identity_number, fullname, address, post_code, email, gender)
		insertForm.Exec(identity, fullname, address, telnum, email)
		log.Println("INSERT: Identity Number:"+identity+" |Fullname:"+fullname+" |Address:"+address+
			" |Phone Number:"+telnum+" |email:"+email)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

//update data person table from first_go database by id
func Update(w http.ResponseWriter, r *http.Request)  {
	db := dbConn()
	if r.Method == "POST" {
		identity := r.FormValue("identity")
		fullname := r.FormValue("fullname")
		address := r.FormValue("address")
		telnum := r.FormValue("telnum")
		email := r.FormValue("email")
		id := r.FormValue("pid")
		//insertForm, error := db.Prepare("UPDATE person SET identity_number=?, " +
		//	"fullname=?, address=?, post_code=?, email=?, gender=? WHERE person_id=?")
		insertForm, error := db.Prepare("UPDATE person SET identity=?, " +
			"fullname=?, address=?, telnum=?, email=? WHERE id=?")
		if error != nil {
			panic(error.Error())
		}
		insertForm.Exec(identity, fullname, address, telnum, email, id)
		log.Println("UPDATE: Identity Number:"+identity+" |Fullname:"+fullname+
			" |Address:"+address+" |Phone Number:"+telnum+" |Email:"+email)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

//delete record in person table by id
func Delete(w http.ResponseWriter, r *http.Request)  {
	db := dbConn()
	nPersonId := r.URL.Query().Get("id")
	deleteForm, error := db.Prepare("DELETE FROM person WHERE id=?")
	if error != nil {
		panic(error.Error())
	}
	deleteForm.Exec(nPersonId)
	log.Println("DELETE successfully")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	//runnning background
	log.Println("Server started on: http://redvelvet.misterc:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	//http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", nil)
}
