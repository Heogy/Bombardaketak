package gramatika

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Denbora string

const (
	Orainaldia Denbora = "Orainaldia"
	Lehenaldia Denbora = "Lehenaldia"
)

type Nor string

const (
	Ni    Nor = "Ni"
	Hura  Nor = "Hura"
	Gu    Nor = "Gu"
	Zu    Nor = "HurZua"
	Zuek  Nor = "Zuek"
	Haiek Nor = "Haiek"
)

type Nori string

const (
	Niri  Nori = "Niri"
	Hari  Nori = "Hari"
	Guri  Nori = "Guri"
	Zuri  Nori = "Zuri"
	Zuei  Nori = "Zuei"
	Haiei Nori = "Haiei"
)

type Nork string

const (
	Nik       Nork = "Niri"
	Hark      Nork = "Hari"
	Guk       Nork = "Guri"
	Zuk       Nork = "Zuk"
	NorkZuek  Nork = "Zuek"
	NorkHaiek Nork = "Haiek"
)

type Guess struct {
	Denbora   Denbora `json:"denbora" validate:"required"`
	Nor       Nor     `json:"nor"`
	Nori      Nori    `json:"nori"`
	Nork      Nork    `json:"nork"`
	Erantzuna string  `json:"erantzuna"`
}

type GuessNorNori struct {
	Denbora   Denbora `json:"denbora" validate:"required"`
	Nor       Nor     `json:"nor"`
	Nori      Nori    `json:"nori"`
	Nondik    string  `json:"nondik"`
	Erantzuna string  `json:"erantzuna"`
}
type NorNori struct {
	RandomNor  Nor  `json:"nor"`
	RandomNori Nori `json:"nori"`
}
type NorNoriNork struct {
	RandomNor  Nor  `json:"nor"`
	RandomNori Nori `json:"nori"`
	RandomNork Nork `json:"nork"`
}

func VerifyNorNori(guess GuessNorNori) (bool, string, error) {
	erantzuna, err := resolveNorNori(guess)
	if err != nil {
		return false, "", err
	}
	return strings.EqualFold(strings.TrimSpace(erantzuna), strings.TrimSpace(guess.Erantzuna)), erantzuna, nil
}

func Verify(guess Guess) (bool, string, error) {
	erantzuna, err := resolve(guess)
	if err != nil {
		return false, "", err
	}
	return strings.EqualFold(strings.TrimSpace(erantzuna), strings.TrimSpace(guess.Erantzuna)), erantzuna, nil
}
func resolveNorNori(guess GuessNorNori) (string, error) {

	db, err := openConnection()
	defer db.Close()

	aditzLaguntzilea := ""
	row := db.QueryRow(`select aditz_lagunzailea from nor_nori where nor = $1 and nori = $2 and denbora = $3;`, guess.Nor, guess.Nori, guess.Denbora)

	switch err := row.Scan(&aditzLaguntzilea); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(aditzLaguntzilea)
	default:
		panic(err)
	}
	return aditzLaguntzilea, err
}
func resolve(guess Guess) (string, error) {

	db, err := openConnection()
	defer db.Close()

	aditzLaguntzilea := ""
	row := db.QueryRow(`select aditz_laguntzilea from aditz_lagunak where nor = $1 and nori = $2 and nork = $3  and denbora = $4;`, guess.Nor, guess.Nori, guess.Nork, guess.Denbora)

	switch err := row.Scan(&aditzLaguntzilea); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(aditzLaguntzilea)
	default:
		panic(err)
	}
	return aditzLaguntzilea, err
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func openConnection() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	psqlconn := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable user=%s password=%s", host, port, dbName, user, password)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	return db, err
}

func RandomGaldera() NorNoriNork { // todo be sure there is a valid answer
	db, err := openConnection()
	defer db.Close()

	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")
	var randomNor Nor
	var randomNori Nori
	var randomNork Nork
	row := db.QueryRow(`select nor, nori, nork from aditz_lagunak where denbora = $1 order by random () limit 1;`, "Orainaldia")
	CheckError(err)

	switch err := row.Scan(&randomNor, &randomNori, &randomNork); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(randomNor, randomNori, randomNork)
	default:
		panic(err)
	}
	fmt.Println(randomNor, randomNori, randomNork)
	return NorNoriNork{randomNor, randomNori, randomNork}

}

func RandomNorNori() NorNori { // todo be sure there is a valid answer
	db, err := openConnection()
	defer db.Close()
	CheckError(err)

	var randomNor Nor
	var randomNori Nori
	row := db.QueryRow(`select nor, nori from nor_nori where denbora = $1 order by random () limit 1;`, "Orainaldia")
	CheckError(err)

	switch err := row.Scan(&randomNor, &randomNori); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(randomNor, randomNori)
	default:
		panic(err)
	}
	fmt.Println(randomNor, randomNori)
	return NorNori{randomNor, randomNori}

}
