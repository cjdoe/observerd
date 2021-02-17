package database

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/vTCP-Foundation/observerd/core/ec"
	"io/ioutil"
)

const (
	filename = "assets/schema.sql"

	// [security]
	// Schema definition SQL is provided as a standalone SQL script.
	// To prevent various attacks through replacing or modifying of the content of this file -
	// hash based validation is used (required script hash is hardcoded into the executable binary).
	requiredSchemaHashHex = "e880e90f54a5f5c7bda6c97b65014a9a7a87a13f093a4ace72d88e8543b68959"
)

// EnsureSchema executes provided schema definition SQL file.
// By default it ensures all required tables / indexes / constraints / triggers are present.
// In case if some table (or index, or ...) is absent - it would be recreated.
// Does not replaces tables in cae if them are present (to not to drop data occasionally).
func EnsureSchema() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("can't ensure database schema: %w", err)
		}
	}()

	script, err := loadSQLScript()
	if err != nil {
		return
	}

	err = validateScriptHash(script)
	if err != nil {
		return
	}

	return executeSQLScript(script)
}

// loadSQLScript returns string representation of the schema definition SQL,
// loaded from the standalone file.
func loadSQLScript() (script string, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		msg := fmt.Sprint("can't read SQL schema definition script: ", filename)
		err = fmt.Errorf(msg+": %w", err)
		return
	}

	script = string(data)
	return
}

// validateScriptHash calculates hash of the provided SQL script and compares it with te expected hash.
// This validation step is required to prevent original SQL instructions replacement.
func validateScriptHash(script string) (err error) {
	hash := sha256.Sum256([]byte(script))
	providedHashHex := hex.EncodeToString(hash[:])

	if providedHashHex != requiredSchemaHashHex {
		msg := fmt.Sprint("hash of the SQL schema definition script (",
			filename, ") does not correspond to the expected one,"+
				" provided hash: '", providedHashHex, "',",
			" required hash: '", requiredSchemaHashHex, "'")

		err = fmt.Errorf(msg+" (%w)", ec.ErrValidation)
		return
	}

	return
}

// executeSQLScript executes schema definition script.
func executeSQLScript(script string) (err error) {
	_, err = db.Exec(context.Background(), script)
	if err != nil {
		return
	}

	return
}
