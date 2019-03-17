<template>
    <div class="table-responsive">
        <table class="table-hover">
            <thead>
            <tr>
                <th>ISIN</th>
                <th>Name</th>
                <th>Amount</th>
                <th>Current Value</th>
                <th>Percentage of portfolio</th>
            </tr>
            </thead>
            <tbody>
            <router-link v-for="holding in this.portfolio" :key="holding.name" :to="{name: 'CompanyPage', params: {isin: holding.isin} }" tag="tr">
                <td>{{ holding.isin }}</td>
                <td>{{ holding.name }}</td>
                <td>{{ holding.scaled.toFixed(4) }}</td>
                <td>£{{ holding.currentValue.toFixed(2) }}</td>
                <td>{{ parseInt(percentageOfPortfolio(holding.currentValue.toFixed(2)))}}</td>
            </router-link>
            </tbody>
        </table>
        <div>
        </div>
    </div>
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
    created () {
        this.getData();
    }
    percentageOfPortfolio (accountValue) {
        return (accountValue / this.$store.getters['returnPortfolioTotal']) * 100;
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
                    this.$store.dispatch('portfolioTotalHandler', currentValue);
                    let newObject = {'isin': holding.isin, 'unscaled': holding.amount.unscaled, 'exponent': holding.amount.exponent, 'scaled': scaled, 'name': accountName, 'price': accountPrice, 'currentValue': currentValue};
                    this.portfolio.push(newObject);
                    console.log(this.portfolio);
                }
            this.$store.dispatch('portfolioHoldingsHandler', this.portfolio);
        }).then(() => {
            this.loaded = true;
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
