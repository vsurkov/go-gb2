# go-gb2


HW3
1. Создайте новый проект с использованием инструментария go mod.
   ➜  go-gb2 git:(main) git branch less3
   ➜  go-gb2 git:(main) git checkout less3
   Switched to branch 'less3'

Создаем проект:
➜  go-gb2 git:(less3) ✗ tree
.
├── README.md
├── cmd
      └── less3
           └── main.go
└── internal
   └── less3
      └── my.go

Инициализируем
➜  go-gb2 git:(less3) ✗ go mod init github.com/vsurkov/go-gb2/internal/less3


2. Опубликуйте проект в репозитории, установив номер версии, указывающий на активный этап
   разработки библиотеки.
3. Обновите номера версий зависимостей в библиотеке.
4. Сделайте изменения в проекте и запушьте их с мажорным обновлением версии пакета.
5. Очистите неиспользуемые библиотеки.
