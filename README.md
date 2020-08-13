# covidcheck

Simple app to track various covid-19 stats 

## Usage 

### Refresh\Init data 

`check-ecdc.sh f` 

### Get 14 day average of new cases per 100k of population


`check-ecdc.sh [Number of Days] [Country 1] [Country 2] [Country.....` 
`check-ecdc.sh 5 IE DE` 

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
`check-ecdc.sh n 5 IE DE` 

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

### Get deaths the previous _n_ days  


`check-ecdc.sh d [Number of Days] [Country 1] [Country 2] [Country.....` 
`check-ecdc.sh d 5 IE DE` 

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







