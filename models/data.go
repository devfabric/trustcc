package models

const (
	TAB_COLID_FSEQ = "TAB-COLID-FSEQ"
)

type ItemKV struct {
	TabName string
	TxID    string
	ColSeq  string
	Val     MetaValue
}

type MetaValue struct {
	FieldVal string
	Version  int64
	Status   int
	Owner    string
}

/*
Key(Tab-TxID-1) => Val(val1,1,0,liuhy)
Key(Tab-TxID-2) => Val(val2,1,0,liuhy)
Key(Tab-TxID-3) => Val(val3,1,0,liuhy)


*/
