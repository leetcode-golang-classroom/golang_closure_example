# golang_closure_sample

This is a repository for demo how closure is work

## 什麼是 closure

每個 function 要被執行之前會根據資料相依性

把該 function 執行所需要的相依資料與邏輯區塊放置於一個記憶體內

## 範例

```golang=
func activateGiftCard() func(int) int {
	amount := 100
	debitFunc := func(debitAmount int) int {
	   amount -= debitAmount
	   return amount
	}
	return debitFunc
}

func main() {
  useGiftCard1 := activateGiftCard()
  useGiftCard2 := activateGiftCard()
  fmt.Println(useGiftCard1(10))
  fmt.Println(useGiftCard2(5))
}
```

在執行時

![](https://i.imgur.com/YTc0H5v.png)

會把 acitvateGiftCard 內的程式區塊與變數各自複製到記憶體內

所以執行的數值各自獨立

分別會印出 90, 95