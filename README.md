#Evolutionary Password Cracker

A lightweight, from-scratch implementation of a Genetic Algorithm written in Go. 

Inspired by Daniel Shiffman's "The Nature of Code", this project demonstrates the core pillars of evolutionary search by evolving a population of random characters into a specific target string over successive generations.

## Why I Built This
I built this project to master the mathematical and programmatic fundamentals of evolutionary algorithms, optimization, and search spaces before applying these concepts to modern neural networks and AI alignment. Go was chosen for its execution speed, allowing thousands of generations to process in milliseconds.

## Core Mechanics

This algorithm relies on standard evolutionary principles:
* **Fitness Evaluation:** Each "DNA" string is scored based on its character-by-character accuracy against the target phrase.
* **Rank Selection:** Parents are chosen for reproduction based on their relative rank in the population, rather than raw fitness, preserving genetic diversity.
* **Crossover:** Two parent strings are sliced at a random midpoint and combined to create a child.
* **Mutation:** A 1% mutation rate introduces random new characters into the gene pool to prevent local minima stagnation.
* **Elitism:** The top-performing DNA sequence is explicitly preserved in the next generation to guarantee the algorithm never regresses.
