<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';
  import { Line, mixins } from 'vue-chartjs';
  const {reactiveProp} = mixins;
    @Component({
        mixins: [Line, reactiveProp]
  })
  export default class HistoryGraph extends Vue {
        chartData:any;
        options:any = {
            maintainAspectRatio: false,
            legend: {
                display: false
            },
            scales: {
                yAxes: [{
                    ticks: {
                        display: false
                    },
                    gridLines: {
                        display: false
                    }
                }],
                xAxes: [{
                    ticks: {
                        display: false
                    },
                    gridLines: {
                        display: false
                    }
                }],
                labels: {
                    display: false
                }
            }
        };
        mounted () {
            // @ts-ignore
            this.renderChart(this.chartData, this.options);
        }
  }
</script>
