<template>
  <div class="ProjectCard">
    <user-pill
      name="Vue ?"
      avatar="https://madewithvuejs.com/mandant/madewithvuejs/images/logo.png"
    />
    <theme-progress :value="95" />
  </div>
</template>

<script>
  import UserPill from '@/components/User/UserPill';
  import ThemeProgress from '@/components/Theme/ThemeProgress';

  export default {
    components: { UserPill, ThemeProgress }
  };
</script>

