// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	fmt.Println(now.Format(time.RFC1123))
// }

// package main

// import (
//     "fmt"
//     "log"
//     "time"
// )

// Преобразовать currentTimeStr в формат time.

// func main() {
//     currentTimeStr := "2021-09-19T15:59:41+03:00"
//     // скопируйте блок себе в IDE и допишите код
//     layout := time.RFC3339
//     currentTime, err := time.Parse(layout, currentTimeStr)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(currentTime)
// } 

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     now := time.Now()
//     fmt.Println("Год:", now.Year())
//     fmt.Println("Месяц:", now.Month())
//     fmt.Println("Число:", now.Day())
//     fmt.Println("День недели:", now.Weekday())
//     hour, min, sec := now.Clock()
//     fmt.Printf("Время: %d, %d, %d\n", hour, min, sec)
//     fmt.Println("Часовой пояс:", now.Location())
//     fmt.Println("timestamp в секундах:", now.Unix())
//     fmt.Println("timestamp в наносекундах:", now.UnixNano())
// }

// package main

// import (
//     "fmt"
//     "log"
//     "time"
// )


// Сравнить время с текущим.

// func main() {
//     currentTimeStr := "2021-09-19T15:59:41+03:00"
//     layout := time.RFC3339
//     currentTime, err := time.Parse(layout, currentTimeStr)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(currentTime)

//     now := time.Now()

//     fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
//     fmt.Println("Is", now, "after", currentTime, "? Answer:", now.After(currentTime))
//     fmt.Println("Is", now, "equal", currentTime, "? Answer:", now.Equal(currentTime))
// } 