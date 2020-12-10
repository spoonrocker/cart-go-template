package commons

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func Connect() {
	db, err := bolt.Open(os.Getenv("DATABASE"), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("carts"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("cart_items"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
