<template>
    <div class="trade-page col-lg-3">
                <div class="row">
                    <div v-if="!showBuy && !showSell" class="col toggle-container">
                        <div class="btn btn-outline-primary" v-on:click="showBuyPane()">Buy</div>
                        <div class="btn btn-outline-primary" v-on:click="showSellPane()">Sell</div>
                    </div>
                        <form v-if="showBuy" class="col buy-form">
                            <div class="form-group d-flex">
                                <h4 class="toggle active">Buy</h4>
                            </div>
                            <div class="form-group">
                                <label for="buyPrice">Price: £</label>
                                <input id="buyPrice" type="text" v-model="buyPrice" :placeholder="(this.data['current-price'] / 100)">
                            </div>
                            <div class="form-group">
                            <label for="buyAmount">Amount</label>
                            <input id="buyAmount" type="text" v-model="buyAmount">
                            </div>

                            <div class="button-holder form-group">
                                <div v-on:click="buyPercent(25)" class="btn btn-dark">25%</div>
                                <div v-on:click="buyPercent(50)" class="btn btn-dark">50%</div>
                                <div v-on:click="buyPercent(75)" class="btn btn-dark">75%</div>
                                <div v-on:click="buyPercent(100)" class="btn btn-dark">100%</div>
                            </div>
                            <div class="form-group">
                                <label for="buyTotal">Total: £</label>
                                <input disabled id="buyTotal" type="text" v-model="computedBuyTotal">
                            </div>
                            <div class="form-group--buttons">
                                <div class="btn btn-dark" v-on:click="closeTradePane()">Cancel</div>
                                <div class="btn btn-primary" v-on:click="buyPost()">Buy</div>
                            </div>
                        </form>
                        <form v-if="showSell" class="col sell-form">
                            <div class="form-group d-flex">
                                <h4 class="h4">Sell</h4>
                            </div>
                            <div class="form-group">
                                <label for="sellPrice">Price: £</label>
                                <input id="sellPrice" type="text" v-model="sellPrice" :placeholder="(this.data['current-price'] / 100)">
                            </div>
                            <div class="form-group">
                            <label for="sellAmount">Amount</label>
                            <input id="sellAmount" type="text" v-model="sellAmount">
                            </div>

                            <div class="button-holder form-group">
                                <div v-on:click="sellPercent(25)" class="btn btn-dark">25%</div>
                                <div v-on:click="sellPercent(50)" class="btn btn-dark">50%</div>
                                <div v-on:click="sellPercent(75)" class="btn btn-dark">75%</div>
                                <div v-on:click="sellPercent(100)" class="btn btn-dark">100%</div>
                            </div>
                            <div class="form-group">
                                <label for="sellTotal">Total: £</label>
                                <input disabled id="sellTotal" type="text" v-model="computedSellTotal">
                            </div>
                            <div class="form-group--buttons">
                                <div class="btn btn-dark" v-on:click="closeTradePane()">Cancel</div>
                                <div class="btn btn-primary" v-on:click="sellPost">Sell</div>

                            </div>
                        </form>
                </div>
                </div>
</template>

<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';
    @Component
  export default class TradesPage extends Vue {
        isin:any;
        data:any = [];
        currentHoldings:number = 0;
        buyAmount:number = 0;
        buyPrice:number = 0;
        sellAmount:number = 0;
        sellPrice:number = 0;
        bodyData:any;
        showBuy:boolean = false;
        showSell:boolean = false;

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
        buyPost () {
            this.bodyData = {fundName: this.isin, amount: parseInt(this.computedBuyTotal), status: 1};
            fetch('http://localhost:8080/portfolios/' + this.$store.getters['returnUsername'] + '/accounts/' + this.$store.getters['returnPortfolioName'] + '/trades', {
                method: 'POST',
                body: JSON.stringify(this.bodyData),
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: 'include'

            }).then(res => {
                if (res.status === 200) {
                    return 'Success';
                } else if (res.status === 401) {
                    throw new Error('username or password incorrect');
                }
            })
                .then((res) => {
                    console.log(res);
                    this.$emit('get-trade-history');
                    this.closeTradePane();
                })
                .catch((error) => {
                    console.log(error);
                });
        }
        sellPost () {
            this.bodyData = {fundName: this.isin, amount: -1 * parseInt(this.computedSellTotal), status: 1};
            fetch('http://localhost:8080/portfolios/' + this.$store.getters['returnUsername'] + '/accounts/' + this.$store.getters['returnPortfolioName'] + '/trades', {
                method: 'POST',
                body: JSON.stringify(this.bodyData),
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: 'include'

            }).then(res => {
                if (res.status === 200) {
                    return 'Success';
                } else if (res.status === 401) {
                    throw new Error('Cannot connect at this time. Please try again later');
                }
            })
                .then((res) => {
                    this.$emit('get-trade-history');
                    this.closeTradePane();
                    console.log(res);
                })
                .catch((error) => {
                    console.log(error);
                });
        }

        showBuyPane () {
            this.showBuy = true;
            this.showSell = false;
        }
        showSellPane () {
            this.showBuy = false;
            this.showSell = true;
        }
        closeTradePane () {
            this.showBuy = false;
            this.showSell = false;
            this.buyAmount = 0;
            this.buyPrice = 0;
            this.sellAmount = 0;
            this.sellPrice = 0;
        }

        // computed
        get computedBuyTotal () {
            return (this.buyAmount * this.buyPrice).toFixed(3);
        }
        get computedSellTotal () {
            return (this.sellAmount * this.sellPrice).toFixed(3);
        }
  };
</script>

<style lang="scss" scoped>
    .trade-page {
        position: relative;
        background: white;
        padding: 10px 0;
        margin:0;
        width:100%;
        -webkit-box-shadow: 1px 2px 3px 2px rgba(184,184,184,0.66);
        -moz-box-shadow: 1px 2px 3px 2px rgba(184,184,184,0.66);
        box-shadow: 1px 2px 3px 2px rgba(184,184,184,0.66);
        display: flex;
        align-items: center;
        align-content: center;
        justify-content: center;
    }
    .toggle-container {
        display: flex;
        justify-content: center;

        .btn {
            margin: 0 10px;
            width: 80px;
            height: 46px;
            display: flex;
            align-items: center;
            justify-content: center;
        }
    }
    .form-group {
        width:100%;
        display: flex;
        justify-content: space-around;
        align-items: center;
    }
    label {
        margin-bottom: 0;
    }
    input {
        justify-content: stretch;
    }
    .form-group--buttons {
        @extend .form-group;
        justify-content: space-around;

        .btn {
            margin:0 10px;
        }
    }
    .btn {
        width: 50px;
        padding: 5px 0px;
    }
    @media only screen and (min-width: 768px) {
        .trade-page {
            margin: 0 auto;
            -webkit-box-shadow: 0 0 1px 1px rgba(184,184,184,0.66);
            -moz-box-shadow: 0 0 1px 1px rgba(184,184,184,0.66);
            box-shadow: 0 0 1px 1px rgba(184,184,184,0.66);
        }
    }
</style>
