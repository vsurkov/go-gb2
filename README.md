# go-gb2


#HW3 
код в less3
1. Создайте новый проект с использованием инструментария go mod.
2. Опубликуйте проект в репозитории, установив номер версии, указывающий на активный этап
   разработки библиотеки.
3. Обновите номера версий зависимостей в библиотеке.
4. Сделайте изменения в проекте и запушьте их с мажорным обновлением версии пакета.
5. Очистите неиспользуемые библиотеки.

#HW5
1. Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех
2. Реализуйте функцию для разблокировки мьютекса с помощью defer
3. Протестируйте производительность операций чтения и записи на множестве
   действительных чисел, безопасность которого обеспечивается sync.Mutex и sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
   
cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz результаты:

##10% запись, 90% чтение
###Mutex set, writeScope:100, readScope:900 of total:1000
BenchmarkSetUsecaseOne/#00-8   	10084777	       112.6 ns/op
BenchmarkSetUsecaseOne/#01-8   	10733632	       105.6 ns/op

###RWMutex set, writeScope:100, readScope:900 of total:1000
BenchmarkRWSetUsecaseOne/#00-8 	 7474468	       173.6 ns/op
BenchmarkRWSetUsecaseOne/#01-8 	 7081318	       172.4 ns/op

##50% запись, 50% чтение
###Mutex set, writeScope:500, readScope:500 of total:1000
BenchmarkSetUsecaseTwo/#00-8   	 9296889	       131.1 ns/op
BenchmarkSetUsecaseTwo/#01-8   	 9873106	       124.8 ns/op

###RWMutex set, writeScope:500, readScope:500 of total:1000
BenchmarkRWSetUsecaseTwo/#00-8 	 6321628	       180.5 ns/op
BenchmarkRWSetUsecaseTwo/#01-8 	 7826137	       164.9 ns/op

##90% запись, 10% чтение
###Mutex set, writeScope:900, readScope:100 of total:1000
BenchmarkSetUsecaseThree/#00-8 	10001839	       125.2 ns/op
BenchmarkSetUsecaseThree/#01-8 	10677118	       124.8 ns/op

###RWMutex set, writeScope:900, readScope:100 of total:1000
BenchmarkRWSetUsecaseThree/#00-8         	 6283282	       193.7 ns/op
BenchmarkRWSetUsecaseThree/#01-8         	 7071294	       171.5 ns/op
