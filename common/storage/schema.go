package storage

func (h *Handler) ensureBlocksTable() (err error) {
	_, err = h.connection.Exec(`
		create table if not exists blocks (
			version UInt16 default 0,
			number UInt64,
			hash String,
			prevHash String,
			signature String,
			nextSigPubKey String,
			recordsNumbers Array(UInt64),
			generated DateTime default now()
		) engine = MergeTree()
			primary key number
			order by number
			partition by number	
    `)

	return
}

func (h *Handler) ensureRecordsTable() (err error) {
	_, err = h.connection.Exec(`
		create table if not exists records (
			number UInt64,
			generated DateTime default now(),
			type UInt16,
			data String
		) engine = MergeTree()
			primary key number
			order by number
			partition by number;	
    `)

	_, err = h.connection.Exec(`
		CREATE TABLE records_buffer AS records ENGINE = Buffer(default, records, 16, 10, 120, 1, 1000, 1, 10485760) 
	`)

	return
}
