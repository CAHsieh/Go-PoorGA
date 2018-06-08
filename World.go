package poorga

//World 是染色體生存的環境
type World struct {
	chromosome    []Chromosome //每個世代存活的染色體
	randomNum     int          //每一世代中會隨機選擇及存活下來的染色體數量
	generationNum int          //每個世代存活的染色體個數
	iterationNum  int          //停止條件1：總迭代次數
	optimizeTims  int          //停止條件2：局部最佳迭代次數
	goodEnough    float64      //停止條件3：足夠停止的準確度
}
