# go-reloaded

Outil en Go qui lit un fichier texte, applique des corrections automatiques et écrit le résultat dans un autre fichier.


## Utilisation

go run . sample.txt result.txt

`sample.txt`: fichier d'entrée
`result.txt`: fichier de sortie 

## Fonctionalités

- `(hex)`: convertit le mot précédent de l'hexadécimal vers le décimal
- `(bin)`: convertit le mot précédent du binaire vers le décimal
- `(up)`, `(low)`, `(cap)` : majuscules, minuscules, initiale en majuscule 
(avec nombre possible : `(up, 2)`)
- Ponctuation collée au mot précédent 
- Apostrophes collés au texte qu'elles entourent 
- `a`devient `an`devant une voyelle ou un `h`

## Structure 

- `main.go` : lecture des arguments et des fichiers 
- `capitalize.go` : fonctions Capitalize, Upper; Lower 
- `applycase.go` : gestion de (up), (low), (cap)
- `convertion.go` : gestion de (hex) et (bin)
- `article.go` : règle a -> an
- `punctuation.go` : ponctuation
- `quotes.go` : apostrophes
- `process.go` : enchaîne toutes les règles
- `process_test.go` : tests unitaires