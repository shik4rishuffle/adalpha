<template>
    <div class="dash-board">
                    <div class="dash-board__portfolio-chart" v-if="this.$store.getters['returnPortfolioLoaded']">
                        <pie-chart></pie-chart>
                        <h3 class="dash-board__portfolio-total">£{{this.$store.getters['returnPortfolioTotal'].toFixed(2)}}</h3>
                    </div>
                    <div class="dash-board__portfolio-table col-lg-7">
                        <the-portfolio></the-portfolio>
                    </div>
            </div>
</template>

<script lang="ts">
  import ThePortfolio from '../components/ThePortfolio';
  import {Vue, Component} from 'vue-property-decorator';
  import PieChart from '../components/PieChart.vue';

  @Component({
      components: {PieChart, ThePortfolio}
  })

  export default class DashBoard extends Vue {
    name: 'DashBoard';
  };
</script>

<style lang="scss" scoped>
    .dash-board {
        width: 100vw;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
.dash-board__portfolio-chart {
    width: 100vw;
    flex-direction: column;
    align-items: center;
}
.dash-board__portfolio-total {
    text-align: center;
    margin-bottom: 15px;
    color: darkslategray;
}
.dash-board__portfolio-table {
    display: flex;
    width: 100%;
}
    @media only screen and (min-width: 768px) {
        .dash-board__portfolio-chart {
            margin-top: 60px;
            margin-bottom: 30px;
        }
        }
</style>
