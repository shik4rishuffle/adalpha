<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';
  import { Doughnut, mixins } from 'vue-chartjs';
  const {reactiveData} = mixins;

    @Component({
        mixins: [Doughnut, reactiveData]
  })
  export default class PieChart extends Vue {
        chartData:any;
        chartOptions:any;
        incData:any = 2;
        mounted () {
            this.setData();
        }

        setData () {
            let dataArray: any = [];
            let labelsArray: any = [];
            for (let holding of this.$store.getters['returnPortfolioHoldings']) {
                dataArray.push(holding.currentValue.toFixed(2));
                labelsArray.push(holding.isin);
            }
            this.chartData = {
                datasets: [{
                    data: dataArray
                }],
                labels: labelsArray
            };
            this.chartOptions = {
                intersect: false,
                animation: {
                    duration: 3000
                }
            };
            // @ts-ignore
            this.renderChart(this.chartData, this.chartOptions);
        }
  }
</script>
