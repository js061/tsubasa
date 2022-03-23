package tsubasago

import (
	"math"
	"time"
	"fmt"
)

func DirectCompute(thres float64, start int, end int) []int {
	// Matrix initiation
  /*matrix := make([][]int, len(dataMap))
  for i := range matrix {
    matrix[i] = make([]int, len(dataMap))
  }*/
  InitMatrix()
  newDataMap := make(map[int][]Point)
  t0 := time.Now()
  CutDataMap(&newDataMap, start, end)
	networkConstructionNaiveParallel(&newDataMap, &matrix, thres)
	elapsed := time.Since(t0)
	fmt.Println("Time:", elapsed)
	checkMatrix(&matrix)

	return GetMatrix()
}

/*func Compute(thres float64, start int, end int, granularity int) []int {
	// Matrix initiation
  matrix := make([][]int, len(dataMap))
  for i := range matrix {
    matrix[i] = make([]int, len(dataMap))
  }

	//networkConstructionBW(&dataMap, &matrix, thres, granularity, 1000, 1000, false, 1.0, start, end)
	var sketchDurations []string = make([]string, getNumCPU()-1)
  var queryDurations []string = make([]string, getNumCPU()-1)
  var queryReadTime []float64 = make([]float64, getNumCPU()-1)

	networkConstructionBWParallelSketch(&dataMap, granularity, 1000, false, 1.0, &sketchDurations)
	networkConstructionBWParallelQuery(&dataMap, &matrix, thres, granularity, 1000, false, start, end, &queryDurations, &queryReadTime)
	DeleteSkecth(false)
	checkMatrix(&matrix)

	return GetMatrix()
}*/

func Sketch(granularity int) {
	var sketchDurations []string = make([]string, getNumCPU()-1)
  networkConstructionBWParallelSketch(&dataMap, granularity, 1000, false, 1.0, &sketchDurations)
}

func Query(thres float64, start int, end int, granularity int) []int {
	InitMatrix()
	var queryDurations []string = make([]string, getNumCPU()-1)
  var queryReadTime []float64 = make([]float64, getNumCPU()-1)
  networkConstructionBWParallelQuery(&dataMap, &matrix, thres, granularity, 1000, false, start, end, &queryDurations, &queryReadTime)
  checkMatrix(&matrix)

  return GetMatrix()
}

func GetMatrix() []int {
	arr := make([]int, len(matrix) * len(matrix))
  index := 0
  for i := 0; i < len(matrix); i += 1 {
  	for j := 0; j < len(matrix); j += 1 {
  		arr[index] = matrix[i][j]
  		index += 1
  	}
  }
  return arr
}

func GetRealMatrix() []float32 {
	realArr := make([]float32, len(realMatrix) * len(realMatrix))
  index := 0
  for i := 0; i < len(realMatrix); i += 1 {
  	for j := 0; j < len(realMatrix); j += 1 {
  		if math.IsNaN(realMatrix[i][j]) {
  			realArr[index] = float32(0.0)
  		} else {
  			realArr[index] = float32(realMatrix[i][j])
  		}
  		index += 1
  	}
  }
  return realArr
}

func ResetSketch() {
	DeleteSkecth(false)
}

func SketchInMem(granularity int) {
	networkConstructionBWParallelSketchInMem(granularity)
}

func QueryInMem(thres float64, queryStart int, queryEnd int) []int {
	networkConstructionBWParallelQueryInMem(queryStart, queryEnd, thres)
	checkMatrix(&matrix)

  return GetMatrix()
}