# Feminicides Info - Cli

Outil en ligne de commande pour le téléchargement et la conversion des données
des féminicides depuis le format KML (des cartographies Google Maps) vers le
format JSON.(utilisables pour différents rendu web).

Les données originales des cartographies Google Maps sont produites par les
autrices du groupe Facebook [Féminicides par compagnons ou
ex](https://www.facebook.com/feminicide/). .


## Installation

Assurez-vous d'avoir installé Go (version 1.12 ou supérieure) sur votre
ordinateur.

Depuis la ligne de commande, tapez :

    $ make


## Usage

Pour télécharger les KML des féminicides :

    $ ./fi-cli fetch 2016
    Fetching KML for year 2016
    Fetching from https://www.google.com/maps/...
    SUCCESS

Pour convertir les KML au format JSON

    $ ./fi-cli convert doc-2019.kml doc-2019.json
    $ ls
    doc-2019.kml doc-2019.json


## Licence et droit d'auteur

[Féminicides.Info CLI](https://github.com/glenux/feminicides-info-cli) est un projet open source sous licence LGPL-3.

Auteur : Glenn Rolland.

<!-- 
## Sponsors et financeurs

[Féminicides.Info]( est un projet indépendant dont le développement continu est rendu possible
grâce au soutien de ses mécènes. Si vous souhaitez vous joindre à eux et soutenir le travail
de son auteur, n'hésitez pas participer :

[Devenez mécène ou sponsor sur Patreon](https://www.patreon.com/glenux)
-->
