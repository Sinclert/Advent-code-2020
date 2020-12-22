package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)


var tileIDRegex = regexp.MustCompile(`\d+`)


//UNFINISHED
func main() {

	tiles, err := parseImageTilesFile("solutions/20/files/image_tiles.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var availableTiles = tiles[1:]
	var matchedTiles   = []ImageTile{tiles[0]}

	for len(availableTiles) > 0 {

		for _, availableTile := range availableTiles {
			var availableEdges = availableTile.GetFreeEdges()

			for _, availableEdge := range availableEdges {

				for _, matchedTile := range matchedTiles {
					matchFound, matchID := matchEdgeWithTile(availableEdge, matchedTile)

					if matchFound == false {
						continue
					}

					edgeStruct := availableTile.edges[matchID]
					edgeStruct.free = false
					availableTile.edges[matchID] = edgeStruct

					matchedTiles = append(availableTiles)
					//Remove from available
				}
			}
		}
	}

	//UNFINISHED
}


func parseImageTilesFile(filepath string) ([]ImageTile, error) {

	var imageTiles []ImageTile

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var currentTileID         string
	var currentTileMatrix [][]string

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if tileIDRegex.MatchString(line) {
			currentTileID = tileIDRegex.FindString(line)
			continue
		}

		if len(line) == 0 {
			imageTiles = append(imageTiles, NewImageTile(currentTileID, &currentTileMatrix))
			continue
		}

		lineCharacters   := strings.Split(line, "")
		currentTileMatrix = append(currentTileMatrix, lineCharacters)
	}

	return imageTiles, nil
}


func matchEdgeWithTile(targetEdge ImageTileEdge, tile ImageTile) (bool, string) {

	for _, edge := range tile.GetFreeEdges() {

		if reflect.DeepEqual(edge.vals, targetEdge.vals) {
			edgeStruct := tile.edges[edge.ID]
			edgeStruct.free = false
			tile.edges[edge.ID] = edgeStruct
			return true, edge.ID
		}
	}

	tile.FlipMatrixHorizontally()

	for _, edge := range tile.GetFreeEdges() {

		if reflect.DeepEqual(edge.vals, targetEdge.vals) {
			edgeStruct := tile.edges[edge.ID]
			edgeStruct.free = false
			tile.edges[edge.ID] = edgeStruct
			return true, edge.ID
		}
	}

	tile.FlipMatrixVertically()

	for _, edge := range tile.GetFreeEdges() {

		if reflect.DeepEqual(edge.vals, targetEdge.vals) {
			edgeStruct := tile.edges[edge.ID]
			edgeStruct.free = false
			tile.edges[edge.ID] = edgeStruct
			return true, edge.ID
		}
	}

	return false, ""
}
