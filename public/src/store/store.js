import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export const store = new Vuex.Store({
  state: {
    loggedIn: false,
    username: '',
    cookieName: 'ADALPHA-TEST-AUTH',
    portfolioName: '',
    portfolioHoldings: [],
    portfolioTrades: false,
    portfolioTotal: 0,
    portfolioLoaded: false,
    currentPage: ''
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
    },
    returnPortfolioName: state => {
      return state.portfolioName;
    },
    returnPortfolioHoldings: state => {
      return state.portfolioHoldings;
    },
    returnPortfolioTrades: state => {
      return state.portfolioTrades;
    },
    returnPortfolioTotal: state => {
      return state.portfolioTotal;
    },
    returnPortfolioLoaded: state => {
      return state.portfolioLoaded;
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
    },
    portfolioNameHandler: (state, payload) => {
      state.portfolioName = payload;
    },
    portfolioHoldingsHandler: (state, payload) => {
      state.portfolioHoldings = payload;
    },
    portfolioTradesHandler: (state, payload) => {
      state.portfolioTrades = payload;
    },
    portfolioTotalHandler: (state, payload) => {
      state.portfolioTotal += payload;
    },
    portfolioLoadedHandler: (state, payload) => {
      state.portfolioLoaded = payload;
    }
  },

  actions: {
    logInHandler: (context, payload) => {
      context.commit('logInHandler', payload);
    },
    logOutHandler: (context) => {
      context.commit('logOutHandler');
    },
    portfolioNameHandler: (context, payload) => {
      context.commit('portfolioNameHandler', payload);
    },
    portfolioHoldingsHandler: (context, payload) => {
      context.commit('portfolioHoldingsHandler', payload);
    },
    portfolioTradesHandler: (context, payload) => {
      context.commit('portfolioTradesHandler', payload);
    },
    portfolioTotalHandler: (context, payload) => {
      context.commit('portfolioTotalHandler', payload);
    },
    portfolioLoadedHandler: (context, payload) => {
      context.commit('portfolioLoadedHandler', payload);
    }
  }

});
