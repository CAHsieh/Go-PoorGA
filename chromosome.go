package poorga

import (
	"math/rand"
)

//Chromosome for genetic algorithm
type Chromosome struct {
	body    []int
	length  int
	fitness float64
}

//Init 用於初始化染色體，必須在Custom中的initChromosome內調用
func (chrom *Chromosome) Init(len int) {
	chrom.fitness = -1
	chrom.length = len
	chrom.body = make([]int, len)
	for i := 0; i < len; i++ {
		chrom.body[i] = rand.Int() % 2
	}
}

//SetFitness 用於設置配適度
func (chrom *Chromosome) SetFitness(fitness float64) {
	chrom.fitness = fitness
}

//GetFitness 用於取得配適度
func (chrom Chromosome) GetFitness() float64 {
	return chrom.fitness
}

//GetBody 用於取得染色體內容
func (chrom Chromosome) GetBody() []int {
	return chrom.body
}

func (chrom *Chromosome) crossover(secChrom Chromosome, crossRate float32) Chromosome {
	var newChrom Chromosome
	len := chrom.length
	newChrom.length = len
	newChrom.body = make([]int, len)
	newChrom.fitness = -1

	motherNum := int(float32(len) * crossRate)
	usedMotherGenetic := make([]bool, len)

	//決定哪些基因來自母親
	for i := 0; i < motherNum; i++ {
		var idx int
		for idx = rand.Int() % len; usedMotherGenetic[idx]; idx = rand.Int() % len {
		}
		usedMotherGenetic[idx] = true
	}

	//產生新基因
	for i := 0; i < len; i++ {
		if usedMotherGenetic[i] {
			newChrom.body[i] = chrom.body[i]
		} else {
			newChrom.body[i] = secChrom.body[i]
		}
	}

	return newChrom
}

func (chrom *Chromosome) mutate(mutationRate float32) Chromosome {
	var newChrom Chromosome
	len := chrom.length
	newChrom.length = len
	newChrom.body = make([]int, len)
	newChrom.fitness = -1

	mutationNum := int(float32(len) * mutationRate)
	mutationGenetic := make([]bool, len)

	//決定哪些基因會突變
	for i := 0; i < mutationNum; i++ {
		var idx int
		for idx = rand.Int() % len; mutationGenetic[idx]; idx = rand.Int() % len {
		}
		mutationGenetic[idx] = true
	}

	//產生新基因
	for i := 0; i < len; i++ {
		if mutationGenetic[i] {
			if chrom.body[i] == 0 {
				newChrom.body[i] = 1
			} else {
				newChrom.body[i] = 0
			}
		} else {
			newChrom.body[i] = chrom.body[i]
		}
	}

	return newChrom
}
