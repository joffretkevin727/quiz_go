# Projet Quiz - Phase 2 (Extension Web)

Ce projet est une application web de quiz interactive développée en **Go**. Il permet de tester ses connaissances sur **Git** et **Golang** à travers une interface web moderne.

---

## 🚀 Fonctionnalités
* **Affichage dynamique** : Les questions s'affichent une par une pour une meilleure expérience utilisateur.
* **Validation en temps réel** : Le programme indique si la réponse saisie est correcte ou incorrecte après chaque validation.
* **Gestion du score** : Le score est calculé tout au long du quiz et affiché à la fin.
* **Interface Web** : Utilisation de templates HTML et de fichiers CSS pour une présentation propre.

---

## 📁 Structure du projet
```text
quiz-go/
├── controller/
│   └── controller.go   # Logique métier et gestion des requêtes
├── router/
│   └── router.go       # Configuration des routes et ressources statiques
├── structure/
│   └── structure.go    # Définition du modèle de données (PageData)
├── static/
│   └── style.css       # Mise en page et design responsive
├── template/
│   └── home.html       # Interface utilisateur via Go Templates
└── main.go             # Point d'entrée de l'application

🛠️ Installation et Lancement
1. Cloner le dépôt :
    git clone https://github.com/joffretkevin727/quiz-go.git

2. Lancer le serveur :
    cd .\quiz-go\
    go run main.go

3. Accéder à l'application :
    Rendez-vous sur http://localhost:8080/home
