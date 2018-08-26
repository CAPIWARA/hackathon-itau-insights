<template>
  <section class="SupporterScreen">
    <h1>Apoiador</h1>

    <project-card
      name="Afro Bella"
      :goal="30000"
      :collected="20400"
      :banner="Banner"
    />
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

