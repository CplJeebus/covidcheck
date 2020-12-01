# covidcheck

Simple app to track various covid-19 stats 

## Usage 

### Refresh\Init data 

`check-ecdc -f`
This has been changed and will now refresh the data file if it is older than 8 hours


### Get 14 day comulative of new cases per 100k of population

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


`./check-ecdc -n [Number of Days] -c [Country 1] [Country 2] [Country.....`

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

### Get deaths in the previous _n_ days  

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

### Get new deaths per million per day in the previous _n_ days  

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

### Get new cases per million per day in the previous _n_ days  

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

### Create a graph of any of the above

This will create a file `points.png` in you current directory
`/check-ecdc -cm -o plot ie de`

This is now the default behaviour


### Output in csv format of any of the above

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

### Events 
Using the `events.yaml` you can specify the dates of key events per country, for example the starting and easing of lockdowns. See the format below. 

```
events:
  - event: "First lockdown"
    geoid: "ie"
    date: "27/03/2020"
  - event: "Phase 1 easing"
    geoid: "ie"
    date: "18/05/2020"
```

Events are added by default add `-xe` to toggle off. 

## Notes 

### Where does the data come from. 

[Here](https://www.ecdc.europa.eu/en/publications-data/download-todays-data-geographic-distribution-covid-19-cases-worldwide) basically it is the daily update from the European Centre for Disease Prevention and Control.

For US states we get the data from the [CDC](https://data.cdc.gov/api/views/9mfq-cb36/rows.csv?accessType=DOWNLOAD), but this a big *pain!* . This site seems to have the most comprehensive data in consistent format. Unlike the ECDC this data is a bit difficult to handle. The ECDC data comes in a single consistent file, that is updated daily. CDC data seems to be updated inconsistently (no covid on the weekends!) And this is different state by state. Also one of the key data points for the E CDC is _the cumulative new cases over 14 days per 100K of population_ as defined [here](https://www.ecdc.europa.eu/en/publications-data/download-todays-data-geographic-distribution-covid-19-cases-worldwide). 

```
The formula to calculate the 14-day cumulative number of reported COVID-19 cases per 100 000 population is  (New cases over 14 day period)/Population)*100 000.
```

So to compare like with like I need to calculate this for the US states, and have created a `method` to do so.  However very few US states report data this way so it is hard to validate my calculations with these States. Long story I have a sore head!! And I’m not convinced that the data is correct, but it is _truthy_ and that seems to be good enough these days.  

Either way the point of this exercise was to practice `go` and the fact that there is no functional national government and no two US states can seemingly agree as to whether or not to festoon their health advice sites with adverts or not, never mind give good data, means I’m learning a lot! 

### Some strangness in the data 

* For some reason some of the data points are weird for instance negative numbers 
* The Country code are reasonably consistent with [ISO 3166](https://www.iban.com/country-codes) however there are some funnies for example `UK` rather than `GB`.
* See rant above about the US.


# Todo

Where to start! Here I suppose;
## Functionality 
  * See if I can add smoothing functions to the -dm and -cm flags. `us` data is especially bad seems the CDC does not work Sundays! 
  * Fiddle with the scales of the graphs. Would like to do a lot with the graphs might be the next big change
  * Allow different file names for the output plots.
  * Some cross platform way to open the graphs automatically.
  * Containerize this and add a front end?? 

## Codey stuff
  * Refactor 
  * To much messing around with files 
  * Clean up the flag stuff. 
  * figure out why messing with the CDC file is so slow 
  * Maybe insert the lot into a DB or key value store.  
    - Can mess with SQLlite locally. 
    - Not sure about a KV store 
    - This is a learning exercise.


### Shell Script
Deprecated, will leave the script there for ~the craic~ interest.


