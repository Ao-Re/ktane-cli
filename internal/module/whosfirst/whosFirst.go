package whosfirst

import (
	"container/list"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Position int

const (
	TopLeft Position = iota + 1
	MiddleLeft
	BottomLeft
	TopRight
	MiddleRight
	BottomRight
)

var DisplayTextList = map[string]Position{
	"yes":      MiddleLeft,
	"first":    TopRight,
	"display":  BottomRight,
	"okay":     TopRight,
	"says":     BottomRight,
	"nothing":  MiddleLeft,
	"":         BottomLeft,
	"blank":    MiddleRight,
	"no":       BottomRight,
	"led":      MiddleLeft,
	"lead":     BottomRight,
	"read":     MiddleRight,
	"red":      MiddleRight,
	"reed":     BottomLeft,
	"leed":     BottomLeft,
	"hold on":  BottomRight,
	"you":      MiddleRight,
	"you are":  BottomRight,
	"your":     MiddleRight,
	"you're":   MiddleRight,
	"ur":       TopLeft,
	"there":    BottomRight,
	"they're":  BottomLeft,
	"their":    MiddleRight,
	"they are": MiddleLeft,
	"see":      BottomRight,
	"c":        TopRight,
	"cee":      BottomRight,
}

var LabelTextList = map[string][]string{
	"ready": {
		"YES", "OKAY", "WHAT", "MIDDLE",
		"LEFT", "PRESS", "RIGHT", "BLANK",
		"READY",
	},

	"first": {
		"LEFT", "OKAY", "YES", "MIDDLE",
		"NO", "RIGHT", "NOTHING", "UHHH",
		"WAIT", "READY", "BLANK", "WHAT",
		"PRESS", "FIRST",
	},

	"no": {
		"BLANK", "UHHH", "WAIT", "FIRST",
		"WHAT", "READY", "RIGHT", "YES",
		"NOTHING", "LEFT", "PRESS", "OKAY",
		"NO",
	},

	"blank": {
		"WAIT", "RIGHT", "OKAY", "MIDDLE",
		"BLANK",
	},

	"nothing": {
		"UHHH", "RIGHT", "OKAY", "MIDDLE",
		"YES", "BLANK", "NO", "PRESS",
		"LEFT", "WHAT", "WAIT", "FIRST",
		"NOTHING",
	},

	"yes": {
		"OKAY", "RIGHT", "UHHH", "MIDDLE",
		"FIRST", "WHAT", "PRESS", "READY",
		"NOTHING", "YES",
	},

	"what": {
		"UHHH", "WHAT",
	},

	"uhhh": {
		"READY", "NOTHING", "LEFT", "WHAT",
		"OKAY", "YES", "RIGHT", "NO",
		"PRESS", "BLANK", "UHHH",
	},

	"left": {
		"RIGHT", "LEFT",
	},

	"right": {
		"YES", "NOTHING", "READY", "PRESS",
		"NO", "WAIT", "WHAT", "RIGHT",
	},

	"middle": {
		"BLANK", "READY", "OKAY", "WHAT",
		"NOTHING", "PRESS", "NO", "WAIT",
		"LEFT", "MIDDLE",
	},

	"okay": {
		"MIDDLE", "NO", "FIRST", "YES",
		"UHHH", "NOTHING", "WAIT", "OKAY",
	},

	"wait": {
		"UHHH", "NO", "BLANK", "OKAY",
		"YES", "LEFT", "FIRST", "PRESS",
		"WHAT", "WAIT",
	},

	"press": {
		"RIGHT", "MIDDLE", "YES", "READY",
		"PRESS",
	},

	"you": {
		"SURE", "YOU ARE", "YOUR", "YOU'RE",
		"NEXT", "UH HUH", "UR", "HOLD",
		"WHAT?", "YOU",
	},

	"you are": {
		"YOUR", "NEXT", "LIKE", "UH HUH",
		"WHAT?", "DONE", "UH UH", "HOLD",
		"YOU", "U", "YOU'RE", "SURE",
		"UR", "YOU ARE",
	},

	"your": {
		"UH UH", "YOU ARE", "UH HUH", "YOUR",
	},

	"you're": {
		"YOU", "YOU'RE",
	},

	"ur": {
		"DONE", "U", "UR",
	},

	"u": {
		"UH HUH", "SURE", "NEXT", "WHAT?",
		"YOU'RE", "UR", "UH UH", "DONE",
		"U",
	},

	"uh huh": {
		"UH HUH",
	},

	"uh uh": {
		"UR", "U", "YOU ARE", "YOU'RE",
		"NEXT", "UH UH",
	},

	"what?": {
		"YOU", "HOLD", "YOU'RE", "YOUR",
		"U", "DONE", "UH UH", "LIKE",
		"YOU ARE", "UH HUH", "UR", "NEXT",
		"WHAT?",
	},

	"done": {
		"SURE", "UH HUH", "NEXT", "WHAT?",
		"YOUR", "UR", "YOU'RE", "HOLD",
		"LIKE", "YOU", "U", "YOU ARE",
		"UH UH", "DONE",
	},

	"next": {
		"WHAT?", "UH HUH", "UH UH", "YOUR",
		"HOLD", "SURE", "NEXT",
	},

	"hold": {
		"YOU ARE", "U", "DONE", "UH UH",
		"YOU", "UR", "SURE", "WHAT?",
		"YOU'RE", "NEXT", "HOLD",
	},

	"sure": {
		"YOU ARE", "DONE", "LIKE", "YOU'RE",
		"YOU", "HOLD", "UH HUH", "UR",
		"SURE",
	},

	"like": {
		"YOU'RE", "NEXT", "U", "UR",
		"HOLD", "DONE", "UH UH", "WHAT?",
		"UH HUH", "YOU", "LIKE",
	},
}

type WhosFirst struct {
	DisplayText string
	LabelTexts  []string
}

func (wf *WhosFirst) Init() {
	rand.Seed(time.Now().UnixNano())

	wf.DisplayText = generateDisplayText()
	wf.LabelTexts = generateLabelTexts(wf.DisplayText)
}

func generateDisplayText() string {
	displayTexts := make([]string, len(DisplayTextList))

	for k := range DisplayTextList {
		displayTexts = append(displayTexts, k)
	}

	return displayTexts[rand.Intn(len(displayTexts))]
}

// label order: TopLeft, MiddleLeft, BottomLeft, TopRight, MiddleRight, BottomRight
func generateLabelTexts(display string) []string {
	selectedPosition := DisplayTextList[display]
	labels := make([]string, len(LabelTextList))

	for k := range LabelTextList {
		labels = append(labels, k)
	}

	selectedLabel := labels[rand.Intn(len(labels))]

	labelTextsList := make([]string, len(LabelTextList[selectedLabel]))

	for _, v := range LabelTextList[selectedLabel] {
		if v != selectedLabel {
			labelTextsList = append(labelTextsList, v)
		}
	}

	rand.Shuffle(len(labelTextsList), func(i, j int) {
		labelTextsList[i], labelTextsList[j] = labelTextsList[j], labelTextsList[i]
	})

	labelTextsQueue := list.New()

	for _, v := range labelTextsList {
		labelTextsQueue.PushBack(v)
	}

	labelTexts := make([]string, 6)

	for i := range labelTexts {
		if i == int(selectedPosition)-1 { // Put the selectedLabel in the correct Position
			labelTexts = append(labelTexts, selectedLabel)
		} else if labelTextsQueue.Len() > 0 { // Put the randomized possible labels as long as the queue is not empty
			labelTexts = append(labelTexts, fmt.Sprintf("%v", labelTextsQueue.Front().Value))
			labelTextsQueue.Remove(labelTextsQueue.Front())
		} else { // Put random label to fill in the rest of the labels if the queue is empty
			tempLabel := labels[rand.Intn(len(labels))]

			for checkIfExist(labelTexts, tempLabel) { // Randomize the selected label while the it exist in the slice
				tempLabel = labels[rand.Intn(len(labels))]
			}

			labelTexts = append(labelTexts, tempLabel)
		}
	}

	return labelTexts
}

func checkIfExist(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}

	return false
}

func (wf *WhosFirst) CheckLogic(ans string) bool {
	solution := wf.SolveWhosFirst()
	return ans == solution
}

func (wf *WhosFirst) SolveWhosFirst() string {
	var labelPosition Position
	labelPosition = DisplayTextList[strings.ToLower(wf.DisplayText)]

	labelText := strings.ToLower(wf.LabelTexts[labelPosition-1])

	for _, possibleAnswer := range LabelTextList[labelText] {
		for _, label := range wf.LabelTexts {
			if possibleAnswer == label {
				return possibleAnswer
			}
		}
	}

	return ""

}
