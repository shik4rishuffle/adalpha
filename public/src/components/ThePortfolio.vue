<template>
        <table class="table portfolio__table table-hover table-responsive">
            <thead class="portfolio__table-head">
            <tr>
                <th class="portfolio__table-cell">ISIN</th>
                <th class="portfolio__table-cell--name">Name</th>
                <th class="portfolio__table-cell">Amount</th>
                <th class="portfolio__table-cell">Value</th>
            </tr>
            </thead>
            <tbody class="portfolio__table-body">
            <router-link v-for="holding in this.portfolio" :key="holding.name" :to="{name: 'CompanyPage', params: {isin: holding.isin} }" tag="tr">
                <td class="portfolio__table-cell">{{ holding.isin }}</td>
                <td class="portfolio__table-cell--name">{{ holding.name }}</td>
                <td class="portfolio__table-cell">{{ holding.scaled.toFixed(4) }}</td>
                <td class="portfolio__table-cell">£{{ holding.currentValue.toFixed(2) }}</td>
            </router-link>
            </tbody>
        </table>
</template>

<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';

  @Component
  export default class ThePortfolio extends Vue {
    name: 'ThePortfolio';
    holdings: any = [];
    holdingValues: any = [];
    portfolio: any = [];
    portfolioTotal: number = 0;
    loaded: boolean = false;
    currentValueTally:number = 0;

      mounted () {
        this.currentValueTally = 0;
        this.getData();
    }

    // methods
    getData () {
    fetch('http://localhost:8080/portfolios/' + this.$store.getters['returnUsername'] + '/accounts/', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    },
    credentials: 'include'
    }).then(response => {
          return response.json();
    })
    .then((res) => {
        for (let value of JSON.parse(JSON.stringify(res))) {
            this.$store.dispatch('portfolioNameHandler', value.name);

            for (let [holding, amount] of Object.entries(value.holdings)) {
                let newHolding = {'isin': holding, 'amount': amount};

                this.holdings.push(newHolding);
            }

            this.$store.dispatch('portfolioTradesHandler', value.trades);
        }
    }).then(() => {
        fetch('http://localhost:8080/funds/', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: 'include'
        }).then(response => {
            return response.json();
        })
            .then((res) => {
                this.holdingValues = res;
            }).then(() => {
                for (let holding of this.holdings) {
                    const scaled = holding.amount.unscaled * Math.pow(10, holding.amount.exponent);
                    let accountName:string;
                    let accountPrice:number;
                    this.holdingValues.filter(function (item) {
                        return (item.isin === holding.isin);
                    }).map(function (item) {
                        accountName = item.name.toString();
                        accountPrice = parseInt(item.price.toString());
                        });
                    const currentValue = (accountPrice * scaled) / 100;
                    this.currentValueTally += currentValue;
                    let newObject = {'isin': holding.isin, 'unscaled': holding.amount.unscaled, 'exponent': holding.amount.exponent, 'scaled': scaled, 'name': accountName, 'price': accountPrice, 'currentValue': currentValue};
                    this.portfolio.push(newObject);
                }
            this.$store.dispatch('portfolioTotalHandler', this.currentValueTally);
            this.$store.dispatch('portfolioHoldingsHandler', this.portfolio);
        }).then(() => {
            this.loaded = true;
            this.$store.dispatch('portfolioLoadedHandler', this.loaded);
        })
            .catch((error) => {
                console.log(error);
            });
        }
    )
      .catch((error) => {
        console.log(error);
      });
    }
  };
</script>

<style lang="scss" scoped>
    .portfolio__table {
        display: table;
    }
    .portfolio__table-cell {
        text-align: center;
        padding:3px;
    }
    .portfolio__table-cell--name {
        @extend .portfolio__table-cell;
        display: none;
    }

    .table {
        text-align: center;
        tr {
            text-align: center;
        }
        th, td {
            padding: 5px 5px;
            text-align: center;
            vertical-align: middle;
        }
    }
    @media only screen and (min-width: 768px) {
        .portfolio__table-cell--name {
            display: table-cell;
        }
    }
</style>
