## REST API
 
#### Cel zadania: 
  Wprowadzenie do REST API i komunikacji z bazą danych

#### Treść zadania:
Stworzyć prostą aplikację typu CRUD. 
Aplikacja ta ma połączyć się z bazą danych, której URL podany jest w dodatkowych informacjach. 
Baza ma stworzoną tabelę `todos` z 5 kolumnami: 
   - id
   - name
   - description
   - created_at
   - updated_at

Głównym zadaniem jest stworzenie aplikacji z 5 endpointami API:
   - `POST /api/todos/` - tworzenie todo
   - `GET /api/todos/` - odczytywanie wszystkich todo
   - `GET /api/todos/:id/` - odczytywanie todo na podstawie ID
   - `PATCH /api/todos/:id/` - aktualizacja todo na podstawie ID
   - `DELETE /api/todos/:id/` - usuwanie todo na podstawie ID
  
Struktura aplikacji powinna być odpowiednio podzielona na `handlery`(zajmujące się otrzymywaniem i weryfikacją żądań) oraz `modele`(do wykonywania żądań).

#### Dodatkowe informacje:
Aplikacja posiada załączoną gotową bazę danych ze stworzonym zasobem `todos`.
Należy uruchomić polecenie `docker-compose up` w głównym katalogu, co sprawi, że
pod URL `postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable` dostępna 
będzie baza danych.

#### Dodatkowe zadania:
  - weryfikacja przychodzących zapytań
  - podawać URL do bazy za pomocą flagi `db` 

#### Przydatne materiały:
  - [Router httprouter](https://github.com/julienschmidt/httprouter)
  - [Przykładowa aplikacja z todos](https://github.com/westonplatter/example-golang-todo) - nie wzorować się za mocno, jest słabo napisana i ma służyć tylko jako pomoc
  - [Obiekt sql.DB](https://golang.org/pkg/database/sql/#DB)
  - [Tutorial komunikacji Go z PostgreSQL](https://flaviocopes.com/golang-sql-database/)
  - [Tutorial REST API w Go](http://web.archive.org/web/20190305103238/https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo)
########

Zadanie polega na rozszerzeniu aplikacji stworzonej w zadaniu `rest-api`.

Lista rzeczy do zrobienia:
 - stworzyć 2 pliki w katalogu `migrations`, `1560243111_create_users_table.up.sql` oraz `1560243111_create_users_table.down.sql`.
 - w `1560243111_create_users_table.up.sql` napisać komendę do tworzenia tabeli `users` z polami `id`, `email`, `password`, `created_at`. 
 - pole `id` ma mieć typ `bigserial` i być `PRIMARY KEY`. 
 - pole `email` ma miec typ `text` oraz być `NOT NULL` i `UNIQUE`
 - pole `password` ma mieć typ `bytea`
 - pole `created_at` ma miec typ `timestamp` oraz być `NOT NULL` z default `now() at time zone 'utc'`
 - w `1560243111_create_users_table.down.sql` napisać komendę do usuwania tabeli `users`
 - w Go stworzyć typ User
 - stworzyć endpoint do rejestracji użytkownika
 - stworzyć endpoint do logowania użytkownika (póki co bez utrzymywania sesji, zwyczajne sprawdzenie hasła i emaila)
 - o hashowaniu haseł możecie poczytać [tutaj](https://gowebexamples.com/password-hashing/)