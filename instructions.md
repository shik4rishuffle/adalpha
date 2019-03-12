# Frontend Technical Test

Welcome to the AdAlpha frontend technical test. The purpose of this test is to
get an example of what you consider good quality code. As such you should
approach this in that vein and use it as an opportunity to show off and not
just meet the requirements with a working solution.


## The challenge

An investor has an account that has the following investments that have a total
value of`£220,000`. The account has the following assets in it.

| ISIN         | ASSET                                                             | % Of Portfolio |
|--------------|-------------------------------------------------------------------|---------------:|
| IE00B52L4369 | BlackRock Institutional Cash Series Sterling Liquidity Agency Inc |             20 |
| GB00BQ1YHQ70 | Threadneedle UK Property Authorised Investment Net GBP 1 Acc      |             20 |
| GB00B3X7QG63 | Vanguard FTSE U.K. All Share Index Unit Trust Accumulation        |             10 |
| GB00BG0QP828 | Legal & General Japan Index Trust C Class Accumulation            |              5 |
| GB00BPN5P238 | Vanguard US Equity Index Institutional Plus GBP Accumulation      |             15 |
| IE00B1S74Q32 | Vanguard U.K. Investment Grade Bond Index Fund GBP Accumulation   |             30 |


Your challenge is build an interface that displays these investments, buy new
investments and sell existing. For the purpose of this test an asset can be
bought and sold in GBP amounts only.

You should also consider adding any other functionality you deem appropriate.


## What we want to see

* A `Vue.js` based front end written with typescript.
* A `Git` repository with well structured commits.
* API calls to the supplied Go backend to retrieve data.
* A write up that explains your decisions and any assumptions.

You are free to use any other technologies or libraries that you like. Any
evidencing of skills and technologies on the job specification would be
advantageous.


## The Sample Backend

A backend has been provided to get you started. Feel free to modify and/or
refactor this to suit your needs and interpretation of the requirements. The
simplest way to get started is to install Docker. From the `server` directory,
to build, run, view logs and stop:

```
docker build -t adalpha-server .
docker run -d --rm --name adalpha-server -p 8080:8080 adalpha-server
docker logs adalpha-server
docker stop adalpha-server
```

The main endpoints patterns are:

```
/login - POST {username, password}, returns cookie
/funds/{isin} - fund collection or individual details
/portfolios/{name}/accounts/{name}/trades - GET at any level, POST to trades
```

For simplicity, usernames are simply portfolio names.

## Submitting

You can submit your application as a zip archvie to jobs@adalphasolutions.com
with the zip in the format `{first-name}.{surname}-{YYYY-MM-DD}.zip`. Be sure
to exclude any built binaries and ignored files.