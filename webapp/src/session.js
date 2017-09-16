import Vue from 'vue'
import Session from './components/Session'
import router from './router/SessionRouter'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#session',
  router,
  template: '<session/>',
  components: { Session }
})
