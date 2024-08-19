package database

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

// Account is used to keep track of the
type Account struct {
	AccountID           int
	AccountUUID         uuid.UUID
	Integration         string
	IntegrationBlob     []byte
	LastTransactionID   string
	LastTransactionDate time.Time
}

// SelectAccounts fetches all the accounts from the database.
func SelectAccounts(db *sql.DB) (*[]Account, error) {
	rows, err := db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	var xa *[]Account
	for rows.Next() {
		// scan reads the raw bytes into xb, however, the value is not copied
		// but still owned by the database and must be used before the following
		// Next call. Otherwise, the number of ...dest need to match the number
		// of columns from the query. Or, we can just read it directly into a
		// struct where the names match the columns.
		a := Account{}
		err = rows.Scan(&a.AccountID, &a.AccountUUID, &a.Integration, &a.IntegrationBlob, &a.LastTransactionID, &a.LastTransactionDate)
		if err != nil {
			log.Println("select accounts scan", err)
			return nil, err
		}
		*xa = append(*xa, a)
	}

	return xa, nil
}

// SelectAccount fetches a single account entry from the database based on the
// account's UUID.
func SelectAccount(db *sql.DB, UUID uuid.UUID) (*Account, error) {
	row := db.QueryRow(`SELECT * FROM account WHERE account_uuid=?`, UUID.String())
	if row == nil {
		return nil, fmt.Errorf("account not found")
	}

	a := Account{}
	err := row.Scan(a.AccountID, a.AccountUUID, a.Integration, a.IntegrationBlob, a.LastTransactionID, a.LastTransactionDate)
	if err != nil {
		log.Println("select account scan:", err)
		return nil, err
	}

	return &a, nil
}

// CreateTable creates the account table in the database.
func CreateTable(db *sql.DB) {
	// in Go best is to have not null with most columns, as Go cannot convert
	// nil to other types such as a string, int, float or time. Go does not
	// resolve a nil into to zero-type of a value.
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS account (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		account_uuid TEXT NOT NULL UNIQUE,
		integration TEXT NOT NULL,
		integration_blob BLOB NOT NULL,
		last_transaction_id TEXT NOT NULL,
		last_transaction_date TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatalln("create", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("stmt exec", err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatalln("stmt close", err)
	}
}

// Insert creates a new account entry.
func Insert(db *sql.DB, account Account) error {
	stmt, err := db.Prepare(`
	INSERT INTO account (account_uuid, integration, integration_blob, last_transaction_id, last_transaction_date) VALUES (?, ?, ?, '', ?)
	`)
	if err != nil {
		log.Println("insert account stmt", err)
		return err
	}
	// now we can do the insert
	_, err = stmt.Exec(account.AccountUUID.String(), account.Integration, account.IntegrationBlob, time.Now())
	if err != nil {
		log.Println("insert account exec", err)
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println("insert account close stmt", err)
		return err
	}
	return nil
}

// UpdateLastTransaction updates a specific account entry in the account table.
func UpdateLastTransaction(db *sql.DB, account Account) error {
	stmt, err := db.Prepare(`
	UPDATE account SET last_transaction_id=?, last_transaction_date=? WHERE account_id=?
	`)
	if err != nil {
		log.Println("update account transaction stmt", err)
		return err
	}

	_, err = stmt.Exec(account.LastTransactionID, account.LastTransactionDate, account.AccountID)
	if err != nil {
		log.Println("update account transaction", err)
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println("update stmt close", err)
		return err
	}
	return nil
}
