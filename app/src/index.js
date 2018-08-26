import Vue from 'vue';
import App from '@/App';
import router from '@/router';

Vue.filter('toMoney', (value) => {
  const formatter = Intl.NumberFormat('pt-BR', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 3
  });

  return 'R$ ' + formatter.format(value);
});

export default new Vue({
  el: '#app',
  router,
  render: (λ) => λ(App)
});
