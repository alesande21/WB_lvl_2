## L 2.1

### Задача
Реализовать паттерн, объяснить применимость паттерна, плюсы и минусы, а также реальные примеры использования паттерна на практике.
* Паттерн «Строитель».

### Комментарий
Шаблон Строитель —  это порождающий шаблон проектирования, который используется для пошагового создания сложных
объектов. Он позволяет отделить процесс построения объекта от его представления, так что один и тот же процесс
построения может создавать различные представления объекта.

Применимость

Паттерн "Строитель" часто используется, когда:
- Необходимо создавать сложные объекты с множеством опций и параметров.
- Важно избежать перегруженных конструкторов с множеством параметров.
- Требуется пошаговое построение объекта, которое позволяет контролировать и изменять процесс создания.

* Преимущества использования шаблона проектирования cтроитель:
    - помогает котролировать создание объекта
    - устраняяет необходимость в конструкторах со множеством параметров и сокращает количество избыточных конструктором
    - позволяет создавать объекты, которые не могут быть изменены после их создания.

* Недостатки использования шаблона проектирования cтроитель:
  - самый большой недостаток строителя — это возможность дублирования кода. Это связано с тем, что классу 
  строителя часто требуется полностью копировать все атрибуты класса Объект.
  - строитель будет становиться все более раздутым по мере добавления новых атрибутов.

В нашем случае нам необходимо построить объект настроек ObjectSettings на основе конфига:
1. У нас есть Builder который определяет этапы построенния объекта настроек.
2. У нас есть структура ObjectBuilder, который реализует интрефейс Builder и предоставляет способ построения объекта настроек.
3. У нас есть структура Controller которая использует Builder для создания объекта.

Реальные примеры использования паттерна "Строитель" на практике:
Создание HTTP-запросов: В сетевых библиотеках, таких как http.Client паттерн Строитель
используется для формирования HTTP-запросов с множеством заголовков, параметров и тел запроса.

https://otus.ru/nest/post/2212/
https://habr.com/ru/articles/87110/