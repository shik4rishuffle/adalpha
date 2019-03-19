<template>
    <div class="the-menu">
        <div v-on:click="show = true" class="burger-icon">Menu</div>
        <div v-on:click="show = false" v-if="show" class="menu-takeover">
            <ul class="card col-4">
                <router-link to="/DashBoard"><li class="menu-item">Dashboard</li></router-link>
                <router-link to="/TradesPage"><li class="menu-item">Trades</li></router-link>
                <li v-on:click="logout" class="menu-item">Logout</li>
            </ul>
        </div>
    </div>
</template>

<script lang="ts">
  import {Vue, Component, Prop} from 'vue-property-decorator';

  @Component
  export default class TheMenu extends Vue {
    @Prop({
        default: false
    })
    show: Boolean;

  // Methods
      logout () {
          this.$store.dispatch('logOutHandler');
          this.$cookies.remove(this.$store.getters['returnCookie']);
          this.$router.push('/LogIn');
      }
  };
</script>

<style scoped>
    .menu-takeover {
        position: absolute;
        top: 0;
        left: 0;
        width: 100vw;
        height: 100vh;
        background: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .menu-item {
        list-style-type: none;
        padding: 10px;
        text-align: center;
    }

    .burger-icon {
        position: absolute;
        top: 10px;
        right: 20px;
    }

</style>
