<template>
  <theme-card class="ProjectCard">
    <img
      :src="banner"
      :style="{ backgroundImage: 'url(' + banner + ')' }"
      class="banner"
    />
    <div class="inner">
      <h1 class="name">{{ name }}</h1>

      <p class="goal">
        Arrecadado
        <strong>{{ collected | toMoney }}</strong>
        <small>(meta: {{ goal | toMoney }})</small>
      </p>

      <theme-progress :value="progress" />

      <p class="details">mais detalhes</p>
    </div>
  </theme-card>
</template>

<script>
  import UserPill from '@/components/User/UserPill';
  import ThemeCard from '@/components/Theme/ThemeCard';
  import ThemeProgress from '@/components/Theme/ThemeProgress';

  export default {
    components: { UserPill, ThemeCard, ThemeProgress },
    props: {
      name: {
        type: String,
        required: true
      },
      goal: {
        type: Number,
        required: true
      },
      collected: {
        type: Number,
        required: true
      },
      banner: {
        type: String,
        required: true
      }
    },
    computed: {
      progress () {
        return (this.collected / this.goal * 100);
      }
    }
  };
</script>

<style lang="stylus">
  .ProjectCard
    overflow: hidden
    box-shadow: 0 0 10px rgba(#000, .3)

    > .banner
      min-width: 100%
      height: 80px
      padding: 50px 50%
      background-size: cover
      background-position: center


    > .inner
      padding: 15px

      > .name
        margin-top: 0

      > .details
        text-align: center
</style>

