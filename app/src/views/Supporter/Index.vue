<template>
  <section class="SupporterScreen">
    <h1>Apoiador</h1>
  </section>
</template>

<script>
  import ProjectCard from '@/components/Project/ProjectCard';
  import Banner from '@/assets/images/Banner.jpeg';

  export default {
    data: () => ({ Banner }),
    components: { ProjectCard }
  };
</script>

<style lang="stylus">
  .SupporterScreen
    padding: 15px
</style>

