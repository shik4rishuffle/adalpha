import VueRouter, { Route } from 'vue-router'
import {VueCookies} from "vue-cookies";
import Vue from 'vue';

// Fix 'cannot find module ./App'
declare module '*.vue' {
    export default Vue;
}
declare module 'vue/types/vue' {
    interface Vue {
        $router: VueRouter,
        $cookies: VueCookies
    }
}