<template>
    <div id="app">
        <router-view/>
        <!--@ToDo: take this off the login screen-->
        <the-menu></the-menu>
    </div>
</template>

<script lang="ts">
  import TheMenu from './components/TheMenu';
  import {Vue, Component} from 'vue-property-decorator';

  @Component({
      components: {TheMenu}
  })
  export default class App extends Vue {
    name: 'App';
  };
</script>
