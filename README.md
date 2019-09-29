# Féminicides·Info - CLI

Outil en ligne de commande (CLI) pour le téléchargement et la conversion des données
des féminicides depuis le format KML (des cartographies Google Maps) vers le
format JSON (utilisables pour différents rendu web).

Les données originales des cartographies Google Maps sont produites par les
autrices du groupe Facebook [Féminicides par compagnons ou
ex](https://www.facebook.com/feminicide/). .


## Installation

Assurez-vous d'avoir installé Go (version 1.12 ou supérieure) sur votre
ordinateur et que $GOPATH/bin soit bien configuré dans le $PATH.

Depuis la ligne de commande, tapez :

    $ git clone https://github.com/glenux/feminicides-info-cli/
    $ go install  ./...


## Usage

### Vue d'ensemble

Commande principale

    Usage:
      fi-cli [command]

    Available Commands:
      convert     Convert feminicide-related KML into JSON
      fetch       Fetch KML data for feminicide in specified year
      help        Help about any command

    Flags:
      -h, --help   help for fi-cli

    Use "fi-cli [command] --help" for more information about a command.

Sous-commande `fetch` :

    Fetch KML data for feminicide in specified year

    Usage:
      fi-cli fetch YEAR [flags]

    Flags:
      -h, --help            help for fetch
      -o, --output string   Output file name for JSON (default "-")

Sous-commande `convert` :

    Convert feminicide-related KML into JSON (long)

    Usage:
      fi-cli convert SOURCE_KML [flags]

    Flags:
      -h, --help            help for convert
      -o, --output string   Output file name for KML (default "-")


### Téléchargement des données

Pour télécharger les données des féminicides au format KML (de 2016 à 2019) :

    $ fi-cli fetch YEAR
    
Par exemple pour 2016 :

    $ fi-cli fetch -o feminicides-2016 2016
    Fetching KML for year 2016
    Fetching from https://www.google.com/maps/...
    SUCCESS

    $ ls
    feminicides-2016.kml

### Conversion des données

Pour convertir les fichiers KML au format JSON :

    $ fi-cli convert FICHIER_KML FICHIER_JSON

Par exemple pour 2016 :

    $ fi-cli convert -o feminicides-2016.json feminicides-2016.kml

    $ ls
    feminicides-2016.kml feminicides-2016.json


## Licence et droit d'auteur

[Féminicides·Info CLI](https://github.com/glenux/feminicides-info-cli) est un projet open source sous licence LGPL-3.

Auteur : Glenn Y. Rolland ([@glenux](https://twitter.com/glenux))

## Sponsors et financement

[Féminicides·Info](https://github.com/glenux/feminicides-info-cli) est un projet indépendant dont le développement continu est rendu possible grâce au soutien de ses mécènes.

Si vous souhaitez vous joindre à eux et soutenir le travail de son auteur, il suffit de participer avec ce lien :

&gt;&gt; [Devenez mécène ou sponsor sur Patreon](https://www.patreon.com/glenux) &lt;&lt;

