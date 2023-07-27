# Groupie Tracker Filters


## Table of Contents
1. [Description](#description)
2. [Authors](#authors)
3. [Usage:](#usage)
4. [Implementation details: algorithm](#implementation-details-algorithm)
5. [Testing](#testing)

### Description:
***
*Hi *Talent*!* 
Groupie Trackers localization is a continuity of the previous Groupie tracker projects which implements a geolocalization functionnality allowing to locate the concerts of artists collected in our API by converting addresses into geographic coordinates.
This project's server (back-end) is written using GO and the web documents (front-end) using HTML & CSS

#### Go,Html,Css:
***
- **GO**, also called Golang or Go language, is an open source programming language that Google developed.
* **HTML**,The HyperText Markup Language is the standard markup language for documents designed to be displayed in at web browser.
+ **CSS**,Cascading Style Sheets, form a computer language that describes the presentation of HTML documents.

### Authors:
***
+ Seynabou Niang (The artist Michel-Ange) -FRONTEND- https://learn.zone01dakar.sn/git/sniang
* Masseck Thiaw (the genius and **captain**) -BACKEND- https://learn.zone01dakar.sn/git/mthiaw
- Vincent Félix Ndour (Writer ,test-file) -README/TEST_FILE- https://learn.zone01dakar.sn/git/vindour

### Usage:
***
A little intro about how to install:
```
$ git clone https://learn.zone01dakar.sn/git/mthiaw/groupie-tracker-geolocatization
$ cd groupie-tracker-geolocalization
```

[def]: #usage-how-to-run
 go run main.go
## Implementation-details-algorithm:
***
Our program collects the information entered by the user in the localization bar, stores them in variables then performs manipulations on these datas while being conformed to the user specifications finally to search for an occurrences in our database and display it.

- **func Api_artists** Get all informations of the API's artist part from "https://groupietrackers.herokuapp.com/api/artists"

- **func api_locations** Gets datas from API's locations part --> "https://groupietrackers.herokuapp.com/api/locations".

+ **func api_dates** Gets datas from API's dates part --> "https://groupietrackers.herokuapp.com/api/dates" .

* **func api_relation** Gets datas from API's dates part --> "https://groupietrackers.herokuapp.com/api/dates" returns relations.

## Testing :
***
For the unit test file two file are used one that helps to test the API with using the Rest Client extension and the other for our programme.

- NB: For unit testing of the API the Rest Client extension must be installed on your PC.