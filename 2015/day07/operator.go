package day7

import "strings"

type operator func(s1, s2, t *wire)

func not(source1, source2, target *wire) {
	if !isValidOperation(source1, source2) {
		return
	}
	target.signal.value = ^source1.signal.value
	target.signal.valid = true
}

func and(source1, source2, target *wire) {
	if !isValidOperation(source1, source2) {
		return
	}
	target.signal.value = source1.signal.value & source2.signal.value
	target.signal.valid = true
}

func or(source1, source2, target *wire) {
	if !isValidOperation(source1, source2) {
		return
	}
	target.signal.value = source1.signal.value | source2.signal.value
	target.signal.valid = true
}

func lshift(source1, source2, target *wire) {
	if !isValidOperation(source1, source2) {
		return
	}
	target.signal.value = source1.signal.value << source2.signal.value
	target.signal.valid = true
}

func rshift(source1, source2, target *wire) {
	if !isValidOperation(source1, source2) {
		return
	}
	target.signal.value = source1.signal.value >> source2.signal.value
	target.signal.valid = true
}

func isValidOperation(source1, source2 *wire) bool {
	if source2 != nil {
		return source1.signal.valid && source2.signal.valid
	} else {
		return source1.signal.valid
	}
}

func getOperator(operatorString string) operator {
	if strings.EqualFold(operatorString, "and") {
		return and
	}

	if strings.EqualFold(operatorString, "or") {
		return or
	}

	if strings.EqualFold(operatorString, "lshift") {
		return lshift
	}

	if strings.EqualFold(operatorString, "rshift") {
		return rshift
	}

	return nil
}
