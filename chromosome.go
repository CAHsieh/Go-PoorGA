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

func (chrom *Chromosome) init(len int) {
	chrom.fitness = -1
	chrom.length = len
	chrom.body = make([]int, len)
	for i := 0; i < len; i++ {
		chrom.body[i] = rand.Int() % 2
	}
}

func (chrom *Chromosome) crossover(secChrom Chromosome, crossRate float64) Chromosome {
	var newChrom Chromosome
	len := chrom.length
	newChrom.length = len
	newChrom.body = make([]int, len)
	newChrom.fitness = -1

	motherNum := int(float64(len) * crossRate)
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

func (chrom *Chromosome) mutate(mutationRate float64) Chromosome {
	var newChrom Chromosome
	len := chrom.length
	newChrom.length = len
	newChrom.body = make([]int, len)
	newChrom.fitness = -1

	mutationNum := int(float64(len) * mutationRate)
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
