# Skcli

## Introduction

Skcli (sky command line interface) is a command line tools for searching flights. It is designed to help you find flights when: 
* your don't have a specific destination in mind
* you know when you want a leave and return (day and time)
* you want to quickly compare flight accross different possible destinations

## User guide

The help menu is pretty straighforward

```shell
$ skcli -h
Usage of skcli:
  -country string
    	Country of the user (default "GB")
  -currency string
    	Currency of the user (default "GBP")
  -delay
    	Simulate random delay in query
  -departAfter string
    	Minimum departure time
  -direct
    	Direct flights only
  -from string
    	Origin Town/Airport (default "LON")
  -in string
    	Date of return/inbound flight (default "20161103")
  -keyFile string
    	API key provided by skyscanner (default "key")
  -language string
    	Language of the user (default "en-GB")
  -noCache
    	Do not read from cache
  -out string
    	Date of departure/outbound flight (default "20161101")
  -returnAfter string
    	Minimum return time
  -to string
    	selected destinations
```

The first step is to search for all possible destinations. Say I want to fly out from london (LON) the first weekend of december, leaving on friday the 2nd and coming back on sunday the 4th, with direct flights only:

```shell
$skcli -out 20161202 -in 20161204 -from LON -direct
```
The app will look for all possible destinations, starting by looking for all countries and then all towns. It will then display all the possible towns, sorted by price.

```shell
...
LPA 49 Gran Canaria Las Palmas https://www.skyscanner.net/transport/flights/LON/LPA/20161202/20161204/
LJU 48 Ljubljana https://www.skyscanner.net/transport/flights/LON/LJU/20161202/20161204/
FAO 48 Faro https://www.skyscanner.net/transport/flights/LON/FAO/20161202/20161204/
GLA 48 Glasgow International https://www.skyscanner.net/transport/flights/LON/GLA/20161202/20161204/
GDN 47 Gdansk https://www.skyscanner.net/transport/flights/LON/GDN/20161202/20161204/
WAW 47 Warsaw Chopin https://www.skyscanner.net/transport/flights/LON/WAW/20161202/20161204/
OVD 47 Asturias https://www.skyscanner.net/transport/flights/LON/OVD/20161202/20161204/
GOT 47 Gothenburg Landvetter https://www.skyscanner.net/transport/flights/LON/GOT/20161202/20161204/
PMI 47 Palma - Majorca https://www.skyscanner.net/transport/flights/LON/PMI/20161202/20161204/
MXP 46 Milan Malpensa https://www.skyscanner.net/transport/flights/LON/MXP/20161202/20161204/
PLQ 46 Palanga International https://www.skyscanner.net/transport/flights/LON/PLQ/20161202/20161204/
TAT 45 Poprad-Tatry https://www.skyscanner.net/transport/flights/LON/TAT/20161202/20161204/
KTW 44 Katowice https://www.skyscanner.net/transport/flights/LON/KTW/20161202/20161204/
GVA 44 Geneva https://www.skyscanner.net/transport/flights/LON/GVA/20161202/20161204/
AMS 43 Amsterdam https://www.skyscanner.net/transport/flights/LON/AMS/20161202/20161204/
CGN 43 Cologne https://www.skyscanner.net/transport/flights/LON/CGN/20161202/20161204/
INV 43 Inverness https://www.skyscanner.net/transport/flights/LON/INV/20161202/20161204/
LUX 42 Luxembourg https://www.skyscanner.net/transport/flights/LON/LUX/20161202/20161204/
AGP 42 Malaga https://www.skyscanner.net/transport/flights/LON/AGP/20161202/20161204/
CPH 39 Copenhagen https://www.skyscanner.net/transport/flights/LON/CPH/20161202/20161204/
POZ 39 Poznan https://www.skyscanner.net/transport/flights/LON/POZ/20161202/20161204/
ALC 38 Alicante https://www.skyscanner.net/transport/flights/LON/ALC/20161202/20161204/
SKP 38 Skopje https://www.skyscanner.net/transport/flights/LON/SKP/20161202/20161204/
BRQ 38 Brno-Turany https://www.skyscanner.net/transport/flights/LON/BRQ/20161202/20161204/
KSC 37 Kosice https://www.skyscanner.net/transport/flights/LON/KSC/20161202/20161204/
WRO 37 Wroclaw https://www.skyscanner.net/transport/flights/LON/WRO/20161202/20161204/
TRN 31 Turin https://www.skyscanner.net/transport/flights/LON/TRN/20161202/20161204/
SZZ 28 Szczecin Goleniow https://www.skyscanner.net/transport/flights/LON/SZZ/20161202/20161204/
BSL 16 Basel Mulhouse Freiburg https://www.skyscanner.net/transport/flights/LON/BSL/20161202/20161204/
158 results
```
I can then from this list chose a subset of destination I want to explore. Say I'm considering  When doing so I can add additional filter. Actually I want to leave after 4pm and come back after 7pm.

```shell
skcli -out 20161202 -in 20161204 -from LON -direct -to BSL,CGN -departAfter 1600 -returnAfter 1900
```

This can take a little bit of time as specific queries take more time. After a while I can see all the matching flights sorted by price:

```shell
LCY=>BSL 1 2016-12-02 19:15:00 +0000 UTC 2016-12-02 21:50:00 +0000 UTC BSL=>LGW 1 2016-12-04 21:10:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 226.85
LHR=>BSL 1 2016-12-02 19:10:00 +0000 UTC 2016-12-02 21:45:00 +0000 UTC BSL=>LGW 1 2016-12-04 21:10:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 219.07
LCY=>BSL 1 2016-12-02 19:15:00 +0000 UTC 2016-12-02 21:50:00 +0000 UTC BSL=>LTN 1 2016-12-04 20:55:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 203.27
LHR=>CGN 1 2016-12-02 15:50:00 +0000 UTC 2016-12-02 18:15:00 +0000 UTC CGN=>STN 1 2016-12-04 19:45:00 +0000 UTC 2016-12-04 19:55:00 +0000 UTC 199.98
LHR=>BSL 1 2016-12-02 19:10:00 +0000 UTC 2016-12-02 21:45:00 +0000 UTC BSL=>LTN 1 2016-12-04 20:55:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 194.48
STN=>CGN 1 2016-12-02 19:35:00 +0000 UTC 2016-12-02 21:45:00 +0000 UTC CGN=>STN 1 2016-12-04 19:45:00 +0000 UTC 2016-12-04 19:55:00 +0000 UTC 179.98
STN=>CGN 1 2016-12-02 16:30:00 +0000 UTC 2016-12-02 18:40:00 +0000 UTC CGN=>STN 1 2016-12-04 19:45:00 +0000 UTC 2016-12-04 19:55:00 +0000 UTC 179.98
LHR=>CGN 1 2016-12-02 20:00:00 +0000 UTC 2016-12-02 22:25:00 +0000 UTC CGN=>STN 1 2016-12-04 19:45:00 +0000 UTC 2016-12-04 19:55:00 +0000 UTC 179.98
LCY=>BSL 1 2016-12-02 19:20:00 +0000 UTC 2016-12-02 21:55:00 +0000 UTC BSL=>LTN 1 2016-12-04 20:55:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 170.78
LHR=>BSL 1 2016-12-02 19:10:00 +0000 UTC 2016-12-02 21:45:00 +0000 UTC BSL=>LGW 1 2016-12-04 21:10:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 150.27
LHR=>BSL 1 2016-12-02 14:00:00 +0000 UTC 2016-12-02 16:35:00 +0000 UTC BSL=>LGW 1 2016-12-04 21:10:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 150.27
LHR=>BSL 1 2016-12-02 19:10:00 +0000 UTC 2016-12-02 21:45:00 +0000 UTC BSL=>LTN 1 2016-12-04 20:55:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 125.68
LHR=>BSL 1 2016-12-02 14:00:00 +0000 UTC 2016-12-02 16:35:00 +0000 UTC BSL=>LTN 1 2016-12-04 20:55:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 125.68
LGW=>BSL 1 2016-12-02 18:05:00 +0000 UTC 2016-12-02 20:40:00 +0000 UTC BSL=>LGW 1 2016-12-04 21:10:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 94.98
LGW=>BSL 1 2016-12-02 18:05:00 +0000 UTC 2016-12-02 20:40:00 +0000 UTC BSL=>LTN 1 2016-12-04 20:55:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 73.98
LGW=>BSL 1 2016-12-02 14:00:00 +0000 UTC 2016-12-02 16:35:00 +0000 UTC BSL=>LGW 1 2016-12-04 21:10:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 72.98
LGW=>BSL 1 2016-12-02 14:00:00 +0000 UTC 2016-12-02 16:35:00 +0000 UTC BSL=>LTN 1 2016-12-04 20:55:00 +0000 UTC 2016-12-04 21:40:00 +0000 UTC 51.98
17 results
```

I can then decide that base is the cheapest option for me, with a gets flight from gatwick at 72.98.

## Installation

- Skcli is writen in go. 
- It uses the [Skyscanner](http://skycanner.net) [api](http://business.skyscanner.net/portal/en-GB/Documentation/ApiOverview). 
- To use the API you'll need a key. 
- Most of the code is in [sklib](https://github.com/arthurandres/sklib)
- It requires a few dependencies that are trivial to install
  - [bolt](https://github.com/boltdb/bolt) for caching results
  - [testify](github.com/stretchr/testify/assert) for testing

So quite simply:
```
go get github.com/arthurandres/skcli
go get github.com/arthurandres/sklib
go get github.com/boltdb/bolt
go get github.com/stretchr/testify/assert 
go install github.com/arthurandres/skcli
```
Then ```cd``` to where you want to run the app. Create a text file called ```key``` and just put your skyscanner API key value in there. 

Run ```$GOPATH/bin/skcli -h``` and voila.





