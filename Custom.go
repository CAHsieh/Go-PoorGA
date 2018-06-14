package poorga

//Custom 用來導入需要客製化的方法
type Custom interface {
	InitChromosome(chromosome *Chromosome)        // 用來做染色體的初始化，請調用chromosome的Init，並傳入染色體長度即可。
	Fitness(chromosome *Chromosome)               //配適度函數，所有的染色體會透過此方法來計算其配適度，並根據配適度由大至小來擇優保留至下一代。
	Print(iteration int, chromosome []Chromosome) //若World設置isPrint為true的話，每一代染色體擇優完後會調用此方法。
	PrintResult(chromosome []Chromosome)          //基因演算法結束(達成任一停止條件時)，會調用此方法。
}
