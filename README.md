# go-gb2


HW3 
код в less3
1. Создайте новый проект с использованием инструментария go mod.
2. Опубликуйте проект в репозитории, установив номер версии, указывающий на активный этап
   разработки библиотеки.
3. Обновите номера версий зависимостей в библиотеке.
4. Сделайте изменения в проекте и запушьте их с мажорным обновлением версии пакета.
5. Очистите неиспользуемые библиотеки.


## HW6
1. Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких потоков. Выполните трассировку программы

- инструментируем:
	trace.Start(os.Stderr)
	defer trace.Stop()
- собираем трассировку:
GOMAXPROCS=1 run main.go serve.go 2>trace.out
или другое количество GOMAXPROCS
- запускаем утилиту и смотрим:
go tool trace trace.out 
➜  1 git:(less6) ✗ go tool trace trace.out                         
2022/06/19 00:59:19 Parsing trace...
2022/06/19 00:59:19 Splitting trace...
2022/06/19 00:59:19 Opening browser. Trace viewer is listening on http://127.0.0.1:55543


2. Написать многопоточную программу, в которой будет использоваться явный вызов планировщика. Выполните трассировку программы
- в горутину воркеров добавить runtime.Gosched()
- остальное как в п.1

3. Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”
- моделируем не используя мьютекс или атомик в горутине листнера
- запускаем go run **-race** main.go
и получаем предупреждение:

>WARNING: DATA RACE
>Read at 0x00c00013c038 by goroutine 8:
>  main.main.func1()
>      /Users/HOMEr/Dropbox/Geekbrains/Go_course/go-gb-repo/go-gb2/less6/3/main.go:29 +0xcc

>Previous write at 0x00c00013c038 by goroutine 7:
>  main.main.func1()
>      /Users/HOMEr/Dropbox/Geekbrains/Go_course/go-gb-repo/go-gb2/less6/3/main.go:29 +0xe6

>Goroutine 8 (running) created at:
>  main.main()
>      /Users/HOMEr/Dropbox/Geekbrains/Go_course/go-gb-repo/go-gb2/less6/3/main.go:25 +0xfc

>Goroutine 7 (running) created at:
>  main.main()
>      /Users/HOMEr/Dropbox/Geekbrains/Go_course/go-gb-repo/go-gb2/less6/3/main.go:25 +0xfc

