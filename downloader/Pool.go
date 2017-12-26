package downloader

//网页下载器池的接口类型
type PageDownLoaderPool interface {
	Take() (PageDownLoader, error)              //从池中获取一个网页下载器
	Return(pageDownLoader PageDownLoader) error //把一个网页下载器归还给池
	Total() uint32                              //网页下载器池的总容量
	Userd() uint32                              //正在被使用的网页下载器数量
}
