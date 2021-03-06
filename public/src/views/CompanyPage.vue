<template>
    <div class="company-page">
        <div class="company-page__history-graph">
            <history-graph :chart-data="datacollection" :options="options"></history-graph>
            <h4 class="company-page__name">{{this.data.name}}</h4>
            <div class="company-page__details">
                <div class="company-page__detail">ISIN: {{ this.data.isin }}</div>
                <div class="company-page__detail">Current Price: £{{(this.data['current-price'] / 100).toFixed(2)}}</div>
                <div class="company-page__detail">Current Holdings: {{this.currentHoldings.toFixed(2)}}</div>
            </div>
        </div>
        <div class="container-fluid">
        <div class="row">
            <div class="col">

                <div class="row">
                    <div class="col-12">
                        <the-trader
                        v-on:get-trade-history="getTradeHistory()"
                        ></the-trader>
                    </div>
                </div>
                 <table class="table table table-striped">
                            <thead>
                            <tr>
                                <th scope="col">Buy/Sell</th>
                                <th scope="col">Fund Name</th>
                                <th scope="col">Amount</th>
                                <th scope="col">Status</th>
                            </tr>
                            </thead>
                <tbody>
                     <tr v-for="trade in this.trades"
                                :key="trade.name"
                                :class="trade.type"
                                v-if="(trade.status >0 && trade.status < 5)"
                            >
                                <td scope="col">{{trade.type}}</td>
                                <td scope="col">{{trade.fundName}}</td>
                                <td scope="col">{{Math.abs(trade.amount)}}</td>
                                <td scope="col">{{getTradeStatus(trade.status)}}</td>
                            </tr>
                </tbody>
                        </table>
            </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';
  import TheTrader from '../components/TheTrader.vue';
  import HistoryGraph from '../components/HistoryGraph.vue';
  @Component({
      components: {HistoryGraph, TheTrader}
  })
  export default class CompanyPage extends Vue {
      isin:string;
      data:any = [];
      currentHoldings:number = 0;
      trades:any = [];
      datacollection: any = {};
      options: any = {};

      created () {
          this.isin = this.$route.params.isin;
          this.getPriceHistory();
          this.getTradeHistory();
      }

      getTradeHistory () {
          fetch('http://localhost:8080/portfolios/' + this.$store.getters['returnUsername'] + '/accounts/' + this.$store.getters['returnPortfolioName'] + '/trades', {
              method: 'GET',
              headers: {
                  'Content-Type': 'application/json'
              },
              credentials: 'include'

          }).then(response => {
              if (response.status === 200) {
                  return response.json();
              } else if (response.status === 401) {
                  throw new Error('Cannot connect at this time. Please try again later');
              }
          })
              .then((res) => {
                  const fundName = this.isin;
                  let transName: string;
                  let transAmount: number;
                  let transStatus: number;
                  let transType: string;
                  let transArray:any = [];
                  res.filter(function (item) {
                      return (item.fundName === fundName);
                  }).map(function (item) {
                      transName = item.fundName;
                      transAmount = item.amount;
                      transStatus = item.status;
                        if (transAmount > 0) {
                            transType = 'buy';
                        } else if (transAmount < 0) {
                            transType = 'sell';
                        }
                      transArray.push({'fundName': transName, 'amount': transAmount, 'status': transStatus, 'type': transType});
                  });
                  this.trades = transArray;
              })
              .catch((error) => {
                  console.log(error);
              });
      }

      getPriceHistory () {
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
                  this.data = res;

                  // Build Graph data and pass through to HistoryGraph.vue component
                  let dataArray = [];
                  let labelsArray = [];
                  for (let result of this.data.history) {
                      dataArray.push(result.price);
                      labelsArray.push(result.date);
                  }
                  this.datacollection = dataArray;
                  this.datacollection = {
                      labels: labelsArray,
                      datasets: [{
                          data: dataArray,
                          lineTension: 0
                      }]
                  };

                  this.getCurrentHoldings();
              })
              .catch((error) => {
                  console.log(error);
              });
      }

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

      getTradeStatus (status) {
          switch (status) {
              case 1:
                  return 'Trade Pending';
              case 2:
                  return 'Trade Cancelled';
              case 3:
                  return 'Trade Contracted';
              case 4:
                  return 'Trade Settled';
          }
      }
  };
</script>
<style lang="scss" scoped>
    .company-page__history-graph {
        position: relative;
        width:100vw;
    }
    .company-page__name {
        position: absolute;
        top:0;
        padding: 0 10px;
        margin-top: 60px;
        margin-bottom: 10px;
        text-align: center;
    }
    .company-page__details {
        position: absolute;
        bottom: 20px;
        padding:0 10px;
    }
    .company-page__detail {
        margin-top: 5px;
        font-size: 0.8em;
    }
    .table {
        width:100vw;
        padding: 0;
        margin: 20px -15px;
        text-align: center;
        tr {
            text-align: center;
        }
        th, td {
            padding: 0;
            text-align: center;
            vertical-align: middle;
        }
    }
    @media only screen and (min-width: 768px) {
        .company-page__detail {
            margin-top: 5px;
            font-size: 1em;
        }
    }
</style>
