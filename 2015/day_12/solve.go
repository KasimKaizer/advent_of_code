package day12

import "encoding/json"

func SolveOne(input string) (int, error) {
	var jsonMap map[string]any
	err := json.Unmarshal([]byte(input), &jsonMap)
	if err != nil {
		return 0, err
	}
	return calculateTotal(jsonMap, func(_ any) bool { return false }), nil
}

func SolveTwo(input string) (int, error) {
	var jsonMap map[string]any
	err := json.Unmarshal([]byte(input), &jsonMap)
	if err != nil {
		return 0, err
	}
	return calculateTotal(jsonMap, func(v any) bool {
		if v, ok := v.(string); ok {
			return v == "red"
		}
		return false
	}), nil
}

func calculateTotal(jsonMap map[string]any, pred func(any) bool) int {
	var total int
	for _, val := range jsonMap {
		if pred(val) {
			return 0
		}
		switch v := val.(type) {
		case int:
			total += v
		case float64:
			total += int(v)
		case []any:
			total += reduce(v, pred)
		case map[string]any:
			total += calculateTotal(v, pred)
		}
	}
	return total
}

func reduce(slc []any, pred func(any) bool) int {
	var total int
	for _, v := range slc {
		switch i := v.(type) {
		case int:
			total += i
		case float64:
			total += int(i)
		case []any:
			total += reduce(i, pred)
		case map[string]any:
			total += calculateTotal(i, pred)
		}
	}
	return total
}

// Failed attempt, it worked for part 1 & part 2 example inputs,
// but didn't work for actual input (for part 2).
// probably because I tried to parse the json myself instead of using the stdlib.
//
// Gonna do what i did for one of the problems on exercism.
// For the first part, I just need to convert every number (which would be in form of a string)
// into an int and add it to some variable.
//
// Remember whenever we are inside an object
//
// func SolveOne(input string) (int, error) {
// 	var sum int
// 	for i := 0; i < len(input); i++ {
// 		if input[i] != '-' && !(input[i] >= '0' && input[i] <= '9') {
// 			continue // skip over all character that are't numbers or a negative sign.
// 		}
// 		// the character at the current index is number
// 		start := i
// 		for input[i] == '-' || input[i] >= '0' && input[i] <= '9' {
// 			i++
// 		}
//
// 		num, err := strconv.Atoi(input[start:i])
// 		if err != nil {
// 			// handle error
// 		}
// 		sum += num
// 	}
// 	return sum
// }
//
// func SolveTwo(input string) (int, error) {
// 	var sum int
// 	var eval, redEval int
// 	evalSum := make(map[int]int)
// 	for i := 0; i < len(input); i++ {
// 		if input[i] == '{' {
// 			eval++
// 		}
// 		if input[i] == '}' {
// 			if eval < redEval || redEval == 0 {
// 				sum += evalSum[eval]
// 				evalSum[eval] = 0
// 				redEval = 0
// 			}
// 			eval--
// 		}
// 		if input[i] == ':' && i < len(input)-5 && input[i:i+6] == `:"red"` {
// 			if redEval == 0 {
// 				redEval = eval
// 			} else {
// 				redEval = min(redEval, eval)
// 			}
// 		}
// 		if input[i] == '-' || input[i] >= '0' && input[i] <= '9' {
// 			start := i
// 			for input[i] == '-' || input[i] >= '0' && input[i] <= '9' {
// 				i++
// 			}
// 			num, err := strconv.Atoi(input[start:i])
// 			if err != nil {
// 				// handle error
// 			}
// 			evalSum[eval] += num
// 			i--
// 		}
// 	}
// 	return sum
// }
