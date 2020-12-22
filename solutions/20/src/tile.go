package main


type ImageTileEdge struct {
	ID             string
	vals           *[]string
	free           bool
}


type ImageTile struct {
	ID             string
	matrix         *[][]string
	matrixHeight   int
	matrixWidth    int
	edges          map[string]ImageTileEdge
	matches        map[string]string
}


func NewImageTile(id string, matrix *[][]string) ImageTile {

	return ImageTile{
		ID: id,
		matrix: matrix,
		matrixHeight: len(*matrix),
		matrixWidth: len((*matrix)[0]),
		edges: map[string]ImageTileEdge{
			"Upper": {ID: "Upper", vals: &(*matrix)[0],                 free: true},
			"Lower": {ID: "Lower", vals: &(*matrix)[len(*matrix)-1],    free: true},
			"Left":  {ID: "Left",  vals: &(*matrix)[:][0],              free: true},
			"Right": {ID: "Right", vals: &(*matrix)[:][len(*matrix)-1], free: true},
		},
		matches: map[string]string{
			"Upper": "NONE",
			"Lower": "NONE",
			"Left":  "NONE",
			"Right": "NONE",
		},
	}
}


func (t ImageTile) _reverseArray(array []string) []string {

	var i = 0
	var j = len(array) - 1

	for i < j {
		array[i], array[j] = array[j], array[i]
		i += 1
		j -= 1
	}

	return array
}


func (t ImageTile) GetFreeEdges() []ImageTileEdge {

	var freeEdges []ImageTileEdge

	for _, edge := range t.edges {
		if edge.free {
			freeEdges = append(freeEdges, edge)
		}
	}

	return freeEdges
}


func (t ImageTile) FlipMatrixHorizontally() {

	var flippedMatrix = make([][]string, t.matrixHeight)

	for i, row := range *t.matrix {
		newRow := t._reverseArray(row)
		flippedMatrix[i] = newRow
	}

	t.matrix = &flippedMatrix
}


func (t ImageTile) FlipMatrixVertically() {

	var flippedMatrix = make([][]string, t.matrixHeight)

	for i, _ := range *t.matrix {
		reverseRowIndex := t.matrixHeight-(i+1)
		flippedMatrix[i] = (*t.matrix)[reverseRowIndex]
	}

	t.matrix = &flippedMatrix
}


func (t ImageTile) RotateMatrix() {

	var flippedMatrix = make([][]string, t.matrixHeight)

	// Initiate all matrix positions
	for i := 0; i < t.matrixWidth; i++ {
		flippedMatrix[i] = make([]string, t.matrixWidth)
	}

	// Rotates 90º in a clock-wise manner
	for i, _ := range *t.matrix {
		for j, _ := range (*t.matrix)[0] {
			flippedMatrix[j][t.matrixWidth-(i+1)] = (*t.matrix)[i][j]
		}
	}

	t.matrix = &flippedMatrix
}
