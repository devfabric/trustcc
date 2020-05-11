package models

type MetaTab struct {
	TabName string //表名
	Owner   string //元数据
	Cols    []struct {
		colSeq  int    //列序
		colName string //列名
	}
	Inds    map[string][]string //索引数据
	Version int64
}
