package poorga

//StringMatcher 用來測試基因演算法匹配輸入的內容
type StringMatcher struct {
	target string
}

//SetTarget 用於設置目標字串
func (matcher *StringMatcher) SetTarget(target string) {
	matcher.target = target
}

func (matcher StringMatcher) initChromosome(chromosome Chromosome) {

}
func (matcher StringMatcher) fitness(chromosome []Chromosome) {

}
