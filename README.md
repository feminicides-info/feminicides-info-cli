# Feminicides Info - Cli

Outil de téléchargement et de conversion des données des féminicides depuis le
format KML (des cartographies Google Maps) vers le format JSON.(utilisables
pour différents rendu web).

Les données originales des cartographies Google Maps sont produites par les
autrices du groupe Facebook [Féminicides par compagnons ou
ex](https://www.facebook.com/feminicide/). .

## Installation

    $ make


## Usage

Pour télécharger les KML des féminicides 

    $ fi-cli fetch 2019
    $ fi-cli fetch 2018
    $ fi-cli fetch 2017
    $ fi-cli fetch 2018

Pour convertir les KML au format JSON

    $ fi-cli convert doc-2019.kml doc-2019.json



