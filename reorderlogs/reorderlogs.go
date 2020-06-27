package reorderlogs

import (
	"sort"
	"strconv"
	"strings"
)

type logRecord struct {
	Identifier, Values string
	IsDigit            bool
}

func parseLogRecord(log string) logRecord {
	logParts := strings.SplitN(log, " ", 2)
	logIdentifier := logParts[0]
	logValues := logParts[1]

	_, err := strconv.Atoi(string(logValues[0]))

	return logRecord{
		Identifier: logIdentifier,
		Values:     logValues,
		IsDigit:    err == nil,
	}
}

func compareLogRecords(log1, log2 logRecord) int {

	// Compare letter log records lexicographically
	if !log1.IsDigit && !log2.IsDigit {
		comparsionResult := strings.Compare(log1.Values, log2.Values)

		// Compare identifiers if values are equal
		if comparsionResult == 0 {
			comparsionResult = strings.Compare(log1.Identifier, log2.Identifier)
		}

		return comparsionResult
	}

	switch {
	case log1.IsDigit && log2.IsDigit:
		return 0
	case log1.IsDigit && !log2.IsDigit:
		return 1
	default:
		return -1
	}
}

// ReorderLogFiles implemens a solution for "Reorder Data in Log Files" problem
// 	see: https://leetcode.com/problems/reorder-data-in-log-files/
func ReorderLogFiles(logs []string) []string {

	sort.SliceStable(logs, func(i, j int) bool {
		logRecord1 := parseLogRecord(logs[i])
		logRecord2 := parseLogRecord(logs[j])

		comparsionResult := compareLogRecords(logRecord1, logRecord2)

		return comparsionResult < 0
	})

	return logs
}
