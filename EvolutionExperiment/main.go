package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// My target phrase that I am trying to get to
const target = "ThisIsAPassword12345"
const phraselen = len(target)

// List of all characters
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

// DNA struct for storing each "guess" in my population
type DNA struct {
	gene    []byte
	fitness float64
}

func NewDNA(length int) DNA {
	return DNA{
		gene:    createPhrase(length),
		fitness: 0.0,
	}
}

// create phrase function that takes a length of how long you want your phrase to be and returns a random slice of bytes.
func createPhrase(n int) []byte {
	finalPhrase := make([]byte, n)
	for i := 0; i < n; i++ {
		finalPhrase[i] = letters[rand.Intn(len(letters))]
	}
	return finalPhrase
}

// fitness calculator. *DNA allows me to modify the actual fitness value using the pointer. This prevents Go from creating a copy of the struct, update the copy's fitness, and immediately throwing it away.
func (d *DNA) calculateFitness(target string) {
	score := 0
	for i := 0; i < len(d.gene); i++ {
		if d.gene[i] == target[i] {
			score++
		}
	}
	d.fitness = float64(score) / float64(len(target))
}

func rankSelection(sortedPop []DNA) DNA {
	popSize := len(sortedPop)

	totalRanks := (popSize * (popSize + 1)) / 2

	target := rand.Intn(totalRanks)

	runningSum := 0
	for i := 0; i < popSize; i++ {
		rank := i + 1
		runningSum += rank

		if runningSum > target {
			return sortedPop[i]
		}
	}
	return sortedPop[popSize-1]
}

func (d *DNA) crossover(partner DNA) DNA {
	child := NewDNA(len(d.gene))
	midpoint := rand.Intn(len(d.gene))

	for i := 0; i < len(d.gene); i++ {
		if i > midpoint {
			child.gene[i] = d.gene[i]
		} else {
			child.gene[i] = partner.gene[i]
		}
	}
	return child
}

func (d *DNA) mutate(rate float64) {
	for i := 0; i < len(d.gene); i++ {
		if rand.Float64() < rate {
			d.gene[i] = letters[rand.Intn(len(letters))]
		}
	}
}

func main() {
	start := time.Now()

	popSize := 150
	mutationRate := 0.01

	population := make([]DNA, popSize)
	for i := 0; i < popSize; i++ {
		population[i] = NewDNA((phraselen))
		population[i].calculateFitness(target)
	}

	generations := 0
	finished := false

	for !finished {
		generations++

		// STEP 1: RANK THE POPULATION (Sort worst to best)
		sort.Slice(population, func(i, j int) bool {
			return population[i].fitness < population[j].fitness
		})

		// Create a temporary slice to hold the next generation
		newPopulation := make([]DNA, popSize)

		elite := population[popSize-1]
		newPopulation[0] = elite

		// STEP 2: REPRODUCTION
		for i := 1; i < popSize; i++ {
			// Pick parents using Rank Selection
			parentA := rankSelection(population)
			parentB := rankSelection(population)

			// Prevent the same parent from being picked twice
			attempts := 0
			for string(parentA.gene) == string(parentB.gene) && attempts < 10 {
				parentB = rankSelection(population)
				attempts++
			}

			// Crossover & Mutate
			child := parentA.crossover(parentB)
			child.mutate(mutationRate)

			newPopulation[i] = child
		}

		// Replace the old population with the new one
		population = newPopulation

		// STEP 3: EVALUATE NEW GENERATION
		bestFitness := 0.0
		bestPhrase := ""

		for i := 0; i < popSize; i++ {
			population[i].calculateFitness(target)

			if population[i].fitness > bestFitness {
				bestFitness = population[i].fitness
				bestPhrase = string(population[i].gene)
			}

			if population[i].fitness == 1.0 {
				finished = true
			}
		}

		fmt.Printf("Gen %d: %s (Fitness: %.2f)\n", generations, bestPhrase, bestFitness)
	}

	elapsed := time.Since(start)
	fmt.Printf("\nSuccess! Target matched in %d generations.\n", generations)
	fmt.Printf("Total time taken: %s\n", elapsed)
}
