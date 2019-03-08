import Vue from 'vue';
import App from './App';
import router from './router';
import 'bootstrap/dist/css/bootstrap.min.css';

Vue.config.productionTip = false;

const VueApp: any = Vue;
/* eslint-disable no-new */
const MyApp = new VueApp({
  el: '#app',
  router,
  components: {App},
  template: '<App/>'
});
