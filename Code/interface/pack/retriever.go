package abc

type MyRetriever struct {
	Something string
}

func (r *MyRetriever) Get(aaa string) string {
	return r.Something + aaa
}

// func (r *MyRetriever) Post(aaa string) string {
// 	return r.Something + aaa
// }
