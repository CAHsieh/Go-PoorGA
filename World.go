package poorga

//World 是染色體生存的環境
type World struct {
	chromosome    []Chromosome //每個世代存活的染色體
	randomNum     int          //每一世代中會隨機選擇及存活下來的染色體數量
	generationNum int          //每個世代存活的染色體個數
	iterationNum  int          //停止條件1：總迭代次數
	optimizeTimes int          //停止條件2：局部最佳迭代次數
	goodEnough    float64      //停止條件3：足夠停止的準確度
	customMethod  Custom       //需要客製化方法
}

//Custom 用來導入需要客製化的方法
type Custom interface {
	initChromosome(chromosome Chromosome)
	fitness(chromosome []Chromosome)
}

//Initial 用於設置相關參數
func (world *World) Initial(generationNum int, randomNum int,
	iterationNum int, optimizeTimes int, goodEnough float64, customMethod Custom) {

	world.generationNum = generationNum
	world.randomNum = randomNum
	world.iterationNum = iterationNum
	world.optimizeTimes = optimizeTimes
	world.goodEnough = goodEnough
	world.customMethod = customMethod

	world.chromosome = make([]Chromosome, generationNum)
	for i := 0; i < generationNum; i++ {
		customMethod.initChromosome(world.chromosome[i])
		customMethod.fitness(world.chromosome)
	}
}

//StartWorld 開始執行基因演算法
func (world *World) StartWorld() {
	optTimes := 0
	for i := 0; i < world.iterationNum; i++ {
		if optTimes >= world.optimizeTimes {
			break
		}

	}
	println("start!!")
}
