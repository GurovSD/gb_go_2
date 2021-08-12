# Урок 2 #

## Задание 1 ##

_Выполните сборку ваших предыдущих программ под операционную систему, отличающуюся от текущей. Проанализируйте вывод команды file для полученного исполняемого файла. Попробуйте запустить исполняемый файл_

команду file в windows использовал через консоль git bash

    Под свою ос:
    > go build main.go
   
    $ file a.out.exe
    a.out.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
    
    Под другую ОС:
    > set GOOS=linux
    > go build main.go
    
    
    $ file main
    main: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=eu8xIk1l1qucWnsc9Mi0/6Od6-nC83ewAvHFucwMX/rhPd3mPu7GBbaO4ftzft/9Il1lncZIEVgNXn4LerR, not stripped

## Задание 2 ##

_Напишите документацию для одной из предыдущих программ. Запустите сервер документации локально. Убедитесь, что документация отображается корректно._

Документацию к своему проекту запустить не получилось.

Обратил внимание, что при запуске godoc отличается вывод консоли от вашего.
У вас - `using GOPATH mode`
у меня - `using module mode; GOMOD=NUL`

Видимо поэтому он не видит папку с проектом и не добавляет на страницу документации.
Задам вопрос по этому поводу в телеграме.
А так сам сервер документации запустился:

![godoc screenshot](https://github.com/GurovSD/gb_go_2/blob/les2/les2/godoc%20screenhot.PNG)
