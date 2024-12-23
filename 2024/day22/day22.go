package day22

import (
	"strconv"
	"strings"
)

func ComputeSumOfSecrets(input string, nbIterations int) int {
	currentSecrets := parseInput(input)
	var nextSecrets []int
	for i := 0; i < nbIterations; i++ {
		for _, currentSecret := range currentSecrets {
			nextSecret := computeNextSecret(currentSecret)
			nextSecrets = append(nextSecrets, nextSecret)
		}
		currentSecrets = nextSecrets
		nextSecrets = nil
	}
	sum := 0
	for _, currentSecret := range currentSecrets {
		sum += currentSecret
	}
	return sum
}

func ComputeMostBananas(input string, nbIterations int) int {
	currentSecrets := parseInput(input)
	var nextSecrets []int
	lastDigits := make([][]int, len(currentSecrets))
	for idx, currentSecret := range currentSecrets {
		lastDigits[idx] = make([]int, nbIterations+1)
		lastDigits[idx][0] = currentSecret % 10
	}

	for i := 0; i < nbIterations-1; i++ {
		for idx, currentSecret := range currentSecrets {
			nextSecret := computeNextSecret(currentSecret)
			nextSecrets = append(nextSecrets, nextSecret)
			lastDigits[idx][i+1] = nextSecret % 10
		}
		currentSecrets = nextSecrets
		nextSecrets = nil
	}
	differences := calculateDifferences(lastDigits)
	sequences := calculateSequences(lastDigits, differences)
	existingSequences := findExistingSequences(sequences)
	return findMostBananas(sequences, existingSequences)
}

func findExistingSequences(sequences []map[string]int) map[string]bool {
	existingSequences := make(map[string]bool)
	for _, sequence := range sequences {
		for val, _ := range sequence {
			existingSequences[val] = true
		}
	}
	return existingSequences
}

func findMostBananas(sequences []map[string]int, existingSequences map[string]bool) int {
	maxBananas := 0
	for existingSequence, _ := range existingSequences {
		bananas := findSumBananasForSequence(sequences, existingSequence)
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}
	return maxBananas
}

func findSumBananasForSequence(otherSequences []map[string]int, sequence string) int {
	count := 0
	for _, otherSequence := range otherSequences {
		if val, ok := otherSequence[sequence]; ok {
			count += val
		}
	}
	return count
}

func calculateSequences(secretsLastDigits [][]int, secretsDifferences [][]int) []map[string]int {
	sequences := make([]map[string]int, len(secretsDifferences))
	for secretsIdx, secretDifferences := range secretsDifferences {
		secretSequence := make(map[string]int)
		for secretIdx := 3; secretIdx < len(secretDifferences); secretIdx++ {
			sequence := strconv.Itoa(secretDifferences[secretIdx-3]) + "," +
				strconv.Itoa(secretDifferences[secretIdx-2]) + "," +
				strconv.Itoa(secretDifferences[secretIdx-1]) + "," +
				strconv.Itoa(secretDifferences[secretIdx])
			if _, ok := secretSequence[sequence]; !ok {
				secretSequence[sequence] = secretsLastDigits[secretsIdx][secretIdx+1]
			}
		}
		sequences[secretsIdx] = secretSequence
	}
	return sequences
}

func calculateDifferences(secretsLastDigits [][]int) [][]int {
	differences := make([][]int, len(secretsLastDigits))
	for idx, secretLastDigits := range secretsLastDigits {
		secretDifferences := make([]int, len(secretLastDigits)-1)
		previousLastDigit := secretLastDigits[0]
		for i := 1; i < len(secretLastDigits); i++ {
			secretDifferences[i-1] = secretLastDigits[i] - previousLastDigit
			previousLastDigit = secretLastDigits[i]
		}
		differences[idx] = secretDifferences
	}
	return differences
}

func computeNextSecret(currentSecret int) int {
	firstStepValue := firstStep(currentSecret)
	secondStepValue := secondStep(firstStepValue)
	return thirdStep(secondStepValue)
}

func thirdStep(value int) int {
	mixedValue := mix(value*2048, value)
	return prune(mixedValue)
}

func secondStep(value int) int {
	mixedValue := mix(value/32, value)
	return prune(mixedValue)
}

func firstStep(value int) int {
	mixedValue := mix(value*64, value)
	return prune(mixedValue)
}

func prune(value int) int {
	return value % 16777216
}

func mix(val1, val2 int) int {
	return val1 ^ val2
}

func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	secrets := make([]int, len(lines))
	for idx, line := range lines {
		secret, _ := strconv.Atoi(line)
		secrets[idx] = secret
	}
	return secrets
}
