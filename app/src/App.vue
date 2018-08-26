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
  @import url('https://fonts.googleapis.com/css?family=Comfortaa:300,400,700')

  *
  *::after
  *::before
    box-sizing: border-box

  body
    margin: 0
    font-family: 'Comfortaa', sans-serif
    background-color: #F1F1F1
</style>
