<template>
    <div class="the-menu">
        <div v-on:click="show = true" class="burger-icon">Menu</div>
        <div v-on:click="show = false" v-if="show" class="menu-takeover">
            <ul class="card col-8 col-md-6">
                <router-link to="/DashBoard"><li class="menu-item">Dashboard</li></router-link>
                <router-link to="/TradesPage"><li class="menu-item">Trades</li></router-link>
                <li v-on:click="logout" class="menu-item--logout">Logout</li>
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

<style lang="scss" scoped>
    .menu-takeover {
        position: fixed;
        top: 0;
        left: 0;
        width: 100vw;
        height: 100%;
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
    .menu-item--logout {
        @extend .menu-item;
        color: darkred;
    }

    .burger-icon {
        position: absolute;
        top: 10px;
        right: 20px;
    }

</style>
