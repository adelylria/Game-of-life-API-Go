package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// gridSize define el tamaño de la cuadrícula (grid) del juego.
// Es una constante que establece que la cuadrícula será de 10x10 celdas.
const gridSize = 10

// GameState es una estructura que representa el estado actual del juego.
// Contiene un campo Grid, que es una matriz bidimensional de booleanos.
// El etiquetado `json:"grid"` indica que este campo se serializará/deserializará como "grid" en formato JSON.
type GameState struct {
	Grid [][]bool `json:"grid"`
}

// initializeGrid inicializa una nueva cuadrícula con celdas vivas y muertas aleatorias.
// Devuelve una matriz bidimensional (grid) de booleanos.
// El valor `true` representa una célula viva y `false` representa una célula muerta.
func initializeGrid() [][]bool {
	// Se crea una matriz bidimensional (grid) de tamaño gridSize x gridSize.
	grid := make([][]bool, gridSize)

	// Se recorre cada fila de la cuadrícula.
	for i := range grid {
		// Para cada fila, se crea una nueva fila de tamaño gridSize.
		grid[i] = make([]bool, gridSize)

		// Se recorre cada celda de la fila.
		for j := range grid[i] {
			// Se asigna a la celda un valor booleano aleatorio.
			// Si rand.Float32() devuelve un número menor que 0.5, la celda estará viva (true).
			grid[i][j] = rand.Float32() < 0.5
		}
	}

	// Se devuelve la cuadrícula inicializada.
	return grid
}

// nextGeneration toma una cuadrícula existente y calcula la siguiente generación basada en las reglas del juego de la vida.
// Devuelve una nueva cuadrícula que representa la siguiente generación.
func nextGeneration(grid [][]bool) [][]bool {
	// Se inicializa una nueva cuadrícula vacía que contendrá la siguiente generación.
	newGrid := initializeGrid()

	// Se recorre cada fila de la cuadrícula actual.
	for i := range grid {
		// Se recorre cada celda en la fila actual.
		for j := range grid[i] {
			// Se cuenta el número de vecinos vivos alrededor de la celda actual.
			liveNeighbors := countLiveNeighbors(grid, i, j)

			// Se aplican las reglas del juego de la vida:
			if grid[i][j] {
				// Si la celda está viva, permanecerá viva si tiene 2 o 3 vecinos vivos.
				newGrid[i][j] = liveNeighbors == 2 || liveNeighbors == 3
			} else {
				// Si la celda está muerta, se convertirá en viva si tiene exactamente 3 vecinos vivos.
				newGrid[i][j] = liveNeighbors == 3
			}
		}
	}

	// Se devuelve la nueva cuadrícula con la siguiente generación.
	return newGrid
}

// countLiveNeighbors cuenta cuántos vecinos vivos hay alrededor de una celda específica en la cuadrícula.
// Toma como parámetros la cuadrícula, y las coordenadas x e y de la celda cuya vecindad se está evaluando.
// Devuelve el número de vecinos vivos como un entero.
func countLiveNeighbors(grid [][]bool, currentRow, currentCol int) int {
	// liveNeighbors almacena el número de vecinos vivos.
	liveNeighbors := 0

	// Se recorren las posiciones alrededor de la celda actual (incluyendo diagonales).
	for rowOffset := -1; rowOffset <= 1; rowOffset++ {
		for colOffset := -1; colOffset <= 1; colOffset++ {
			// Si ambos índices son 0, significa que estamos en la celda actual, no en un vecino.
			if rowOffset == 0 && colOffset == 0 {
				continue // Se omite la celda actual.
			}

			// Calcula las coordenadas del vecino (ni, nj).
			neighborRow, neighborCol := currentRow+rowOffset, currentCol+colOffset

			// Verifica que las coordenadas del vecino están dentro de los límites de la cuadrícula.
			// Si el vecino está dentro de los límites y está vivo, se incrementa el contador de vecinos vivos.
			/*
				[0, 1, 0]   -> [0, 0]  [0, 1]  [0, 2]
				[1, 1, 0]   -> [1, 0]  [1, 1]  [1, 2]
				[0, 0, 1]   -> [2, 0]  [2, 1]  [2, 2]

				[0, 0], [0, 1], [0, 2]   (Fila superior)
				[1, 0],         [1, 2]   (Fila actual, vecinos laterales)
				[2, 0], [2, 1], [2, 2]   (Fila inferior)

				2. Vecino (0, 1):

				Coordenadas: neighborRow = 1 + (-1) = 0, neighborCol = 1 + 0 = 1.
				Estado: grid[0][1] = 1 (viva).
				Resultado: Se incrementa el contador liveNeighbors++.

				Vecino (0, 2):

				Coordenadas: neighborRow = 1 + (-1) = 0, neighborCol = 1 + 1 = 2.
				Estado: grid[0][2] = 0 (muerta).
				Resultado: No se incrementa el contador.
			*/
			if neighborRow >= 0 && neighborRow < gridSize && neighborCol >= 0 && neighborCol < gridSize && grid[neighborRow][neighborCol] {
				liveNeighbors++
			}
		}
	}

	// Se devuelve el número total de vecinos vivos.
	return liveNeighbors
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	grid := initializeGrid()
	grid = nextGeneration(grid)
	json.NewEncoder(w).Encode(GameState{Grid: grid})
}
