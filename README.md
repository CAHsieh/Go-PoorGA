# **這是一個簡陋的基因演算法框架**

*Thie is a poor genetic algorithm framework*

## **Overview**

本專案只有簡單的三個主體：chromosome.go、World.go、Custom.go。

> **chromosome.go**
> ---
>> Chromosome為染色體物件，其中包含三屬性body、length、fitness及三方法init、crossover、mutate：
>>
>> |屬性|說明|
>> | :---: | :--- |
>> |body|染色體的本體，由一串01組成。|
>> |legth|染色體長度。|
>> |fitness|此染色體的環境配適度。|
>> ---
>> |方法|說明|
>> | :---: | :--- |
>> |Init|用於初始化染色體，必須在Custom中的initChromosome內調用|
>> |SetFitness|用於設置配適度|
>> |GetFitness|用於取得配適度|
>> |GetBody|用於取得染色體內容|
>> |crossover|傳入另一條染色體及交配率，產生交配後的下一代染色體。|
>> |mutate|傳入突變率，產生突變後的下一代染色體。|
> **World.go**
> ---
>> World為模擬染色體生存環境的物件，在使用時需建立此物件並調用其方法來執行基因演算法。
>>
>> |屬性|說明|
>> | :---: | :--- |
>> | chromosome    | 每個世代存活的染色體 |
>> | randomNum     | 每一世代中會隨機選擇及存活下來的染色體數量 |
>> | generationNum | 每個世代存活的染色體個數 |
>> | iterationNum  | 停止條件1：總迭代次數 |
>> | optimizeTimes | 停止條件2：局部最佳迭代次數 |
>> | goodEnough    | 停止條件3：足夠停止的準確度 |
>> | customMethod  | 客製化介面 |
>> | isPrint       | 是否顯示當前狀態 |
>> ---
>> |方法|說明|
>> | :---: | :--- |
>> |Initial|用於設置相關參數|
>> |SetIsPrint|用來控制是否在GA過程中調用Custom介面中的Print|
>> |StartWorld|開始執行基因演算法|
>> |crossover|產生交配後的下一世代|
>> |mutate|產生突變後的下一世代|
>> |selection|根據配適度選擇下一世代存活下來的染色體群，若配適度相同則優先序為突變 > 交配 > 本世代染色體|
> **Custom.go**
> ---
>> Custom為需要使用者客製化的介面，其中包含InitChromosome、Fitness、Print、PrintResult
>>
>> |方法|說明|
>> | :---: | :--- |
>> |InitChromosome|用來做染色體的初始化，請調用chromosome的Init，並傳入染色體長度即可。|
>> |Fitness|配適度函數，所有的染色體會透過此方法來計算其配適度，並根據配適度由大至小來擇優保留至下一代。|
>> |Print|若World設置isPrint為true的話，每一代染色體擇優完後會調用此方法。|
>> |PrintResult|基因演算法結束(達成任一停止條件時)，會調用此方法。|

## **What you need to do**

1. 客製化Module並實現Custom.go中的interface方法

1. 建立World物件，依序調用Initial、StartWorld。

1. SetIsPrint自行選擇是否調用

## **Example**

```go
package main

import poorga "go-poorGA"

func main() {

	var stringMatcher poorga.StringMatcher
	stringMatcher.SetTarget("Hello World!")
	var world poorga.World
	world.Initial(50, 20, 300000, 300000, 1.0, stringMatcher)
	world.SetIsPrint(true)
	world.StartWorld()

}
```
### **Result**
![result](https://github.com/CAHsieh/Go-PoorGA/blob/master/result.png)