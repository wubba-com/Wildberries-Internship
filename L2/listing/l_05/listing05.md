Что выведет программа? Объяснить вывод программы

```

package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}

```

Ответ:

Внутри интерфейсы реализованы в виде двух элементов, типа T и значения V. 
V - это конкретное значение, такое как int, struct 
или указатель, никогда не сам интерфейс, и имеет тип T.

Значение интерфейса равно nil, только если V и T оба не заданы, т.е nil.nil, а после присваивания err имеет *main.customError.nil

Следовательно:
```
1. Будет false, т.к T-*main.customError V-nil != nil.nil или *main.customError.<nil> != <nil> <nil>
```