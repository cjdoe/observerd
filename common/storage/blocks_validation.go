package storage

//// LastBlockNumber returns number of the last block written to the storage.
//// In case if there is no any block in storage - returns ErrNoBlocks.
//func (h *Handler) LastBlockNumber() (number uint64, err error) {
//	rows, err = h.connection.Query("SELECT max(number) FROM blocks")
//	defer ensureClose(rows)
//	if err != nil {
//		return
//	}
//
//	for rows.Next() {
//		err = rows.Scan(&number)
//		if err != nil {
//		    return
//		}
//	}
//
//	if number == 0 {
//		number, err = h.BlocksCount()
//		if err != nil {
//		    return 0, err
//		}
//
//		if number == 0 {
//			err = e.ErrNoBlocks
//		}
//	}
//
//	return
//}

// BlocksCount returns amount of blocks present in storage.
func (h *Handler) BlocksCount() (amount uint64, err error) {
	rows, err = h.connection.Query("SELECT count(number) FROM blocks")
	defer ensureClose(rows)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&amount)
		return
	}
	return
}
