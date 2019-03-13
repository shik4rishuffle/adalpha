import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export const store = new Vuex.Store({
  state: {
    loggedIn: false,
    username: '',
    cookieName: 'ADALPHA-TEST-AUTH'
  },
  getters: {
    returnLoginAuth: state => {
      return state.loggedIn;
    },

    returnCookie: state => {
      return state.cookieName;
    },

    returnUsername: state => {
      return state.username;
    }
  },

  mutations: {
    logInHandler: (state, payload) => {
      state.loggedIn = true;
      state.username = payload;
    },
    logOutHandler: (state) => {
      state.loggedIn = false;
      state.username = '';
    }
  },

  actions: {
    logInHandler: (context, payload) => {
      context.commit('logInHandler', payload);
    },
    logOutHandler: (context) => {
      context.commit('logOutHandler');
    }
  }

});
