## L 2.2

### Задача
Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module. Использовать
библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования
Программа должна быть оформлена как go module.

Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS.

### Комментарий
Созданный модуль github.com/alesande21/exactTime/v2 предоставляет функциональность для получения точного времени с использованием NTP.
Модуль импортируется в основное приложение и вызов метода GetTIme позволяет получить сетевое время с ntp сервера, а также сравнить его с текущем временем.
Если программа завершает выполение с ненулевым кодом ошибки то ошибка выволедится в стандарный поток ошибок.
Созданный модуль доступень по [ссылке](https://github.com/alesande21/exactTime/blob/v2/getTime.go).

https://www.socketloop.com/tutorials/golang-get-current-time-from-the-internet-time-server-ntp-example