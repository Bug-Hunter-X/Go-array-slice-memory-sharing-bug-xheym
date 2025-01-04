func main() {
        var a [10]int
        fmt.Println(a)
        for i := 0; i < 10; i++ {
                a[i] = i
        }
        fmt.Println(a)
        b := make([]int, len(a[:5]))
        copy(b, a[:5])
        fmt.Println(b)
        for i := 0; i < 5; i++ {
                b[i] = i * 10
        }
        fmt.Println(b)
        fmt.Println(a)
} 