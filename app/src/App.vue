<template>
  <theme-layout>
    <router-view />
  </theme-layout>
</template>

<script>
  import ThemeLayout from '@/components/Theme/ThemeLayout';

  export default {
    components: { ThemeLayout }
  };
</script>

<style lang="stylus">
  @import url('https://fonts.googleapis.com/css?family=Work+Sans')

  body
    margin: 0
    font-family: 'Work Sans', sans-serif
</style>
