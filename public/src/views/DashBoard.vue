<template>
    <div class="dash-board container-fluid">
        <div class="row">
            <div class="col">
                <h1 class="h1">Welcome {{this.$store.getters['returnUsername']}}</h1>
            </div>
        </div>
        <div class="row graph-row">
            <div class="col-8 graph-col">
                <div class="col-11 graph">
                    <InvestmentGraph></InvestmentGraph>
                </div>
                <div class="col-1">
                    <div>Toggle</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
  import InvestmentGraph from '../components/InvestmentGraph';
  import {Vue, Component} from 'vue-property-decorator';

  @Component({
      components: {InvestmentGraph}
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
