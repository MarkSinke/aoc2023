package aoc2023

import (
	"regexp"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

func (p Part) GetField(fieldName string) int {
	switch fieldName {
	case "x":
		return p.x
	case "m":
		return p.m
	case "a":
		return p.a
	case "s":
		return p.s
	}
	panic("unknown field name " + string(fieldName))
}

type Rule interface {
	// yield the name of the destination, or "" if no match
	DestinationFor(p Part) string
	// inspect the rule, return a new inspection and a potential left-over range
	Inspect(pr PartRange) (ToInspect, PartRange)
}

type FixedRule struct {
	target string
}

func (r FixedRule) DestinationFor(p Part) string {
	return r.target
}

func (r FixedRule) Inspect(pr PartRange) (ToInspect, PartRange) {
	return ToInspect{r.target, pr}, PartRange{}
}

type GreaterThanRule struct {
	fieldName string
	value     int
	target    string
}

func (r GreaterThanRule) DestinationFor(p Part) string {
	if p.GetField(r.fieldName) > r.value {
		return r.target
	}
	return ""
}

func (r GreaterThanRule) Inspect(pr PartRange) (ToInspect, PartRange) {
	pr0, pr1 := pr.SplitBefore(r.fieldName, r.value+1)
	return ToInspect{r.target, pr1}, pr0
}

type LessThanRule struct {
	fieldName string
	value     int
	target    string
}

func (r LessThanRule) DestinationFor(p Part) string {
	if p.GetField(r.fieldName) < r.value {
		return r.target
	}
	return ""
}

func (r LessThanRule) Inspect(pr PartRange) (ToInspect, PartRange) {
	pr0, pr1 := pr.SplitBefore(r.fieldName, r.value)
	return ToInspect{r.target, pr0}, pr1
}

type Workflow []Rule

type Program map[string]Workflow

func ReadWorkflowsAndParts(path string) (Program, []Part) {
	lines := ReadFile(path)

	program := Program{}
	var parts []Part

	readingProgram := true

	for _, line := range lines {
		if line == "" {
			readingProgram = false
			continue
		}

		if readingProgram {
			name, workflow := parseWorkflow(line)
			program[name] = workflow
		} else {
			parts = append(parts, parsePart(line))
		}
	}

	return program, parts
}

var workflowRegex = regexp.MustCompile(`(.*){(.*)}`)

func parseWorkflow(str string) (string, Workflow) {
	matches := workflowRegex.FindStringSubmatch(str)

	name := matches[1]
	ruleStrs := strings.Split(matches[2], ",")

	var rules []Rule
	for _, ruleStr := range ruleStrs {
		rules = append(rules, parseRule(ruleStr))
	}

	return name, rules
}

var lessThanRegex = regexp.MustCompile(`([xmas])<([0-9]+):(.*)`)
var greaterThanRegex = regexp.MustCompile(`([xmas])>([0-9]+):(.*)`)

func parseRule(str string) Rule {
	matches := lessThanRegex.FindStringSubmatch(str)

	if matches != nil {
		value, _ := strconv.Atoi(matches[2])
		return LessThanRule{matches[1], value, matches[3]}
	}

	matches = greaterThanRegex.FindStringSubmatch(str)

	if matches != nil {
		value, _ := strconv.Atoi(matches[2])
		return GreaterThanRule{matches[1], value, matches[3]}
	}

	return FixedRule{str}
}

var partsRegex = regexp.MustCompile(`{x=([0-9]+),m=([0-9]+),a=([0-9]+),s=([0-9]+)}`)

func parsePart(str string) Part {
	matches := partsRegex.FindStringSubmatch(str)
	x, _ := strconv.Atoi(matches[1])
	m, _ := strconv.Atoi(matches[2])
	a, _ := strconv.Atoi(matches[3])
	s, _ := strconv.Atoi(matches[4])
	return Part{x, m, a, s}
}

func ExecuteProgram(program Program, parts []Part) []Part {
	accepted := []Part{}

	for _, part := range parts {
		result := executeProgramPart(program, part)
		if result == "A" {
			accepted = append(accepted, part)
		}
	}

	return accepted
}

func executeProgramPart(program Program, part Part) string {
	workflowName := "in"

	for workflowName != "A" && workflowName != "R" {
		wf := program[workflowName]
		workflowName = executeWorkflow(wf, part)
	}

	return workflowName
}

func executeWorkflow(wf Workflow, part Part) string {
	for _, rule := range wf {
		result := rule.DestinationFor(part)
		if result != "" {
			return result
		}
	}
	panic("missing fixed rule")
}

func SumRatings(parts []Part) int {
	sum := 0

	for _, part := range parts {
		sum += part.x + part.m + part.a + part.s
	}

	return sum
}

// range of values, min included, max exluded
type PartRange struct {
	xMin int
	xMax int
	mMin int
	mMax int
	aMin int
	aMax int
	sMin int
	sMax int
}

func (pr PartRange) IsEmpty() bool {
	return pr.xMin >= pr.xMax || pr.mMin >= pr.mMax || pr.aMin >= pr.aMax || pr.sMin >= pr.sMax
}

func (pr PartRange) SplitBefore(fieldName string, value int) (before PartRange, after PartRange) {
	switch fieldName {
	case "x":
		before = PartRange{pr.xMin, value, pr.mMin, pr.mMax, pr.aMin, pr.aMax, pr.sMin, pr.sMax}
		after = PartRange{value, pr.xMax, pr.mMin, pr.mMax, pr.aMin, pr.aMax, pr.sMin, pr.sMax}
	case "m":
		before = PartRange{pr.xMin, pr.xMax, pr.mMin, value, pr.aMin, pr.aMax, pr.sMin, pr.sMax}
		after = PartRange{pr.xMin, pr.xMax, value, pr.mMax, pr.aMin, pr.aMax, pr.sMin, pr.sMax}
	case "a":
		before = PartRange{pr.xMin, pr.xMax, pr.mMin, pr.mMax, pr.aMin, value, pr.sMin, pr.sMax}
		after = PartRange{pr.xMin, pr.xMax, pr.mMin, pr.mMax, value, pr.aMax, pr.sMin, pr.sMax}
	case "s":
		before = PartRange{pr.xMin, pr.xMax, pr.mMin, pr.mMax, pr.aMin, pr.aMax, pr.sMin, value}
		after = PartRange{pr.xMin, pr.xMax, pr.mMin, pr.mMax, pr.aMin, pr.aMax, value, pr.sMax}
	default:
		panic("unknown field name")
	}
	return
}

type ToInspect struct {
	workflowName string
	partRange    PartRange
}

func InspectProgram(program Program) []PartRange {
	accepted := []PartRange{}
	rangesToInspect := []ToInspect{{"in", PartRange{1, 4001, 1, 4001, 1, 4001, 1, 4001}}}

	for len(rangesToInspect) > 0 {
		inspect := rangesToInspect[0]
		rangesToInspect = rangesToInspect[1:]

		newInspections, newAccepted := inspectRange(program, inspect)

		rangesToInspect = append(rangesToInspect, newInspections...)
		accepted = append(accepted, newAccepted...)
	}

	return accepted
}

func inspectRange(program Program, inspect ToInspect) ([]ToInspect, []PartRange) {
	if inspect.workflowName == "R" || inspect.partRange.IsEmpty() {
		return []ToInspect{}, []PartRange{}
	}

	if inspect.workflowName == "A" {
		return []ToInspect{}, []PartRange{inspect.partRange}
	}

	wf := program[inspect.workflowName]
	return inspectWorkflow(wf, inspect.partRange), []PartRange{}
}

func inspectWorkflow(wf Workflow, partRange PartRange) []ToInspect {
	inspects := []ToInspect{}
	for _, rule := range wf {
		inspect, residue := rule.Inspect(partRange)
		inspects = append(inspects, inspect)
		partRange = residue
	}
	return inspects
}

func SumCombinations(prs []PartRange) int {
	sum := 0
	for _, pr := range prs {
		sum += (pr.xMax - pr.xMin) * (pr.mMax - pr.mMin) * (pr.aMax - pr.aMin) * (pr.sMax - pr.sMin)
	}
	return sum
}
