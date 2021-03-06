package analyzer

//分析器池的接口类型
type AnalyzerPool interface {
	Take() (Analyzer, error)        //从池中取出一个分析器
	Return(analyzer Analyzer) error //把一个分析器归还给池
	Total() uint32                  //获取分析器池的总容量
	Used() uint32                   // 正在被使用的分析器的数量
}
