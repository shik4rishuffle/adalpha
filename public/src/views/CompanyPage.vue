<template>
    <div class="company-page container-fluid">
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

                <div class="row">
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
                                <th scope="col">{{trade.date}}</th>
                                <th scope="col">£{{(trade.price / 100).toFixed(2)}}</th>
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

<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';
  @Component
  export default class CompanyPage extends Vue {
      isin:any;
      data:any = [];
      currentHoldings:number = 0;

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
  };
</script>

<style scoped>

</style>
