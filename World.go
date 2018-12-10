package poorga

import (
	"fmt"
	"math/rand"
	"runtime"
)

//World 是染色體生存的環境
type World struct {
	chromosome    []Chromosome //每個世代存活的染色體
	randomNum     int          //每一世代中會隨機選擇及存活下來的染色體數量
	generationNum int          //每個世代存活的染色體個數
	iterationNum  int          //停止條件1：總迭代次數
	optimizeTimes int          //停止條件2：局部最佳迭代次數
	goodEnough    float64      //停止條件3：足夠停止的準確度
	customMethod  Custom       //需要客製化方法
	isPrint       bool         //是否顯示當前狀態
	cpuNums       int          //CPU使用數量上限
}

//Initial 用於設置相關參數
func (world *World) Initial(generationNum int, randomNum int,
	iterationNum int, optimizeTimes int, goodEnough float64, customMethod Custom) {

	if randomNum > generationNum {
		fmt.Println("randomNum can not larger than generationNum")
		panic(world)
	}

	world.generationNum = generationNum
	world.randomNum = randomNum
	world.iterationNum = iterationNum
	world.optimizeTimes = optimizeTimes
	world.goodEnough = goodEnough
	world.customMethod = customMethod
	world.cpuNums = 1
}

//SetIsPrint 用來控制是否在GA過程中調用Custom介面中的Print
func (world *World) SetIsPrint(isPrint bool) {
	world.isPrint = isPrint
}

//SetMAXCPUs sets the maximum number of CPUs that can be executing
func (world *World) SetMAXCPUs(cpuNums int) {
	world.cpuNums = cpuNums
}

//StartWorld 開始執行基因演算法
func (world *World) StartWorld() {

	//使用多核心
	runtime.GOMAXPROCS(world.cpuNums)

	optTimes := 0
	lastFitness := float64(0)

	//initial
	world.chromosome = make([]Chromosome, world.generationNum)
	for i := 0; i < world.generationNum; i++ {
		world.customMethod.InitChromosome(&world.chromosome[i])
		world.customMethod.Fitness(&world.chromosome[i])
	}

	i := 0
	for ; i < world.iterationNum; i++ {
		if optTimes >= world.optimizeTimes {
			break
		}

		crossChannel := make(chan Chromosome, world.generationNum)  //交配
		mutateChannel := make(chan Chromosome, world.generationNum) //突變

		crossChr := make([]Chromosome, world.generationNum)
		mutatedChr := make([]Chromosome, world.generationNum)

		for i := 0; i < world.generationNum; i++ {
			go world.crossoverChannel(crossChannel)
			go world.mutateChannel(i, mutateChannel)
		}
		for i := 0; i < world.generationNum; i++ {
			crossChr[i] = <-crossChannel
			mutatedChr[i] = <-mutateChannel
		}

		world.chromosome = world.selection(crossChr, mutatedChr) // 選擇

		if world.isPrint {
			world.customMethod.Print(i, world.chromosome)
		}

		if world.chromosome[0].fitness >= world.goodEnough {
			// good enough
			break
		}

		if lastFitness == world.chromosome[0].fitness {
			optTimes++
		} else {
			lastFitness = world.chromosome[0].fitness
			optTimes = 0
		}

	}
	world.customMethod.PrintResult(world.chromosome)
	fmt.Printf("end at iteration %d !!", i)
}

/**
 * 產生交配後的下一代
 */
func (world World) crossover() []Chromosome {

	crossChr := make([]Chromosome, world.generationNum)

	for i := 0; i < world.generationNum; i++ {
		firstIdx := rand.Int() % world.generationNum
		secIdx := rand.Int() % world.generationNum
		for secIdx == firstIdx {
			secIdx = rand.Int() % world.generationNum
		}
		crossChr[i] = world.chromosome[firstIdx].crossover(world.chromosome[secIdx], rand.Float32()*0.5)
		world.customMethod.Fitness(&crossChr[i])
	}

	return crossChr
}

/**
 * 產生交配後的下一代
 * Concurrency Version
 */
func (world World) crossoverChannel(channel chan Chromosome) {

	var crossChr Chromosome

	firstIdx := rand.Int() % world.generationNum
	secIdx := rand.Int() % world.generationNum
	for secIdx == firstIdx {
		secIdx = rand.Int() % world.generationNum
	}
	crossChr = world.chromosome[firstIdx].crossover(world.chromosome[secIdx], rand.Float32()*0.5)
	world.customMethod.Fitness(&crossChr)

	channel <- crossChr
}

/**
 * 產生突變後的下一代
 */
func (world World) mutate() []Chromosome {

	mutatedChr := make([]Chromosome, world.generationNum)

	for i := 0; i < world.generationNum; i++ {
		mutatedChr[i] = world.chromosome[i].mutate(rand.Float32() * 0.7)
		world.customMethod.Fitness(&mutatedChr[i])
	}

	return mutatedChr
}

/**
 * 產生突變後的下一代
 * Concurrency Version
 */
func (world World) mutateChannel(index int, channel chan Chromosome) {

	var mutatedChr Chromosome

	mutatedChr = world.chromosome[index].mutate(rand.Float32() * 0.7)
	world.customMethod.Fitness(&mutatedChr)

	channel <- mutatedChr
}

/**
 * 根據配適度選擇下一世代存活下來的染色體群，
 * 若配適度相同則優先序為突變 > 交配 > 本世代染色體
 */
func (world World) selection(crossChr []Chromosome, mutatedChr []Chromosome) []Chromosome {

	remainingPlace := world.generationNum - world.randomNum
	survivedChr := make([]Chromosome, world.generationNum)

	oriIdx := 0
	crxIdx := 0
	mutIdx := 0

	for i := 0; i < remainingPlace; i++ {
		if mutatedChr[mutIdx].fitness >= world.chromosome[oriIdx].fitness && mutatedChr[mutIdx].fitness >= crossChr[crxIdx].fitness {
			survivedChr[i] = mutatedChr[mutIdx]
			mutIdx++
		} else if crossChr[crxIdx].fitness >= world.chromosome[oriIdx].fitness && crossChr[crxIdx].fitness >= mutatedChr[mutIdx].fitness {
			survivedChr[i] = crossChr[crxIdx]
			crxIdx++
		} else if world.chromosome[oriIdx].fitness >= crossChr[crxIdx].fitness && world.chromosome[oriIdx].fitness >= mutatedChr[mutIdx].fitness {
			survivedChr[i] = world.chromosome[oriIdx]
			oriIdx++
		}
	}

	for i := 0; i < world.randomNum; i++ {
		which := rand.Int() % 3
		switch which {
		case 0:
			survivedChr[i+remainingPlace] = world.chromosome[(rand.Int()%(world.generationNum-oriIdx))+oriIdx]
		case 1:
			survivedChr[i+remainingPlace] = crossChr[(rand.Int()%(world.generationNum-crxIdx))+crxIdx]
		case 2:
			survivedChr[i+remainingPlace] = mutatedChr[(rand.Int()%(world.generationNum-mutIdx))+mutIdx]
		}
	}

	return survivedChr
}
