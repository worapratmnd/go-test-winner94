package model

type TransactionDb struct {
	Id         int     `db:"id"`
	Input      *string `db:"input"`
	Type       *string `db:"type"`
	S1         *int    `db:"s1"`
	S2         *int    `db:"s2"`
	S3         *int    `db:"s3"`
	Total      *string `db:"total"`
	Status     *int    `db:"status"`
	CreateBy   *string `db:"createBy"`
	CreateDttm *string `db:"createDttm"`
	UpdateBy   *string `db:"updateBy"`
	UpdateDttm *string `db:"updateDttm"`
	Tables     *int    `db:"tables"`
	RoundName  *string `db:"roundName"`
	RoundNo    *int    `db:"roundNo"`
	Result     *string `db:"result"`
}
