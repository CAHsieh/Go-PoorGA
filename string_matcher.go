package poorga

import (
	"fmt"
	"math"
)

//StringMatcher 用來測試基因演算法匹配輸入的內容
type StringMatcher struct {
	target []rune
}

//SetTarget 用於設置目標字串
func (matcher *StringMatcher) SetTarget(target string) {
	matcher.target = []rune(target)
}

func (matcher StringMatcher) initChromosome(chromosome *Chromosome) {
	//ascii 32~126 -> 0~94 + 32
	//2^6 < 94 < 2^7
	chromosome.init(8 * len(matcher.target))
}
func (matcher StringMatcher) fitness(chromosome *Chromosome) {
	targetLen := len(matcher.target)
	// for i := 0; i < len(*chromosome); i++ {
	currentString := matcher.getString(targetLen, chromosome)
	acc := 0
	for j := 0; j < targetLen; j++ {
		if matcher.target[j] == currentString[j] {
			acc++
		}
	}
	chromosome.fitness = float64(acc) / float64(targetLen)
	// fmt.Print(string(currentString) + " ")
	// fmt.Println(chromosome.fitness)
	// }
}

func (matcher StringMatcher) print(iteration int, chromosome []Chromosome) {
	if iteration%300 != 0 {
		return
	}
	targetLen := len(matcher.target)
	currentString := matcher.getString(targetLen, &chromosome[0])
	fmt.Printf("%d: %s %f\n", iteration, string(currentString), chromosome[0].fitness)
}

func (matcher StringMatcher) printResult(chromosome []Chromosome) {
	for i, chr := range chromosome {
		fmt.Printf("Chromosome No.%d\tfitness:%f\n", i, chr.fitness)
		fmt.Printf("Result: %s\n", string(matcher.getString(len(matcher.target), &chr)))
		fmt.Print("Chromosome body: ")
		fmt.Println(chr.body)
	}
}

func (matcher StringMatcher) getString(targetLen int, chromosome *Chromosome) []rune {
	result := make([]rune, targetLen)
	for j := 0; j < targetLen; j++ {
		st := j * 8
		end := (j + 1) * 8
		value := 0
		for k := st; k < end; k++ {
			if chromosome.body[k] == 1 {
				value += int(math.Pow(2, float64(k-st)))
			}
		}
		result[j] = rune(value)
	}
	return result
}
