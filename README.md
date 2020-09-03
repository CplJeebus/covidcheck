# covidcheck

Simple app to track various covid-19 stats 

## Usage 

### Refresh\Init data 

Script Version
`check-ecdc.sh f` 

Go Version
`check-ecdc -f`


### Get 14 day average of new cases per 100k of population

Shell version

`check-ecdc.sh [Number of Days] [Country 1] [Country 2] [Country.....` 

`./check-ecdc.sh 5 IE DE` 

Go Version
`./check-ecdc -n [Number of Days] [Country 1] [Country 2] [Country.....`

```
18.27	IE	13/08/2020
17.78	IE	12/08/2020
17.86	IE	11/08/2020
16.94	IE	10/08/2020
15.80	IE	09/08/2020
14.62	DE	13/08/2020
13.96	DE	12/08/2020
13.31	DE	11/08/2020
12.91	DE	10/08/2020
12.79	DE	09/08/2020
```

### Get new cases in the previous _n_ days  


`check-ecdc.sh n [Number of Days] [Country 1] [Country 2] [Country.....` 

`./check-ecdc.sh n 5 IE DE` 

```
37	IE	13/08/2020
33	IE	12/08/2020
56	IE	11/08/2020
68	IE	10/08/2020
174	IE	09/08/2020
1445	DE	13/08/2020
1226	DE	12/08/2020
966	DE	11/08/2020
436	DE	10/08/2020
555	DE	09/08/2020
```

Go Version
`./check-ecdc -n [Number of Days] -c [Country 1] [Country 2] [Country.....`


### Get deaths in the previous _n_ days  


`check-ecdc.sh d [Number of Days] [Country 1] [Country 2] [Country.....` 

`./check-ecdc.sh d 5 IE DE` 

Go Version
`./check-ecdc -n [Number of Days] -d [Country 1] [Country 2] [Country.....`


```
1	IE	13/08/2020
1	IE	12/08/2020
0	IE	11/08/2020
0	IE	10/08/2020
0	IE	09/08/2020
4	DE	13/08/2020
6	DE	12/08/2020
4	DE	11/08/2020
1	DE	10/08/2020
1	DE	09/08/2020
```

##Go Version Only##

###Get new deaths per million per day in the previous _n_ days  

`./check-ecdc -n [Number of Days] -dm [Country 1] [Country 2] [Country.....`

```
0.0000	IE	31/08/2020
0.0000	IE	30/08/2020
0.0000	IE	29/08/2020
0.0000	IE	28/08/2020
0.0000	IE	27/08/2020
0.0361	DE	31/08/2020
0.0723	DE	30/08/2020
0.0120	DE	29/08/2020
0.0361	DE	28/08/2020
0.0602	DE	27/08/2020
```

###Get new cases per million per day in the previous _n_ days  

`./check-ecdc -n [Number of Days] -cm [Country 1] [Country 2] [Country.....`

```
8.1562	IE	31/08/2020
28.7506	IE	30/08/2020
25.6921	IE	29/08/2020
18.3515	IE	28/08/2020
33.0326	IE	27/08/2020
7.3477	DE	31/08/2020
9.4556	DE	30/08/2020
17.8152	DE	29/08/2020
18.9233	DE	28/08/2020
18.1524	DE	27/08/2020
```

###Create a graph of any of the above###

This will create a file `points.png` in you current directory
`/check-ecdc -cm -o plot ie de`


###Create a graph of any of the above###

Output csv formated to stdout.

`./check-ecdc -c -o csv ie de`

```
Date,ie,de
31/08/2020,40,610
30/08/2020,141,785
29/08/2020,126,1479
28/08/2020,90,1571
27/08/2020,162,1507
```


## Notes 

### Where does the data come from. 

[Here](https://www.ecdc.europa.eu/en/publications-data/download-todays-data-geographic-distribution-covid-19-cases-worldwide) basically it is the daily update from the European Centre for Disease Prevention and Control.

### Some strangness in the data 

* For some reason some of the data points are weird for instance negative numbers 
* The Country code are reasonably consistent with [ISO 3166](https://www.iban.com/country-codes) however there are some funnies for example `UK` rather than `GB`.
`


# Todo

### Shell Script
Probably won't do much more but;
  * A `usage` function would be nice
  * Better arg handling
  * Prereq check for things like `jq`

### Go Version
Where to start! Here I suppose;
  * Get functionality parity to the shell script *prioity 1* - **DONE**
  * Refactor to get modules working properly and figure out my hinky `$GOPATH` issues. **DONE** _Note:_ My vim config still insist my GOPATH is wrong!
  * Refactor some more for the sake of it **DONE**
  * Report usage and make interface better the flags are confusing **DONE** or at least better.
  * Better arg handling logic, especially aroud required args and defaults and stuff. **DONE** Again its better.
  * Arg validation for Country codes, I actually like the silent fail but a simple `toupper` would be cool **DONE** 
  * Output validation. This might require some research to find the background, but app should filter nonsensical output. _"-48" deaths UK._ I think we have a bigger problem than covid!
  * Nicer formatting tabulating via date for instance. **STARTED** Now can output in _csv_ format 
  * Add\Improve output formats csv,json etc. **STARTED** See above.
  * Fiddle with the scales of the graphs. 
  * Fix titles on the graphs and general clean up **DONE**
  * Allow different file names for the output plots.
  * Do per capita (or deaths per 10^6) calcs for new cases and deaths! **DONE** _-dm_ and _-cm_ will give you cases per 10^6


