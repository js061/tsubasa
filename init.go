package tsubasago

func Init(username string, password_ string) {
	dataMap = make(map[int][]Point)
	user = username
	password = password_

	InitMatrix()

	NCPU := getNumCPU()
	pairWindowsList = make([][]BasicWindowResult, NCPU)
}

func InitMatrix() {
	matrix = make([][]int, len(dataMap))
  for i := range matrix {
    matrix[i] = make([]int, len(dataMap))
  }

  realMatrix = make([][]float64, len(dataMap))
  for i := range realMatrix {
    realMatrix[i] = make([]float64, len(dataMap))
  }
}

