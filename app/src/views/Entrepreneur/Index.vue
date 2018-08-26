<template>
  <section class="EntrepreneurScreen">
    <user-pill
      name="Made with Vue.js"
      avatar="https://madewithvuejs.com/mandant/madewithvuejs/images/logo.png"
    />

    <h2 class="subtitle">Meu projeto</h2>

    <project-card
      name="Afro Bella"
      :goal="30000"
      :collected="20400"
      :banner="Banner"
    />
  </section>
</template>

<script>
  import UserPill from '@/components/User/UserPill';
  import Banner from '@/assets/images/Banner.jpeg';

  export default {
    data: () => ({ Banner }),
    components: { UserPill }
  };
</script>

<style lang="stylus">
  .EntrepreneurScreen
    padding: 15px
    width: 100%
</style>
