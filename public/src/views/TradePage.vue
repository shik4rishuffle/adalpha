<template>
    <div class="trade-page container-fluid">
        <div class="row">
            <div class="col">
                <div class="row">
                    <div class="col flex-row">
                        <h1 class="h1 company-name">{{this.data.name}}</h1>
                        <h3 class="h3">ISIN: {{ this.data.isin }}</h3>
                        <h3 class="h3">Current Price: £{{(this.data['current-price'] / 100).toFixed(2)}}</h3>
                        <h3 class="h3">Current Holdings: {{this.currentHoldings.toFixed(2)}}</h3>

                    </div>
                </div>

                <div class="row justify-content-center">
                    <div class="col-2">
                        <div class="form-group d-flex"><h4 class="h4">Buy</h4></div>
                        <form class="buy-form">
                            <div class="form-group">
                                <label for="buyPrice">Price: £</label>
                                <input id="buyPrice" type="text" v-model="buyPrice" :placeholder="(this.data['current-price'] / 100)">
                            </div>
                            <div class="form-group">
                            <label for="buyAmount">Amount</label>
                            <input id="buyAmount" type="text" v-model="buyAmount">
                            </div>

                            <div class="button-holder form-group">
                                <div v-on:click="buyPercent(25)" class="btn btn-primary">25%</div>
                                <div v-on:click="buyPercent(50)" class="btn btn-primary">50%</div>
                                <div v-on:click="buyPercent(75)" class="btn btn-primary">75%</div>
                                <div v-on:click="buyPercent(100)" class="btn btn-primary">100%</div>
                            </div>
                            <div class="form-group">
                                <label for="buyTotal">Total: £</label>
                                <input disabled id="buyTotal" type="text" v-model="computedBuyTotal">
                            </div>
                            <div class="form-group d-flex justify-content-between">
                                <div class="btn btn-secondary">Clear</div>
                                <div class="btn btn-success">Buy</div>
                            </div>
                        </form>
                    </div>
                    <div class="col-2">
                        <div class="form-group d-flex justify-content-between align-items-center"><h4 class="h4">Sell</h4>
                            <small>{{this.currentHoldings}}</small>
                        </div>
                        <form class="sell-form">
                            <div class="form-group">
                                <label for="sellPrice">Price: £</label>
                                <input id="sellPrice" type="text" v-model="sellPrice" :placeholder="(this.data['current-price'] / 100)">
                            </div>
                            <div class="form-group">
                            <label for="sellAmount">Amount</label>
                            <input id="sellAmount" type="text" v-model="sellAmount">
                            </div>

                            <div class="button-holder form-group">
                                <div v-on:click="sellPercent(25)" class="btn btn-primary">25%</div>
                                <div v-on:click="sellPercent(50)" class="btn btn-primary">50%</div>
                                <div v-on:click="sellPercent(75)" class="btn btn-primary">75%</div>
                                <div v-on:click="sellPercent(100)" class="btn btn-primary">100%</div>
                            </div>
                            <div class="form-group">
                                <label for="sellTotal">Total: £</label>
                                <input disabled id="sellTotal" type="text" v-model="computedSellTotal">
                            </div>
                            <div class="form-group d-flex justify-content-between">
                                <div class="btn btn-secondary">Clear</div>
                                <div class="btn btn-primary" v-on:click="sellPost">Sell</div>
                            </div>
                        </form>
                    </div>
                </div>

                <div class="row collapse">
                    <div class="col">
                        <h3 class="h3">History</h3>
                        <table class="table table-hover">
                            <thead>
                            <tr>
                                <th scope="col">Date</th>
                                <th scope="col">Price</th>
                                <th scope="col">Status</th>
                            </tr>
                            <tr v-for="trade in this.data.history" :key="trade.name">
                                <td scope="col">{{trade.date}}</td>
                                <td scope="col">£{{(trade.price / 100).toFixed(2)}}</td>
                                <td scope="col">Processed</td>
                            </tr>
                            </thead>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';
    @Component
  export default class TradePage extends Vue {
        isin:any;
        data:any = [];
        currentHoldings:number = 0;
        buyAmount: number = 0;
        buyPrice: number = 0;
        sellAmount: number = 0;
        sellPrice: number = 0;

        created () {
            this.isin = this.$route.params.isin;

            fetch('http://localhost:8080/funds/' + this.isin, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: 'include'
            }).then(response => {
                return response.json();
            })
                .then((res) => {
                    console.log(res);
                    this.data = res;
                    this.getCurrentHoldings();
                })
                .catch((error) => {
                    console.log(error);
                });
        }

        // methods
        getCurrentHoldings () {
            const portfolio = this.$store.getters['returnPortfolioHoldings'];
            const thisIsin = this.data.isin;
            const find = portfolio.filter(function (item) {
                return (item.isin === thisIsin);
            }).map(function (item) {
                return item.scaled.toString();
            });
            this.currentHoldings = parseFloat(find);
        }
        sellPercent (percent) {
            this.sellAmount = (this.currentHoldings / 100) * percent;
        }
        buyPercent (percent) {
            this.buyAmount = (this.currentHoldings / 100) * percent;
        }
        sellPost () {
            fetch('http://localhost:8080/portfolios/' + this.$store.getters['returnUsername'] + '/accounts/' + this.$store.getters['returnPortfolioName'] + '/trades', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: 'include',

                body: JSON.stringify({
                    // @ts-ignore
                    Status: 'TradePending',
                    FundName: this.isin,
                    Amount: this.computedSellTotal
                    })
            }).then(response => {
                return response.json();
            })
                .then((res) => {
                    console.log(res);
                    this.data = res;
                    this.getCurrentHoldings();
                })
                .catch((error) => {
                    console.log(error);
                });
        }

        // computed
        get computedBuyTotal () {
            return (this.buyAmount * this.buyPrice).toFixed(4);
        }
        get computedSellTotal () {
            return (this.sellAmount * this.sellPrice).toFixed(4);
        }
  };
</script>
