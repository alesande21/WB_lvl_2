## L 2.1

### Задача
* Паттерн «Посетитель».

### Комментарий
в сложной программе.
* Преимущества использования шаблона проектирования Посетитель:
  
* Минусы использования шаблона проектирования Посетитель:
 
* Реальные примеры использования паттерна Посетитель:

В моей реализации ITaxi объявляет операцию для каждого контрентого элемнта (человка, животного багажа). Имя и сигнатура 
операции идентифицируют структуру, которая отправляет запрос Visit посетителю. Это позволяет посетителю определить 
конкретный класс посещаемого элемента. Затем посетитель может получить доступ к элементам напрямую через его определенный интерфейс.
Taxi реализуют каждую операцию, объявленную ITaxi. Transportable определяет операцию Accept, которая принимает посетителя в качестве аргумента.
Person, Animal и Luggage реализуют операцию Accept, которая принимает ITaxi в качестве аргумента.



https://refactoring.guru/design-patterns/visitor
https://sourcemaking.com/design_patterns/visitor
https://www.opencodez.com/java/visitor-design-pattern.htm
https://touch.ethz.ch/visitor.pdf
https://stackoverflow.com/questions/2604169/visitor-patterns-purpose-with-examples
https://mydesignpatterns.wordpress.com/2009/03/20/visitor-design-pattern/