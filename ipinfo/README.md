## Pobieranie i obróbka danych z API [ipinfo](https://ipinfo.io/)
 
#### Cel zadania: 
  Nauka komunikacji z API przy pomocy formatu [JSON](https://www.json.org/)

#### Treść zadania:
  Stworzyć program, który przyjmuje adres IP jako argument(flagę) o nazwie. 
  Następnie program wykonuje zapytanie do [API ipinfo](https://ipinfo.io/developers), którego odpowiedź ma zdekodować do wcześniej zadeklarowanego typu(struktury).
  Po zdekodowaniu odpowiedzi, program ma wyświetlić pola struktury w następującej formie:
  ```
    IP address: xxx.xxx.xxx.xxx
    Hostname: xxxx.xxxx.xxxxx
    City: xxxxxx
    Region: xxxxxx
    Country: xxx
    Loc: xxxx
    Postal: xxxxx
  ```

##### Przykładowe użycie programu ma wyglądać następująco:
  ```
  go run *.go -ip 81.190.40.214

  IP address: 81.190.40.214
  Hostname: mail.idesign.center
  City: Szczecin
  Region: West Pomerania
  Country: PL
  Loc: 53.4358, 14.5094
  Postal: 71-132
  ```

#### Przydatne materiały:
  - [dokumentacja API ipinfo.io](https://ipinfo.io/developers)
  - [dokumentacja pakietu net/http](https://golang.org/pkg/net/http/)
  - [przykładowe zapytania przy użyciu pakietu net/http](https://dlintw.github.io/gobyexample/public/http-client.html)
  - [dokumentacja pakietu encoding/json](https://golang.org/pkg/encoding/json/)
  - [przykłady kodowania i dekodowania JSON](https://gobyexample.com/json)
  - [dokumentacja pakietu flag](https://golang.org/pkg/flag/)
  - [przykłady użycia flag w Go](https://gobyexample.com/command-line-flags)