<template>
  <figure class="UserPill">
    <img :src="avatar" class="avatar" />
    <p class="name">{{ name }}</p>
  </figure>
</template>

<script>
  export default {
    props: {
      avatar: String,
      name: {
        type: String,
        required: true
      }
    }
  };
</script>

<style lang="stylus">
  .UserPill
    margin: 0
    display: flex
    align-items: center
    flex-direction: column
    justify-content: center

    > .avatar
      width: 68px
      margin: 8px
      height: 68px
      border: 1px solid #000
      border-radius: 50%
      background-color: #FFF

    > .name
      margin: 0
      font-size: 16pt
      font-weight: bold
</style>
