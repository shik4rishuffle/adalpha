import Vue from 'vue';
import App from './App';
import router from './router';
import 'bootstrap/dist/css/bootstrap.min.css';
import {store} from './store/store';
import Vuex from 'vuex';
import VueCookies from 'vue-cookies';

Vue.use(Vuex);
Vue.use(VueCookies);

Vue.config.productionTip = false;
const VueApp: any = Vue;
/* eslint-disable no-new */
const MyApp = new VueApp({
  router,
  store: store,
  el: '#app',
  beforeMount: () => {
    const cookieName = store.getters['returnCookie'];
    if(Vue.cookies.isKey(cookieName)){
      console.log(Vue.cookies.get(cookieName));
      const cookieValue = Vue.cookies.get(cookieName);
      store.dispatch('logInHandler', cookieValue);
    }
  },
  components: {App},
  template: '<App/>'
});
