Когда завершите задачу, в этом README опишите свой ход мыслей: как вы пришли к решению, какие были варианты и почему выбрали именно этот. 

# Что нужно сделать

Реализовать интерфейс с методом для проверки правил флуд-контроля. Если за последние N секунд вызовов метода Check будет больше K, значит, проверка на флуд-контроль не пройдена.

- Интерфейс FloodControl располагается в файле main.go.

- Флуд-контроль может быть запущен на нескольких экземплярах приложения одновременно, поэтому нужно предусмотреть общее хранилище данных. Допустимо использовать любое на ваше усмотрение. 

# Необязательно, но было бы круто

Хорошо, если добавите поддержку конфигурации итоговой реализации. Параметры — на ваше усмотрение.

# Цепочка мыслей

По началу не очень понял задание и подумал, что под хранилищем данных подразумевается база данных или нечто подобное. Уже написав программу для работы с базой данных понял, что это совсем не то, что имеллось в виду по тз. Перечитав тз понял что для конечного ответа от функции Check нужно всего три параметра, два из них уже даны в тз время за которое отсылаются запросы и колличество хапросов, которое является верхней планкой по колличесву запросов. Последний пункт, который я понял уже после написания теста по отправке запроса, что по истечению времени (секунд) буффер запросов должен сам собой подчищаться, иначе при отправке спустя некоторое время пользователь автоматически будет в флуд-блоке, что делает систему одноразовой. Для этого была реализованна функция cleanup() в файле foottest.go. Для возможности использования сразу несколькими пользователями реалазация была выполнена при помощи канала/горутины. Что касается вопроса конфигурации, решить ее можно было несколькими вариантами, простой docx файл, exel таблицы или просто чтение из баз данных, я выбрал вариант с Json форматом потому, что работа с ним довольно проста и быстра в разработке. В файле floodtest начиная с 93 строки идет закоментированный код, который является основой для main2.go и main4.go. Вот пожалуй и все, что я могу сказать об этом мини проекте.
