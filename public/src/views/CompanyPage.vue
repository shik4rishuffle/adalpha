<template>
    <div class="company-page container-fluid">
        <div class="row">
            <div class="col-12">
                <investment-graph></investment-graph>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <div class="row">
                    <div class="col flex-row">
                        <h1 class="h1 company-name">{{companyName}}</h1>
                        <h1 class="h1">£{{currentHoldings}}</h1>
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

<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';

    @Component
  export default class CompanyPage extends Vue {
    name: string;
    companyName: string;
    currentHoldings: number;
  };
</script>

<style scoped>

</style>
