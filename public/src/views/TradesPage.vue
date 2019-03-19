<template>
    <div class="company-page">
                        <h3 class="company-page-title">All Trades</h3>
                        <div class="table-responsive">
                            <table class="table table-hover table-striped">
                                <thead>
                                <tr>
                                    <th>Buy/Sell</th>
                                    <th>Fund Name</th>
                                    <th>Amount</th>
                                    <th>Status</th>
                                </tr>
                                </thead>
                                <tbody>
                                <router-link v-for="trade in this.trades"
                                    :key="trade.fundName"
                                    :class="'trade-row-' + trade.type"
                                    v-if="(trade.status >0 && trade.status < 5 && trade.amount !== 0)"
                                    tag="tr"
                                    :to="{name: 'CompanyPage', params: {isin: trade.fundName}}"
                                >
                                    <td >{{trade.type}}</td>
                                    <td>{{trade.fundName}}</td>
                                    <td>{{trade.amount}}</td>
                                    <td>{{getTradeStatus(trade.status)}}</td>
                                </router-link>
                                </tbody>
                            </table>
                        </div>
                    </div>

</template>

<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';
    @Component
  export default class TradesPage extends Vue {
        isin:string;
        data:any = [];
        currentHoldings:number = 0;
        trades:any = [];

        created () {
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
                    let transName: string;
                    let transAmount: number;
                    let transStatus: number;
                    let transType: string;
                    let transArray:any = [];
                    res.map(function (item) {
                        transName = item.fundName;
                        transAmount = item.amount;
                        transStatus = item.status;
                        if (transAmount > 0) {
                            transType = 'buy';
                        } else if (transAmount < 0) {
                            transType = 'sell';
                        };
                        transArray.push({'fundName': transName, 'amount': transAmount, 'status': transStatus, 'type': transType});
                    });
                    this.trades = transArray;
                })
                .catch((error) => {
                    console.log(error);
                });
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
    .company-page {
        margin:0;
        padding:0;
    }
    .company-page-title {
        margin-top:30px;
        text-align: center;
    }
    .table {
        width:100vw;
        padding: 0;
        margin: 10px 0;
        text-align: center;
    tr {
        text-align: center;
    }
    th, td {
        padding: 5px 3px;
        text-align: center;
        vertical-align: middle;
        font-size: 0.99em;
    }
    }
</style>
