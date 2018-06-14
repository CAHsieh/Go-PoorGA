package poorga

//Custom 用來導入需要客製化的方法
type Custom interface {
	initChromosome(chromosome *Chromosome)
	fitness(chromosome *Chromosome)
	print(iteration int, chromosome []Chromosome)
	printResult(chromosome []Chromosome)
}
