# Projet Hangman-Web

## Présentation de l'équipe et du groupe

**Groupe : Hangweb**

- Johann Chiarotto
- Baptiste Renou
- Adrien Dulou

---

## Sommaire

1. [Présentation du projet](#présentation-du-projet)
2. [Pré-requis](#pré-requis)
3. [Installation et démarrage](#installation-et-démarrage)
4. [Fonctionnalités](#fonctionnalités)
5. [Structure du projet](#structure-du-projet)
6. [Points importants](#points-importants)
7. [Crédits](#crédits)

---

## Présentation du projet

**Hangman-Web** est une adaptation web du célèbre jeu du pendu. Le jeu est développé en langage Golang et s’exécute via un serveur web local. L'objectif est de deviner un mot choisit aussi aléatoirement que possible, en proposant des lettres et avec un nombre limité de tentatives.

Ce projet a été réalisé dans le cadre d’un travail de groupe par trois étudiants de 1ere année.

---

## Pré-requis

Pour exécuter ce projet, vous aurez besoin de :

1. **Git** – pour cloner le dépôt.
2. **Go (Golang)** – version 1.20 ou ultérieure recommandée.
3. **Un navigateur web** – pour interagir avec le jeu.


Assurez-vous que ces outils soient installés sur votre système avant de commencer.

---

## Installation et démarrage


Choisir un emplacement pour le programme, puis entrer la commande ``` git clone  https://ytrack.learn.ynov.com/git/cjohann/hangman-web.git ``` dans une fenêtre de commande bash, afin de cloner le reporsitory git.

Il faut ensuite se déplacer dans le dossier "server" (```cd hangman-web/server```) et enfin lancer le programme avec ```go run main.go```. Le jeu sera alors accessible en ouvrant un naviguateur, puis en entrant l'url: http://localhost:8080


---

## Fonctionnalités

* Interface web simple et intuitive permettant de jouer au pendu via un navigateur.

* Choix 'aléatoire' des mots mystères, et de leur ordre pour une expérience unique à chaque partie.

* Historique des lettres déjà proposées, pour aider le joueur à ne pas répéter ses choix.

* Indicateur de vies restantes, limitant le nombre de tentatives possibles.

* Gestion des entrées utilisateur avec traitement des erreurs.

---

## Structure du projet

* **/server/**: Contient le code source du serveur web et du jeu.

* **/CSS/** : Contient les fichiers de style pour l’interface utilisateur.

* **index.html** : La page principale du jeu.

---

## Points importants

1. **URL d’accès au jeu :** Le serveur est actif sur le port 8080 (http://localhost:8080).

2. **Modifications du code :** Toutes les fonctions liées à la logique du jeu se trouvent dans le fichier Fonctions.

3. **Personnalisation :** Vous pouvez modifier les mots mystères ou les styles en éditant les fichiers correspondants.



---

Nous espérons que vous apprécierez jouer à notre version web du jeu du pendu !
