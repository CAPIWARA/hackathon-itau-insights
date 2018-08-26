<template>
  <section class="EntrepreneurScreen">
    <div class="inner">
      <user-pill
        name="Made with Vue.js"
        avatar="https://madewithvuejs.com/mandant/madewithvuejs/images/logo.png"
      />

      <h2 class="subtitle">Meu projeto</h2>

      <project-card
        name="Afro Bella"
        :goal="30000"
        :collected="20400"
        :banner="BannerImage"
      />

      <h2 class="subtitle">Destaques</h2>
    </div>

    <carousel :pagination-enabled="false">
      <slide v-for="_ in projects" :key="_.name">
        <theme-card
          class="ProjectItem"
          :style="{ 'background-image': 'url(' + _.banner + ')' }"
        >
          <div class="description">
            <h4 class="name">{{ _.name }}</h4>
            <p class="author">por {{ _.author }}</p>
            <theme-progress :value="_.progress" />
          </div>
        </theme-card>
      </slide>
    </carousel>
  </section>
</template>

<script>
  import UserPill from '@/components/User/UserPill';
  import ProjectCard from '@/components/Project/ProjectCard';
  import ThemeCard from '@/components/Theme/ThemeCard';
  import ThemeProgress from '@/components/Theme/ThemeProgress';
  import BannerImage from '@/assets/images/Banner.jpeg';
  import { Carousel, Slide } from 'vue-carousel';

  const projects = [
    {
      banner: 'https://sunflourbakingcompany.com/wp-content/uploads/2017/03/sunflour-cupcakes-3.jpg',
      name: 'Cupcakes da Jú',
      author: 'Juliana Sampaio',
      progress: 78
    },
    {
      banner: 'http://iphotochannel.com.br/wp-content/uploads/2014/08/shutterstock_1526248791.jpg',
      name: 'Modernização das ferramentas',
      author: 'Rafael Pereira Lopes',
      progress: 43
    },
    {
      banner: 'http://iphotochannel.com.br/wp-content/uploads/2014/08/shutterstock_1526248791.jpg',
      name: 'Modernização das ',
      author: 'Rafael Pereira Lopes',
      progress: 43
    }
  ];

  export default {
    data: () => ({ projects, BannerImage }),
    components: {
      Carousel, Slide, UserPill,
      ProjectCard, ThemeCard, ThemeProgress
    }
  };
</script>

<style lang="stylus">
  .EntrepreneurScreen
    width: 100%
    padding-bottom: 30px

    > .inner
      padding: 15px

  .ProjectItem
    display: flex
    flex-direction: column
    width: 210px
    height: 145px
    margin-right: 20px
    overflow: hidden
    background-size: cover
    background-position: center

    > .description
      margin-top: auto
      background-color: rgba(#FFF, .9)
      width: 100%
      height: 90px
      border-bottom-left-radius: 12px
      border-bottom-right-radius: 12px
      padding: 10px 15px 15px

      > .name
        margin: 0
        font-size: 12pt

      > .author
        margin: 0
        font-size: 10pt
</style>
