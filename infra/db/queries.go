package db

const (
	QueryUser = `SELECT r.id,
			         r.name,
			         r.email,
			         r.document,
			         r.pixKey,
			         pkt.name pix_type,
			         case r.status
			             when 0 then 'draft'
			             when 1 then 'active' END AS status 
				 FROM receiver r
					      LEFT JOIN pix_key_type pkt on r.pixKeyType = pkt.id
				 WHERE r.name LIKE $1
				 OR r.document LIKE $1
				 OR r.email LIKE $1
				 OR pkt.name LIKE $1
				 ORDER BY r.id
			     LIMIT $2;`

	QueryUserByID = `SELECT r.id,
			         r.name,
			         r.email,
			         r.document,
			         r.pixKey,
			         pkt.name pix_type,
			         case r.status
			             when 0 then 'draft'
			             when 1 then 'active' END AS status
					 FROM receiver r
					          LEFT JOIN pix_key_type pkt on r.pixKeyType = pkt.id
					 WHERE r.id = $1 LIMIT 1`

	QueryPixTypeByName = `SELECT id FROM pix_key_type WHERE name = $1 LIMIT 1`

	QueryListOfReceivers = `SELECT r.id,
								   r.name,
								   r.document,
								   case r.status
									   WHEN 0 THEN 'draft'
									   when 1 then 'active' END AS status
								FROM receiver r
										 LEFT JOIN pix_key_type pkt on r.pixKeyType = pkt.id
								ORDER BY r.id
								OFFSET $2
								LIMIT $1`

	InsertNewReceiverQuery = `INSERT INTO receiver (id, name, email, document, pixKey, pixKeyType, status)
							  VALUES ($1, $2, $3, $4, $5, $6, $7)`

	UpdateReceiverByID = `UPDATE receiver
					  SET name       = $1,
					      email      = $2,
					      document   = $3,
					      pixKey     = $4,
					      pixKeyType = $5,
					      status     = $6
					  WHERE receiver.id = $7`

	UpdateReceiverEmailByID = `UPDATE receiver
							   SET email = $1
							   WHERE receiver.id = $2`

	DelteReceiversByID = `DELETE FROM receiver WHERE id IN (?)`
)
