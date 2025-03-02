package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

// 品目ごとの栄養素情報を持つ構造体
type PFC struct {
	Item         string  `json:"item"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Carbohydrate float64 `json:"carbohydrate"`
}

// カロリー計算関数
func CalculateCalorie(protein, fat, carbohydrate float64) int {
	// タンパク質、脂質、炭水化物のカロリーを計算
	proteincalorie := protein * 4
	fatcalorie := fat * 9
	carbohydratecalorie := carbohydrate * 4

	// カロリーを合計して四捨五入する
	return int(math.Round(proteincalorie + fatcalorie + carbohydratecalorie))
}

func main() {
	// 引数がない場合はエラーを表示して終了
	if len(os.Args) < 2 {
		fmt.Println("JSON データを引数に渡してください")
		os.Exit(1)
	}

	// JSONデータをパース
	var pfcs []PFC
	err := json.Unmarshal([]byte(os.Args[1]), &pfcs)
	if err != nil {
		fmt.Printf("JSON データのパースに失敗しました: %v\n", err)
		fmt.Printf("入力データ: %s\n", os.Args[1])
		os.Exit(1)
	}

	totalCalories := 0
	for _, pfc := range pfcs {
		calories := CalculateCalorie(pfc.Protein, pfc.Fat, pfc.Carbohydrate)
		fmt.Printf("%sのカロリー: %d kcal\n", pfc.Item, calories)
		totalCalories += calories
	}

	fmt.Printf("総カロリーは %d kcal です。\n", totalCalories)
}
