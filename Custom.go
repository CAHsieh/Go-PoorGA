package poorga

//Custom 用來導入需要客製化的方法
type Custom interface {
	InitChromosome(chromosome *Chromosome)
	Fitness(chromosome *Chromosome)
	Print(iteration int, chromosome []Chromosome)
	PrintResult(chromosome []Chromosome)
}
