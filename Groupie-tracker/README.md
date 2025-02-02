# Projet Groupie-Tracker


## Présentation du groupe et de l'équipe

*Groupe : **API-Crackers***

- **Johann Chiarotto**
- **Baptiste Renou** 
- **Adrien Dulou** 

Notre équipe a collaboré sur ce projet dans le cadre du cours Groupie-tracker. Nous avons pris en charge différents aspects du projet, allant du backend (Go) au frontend (HTML, CSS, JS).


[Accédez au tableau Trello du projet](https://trello.com/invite/b/6763f98e151861c27b6a449d/ATTI63eea7e189237cef46592cb25221c22c27F8374B/groupie-tracker)

## Présentation du projet

Le projet **Groupie-Tracker** consiste à manipuler une API RESTful pour créer un site web interactif qui affiche des informations sur des groupes de musique, leurs concerts et leurs lieux. Ce projet permet également aux utilisateurs de rechercher des informations spécifiques à travers une barre de recherche et d'interagir avec le serveur pour récupérer des données dynamiques.

---

## Pré-requis

- Go (version 1.16 ou supérieure)
- Un navigateur web moderne
- Accès à Internet pour télécharger les dépendances (si nécessaire)

---

## Installation et démarrage

1. Clonez le repository :
    ```bash
    git clone https://ytrack.learn.ynov.com/git/cjohann/Groupie-tracker.git
    ```

2. Allez dans le répertoire du projet :
    ```bash
    cd Groupie-tracker/projet/Server
    ```


3. Lancez le serveur :
    ```bash
    go run main.go
    ```

4. Ouvrez votre navigateur et accédez à :
    ```bash
    http://localhost:8080
    ```

---

## Fonctionnalités

### 1. Affichage des informations sur les groupes

- Les utilisateurs peuvent explorer les informations sur les artistes et groupes, y compris leur nom, date de création, membres et image.
  
### 2. Recherche dynamique avec la barre de recherche (bonus)

- La barre de recherche permet aux utilisateurs de rechercher plus facilement des artistes. 


### 3. Affichage des concerts à venir

- Les utilisateurs peuvent voir les concerts passés et à venir, avec des informations sur les lieux et dates.

---



## Points importants

- Le backend est écrit en Go, et toutes les pages doivent être servies correctement (sans provoquer d'erreurs).
- Le code respecte les bonnes pratiques de développement, notamment la gestion des erreurs et la structure logique du projet.
- La barre de recherche offre une fonctionnalité bonus pour améliorer l'expérience utilisateur.
- La page d'accueil propose des artistes recommandés ainsi que différents types de musiques.
- Possibilité futur simple d'améliorations en nombre de données, grâce à deux fonctions supplémentaires qui vont chercher des données dans les API locations et dates.

---
