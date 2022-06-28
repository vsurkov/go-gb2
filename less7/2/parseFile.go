package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func parseFile(srcFileName, funcName string) (int, error) {
	var count = 0
	fset := token.NewFileSet()

	// парсим файл, чтобы получить AST
	astFile, err := parser.ParseFile(fset, srcFileName, nil, 0)
	if err != nil {
		return 0, err
	}

	// получаем Scope
	for _, obj := range astFile.Scope.Objects {
		// проверяем что объект это функция и имя == искомое
		if obj.Kind.String() == "func" && obj.Name == funcName {

			// получаем для объекта Decl, и приводим к типу *ast.FuncDecl что бы получить Body.List
			decl := obj.Decl
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}

			// получаем список объектов из тела, и если очередной элемент соответствует *ast.GoStmt увеличиваем счетчик
			for _, li := range funcDecl.Body.List {
				_, ok := li.(*ast.GoStmt)
				if !ok {
					continue
				}
				count += 1
			}

		}
	}
	//TODO добавить шагание вглубь, потому что если положить в цикл go func() то это не будет посчитано
	go func() {
	}()
	go func() {
	}()
	return count, nil
}
