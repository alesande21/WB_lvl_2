## L 2.1

### Задача
Реализовать паттерн, объяснить применимость паттерна, плюсы и минусы, 
а также реальные примеры использования паттерна на практике.
* Паттерн «Команда».

### Комментарий
* Описание паттерна:
  Паттерн «Команда» — это поведенческий шаблон проектирования, который превращает запросы в объекты. 
  Это позволяет параметризовать объекты другими запросами, ставить запросы в очередь, протоколировать их, 
  а также поддерживать операции отмены. Основная идея паттерна заключается в том, чтобы отделить отправителя 
  команды от ее исполнителя.

  * Когда использовать:
  -  Если необходимо инкапсулировать запросы как объекты для последующего протоколирования, логирования, очередности выполнения или возможности отмены операций.
  - Когда нужно обеспечить гибкое управление выполнением действий в ответ на события.
  - Паттерн позволяет легко реализовать операции отмены или повторного выполнения.

* Преимущества использования шаблона проектирования Команда:
  -  Паттерн отделяет объект, который инициирует выполнение команды (отправитель), от объекта, который выполняет эту команду (получатель). Это повышает гибкость и упрощает поддержку кода.
  - Добавление новой команды не требует изменений в существующем клиентском коде, что соответствует принципу открытости/закрытости.
  - Логика команд изолирована в отдельных объектах, что упрощает их тестирование.

* Минусы использования шаблона проектирования Команда:
  - Паттерн может привести к значительному увеличению количества классов, что может усложнить проект.
  - Если в системе много различных команд, управление их реализацией и поддержка могут оказаться затратными.

* Реальные примеры использования паттерна Команда:
  - Для управления движением персонажей или выполнением действий по нажатию клавиш.
  - Используется для обработки команд, таких как действия меню или кнопки «Отменить» и «Повторить» в приложениях.


В моей реализации Invoker это объект, который хранит и запускает команды.  MainWindow это получатель, который 
содержит методы для переключения калькулятора в разные режимы. Команды (SwitchToBasic, SwitchToProgrammer, 
SwitchToScientific): каждая команда инкапсулирует конкретное действие и знает, как связаться с получателем (MainWindow) для выполнения действия.