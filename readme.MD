# Fizzbuzz
#### Mais cette fois, c'est un web service!
## Sujet
Exercise: Write a simple fizz-buzz REST server.
 
The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by `fizz`, all multiples of 5 by `buzz`, and all multiples of 15 by `fizzbuzz`.
The output would look like this: `1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...`.
 
Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters: three integers `int1`, `int2` and limit, and two strings `str1` and `str2`.
- Returns a list of strings with numbers from `1` to `limit`, where: all multiples of `int1` are replaced by `str1`, all multiples of `int2` are replaced by `str2`, all multiples of `int1` and `int2` are replaced by `str1str2`.
 
The server needs to be:
- Ready for production
- Easy to maintain by other developers
 
Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request
## Rendu
Le rendu ne contient aucune dépendance externe. Le module n'utilise que la librairie standard afin de permettre une plus grande extensibilité à l'aide des interfaces standards
### `/pkg`
`/pkg` contient la core library du module, ainsi que la définition des interfaces à utiliser pour intéragir avec cette library
- StatCounter est une interface qui définit deux méthodes, `Add()` et `GetMostRequested()`
La librairie utilise par défaut un compteur basé sur une map et un `Sync.RWMutex`. Le compteur peut-être remplacé par n'importe quelle structure remplissant les conditions de l'interface avec `pkg.SetStatCounter()`
- Request est une interface qui va définir une requête. La librairie possède deux fonctions, `FizzBuzz()` pour une requête standalone, et `FizzBuzzWithStats()` pour une requête qui s'enregistre dans le StatCounter, qui prennent en argument une requête. Il est possible d'en générer une avec `NewRequest()` 
### `/cmd/cli`
Dans `/cmd/cli`, un programme standalone permet de tester la logique. Il s'invoque comme ceci:
```
> ./cli 3 5 16 fizz buzz
1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16
```
### `/cmd/server`
Dans `/cmd/server` se trouve le serveur du webservice. Il se lance comme ceci:
```
> ./server --port=8080
```
Il n'utilise que la librairie standard `net/http`, son port est configurable à l'aide du packet standard `flag`
### `/internal`
`/internal` contient les routes à utiliser par `/cmd/server`. Il utilise le package `log` par defaut de go pour log les erreurs, et ne repose que sur le packet `net/http`