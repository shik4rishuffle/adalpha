import Vue from 'vue';
import Router from 'vue-router';
import {store} from '../store/store';

import LogIn from '@/views/LogIn';
import CompanyPage from '@/views/CompanyPage';
import DashBoard from '@/views/DashBoard';
import TradesPage from '@/views/TradesPage';

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
      path: '/TradesPage',
      name: 'TradesPage',
      component: TradesPage
    }
  ],
  mode: 'history'
});
router.beforeEach((to, from, next) => {
  const cookieName = store.getters['returnCookie'];
  if (!to.matched.some(record => record.meta.noRequireAuth) && !Vue.cookies.isKey(cookieName)) {
    if (!store.getters.returnLoginAuth) {
      next({ name: 'LogIn' });
    } else {
      next();
    }
  } else {
    next();
  }
});
router.afterEach((to, from) => {
  if (from.name !== LogIn) {
    localStorage.setItem('currentPage', to.name);
  }
});
export default router;
