package models

import (
	"fmt"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/consts"
	"strings"
)

type PreflopRFIRange struct {
	PlayerName  string             `json:"playerName"`
	Position    string             `json:"-"`
	HolecardMap map[string]Actions `json:"holecardMap"`
}

type Actions struct {
	Raise int			`json:"raise"`
	Fold  int			`json:"fold"`
	Call  int			`json:"call"`
}

func (actions Actions) ToString() string {
	return fmt.Sprintf("[r%df%dc%d]", actions.Raise, actions.Fold, actions.Call)
}

func InitPreflopRFIRange(name, position string) PreflopRFIRange {
	return PreflopRFIRange{
		PlayerName:  name,
		Position:    position,
		HolecardMap: make(map[string]Actions),
	}
}

func (rfi PreflopRFIRange) PrintRFIMatrix() {
	fmt.Println(rfi.Position)
	for i := 0; i < 13; i++ {
		holecardsSlice := make([]string, 0)
		for j := 0; j < 13; j++ {
			holecardStr := MatrixToHolecards(i, j)
			if freq, exists := rfi.HolecardMap[holecardStr]; exists {
				holecardsSlice = append(holecardsSlice, fmt.Sprintf("%s %v", holecardStr, freq.ToString()))
			} else {
				holecardsSlice = append(holecardsSlice, fmt.Sprintf("%s %v", holecardStr, Actions{}.ToString()))
			}
		}
		rowStr := strings.Join(holecardsSlice, "|")
		fmt.Println(rowStr)
	}
}

func MatrixToHolecards(i, j int) string {
	iWeight := 14 - i
	jWeight := 14 - j
	suited := ""
	if i < j {
		suited = "s"
		return consts.WeightToCardChar[iWeight] + consts.WeightToCardChar[jWeight] + suited
	} else if i > j {
		suited = "o"
		return consts.WeightToCardChar[jWeight] + consts.WeightToCardChar[iWeight] + suited
	} else {
		return consts.WeightToCardChar[iWeight] + consts.WeightToCardChar[jWeight]
	}
}
