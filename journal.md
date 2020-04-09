# Daily thread

## Semaine 1 (02/03 - 06/03)
Rencontre avec les équipes, découverte des concepts de batsim, nix, kubernetes.

Etat de l'art.

Formation Go.

## Semaine 2 (09/03 - 13/03)
Etat de l'art.

Fin formation Go.

Tutos Nix.

Bcp de discussions autour de l'approche à prendre -> Proof of concept sur le
bashScheduler, avec une implémentation bas niveau directement en tant qu'API.

## Semaine 3
### 16/03
- Mise en place de l'environnement de travail sur ma machine à la maison.
- Tuto batsim "doing a reproducible environnment"
- Tuto api en go, avec un bon rappel des paramètres d'une API. Premières lignes de test, pour tester la bonne communication entre bat-kube et ./scheduler.sh

### 17/03
- Etabli des réponses mock du simulateur au scheduler.
- Utilisation de packr pour la gestion des fichiers. Essayé d'utiliser Pkger, pas réussi
- Début de traitement de requête post pour binder. Le scheduler finit toujours sur une erreur interne qu'il faut débugger

### 18/03
- Version fonctionnelle d'une api qui ne renvoie que des mocks prédéfinis.
- Exploré la piste des mock de cluster. Peut être utilisé pour rediriger les points de l'api que l'on ne veut pas implémenter.
- Next step : bashScheduler avec des mocks -> Probleme conceptuel, les mocks sont utiles côté client.
- Etat de l'art sur les mock servers / nodes, rédaction avantages/inconvénients approches d'architecture, proposé une extension de la solution custom client.

### 19/03
- Réunion avec Olivier. Next step : simulateur de bout en bout pour le bashScheduler
- Exploré la doc de batsim (Plateforme, workloads, events). Mise en parallèle des entités batsim et kube.
- Nouvelle organisation des repos

### 20/03
- Réunion avec michael et olivier, mise au point collective. Mises au clair sur les traductions pods/jobs, clarifications de qlq points sur la prochaine étape incrémentale.
- Expérimentations avec les objets kube et le fake client go

## Semaine 4
### 23/03
- Parvenu à parser un deployment.yaml en objet v1
- Remise au point avec michael sur l'architecture du simulateur
- mise en place de multipass + k3s pour lancer des tests locaux
- lecture et discussions sur les packages / modules Go

### 24/03
- Réflexions sur l'architecture du code
- Lectures go sur la gestion et propagation des erreurs
- Deserialisation des compute_resources en node
- Communication des objets entre broker et api

### 25/03
- Tenté de sérialiser des objets kube -> besoin de générer les fonctions DeepCopyObject [reference](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/)
- Nouvelle piste à explorer : la librairie [apiserver](https://github.com/kubernetes/apiserver)
- réunion hebdo avec Olivier
- Doc et préparation de demo

### 26/03
- Demo
- Corrigé le bug des runtime.Object
- Ajouté des types Batsim.

### 27/03
- Réu avec michael et Olivier. Prochaines étapes : 1) finir la v0 end to end 2) Diagramme de séquence
3) Etudier la faisabilité de go-swagger 4) Problème du temps
- Sérialisation propre de SIMULATION_BEGINS en struct go

## Semaine 5
### 30/03
- Encore de la sérialisation et du travail sur les types liés à batsim. Très certainement d'autres retouches qui arrivent.
- Réflexions sur une autre architecture possible de batkube, ne définissant pas une nouvelle api
- Réflexion sur la gestion de la conversion pod <-> job

### 31/03
- Rapide appel avec Michael à propos de l'idée d'architecture -> décidée pas viable. Doc mise à jour.
- re-boulot sur les types Profile. Abandon de l'interface Profile (pour le moment).
- début d'implémentation de l'idée d'une convertion pod <-> job basée sur des .yml, abandon (pas du tout une bonne idée)
- Remise en cause du rôle des différents éléments de batkube nottament par rapport à la concurrence. Mise au propre
d'une vraie organisation plus adaptée, modulaire, claire. Implémentation demain.

### 01/04
- Changé l'architecture de batkube : le traitement des messages se fait dans l'api et non plus daans le broker. Code plus clair.
- Support de JOB_SUBMITTED
- Support du query paramater fieldSelector pour /pods
- Expériences sur la synchro -> Stocker les objets v1 en variable globale de l'api ne fonctionne pas. Il faut trouver une autre solution.

### 02/04
- Première version fonctionnelle de bout en bout. Améliorations sur les logs. README mis à jour.
- Premier jet sur un diagramme de séquence. Discuté avec Olivier via telegram sur la synchro du temps et l'architecture du code.
- Pas fait grand chose d'autre, fuite d'eau à la maison qui requiert mon attention

### 03/04
- Réu avec michael et olivier
- Recherches sur les interactions entre kube et scheduler (watch, informers)
- Manips en essayant de lancer des expériences avec un scheduler utilisant go -> problème de la config kube.

### 06/04
- Recherches sur l'authentification avec l'api server : config kube, https
- Expériences avec swagger api. Authentification réussie avec les certificats ssl, mais le serveur demande des credentials dont je n'ai aucune idée...

### 07/04
- Passé la journée à tenter de débugger la swagger api. Pas réussi -> handler mytérieux qui renvoie une 401 unauthorized
- Ce qui a été tenté : déboggage avec curl (et postman), avec et sans certificats, implémentation
basique de BearerTokenAuth, déboggage pas à pas de l'api, mode debug de swagger (DEBUG=1)

### 08/04
- Solution au problème trouvé par Olivier
- Encore beaucoup de temps passé sur l'authentification et à comprendre le fonctionnement de swagger api
- Slides pour la démo

### 09/04
- matin: demo
- Testé le code exemple de swagger pour les json en chunk
- Migré le code de l'api swagger à la racine du projet, enlevé l'ancien package api
