<template>
    <div class="dash-board container-fluid">
        <div class="row">
            <div class="col">
                <h1 class="h1">Welcome {{username}}</h1>
            </div>
        </div>
        <div class="row graph-row">
            <div class="col-8 graph-col">
                <div class="col-11 graph">
                    <investment-graph></investment-graph>
                </div>
                <div class="col-1">
                    <div>Toggle</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
  import InvestmentGraph from '../components/InvestmentGraph';

  export default {
    name: 'DashBoard',
    components: {InvestmentGraph},
    data () {
      return {
        username: 'Flynn'
      };
    }
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
