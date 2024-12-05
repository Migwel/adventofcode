package main

import (
	"math"
	"strconv"
	"strings"
)

func orderPages(rulesStr string, updatesStr string) int {
	rules := parseRules(rulesStr)
	updates := parseUpdates(updatesStr)
	result := 0
outer:
	for _, update := range updates {
		for i := 0; i < len(update)-1; i++ {
			currVal := update[i]
			isUpdateValid, _ := isUpdateValid(currVal, update[i+1:], rules)
			if !isUpdateValid {
				continue outer
			}
		}
		result += findMidValue(update)
	}
	return result
}

func fixUnorderedPages(rulesStr string, updatesStr string) int {
	rules := parseRules(rulesStr)
	updates := parseUpdates(updatesStr)
	result := 0
outer:
	for _, update := range updates {
		for i := 0; i < len(update)-1; i++ {
			currVal := update[i]
			isUpdateValid, _ := isUpdateValid(currVal, update[i+1:], rules)
			if !isUpdateValid {
				fixedUpdate := fixUpdate(update, rules)
				result += findMidValue(fixedUpdate)
				continue outer
			}
		}
	}
	return result
}

func fixUpdate(update []string, rules map[string][]string) []string {
	for i := 0; i < len(update); i++ {
		currVal := update[i]
		isUpdateValid, idx := isUpdateValid(currVal, update[i+1:], rules)
		if !isUpdateValid {
			fixedUpdated := moveWrongValue(update, i, i+1+idx)
			return fixUpdate(fixedUpdated, rules)
		}
	}
	return update
}

func moveWrongValue(update []string, currIdx, targetIdx int) []string {
	currVal := update[currIdx]
	var fixedUpdate []string
	for i := 0; i < len(update); i++ {
		if i == currIdx {
			continue
		}
		fixedUpdate = append(fixedUpdate, update[i])
		if targetIdx == i {
			fixedUpdate = append(fixedUpdate, currVal)
		}
	}
	return fixedUpdate
}

func findMidValue(update []string) int {
	midIndex := (int)(math.Floor((float64)(len(update) / 2)))
	val, _ := strconv.Atoi(update[midIndex])
	return val

}

func isUpdateValid(currVal string, remValues []string, rules map[string][]string) (bool, int) {
	for i := 0; i < len(remValues); i++ {
		remValue := remValues[i]
		rule := rules[remValue]
		if len(rule) == 0 {
			continue
		}
		for _, afterValue := range rule {
			if currVal == afterValue {
				return false, i
			}
		}
	}
	return true, -1
}

func parseUpdates(updatesStr string) [][]string {
	rows := strings.Split(updatesStr, "\n")
	updates := make([][]string, len(rows))
	for i := 0; i < len(rows); i++ {
		updateSplit := strings.Split(rows[i], ",")
		row := make([]string, len(updateSplit))
		for j := 0; j < len(updateSplit); j++ {
			row[j] = updateSplit[j]
		}
		updates[i] = row
	}
	return updates
}

func parseRules(rulesStr string) map[string][]string {
	rows := strings.Split(rulesStr, "\n")
	rules := make(map[string][]string)
	for i := 0; i < len(rows); i++ {
		ruleSplit := strings.Split(rows[i], "|")
		existingRule := rules[ruleSplit[0]]
		if len(existingRule) == 0 {
			rules[ruleSplit[0]] = []string{ruleSplit[1]}
		} else {
			rules[ruleSplit[0]] = append(rules[ruleSplit[0]], ruleSplit[1])
		}
	}
	return rules
}
