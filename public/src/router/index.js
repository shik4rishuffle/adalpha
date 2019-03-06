import Vue from 'vue';
import Router from 'vue-router';
import LogIn from '@/views/LogIn';
import CompanyPage from '@/views/CompanyPage';
import DashBoard from '@/views/DashBoard';
import TradePage from '@/views/TradePage';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/LogIn'
    },
    {
      path: '/LogIn',
      name: 'LogIn',
      component: LogIn
    },
    {
      path: '/CompanyPage',
      name: 'CompanyPage',
      component: CompanyPage
    },
    {
      path: '/DashBoard',
      name: 'DashBoard',
      component: DashBoard
    },
    {
      path: '/TradePage',
      name: 'TradePage',
      component: TradePage
    }
  ]
});
