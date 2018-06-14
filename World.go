package poorga

import (
	"fmt"
	"math/rand"
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
}

//SetIsPrint 用來控制是否輸出GA過程
func (world *World) SetIsPrint(isPrint bool) {
	world.isPrint = isPrint
}

//StartWorld 開始執行基因演算法
func (world *World) StartWorld() {
	optTimes := 0
	lastFitness := float64(0)

	//initial
	world.chromosome = make([]Chromosome, world.generationNum)
	for i := 0; i < world.generationNum; i++ {
		world.customMethod.InitChromosome(&world.chromosome[i])
		world.customMethod.Fitness(&world.chromosome[i])
	}

	for i := 0; i < world.iterationNum; i++ {
		if optTimes >= world.optimizeTimes {
			break
		}

		crossChr := world.crossover() // 交配
		mutatedChr := world.mutate()  // 突變

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
	println("end!!")
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

func (world World) selection(crossChr []Chromosome, mutatedChr []Chromosome) []Chromosome {

	remainingPlace := world.generationNum - world.randomNum
	survivedChr := make([]Chromosome, world.generationNum)

	oriIdx := 0
	crxIdx := 0
	mutIdx := 0

	for i := 0; i < remainingPlace; i++ {
		if crossChr[crxIdx].fitness >= world.chromosome[oriIdx].fitness && crossChr[crxIdx].fitness >= mutatedChr[mutIdx].fitness {
			survivedChr[i] = crossChr[crxIdx]
			crxIdx++
		} else if mutatedChr[mutIdx].fitness >= world.chromosome[oriIdx].fitness && mutatedChr[mutIdx].fitness >= crossChr[crxIdx].fitness {
			survivedChr[i] = mutatedChr[mutIdx]
			mutIdx++
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
