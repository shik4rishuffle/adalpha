<template>
    <div class="dash-board container-fluid">
        <div class="row">
            <div class="col">
                <h3 class="h3">{{this.$store.getters['returnPortfolioName']}} - £{{this.$store.getters['returnPortfolioTotal'].toFixed(2)}}</h3>
            </div>
        </div>
        <div class="row graph-row">
            <div class="col-6 graph-col">
                <div class="col-12 graph">
                    <the-portfolio></the-portfolio>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
  import ThePortfolio from '../components/ThePortfolio';
  import {Vue, Component} from 'vue-property-decorator';

  @Component({
      components: {ThePortfolio}
  })

  export default class DashBoard extends Vue {
    name: 'DashBoard';
  };
</script>

<style scoped>
    .graph-row {
        justify-content: center;
        align-items: center;
    }

    .graph-col {
        display: flex;
        flex-direction: row;
        text-align: center;
    }
</style>
