# Forum

## Présentation et Membres du groupe

Groupe : **FORUMIX**

- Johann Chiarotto
- Adrien Dulou
- Baptiste Renou
- Ulysse Prevost Lacaze

## Description

**Forum de dilemmes**

*Forum* est une plateforme de discussion en ligne permettant à des utilisateurs de s'inscrire, de créer des posts, de commenter, d'aimer/disliker des posts et commentaires, et de filtrer le contenu selon différentes catégories. Ce projet utilise du Go pour la logique backend et SQLite pour la base de données.

## Objectifs du projet

- Permettre la communication entre utilisateurs via des posts et commentaires.
- Associer des catégories aux posts.
- Ajouter des fonctionnalités de likes et dislikes pour les posts et commentaires.
<!--- Implémenter un mécanisme de filtrage des posts.-->
- Assurer la sécurité des données des utilisateurs (inscription, authentification, etc.).

## Fonctionnalités

- **Inscription et connexion** : Les utilisateurs peuvent s’inscrire, se connecter et maintenir une session active grâce aux cookies.
- **Création de posts et commentaires** : Les utilisateurs enregistrés peuvent créer des posts et ajouter des commentaires.
- **Likes et Dislikes** : Les utilisateurs peuvent aimer ou ne pas aimer des posts et des commentaires.
- **Filtrage des posts** : Filtrer les posts par catégories, par les posts créés par l'utilisateur et les posts aimés par l'utilisateur.
- **Sécurisation des mots de passe** : Les mots de passe sont cryptés avant d’être stockés.
- **Base de données SQLite** : Utilisation d'une base de données SQLite pour stocker les données du forum.
- **Modifications** : Changement de mot de passe et de photo de profil ( a travers une adresse de l'image )

## Points clefs

- **Changer sa photo de profil** : Aller chercher une image dans google photos, cliquer sur l'image voulu de sorte a l'avoir sur le coté de votre écran et cliquer sur "Copier l'adresse de l'image"
- **Modération** : Administrateur en possibilité de supprimer commentaires, posts et comptes en cas de non respect d'autruit

## Prérequis

Avant de commencer, assurez-vous d’avoir les éléments suivants installés sur votre machine :

- **Golang** 
- **Un accès internet**
  
## Installation

1. Clonez le dépôt :

```bash
git clone https://ytrack.learn.ynov.com/git/cjohann/FORUM.git
```

2. Accédez au répertoire du projet :

```bash
cd Projet/Server/
```

3. Exécutez le serveur avec :

```bash
go run main.go
```

## Sécurité

- **Cryptage des mots de passe** : Les mots de passe sont cryptés avec bcrypt avant d’être stockés.
- **Sessions et cookies** : Un cookie est créé pour chaque utilisateur connecté et expirera après un délai de 24 heures.
- **Protection contre les attaques courantes** : Le projet prend en compte les meilleures pratiques de sécurité, telles que la protection contre les injections SQL et la gestion des erreurs.

