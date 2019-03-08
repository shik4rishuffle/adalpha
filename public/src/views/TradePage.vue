<template>
    <div class="trade-page container-fluid">
        <div class="row">
            <div class="col-12">
                <investment-graph></investment-graph>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1>{{companyName}}</h1>
                        <h1 class="h1">£{{currentHoldings}}</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col trade-buttons">
                        <button class="btn-primary btn">Buy</button>
                    </div>
                    <div class="col trade-buttons">
                        <button class="btn-primary btn">Sell</button>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <table class="table table-hover">
                            <thead>
                                <tr>
                                    <th scope="col">Date</th>
                                    <th scope="col">Buy/Sell</th>
                                    <th scope="col">Amount</th>
                                    <th scope="col">Processed</th>
                                </tr>
                            </thead>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
  import InvestmentGraph from '../components/InvestmentGraph';

  export default {
    name: 'TradePage',
    components: {InvestmentGraph},
    data () {
      return {
        companyName: 'test',
        currentHoldings: 33
      };
    }
  };
</script>

<style scoped>

</style>
