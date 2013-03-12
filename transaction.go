package goci

/*
#include <oci.h>
#include <stdlib.h>
#include <string.h>

#cgo CFLAGS: -I/home/oracle/app/oracle/product/11.2.0/client_1/rdbms/public
#cgo LDFLAGS: -lclntsh -L/home/oracle/app/oracle/product/11.2.0/client_1/lib

*/
import "C"

type transaction struct {
	conn *connection
}

func (tx *transaction) Commit() error {
	if err := tx.conn.exec("COMMIT"); err != nil {
		return err
	}
	return nil
}

func (tx *transaction) Rollback() error {
	if err := tx.conn.exec("ROLLBACK"); err != nil {
		return err
	}
	return nil
}
