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

//InitChromosome is the method to implement Custom interface
func (matcher StringMatcher) InitChromosome(chromosome *Chromosome) {
	//ascii 32~126 -> 0~94 + 32
	//2^6 < 94 < 2^7
	chromosome.Init(8 * len(matcher.target))
}

//Fitness is the method to implement Custom interface
func (matcher StringMatcher) Fitness(chromosome *Chromosome) {
	targetLen := len(matcher.target)
	// for i := 0; i < len(*chromosome); i++ {
	currentString := matcher.getString(targetLen, chromosome)
	acc := 0
	for j := 0; j < targetLen; j++ {
		if matcher.target[j] == currentString[j] {
			acc++
		}
	}
	chromosome.SetFitness(float64(acc) / float64(targetLen))
	// fmt.Print(string(currentString) + " ")
	// fmt.Println(chromosome.fitness)
	// }
}

//Print is the method to implement Custom interface
func (matcher StringMatcher) Print(iteration int, chromosome []Chromosome) {
	if iteration%300 != 0 {
		return
	}
	targetLen := len(matcher.target)
	currentString := matcher.getString(targetLen, &chromosome[0])
	fmt.Printf("%d: %s %f\n", iteration, string(currentString), chromosome[0].GetFitness())
}

//PrintResult is the method to implement Custom interface
func (matcher StringMatcher) PrintResult(chromosome []Chromosome) {
	for i, chr := range chromosome {
		fmt.Printf("Chromosome No.%d\tfitness:%f\n", i, chr.GetFitness())
		fmt.Printf("Result: %s\n", string(matcher.getString(len(matcher.target), &chr)))
		fmt.Print("Chromosome body: ")
		fmt.Println(chr.GetBody())
	}
}

func (matcher StringMatcher) getString(targetLen int, chromosome *Chromosome) []rune {
	result := make([]rune, targetLen)
	body := chromosome.GetBody()
	for j := 0; j < targetLen; j++ {
		st := j * 8
		end := (j + 1) * 8
		value := 0
		for k := st; k < end; k++ {
			if body[k] == 1 {
				value += int(math.Pow(2, float64(k-st)))
			}
		}
		result[j] = rune(value)
	}
	return result
}
