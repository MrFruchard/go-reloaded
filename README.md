# go-reloaded

# Transformations de Texte

Ce programme Go permet de lire un fichier texte, d'appliquer diverses transformations sur le contenu, puis d'enregistrer le résultat dans un autre fichier. Il peut être utilisé pour convertir des nombres hexadécimaux et binaires en décimal, modifier la casse de certains mots, remplacer "a" par "an" devant des voyelles, et corriger la ponctuation.

## Fonctionnalités

- **Conversion Hexadécimale et Binaire** : Convertit les nombres hexadécimaux et binaires présents dans le texte en décimaux.
- **Transformations de Casse** : Modifie la casse des mots selon les instructions suivantes :
  - `mot (up, n)`: Met en majuscules les `n` derniers mots.
  - `mot (low, n)`: Met en minuscules les `n` derniers mots.
  - `mot (cap, n)`: Met en majuscule la première lettre des `n` derniers mots.
- **Remplacement de "a" par "an"** : Remplace le mot "a" par "an" devant les mots qui commencent par une voyelle ou un "h" muet.
- **Correction de Ponctuation** : Supprime les espaces inutiles autour des signes de ponctuation et corrige les espaces autour des apostrophes.

## Installation

1. Assurez-vous d'avoir Go installé sur votre machine. Vous pouvez le télécharger depuis [le site officiel de Go](https://golang.org/dl/).
2. Clonez ce dépôt ou téléchargez le fichier `main.go`.

## Utilisation

Pour exécuter le programme, utilisez la commande suivante dans votre terminal :

```bash
go run main.go <input_file> <output_file>
