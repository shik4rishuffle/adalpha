<template>
    <div id="app">
        <router-view/>
        <!--@ToDo: take this off the login screen-->
        <TheMenu></TheMenu>
    </div>
</template>

<script>
  import TheMenu from './components/TheMenu';

  export default {
    name: 'App',
    components: {TheMenu}
  };
</script>
