import Vue from 'vue';
import Router from 'vue-router';
import {store} from '../store/store';

import LogIn from '@/views/LogIn';
import CompanyPage from '@/views/CompanyPage';
import DashBoard from '@/views/DashBoard';
import TradePage from '@/views/TradePage';

Vue.use(Router);
const router = new Router({
  routes: [
    {
      path: '/',
      redirect: '/DashBoard'
    },
    {
      path: '/LogIn',
      name: 'LogIn',
      component: LogIn,
      meta: {
        noRequireAuth: true
      }
    },
    {
      path: '/CompanyPage/:isin',
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
  ],
  mode: 'history'
});
router.beforeEach((to, from, next) => {
  if (!to.matched.some(record => record.meta.noRequireAuth)) {
    if (!store.getters.returnLoginAuth) {
      next({ name: 'LogIn' });
    } else {
      next();
    }
  } else {
    next();
  }
});
router.afterEach((to) => {
  store.dispatch('currentPageHandler', to.name);
});
export default router;
