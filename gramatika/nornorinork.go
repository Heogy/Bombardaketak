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

type NorNoriNork struct {
	RandomNor  Nor  `json:"nor"`
	RandomNori Nori `json:"nori"`
	RandomNork Nork `json:"nork"`
}

func Verify(guess Guess) (bool, string, error) {
	erantzuna, err := resolve(guess)
	if err != nil {
		return false, "", err
	}
	return strings.EqualFold(strings.TrimSpace(erantzuna), strings.TrimSpace(guess.Erantzuna)), erantzuna, nil
}
func resolve(guess Guess) (string, error) {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	psqlconn := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable user=%s password=%s", host, port, dbName, user, password)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	aditzLaguntzilea := ""
	row := db.QueryRow(`select aditz_laguntzilea from aditz_lagunak where nor = $1 and nori = $2 and nork = $3  and denbora = $4;`, guess.Nor, guess.Nori, guess.Nork, guess.Denbora)
	CheckError(err)

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

// func resolve(guess Guess) (string, error) {
// 	// println(guess.Denbora)
// 	if guess.Denbora == Orainaldia {
// 		// println(guess.Nor)

// 		// NOR NORI NORK
// 		if guess.Nor == Ni || guess.Nor == Zu || guess.Nor == Hura {
// 			if guess.Nori == Niri {
// 				if guess.Nork == Nik {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == Hark {
// 					return "dit", nil
// 				} else if guess.Nork == Guk {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == Zuk {
// 					return "didazu", nil
// 				} else if guess.Nork == NorkZuek {
// 					return "didazue", nil
// 				} else if guess.Nork == NorkHaiek {
// 					return "didate", nil
// 				}
// 			} else if guess.Nori == Hari {
// 				if guess.Nork == Nik {
// 					return "diot", nil
// 				} else if guess.Nork == Hark {
// 					return "dio", nil
// 				} else if guess.Nork == Guk {
// 					return "diogu", nil
// 				} else if guess.Nork == Zuk {
// 					return "diozu", nil
// 				} else if guess.Nork == NorkZuek {
// 					return "diozue", nil
// 				} else if guess.Nork == NorkHaiek {
// 					return "diote", nil
// 				}
// 			} else if guess.Nori == Guri {
// 				if guess.Nork == Nik {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == Hark {
// 					return "digu", nil
// 				} else if guess.Nork == Guk {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == Zuk {
// 					return "diguzu", nil
// 				} else if guess.Nork == NorkZuek {
// 					return "diguzue", nil
// 				} else if guess.Nork == NorkHaiek {
// 					return "digute", nil
// 				}
// 			} else if guess.Nori == Zuri {
// 				if guess.Nork == Nik {
// 					return "dizut", nil
// 				} else if guess.Nork == Hark {
// 					return "dizu", nil
// 				} else if guess.Nork == Guk {
// 					return "dizugu", nil
// 				} else if guess.Nork == Zuk {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == NorkZuek {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == NorkHaiek {
// 					return "dizute", nil
// 				}
// 			} else if guess.Nori == Zuei {
// 				if guess.Nork == Nik {
// 					return "dizuet", nil
// 				} else if guess.Nork == Hark {
// 					return "dizue", nil
// 				} else if guess.Nork == Guk {
// 					return "dizuegu", nil
// 				} else if guess.Nork == Zuk {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == NorkZuek {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == NorkHaiek {
// 					return "dizuete", nil
// 				}
// 			} else if guess.Nori == Haiei {
// 				if guess.Nork == Nik {
// 					return "diet", nil
// 				} else if guess.Nork == Hark {
// 					return "die", nil
// 				} else if guess.Nork == Guk {
// 					return "diegu", nil
// 				} else if guess.Nork == Zuk {
// 					return "diezu", nil
// 				} else if guess.Nork == NorkZuek {
// 					return "diezue", nil
// 				} else if guess.Nork == NorkHaiek {
// 					return "diete", nil
// 				}
// 			}
// 		} else if guess.Nor == Gu || guess.Nor == Zuek || guess.Nor == Haiek {
// 			if guess.Nori == Niri {
// 				if guess.Nork == Nik {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == Hark {
// 					return "dizkit", nil
// 				} else if guess.Nork == Guk {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == Zuk {
// 					return "dizkidazu", errors.New("Not implemented")
// 				} else if guess.Nork == NorkZuek {
// 					return "dizkidazue", errors.New("Not implemented")
// 				} else if guess.Nork == NorkHaiek {
// 					return "dizkidate", errors.New("Not implemented")
// 				}
// 			} else if guess.Nori == Hari {
// 				if guess.Nork == Nik {
// 					return "dizkiot", nil
// 				} else if guess.Nork == Hark {
// 					return "dizkio", nil
// 				} else if guess.Nork == Guk {
// 					return "dizkiogu", nil
// 				} else if guess.Nork == Zuk {
// 					return "dizkiozu", nil
// 				} else if guess.Nork == NorkZuek {
// 					return "dizkiozue", nil
// 				} else if guess.Nork == NorkHaiek {
// 					return "dizkiote", nil
// 				}
// 			} else if guess.Nori == Guri {
// 				if guess.Nork == Nik {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == Hark {
// 					return "dizkigu", nil
// 				} else if guess.Nork == Guk {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == Zuk {
// 					return "dizkiguzu", nil
// 				} else if guess.Nork == NorkZuek {
// 					return "dizkiguzue", nil
// 				} else if guess.Nork == NorkHaiek {
// 					return "dizkigute", nil
// 				}
// 			} else if guess.Nori == Zuri {
// 				if guess.Nork == Nik {
// 					return "dizkizut", nil
// 				} else if guess.Nork == Hark {
// 					return "dizkizu", nil
// 				} else if guess.Nork == Guk {
// 					return "dizkizugu", nil
// 				} else if guess.Nork == Zuk {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == NorkZuek {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == NorkHaiek {
// 					return "dizkizute", nil
// 				}
// 			} else if guess.Nori == Zuei {
// 				if guess.Nork == Nik {
// 					return "dizkizuet", nil
// 				} else if guess.Nork == Hark {
// 					return "dizkizue", nil
// 				} else if guess.Nork == Guk {
// 					return "dizkizuegu", nil
// 				} else if guess.Nork == Zuk {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == NorkZuek {
// 					return "", errors.New("Not implemented")
// 				} else if guess.Nork == NorkHaiek {
// 					return "dizkizuete", nil
// 				}
// 			} else if guess.Nori == Haiei {
// 				if guess.Nork == Nik {
// 					return "dizkiet", nil
// 				} else if guess.Nork == Hark {
// 					return "dizkie", nil
// 				} else if guess.Nork == Guk {
// 					return "dizkiegu", nil
// 				} else if guess.Nork == Zuk {
// 					return "dizkiezu", nil
// 				} else if guess.Nork == NorkZuek {
// 					return "dizkiezue", nil
// 				} else if guess.Nork == NorkHaiek {
// 					return "dizkiete", nil
// 				}
// 			}
// 		}

// 		// NOR NORI
// 		if guess.Nor == Hura {
// 			if guess.Nori == Niri {
// 				return "zait", nil
// 			} else if guess.Nori == Hari {
// 				return "zaio", nil
// 			} else if guess.Nori == Guri {
// 				return "zaigu", nil
// 			} else if guess.Nori == Zuri {
// 				return "zaizu", nil
// 			} else if guess.Nori == Zuei {
// 				return "zaizue", nil
// 			} else if guess.Nori == Haiei {
// 				return "zaie", nil
// 			}
// 		} else if guess.Nor == Haiek {
// 			if guess.Nori == Niri {
// 				return "zaizkit", nil
// 			} else if guess.Nori == Hari {
// 				return "zaizkio", nil
// 			} else if guess.Nori == Guri {
// 				return "zaizkigu", nil
// 			} else if guess.Nori == Zuri {
// 				return "zaizkizu", nil
// 			} else if guess.Nori == Zuei {
// 				return "zaizkizue", nil
// 			} else if guess.Nori == Haiei {
// 				return "zaizkie", nil
// 			}
// 		}

// 		// NOR NORK
// 		if guess.Nor == Ni {
// 			if guess.Nork == Nik {
// 				return "", errors.New("Not implemented")
// 			} else if guess.Nork == Hark {
// 				return "nau", nil
// 			} else if guess.Nork == Guk {
// 				return "", errors.New("Not implemented")
// 			} else if guess.Nork == Zuk {
// 				return "nauzu", nil
// 			} else if guess.Nork == NorkZuek {
// 				return "nauzue", nil
// 			} else if guess.Nork == NorkHaiek {
// 				return "naute", nil
// 			}
// 		} else if guess.Nor == Hura {
// 			if guess.Nork == Nik {
// 				return "dut", nil
// 			} else if guess.Nork == Hark {
// 				return "du", nil
// 			} else if guess.Nork == Guk {
// 				return "dugu", nil
// 			} else if guess.Nork == Zuk {
// 				return "duzu", nil
// 			} else if guess.Nork == NorkZuek {
// 				return "duzue", nil
// 			} else if guess.Nork == NorkHaiek {
// 				return "dute", nil
// 			}
// 		} else if guess.Nor == Gu {
// 			if guess.Nork == Nik {
// 				return "", errors.New("Not implemented")
// 			} else if guess.Nork == Hark {
// 				return "gaitu", nil
// 			} else if guess.Nork == Guk {
// 				return "", errors.New("Not implemented")
// 			} else if guess.Nork == Zuk {
// 				return "gaituzu", nil
// 			} else if guess.Nork == NorkZuek {
// 				return "gaituzue", nil
// 			} else if guess.Nork == NorkHaiek {
// 				return "gaitute", nil
// 			}
// 		} else if guess.Nor == Zu {
// 			if guess.Nork == Nik {
// 				return "zaitut", nil
// 			} else if guess.Nork == Hark {
// 				return "zaitu", nil
// 			} else if guess.Nork == Guk {
// 				return "zaitugu", nil
// 			} else if guess.Nork == Zuk {
// 				return "", errors.New("Not implemented")
// 			} else if guess.Nork == NorkZuek {
// 				return "", errors.New("Not implemented")
// 			} else if guess.Nork == NorkHaiek {
// 				return "zaitute", nil
// 			}
// 		} else if guess.Nor == Zuek {
// 			if guess.Nork == Nik {
// 				return "zaituztet", nil
// 			} else if guess.Nork == Hark {
// 				return "zaituzte", nil
// 			} else if guess.Nork == Guk {
// 				return "zaituztegu", nil
// 			} else if guess.Nork == Zuk {
// 				return "", errors.New("Not implemented")
// 			} else if guess.Nork == NorkZuek {
// 				return "", errors.New("Not implemented")
// 			} else if guess.Nork == NorkHaiek {
// 				return "zaituztete", nil
// 			}
// 		} else if guess.Nor == Haiek {
// 			if guess.Nork == Nik {
// 				return "ditut", nil
// 			} else if guess.Nork == Hark {
// 				return "ditu", nil
// 			} else if guess.Nork == Guk {
// 				return "ditugu", nil
// 			} else if guess.Nork == Zuk {
// 				return "dituzu", nil
// 			} else if guess.Nork == NorkZuek {
// 				return "dituzue", nil
// 			} else if guess.Nork == NorkHaiek {
// 				return "dituzte", nil
// 			}
// 		}

// 	}
// 	return "", errors.New("Not implemented")
// }

func RandomGaldera() NorNoriNork { // todo be sure there is a valid answer
	// open database
	// close database
	// check db
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	psqlconn := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable user=%s password=%s", host, port, dbName, user, password)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

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
