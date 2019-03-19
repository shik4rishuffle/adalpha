<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';
  import { Doughnut, mixins } from 'vue-chartjs';
  const {reactiveData} = mixins;

    @Component({
        mixins: [Doughnut, reactiveData]
  })
  export default class PieChart extends Vue {
        chartData:any;
        options:any;
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
                labels: labelsArray,
                centerText: {
                    display: true,
                    text: '20'
                }
            };
            this.options = {
                maintainAspectRatio: false,
                intersect: false,
                animation: {
                    duration: 1000
                },
                legend: {
                    display: false
                }
            };
            // @ts-ignore
            this.renderChart(this.chartData, this.options);
        }
  }
</script>
