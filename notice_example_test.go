// +build go1.10

package pq_test

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func ExampleConnectorWithNoticeHandler() {
	name := ""
	// Base connector to wrap
	connector, err := pq.NewConnector(name)
	if err != nil {
		log.Fatal(err)
	}
	// Wrap the connector to simply print out the message
	connector = pq.ConnectorWithNoticeHandler(connector, func(notice *pq.Error) {
		fmt.Println("Notice sent: " + notice.Message)
	}).(*pq.Connector)
	db := sql.OpenDB(connector)
	defer db.Close()
	// Raise a notice
	sql := "DO language plpgsql $$ BEGIN RAISE NOTICE 'test notice'; END $$"
	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
	// Output:
	// Notice sent: test notice
}
